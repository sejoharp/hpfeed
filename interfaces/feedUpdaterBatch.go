package interfaces

import (
	"bitbucket.org/joscha/hpfeed/usecases"
	"math/rand"
	"time"
)

type FeedUpdater struct {
	updateInterval int
	service        usecases.MessageInteractorInterface
	forumReader    ForumReaderInterface
}

func CreateNewFeedUpdater(updateInterval int, service usecases.MessageInteractorInterface, forumReader ForumReaderInterface) *FeedUpdater {
	return &FeedUpdater{updateInterval: updateInterval, service: service, forumReader: forumReader}
}

func (this *FeedUpdater) StartFeedUpdateCycle() {
	this.updateFeedData()
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(this.updateInterval)) * time.Minute)
			this.updateFeedData()
		}
	}()
}

func (this *FeedUpdater) updateFeedData() {
	rawData := this.forumReader.GetData()
	doc := GenerateDocument(rawData)
	threads := ParseThreads(doc)
	this.service.StoreNewMessages(convertAllThreadsToMesssages(threads))
}

func convertThreadToMessage(thread *Thread) *usecases.Message {
	return &usecases.Message{Date: thread.Date, Link: thread.Link, Topic: thread.Topic}
}

func convertAllThreadsToMesssages(threads []*Thread) []*usecases.Message {
	messages := make([]*usecases.Message, 0)
	for _, thread := range threads {
		messages = append(messages, convertThreadToMessage(thread))
	}
	return messages
}
