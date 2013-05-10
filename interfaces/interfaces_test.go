package interfaces

import (
	"github.com/ghthor/gospec"
	"testing"
)

func TestAllInterfacesSpecs(t *testing.T) {
	r := gospec.NewRunner()

	r.AddSpec(ForumReaderSpec)
	r.AddSpec(ParserSpec)

	gospec.MainGoTest(r, t)
}
