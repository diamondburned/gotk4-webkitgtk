// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

//export _gotk4_soup3_Session_ConnectRequestQueued
func _gotk4_soup3_Session_ConnectRequestQueued(arg0 C.gpointer, arg1 *C.SoupMessage, arg2 C.guintptr) {
	var f func(msg *Message)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(msg *Message))
	}

	var _msg *Message // out

	_msg = wrapMessage(coreglib.Take(unsafe.Pointer(arg1)))

	f(_msg)
}

//export _gotk4_soup3_Session_ConnectRequestUnqueued
func _gotk4_soup3_Session_ConnectRequestUnqueued(arg0 C.gpointer, arg1 *C.SoupMessage, arg2 C.guintptr) {
	var f func(msg *Message)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(msg *Message))
	}

	var _msg *Message // out

	_msg = wrapMessage(coreglib.Take(unsafe.Pointer(arg1)))

	f(_msg)
}
