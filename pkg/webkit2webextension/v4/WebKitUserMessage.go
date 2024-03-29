// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

// UserMessageErrorQuark gets the quark for the domain of user message errors.
//
// The function returns the following values:
//
//   - quark: user message error domain.
//
func UserMessageErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.webkit_user_message_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

// UserMessageClass: instance of this type is always passed by reference.
type UserMessageClass struct {
	*userMessageClass
}

// userMessageClass is the struct that's finalized.
type userMessageClass struct {
	native *C.WebKitUserMessageClass
}
