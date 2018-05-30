package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/prantoran/pubsub-gokafka/conf"
	"github.com/prantoran/pubsub-gokafka/data"
)

func InitProducer(brokerFlag int, clientID string) (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	// sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// for verbose logger
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = clientID

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer(conf.AllBrokers(brokerFlag), config)

	return prd, err
}

func Publish(topics *[]string, message *data.KafkaMsg, producer sarama.SyncProducer) error {
	// publish sync
	
	b, err := json.Marshal(&message)
	if err != nil {
		return nil
	}
	for _, topic := range *topics {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(b),
		}

		fmt.Println("topic:", topic, " msg:", msg)
		p, o, err := producer.SendMessage(msg)
		if err != nil {
			return fmt.Errorf("Error publish: %v", err.Error())
		}

		// publish async
		//producer.Input() <- &sarama.ProducerMessage{

		fmt.Println("Partition: ", p)
		fmt.Println("Offset: ", o)
	}
	return nil
}
