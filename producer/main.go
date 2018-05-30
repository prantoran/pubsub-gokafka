package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/prantoran/pubsub-gokafka/api"
	"github.com/prantoran/pubsub-gokafka/data"
)

func main() {
	// create producer
	producer, err := api.InitProducer(1, "pubsub_sarama_client")
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

		if err := api.Publish(&topics, data.NewKafkaMsg(msg), producer); err != nil {
			log.Println("could not publish: ", err)
		}

		// publish with go routene
		// go publish(msg, producer)
	}
}
