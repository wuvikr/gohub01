package sms

import "sync"

type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

type SMS struct {
	Driver
}

var (
	once        sync.Once
	internalSMS *SMS
)

func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})

	return internalSMS
}

func (s *SMS) SendMessage(phone string, message Message) bool {
	return s.Send(phone, message)
}
