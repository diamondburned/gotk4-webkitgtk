// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"

//export _gotk4_webkit24_Download_ConnectCreatedDestination
func _gotk4_webkit24_Download_ConnectCreatedDestination(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(destination string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(destination string))
	}

	var _destination string // out

	_destination = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_destination)
}

//export _gotk4_webkit24_Download_ConnectDecideDestination
func _gotk4_webkit24_Download_ConnectDecideDestination(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) (cret C.gboolean) {
	var f func(suggestedFilename string) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(suggestedFilename string) (ok bool))
	}

	var _suggestedFilename string // out

	_suggestedFilename = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := f(_suggestedFilename)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_webkit24_Download_ConnectFailed
func _gotk4_webkit24_Download_ConnectFailed(arg0 C.gpointer, arg1 *C.GError, arg2 C.guintptr) {
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

//export _gotk4_webkit24_Download_ConnectFinished
func _gotk4_webkit24_Download_ConnectFinished(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_webkit24_Download_ConnectReceivedData
func _gotk4_webkit24_Download_ConnectReceivedData(arg0 C.gpointer, arg1 C.guint64, arg2 C.guintptr) {
	var f func(dataLength uint64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(dataLength uint64))
	}

	var _dataLength uint64 // out

	_dataLength = uint64(arg1)

	f(_dataLength)
}
