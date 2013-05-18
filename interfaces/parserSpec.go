package interfaces

import (
	"code.google.com/p/go.net/html"
	"fmt"
	. "github.com/ghthor/gospec"
	"github.com/puerkitobio/goquery"
	"os"
)

func loadDoc(page string) *goquery.Document {
	if f, e := os.Open(fmt.Sprintf("testdata/%s", page)); e != nil {
		panic(e.Error())
	} else {
		defer f.Close()
		if node, e := html.Parse(f); e != nil {
			panic(e.Error())
		} else {
			return goquery.NewDocumentFromNode(node)
		}
	}
	return nil
}

func ParserSpec(c Context) {
	c.Specify("It parses the date.", func() {
		expectedDateString := "13.01.2013, 18:52"
		resultDate := parseDate(expectedDateString)
		resultDateString := resultDate.Format(RAW_DATE_FORMAT)
		c.Expect(resultDateString, Equals, expectedDateString)
	})
	c.Specify("It parses the data from the latest posts.", func() {
		doc := loadDoc("forum.html")
		threads := ParseThreads(doc)
		c.Expect(len(threads), Equals, 26)

		thread := threads[0]
		c.Expect(thread.Topic, Equals, "Saison 2012")

		c.Expect(thread.Date.Format(RAW_DATE_FORMAT), Equals, "13.01.2013, 18:52")
		c.Expect(thread.Link, Equals, "http://www.kickern-hamburg.de/phpBB2/viewtopic.php?p=55873#55873")
	})
}
