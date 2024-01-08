// Code generated by girgen. DO NOT EDIT.

package webkitwebprocessextension

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit/webkit-web-process-extension.h>
import "C"

//export _gotk4_webkitwebprocessextension6_ScriptWorld_ConnectWindowObjectCleared
func _gotk4_webkitwebprocessextension6_ScriptWorld_ConnectWindowObjectCleared(arg0 C.gpointer, arg1 *C.WebKitWebPage, arg2 *C.WebKitFrame, arg3 C.guintptr) {
	var f func(page *WebPage, frame *Frame)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(page *WebPage, frame *Frame))
	}

	var _page *WebPage // out
	var _frame *Frame  // out

	_page = wrapWebPage(coreglib.Take(unsafe.Pointer(arg1)))
	_frame = wrapFrame(coreglib.Take(unsafe.Pointer(arg2)))

	f(_page, _frame)
}
