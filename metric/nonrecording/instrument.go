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

package nonrecording // import "go.opentelemetry.io/otel/metric/nonrecording"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
)

// NewNoopObservableFloat64Instruments returns a nonrecording implementation of metric.Float64ObserverInstruments
func NewNoopObservableFloat64Instruments() metric.ObservableFloat64Instruments {
	return nonRecordingObservableFloat64Instrument{}
}

var (
	_ instrument.ObservableFloat64Counter       = nonRecordingObservableFloat64Instrument{}
	_ instrument.ObservableFloat64UpDownCounter = nonRecordingObservableFloat64Instrument{}
	_ instrument.ObservableFloat64Gauge         = nonRecordingObservableFloat64Instrument{}
)

type nonRecordingObservableFloat64Instrument struct {
	instrument.Asynchronous
}

func (i nonRecordingObservableFloat64Instrument) Counter(name string, opts ...instrument.Option) (instrument.ObservableFloat64Counter, error) {
	return i, nil
}
func (i nonRecordingObservableFloat64Instrument) UpDownCounter(name string, opts ...instrument.Option) (instrument.ObservableFloat64UpDownCounter, error) {
	return i, nil
}
func (i nonRecordingObservableFloat64Instrument) Gauge(name string, opts ...instrument.Option) (instrument.ObservableFloat64Gauge, error) {
	return i, nil
}
func (i nonRecordingObservableFloat64Instrument) Observe(ctx context.Context, x float64, attrs ...attribute.KeyValue) {
}

// NewNoopFloat64Instruments returns a nonrecording implementation of metric.Float64ObserverInstruments and metric.Float64Instruments
func NewNoopFloat64Instruments() metric.Float64Instruments {
	return nonRecordingFloat64Instrument{}
}

var (
	_ instrument.Float64Counter       = nonRecordingFloat64Instrument{}
	_ instrument.Float64UpDownCounter = nonRecordingFloat64Instrument{}
	_ instrument.Float64Histogram     = nonRecordingFloat64Instrument{}
)

type nonRecordingFloat64Instrument struct {
	instrument.Synchronous
}

func (i nonRecordingFloat64Instrument) Counter(name string, opts ...instrument.Option) (instrument.Float64Counter, error) {
	return i, nil
}
func (i nonRecordingFloat64Instrument) UpDownCounter(name string, opts ...instrument.Option) (instrument.Float64UpDownCounter, error) {
	return i, nil
}
func (i nonRecordingFloat64Instrument) Histogram(name string, opts ...instrument.Option) (instrument.Float64Histogram, error) {
	return i, nil
}
func (i nonRecordingFloat64Instrument) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
}
func (i nonRecordingFloat64Instrument) Record(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
}

// NewNoopObservableInt64Instruments returns a nonrecording implementation of metric.Int64ObserverInstruments
func NewNoopObservableInt64Instruments() metric.ObservableInt64Instruments {
	return nonRecordingObservableInt64Instrument{}
}

var (
	_ instrument.ObservableInt64Counter       = nonRecordingObservableInt64Instrument{}
	_ instrument.ObservableInt64UpDownCounter = nonRecordingObservableInt64Instrument{}
	_ instrument.ObservableInt64Gauge         = nonRecordingObservableInt64Instrument{}
)

type nonRecordingObservableInt64Instrument struct {
	instrument.Asynchronous
}

func (i nonRecordingObservableInt64Instrument) Counter(name string, opts ...instrument.Option) (instrument.ObservableInt64Counter, error) {
	return i, nil
}
func (i nonRecordingObservableInt64Instrument) UpDownCounter(name string, opts ...instrument.Option) (instrument.ObservableInt64UpDownCounter, error) {
	return i, nil
}
func (i nonRecordingObservableInt64Instrument) Gauge(name string, opts ...instrument.Option) (instrument.ObservableInt64Gauge, error) {
	return i, nil
}
func (i nonRecordingObservableInt64Instrument) Observe(ctx context.Context, x int64, attrs ...attribute.KeyValue) {
}

// NewNoopInt64Instruments returns a nonrecording implementation of metric.Int64ObserverInstruments and metric.Int64Instruments
func NewNoopInt64Instruments() metric.Int64Instruments {
	return nonRecordingInt64Instrument{}
}

var (
	_ instrument.Int64Counter       = nonRecordingInt64Instrument{}
	_ instrument.Int64UpDownCounter = nonRecordingInt64Instrument{}
	_ instrument.Int64Histogram     = nonRecordingInt64Instrument{}
)

type nonRecordingInt64Instrument struct {
	instrument.Synchronous
}

func (i nonRecordingInt64Instrument) Counter(name string, opts ...instrument.Option) (instrument.Int64Counter, error) {
	return i, nil
}
func (i nonRecordingInt64Instrument) UpDownCounter(name string, opts ...instrument.Option) (instrument.Int64UpDownCounter, error) {
	return i, nil
}
func (i nonRecordingInt64Instrument) Histogram(name string, opts ...instrument.Option) (instrument.Int64Histogram, error) {
	return i, nil
}
func (i nonRecordingInt64Instrument) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
}
func (i nonRecordingInt64Instrument) Record(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
}
