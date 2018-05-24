### PubSub using Kafka in Go


Kafka uses Zookeeper to mange following tasks

* Electing a controller - The controller is one of the brokers and is responsible for maintaining the leader/follower relationship for all the partitions. When a node shuts down, it is the controller that tells other replicas to become partition leaders to replace the partition leaders on the node that is going away. Zookeeper is used to elect a controller, make sure there is only one and elect a new one it if it crashes.
* Cluster membership - Which brokers are alive and part of the cluster? this is also managed through ZooKeeper.
* Topic configuration - Which topics exist, how many partitions each has, where are the replicas, who is the preferred leader, what configuration overrides are set for each topic
* Manage Quotas - How much data is each client allowed to read and write
* Access control - Who is allowed to read and write to which topic (old high level consumer). Which consumer groups exist, who are their members and what is the latest offset each group got from each partition.


KAFKA_ADVERTISED_HOST_NAME is the IP address of the machine(my local machine) which Kafka container running. ZOOKEEPER_IP is the Zookeeper container running machines IP. By this way producers and consumers can access Kafka and Zookeeper by using that IP address.

Same producer can publish to multiple topics

Consumer group is a set of consumers which has a unique group id. 
Each consumer group is a subscriber to one or more kafka topics. 
Each consumer group maintains its offset per topic partition. 
A record gets delivered to only one consumer in a consumer group. 
Each consumer in a consumer group processes records and only one consumer in that group will get the same record. 
Consumers in a consumer group load balance record processing.

#### Steps

* To check which IP addresses to use

`ifconfig | grep inet`

* Run Zookeeper and Kafka

`docker-compose up`

* Create topic

`docker run --rm ches/kafka kafka-topics.sh --create --zookeeper 192.168.4.93:2181 --replication-factor 1 --partitions 1 --topic senz`

* List topics

`docker run --rm ches/kafka kafka-topics.sh --list --zookeeper 192.168.4.93:2181`

* Create publisher

`docker run --rm --interactive ches/kafka kafka-console-producer.sh --broker-list 192.168.4.93:9092 --topic senz --timeout 2000`
- - creates an interactive pub, taking inp from commandline and pub them to kafka
- - `broker-list` specifies the host and port of kafka

* create consumer

<!-- `docker run --rm ches/kafka kafka-console-consumer.sh --topic senz --from-beginning --zookeeper 192.168.4.93:2181` -->
`docker run --rm ches/kafka kafka-console-consumer.sh --bootstrap-server 192.168.4.93:9092 --topic senz --from-beginning`

##### Nuances
- set `KAFKA_ADVERTISED_HOST_NAME` in `docker-compose.yml` if producer cannot publish.


* Tutorials:

[Kafka and Zookeeper with Docker](https://medium.com/@itseranga/kafka-and-zookeeper-with-docker-65cff2c2c34f)
[Kafka consumer](https://medium.com/@itseranga/kafka-consumer-with-golang-a93db6131ac2)
[Kafka producer](https://medium.com/@itseranga/kafka-producer-with-golang-fab7348a5f9a)