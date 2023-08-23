package logbridge

import (
	"context"

	"go.opentelemetry.io/otel/logbridge"
	"go.opentelemetry.io/otel/sdk/instrumentation"
)

type producer struct {
	processor Processor
}

func NewProducer(opts ...Option) logbridge.LoggerProducer {
	p := &producer{}
	for _, opt := range opts {
		opt.apply(p)
	}

	return p
}

func (p *producer) Logger(name string, options ...logbridge.LoggerOption) logbridge.Logger {
	cfg := logbridge.NewLoggerConfig(options...)
	return &logger{
		scope: instrumentation.Scope{
			Name:      name,
			Version:   cfg.Version(),
			SchemaURL: cfg.SchemaURL(),
		},
		processor: p.processor,
	}
}

type logger struct {
	scope     instrumentation.Scope
	processor Processor
}

func (l *logger) Emit(ctx context.Context, record *logbridge.Record) {
	l.processor.Process(ctx, record)
}

type Option interface {
	apply(*producer)
}
type optionFunc func(*producer)

func (o optionFunc) apply(p *producer) { o(p) }

func WithProcessor(p Processor) Option {
	return optionFunc(func(prod *producer) {
		prod.processor = p
	})
}

type Processor interface {
	Process(context.Context, *logbridge.Record)
	// Shutdown()?
	// Flush()?
}

func NewSimpleProcessor(exporter Exporter) *simpleProcessor {
	return &simpleProcessor{exporter: exporter}
}

type simpleProcessor struct {
	exporter Exporter
}

func (p *simpleProcessor) Process(ctx context.Context, record *logbridge.Record) {
	p.exporter.Export(ctx, []*logbridge.Record{record})
}

type Exporter interface {
	Export(context.Context, []*logbridge.Record)
	// Shutdown()?
	// Flush()?
}
