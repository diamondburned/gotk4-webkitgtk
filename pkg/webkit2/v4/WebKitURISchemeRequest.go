// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_uri_scheme_request_get_type()), F: marshalURISchemeRequester},
	})
}

type URISchemeRequest struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*URISchemeRequest)(nil)
)

func wrapURISchemeRequest(obj *externglib.Object) *URISchemeRequest {
	return &URISchemeRequest{
		Object: obj,
	}
}

func marshalURISchemeRequester(p uintptr) (interface{}, error) {
	return wrapURISchemeRequest(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Finish a KitURISchemeRequest by setting the contents of the request and its
// mime type.
//
// The function takes the following parameters:
//
//    - stream to read the contents of the request.
//    - streamLength: length of the stream or -1 if not known.
//    - contentType (optional): content type of the stream or NULL if not known.
//
func (request *URISchemeRequest) Finish(stream gio.InputStreamer, streamLength int64, contentType string) {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _arg1 *C.GInputStream           // out
	var _arg2 C.gint64                  // out
	var _arg3 *C.gchar                  // out

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(request.Native()))
	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = C.gint64(streamLength)
	if contentType != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	C.webkit_uri_scheme_request_finish(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(request)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(streamLength)
	runtime.KeepAlive(contentType)
}

// FinishError: finish a KitURISchemeRequest with a #GError.
//
// The function takes the following parameters:
//
//    - err that will be passed to the KitWebView.
//
func (request *URISchemeRequest) FinishError(err error) {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _arg1 *C.GError                 // out

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(request.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.webkit_uri_scheme_request_finish_error(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(err)
}

// Path: get the URI path of request.
//
// The function returns the following values:
//
//    - utf8: URI path of request.
//
func (request *URISchemeRequest) Path() string {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_uri_scheme_request_get_path(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Scheme: get the URI scheme of request.
//
// The function returns the following values:
//
//    - utf8: URI scheme of request.
//
func (request *URISchemeRequest) Scheme() string {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_uri_scheme_request_get_scheme(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URI: get the URI of request.
//
// The function returns the following values:
//
//    - utf8: full URI of request.
//
func (request *URISchemeRequest) URI() string {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_uri_scheme_request_get_uri(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WebView: get the KitWebView that initiated the request.
//
// The function returns the following values:
//
//    - webView that initiated request.
//
func (request *URISchemeRequest) WebView() *WebView {
	var _arg0 *C.WebKitURISchemeRequest // out
	var _cret *C.WebKitWebView          // in

	_arg0 = (*C.WebKitURISchemeRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_uri_scheme_request_get_web_view(_arg0)
	runtime.KeepAlive(request)

	var _webView *WebView // out

	_webView = wrapWebView(externglib.Take(unsafe.Pointer(_cret)))

	return _webView
}
