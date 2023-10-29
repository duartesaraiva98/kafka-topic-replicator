package configuration

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

type MultiClusterConfig struct {
	source      MultiClusterSourceConfig
	destination MultiClusterDestinationConfig
}

type MultiClusterSourceConfig struct {
	consumer MultiClusterSourceConsumer
	client   kafka.ConfigMap
}

type MultiClusterSourceConsumer struct {
	topic            string
	groupId          string
	bootstrapServers string
	resetStrategy    string
}

type MultiClusterDestinationConfig struct {
	producer MultiClusterDestinationProducer
	client   kafka.ConfigMap
}

type MultiClusterDestinationProducer struct {
	topic            string
	bootstrapServers string
}

func (cfg MultiClusterConfig) SourceTopic() string {
	return cfg.source.consumer.topic
}

func (cfg MultiClusterConfig) SourceClientConfiguration() kafka.ConfigMap {
	client := cfg.source.client
	client.SetKey("group.id", cfg.source.consumer.groupId)
	client.SetKey("bootstrap.servers", cfg.source.consumer.bootstrapServers)
	client.SetKey("auto.offset.reset", cfg.source.consumer.resetStrategy)

	return client
}

func (cfg MultiClusterConfig) DestinationTopic() string {
	return cfg.destination.producer.topic
}

func (cfg MultiClusterConfig) DestinationClientConfiguration() kafka.ConfigMap {
	client := cfg.source.client
	client.SetKey("bootstrap.servers", cfg.destination.producer.bootstrapServers)

	return client
}

func resolveMultiClusterConfig() Config {
	var sourceConsumer MultiClusterSourceConsumer

	err := viper.UnmarshalKey("source.consumer", &sourceConsumer)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	cons := MultiClusterSourceConfig{
		consumer: sourceConsumer,
		client:   makeConfigMapFrom(viper.GetStringMap("source.client")),
	}

	var destinationProducer MultiClusterDestinationProducer

	err = viper.UnmarshalKey("destination.producer", &destinationProducer)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	prod := MultiClusterDestinationConfig{
		producer: destinationProducer,
		client:   makeConfigMapFrom(viper.GetStringMap("destination.client")),
	}

	return MultiClusterConfig{
		source:      cons,
		destination: prod,
	}
}
