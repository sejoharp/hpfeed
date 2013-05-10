package interfaces

import (
	. "github.com/ghthor/gospec"
)

func ForumReaderSpec(c Context) {
	c.Specify("Forum is available", func() {
		available := isWebsiteAvailable("kickern-hamburg.de")
		c.Expect(available, IsTrue)
	})
	c.Specify("localhost on port 80 is not available", func() {
		available := isWebsiteAvailable("localhost")
		c.Expect(available, IsFalse)
	})
}
