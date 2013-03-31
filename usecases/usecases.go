package usecases

import (
	"hpfeed/domain"
	"time"
)

type Message struct {
	Topic string
	Date  time.Time
	Link  string
	ID    string
}

type MessageInteractorInterface interface {
	GetAllMessages() []*Message
	StoreNewMessages(messages []*Message)
}

func convertToDomainMessage(message *Message) *domain.Message {
	return &domain.Message{Topic: message.Topic, Date: message.Date, Link: message.Link, ID: message.ID}
}

func convertToUsecaseMessage(message *domain.Message) *Message {
	return &Message{Topic: message.Topic, Date: message.Date, Link: message.Link, ID: message.ID}
}

func convertAllToUsecaseMessages(messages []*domain.Message) []*Message {
	usecaseMessages := make([]*Message, 0)
	for _, message := range messages {
		usecaseMessages = append(usecaseMessages, convertToUsecaseMessage(message))
	}
	return usecaseMessages
}

func convertAllToDomainMessages(messages []*Message) []*domain.Message {
	domainMessages := make([]*domain.Message, 0)
	for _, message := range messages {
		domainMessages = append(domainMessages, convertToDomainMessage(message))
	}
	return domainMessages
}
