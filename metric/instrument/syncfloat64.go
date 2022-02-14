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
	"context"

	"go.opentelemetry.io/otel/attribute"
)

// Float64Counter is an instrument that records increasing values.
type Float64Counter interface {
	// Add records a change to the counter.
	Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue)

	Synchronous
}

// Float64UpDownCounter is an instrument that records increasing or decresing values.
type Float64UpDownCounter interface {
	// Add records a change to the counter.
	Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue)

	Synchronous
}

// Float64Histogram is an instrument that records a distribution of values.
type Float64Histogram interface {
	// Record adds an additional value to the distribution.
	Record(ctx context.Context, incr float64, attrs ...attribute.KeyValue)

	Synchronous
}
