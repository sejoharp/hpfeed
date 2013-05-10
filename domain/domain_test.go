package domain


import (
	"github.com/ghthor/gospec"
	"testing"
)

func TestAllDomainSpecs(t *testing.T) {
	r := gospec.NewRunner()
	
	// r.AddSpec(CouchDbRepoSpec)
	// r.AddSpec(MongoDbRepoSpec)

	gospec.MainGoTest(r, t)
}
