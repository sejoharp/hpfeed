package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bytes"
	"exp/html"
	"github.com/puerkitobio/goquery"
	"strings"
	"time"
)

type Thread struct {
	Topic string
	Date  time.Time
	Link  string
}

const RAW_DATE_FORMAT = "02.01.2006, 15:04"

func GenerateDocument(rawData []byte) *goquery.Document {
	node, err := html.Parse(bytes.NewReader(rawData))
	helper.HandleFatalError("document generation failed:", err)
	return goquery.NewDocumentFromNode(node)
}

func ParseThreads(doc *goquery.Document) []*Thread {
	threads := make([]*Thread, 0)
	doc.Find("#Content table tbody > tr").Each(func(i int, s *goquery.Selection) {
		if s.Children().Length() == 6 {
			thread := &Thread{}
			thread.Topic = strings.TrimSpace(s.Find(":nth-child(2) a.genmed").Text())
			thread.Date = parseDate(strings.TrimSpace(s.Find(":nth-child(6) a").First().Text()))
			link, _ := s.Find(":nth-child(6) a").First().Attr("href")
			thread.Link = "http://www.kickern-hamburg.de/phpBB2/" + strings.TrimSpace(link)
			threads = append(threads, thread)
		}
	})
	return threads
}

func parseDate(rawDate string) time.Time {
	date, _ := time.Parse(RAW_DATE_FORMAT, rawDate)
	return overrideLocation(date)
}

func overrideLocation(t time.Time) time.Time {
	y, m, d := t.Date()
	H, M, S := t.Clock()
	return time.Date(y, m, d, H, M, S, 0, time.Local)
}
