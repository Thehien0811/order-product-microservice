package kafka

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

func (uc implUseCase) SendMsg(topic string, par int, msg string) {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	msgSend := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(bytes),
	}

	_, _, err = producer.SendMessage(msgSend)
	if err != nil {
		log.Fatal(err)
	}
}
func (uc implUseCase) ReceiveMsg(topic string, par int, msg string) {
	
}
