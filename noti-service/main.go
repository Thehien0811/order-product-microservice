package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Cannot create consumer")
	}

	partion := 0
	consumerPart, err := consumer.ConsumePartition("create_order", int32(partion), sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Failed to consume partition")
	}
	defer consumerPart.Close()
	for {
		select {
		case msg, ok := <-consumerPart.Messages():
			if !ok {
				log.Println("Channel closed, exiting goroutine")
				return
			}
			fmt.Println(string(msg.Value))
		}
	}

}
