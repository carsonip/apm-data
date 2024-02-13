// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package modelprocessor

import (
	"context"
	"encoding/binary"
	"fmt"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"strconv"
	"time"

	"github.com/elastic/apm-data/model/modelpb"
)

// RedirectLogs extracts all APMEvent logs from Batch, convert them to otel logs and send them to otelLogsProcessor,
// and remove them from the original batch.
type RedirectLogs struct {
	OTelLogsProcessor modelpb.OTelLogsProcessor
}

func (r RedirectLogs) ProcessBatch(ctx context.Context, batch *modelpb.Batch) error {
	// FIXME(carsonip): alloc
	b := make(modelpb.Batch, 0, len(*batch))
	for _, e := range *batch {
		if e.Type() != modelpb.LogEventType {
			b = append(b, e)
			continue
		}
		logs := plog.NewLogs()
		rl := logs.ResourceLogs().AppendEmpty()
		rAttrs := rl.Resource().Attributes()
		rAttrs.PutStr("service.name", e.GetService().GetName())
		rAttrs.PutStr("service.version", e.GetService().GetVersion())
		sl := rl.ScopeLogs().AppendEmpty()
		sAttrs := sl.Scope().Attributes()
		sl.Scope().SetName(e.GetService().GetFramework().GetName())
		sl.Scope().SetVersion(e.GetService().GetFramework().GetVersion())
		sAttrs.PutStr("log.logger", e.GetLog().GetLogger())
		lr := sl.LogRecords().AppendEmpty()
		lr.SetTimestamp(pcommon.Timestamp(e.GetTimestamp()))
		lr.SetObservedTimestamp(pcommon.Timestamp(time.Now().UnixNano()))
		lr.Body().SetStr(e.GetMessage())
		if spanID := e.GetSpan().GetId(); spanID != "" {
			if len(spanID) != 16 {
				return fmt.Errorf("spanID len != 16; spanID = %q", spanID)
			}
			i, err := strconv.ParseUint(spanID, 16, 64)
			if err != nil {
				return err
			}
			var bytes [8]byte
			binary.BigEndian.PutUint64(bytes[:], i)
			lr.SetSpanID(bytes)
		}
		if traceID := e.GetTrace().GetId(); traceID != "" {
			if len(traceID) != 32 {
				return fmt.Errorf("traceID len != 32; traceID = %q", traceID)
			}
			upper, err := strconv.ParseUint(traceID[:16], 16, 64)
			if err != nil {
				return err
			}
			lower, err := strconv.ParseUint(traceID[16:], 16, 64)
			if err != nil {
				return err
			}
			var bytes [16]byte
			binary.BigEndian.PutUint64(bytes[:8], upper)
			binary.BigEndian.PutUint64(bytes[8:], lower)
			lr.SetTraceID(bytes)
		}
		// FIXME(carsonip): there are many more other fields to be translated
		if err := r.OTelLogsProcessor.ProcessLogs(ctx, logs); err != nil {
			return err
		}
	}
	// Mutate batch to filter away logs
	*batch = b
	return nil
}
