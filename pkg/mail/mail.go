package mail

import "sync"

type Mailer struct {
	Driver Driver
}

var (
	once           sync.Once
	internalMailer *Mailer
)

func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})

	return internalMailer
}

func (m *Mailer) Send(email Email) error {
	return m.Driver.Send(email)
}
