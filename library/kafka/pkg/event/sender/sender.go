package sender

import (
	"encoding/json"

	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/kafka/pkg/config"
	"github.com/Finance-Tracker-MHS-DevDays-Fall-2025/notification/library/kafka/pkg/event"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type EventSender interface {
	SendEvent(event event.Wrapper) error
}

type eventSender struct {
	producer *kafka.Producer
	topic    string
}

func NewEventSender(
	config config.EventSenderConfig,
	logger *logrus.Logger,
	senderName string,
) (EventSender, error) {
	producer, err := kafka.NewProducer(config.ToKafkaConfigMap())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	go listener(producer, logger, senderName)

	return &eventSender{
		producer: producer,
		topic:    config.Topic,
	}, nil
}

func (sender *eventSender) SendEvent(event event.Wrapper) error {
	value, err := json.Marshal(event)
	if err != nil {
		return errors.WithStack(err)
	}

	message := &kafka.Message{
		Key:   []byte(event.RequestID.String()),
		Value: value,
		TopicPartition: kafka.TopicPartition{
			Topic:     &sender.topic,
			Partition: kafka.PartitionAny,
		},
		Timestamp: event.Timestamp,
	}

	err = sender.producer.Produce(message, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
