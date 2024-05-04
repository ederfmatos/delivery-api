package events

import (
	"delivery-api/application/events"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaEventEmitter struct {
	kafkaProducer *kafka.Producer
}

func NewKafkaEventEmitter(server string) *KafkaEventEmitter {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":   server,
		"delivery.timeout.ms": "0",
		"enable.idempotence":  "true",
	})
	if err != nil {
		panic(err)
	}

	return &KafkaEventEmitter{kafkaProducer: producer}
}

func (emitter *KafkaEventEmitter) Emit(event events.Event) error {
	topic := event.Name()
	messageBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	message := &kafka.Message{
		Value:          messageBytes,
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(event.Id()),
	}
	return emitter.kafkaProducer.Produce(message, nil)
}
