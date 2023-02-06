package services

type SmsSender interface {
	Send(to, msg string) error
}

type EmailSender interface {
	Send(email, sub, body string) error
}
