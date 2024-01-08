// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4-webkitgtk/pkg/soup/v3"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

// GType values.
var (
	GTypeURIResponse = coreglib.Type(C.webkit_uri_response_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeURIResponse, F: marshalURIResponse},
	})
}

// URIResponseOverrides contains methods that are overridable.
type URIResponseOverrides struct {
}

func defaultURIResponseOverrides(v *URIResponse) URIResponseOverrides {
	return URIResponseOverrides{}
}

// URIResponse represents an URI response.
//
// A KitURIResponse contains information such as the URI, the status code,
// the content length, the mime type, the HTTP status or the suggested filename.
type URIResponse struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*URIResponse)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*URIResponse, *URIResponseClass, URIResponseOverrides](
		GTypeURIResponse,
		initURIResponseClass,
		wrapURIResponse,
		defaultURIResponseOverrides,
	)
}

func initURIResponseClass(gclass unsafe.Pointer, overrides URIResponseOverrides, classInitFunc func(*URIResponseClass)) {
	if classInitFunc != nil {
		class := (*URIResponseClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapURIResponse(obj *coreglib.Object) *URIResponse {
	return &URIResponse{
		Object: obj,
	}
}

func marshalURIResponse(p uintptr) (interface{}, error) {
	return wrapURIResponse(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ContentLength: get the expected content length of the KitURIResponse.
//
// It can be 0 if the server provided an incorrect or missing Content-Length.
//
// The function returns the following values:
//
//   - guint64: expected content length of response.
//
func (response *URIResponse) ContentLength() uint64 {
	var _arg0 *C.WebKitURIResponse // out
	var _cret C.guint64            // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))

	_cret = C.webkit_uri_response_get_content_length(_arg0)
	runtime.KeepAlive(response)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// HTTPHeaders: get the HTTP headers of a KitURIResponse as a MessageHeaders.
//
// The function returns the following values:
//
//   - messageHeaders with the HTTP headers of response or NULL if response is
//     not an HTTP response.
//
func (response *URIResponse) HTTPHeaders() *soup.MessageHeaders {
	var _arg0 *C.WebKitURIResponse  // out
	var _cret *C.SoupMessageHeaders // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))

	_cret = C.webkit_uri_response_get_http_headers(_arg0)
	runtime.KeepAlive(response)

	var _messageHeaders *soup.MessageHeaders // out

	_messageHeaders = (*soup.MessageHeaders)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.soup_message_headers_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_messageHeaders)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_message_headers_unref((*C.SoupMessageHeaders)(intern.C))
		},
	)

	return _messageHeaders
}

// MIMEType gets the MIME type of the response.
//
// The function returns the following values:
//
//   - utf8: MIME type, as a string.
//
func (response *URIResponse) MIMEType() string {
	var _arg0 *C.WebKitURIResponse // out
	var _cret *C.gchar             // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))

	_cret = C.webkit_uri_response_get_mime_type(_arg0)
	runtime.KeepAlive(response)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// StatusCode: get the status code of the KitURIResponse.
//
// Get the status code of the KitURIResponse as returned by the server. It will
// normally be a KnownStatusCode, for example SOUP_STATUS_OK, though the server
// can respond with any unsigned integer.
//
// The function returns the following values:
//
//   - guint status code of response.
//
func (response *URIResponse) StatusCode() uint {
	var _arg0 *C.WebKitURIResponse // out
	var _cret C.guint              // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))

	_cret = C.webkit_uri_response_get_status_code(_arg0)
	runtime.KeepAlive(response)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SuggestedFilename: get the suggested filename for response.
//
// Get the suggested filename for response, as specified by the
// 'Content-Disposition' HTTP header, or NULL if it's not present.
//
// The function returns the following values:
//
//   - utf8: suggested filename or NULL if the 'Content-Disposition' HTTP header
//     is not present.
//
func (response *URIResponse) SuggestedFilename() string {
	var _arg0 *C.WebKitURIResponse // out
	var _cret *C.gchar             // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))

	_cret = C.webkit_uri_response_get_suggested_filename(_arg0)
	runtime.KeepAlive(response)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URI gets the URI which resulted in the response.
//
// The function returns the following values:
//
//   - utf8: response URI, as a string.
//
func (response *URIResponse) URI() string {
	var _arg0 *C.WebKitURIResponse // out
	var _cret *C.gchar             // in

	_arg0 = (*C.WebKitURIResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))

	_cret = C.webkit_uri_response_get_uri(_arg0)
	runtime.KeepAlive(response)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URIResponseClass: instance of this type is always passed by reference.
type URIResponseClass struct {
	*uriResponseClass
}

// uriResponseClass is the struct that's finalized.
type uriResponseClass struct {
	native *C.WebKitURIResponseClass
}
