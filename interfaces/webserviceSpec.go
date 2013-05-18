package interfaces

import (
	"bitbucket.org/joscha/hpfeed/usecases"
	. "github.com/ghthor/gospec"
)

type FeedBuilderMock struct{}

func (this *FeedBuilderMock) Generate(messages []*usecases.UcMessage) []byte {
	return []byte("myMock")
}

type InteractorMock struct{}

func (this *InteractorMock) GetAllMessages() []*usecases.UcMessage {
	return make([]*usecases.UcMessage, 0)
}
func (this *InteractorMock) StoreNewMessages(messages []*usecases.UcMessage) {}

func WebserviceSpec(c Context) {
	c.Specify("webservice is stoppable.", func() {

		webservice := CreateNewWebservice(&InteractorMock{}, &FeedBuilderMock{}, 8000, "/news")
		webservice.Start()
		c.Expect(webservice.Stop(), Equals, nil)
	})
}
