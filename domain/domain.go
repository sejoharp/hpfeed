package domain

import (
	"time"
)

type NewsRepositoryInterface interface {
	StoreAll(messages []*Message)
	GetLatestMessageDate() time.Time
	GetAllMessages() []*Message
}

type Message struct {
	Topic string
	Date  time.Time
	Link  string
}
