// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"

//export _gotk4_webkit24_AuthenticationRequest_ConnectAuthenticated
func _gotk4_webkit24_AuthenticationRequest_ConnectAuthenticated(arg0 C.gpointer, arg1 *C.WebKitCredential, arg2 C.guintptr) {
	var f func(credential *Credential)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(credential *Credential))
	}

	var _credential *Credential // out

	_credential = (*Credential)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_credential)
}

//export _gotk4_webkit24_AuthenticationRequest_ConnectCancelled
func _gotk4_webkit24_AuthenticationRequest_ConnectCancelled(arg0 C.gpointer, arg1 C.guintptr) {
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