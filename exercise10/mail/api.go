package mail
import "github.com/sirupsen/logrus"

type Message struct {
	To string
	From string
	Subject string
	Body string 
}

type Sender interface {
	Send(message Message) error
}