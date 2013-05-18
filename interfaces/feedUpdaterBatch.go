package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bitbucket.org/joscha/hpfeed/usecases"
	"math/rand"
	"time"
)

type FeedUpdater struct {
	updateInterval int
	interactor     usecases.MessageInteractorInterface
	forumReader    ForumReaderInterface
	ticker         *time.Ticker
	parser         ParserInterface
}

func CreateNewFeedUpdater(updateInterval int, interactor usecases.MessageInteractorInterface, forumReader ForumReaderInterface, parser ParserInterface) *FeedUpdater {
	return &FeedUpdater{updateInterval: updateInterval, interactor: interactor, forumReader: forumReader, parser: parser}
}

func (this *FeedUpdater) Stop() {
	this.ticker.Stop()
}

func (this *FeedUpdater) Start() {
	this.updateFeedData()
	duration := time.Duration(rand.Intn(this.updateInterval)) * time.Second
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
		doc := this.parser.GenerateDocument(rawData)
		threads := this.parser.ParseThreads(doc)
		this.interactor.StoreNewMessages(convertAllThreadsToMesssages(threads))
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
