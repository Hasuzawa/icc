package tag_test

import (
	"testing"
	"unsafe"

	"github.com/Hasuzawa/icc/tag"
	"github.com/stretchr/testify/assert"
)

func TestTagEntrySize(t *testing.T) {
	assert.Equal(t, uintptr(12), unsafe.Sizeof(tag.TagEntry{}))
}
