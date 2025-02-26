package mail

type From struct {
	Address string
	Name    string
}

type Email struct {
	From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    string
	HTML    string
}

type Driver interface {
	Send(email Email) error
}
