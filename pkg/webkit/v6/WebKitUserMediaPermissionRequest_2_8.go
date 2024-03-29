// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
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
	GTypeUserMediaPermissionRequest = coreglib.Type(C.webkit_user_media_permission_request_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeUserMediaPermissionRequest, F: marshalUserMediaPermissionRequest},
	})
}

// UserMediaPermissionIsForAudioDevice: check whether the permission request is
// for an audio device.
//
// The function takes the following parameters:
//
//   - request: KitUserMediaPermissionRequest.
//
// The function returns the following values:
//
//   - ok: TRUE if access to an audio device was requested.
//
func UserMediaPermissionIsForAudioDevice(request *UserMediaPermissionRequest) bool {
	var _arg1 *C.WebKitUserMediaPermissionRequest // out
	var _cret C.gboolean                          // in

	_arg1 = (*C.WebKitUserMediaPermissionRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_user_media_permission_is_for_audio_device(_arg1)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UserMediaPermissionIsForVideoDevice: check whether the permission request is
// for a video device.
//
// The function takes the following parameters:
//
//   - request: KitUserMediaPermissionRequest.
//
// The function returns the following values:
//
//   - ok: TRUE if access to a video device was requested.
//
func UserMediaPermissionIsForVideoDevice(request *UserMediaPermissionRequest) bool {
	var _arg1 *C.WebKitUserMediaPermissionRequest // out
	var _cret C.gboolean                          // in

	_arg1 = (*C.WebKitUserMediaPermissionRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_user_media_permission_is_for_video_device(_arg1)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UserMediaPermissionRequestOverrides contains methods that are overridable.
type UserMediaPermissionRequestOverrides struct {
}

func defaultUserMediaPermissionRequestOverrides(v *UserMediaPermissionRequest) UserMediaPermissionRequestOverrides {
	return UserMediaPermissionRequestOverrides{}
}

// UserMediaPermissionRequest: permission request for accessing user's
// audio/video devices.
//
// WebKitUserMediaPermissionRequest represents a request for permission to
// decide whether WebKit should be allowed to access the user's audio and video
// source devices when requested through the getUserMedia API.
//
// When a WebKitUserMediaPermissionRequest is not handled by the user, it is
// denied by default.
type UserMediaPermissionRequest struct {
	_ [0]func() // equal guard
	*coreglib.Object

	PermissionRequest
}

var (
	_ coreglib.Objector = (*UserMediaPermissionRequest)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*UserMediaPermissionRequest, *UserMediaPermissionRequestClass, UserMediaPermissionRequestOverrides](
		GTypeUserMediaPermissionRequest,
		initUserMediaPermissionRequestClass,
		wrapUserMediaPermissionRequest,
		defaultUserMediaPermissionRequestOverrides,
	)
}

func initUserMediaPermissionRequestClass(gclass unsafe.Pointer, overrides UserMediaPermissionRequestOverrides, classInitFunc func(*UserMediaPermissionRequestClass)) {
	if classInitFunc != nil {
		class := (*UserMediaPermissionRequestClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapUserMediaPermissionRequest(obj *coreglib.Object) *UserMediaPermissionRequest {
	return &UserMediaPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalUserMediaPermissionRequest(p uintptr) (interface{}, error) {
	return wrapUserMediaPermissionRequest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
