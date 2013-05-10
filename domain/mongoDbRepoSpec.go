package domain

import (
	. "github.com/ghthor/gospec"
	"labix.org/v2/mgo/bson"
	"time"
)

// EXPLORATION TESTS
func MongoDbRepoSpec(c Context) {
	c.Specify("It stores a message to db.", func() {
		date := time.Date(1990, 01, 01, 1, 0, 0, 0, time.Local)
		link := "www.test.de/test?param=test"
		topic := "testtopic"
		news := &DbMessage{Date: date, Link: link, Topic: topic}
		dataAccess := CreateMongoDbNewsRepo("localhost", "hpfeed")
		array := make([]*DbMessage, 0)
		array = append(array, news)
		dataAccess.StoreAll(array)

		result := News{}
		dataAccess.getNewsCollection().Find(bson.M{"topic": "testtopic"}).One(&result)
		c.Expect(result.Topic, Equals, topic)
		c.Expect(result.Link, Equals, link)
		c.Expect(result.Date, Equals, date)
	})
	c.Specify("It fetches the date of the latest message.", func() {
		dataAccess := CreateMongoDbNewsRepo("localhost", "hpfeed")

		date := time.Now()
		newsArray := make([]*DbMessage, 0)
		newsArray = append(newsArray, &DbMessage{Date: time.Date(1990, 01, 01, 1, 0, 0, 0, time.Local), Link: "www.test.de/test?param=test", Topic: "testtopic"})
		newsArray = append(newsArray, &DbMessage{Date: date, Link: "www.test.de/test?param=test", Topic: "testtopic"})
		dataAccess.StoreAll(newsArray)
		latestDate := dataAccess.GetLatestMessageDate()
		c.Expect(latestDate.Format(time.RFC3339), Equals, date.Format(time.RFC3339))
	})
	c.Specify("It fetches all messages.", func() {
		dataAccess := CreateMongoDbNewsRepo("localhost", "hpfeed")
		dataAccess.getNewsCollection().RemoveAll(nil)
		date := time.Now()
		newsExpected := make([]*DbMessage, 0)
		newsExpected = append(newsExpected, &DbMessage{Date: time.Date(1990, 01, 01, 1, 0, 0, 0, time.Local), Link: "www.test.de/test?param=test", Topic: "testtopic"})
		newsExpected = append(newsExpected, &DbMessage{Date: date, Link: "www.test.de/test?param=test", Topic: "testtopic"})
		dataAccess.StoreAll(newsExpected)

		newsResult := dataAccess.GetAllMessages()
		c.Expect(len(newsResult), Equals, len(newsExpected))
	})
}
