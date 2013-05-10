package domain

import (
	. "github.com/ghthor/gospec"
	"reflect"
)

// EXPLORATION TESTS
func CouchDbRepoSpec(c Context) {
	c.Specify("db is available", func() {
		newsRepo := CreateCouchDbRepo("localhost", "5678", "hpfeed")
		messages := newsRepo.GetAllMessages()
		c.Expect(len(messages) > 0, IsTrue)
	})
	c.Specify("fetches all values from a hpfeed message", func() {
		newsRepo := CreateCouchDbRepo("localhost", "5678", "hpfeed")
		messages := newsRepo.GetAllMessages()

		c.Expect(messages[0].ID != "", IsTrue)
		c.Expect(messages[0].Topic != "", IsTrue)
		c.Expect(reflect.TypeOf(messages[0].Date).Name() == "time.Time", IsTrue)
		c.Expect(messages[0].Link != "", IsTrue)
	})
}
