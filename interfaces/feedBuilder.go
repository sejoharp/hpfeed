package interfaces

import (
	"github.com/baliw/moverss"
	"hpfeed/usecases"
	"strconv"
)

type FeedBuilder struct {
	updateInterval int
}

func CreateNewFeedBuilder(updateInterval int) *FeedBuilder {
	return &FeedBuilder{updateInterval: updateInterval}
}

func (this *FeedBuilder) Generate(messages []*usecases.Message) []byte {
	channel := moverss.ChannelFactory("Hamburg Privateers", "http://www.kickern-hamburg.de/phpBB2/viewforum.php?f=15", "Hamburg Privateers feed")
	channel.SetTTL(strconv.Itoa(this.updateInterval))
	for _, message := range messages {
		description := "<a href=" + message.Link + "> ..zur Nachricht</a>"
		item := &moverss.Item{Title: message.Topic, Link: message.Link, Description: description}
		item.SetPubDate(message.Date)
		channel.AddItem(item)
	}
	return channel.Publish()
}
