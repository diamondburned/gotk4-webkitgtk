// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
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
	GTypeDeviceInfoPermissionRequest = coreglib.Type(C.webkit_device_info_permission_request_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDeviceInfoPermissionRequest, F: marshalDeviceInfoPermissionRequest},
	})
}

// DeviceInfoPermissionRequestOverrides contains methods that are overridable.
type DeviceInfoPermissionRequestOverrides struct {
}

func defaultDeviceInfoPermissionRequestOverrides(v *DeviceInfoPermissionRequest) DeviceInfoPermissionRequestOverrides {
	return DeviceInfoPermissionRequestOverrides{}
}

// DeviceInfoPermissionRequest: permission request for accessing user's
// audio/video devices.
//
// WebKitUserMediaPermissionRequest represents a request for permission to
// whether WebKit should be allowed to access the user's devices information
// when requested through the enumerateDevices API.
//
// When a WebKitDeviceInfoPermissionRequest is not handled by the user, it is
// denied by default.
type DeviceInfoPermissionRequest struct {
	_ [0]func() // equal guard
	*coreglib.Object

	PermissionRequest
}

var (
	_ coreglib.Objector = (*DeviceInfoPermissionRequest)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DeviceInfoPermissionRequest, *DeviceInfoPermissionRequestClass, DeviceInfoPermissionRequestOverrides](
		GTypeDeviceInfoPermissionRequest,
		initDeviceInfoPermissionRequestClass,
		wrapDeviceInfoPermissionRequest,
		defaultDeviceInfoPermissionRequestOverrides,
	)
}

func initDeviceInfoPermissionRequestClass(gclass unsafe.Pointer, overrides DeviceInfoPermissionRequestOverrides, classInitFunc func(*DeviceInfoPermissionRequestClass)) {
	if classInitFunc != nil {
		class := (*DeviceInfoPermissionRequestClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDeviceInfoPermissionRequest(obj *coreglib.Object) *DeviceInfoPermissionRequest {
	return &DeviceInfoPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalDeviceInfoPermissionRequest(p uintptr) (interface{}, error) {
	return wrapDeviceInfoPermissionRequest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
