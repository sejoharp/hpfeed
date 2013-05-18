package interfaces

import (
	"bytes"
	"code.google.com/p/go.net/html"
	. "github.com/ghthor/gospec"
	"github.com/puerkitobio/goquery"
	"time"
)

type ForumReaderMock struct{}

func (this *ForumReaderMock) GetData() []byte   { return []byte("myMock") }
func (this *ForumReaderMock) IsAvailable() bool { return true }

type ParserMock struct{}

func (this *ParserMock) GenerateDocument(rawData []byte) *goquery.Document {
	node, _ := html.Parse(bytes.NewReader([]byte("myMock")))
	return goquery.NewDocumentFromNode(node)
}
func (this *ParserMock) ParseThreads(doc *goquery.Document) []*Thread { return make([]*Thread, 0) }

func FeedUpdaterBatchSpec(c Context) {
	c.Specify("UpdateCycle is stoppable.", func() {
		forumUpdater := CreateNewFeedUpdater(2, &InteractorMock{}, &ForumReaderMock{}, &ParserMock{})
		forumUpdater.Start()
		forumUpdater.Stop()
		time.Sleep(2 * time.Second)
		select {
		case <-forumUpdater.ticker.C:
			c.Expect("ticker did not shut down", Equals, "")
		default:
			c.Expect("ticker did shut down", Equals, "ticker did shut down")
		}
	})
}
