package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/wvanbergen/kafka/consumergroup"
)

const (
	// zookeeperConn = "10.4.1.29:2181"
	zookeeperConn = "192.168.4.93:2181"
)

func parseFlags() {
	flag.String("topics", "senz", "help message for flagname")
	flag.String("cg", "zgroup", "Consumer group for the client")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func getTopics() []string {
	tf := viper.GetString("topics")
	return strings.Split(tf, ",")
}

func main() {
	parseFlags()

	cgName := viper.GetString("cg")
	fmt.Println("cgName:", cgName)

	topics := getTopics()

	for i, u := range topics {
		fmt.Println("topic[", i, "]: ", u)
	}

	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// init consumer croup
	cg, err := initConsumer(&topics, cgName)
	if err != nil {
		fmt.Println("Error consumer goup: ", err.Error())
		os.Exit(1)
	}
	defer cg.Close()

	// run consumer
	consume(cg)
}

func initConsumer(topics *[]string, cgroup string) (*consumergroup.ConsumerGroup, error) {
	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup(cgroup, *topics, []string{zookeeperConn}, config)
	if err != nil {
		return nil, err
	}

	return cg, err
}

func consume(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			// messages coming through chanel
			// only take messages from subscribed topic

			fmt.Println("msg\ntopic:", msg.Topic,
				"\nkey:", string(msg.Key),
				"\nval:", string(msg.Value),
				"\noffset:", msg.Offset,
				"\npartition:", msg.Partition,
				"\ntimestamp:", msg.Timestamp,
				"\nblocktimestamp:", msg.BlockTimestamp,
				"\nheaders:", msg.Headers)
			fmt.Println("Value: ", string(msg.Value))

			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}
