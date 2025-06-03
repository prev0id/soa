package consumer

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"statistics/internal/domain"

	"github.com/segmentio/kafka-go"
)

func StartConsumers(db *sql.DB) {
	topics := []string{"view_promo", "like_promo", "comment_promo", "view_post", "like_post", "comment_post"}

	for _, topic := range topics {
		go consumeTopic(topic, db)
	}
}

func consumeTopic(topic string, db *sql.DB) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		GroupID: "statistics-servcie",
		Topic:   topic,
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Failed to read message from topic %s: %v", topic, err)
			continue
		}

		var event domain.Event
		if err = json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("Failed to unmarshal event from topic %s: %v", topic, err)
			continue
		}

		date := time.Now().Format("2006-01-02")
		switch topic {
		case "view_promo":
			_, err = db.Exec("INSERT INTO promo_stats (promo_id, date, views, likes, comments) VALUES (?, ?, 1, 0, 0)", event.PromoID, date)
		case "like_promo":
			_, err = db.Exec("INSERT INTO promo_stats (promo_id, date, views, likes, comments) VALUES (?, ?, 0, 1, 0)", event.PromoID, date)
		case "comment_promo":
			_, err = db.Exec("INSERT INTO promo_stats (promo_id, date, views, likes, comments) VALUES (?, ?, 0, 0, 1)", event.PromoID, date)
		case "view_post":
			_, err = db.Exec("INSERT INTO post_stats (post_id, date, views, likes, comments) VALUES (?, ?, 1, 0, 0)", event.PostID, date)
		case "like_post":
			_, err = db.Exec("INSERT INTO post_stats (post_id, date, views, likes, comments) VALUES (?, ?, 0, 1, 0)", event.PostID, date)
		case "comment_post":
			_, err = db.Exec("INSERT INTO post_stats (post_id, date, views, likes, comments) VALUES (?, ?, 0, 0, 1)", event.PostID, date)
		}

		if err != nil {
			log.Printf("Failed to insert into ClickHouse for topic %s: %v", topic, err)
		}
	}
}
