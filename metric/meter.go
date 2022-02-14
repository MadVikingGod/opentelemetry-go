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
	"context"

	"go.opentelemetry.io/otel/metric/instrument"
)

// MeterProvider provides access to named Meter instances, for instrumenting
// an application or library.
type MeterProvider interface {
	// Meter creates an instance of a `Meter` interface. The instrumentationName
	// must be the name of the library providing instrumentation. This name may
	// be the same as the instrumented code only if that code provides built-in
	// instrumentation. If the instrumentationName is empty, then a
	// implementation defined default name will be used instead.
	Meter(instrumentationName string, opts ...MeterOption) Meter
}

// Meter provides access to instrument instances for recording metrics.
type Meter interface {
	// Float64 is the namespace for the Synchronous Float instruments
	Float64() Float64Instruments

	// Int64 is the namespace for the Synchronous Integer instruments
	Int64() Int64Instruments

	// Float64Observer is the namespace for the Asynchronous Float instruments
	//
	// To Observe data with instruments it must be registered in a callback.
	Float64Observer() Float64ObserverInstruments

	// Int64Observer is the namespace for the Asynchronous Integer instruments.
	//
	// To Observe data with instruments it must be registered in a callback.
	Int64Observer() Int64ObserverInstruments

	// RegisterCallback captures the function that will be called during Collect.
	//
	// It is only valid to call Observe within the scope of the passed function,
	// and only on the instruments that were registered with this call.
	RegisterCallback(insts []instrument.Asynchronous, function func(context.Context)) error
}

// Instruments provides access to individual instruments.
type Float64ObserverInstruments interface {
	// Counter creates an instrument for recording increasing values.
	Counter(name string, opts ...instrument.Option) (instrument.Float64ObserverCounter, error)

	// UpDownCounter creates an instrument for recording changes of a value.
	UpDownCounter(name string, opts ...instrument.Option) (instrument.Float64ObserverUpDownCounter, error)

	// Gauge creates an instrument for recording the current value.
	Gauge(name string, opts ...instrument.Option) (instrument.Float64ObserverGauge, error)
}

type Int64ObserverInstruments interface {
	// Counter creates an instrument for recording increasing values.
	Counter(name string, opts ...instrument.Option) (instrument.Int64ObserverCounter, error)

	// UpDownCounter creates an instrument for recording changes of a value.
	UpDownCounter(name string, opts ...instrument.Option) (instrument.Int64ObserverUpDownCounter, error)

	// Gauge creates an instrument for recording the current value.
	Gauge(name string, opts ...instrument.Option) (instrument.Int64ObserverGauge, error)
}

// Instruments provides access to individual instruments.
type Float64Instruments interface {
	// Counter creates an instrument for recording increasing values.
	Counter(name string, opts ...instrument.Option) (instrument.Float64Counter, error)
	// UpDownCounter creates an instrument for recording changes of a value.
	UpDownCounter(name string, opts ...instrument.Option) (instrument.Float64UpDownCounter, error)
	// Histogram creates an instrument for recording a distribution of values.
	Histogram(name string, opts ...instrument.Option) (instrument.Float64Histogram, error)
}

// Instruments provides access to individual instruments.
type Int64Instruments interface {
	// Counter creates an instrument for recording increasing values.
	Counter(name string, opts ...instrument.Option) (instrument.Int64Counter, error)
	// UpDownCounter creates an instrument for recording changes of a value.
	UpDownCounter(name string, opts ...instrument.Option) (instrument.Int64UpDownCounter, error)
	// Histogram creates an instrument for recording a distribution of values.
	Histogram(name string, opts ...instrument.Option) (instrument.Int64Histogram, error)
}
