package usecases

import (
	"hpfeed/domain"
	"time"
)

type MessageInteractor struct {
	newsRepo domain.NewsRepositoryInterface
}

func CreateNewMessageInteractor(newsRepo domain.NewsRepositoryInterface) *MessageInteractor {
	return &MessageInteractor{newsRepo: newsRepo}
}

func (this *MessageInteractor) GetAllMessages() []*Message {
	return convertAllToUsecaseMessages(this.newsRepo.GetAllMessages())
}

func (this *MessageInteractor) StoreNewMessages(messages []*Message) {
	newMessages := determineNewMessage(messages, this.newsRepo.GetLatestMessageDate())
	this.newsRepo.StoreAll(convertAllToDomainMessages(newMessages))
}

func determineNewMessage(allMessages []*Message, latestMessageDate time.Time) []*Message {
	newMessages := make([]*Message, 0)
	for _, message := range allMessages {
		if message.Date.After(latestMessageDate) {
			newMessages = append(newMessages, message)
		}
	}
	return newMessages
}

func convertToDomainMessage(message *Message) *domain.Message {
	return &domain.Message{Topic: message.Topic, Date: message.Date, Link: message.Link}
}

func convertToUsecaseMessage(message *domain.Message) *Message {
	return &Message{Topic: message.Topic, Date: message.Date, Link: message.Link}
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
