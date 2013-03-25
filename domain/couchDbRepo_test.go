package domain

import (
couch "code.google.com/p/dsallings-couch-go"
"testing"
"time"
)

// EXPLORATION TESTS

func TestFetchingAllMessages(t *testing.T) {
	rows, err := getAllMessages()
	if err != nil {
		t.Errorf("fetching failed", err)
	}

	if amount := len(rows); amount != 4 {
		t.Errorf("fetching false amount, expected: %d, result: %d", 4, amount)
	}
	if rows[0].Id == "" {
		t.Errorf("fetching failed. Id must be filled, result: %s", rows[0].Id)
	}
	if rows[0].Rev == "" {
		t.Errorf("fetching failed. Rev must be filled, result: %s", rows[0].Rev)
	}
	if rows[0].ContentType == "" {
		t.Errorf("fetching failed. ContentType must be filled, result: %s", rows[0].ContentType)
	}
	if rows[0].Topic == "" {
		t.Errorf("fetching failed. Topic must be filled, result: %s", rows[0].Topic)
	}
	if rows[0].Date == "" {
		t.Errorf("fetching failed. Date must be filled, result: %s", rows[0].Date)
	}
	if rows[0].Link == "" {
		t.Errorf("fetching failed. Link must be filled, result: %s", rows[0].Link)
	}
}

func TestInsertAndRetrieve(t *testing.T) {
	conn, _ := couch.NewDatabase("localhost", "5984", "hpnews")
	msg := Message{Topic: "test", Date: "2012-01-10 14:10", Link: "www.test.de", ContentType: "news"}

	id, rev, err := conn.Insert(msg)
	if err != nil {
		t.Errorf("Insert failed", err)
	}

	result := new(Message)
	err = conn.Retrieve(id, result)
	if err != nil {
		t.Errorf("Retrieving failed", err)
	}
	if result.Id == "" {
		t.Errorf("fetching failed. Id must be filled, result: %s", result.Id)
	}
	if result.Rev == "" {
		t.Errorf("fetching failed. Rev must be filled, result: %s", result.Rev)
	}
	if result.ContentType == "" {
		t.Errorf("fetching failed. ContentType must be filled, result: %s", result.ContentType)
	}
	if result.Topic == "" {
		t.Errorf("fetching failed. Topic must be filled, result: %s", result.Topic)
	}
	if result.Date == "" {
		t.Errorf("fetching failed. Date must be filled, result: %s", result.Date)
	}
	if result.Link == "" {
		t.Errorf("fetching failed. Link must be filled, result: %s", result.Link)
	}
	err = conn.Delete(id, rev)
}

func TestLatestDate(t *testing.T) {
	date, err := GetLatestMessageDate()
	if err != nil {
		t.Errorf("Retrieving failed", err)
	}
	if expectedDate, _ := time.Parse("2006-01-02 15:04", "2013-01-10 14:21"); date != expectedDate {
		t.Errorf("false date, result: ", date)
	}
}


func TestFormatDate(t *testing.T) {
	date := time.Date(1990, 01, 01, 0, 0, 0, 0, time.UTC)
	expected := "01 Jan 90 00:00 UTC"
	if dateString := formatDate(date); dateString != expected {
		t.Errorf("formatting date failed, expected: %s result: %s", expected, dateString)
	}
}

