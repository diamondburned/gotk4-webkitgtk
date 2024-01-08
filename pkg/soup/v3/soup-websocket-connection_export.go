// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

//export _gotk4_soup3_WebsocketConnection_ConnectClosed
func _gotk4_soup3_WebsocketConnection_ConnectClosed(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_soup3_WebsocketConnection_ConnectClosing
func _gotk4_soup3_WebsocketConnection_ConnectClosing(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_soup3_WebsocketConnection_ConnectError
func _gotk4_soup3_WebsocketConnection_ConnectError(arg0 C.gpointer, arg1 *C.GError, arg2 C.guintptr) {
	var f func(err error)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(err error))
	}

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(arg1))

	f(_err)
}

//export _gotk4_soup3_WebsocketConnection_ConnectMessage
func _gotk4_soup3_WebsocketConnection_ConnectMessage(arg0 C.gpointer, arg1 C.gint, arg2 *C.GBytes, arg3 C.guintptr) {
	var f func(typ int, message *glib.Bytes)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(typ int, message *glib.Bytes))
	}

	var _typ int             // out
	var _message *glib.Bytes // out

	_typ = int(arg1)
	_message = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.g_bytes_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	f(_typ, _message)
}

//export _gotk4_soup3_WebsocketConnection_ConnectPong
func _gotk4_soup3_WebsocketConnection_ConnectPong(arg0 C.gpointer, arg1 *C.GBytes, arg2 C.guintptr) {
	var f func(message *glib.Bytes)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(message *glib.Bytes))
	}

	var _message *glib.Bytes // out

	_message = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.g_bytes_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	f(_message)
}