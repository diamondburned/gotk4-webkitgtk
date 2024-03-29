// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

// HeadersParse parses the headers of an HTTP request or response in str and
// stores the results in dest. Beware that dest may be modified even on failure.
//
// This is a low-level method; normally you would use
// soup_headers_parse_request() or soup_headers_parse_response().
//
// The function takes the following parameters:
//
//   - str: header string (including the Request-Line or Status-Line, but not
//     the trailing blank line).
//   - len: length of str.
//   - dest to store the header values in.
//
// The function returns the following values:
//
//   - ok success or failure.
//
func HeadersParse(str string, len int, dest *MessageHeaders) bool {
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg3 *C.SoupMessageHeaders // out
	var _cret C.gboolean            // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(len)
	_arg3 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(dest)))

	_cret = C.soup_headers_parse(_arg1, _arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(len)
	runtime.KeepAlive(dest)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
