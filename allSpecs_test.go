package main

import (
	"bitbucket.org/joscha/hpfeed/interfaces"
	"github.com/ghthor/gospec"
	"testing"
	//"bitbucket.org/joscha/hpfeed/domain"
)

func TestSpecs(t *testing.T) {
	r := gospec.NewRunner()

	r.AddSpec(interfaces.ForumReaderSpec)
	r.AddSpec(interfaces.ParserSpec)
	r.AddSpec(interfaces.WebserviceSpec)
	r.AddSpec(interfaces.FeedUpdaterBatchSpec)
	//r.AddSpec(domain.CouchDbRepoSpec)
	//r.AddSpec(domain.MongoDbRepoSpec)

	gospec.MainGoTest(r, t)
}
