package message_producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type TravelMsgProducer interface {
	TravelSaved()
	TravelUpdated()
	TravelDeleted()
}

func NewTravelMsgProducer() TravelMsgProducer {
	return &kafkaProducer{
		k: &kafka.Writer{
			Addr:     kafka.TCP("localhost:9092"),
			Topic:    "topic-Travel",
			Balancer: &kafka.LeastBytes{},
		},
	}
}

type kafkaProducer struct {
	k *kafka.Writer
}

func (k kafkaProducer) TravelSaved() {
	msg := kafka.Message{
		Key:   []byte("New Trip"),
		Value: []byte("New trip was created"),
	}

	err := k.k.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Fatalln(err)
	}
}

func (k kafkaProducer) TravelUpdated() {
	msg := kafka.Message{
		Key:   []byte("Trip updated"),
		Value: []byte("Existing trip was updated"),
	}

	err := k.k.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Fatalln(err)
	}
}

func (k kafkaProducer) TravelDeleted() {
	msg := kafka.Message{
		Key:   []byte("Trip was deleted"),
		Value: []byte("Existing trip was deleted"),
	}

	err := k.k.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Fatalln(err)
	}
}
