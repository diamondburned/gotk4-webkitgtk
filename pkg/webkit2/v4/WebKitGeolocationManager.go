// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// extern gboolean _gotk4_webkit24_GeolocationManager_ConnectStart(gpointer, guintptr);
// extern void _gotk4_webkit24_GeolocationManager_ConnectStop(gpointer, guintptr);
import "C"

// glib.Type values for WebKitGeolocationManager.go.
var (
	GTypeGeolocationManager  = externglib.Type(C.webkit_geolocation_manager_get_type())
	GTypeGeolocationPosition = externglib.Type(C.webkit_geolocation_position_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGeolocationManager, F: marshalGeolocationManager},
		{T: GTypeGeolocationPosition, F: marshalGeolocationPosition},
	})
}

// GeolocationManagerOverrider contains methods that are overridable.
type GeolocationManagerOverrider interface {
}

type GeolocationManager struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*GeolocationManager)(nil)
)

func classInitGeolocationManagerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGeolocationManager(obj *externglib.Object) *GeolocationManager {
	return &GeolocationManager{
		Object: obj,
	}
}

func marshalGeolocationManager(p uintptr) (interface{}, error) {
	return wrapGeolocationManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_webkit24_GeolocationManager_ConnectStart
func _gotk4_webkit24_GeolocationManager_ConnectStart(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectStart: signal is emitted to notify that manager needs to start
// receiving position updates. After this signal is emitted the user should
// provide the updates using webkit_geolocation_manager_update_position() every
// time the position changes, or use webkit_geolocation_manager_failed() in case
// it isn't possible to determine the current position.
//
// If the signal is not handled, WebKit will try to determine the position using
// GeoClue if available.
func (manager *GeolocationManager) ConnectStart(f func() (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(manager, "start", false, unsafe.Pointer(C._gotk4_webkit24_GeolocationManager_ConnectStart), f)
}

//export _gotk4_webkit24_GeolocationManager_ConnectStop
func _gotk4_webkit24_GeolocationManager_ConnectStop(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectStop: signal is emitted to notify that manager doesn't need to receive
// position updates anymore.
func (manager *GeolocationManager) ConnectStop(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(manager, "stop", false, unsafe.Pointer(C._gotk4_webkit24_GeolocationManager_ConnectStop), f)
}

// Failed: notify manager that determining the position failed.
//
// The function takes the following parameters:
//
//    - errorMessage: error message.
//
func (manager *GeolocationManager) Failed(errorMessage string) {
	var _arg0 *C.WebKitGeolocationManager // out
	var _arg1 *C.char                     // out

	_arg0 = (*C.WebKitGeolocationManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(errorMessage)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_geolocation_manager_failed(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(errorMessage)
}

// EnableHighAccuracy: get whether high accuracy is enabled.
//
// The function returns the following values:
//
func (manager *GeolocationManager) EnableHighAccuracy() bool {
	var _arg0 *C.WebKitGeolocationManager // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitGeolocationManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_geolocation_manager_get_enable_high_accuracy(_arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UpdatePosition: notify manager that position has been updated to position.
//
// The function takes the following parameters:
//
//    - position: KitGeolocationPosition.
//
func (manager *GeolocationManager) UpdatePosition(position *GeolocationPosition) {
	var _arg0 *C.WebKitGeolocationManager  // out
	var _arg1 *C.WebKitGeolocationPosition // out

	_arg0 = (*C.WebKitGeolocationManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.WebKitGeolocationPosition)(gextras.StructNative(unsafe.Pointer(position)))

	C.webkit_geolocation_manager_update_position(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(position)
}

// GeolocationPosition is an opaque struct used to provide position updates to a
// KitGeolocationManager using webkit_geolocation_manager_update_position().
//
// An instance of this type is always passed by reference.
type GeolocationPosition struct {
	*geolocationPosition
}

// geolocationPosition is the struct that's finalized.
type geolocationPosition struct {
	native *C.WebKitGeolocationPosition
}

func marshalGeolocationPosition(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &GeolocationPosition{&geolocationPosition{(*C.WebKitGeolocationPosition)(b)}}, nil
}

// NewGeolocationPosition constructs a struct GeolocationPosition.
func NewGeolocationPosition(latitude float64, longitude float64, accuracy float64) *GeolocationPosition {
	var _arg1 C.double                     // out
	var _arg2 C.double                     // out
	var _arg3 C.double                     // out
	var _cret *C.WebKitGeolocationPosition // in

	_arg1 = C.double(latitude)
	_arg2 = C.double(longitude)
	_arg3 = C.double(accuracy)

	_cret = C.webkit_geolocation_position_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(latitude)
	runtime.KeepAlive(longitude)
	runtime.KeepAlive(accuracy)

	var _geolocationPosition *GeolocationPosition // out

	_geolocationPosition = (*GeolocationPosition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_geolocationPosition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_geolocation_position_free((*C.WebKitGeolocationPosition)(intern.C))
		},
	)

	return _geolocationPosition
}

// Copy: make a copy of the KitGeolocationPosition.
//
// The function returns the following values:
//
//    - geolocationPosition: copy of position.
//
func (position *GeolocationPosition) Copy() *GeolocationPosition {
	var _arg0 *C.WebKitGeolocationPosition // out
	var _cret *C.WebKitGeolocationPosition // in

	_arg0 = (*C.WebKitGeolocationPosition)(gextras.StructNative(unsafe.Pointer(position)))

	_cret = C.webkit_geolocation_position_copy(_arg0)
	runtime.KeepAlive(position)

	var _geolocationPosition *GeolocationPosition // out

	_geolocationPosition = (*GeolocationPosition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_geolocationPosition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_geolocation_position_free((*C.WebKitGeolocationPosition)(intern.C))
		},
	)

	return _geolocationPosition
}

// SetAltitude: set the position altitude.
//
// The function takes the following parameters:
//
//    - altitude in meters.
//
func (position *GeolocationPosition) SetAltitude(altitude float64) {
	var _arg0 *C.WebKitGeolocationPosition // out
	var _arg1 C.double                     // out

	_arg0 = (*C.WebKitGeolocationPosition)(gextras.StructNative(unsafe.Pointer(position)))
	_arg1 = C.double(altitude)

	C.webkit_geolocation_position_set_altitude(_arg0, _arg1)
	runtime.KeepAlive(position)
	runtime.KeepAlive(altitude)
}

// SetAltitudeAccuracy: set the accuracy of position altitude.
//
// The function takes the following parameters:
//
//    - altitudeAccuracy: accuracy of position altitude in meters.
//
func (position *GeolocationPosition) SetAltitudeAccuracy(altitudeAccuracy float64) {
	var _arg0 *C.WebKitGeolocationPosition // out
	var _arg1 C.double                     // out

	_arg0 = (*C.WebKitGeolocationPosition)(gextras.StructNative(unsafe.Pointer(position)))
	_arg1 = C.double(altitudeAccuracy)

	C.webkit_geolocation_position_set_altitude_accuracy(_arg0, _arg1)
	runtime.KeepAlive(position)
	runtime.KeepAlive(altitudeAccuracy)
}

// SetHeading: set the position heading, as a positive angle between the
// direction of movement and the North direction, in clockwise direction.
//
// The function takes the following parameters:
//
//    - heading in degrees.
//
func (position *GeolocationPosition) SetHeading(heading float64) {
	var _arg0 *C.WebKitGeolocationPosition // out
	var _arg1 C.double                     // out

	_arg0 = (*C.WebKitGeolocationPosition)(gextras.StructNative(unsafe.Pointer(position)))
	_arg1 = C.double(heading)

	C.webkit_geolocation_position_set_heading(_arg0, _arg1)
	runtime.KeepAlive(position)
	runtime.KeepAlive(heading)
}

// SetSpeed: set the position speed.
//
// The function takes the following parameters:
//
//    - speed in meters per second.
//
func (position *GeolocationPosition) SetSpeed(speed float64) {
	var _arg0 *C.WebKitGeolocationPosition // out
	var _arg1 C.double                     // out

	_arg0 = (*C.WebKitGeolocationPosition)(gextras.StructNative(unsafe.Pointer(position)))
	_arg1 = C.double(speed)

	C.webkit_geolocation_position_set_speed(_arg0, _arg1)
	runtime.KeepAlive(position)
	runtime.KeepAlive(speed)
}

// SetTimestamp: set the position timestamp. By default it's the time when the
// position was created.
//
// The function takes the following parameters:
//
//    - timestamp in seconds since the epoch, or 0 to use current time.
//
func (position *GeolocationPosition) SetTimestamp(timestamp uint64) {
	var _arg0 *C.WebKitGeolocationPosition // out
	var _arg1 C.guint64                    // out

	_arg0 = (*C.WebKitGeolocationPosition)(gextras.StructNative(unsafe.Pointer(position)))
	_arg1 = C.guint64(timestamp)

	C.webkit_geolocation_position_set_timestamp(_arg0, _arg1)
	runtime.KeepAlive(position)
	runtime.KeepAlive(timestamp)
}
