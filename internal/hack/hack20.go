//go:build go1.20
// +build go1.20

package hack

import (
	"unsafe"
)

// hack 减少内存分配

// SliceToString slice to string with out data copy.
// Warning: read-only
// https://groups.google.com/g/Golang-Nuts/c/ENgbUzYvCuU/m/90yGx7GUAgAJ%20
func SliceToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// StringToSlice string to slice with out data copy
func StringToSlice(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
