package consumer

import (
	"fmt"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/golibs/kafka"
	"github.com/rintik-io/rintik-auth/configs"
)

type KafkaConsumer struct{}

// Start : Starting Kafka Consumer
func (p KafkaConsumer) Start() error {
	var err error

	golibs.Log.Println("| Consumer Type |", configs.Properties.Services.Kafka.Consumer.Type)

	switch configs.Properties.Services.Kafka.Consumer.Type {
	case "single":
		consumerSingle := kafka.NewConsumerSingle(
			configs.Properties.Services.Kafka.Brokers,
			configs.Properties.Services.Kafka.Consumer.Topic,
			configs.Properties.Services.Kafka.Verbose,
			configs.Properties.Services.Kafka.Consumer.Oldest,
			true,
			nil,
		)
		consumerSingle.Logger = golibs.Log

		// consumerSingle.AddHandler("t-favorites", &handlers.Favorites{}, "Handle")
		err = consumerSingle.Consume()
	case "group":
		consumerGroup := kafka.NewConsumerGroup(
			configs.Properties.Services.Kafka.Brokers,
			configs.Properties.Services.Kafka.Consumer.Topic,
			configs.Properties.Services.Kafka.Consumer.Group,
			configs.Properties.Services.Kafka.Assignor,
			configs.Properties.Services.Kafka.Version,
			configs.Properties.Services.Kafka.Verbose,
			configs.Properties.Services.Kafka.Consumer.Oldest,
			true,
			nil,
		)
		consumerGroup.Logger = golibs.Log

		// consumerGroup.AddHandler("t-favorites", &handlers.Favorites{}, "Handle")
		err = consumerGroup.Consume()
	default:
		return fmt.Errorf("Unrecognized Consumer Type")
	}

	return err
}
