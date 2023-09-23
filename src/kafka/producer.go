package replicator

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

/*
kafka.ConfigMap{"bootstrap.servers": "localhost"}
*/

func NewProducer(config *kafka.ConfigMap) *kafka.Producer {
	p, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}

	return p
}
