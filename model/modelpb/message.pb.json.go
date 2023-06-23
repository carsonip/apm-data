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

package modelpb

import "github.com/elastic/apm-data/model/internal/modeljson"

func (m *Message) toModelJSON(out *modeljson.Message) {
	*out = modeljson.Message{
		Body:       m.Body,
		Age:        modeljson.MessageAge{Millis: m.AgeMillis},
		Queue:      modeljson.MessageQueue{Name: m.QueueName},
		RoutingKey: m.RoutingKey,
	}
	if n := len(m.Headers); n > 0 {
		headers := make(map[string][]string, n)
		for _, hv := range m.Headers {
			if hv != nil {
				headers[hv.Key] = hv.Value
			}
		}
		out.Headers = headers
	}
}
