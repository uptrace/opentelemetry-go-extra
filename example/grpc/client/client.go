package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/example/grpc/api"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	target := os.Getenv("GRPC_TARGET")
	if target == "" {
		target = ":9999"
	}

	log.Println("connecting to", target)

	conn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() { _ = conn.Close() }()

	client := api.NewHelloServiceClient(conn)
	if err := sayHello(client); err != nil {
		log.Fatal(err)
		return
	}
}

func sayHello(client api.HelloServiceClient) error {
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client",
		"user-id", "test-user",
	))

	resp, err := client.SayHello(ctx, &api.HelloRequest{Greeting: "World"})
	if err != nil {
		return err
	}
	log.Println("reply:", resp.Reply)

	return nil
}
