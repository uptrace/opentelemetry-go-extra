package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/example/grpc/api"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

const addr = ":9999"

type helloServer struct {
	api.HelloServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	fmt.Println("trace", otelplay.TraceURL(trace.SpanFromContext(ctx)))

	time.Sleep(50 * time.Millisecond)
	return &api.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	log.Println("serving on", addr)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	api.RegisterHelloServiceServer(server, &helloServer{})
	if err := server.Serve(ln); err != nil {
		log.Fatal(err)
		return
	}
}
