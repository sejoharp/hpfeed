package domain

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type News struct {
	Topic string
	Date  time.Time
	Link  string
	ID    bson.ObjectId `bson:"_id,omitempty"`
}

type MongoDbNewsRepo struct {
	dbhost string
	dbname string
}

func CreateMongoDbNewsRepo(dbhost string, dbname string) *MongoDbNewsRepo {
	return &MongoDbNewsRepo{dbhost: dbhost, dbname: dbname}
}

func (this *MongoDbNewsRepo) getNewsCollection() *mgo.Collection {
	session, err := mgo.Dial(this.dbhost)
	helper.HandleFatalError("error getting connection to mongodb:", err)
	return session.DB(this.dbname).C("news")
}

func (this *MongoDbNewsRepo) StoreAll(messages []*DbMessage) {
	collection := this.getNewsCollection()
	defer collection.Database.Session.Close()
	newsList := convertAllMessagesToNews(messages)
	for _, news := range newsList {
		err := collection.Insert(news)
		helper.HandleFatalError("error inserting news:", err)
	}
}

func (this *MongoDbNewsRepo) GetLatestMessageDate() time.Time {
	collection := this.getNewsCollection()
	defer collection.Database.Session.Close()

	result := News{}
	err := collection.Find(nil).Select(bson.M{"date": 1}).Sort("-date").One(&result)
	helper.HandleFatalError("error getting latest News:", err)
	return result.Date
}

func (this *MongoDbNewsRepo) GetAllMessages() []*DbMessage {
	collection := this.getNewsCollection()
	defer collection.Database.Session.Close()

	result := make([]*News, 0)
	err := collection.Find(nil).Sort("-date").All(&result)
	helper.HandleFatalError("error getting all News:", err)
	return convertAllNewsToMesssages(result)
}

func messageToNews(message *DbMessage) *News {
	if message.ID == "" {
		return &News{Topic: message.Topic, Date: message.Date, Link: message.Link}
	}
	return &News{
		Topic: message.Topic,
		Date:  message.Date,
		Link:  message.Link,
		ID:    bson.ObjectIdHex(message.ID)}
}

func convertAllMessagesToNews(messages []*DbMessage) []*News {
	newsList := make([]*News, 0)
	for _, message := range messages {
		newsList = append(newsList, messageToNews(message))
	}
	return newsList
}

func newsToMessage(news *News) *DbMessage {
	return &DbMessage{
		Topic: news.Topic,
		Date:  news.Date,
		Link:  news.Link,
		ID:    news.ID.Hex()}
}

func convertAllNewsToMesssages(newsList []*News) []*DbMessage {
	messages := make([]*DbMessage, 0)
	for _, news := range newsList {
		messages = append(messages, newsToMessage(news))
	}
	return messages
}
