package domain

import (
	"bitbucket.org/joscha/hpfeed/helper"
	couch "code.google.com/p/dsallings-couch-go"
	"time"
)

type Result struct {
	Rows []MessageResult
}

type MessageResult struct {
	Id  string `json:"_id"`
	Key string
	Doc MessageRow
}

type MessageRow struct {
	Id          string `json:"_id"`
	Rev         string `json:"_rev"`
	ContentType string `json:"type"`
	Topic       string `json:"topic"`
	Date        string `json:"date"`
	Link        string `json:"link"`
}

type LatestDate struct {
	Key string
}

type CouchDbNewsRepo struct {
	conn *couch.Database
}

func newCouchDbNewsRepo(dbhost string, dbport string, dbname string) *CouchDbNewsRepo {
	conn, err := couch.NewDatabase(dbhost, dbport, dbname)
	helper.HandleFatalError("db connection error:", err)
	return &CouchDbNewsRepo{conn: &conn}
}

func (this *CouchDbNewsRepo) getAllMessages() []*MessageRow {
	result := Result{}
	err := this.conn.Query("_design/messages/_view/by_date", map[string]interface{}{"include_docs": true, "descending": true}, &result)
	helper.HandleFatalError("query all messages failed:", err)

	messages := make([]*MessageRow, 0)
	for i, _ := range result.Rows {
		messages = append(messages, &result.Rows[i].Doc)
	}
	return messages
}

func (this *CouchDbNewsRepo) getLatestMessageDate() time.Time {
	var result struct {
		Rows []struct {
			Key string
		}
	}
	err := this.conn.Query("_design/messages/_view/by_date", map[string]interface{}{"limit": 1, "descending": true}, &result)
	helper.HandleFatalError("query query messages_by_date failed:", err)
	var dateResult time.Time
	if len(result.Rows) == 0 {
		return time.Date(1990, 01, 01, 0, 0, 0, 0, time.UTC)
	}
	dateResult, err = time.Parse(time.RFC3339, result.Rows[0].Key)
	helper.HandleFatalError("parsing last date from db failed:", err)
	return dateResult
}

func (this *CouchDbNewsRepo) saveMessage(message *MessageRow) (string, string, error) {
	message.ContentType = "news"
	return this.conn.Insert(message)
}

func (this *CouchDbNewsRepo) saveNewMessages(connection *couch.Database, messages []*MessageRow) {
	for _, message := range messages {
		message.ContentType = "news"
		_, _, err := this.conn.Insert(message)
		helper.HandleFatalError("error inserting message:", err)
	}
}
