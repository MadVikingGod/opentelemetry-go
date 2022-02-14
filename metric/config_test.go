// Copyright The OpenTelemetry Authors
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

package metric // import "go.opentelemetry.io/otel/metric"

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name    string
		options []MeterOption

		wantInstrumentationVersion string
		wantSchemaURL              string
	}{
		{
			name: "Default is empty",
		},
		{
			name: "both options works",
			options: []MeterOption{
				WithInstrumentationVersion("iv1.0"),
				WithSchemaURL("schema2"),
			},
			wantInstrumentationVersion: "iv1.0",
			wantSchemaURL:              "schema2",
		},
		{
			name: "options take last value",
			options: []MeterOption{
				WithInstrumentationVersion("iv1.0"),
				WithSchemaURL("schema2"),
				WithInstrumentationVersion("iv2.0"),
				WithSchemaURL("schema4"),
			},
			wantInstrumentationVersion: "iv2.0",
			wantSchemaURL:              "schema4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewMeterConfig(tt.options...)

			require.Equal(t, tt.wantInstrumentationVersion, config.InstrumentationVersion())
			require.Equal(t, tt.wantSchemaURL, config.SchemaURL())
		})
	}
}
