package hack

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestConvertEqual(t *testing.T) {
	t.Parallel()

	assert.Equal(t, stringToSlice("hello world😄"), string2Bytes("hello world😄"))
	assert.Equal(t, stringToSlice(""), string2Bytes(""))

	assert.Equal(t, sliceToString(nil), bytes2String(nil))
	assert.Equal(t, sliceToString([]byte{}), bytes2String([]byte{}))
	assert.Equal(t, sliceToString([]byte("123🤣")), bytes2String([]byte("123🤣")))
}

// sliceToString slice to string with out data copy
func sliceToString(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// stringToSlice string to slice with out data copy
func stringToSlice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

// string2Bytes converts string to []byte without memory allocation.
// Go1.20 later, note it may break if string and/or slice header will change
// in the future go versions.
func string2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// bytes2String converts []byte to string without memory allocation.
// go1.20 later.
func bytes2String(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
