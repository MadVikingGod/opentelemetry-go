module go.opentelemetry.io/otel/exporters/stdout/stdoutlogging

go 1.21

replace (
	go.opentelemetry.io/otel/bridge/slog => ../../../bridge/slog
	go.opentelemetry.io/otel/logbridge => ../../../logbridge
	go.opentelemetry.io/otel/sdk/logbridge => ../../../sdk/logbridge
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/bridge/slog v0.0.1 // indirect
	go.opentelemetry.io/otel/logbridge v0.0.1 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
)
