// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// glib.Type values for WebKitWindowProperties.go.
var GTypeWindowProperties = externglib.Type(C.webkit_window_properties_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeWindowProperties, F: marshalWindowProperties},
	})
}

// WindowPropertiesOverrider contains methods that are overridable.
type WindowPropertiesOverrider interface {
}

type WindowProperties struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*WindowProperties)(nil)
)

func classInitWindowPropertieser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWindowProperties(obj *externglib.Object) *WindowProperties {
	return &WindowProperties{
		Object: obj,
	}
}

func marshalWindowProperties(p uintptr) (interface{}, error) {
	return wrapWindowProperties(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Fullscreen: get whether the window should be shown in fullscreen state or
// not.
//
// The function returns the following values:
//
//    - ok: TRUE if the window should be fullscreen or FALSE otherwise.
//
func (windowProperties *WindowProperties) Fullscreen() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	_cret = C.webkit_window_properties_get_fullscreen(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Geometry: get the geometry the window should have on the screen when shown.
//
// The function returns the following values:
//
//    - geometry: return location for the window geometry.
//
func (windowProperties *WindowProperties) Geometry() *gdk.Rectangle {
	var _arg0 *C.WebKitWindowProperties // out
	var _arg1 C.GdkRectangle            // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	C.webkit_window_properties_get_geometry(_arg0, &_arg1)
	runtime.KeepAlive(windowProperties)

	var _geometry *gdk.Rectangle // out

	_geometry = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _geometry
}

// LocationbarVisible: get whether the window should have the locationbar
// visible or not.
//
// The function returns the following values:
//
//    - ok: TRUE if locationbar should be visible or FALSE otherwise.
//
func (windowProperties *WindowProperties) LocationbarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	_cret = C.webkit_window_properties_get_locationbar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MenubarVisible: get whether the window should have the menubar visible or
// not.
//
// The function returns the following values:
//
//    - ok: TRUE if menubar should be visible or FALSE otherwise.
//
func (windowProperties *WindowProperties) MenubarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	_cret = C.webkit_window_properties_get_menubar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Resizable: get whether the window should be resizable by the user or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the window should be resizable or FALSE otherwise.
//
func (windowProperties *WindowProperties) Resizable() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	_cret = C.webkit_window_properties_get_resizable(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ScrollbarsVisible: get whether the window should have the scrollbars visible
// or not.
//
// The function returns the following values:
//
//    - ok: TRUE if scrollbars should be visible or FALSE otherwise.
//
func (windowProperties *WindowProperties) ScrollbarsVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	_cret = C.webkit_window_properties_get_scrollbars_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StatusbarVisible: get whether the window should have the statusbar visible or
// not.
//
// The function returns the following values:
//
//    - ok: TRUE if statusbar should be visible or FALSE otherwise.
//
func (windowProperties *WindowProperties) StatusbarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	_cret = C.webkit_window_properties_get_statusbar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ToolbarVisible: get whether the window should have the toolbar visible or
// not.
//
// The function returns the following values:
//
//    - ok: TRUE if toolbar should be visible or FALSE otherwise.
//
func (windowProperties *WindowProperties) ToolbarVisible() bool {
	var _arg0 *C.WebKitWindowProperties // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitWindowProperties)(unsafe.Pointer(externglib.InternObject(windowProperties).Native()))

	_cret = C.webkit_window_properties_get_toolbar_visible(_arg0)
	runtime.KeepAlive(windowProperties)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
