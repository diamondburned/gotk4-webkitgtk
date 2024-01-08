// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"

//export _gotk4_webkit24_WebInspector_ConnectAttach
func _gotk4_webkit24_WebInspector_ConnectAttach(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_webkit24_WebInspector_ConnectBringToFront
func _gotk4_webkit24_WebInspector_ConnectBringToFront(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_webkit24_WebInspector_ConnectClosed
func _gotk4_webkit24_WebInspector_ConnectClosed(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_webkit24_WebInspector_ConnectDetach
func _gotk4_webkit24_WebInspector_ConnectDetach(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_webkit24_WebInspector_ConnectOpenWindow
func _gotk4_webkit24_WebInspector_ConnectOpenWindow(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
