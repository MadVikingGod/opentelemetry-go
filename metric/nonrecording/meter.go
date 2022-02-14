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

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
)

func NewNoopMeterProvider() metric.MeterProvider {
	return noopMeterProvider{}
}

type noopMeterProvider struct{}

var _ metric.MeterProvider = noopMeterProvider{}

func (noopMeterProvider) Meter(instrumentationName string, opts ...metric.MeterOption) metric.Meter {
	return noopMeter{}
}

type noopMeter struct{}

var _ metric.Meter = noopMeter{}

func (noopMeter) Int64Observer() metric.Int64ObserverInstruments {
	return NewNoopInt64ObserverInstruments()
}
func (noopMeter) Float64Observer() metric.Float64ObserverInstruments {
	return NewNoopFloat64ObserverInstruments()
}
func (noopMeter) Int64() metric.Int64Instruments {
	return NewNoopInt64Instruments()
}
func (noopMeter) Float64() metric.Float64Instruments {
	return NewNoopFloat64Instruments()
}
func (noopMeter) RegisterCallback([]instrument.Asynchronous, func(context.Context)) error {
	return nil
}
