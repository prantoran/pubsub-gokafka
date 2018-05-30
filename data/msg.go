package data

type KafkaMsg struct {
	Data       interface{}
	RetryCount int
}

func NewKafkaMsg(data interface{}) *KafkaMsg {
	return &KafkaMsg{
		Data:       data,
		RetryCount: 0,
	}
}
