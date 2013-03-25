package domain

import (
	"labix.org/v2/mgo/bson"
	"testing"
	"time"
)

func TestSaveAllNews(t *testing.T) {
	date := time.Date(1990, 01, 01, 1, 0, 0, 0, time.Local)
	link := "www.test.de/test?param=test"
	topic := "testtopic"
	news := &News{Date: date, Link: link, Topic: topic}
	dataAccess := createDataAccess("localhost", "hpfeed")
	array := make([]*News, 0)
	array = append(array, news)
	dataAccess.saveAllNews(array)

	result := News{}
	dataAccess.newsCollection.Find(bson.M{"topic": "testtopic"}).One(&result)
	if result.Topic != topic {
		t.Errorf("loading saved data failed, expected: %s result: %s", topic, result.Topic)
	}
	if result.Link != link {
		t.Errorf("loading saved data failed, expected: %s result: %s", link, result.Link)
	}
	if result.Date != date {
		t.Errorf("loading saved data failed, expected: %s result: %s", date, result.Date)
	}
}

func TestGetLatestNewsDate(t *testing.T) {
	dataAccess := createDataAccess("localhost", "hpfeed")

	date := time.Now()
	newsArray := make([]*News, 0)
	newsArray = append(newsArray, &News{Date: time.Date(1990, 01, 01, 1, 0, 0, 0, time.Local), Link: "www.test.de/test?param=test", Topic: "testtopic"})
	newsArray = append(newsArray, &News{Date: date, Link: "www.test.de/test?param=test", Topic: "testtopic"})
	dataAccess.saveAllNews(newsArray)
	latestDate := dataAccess.getLatestNewsDate()
	if latestDate.Format(time.RFC3339) != date.Format(time.RFC3339) {
		t.Errorf("loading latest Date failed, expected: %s result: %s", date, latestDate)
	}
}

func TestGetAllNews(t *testing.T) {
	dataAccess := createDataAccess("localhost", "hpfeed")
	dataAccess.newsCollection.RemoveAll(nil)
	date := time.Now()
	newsExpected := make([]*News, 0)
	newsExpected = append(newsExpected, &News{Date: time.Date(1990, 01, 01, 1, 0, 0, 0, time.Local), Link: "www.test.de/test?param=test", Topic: "testtopic"})
	newsExpected = append(newsExpected, &News{Date: date, Link: "www.test.de/test?param=test", Topic: "testtopic"})
	dataAccess.saveAllNews(newsExpected)

	newsResult := dataAccess.getAllNews()
	if len(newsResult) != len(newsExpected) {
		t.Errorf("loading allNews failed, expected: %s result: %s", newsExpected, newsResult)
	}
}
