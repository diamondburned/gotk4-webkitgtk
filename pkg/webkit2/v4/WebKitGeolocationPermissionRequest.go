// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_geolocation_permission_request_get_type()), F: marshalGeolocationPermissionRequester},
	})
}

type GeolocationPermissionRequest struct {
	*externglib.Object

	PermissionRequest
}

func wrapGeolocationPermissionRequest(obj *externglib.Object) *GeolocationPermissionRequest {
	return &GeolocationPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalGeolocationPermissionRequester(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGeolocationPermissionRequest(obj), nil
}

func (*GeolocationPermissionRequest) privateGeolocationPermissionRequest() {}
