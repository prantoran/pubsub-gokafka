package conf

const (
	kafka1 = "192.168.4.93:9092"
	kafka2 = "192.168.4.93:9093"
	kafka3 = "192.168.4.93:9094"
)

func AllBrokers() []string {
	return []string{kafka1, kafka2, kafka3}
}
