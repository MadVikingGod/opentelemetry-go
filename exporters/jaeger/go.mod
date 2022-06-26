module go.opentelemetry.io/otel/exporters/jaeger

go 1.16

require (
	github.com/google/go-cmp v0.5.8
	github.com/stretchr/testify v1.7.5
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
)

replace go.opentelemetry.io/otel/trace => ../../trace

replace go.opentelemetry.io/otel => ../..

replace go.opentelemetry.io/otel/sdk => ../../sdk
