// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// glib.Type values for WebKitResponsePolicyDecision.go.
var GTypeResponsePolicyDecision = externglib.Type(C.webkit_response_policy_decision_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeResponsePolicyDecision, F: marshalResponsePolicyDecision},
	})
}

// ResponsePolicyDecisionOverrider contains methods that are overridable.
type ResponsePolicyDecisionOverrider interface {
}

type ResponsePolicyDecision struct {
	_ [0]func() // equal guard
	PolicyDecision
}

var (
	_ PolicyDecisioner = (*ResponsePolicyDecision)(nil)
)

func classInitResponsePolicyDecisioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapResponsePolicyDecision(obj *externglib.Object) *ResponsePolicyDecision {
	return &ResponsePolicyDecision{
		PolicyDecision: PolicyDecision{
			Object: obj,
		},
	}
}

func marshalResponsePolicyDecision(p uintptr) (interface{}, error) {
	return wrapResponsePolicyDecision(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Request: return the KitURIRequest associated with the response decision.
// Modifications to the returned object are <emphasis>not</emphasis> taken into
// account when the request is sent over the network, and is intended only to
// aid in evaluating whether a response decision should be taken or not. To
// modify requests before they are sent over the network the
// KitPage::send-request signal can be used instead.
//
// The function returns the following values:
//
//    - uriRequest: URI request that is associated with this policy decision.
//
func (decision *ResponsePolicyDecision) Request() *URIRequest {
	var _arg0 *C.WebKitResponsePolicyDecision // out
	var _cret *C.WebKitURIRequest             // in

	_arg0 = (*C.WebKitResponsePolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_response_policy_decision_get_request(_arg0)
	runtime.KeepAlive(decision)

	var _uriRequest *URIRequest // out

	_uriRequest = wrapURIRequest(externglib.Take(unsafe.Pointer(_cret)))

	return _uriRequest
}

// Response gets the value of the KitResponsePolicyDecision:response property.
//
// The function returns the following values:
//
//    - uriResponse: URI response that is associated with this policy decision.
//
func (decision *ResponsePolicyDecision) Response() *URIResponse {
	var _arg0 *C.WebKitResponsePolicyDecision // out
	var _cret *C.WebKitURIResponse            // in

	_arg0 = (*C.WebKitResponsePolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_response_policy_decision_get_response(_arg0)
	runtime.KeepAlive(decision)

	var _uriResponse *URIResponse // out

	_uriResponse = wrapURIResponse(externglib.Take(unsafe.Pointer(_cret)))

	return _uriResponse
}

// IsMIMETypeSupported gets whether the MIME type of the response can be
// displayed in the KitWebView that triggered this policy decision request. See
// also webkit_web_view_can_show_mime_type().
//
// The function returns the following values:
//
//    - ok: TRUE if the MIME type of the response is supported or FALSE
//      otherwise.
//
func (decision *ResponsePolicyDecision) IsMIMETypeSupported() bool {
	var _arg0 *C.WebKitResponsePolicyDecision // out
	var _cret C.gboolean                      // in

	_arg0 = (*C.WebKitResponsePolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_response_policy_decision_is_mime_type_supported(_arg0)
	runtime.KeepAlive(decision)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
