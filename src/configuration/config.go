package configuration

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

type Config interface {
	SourceTopic() string
	SourceClientConfiguration() kafka.ConfigMap

	DestinationTopic() string
	DestinationClientConfiguration() kafka.ConfigMap
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
