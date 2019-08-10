package main

import (
	"io"
	"log"
	"net/http"
	"os"

	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/graphql/resolver"
	"github.com/jloom6/phishql/graphql/schema"
	"github.com/jloom6/phishql/mapper"
	"google.golang.org/grpc"
)

const (
	port = ":8420"
)

var (
	endpoint = os.Getenv("PHISHQL_API_ENDPOINT")
)

func main() {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer closeConn(conn)

	// Init the GraphQL types since a circular dependency needs to be defined
	schema.InitTypes()

	http.HandleFunc("/graphql", makeHandleFunc(conn))

	log.Printf("listening on port %s (GraphQL)\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func closeConn(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatalf("failed to close connection: %v", err)
	}
}

func makeHandleFunc(cc *grpc.ClientConn) func(http.ResponseWriter, *http.Request) {
	r := resolver.New(resolver.Params{
		Client: phishqlpb.NewPhishQLServiceClient(cc),
		Mapper: mapper.New(),
	})
	f, err := schema.NewHandleFunc(schema.Params{
		Resolver: r,
	})
	if err != nil {
		log.Fatalf("failed to create schema, error: %v", err)
	}

	return f
}
