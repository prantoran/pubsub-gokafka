docker-compose exec kafka1 kafka-topics --create --topic senz --partitions 2 --replication-factor 1 --if-not-exists --zookeeper 192.168.4.93:2181

docker-compose exec kafka1 kafka-topics --create --topic renz --partitions 3 --replication-factor 1 --if-not-exists --zookeeper 192.168.4.93:2181

docker run --rm ches/kafka kafka-topics.sh --list --zookeeper 192.168.4.93:2181