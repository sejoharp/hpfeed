package interfaces

import (
	"exp/html"
	"fmt"
	"github.com/puerkitobio/goquery"
	"os"
	"testing"
)

func loadDoc(page string) *goquery.Document {
	if f, e := os.Open(fmt.Sprintf("./testdata/%s", page)); e != nil {
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

func TestParseDate(t *testing.T) {
	expectedDateString := "13.01.2013, 18:52"
	resultDate := parseDate(expectedDateString)
	resultDateString := resultDate.Format(RAW_DATE_FORMAT)
	if expectedDateString != resultDateString {
		t.Errorf("Parsing date failed. expected: %s, result: %s", expectedDateString, resultDateString)
	}
}

func TestParsing(t *testing.T) {
	doc := loadDoc("forum.html")
	threads := parseThreads(doc)
	if amount := len(threads); amount != 26 {
		t.Errorf("Parsing amount failed. Parsed: %d", amount)
	}
	thread := threads[0]
	if topic := thread.topic; topic != "Saison 2012" {
		t.Errorf("Parsing amount failed. Parsed: %s", topic)
	}
	inputDate := "13.01.2013, 18:52"
	if date := thread.date.Format(RAW_DATE_FORMAT); date != inputDate {
		t.Errorf("Parsing date failed. Parsed: %s", date)
	}
	if link := thread.link; link != "http://www.kickern-hamburg.de/phpBB2/viewtopic.php?p=55873#55873" {
		t.Errorf("Parsing link failed. Parsed: %s", link)
	}
}
