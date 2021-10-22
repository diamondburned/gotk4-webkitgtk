// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
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
		{T: externglib.Type(C.webkit_user_media_permission_request_get_type()), F: marshalUserMediaPermissionRequester},
	})
}

//
// The function takes the following parameters:
//
//    - request: KitUserMediaPermissionRequest.
//
func UserMediaPermissionIsForAudioDevice(request *UserMediaPermissionRequest) bool {
	var _arg1 *C.WebKitUserMediaPermissionRequest // out
	var _cret C.gboolean                          // in

	_arg1 = (*C.WebKitUserMediaPermissionRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_user_media_permission_is_for_audio_device(_arg1)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

//
// The function takes the following parameters:
//
//    - request: KitUserMediaPermissionRequest.
//
func UserMediaPermissionIsForVideoDevice(request *UserMediaPermissionRequest) bool {
	var _arg1 *C.WebKitUserMediaPermissionRequest // out
	var _cret C.gboolean                          // in

	_arg1 = (*C.WebKitUserMediaPermissionRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_user_media_permission_is_for_video_device(_arg1)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

type UserMediaPermissionRequest struct {
	*externglib.Object

	PermissionRequest
}

func wrapUserMediaPermissionRequest(obj *externglib.Object) *UserMediaPermissionRequest {
	return &UserMediaPermissionRequest{
		Object: obj,
		PermissionRequest: PermissionRequest{
			Object: obj,
		},
	}
}

func marshalUserMediaPermissionRequester(p uintptr) (interface{}, error) {
	return wrapUserMediaPermissionRequest(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
