package mail

import log "github.com/sirupsen/logrus"

type SenderImpl struct {}

func (s SenderImpl) Send(message Message) error {

	log.WithFields(log.Fields{
		"from":message.From,
		"to":message.To,
		"size":len(message.Body),
		}).Info("Sending message....")
	return nil
}