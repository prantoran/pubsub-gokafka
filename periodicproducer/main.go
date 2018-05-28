package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "192.168.4.93:9092"
)

var curPartion = map[string]int32{}

func main() {
	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}
	// r := rand.New(rand.NewSource(99))
	cnt := 0
	for {
		// msk := r.Intn(3)
		topics := []string{"senz"}
		fmt.Println("topics:", topics)
		msg := "periodicmsg:" + strconv.Itoa(cnt)
		cnt++
		publish(&topics, msg, producer)
		time.Sleep(time.Millisecond * 500)
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

		pp, ok := curPartion[topic]
		if !ok {
			pp = 0
		}
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Value:     sarama.StringEncoder(message),
			Partition: pp,
		}
		curPartion[topic] = (pp + 1) % 2
		fmt.Println("topic:", topic, " msg:", msg)
		fmt.Println("pp:", pp)
		p, o, err := producer.SendMessage(msg)
		fmt.Println("msg topic:", msg.Topic, " value:", msg.Value, " msg partition:", msg.Partition)
		if err != nil {
			fmt.Println("Error publish: ", err.Error())
		}

		// publish async
		//producer.Input() <- &sarama.ProducerMessage{

		fmt.Println("Partition: ", p)
		fmt.Println("Offset: ", o)
	}

}
