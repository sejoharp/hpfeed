package interfaces

import (
	"bitbucket.org/joscha/hpfeed/usecases"
)

type ForumReaderInterface interface {
	GetData() []byte
}

type FeedBuilderInterface interface {
	Generate(messages []*usecases.Message) []byte
}
