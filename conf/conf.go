package conf

const (
	kafka1 = "192.168.4.93:9092"
	kafka2 = "192.168.4.93:9093"
	kafka3 = "192.168.4.93:9094"
)

func AllBrokers(flag int) []string {
	ret := []string{}
	if flag&1 == 1 {
		ret = append(ret, kafka1)
	}
	if flag&2 == 2 {
		ret = append(ret, kafka2)
	}
	if flag&4 == 4 {
		ret = append(ret, kafka3)
	}
	return ret
}
