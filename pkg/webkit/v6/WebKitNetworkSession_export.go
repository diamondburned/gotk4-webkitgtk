// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

//export _gotk4_webkit6_NetworkSession_ConnectDownloadStarted
func _gotk4_webkit6_NetworkSession_ConnectDownloadStarted(arg0 C.gpointer, arg1 *C.WebKitDownload, arg2 C.guintptr) {
	var f func(download *Download)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(download *Download))
	}

	var _download *Download // out

	_download = wrapDownload(coreglib.Take(unsafe.Pointer(arg1)))

	f(_download)
}