package domain

import (
	"time"
)

type NewsRepositoryInterface interface {
	StoreAll(messages []*DbMessage)
	GetLatestMessageDate() time.Time
	GetAllMessages() []*DbMessage
}

type DbMessage struct {
	Topic string
	Date  time.Time
	Link  string
	ID    string
}
