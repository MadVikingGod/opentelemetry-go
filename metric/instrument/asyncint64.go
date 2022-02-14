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

// Int64ObserverCounter is an instrument that records increasing values.
type Int64ObserverCounter interface {
	// Observe records the state of the instrument.
	//
	// It is only valid to call this within a callback. If called outside of the
	// registered callback it should have no effect on the instrument, and an
	// error will be reported via the error handler.
	Observe(ctx context.Context, x int64, attrs ...attribute.KeyValue)

	Asynchronous
}

// Int64ObserverUpDownCounter is an instrument that records increasing or decresing values.
type Int64ObserverUpDownCounter interface {
	// Observe records the state of the instrument.
	//
	// It is only valid to call this within a callback. If called outside of the
	// registered callback it should have no effect on the instrument, and an
	// error will be reported via the error handler.
	Observe(ctx context.Context, x int64, attrs ...attribute.KeyValue)

	Asynchronous
}

// Int64ObserverGauge is an instrument that records independent readings.
type Int64ObserverGauge interface {
	// Observe records the state of the instrument.
	//
	// It is only valid to call this within a callback. If called outside of the
	// registered callback it should have no effect on the instrument, and an
	// error will be reported via the error handler.
	Observe(ctx context.Context, x int64, attrs ...attribute.KeyValue)

	Asynchronous
}
