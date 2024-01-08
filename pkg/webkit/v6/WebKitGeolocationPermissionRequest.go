// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypeGeolocationPermissionRequest = coreglib.Type(C.webkit_geolocation_permission_request_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGeolocationPermissionRequest, F: marshalGeolocationPermissionRequest},
	})
}

// GeolocationPermissionRequestOverrides contains methods that are overridable.
type GeolocationPermissionRequestOverrides struct {
}

func defaultGeolocationPermissionRequestOverrides(v *GeolocationPermissionRequest) GeolocationPermissionRequestOverrides {
	return GeolocationPermissionRequestOverrides{}
}

// GeolocationPermissionRequest: permission request for sharing the user's
// location.
//
// WebKitGeolocationPermissionRequest represents a request for permission to
// decide whether WebKit should provide the user's location to a website when
// requested through the Geolocation API.
//
// When a WebKitGeolocationPermissionRequest is not handled by the user,
// it is denied by default.
//
// When embedding web views in your application, you *must* configure an
// application identifier to allow web content to use geolocation services.
// The identifier *must* match the name of the .desktop file which describes the
// application, sans the suffix.
//
// If your application uses #GApplication (or any subclass like
// Application), WebKit will automatically use the identifier returned by
// g_application_get_application_id(). This is the recommended approach for
// enabling geolocation in applications.
//
// If an identifier cannot be obtained through #GApplication, the value returned
// by g_get_prgname() will be used instead as a fallback. For programs which
// cannot use #GApplication, calling g_set_prgname() early during initialization
// is needed when the name of the executable on disk does not match the name of
// a valid .desktop file.
type GeolocationPermissionRequest struct {
	_ [0]func() // equal guard
	*coreglib.Object

	PermissionRequest
}

var (
	_ coreglib.Objector = (*GeolocationPermissionRequest)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*GeolocationPermissionRequest, *GeolocationPermissionRequestClass, GeolocationPermissionRequestOverrides](
		GTypeGeolocationPermissionRequest,
		initGeolocationPermissionRequestClass,
		wrapGeolocationPermissionRequest,
		defaultGeolocationPermissionRequestOverrides,
	)
}

func initGeolocationPermissionRequestClass(gclass unsafe.Pointer, overrides GeolocationPermissionRequestOverrides, classInitFunc func(*GeolocationPermissionRequestClass)) {
	if classInitFunc != nil {
		class := (*GeolocationPermissionRequestClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGeolocationPermissionRequest(obj *coreglib.Object) *GeolocationPermissionRequest {
	return &GeolocationPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalGeolocationPermissionRequest(p uintptr) (interface{}, error) {
	return wrapGeolocationPermissionRequest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GeolocationPermissionRequestClass: instance of this type is always passed by
// reference.
type GeolocationPermissionRequestClass struct {
	*geolocationPermissionRequestClass
}

// geolocationPermissionRequestClass is the struct that's finalized.
type geolocationPermissionRequestClass struct {
	native *C.WebKitGeolocationPermissionRequestClass
}