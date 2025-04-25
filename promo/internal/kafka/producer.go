package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type Producer struct {
	writers map[string]*kafka.Writer
}

func NewProducer(broker string, topics ...string) *Producer {
	writers := make(map[string]*kafka.Writer)
	for _, topic := range topics {
		writers[topic] = &kafka.Writer{
			Addr:                   kafka.TCP(broker),
			Topic:                  topic,
			RequiredAcks:           kafka.RequireAll,
			AllowAutoTopicCreation: true,
		}
	}
	return &Producer{writers: writers}
}

func (p *Producer) Send(ctx context.Context, topic string, msg proto.Message) error {
	b, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	writer, ok := p.writers[topic]
	if !ok {
		return fmt.Errorf("no writer for topic %s", topic)
	}

	return writer.WriteMessages(ctx, kafka.Message{
		Value: b,
		Time:  time.Now(),
	})
}

func (p *Producer) Close() error {
	var firstErr error
	for topic, w := range p.writers {
		if err := w.Close(); err != nil && firstErr == nil {
			firstErr = fmt.Errorf("closing writer %s: %w", topic, err)
		}
	}
	return firstErr
}
