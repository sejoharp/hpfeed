package domain

import (
	"bitbucket.org/joscha/hpfeed/helper"
	couch "code.google.com/p/dsallings-couch-go"
	"strconv"
	"time"
)

type QueryResult struct {
	Rows []struct {
		Id  string `json:"_id"`
		Key string
		Doc CouchDBMessage
	}
}

type CouchDBMessage struct {
	Id          string `json:"_id"`
	Rev         string `json:"_rev"`
	ContentType string `json:"type"`
	Topic       string `json:"topic"`
	Date        string `json:"date"`
	Link        string `json:"link"`
}

type CouchDbRepo struct {
	dbhost     string
	dbport     string
	dbname     string
	dbuser     string
	dbpassword string
}

func CreateCouchDbRepo(dbhost, dbport, dbname, dbuser, dbpassword string) *CouchDbRepo {
	return &CouchDbRepo{dbhost, dbport, dbname, dbuser, dbpassword}
}

func (this *CouchDbRepo) getConnection() *couch.Database {
	dburl := "http://" + this.dbuser + ":" + this.dbpassword + "@" + this.dbhost + ":" + this.dbport + "/" + this.dbname
	conn, err := couch.Connect(dburl)
	helper.HandleFatalError("db connection error:", err)
	return &conn
}

func (this *CouchDbRepo) GetAllMessages() []*DbMessage {
	var result QueryResult

	err := this.getConnection().Query("_design/messages/_view/by_date", map[string]interface{}{"include_docs": true, "descending": true}, &result)
	helper.HandleFatalError("query all messages failed:", err)

	messages := make([]*DbMessage, 0)
	for i, _ := range result.Rows {
		messages = append(messages, convertCouchDbToMessage(&result.Rows[i].Doc))
	}
	return messages
}

func (this *CouchDbRepo) GetLatestMessageDate() time.Time {
	var result QueryResult

	err := this.getConnection().Query("_design/messages/_view/by_date", map[string]interface{}{"limit": 1, "descending": true}, &result)
	helper.HandleFatalError("getting latest message date failed:", err)
	if len(result.Rows) == 0 {
		return time.Date(1990, 01, 01, 0, 0, 0, 0, time.UTC)
	}
	unixtimestamp, err := strconv.ParseInt(result.Rows[0].Key, 10, 64)
	helper.HandleFatalError("parsing last date from db failed:", err)
	return time.Unix(unixtimestamp, 0)
}

func (this *CouchDbRepo) StoreAll(messages []*DbMessage) {
	couchDbMessageList := convertAllMessagesToCouchDb(messages)
	for _, couchDbMessage := range couchDbMessageList {
		couchDbMessage.ContentType = "news"
		_, _, err := this.getConnection().Insert(couchDbMessage)
		helper.HandleFatalError("error inserting couchDbMessage:", err)
	}
}

func convertMessageToCouchDb(message *DbMessage) *CouchDBMessage {
	if message.ID == "" {
		return &CouchDBMessage{
			Topic: message.Topic,
			Date:  strconv.FormatInt(message.Date.Unix(), 10),
			Link:  message.Link}
	}
	return &CouchDBMessage{
		Topic: message.Topic,
		Date:  strconv.FormatInt(message.Date.Unix(), 10),
		Link:  message.Link,
		Id:    message.ID}
}

func convertAllMessagesToCouchDb(messages []*DbMessage) []*CouchDBMessage {
	couchDbMessageList := make([]*CouchDBMessage, 0)
	for _, message := range messages {
		couchDbMessageList = append(couchDbMessageList, convertMessageToCouchDb(message))
	}
	return couchDbMessageList
}

func convertCouchDbToMessage(couchDbMessage *CouchDBMessage) *DbMessage {
	return &DbMessage{
		Topic: couchDbMessage.Topic,
		Date:  time.Unix(stringToInt64(couchDbMessage.Date), 0),
		Link:  couchDbMessage.Link,
		ID:    couchDbMessage.Id}
}

func convertAllCouchDbToMesssages(couchDbMessageList []*CouchDBMessage) []*DbMessage {
	messages := make([]*DbMessage, 0)
	for _, couchDbMessage := range couchDbMessageList {
		messages = append(messages, convertCouchDbToMessage(couchDbMessage))
	}
	return messages
}

func stringToInt64(raw string) int64 {
	result, err := strconv.ParseInt(raw, 10, 64)
	helper.HandleFatalError("parsing string to int64 failed:", err)
	return result
}
