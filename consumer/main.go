package main

import (
	"log"

	"github.com/IBM/sarama"
	snappy "github.com/eapache/go-xerial-snappy"
)

func main() {
	brokerUrl := "localhost:9092"

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumer, err := sarama.NewConsumer([]string{brokerUrl}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	topicName := "test"
	consumePartition, err := consumer.ConsumePartition(topicName, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer consumePartition.Close()

	for {
		select {
		case message := <-consumePartition.Messages():
			decompress, err := snappy.Decode([]byte(message.Value))
			if err != nil {
				log.Printf("Error decoding message: %v", err)
			} else {
				log.Printf("Received message: %s\n", decompress)
			}
		case err := <-consumePartition.Errors():
			log.Printf("Error: %v\n", err)
			return
		}
	}
}
