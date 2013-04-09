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
	utf8String := toUtf8(rawData)
	utf8byteArray := []byte(utf8String)
	node, err := html.Parse(bytes.NewReader(utf8byteArray))
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

func toUtf8(iso8859_1_buf []byte) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return string(buf)
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
