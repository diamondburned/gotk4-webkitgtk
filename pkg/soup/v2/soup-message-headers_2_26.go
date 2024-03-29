// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

// Range represents a byte range as used in the Range header.
//
// If end is non-negative, then start and end represent the bounds of of the
// range, counting from 0. (Eg, the first 500 bytes would be represented as
// start = 0 and end = 499.)
//
// If end is -1 and start is non-negative, then this represents a range starting
// at start and ending with the last byte of the requested resource body. (Eg,
// all but the first 500 bytes would be start = 500, and end = -1.)
//
// If end is -1 and start is negative, then it represents a "suffix range",
// referring to the last -start bytes of the resource body. (Eg, the last 500
// bytes would be start = -500 and end = -1.)
//
// An instance of this type is always passed by reference.
type Range struct {
	*_range
}

// _range is the struct that's finalized.
type _range struct {
	native *C.SoupRange
}

// NewRange creates a new Range instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewRange(start, end int64) Range {
	var f0 C.goffset // out
	f0 = C.goffset(start)
	var f1 C.goffset // out
	f1 = C.goffset(end)

	v := C.SoupRange{
		start: f0,
		end:   f1,
	}

	return *(*Range)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// Start: start of the range.
func (r *Range) Start() int64 {
	valptr := &r.native.start
	var _v int64 // out
	_v = int64(*valptr)
	return _v
}

// End: end of the range.
func (r *Range) End() int64 {
	valptr := &r.native.end
	var _v int64 // out
	_v = int64(*valptr)
	return _v
}

// Start: start of the range.
func (r *Range) SetStart(start int64) {
	valptr := &r.native.start
	*valptr = C.goffset(start)
}

// End: end of the range.
func (r *Range) SetEnd(end int64) {
	valptr := &r.native.end
	*valptr = C.goffset(end)
}
