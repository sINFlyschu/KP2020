package client

import mailSender "github.com/sINFlyschu/KP2020/exercise10/mail"

func main() {
	mailSender := mailSender.SenderImpl{}

	mailSender.Send(Message{
		To: "email@email.de",
		From: "Mail@mail.de",
		Subject: "Hello",
		Body: "mail mail mail",
	})
}