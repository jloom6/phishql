package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	endpoint = "localhost:9090"
	port = ":8080"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Printf("listening on port %s (HTTP)\n", port)
	log.Printf("forwarding to %s (gRPC)\n", endpoint)

	mux := runtime.NewServeMux()
	err := phishqlpb.RegisterPhishQLServiceHandlerFromEndpoint(ctx, mux, endpoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}

	return http.ListenAndServe(port, mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
