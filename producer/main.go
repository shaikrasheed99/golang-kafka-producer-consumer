package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/IBM/sarama"
	snappy "github.com/eapache/go-xerial-snappy"
)

func main() {
	brokerUrl := "localhost:9092"

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	producer, err := sarama.NewAsyncProducer([]string{brokerUrl}, config)
	if err != nil {
		log.Fatalln("Error while creating async producer", err)
	}

	defer producer.AsyncClose()

	messages := make(chan string)

	go readFromTerminal(messages)

	sendMessagesToBroker(producer, messages)

	producer.AsyncClose()
}

func readFromTerminal(messages chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter message to send (or type 'exit' to quit): ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error taking input from terminal:", err)
			close(messages)
			return
		}

		message = strings.TrimSpace(message)
		if message == "exit" {
			close(messages)
			return
		}

		messages <- message
	}
}

func sendMessagesToBroker(producer sarama.AsyncProducer, messages chan string) {
	topicName := "test"
	for message := range messages {
		compressed := snappy.Encode([]byte(message))

		msg := &sarama.ProducerMessage{
			Topic: topicName,
			Value: sarama.ByteEncoder(compressed),
		}

		producer.Input() <- msg
	}
}
