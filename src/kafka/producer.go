package replicator

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewProducer(config kafka.ConfigMap) *kafka.Producer {
	p, err := kafka.NewProducer(&config)
	if err != nil {
		panic(err)
	}

	return p
}
