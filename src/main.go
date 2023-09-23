package main

import (
	"fmt"
	"os"
	"time"

	"github.com/duartesaraiva98/kafka-topic-replicator/configuration"
	replicator "github.com/duartesaraiva98/kafka-topic-replicator/kafka"
)

func main() {
	filePath := os.Getenv("CONFIG_FILE")

	if filePath == "" {
		fmt.Println("`CONFIG_FILE` environment variable needs to be set and non-empty")
		os.Exit(1)
	}

	cfg := configuration.ReadConfig(filePath)

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
