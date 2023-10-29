package configuration

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

type SingleClusterConfig struct {
	bootstrapServers    string
	client              kafka.ConfigMap
	sourceConsumer      SingleClusterSourceConsumer
	destinationProducer SingleClusterDestinationProducer
}

type SingleClusterSourceConsumer struct {
	topic         string
	groupId       string
	resetStrategy string
}

type SingleClusterDestinationProducer struct {
	topic string
}

func (cfg SingleClusterConfig) SourceTopic() string {
	return cfg.sourceConsumer.topic
}

func (cfg SingleClusterConfig) SourceClientConfiguration() kafka.ConfigMap {
	client := cfg.client
	client.SetKey("group.id", cfg.sourceConsumer.groupId)
	client.SetKey("bootstrap.servers", cfg.bootstrapServers)
	client.SetKey("auto.offset.reset", cfg.sourceConsumer.resetStrategy)

	return client
}

func (cfg SingleClusterConfig) DestinationTopic() string {
	return cfg.destinationProducer.topic
}

func (cfg SingleClusterConfig) DestinationClientConfiguration() kafka.ConfigMap {
	client := cfg.client
	client.SetKey("bootstrap.servers", cfg.bootstrapServers)

	return client
}

func resolveSingleClusterConfig() Config {
	var sourceConsumer SingleClusterSourceConsumer

	err := viper.UnmarshalKey("source.consumer", &sourceConsumer)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	var destinationProducer SingleClusterDestinationProducer

	err = viper.UnmarshalKey("destination.producer", &destinationProducer)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return SingleClusterConfig{
		bootstrapServers:    viper.GetString("bootstrapServers"),
		client:              makeConfigMapFrom(viper.GetStringMap("source.client")),
		sourceConsumer:      sourceConsumer,
		destinationProducer: destinationProducer,
	}
}
