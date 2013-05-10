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

func (this *MessageInteractor) GetAllMessages() []*UcMessage {
	return convertAllToUsecaseMessages(this.newsRepo.GetAllMessages())
}

func (this *MessageInteractor) StoreNewMessages(messages []*UcMessage) {
	newMessages := determineNewMessage(messages, this.newsRepo.GetLatestMessageDate())
	this.newsRepo.StoreAll(convertAllToDomainMessages(newMessages))
}

func determineNewMessage(allMessages []*UcMessage, latestMessageDate time.Time) []*UcMessage {
	newMessages := make([]*UcMessage, 0)
	for _, message := range allMessages {
		if message.Date.After(latestMessageDate) {
			newMessages = append(newMessages, message)
		}
	}
	return newMessages
}
