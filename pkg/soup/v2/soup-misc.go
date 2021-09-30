// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <libsoup/soup.h>
import "C"

const CHAR_HTTP_CTL = 16
const CHAR_HTTP_SEPARATOR = 8
const CHAR_URI_GEN_DELIMS = 2
const CHAR_URI_PERCENT_ENCODED = 1
const CHAR_URI_SUB_DELIMS = 4

// StrCaseEqual compares v1 and v2 in a case-insensitive manner
func StrCaseEqual(v1 cgo.Handle, v2 cgo.Handle) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v1))
	_arg2 = (C.gconstpointer)(unsafe.Pointer(v2))

	_cret = C.soup_str_case_equal(_arg1, _arg2)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StrCaseHash hashes key in a case-insensitive manner.
func StrCaseHash(key cgo.Handle) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.soup_str_case_hash(_arg1)
	runtime.KeepAlive(key)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}
