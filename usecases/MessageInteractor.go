package usecases

import (
	"bitbucket.org/joscha/hpfeed/domain"
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
