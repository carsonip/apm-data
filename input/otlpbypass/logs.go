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

// Portions copied from OpenTelemetry Collector (contrib), from the
// elastic exporter.
//
// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otlpbypass

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/elastic/apm-data/input"
)

// ConsumeLogs calls ConsumeLogsWithResult but ignores the result.
// It exists to satisfy the go.opentelemetry.io/collector/consumer.Logs interface.
func (c *Consumer) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_, err := c.ConsumeLogsWithResult(ctx, logs)
	return err
}

// ConsumeLogsWithResult consumes OpenTelemetry log data, converting into
// the Elastic APM log model and sending to the reporter.
func (c *Consumer) ConsumeLogsWithResult(ctx context.Context, ld plog.Logs) (input.ConsumeLogsResult, error) {
	if err := semAcquire(ctx, c.sem, 1); err != nil {
		return input.ConsumeLogsResult{}, err
	}
	defer c.sem.Release(1)

	// set ObservedTimestamp
	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rl := rls.At(i)
		ills := rl.ScopeLogs()
		for j := 0; j < ills.Len(); j++ {
			logs := ills.At(j).LogRecords()
			for k := 0; k < logs.Len(); k++ {
				logs.At(k).SetObservedTimestamp(pcommon.NewTimestampFromTime(time.Now()))
			}
		}
	}

	if err := c.config.LogsProcessor.ProcessLogs(ctx, ld); err != nil {
		return input.ConsumeLogsResult{}, err
	}
	return input.ConsumeLogsResult{RejectedLogRecords: 0}, nil
}
