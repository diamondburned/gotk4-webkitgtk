// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"

//export _gotk4_webkit24_FindController_ConnectCountedMatches
func _gotk4_webkit24_FindController_ConnectCountedMatches(arg0 C.gpointer, arg1 C.guint, arg2 C.guintptr) {
	var f func(matchCount uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(matchCount uint))
	}

	var _matchCount uint // out

	_matchCount = uint(arg1)

	f(_matchCount)
}

//export _gotk4_webkit24_FindController_ConnectFailedToFindText
func _gotk4_webkit24_FindController_ConnectFailedToFindText(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_webkit24_FindController_ConnectFoundText
func _gotk4_webkit24_FindController_ConnectFoundText(arg0 C.gpointer, arg1 C.guint, arg2 C.guintptr) {
	var f func(matchCount uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(matchCount uint))
	}

	var _matchCount uint // out

	_matchCount = uint(arg1)

	f(_matchCount)
}