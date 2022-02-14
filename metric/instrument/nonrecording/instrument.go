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

package nonrecording // import "go.opentelemetry.io/otel/metric/instrument/nonrecording"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
)

// NewNoopFloat64ObserverInstruments returns a nonrecording implementation of metric.Float64ObserverInstruments
func NewNoopFloat64ObserverInstruments() metric.Float64ObserverInstruments {
	return nonRecordingFloat64ObserverInstrument{}
}

var (
	_ instrument.Float64ObserverCounter       = nonRecordingFloat64ObserverInstrument{}
	_ instrument.Float64ObserverUpDownCounter = nonRecordingFloat64ObserverInstrument{}
	_ instrument.Float64ObserverGauge         = nonRecordingFloat64ObserverInstrument{}
)

type nonRecordingFloat64ObserverInstrument struct {
	instrument.Asynchronous
}

func (i nonRecordingFloat64ObserverInstrument) Counter(name string, opts ...instrument.Option) (instrument.Float64ObserverCounter, error) {
	return i, nil
}
func (i nonRecordingFloat64ObserverInstrument) UpDownCounter(name string, opts ...instrument.Option) (instrument.Float64ObserverUpDownCounter, error) {
	return i, nil
}
func (i nonRecordingFloat64ObserverInstrument) Gauge(name string, opts ...instrument.Option) (instrument.Float64ObserverGauge, error) {
	return i, nil
}
func (i nonRecordingFloat64ObserverInstrument) Observe(ctx context.Context, x float64, attrs ...attribute.KeyValue) {
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

// NewNoopInt64ObserverInstruments returns a nonrecording implementation of metric.Int64ObserverInstruments
func NewNoopInt64ObserverInstruments() metric.Int64ObserverInstruments {
	return nonRecordingInt64ObserverInstrument{}
}

var (
	_ instrument.Int64ObserverCounter       = nonRecordingInt64ObserverInstrument{}
	_ instrument.Int64ObserverUpDownCounter = nonRecordingInt64ObserverInstrument{}
	_ instrument.Int64ObserverGauge         = nonRecordingInt64ObserverInstrument{}
)

type nonRecordingInt64ObserverInstrument struct {
	instrument.Asynchronous
}

func (i nonRecordingInt64ObserverInstrument) Counter(name string, opts ...instrument.Option) (instrument.Int64ObserverCounter, error) {
	return i, nil
}
func (i nonRecordingInt64ObserverInstrument) UpDownCounter(name string, opts ...instrument.Option) (instrument.Int64ObserverUpDownCounter, error) {
	return i, nil
}
func (i nonRecordingInt64ObserverInstrument) Gauge(name string, opts ...instrument.Option) (instrument.Int64ObserverGauge, error) {
	return i, nil
}
func (i nonRecordingInt64ObserverInstrument) Observe(ctx context.Context, x int64, attrs ...attribute.KeyValue) {
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
