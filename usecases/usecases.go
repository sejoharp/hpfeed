package usecases

import (
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
