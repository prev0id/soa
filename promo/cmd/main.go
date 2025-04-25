package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"promo/internal/api"
	"promo/internal/kafka"
	"promo/internal/store"
	loyaltypb "promo/pkg/loyalty"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "host=promo-postgres port=5432 user=user password=password dbname=promo_db sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer pool.Close()

	store := store.NewStore(pool)

	producer := kafka.NewProducer("kafka:9092", "client-registrations", "likes", "clicks", "views", "comments")
	defer producer.Close()

	grpcService := api.NewService(store, producer)

	grpcServer := grpc.NewServer()
	loyaltypb.RegisterLoyaltyCoreServer(grpcServer, grpcService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":12301")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :12301")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
