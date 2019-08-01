package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	phishqlpb "github.com/jloom6/phishql/.gen/proto/jloom6/phishql"
	"github.com/jloom6/phishql/handler"
	"github.com/jloom6/phishql/internal/db"
	"github.com/jloom6/phishql/mapper"
	"github.com/jloom6/phishql/service"
	"github.com/jloom6/phishql/storage/mysql"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
	dbDriver = "mysql"
	dbConnectionString = "wilson:wilson@/phish?parseTime=true"
)

func main() {
	sqlDB, err := sql.Open(dbDriver, dbConnectionString)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("failed to close db: %v", err)
		}
	}()

	store := mysql.New(mysql.Params{DB: db.New(sqlDB)})
	svc := service.New(service.Params{Store: store})
	h := handler.New(handler.Params{
		Service: svc,
		Mapper: mapper.New(),
	})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	phishqlpb.RegisterPhishQLServiceServer(s, h)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
