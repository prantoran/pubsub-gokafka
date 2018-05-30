package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/prantoran/pubsub-gokafka/api"
	"github.com/prantoran/pubsub-gokafka/data"

	"github.com/prantoran/pubsub-gokafka/conf"
)

func main() {

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	// req to activate msg.Error() chan
	config.Group.Return.Notifications = true

	// init consumer
	topics := []string{"senz", "renz"}
	consumer, err := cluster.NewConsumer(conf.AllBrokers(1), "c1", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {

		case msg, ok := <-consumer.Messages():
			// This channel will only return if Config.Group.Mode option is set to
			// ConsumerModeMultiplex (default).
			if ok {

				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s timestamp: %v\nheaders: %v\n",
					msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value, msg.Timestamp, msg.Headers)

				d := data.KafkaMsg{}
				err := json.Unmarshal(msg.Value, &d)

				if err != nil {
					fmt.Println("could not unmarshal msg.Value: ", err)
					continue
				}

				fmt.Println("d data:", d.Data, " retrycnt:", d.RetryCount)

				hwm := consumer.HighWaterMarks()

				for k, v := range hwm {
					fmt.Println("\tk:", k, " v:", v)
				}

				producer, err := api.InitProducer(1, "multiplexed_consumer")
				d.RetryCount++
				if d.RetryCount > 5 {
					fmt.Println("max retry count reached")
					continue
				}
				if err := api.Publish(&[]string{msg.Topic}, &d, producer); err != nil {
					log.Println("could not publish: ", err)
				} // consumer.MarkOffset(msg, "") // mark message as processed

				// Please be aware that calling this function during an internal rebalance cycle may return
				// broker errors (e.g. sarama.ErrUnknownMemberId or sarama.ErrIllegalGeneration).
				// consumer.CommitOffsets()

			}
		case err, ok := <-consumer.Errors():
			if ok {
				log.Println("err:", err)
			}
		case <-signals:
			return
		}
	}
}
