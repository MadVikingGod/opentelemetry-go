module go.opentelemetry.io/otel/sdk/logbridge

go 1.21

require (
	go.opentelemetry.io/otel/logbridge v0.0.1
	go.opentelemetry.io/otel/sdk v1.16.0
)

require go.opentelemetry.io/otel v1.16.0 // indirect

replace go.opentelemetry.io/otel/logbridge => ../../logbridge
