package logbridge

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type LoggerProducer interface {
	Logger(name string, options ...LoggerOption) Logger
}

type Logger interface {
	Emit(ctx context.Context, record *Record)
}

type LoggerConfig struct {
	version    string
	schemaURL  string
	attributes []attribute.KeyValue
}

func (c LoggerConfig) Version() string {
	return c.version
}
func (c LoggerConfig) SchemaURL() string {
	return c.schemaURL
}
func (c LoggerConfig) Attributes() []attribute.KeyValue {
	return c.attributes
}

func NewLoggerConfig(options ...LoggerOption) LoggerConfig {
	var config LoggerConfig
	for _, o := range options {
		config = o.apply(config)
	}
	return config
}

type LoggerOption interface {
	apply(LoggerConfig) LoggerConfig
}

type Record struct {
	Timestamp         time.Time
	ObservedTimestamp time.Time
	Severity          Severity
	SeverityText      string
	Body              string
	Attributes        []attribute.KeyValue
}
