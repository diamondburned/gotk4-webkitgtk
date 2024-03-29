// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

//export _gotk4_soup2_CookieJar_ConnectChanged
func _gotk4_soup2_CookieJar_ConnectChanged(arg0 C.gpointer, arg1 *C.SoupCookie, arg2 *C.SoupCookie, arg3 C.guintptr) {
	var f func(oldCookie, newCookie *Cookie)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(oldCookie, newCookie *Cookie))
	}

	var _oldCookie *Cookie // out
	var _newCookie *Cookie // out

	_oldCookie = (*Cookie)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_newCookie = (*Cookie)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	f(_oldCookie, _newCookie)
}
