package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"

	"statistics/config"
	"statistics/grpc_server"
	"statistics/kafka_consumer"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// ClickHouse connection
	chConn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.ClickHouseAddr},
		Auth: clickhouse.Auth{
			Database: cfg.ClickHouseDB,
			Username: cfg.ClickHouseUser,
			Password: cfg.ClickHousePassword,
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %v", err)
	}
	defer chConn.Close()

	// Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokers,
		"group.id":          cfg.KafkaGroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to topics
	topics := []string{"view_promo", "like_promo", "comment_promo", "view_post", "like_post", "comment_post"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to Kafka topics: %v", err)
	}

	// Start Kafka consumer in a goroutine
	go kafka_consumer.Consume(consumer, chConn)

	// gRPC server
	grpcServer := grpc.NewServer()
	statisticsService := grpc_server.NewStatisticsService(chConn)
	statistics.RegisterStatisticsServiceServer(grpcServer, statisticsService)

	// Start gRPC server
	go func() {
		if err := grpcServer.Serve(cfg.GRPCListenAddr); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Wait for termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Graceful shutdown
	grpcServer.GracefulStop()
}
