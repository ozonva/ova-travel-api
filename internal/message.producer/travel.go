package message_producer

import (
	"context"
	"fmt"
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
		Key:   []byte(fmt.Sprintf("New Trip")),
		Value: []byte(fmt.Sprintf("New trip was created")),
	}

	err := k.k.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Fatalln(err)
	}
}

func (k kafkaProducer) TravelUpdated() {
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Trip updated")),
		Value: []byte(fmt.Sprintf("Existing trip was updated")),
	}

	err := k.k.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Fatalln(err)
	}
}

func (k kafkaProducer) TravelDeleted() {
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Trip was deleted")),
		Value: []byte(fmt.Sprintf("Existing trip was deleted")),
	}

	err := k.k.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Fatalln(err)
	}
}
