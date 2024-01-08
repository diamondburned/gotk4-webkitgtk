// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// GType values.
var (
	GTypePermissionStateQuery = coreglib.Type(C.webkit_permission_state_query_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePermissionStateQuery, F: marshalPermissionStateQuery},
	})
}

// PermissionStateQuery: this query represents a user's choice to allow or
// deny access to "powerful features" of the platform, as specified in the
// [Permissions W3C Specification](https://w3c.github.io/permissions/).
//
// When signalled by the KitWebView through the query-permission-state
// signal, the application has to eventually respond, via
// webkit_permission_state_query_finish(), whether it grants, denies or requests
// a dedicated permission prompt for the given query.
//
// When a KitPermissionStateQuery is not handled by the user, the user-agent is
// instructed to prompt the user for the given permission.
//
// An instance of this type is always passed by reference.
type PermissionStateQuery struct {
	*permissionStateQuery
}

// permissionStateQuery is the struct that's finalized.
type permissionStateQuery struct {
	native *C.WebKitPermissionStateQuery
}

func marshalPermissionStateQuery(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &PermissionStateQuery{&permissionStateQuery{(*C.WebKitPermissionStateQuery)(b)}}, nil
}

// Finish: notify the web-engine of the selected permission state for the
// given query. This function should only be called as a response to the
// WebKitWebView::query-permission-state signal.
//
// The function takes the following parameters:
//
//   - state: KitPermissionState.
//
func (query *PermissionStateQuery) Finish(state PermissionState) {
	var _arg0 *C.WebKitPermissionStateQuery // out
	var _arg1 C.WebKitPermissionState       // out

	_arg0 = (*C.WebKitPermissionStateQuery)(gextras.StructNative(unsafe.Pointer(query)))
	_arg1 = C.WebKitPermissionState(state)

	C.webkit_permission_state_query_finish(_arg0, _arg1)
	runtime.KeepAlive(query)
	runtime.KeepAlive(state)
}

// Name: get the permission name for which access is being queried.
//
// The function returns the following values:
//
//   - utf8: permission name for query.
//
func (query *PermissionStateQuery) Name() string {
	var _arg0 *C.WebKitPermissionStateQuery // out
	var _cret *C.gchar                      // in

	_arg0 = (*C.WebKitPermissionStateQuery)(gextras.StructNative(unsafe.Pointer(query)))

	_cret = C.webkit_permission_state_query_get_name(_arg0)
	runtime.KeepAlive(query)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SecurityOrigin: get the permission origin for which access is being queried.
//
// The function returns the following values:
//
//   - securityOrigin representing the origin from which the query was emitted.
//
func (query *PermissionStateQuery) SecurityOrigin() *SecurityOrigin {
	var _arg0 *C.WebKitPermissionStateQuery // out
	var _cret *C.WebKitSecurityOrigin       // in

	_arg0 = (*C.WebKitPermissionStateQuery)(gextras.StructNative(unsafe.Pointer(query)))

	_cret = C.webkit_permission_state_query_get_security_origin(_arg0)
	runtime.KeepAlive(query)

	var _securityOrigin *SecurityOrigin // out

	_securityOrigin = (*SecurityOrigin)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.webkit_security_origin_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_securityOrigin)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_security_origin_unref((*C.WebKitSecurityOrigin)(intern.C))
		},
	)

	return _securityOrigin
}