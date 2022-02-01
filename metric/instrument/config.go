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
	"go.opentelemetry.io/otel/metric/unit"
	"golang.org/x/net/context"
)

// Config contains options for metric instrument descriptors.
type Config struct {
	description string
	unit        unit.Unit

	callback func(context.Context, Asynchronous)
}

// Description describes the instrument in human-readable terms.
func (cfg Config) Description() string {
	return cfg.description
}

// Unit describes the measurement unit for a instrument.
func (cfg Config) Unit() unit.Unit {
	return cfg.unit
}

// Callback returns the function that will be registered on asynchronous
// instrument creation.
func (cfg Config) Callback() func(context.Context, Asynchronous) {
	return cfg.callback
}

// Option is an interface for applying metric instrument options.
type Option interface {
	// ApplyMeter is used to set a Option value of a
	// Config.
	applyInstrument(*Config)
}

// NewConfig creates a new Config
// and applies all the given options.
func NewConfig(opts ...Option) Config {
	var config Config
	for _, o := range opts {
		o.applyInstrument(&config)
	}
	return config
}

type optionFunc func(*Config)

func (fn optionFunc) applyInstrument(cfg *Config) {
	fn(cfg)
}

// WithDescription applies provided description.
func WithDescription(desc string) Option {
	return optionFunc(func(cfg *Config) {
		cfg.description = desc
	})
}

// WithUnit applies provided unit.
func WithUnit(unit unit.Unit) Option {
	return optionFunc(func(cfg *Config) {
		cfg.unit = unit
	})
}

func WithCallback(fn func(context.Context, Asynchronous)) Option {
	return optionFunc(func(cfg *Config) {
		cfg.callback = fn
	})
}
