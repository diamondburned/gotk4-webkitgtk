// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
// void _gotk4_webkit6_PermissionRequest_virtual_allow(void* fnptr, WebKitPermissionRequest* arg0) {
//   ((void (*)(WebKitPermissionRequest*))(fnptr))(arg0);
// };
// void _gotk4_webkit6_PermissionRequest_virtual_deny(void* fnptr, WebKitPermissionRequest* arg0) {
//   ((void (*)(WebKitPermissionRequest*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypePermissionRequest = coreglib.Type(C.webkit_permission_request_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePermissionRequest, F: marshalPermissionRequest},
	})
}

// PermissionRequest: permission request.
//
// There are situations where an embedder would need to ask the user for
// permission to do certain types of operations, such as switching to fullscreen
// mode or reporting the user's location through the standard Geolocation API.
// In those cases, WebKit will emit a KitWebView::permission-request signal with
// a KitPermissionRequest object attached to it.
//
// PermissionRequest wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PermissionRequest struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PermissionRequest)(nil)
)

// PermissionRequester describes PermissionRequest's interface methods.
type PermissionRequester interface {
	coreglib.Objector

	// Allow the action which triggered this request.
	Allow()
	// Deny the action which triggered this request.
	Deny()
}

var _ PermissionRequester = (*PermissionRequest)(nil)

func wrapPermissionRequest(obj *coreglib.Object) *PermissionRequest {
	return &PermissionRequest{
		Object: obj,
	}
}

func marshalPermissionRequest(p uintptr) (interface{}, error) {
	return wrapPermissionRequest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Allow the action which triggered this request.
func (request *PermissionRequest) Allow() {
	var _arg0 *C.WebKitPermissionRequest // out

	_arg0 = (*C.WebKitPermissionRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	C.webkit_permission_request_allow(_arg0)
	runtime.KeepAlive(request)
}

// Deny the action which triggered this request.
func (request *PermissionRequest) Deny() {
	var _arg0 *C.WebKitPermissionRequest // out

	_arg0 = (*C.WebKitPermissionRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	C.webkit_permission_request_deny(_arg0)
	runtime.KeepAlive(request)
}

// Allow: allow the action which triggered this request.
func (request *PermissionRequest) allow() {
	gclass := (*C.WebKitPermissionRequestInterface)(coreglib.PeekParentClass(request))
	fnarg := gclass.allow

	var _arg0 *C.WebKitPermissionRequest // out

	_arg0 = (*C.WebKitPermissionRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	C._gotk4_webkit6_PermissionRequest_virtual_allow(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(request)
}

// Deny: deny the action which triggered this request.
func (request *PermissionRequest) deny() {
	gclass := (*C.WebKitPermissionRequestInterface)(coreglib.PeekParentClass(request))
	fnarg := gclass.deny

	var _arg0 *C.WebKitPermissionRequest // out

	_arg0 = (*C.WebKitPermissionRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	C._gotk4_webkit6_PermissionRequest_virtual_deny(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(request)
}

// PermissionRequestInterface: instance of this type is always passed by
// reference.
type PermissionRequestInterface struct {
	*permissionRequestInterface
}

// permissionRequestInterface is the struct that's finalized.
type permissionRequestInterface struct {
	native *C.WebKitPermissionRequestInterface
}