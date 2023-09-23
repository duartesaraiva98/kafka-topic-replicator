package configuration

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

type Config struct {
	Source      SourceConfig
	Destination DestinationConfig
}

type SourceConfig struct {
	Topic          string
	ConsumerConfig *kafka.ConfigMap
}

type DestinationConfig struct {
	Topic          string
	ProducerConfig *kafka.ConfigMap
}

func ReadConfig(filePath string) Config {
	regexPattern := "^(.*)\\/([^\\/]+)$"
	regex := regexp.MustCompile(regexPattern)
	result := regex.Split(filePath, -1)

	if filePath[0] != '/' {
		result = []string{"", filePath}
	}

	viper.AddConfigPath(defaultIfEmpty(result[0], "."))
	viper.SetConfigName(strings.ReplaceAll(result[1], ".yaml", ""))
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	consumerProperties := viper.GetStringMap("source.consumer.properties")

	consConfig := makeConfigMapFrom(consumerProperties)
	consConfig.SetKey("group.id", viper.GetString("source.consumer.group.id"))
	consConfig.SetKey("bootstrap.servers", viper.GetString("source.consumer.bootstrap.servers"))
	consConfig.SetKey("auto.offset.reset", viper.GetString("source.consumer.auto.offset.reset"))

	producerProperties := viper.GetStringMap("destination.producer.properties")

	prodConfig := makeConfigMapFrom(producerProperties)
	prodConfig.SetKey("bootstrap.servers", viper.GetString("destination.producer.bootstrap.servers"))

	sourceConfig := SourceConfig{
		Topic:          viper.GetString("source.topic"),
		ConsumerConfig: consConfig,
	}

	destinationConfig := DestinationConfig{
		Topic:          viper.GetString("destination.topic"),
		ProducerConfig: prodConfig,
	}

	return Config{
		Source:      sourceConfig,
		Destination: destinationConfig,
	}
}

func makeConfigMapFrom(stringMap map[string]interface{}) *kafka.ConfigMap {
	configMap := kafka.ConfigMap{}
	for k, v := range stringMap {
		err := configMap.SetKey(k, v)
		if err != nil {
			panic(err)
		}
	}
	return &configMap
}

func defaultIfEmpty(str string, defaultValue string) string {
	if str == "" {
		return defaultValue
	} else {
		return str
	}
}
