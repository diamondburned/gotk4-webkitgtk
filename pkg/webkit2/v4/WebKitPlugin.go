// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_plugin_get_type()), F: marshalPluginner},
	})
}

type Plugin struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Plugin)(nil)

func wrapPlugin(obj *externglib.Object) *Plugin {
	return &Plugin{
		Object: obj,
	}
}

func marshalPluginner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPlugin(obj), nil
}

// Description: deprecated: since version 2.32.
func (plugin *Plugin) Description() string {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.gchar        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(plugin.Native()))

	_cret = C.webkit_plugin_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// MIMEInfoList: get information about MIME types handled by the plugin, as a
// list of KitMimeInfo.
//
// Deprecated: since version 2.32.
func (plugin *Plugin) MIMEInfoList() *externglib.List {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.GList        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(plugin.Native()))

	_cret = C.webkit_plugin_get_mime_info_list(_arg0)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.WebKitMimeInfo)(_p)
		var dst *MIMEInfo // out
		dst = (*MIMEInfo)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(dst, func(v *MIMEInfo) {
			C.webkit_mime_info_unref((*C.WebKitMimeInfo)(gextras.StructNative(unsafe.Pointer(v))))
		})
		return dst
	})

	return _list
}

// Name: deprecated: since version 2.32.
func (plugin *Plugin) Name() string {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.gchar        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(plugin.Native()))

	_cret = C.webkit_plugin_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Path: deprecated: since version 2.32.
func (plugin *Plugin) Path() string {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.gchar        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(plugin.Native()))

	_cret = C.webkit_plugin_get_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
