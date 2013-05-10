package interfaces

import (
	"bitbucket.org/joscha/hpfeed/usecases"
)

type ForumReaderInterface interface {
	GetData() []byte
	IsAvailable() bool 
}

type FeedBuilderInterface interface {
	Generate(messages []*usecases.UcMessage) []byte
}
