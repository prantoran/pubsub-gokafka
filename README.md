### PubSub using Kafka in Go


Kafka uses Zookeeper to mange following tasks

* Electing a controller - The controller is one of the brokers and is responsible for maintaining the leader/follower relationship for all the partitions. When a node shuts down, it is the controller that tells other replicas to become partition leaders to replace the partition leaders on the node that is going away. Zookeeper is used to elect a controller, make sure there is only one and elect a new one it if it crashes.
* Cluster membership - Which brokers are alive and part of the cluster? this is also managed through ZooKeeper.
* Topic configuration - Which topics exist, how many partitions each has, where are the replicas, who is the preferred leader, what configuration overrides are set for each topic
* Manage Quotas - How much data is each client allowed to read and write
* Access control - Who is allowed to read and write to which topic (old high level consumer). Which consumer groups exist, who are their members and what is the latest offset each group got from each partition.


#### Steps
* Running Zookeeper
`docker run -d --name zookeeper -p 2181:2181 jplock/zookeeper`
* Running Kafka
`docker run -d --name kafka -p 7203:7203 -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=10.4.1.29 -e ZOOKEEPER_IP=10.4.1.29 ches/kafka`


* Tutorials:

[Kafka and Zookeeper with Docker](https://medium.com/@itseranga/kafka-and-zookeeper-with-docker-65cff2c2c34f)