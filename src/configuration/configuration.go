package configuration

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

type Config struct {
	Source      SourceConfig
	Destination DestinationConfig
}

type SourceConfig struct {
	Consumer SourceConsumer
	Client   kafka.ConfigMap
}

type DestinationConfig struct {
	Producer DestinationProducer
	Client   kafka.ConfigMap
}

type SourceConsumer struct {
	Topic            string
	GroupId          string
	BootstrapServers string
	ResetStrategy    string
}

type DestinationProducer struct {
	Topic            string
	BootstrapServers string
}

func ReadConfig(filePath string) Config {
	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	multiCluster := viper.GetBool("multiCluster")

	if multiCluster {
		return resolveMultiClusterConfig()
	} else {
		return resolveSingleClusterConfig()
	}
}

func resolveSingleClusterConfig() Config {

}

func resolveMultiClusterConfig() Config {
	var sourceConsumer SourceConsumer

	err := viper.UnmarshalKey("source.consumer", &sourceConsumer)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	cons := SourceConfig{
		Consumer: sourceConsumer,
		Client:   makeConfigMapFrom(viper.GetStringMap("source.client")),
	}

	var destinationProducer DestinationProducer

	err = viper.UnmarshalKey("destination.producer", &destinationProducer)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	prod := DestinationConfig{
		Producer: destinationProducer,
		Client:   makeConfigMapFrom(viper.GetStringMap("destination.client")),
	}

	return Config{
		Source:      cons,
		Destination: prod,
	}
}

func makeConfigMapFrom(stringMap map[string]interface{}) kafka.ConfigMap {
	configMap := kafka.ConfigMap{}
	for k, v := range stringMap {
		err := configMap.SetKey(k, v)
		if err != nil {
			panic(err)
		}
	}
	return configMap
}
