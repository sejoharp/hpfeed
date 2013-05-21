package interfaces

import (
	"bitbucket.org/joscha/hpfeed/usecases"
	"github.com/baliw/moverss"
	"strconv"
	"time"
)

type FeedBuilder struct {
	updateInterval int
}

func CreateNewFeedBuilder(updateInterval int) *FeedBuilder {
	return &FeedBuilder{updateInterval: updateInterval}
}

func (this *FeedBuilder) Generate(messages []*usecases.UcMessage) []byte {
	channel := moverss.ChannelFactory("Hamburg Privateers", "http://www.kickern-hamburg.de/phpBB2/viewforum.php?f=15", "Hamburg Privateers feed")
	channel.SetTTL(strconv.Itoa(this.updateInterval/120))
	for _, message := range messages {
		item := createItem(message)
		channel.AddItem(item)
	}
	return channel.Publish()
}

func createItem(message *usecases.UcMessage) *moverss.Item {
	item := &moverss.Item{
		Title:   message.Topic,
		Link:    message.Link,
		Guid:    message.Link,
		PubDate: message.Date.UTC().Format(time.RFC822)}
	return item
}
