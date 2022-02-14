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

package instrument // import "go.opentelemetry.io/otel/metric/instrument"

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/metric/unit"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name    string
		options []Option

		wantDescription string
		wantUnit        unit.Unit
	}{
		{
			name: "Default is empty",
		},
		{
			name: "Both options works",
			options: []Option{
				WithDescription("descriptive description"),
				WithUnit(unit.Milliseconds),
			},
			wantDescription: "descriptive description",
			wantUnit:        unit.Milliseconds,
		},
		{
			name: "Options take last value",
			options: []Option{
				WithDescription("descriptive description"),
				WithUnit(unit.Milliseconds),
				WithDescription("not very descriptive"),
				WithUnit(unit.Bytes),
			},
			wantDescription: "not very descriptive",
			wantUnit:        unit.Bytes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig(tt.options...)

			require.Equal(t, tt.wantDescription, config.Description())
			require.Equal(t, tt.wantUnit, config.Unit())
		})
	}
}
