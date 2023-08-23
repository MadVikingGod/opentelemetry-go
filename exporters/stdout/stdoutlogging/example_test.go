package stdoutlogging_test

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlogging"
	"go.opentelemetry.io/otel/logbridge"
	sdklogbridge "go.opentelemetry.io/otel/sdk/logbridge"
)

func Example() {
	ctx := context.Background()

	exp := stdoutlogging.NewExporterNoTime(os.Stdout)
	producer := sdklogbridge.NewProducer(sdklogbridge.WithProcessor(sdklogbridge.NewSimpleProcessor(exp)))

	logger := producer.Logger("example")

	logger.Emit(ctx, &logbridge.Record{
		Body:         "Hello, world!",
		Severity:     logbridge.SeverityInfo,
		SeverityText: logbridge.SeverityInfo.String(),
		Attributes: []attribute.KeyValue{
			attribute.String("example", "String"),
			attribute.Int("number", 123),
		},
	})
	// Output: Info body="Hello, world!" example="String" number="123"
}
