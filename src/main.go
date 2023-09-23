package main

import (
	"time"

	"github.com/duartesaraiva98/kafka-topic-replicator/configuration"
	replicator "github.com/duartesaraiva98/kafka-topic-replicator/kafka"
)

func main() {
	cfg := configuration.ReadConfig("/Users/duarte/Oss/kafka-topic-replicator/config.yaml")

	c := replicator.StartConsumer(cfg.Source.ConsumerConfig, cfg.Source.Topic)
	p := replicator.NewProducer(cfg.Destination.ProducerConfig)

	for {
		if !replicator.PipeTo(c, p, cfg.Destination.Topic) {
			break
		}
	}

	p.Flush(int((5 * time.Second).Milliseconds()))

	p.Close()
	c.Close()
}
