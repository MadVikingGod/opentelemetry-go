package slog

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutlogging"
	sdklogbridge "go.opentelemetry.io/otel/sdk/logbridge"
)

func Example() {
	ctx := context.Background()

	exp := stdoutlogging.NewExporterNoTime(os.Stdout)
	producer := sdklogbridge.NewProducer(sdklogbridge.WithProcessor(sdklogbridge.NewSimpleProcessor(exp)))

	handler := NewHandler(WithLogger(producer.Logger("example")))

	logger := slog.New(handler)

	logger.InfoContext(ctx, "Hello, world!", "example", "String", "number", 123)

	// Output: Info body="Hello, world!" example="String" number="123"
}
