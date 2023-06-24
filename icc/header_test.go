package icc_test

import (
	"testing"
	"unsafe"

	"github.com/Hasuzawa/icc/icc"
	"github.com/stretchr/testify/assert"
)

func TestHeaderSize(t *testing.T) {
	header := icc.Header{}
	assert.Equal(t, uintptr(128), unsafe.Sizeof(header))
}
