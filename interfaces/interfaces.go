package interfaces

import (
	"bitbucket.org/joscha/hpfeed/usecases"
	"github.com/puerkitobio/goquery"
)

type ForumReaderInterface interface {
	GetData() []byte
	IsAvailable() bool
}

type FeedBuilderInterface interface {
	Generate(messages []*usecases.UcMessage) []byte
}

type ParserInterface interface {
	GenerateDocument(rawData []byte) *goquery.Document
	ParseThreads(doc *goquery.Document) []*Thread
}
