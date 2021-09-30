// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4-webkitgtk/pkg/soup/v2"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_uri_response_get_type()), F: marshalURIResponser},
	})
}

type URIResponse struct {
	*externglib.Object
}

func wrapURIResponse(obj *externglib.Object) *URIResponse {
	return &URIResponse{
		Object: obj,
	}
}

func marshalURIResponser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapURIResponse(obj), nil
}

// ContentLength: get the expected content length of the KitURIResponse. It can
// be 0 if the server provided an incorrect or missing Content-Length.
func (response *URIResponse) ContentLength() uint64 {
	var _arg0 *C.WebKitURIResponse // out
	var _cret C.guint64            // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(response.Native()))

	_cret = C.webkit_uri_response_get_content_length(_arg0)
	runtime.KeepAlive(response)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// HttpHeaders: get the HTTP headers of a KitURIResponse as a MessageHeaders.
func (response *URIResponse) HttpHeaders() *soup.MessageHeaders {
	var _arg0 *C.WebKitURIResponse  // out
	var _cret *C.SoupMessageHeaders // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(response.Native()))

	_cret = C.webkit_uri_response_get_http_headers(_arg0)
	runtime.KeepAlive(response)

	var _messageHeaders *soup.MessageHeaders // out

	_messageHeaders = (*soup.MessageHeaders)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _messageHeaders
}

func (response *URIResponse) MIMEType() string {
	var _arg0 *C.WebKitURIResponse // out
	var _cret *C.gchar             // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(response.Native()))

	_cret = C.webkit_uri_response_get_mime_type(_arg0)
	runtime.KeepAlive(response)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// StatusCode: get the status code of the KitURIResponse as returned by the
// server. It will normally be a KnownStatusCode, for example SOUP_STATUS_OK,
// though the server can respond with any unsigned integer.
func (response *URIResponse) StatusCode() uint {
	var _arg0 *C.WebKitURIResponse // out
	var _cret C.guint              // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(response.Native()))

	_cret = C.webkit_uri_response_get_status_code(_arg0)
	runtime.KeepAlive(response)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SuggestedFilename: get the suggested filename for response, as specified by
// the 'Content-Disposition' HTTP header, or NULL if it's not present.
func (response *URIResponse) SuggestedFilename() string {
	var _arg0 *C.WebKitURIResponse // out
	var _cret *C.gchar             // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(response.Native()))

	_cret = C.webkit_uri_response_get_suggested_filename(_arg0)
	runtime.KeepAlive(response)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

func (response *URIResponse) URI() string {
	var _arg0 *C.WebKitURIResponse // out
	var _cret *C.gchar             // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(response.Native()))

	_cret = C.webkit_uri_response_get_uri(_arg0)
	runtime.KeepAlive(response)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
