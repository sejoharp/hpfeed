package domain

import (
	"hpfeed/helper"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

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

func (this *MongoDbNewsRepo) StoreAll(messages []*Message) {
	collection := this.getNewsCollection()
	defer collection.Database.Session.Close()

	for _, message := range messages {
		err := collection.Insert(message)
		helper.HandleFatalError("error inserting message:", err)
	}
}

func (this *MongoDbNewsRepo) GetLatestMessageDate() time.Time {
	collection := this.getNewsCollection()
	defer collection.Database.Session.Close()

	result := Message{}
	err := collection.Find(nil).Select(bson.M{"date": 1}).Sort("-date").One(&result)
	helper.HandleFatalError("error getting latest message:", err)
	return result.Date
}

func (this *MongoDbNewsRepo) GetAllMessages() []*Message {
	collection := this.getNewsCollection()
	defer collection.Database.Session.Close()

	result := make([]*Message, 0)
	err := collection.Find(nil).Sort("-date").All(&result)
	helper.HandleFatalError("error getting all messages:", err)
	return result
}
