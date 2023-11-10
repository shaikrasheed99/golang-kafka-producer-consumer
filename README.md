# Kafka Producer & Consumer applications in Golang

This is a simple example of a Kafka producer and consumer written in Go using the Sarama library. The producer sends messages to a Kafka topic, and the consumer reads and processes those messages.

## Prerequisites

Before running this application, ensure that you have the following installed:

* Go (Golang) - You can download and install Go from the [official website](https://golang.org/dl/).
* Kafka - You need a running Kafka broker. You can download and set up Apache Kafka from the [official website](https://kafka.apache.org/downloads).

## Setup

1. Clone this repository to your local machine
```bash
git clone https://github.com/shaikrasheed99/golang-kafka-producer-consumer.git
```

2. Install the required Go packages
```bash
go get .
```

3. Update the Kafka broker URL in the producer and consumer code if your Kafka broker is running on a different host and port.

4. Build the producer and consumer executables
```bash
go build producer.go
go build consumer.go
```

## Running the Producer

1. Start the Kafka producer by running the producer executable
```bash
./producer
```

2. The producer will prompt you to enter messages. Type a message and press Enter. To exit the producer, type "exit" and press Enter.

3. The producer will send the messages to the Kafka topic specified in the code (default is "test").

## Running the Consumer

1. Start the Kafka consumer by running the consumer executable
```bash
./consumer
```

2. The consumer will subscribe to the Kafka topic specified in the code (default is "test") and start reading messages.

3. The consumer will print received messages to the console, or it will log any errors encountered during message processing.

4. To stop the consumer, press `Ctrl+C`.