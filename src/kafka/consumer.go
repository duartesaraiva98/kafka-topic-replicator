package replicator

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

/*
&kafka.ConfigMap{
		"bootstrap.servers": bootstrap,
		"group.id":          groupId,
		"auto.offset.reset": autoOffsetReset,
	}
*/

func StartConsumer(config kafka.ConfigMap, topic string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&config)

	if err != nil {
		panic(err)
	}

	err2 := c.SubscribeTopics([]string{topic}, nil)

	if err2 != nil {
		panic(err)
	}

	return c
}

func PipeTo(c *kafka.Consumer, p *kafka.Producer, topic string) bool {
	msg, err := c.ReadMessage(10 * time.Second)
	if err == nil {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		newMessage := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            msg.Key,
			Value:          msg.Value,
			Headers:        msg.Headers,
		}

		err2 := p.Produce(newMessage, nil)
		if err2 != nil {
			fmt.Printf("Error on producing message %s\n%v\n", msg.TopicPartition, err2)
		}
	} else if !err.(kafka.Error).IsTimeout() {
		// The client will automatically try to recover from all errors.
		// Timeout is not considered an error because it is raised by
		// ReadMessage in absence of messages.
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	} else {
		fmt.Printf("Timeout should mean no more messages\n")
		return false
	}
	return true
}
