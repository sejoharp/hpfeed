package usecases

import (
	"bitbucket.org/joscha/hpfeed/domain"
	"time"
)

type UcMessage struct {
	Topic string
	Date  time.Time
	Link  string
	ID    string
}

type MessageInteractorInterface interface {
	GetAllMessages() []*UcMessage
	StoreNewMessages(messages []*UcMessage)
}

func convertToDomainMessage(message *UcMessage) *domain.DbMessage {
	return &domain.DbMessage{Topic: message.Topic, Date: message.Date, Link: message.Link, ID: message.ID}
}

func convertToUsecaseMessage(message *domain.DbMessage) *UcMessage {
	return &UcMessage{Topic: message.Topic, Date: message.Date, Link: message.Link, ID: message.ID}
}

func convertAllToUsecaseMessages(messages []*domain.DbMessage) []*UcMessage {
	usecaseMessages := make([]*UcMessage, 0)
	for _, message := range messages {
		usecaseMessages = append(usecaseMessages, convertToUsecaseMessage(message))
	}
	return usecaseMessages
}

func convertAllToDomainMessages(messages []*UcMessage) []*domain.DbMessage {
	domainMessages := make([]*domain.DbMessage, 0)
	for _, message := range messages {
		domainMessages = append(domainMessages, convertToDomainMessage(message))
	}
	return domainMessages
}
