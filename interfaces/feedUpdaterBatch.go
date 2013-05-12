package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bitbucket.org/joscha/hpfeed/usecases"
	"math/rand"
	"time"
)

type FeedUpdater struct {
	updateInterval int
	service        usecases.MessageInteractorInterface
	forumReader    ForumReaderInterface
	ticker         *time.Ticker
}

func CreateNewFeedUpdater(updateInterval int, service usecases.MessageInteractorInterface, forumReader ForumReaderInterface) *FeedUpdater {
	return &FeedUpdater{updateInterval: updateInterval, service: service, forumReader: forumReader}
}

func (this *FeedUpdater) StopFeedUpdateCycle() {
	this.ticker.Stop()
}

func (this *FeedUpdater) StartFeedUpdateCycle() {
	this.updateFeedData()
	duration := time.Duration(rand.Intn(this.updateInterval)) * time.Minute
	this.ticker = time.NewTicker(duration)
	go func() {
		for _ = range this.ticker.C {
			this.updateFeedData()
		}
	}()
}

func (this *FeedUpdater) updateFeedData() {
	if this.forumReader.IsAvailable() {
		rawData := this.forumReader.GetData()
		doc := GenerateDocument(rawData)
		threads := ParseThreads(doc)
		this.service.StoreNewMessages(convertAllThreadsToMesssages(threads))
	} else {
		helper.LogError("The forum is offline.")
	}
}

func convertThreadToMessage(thread *Thread) *usecases.UcMessage {
	return &usecases.UcMessage{Date: thread.Date, Link: thread.Link, Topic: thread.Topic}
}

func convertAllThreadsToMesssages(threads []*Thread) []*usecases.UcMessage {
	messages := make([]*usecases.UcMessage, 0)
	for _, thread := range threads {
		messages = append(messages, convertThreadToMessage(thread))
	}
	return messages
}
