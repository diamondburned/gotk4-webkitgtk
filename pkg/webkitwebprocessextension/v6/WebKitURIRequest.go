// Code generated by girgen. DO NOT EDIT.

package webkitwebprocessextension

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4-webkitgtk/pkg/soup/v3"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit-web-process-extension.h>
import "C"

// GType values.
var (
	GTypeURIRequest = coreglib.Type(C.webkit_uri_request_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeURIRequest, F: marshalURIRequest},
	})
}

// URIRequestOverrides contains methods that are overridable.
type URIRequestOverrides struct {
}

func defaultURIRequestOverrides(v *URIRequest) URIRequestOverrides {
	return URIRequestOverrides{}
}

// URIRequest represents a URI request.
//
// A KitURIRequest can be created with a URI using the webkit_uri_request_new()
// method, and you can get the URI of an existing request with the
// webkit_uri_request_get_uri() one.
type URIRequest struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*URIRequest)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*URIRequest, *URIRequestClass, URIRequestOverrides](
		GTypeURIRequest,
		initURIRequestClass,
		wrapURIRequest,
		defaultURIRequestOverrides,
	)
}

func initURIRequestClass(gclass unsafe.Pointer, overrides URIRequestOverrides, classInitFunc func(*URIRequestClass)) {
	if classInitFunc != nil {
		class := (*URIRequestClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapURIRequest(obj *coreglib.Object) *URIRequest {
	return &URIRequest{
		Object: obj,
	}
}

func marshalURIRequest(p uintptr) (interface{}, error) {
	return wrapURIRequest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewURIRequest creates a new KitURIRequest for the given URI.
//
// The function takes the following parameters:
//
//   - uri: URI.
//
// The function returns the following values:
//
//   - uriRequest: new KitURIRequest.
//
func NewURIRequest(uri string) *URIRequest {
	var _arg1 *C.gchar            // out
	var _cret *C.WebKitURIRequest // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_uri_request_new(_arg1)
	runtime.KeepAlive(uri)

	var _uriRequest *URIRequest // out

	_uriRequest = wrapURIRequest(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _uriRequest
}

// HTTPHeaders: get the HTTP headers of a KitURIRequest as a MessageHeaders.
//
// The function returns the following values:
//
//   - messageHeaders with the HTTP headers of request or NULL if request is not
//     an HTTP request.
//
func (request *URIRequest) HTTPHeaders() *soup.MessageHeaders {
	var _arg0 *C.WebKitURIRequest   // out
	var _cret *C.SoupMessageHeaders // in

	_arg0 = (*C.WebKitURIRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_request_get_http_headers(_arg0)
	runtime.KeepAlive(request)

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

// HTTPMethod: get the HTTP method of the KitURIRequest.
//
// The function returns the following values:
//
//   - utf8: HTTP method of the KitURIRequest or NULL if request is not an HTTP
//     request.
//
func (request *URIRequest) HTTPMethod() string {
	var _arg0 *C.WebKitURIRequest // out
	var _cret *C.gchar            // in

	_arg0 = (*C.WebKitURIRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_request_get_http_method(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URI obtains the request URI.
//
// The function returns the following values:
//
//   - utf8: request URI, as a string.
//
func (request *URIRequest) URI() string {
	var _arg0 *C.WebKitURIRequest // out
	var _cret *C.gchar            // in

	_arg0 = (*C.WebKitURIRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_uri_request_get_uri(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetURI: set the URI of request.
//
// The function takes the following parameters:
//
//   - uri: URI.
//
func (request *URIRequest) SetURI(uri string) {
	var _arg0 *C.WebKitURIRequest // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.WebKitURIRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_uri_request_set_uri(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(uri)
}

// URIRequestClass: instance of this type is always passed by reference.
type URIRequestClass struct {
	*uriRequestClass
}

// uriRequestClass is the struct that's finalized.
type uriRequestClass struct {
	native *C.WebKitURIRequestClass
}
