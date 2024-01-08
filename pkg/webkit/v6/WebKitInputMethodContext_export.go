// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

//export _gotk4_webkit6_InputMethodContext_ConnectCommitted
func _gotk4_webkit6_InputMethodContext_ConnectCommitted(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(text string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(text string))
	}

	var _text string // out

	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_text)
}

//export _gotk4_webkit6_InputMethodContext_ConnectDeleteSurrounding
func _gotk4_webkit6_InputMethodContext_ConnectDeleteSurrounding(arg0 C.gpointer, arg1 C.gint, arg2 C.guint, arg3 C.guintptr) {
	var f func(offset int, nChars uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offset int, nChars uint))
	}

	var _offset int  // out
	var _nChars uint // out

	_offset = int(arg1)
	_nChars = uint(arg2)

	f(_offset, _nChars)
}

//export _gotk4_webkit6_InputMethodContext_ConnectPreeditChanged
func _gotk4_webkit6_InputMethodContext_ConnectPreeditChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_webkit6_InputMethodContext_ConnectPreeditFinished
func _gotk4_webkit6_InputMethodContext_ConnectPreeditFinished(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_webkit6_InputMethodContext_ConnectPreeditStarted
func _gotk4_webkit6_InputMethodContext_ConnectPreeditStarted(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
