// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_web_view_session_state_get_type()), F: marshalWebViewSessionState},
	})
}

// WebViewSessionState: instance of this type is always passed by reference.
type WebViewSessionState struct {
	*webViewSessionState
}

// webViewSessionState is the struct that's finalized.
type webViewSessionState struct {
	native *C.WebKitWebViewSessionState
}

func marshalWebViewSessionState(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &WebViewSessionState{&webViewSessionState{(*C.WebKitWebViewSessionState)(b)}}, nil
}

// NewWebViewSessionState constructs a struct WebViewSessionState.
func NewWebViewSessionState(data *glib.Bytes) *WebViewSessionState {
	var _arg1 *C.GBytes                    // out
	var _cret *C.WebKitWebViewSessionState // in

	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(data)))

	_cret = C.webkit_web_view_session_state_new(_arg1)
	runtime.KeepAlive(data)

	var _webViewSessionState *WebViewSessionState // out

	_webViewSessionState = (*WebViewSessionState)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_webViewSessionState)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_web_view_session_state_unref((*C.WebKitWebViewSessionState)(intern.C))
		},
	)

	return _webViewSessionState
}

// Serialize serializes a KitWebViewSessionState.
//
// The function returns the following values:
//
//    - bytes containing the state serialized.
//
func (state *WebViewSessionState) Serialize() *glib.Bytes {
	var _arg0 *C.WebKitWebViewSessionState // out
	var _cret *C.GBytes                    // in

	_arg0 = (*C.WebKitWebViewSessionState)(gextras.StructNative(unsafe.Pointer(state)))

	_cret = C.webkit_web_view_session_state_serialize(_arg0)
	runtime.KeepAlive(state)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}
