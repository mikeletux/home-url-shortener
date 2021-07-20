package uidgen

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUIDGen(t *testing.T) {
	// Get a uid
	uid := uidgen{}.New()

	t.Run("Check that uuid is valid", func(t *testing.T) {
		assert.Regexp(t, regexp.MustCompile(`\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`), uid, "The uuid generated is not valid")
	})
}
