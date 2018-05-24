### PubSub using Kafka in Go


Kafka uses Zookeeper to mange following tasks

* Electing a controller - The controller is one of the brokers and is responsible for maintaining the leader/follower relationship for all the partitions. When a node shuts down, it is the controller that tells other replicas to become partition leaders to replace the partition leaders on the node that is going away. Zookeeper is used to elect a controller, make sure there is only one and elect a new one it if it crashes.
* Cluster membership - Which brokers are alive and part of the cluster? this is also managed through ZooKeeper.
* Topic configuration - Which topics exist, how many partitions each has, where are the replicas, who is the preferred leader, what configuration overrides are set for each topic
* Manage Quotas - How much data is each client allowed to read and write
* Access control - Who is allowed to read and write to which topic (old high level consumer). Which consumer groups exist, who are their members and what is the latest offset each group got from each partition.


KAFKA_ADVERTISED_HOST_NAME is the IP address of the machine(my local machine) which Kafka container running. ZOOKEEPER_IP is the Zookeeper container running machines IP. By this way producers and consumers can access Kafka and Zookeeper by using that IP address.

#### Steps
* Run Zookeeper and Kafka

`docker-compose up`

* Create topic

`docker run --rm ches/kafka kafka-topics.sh --create --topic senz --replication-factor 1 --partitions 1 --zookeeper 192.168.4.93:2181`


* Tutorials:

[Kafka and Zookeeper with Docker](https://medium.com/@itseranga/kafka-and-zookeeper-with-docker-65cff2c2c34f)