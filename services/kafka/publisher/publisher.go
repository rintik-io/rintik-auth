package publisher

import (
	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/golibs/kafka"
	"github.com/rintik-io/rintik-auth/configs"
)

type KafkaPublisher struct{}

// Connection : Kafka Publisher Open Connection
var Connection *kafka.Publisher

// Start : Starting Kafka Publisher
func (p KafkaPublisher) Start() error {
	kafkaPublisher := kafka.NewKafkaPublisher(
		configs.Properties.Services.Kafka.Brokers,
		configs.Properties.Services.Kafka.Verbose,
		configs.Properties.Services.Kafka.DialTimeout,
		configs.Properties.Services.Kafka.Publisher.RetryMax,
		configs.Properties.Services.Kafka.Publisher.Timeout,
		configs.Properties.Services.Kafka.Publisher.Idempotent,
		configs.Properties.Services.Kafka.Version,
	)
	kafkaPublisher.Logger = golibs.Log
	err := kafkaPublisher.CreatePublisherConnection()
	if err != nil {
		return err
	}

	Connection = kafkaPublisher

	return nil
}
