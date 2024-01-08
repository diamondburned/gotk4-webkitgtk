// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4-webkitgtk/pkg/soup/v3"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypeURISchemeResponse = coreglib.Type(C.webkit_uri_scheme_response_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeURISchemeResponse, F: marshalURISchemeResponse},
	})
}

// URISchemeResponseOverrides contains methods that are overridable.
type URISchemeResponseOverrides struct {
}

func defaultURISchemeResponseOverrides(v *URISchemeResponse) URISchemeResponseOverrides {
	return URISchemeResponseOverrides{}
}

// URISchemeResponse represents a URI scheme response.
//
// If you register a particular URI scheme in a KitWebContext,
// using webkit_web_context_register_uri_scheme(), you have to provide a
// KitURISchemeRequestCallback. After that, when a URI response is made with
// that particular scheme, your callback will be called. There you will be able
// to provide more response parameters when the methods and properties of a
// KitURISchemeRequest is not enough.
//
// When you finished setting up your KitURISchemeResponse, call
// webkit_uri_request_finish_with_response() with it to return the response.
type URISchemeResponse struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*URISchemeResponse)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*URISchemeResponse, *URISchemeResponseClass, URISchemeResponseOverrides](
		GTypeURISchemeResponse,
		initURISchemeResponseClass,
		wrapURISchemeResponse,
		defaultURISchemeResponseOverrides,
	)
}

func initURISchemeResponseClass(gclass unsafe.Pointer, overrides URISchemeResponseOverrides, classInitFunc func(*URISchemeResponseClass)) {
	if classInitFunc != nil {
		class := (*URISchemeResponseClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapURISchemeResponse(obj *coreglib.Object) *URISchemeResponse {
	return &URISchemeResponse{
		Object: obj,
	}
}

func marshalURISchemeResponse(p uintptr) (interface{}, error) {
	return wrapURISchemeResponse(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewURISchemeResponse: create a new KitURISchemeResponse.
//
// The function takes the following parameters:
//
//   - inputStream to read the contents of the request.
//   - streamLength: length of the stream or -1 if not known.
//
// The function returns the following values:
//
//   - uriSchemeResponse: newly created KitURISchemeResponse.
//
func NewURISchemeResponse(inputStream gio.InputStreamer, streamLength int64) *URISchemeResponse {
	var _arg1 *C.GInputStream            // out
	var _arg2 C.gint64                   // out
	var _cret *C.WebKitURISchemeResponse // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(inputStream).Native()))
	_arg2 = C.gint64(streamLength)

	_cret = C.webkit_uri_scheme_response_new(_arg1, _arg2)
	runtime.KeepAlive(inputStream)
	runtime.KeepAlive(streamLength)

	var _uriSchemeResponse *URISchemeResponse // out

	_uriSchemeResponse = wrapURISchemeResponse(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _uriSchemeResponse
}

// SetContentType sets the content type for the response.
//
// The function takes the following parameters:
//
//   - contentType: content type of the stream.
//
func (response *URISchemeResponse) SetContentType(contentType string) {
	var _arg0 *C.WebKitURISchemeResponse // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.WebKitURISchemeResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_uri_scheme_response_set_content_type(_arg0, _arg1)
	runtime.KeepAlive(response)
	runtime.KeepAlive(contentType)
}

// SetHTTPHeaders: assign the provided MessageHeaders to the response.
//
// headers need to be of the type SOUP_MESSAGE_HEADERS_RESPONSE. Any existing
// headers will be overwritten.
//
// The function takes the following parameters:
//
//   - headers: HTTP headers to be set.
//
func (response *URISchemeResponse) SetHTTPHeaders(headers *soup.MessageHeaders) {
	var _arg0 *C.WebKitURISchemeResponse // out
	var _arg1 *C.SoupMessageHeaders      // out

	_arg0 = (*C.WebKitURISchemeResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))
	_arg1 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(headers)))

	C.webkit_uri_scheme_response_set_http_headers(_arg0, _arg1)
	runtime.KeepAlive(response)
	runtime.KeepAlive(headers)
}

// SetStatus sets the status code and reason phrase for the response.
//
// If status_code is a known value and reason_phrase is NULL, the reason_phrase
// will be set automatically.
//
// The function takes the following parameters:
//
//   - statusCode: HTTP status code to be returned.
//   - reasonPhrase (optional): reason phrase.
//
func (response *URISchemeResponse) SetStatus(statusCode uint, reasonPhrase string) {
	var _arg0 *C.WebKitURISchemeResponse // out
	var _arg1 C.guint                    // out
	var _arg2 *C.gchar                   // out

	_arg0 = (*C.WebKitURISchemeResponse)(unsafe.Pointer(coreglib.InternObject(response).Native()))
	_arg1 = C.guint(statusCode)
	if reasonPhrase != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(reasonPhrase)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.webkit_uri_scheme_response_set_status(_arg0, _arg1, _arg2)
	runtime.KeepAlive(response)
	runtime.KeepAlive(statusCode)
	runtime.KeepAlive(reasonPhrase)
}

// URISchemeResponseClass: instance of this type is always passed by reference.
type URISchemeResponseClass struct {
	*uriSchemeResponseClass
}

// uriSchemeResponseClass is the struct that's finalized.
type uriSchemeResponseClass struct {
	native *C.WebKitURISchemeResponseClass
}