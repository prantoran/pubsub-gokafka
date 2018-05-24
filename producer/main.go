package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "192.168.4.93:9092"
)

func main() {
	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	// read command line input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')
		fmt.Printf("msg: %v type: %T \n", msg, msg)
		words := strings.Split(msg, " ")
		msg = ""

		for i, u := range words {
			fmt.Println("i:", i, " u:", u)
			if i == 0 {
				continue
			}
			msg += u
		}

		fmt.Println("upd msg:", msg)
		topics := strings.Split(words[0], ",")
		fmt.Println("topics:", topics)
		// publish without goroutene
		publish(&topics, msg, producer)

		// publish with go routene
		// go publish(msg, producer)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = "pubsub_sarama_client"

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func publish(topics *[]string, message string, producer sarama.SyncProducer) {
	// publish sync

	for _, topic := range *topics {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		fmt.Println("topic:", topic, " msg:", msg)
		p, o, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Error publish: ", err.Error())
		}

		// publish async
		//producer.Input() <- &sarama.ProducerMessage{

		fmt.Println("Partition: ", p)
		fmt.Println("Offset: ", o)
	}

}
