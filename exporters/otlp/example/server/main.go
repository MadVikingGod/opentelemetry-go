package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbtrace "go.opentelemetry.io/otel/exporters/otlp/internal/opentelemetry-proto-gen/collector/trace/v1"
)

func main() {
	lis, _ := net.Listen("tcp", ":9000")
	serv := grpc.NewServer()
	pbtrace.RegisterTraceServiceServer(serv, &traceServer{})
	fmt.Println("Starting Server localhost:9000")
	serv.Serve(lis)
}

type traceServer struct{}

func (t *traceServer) Export(ctx context.Context, request *pbtrace.ExportTraceServiceRequest) (*pbtrace.ExportTraceServiceResponse, error) {
	fmt.Println("Received request")
	return nil, status.Error(codes.Aborted, "Error Message")
}
