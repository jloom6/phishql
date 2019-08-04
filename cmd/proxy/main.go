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
	port = ":8080"
)

var (
	endpoint = "localhost:9090"//os.Getenv("PHISHQL_API_ENDPOINT")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Printf("listening on port %s (HTTP)\n", port)
	log.Printf("forwarding to %s (gRPC)\n", endpoint)

	mux := runtime.NewServeMux()
	err := phishqlpb.RegisterPhishQLServiceHandlerFromEndpoint(ctx, mux, endpoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
