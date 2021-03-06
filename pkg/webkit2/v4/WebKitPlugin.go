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
import "C"

// glib.Type values for WebKitPlugin.go.
var GTypePlugin = externglib.Type(C.webkit_plugin_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePlugin, F: marshalPlugin},
	})
}

// PluginOverrider contains methods that are overridable.
type PluginOverrider interface {
}

type Plugin struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Plugin)(nil)
)

func classInitPluginner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPlugin(obj *externglib.Object) *Plugin {
	return &Plugin{
		Object: obj,
	}
}

func marshalPlugin(p uintptr) (interface{}, error) {
	return wrapPlugin(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Description: deprecated: since version 2.32.
//
// The function returns the following values:
//
//    - utf8: description of the plugin.
//
func (plugin *Plugin) Description() string {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.gchar        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(externglib.InternObject(plugin).Native()))

	_cret = C.webkit_plugin_get_description(_arg0)
	runtime.KeepAlive(plugin)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// MIMEInfoList: get information about MIME types handled by the plugin, as a
// list of KitMimeInfo.
//
// Deprecated: since version 2.32.
//
// The function returns the following values:
//
//    - list of KitMimeInfo.
//
func (plugin *Plugin) MIMEInfoList() []*MIMEInfo {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.GList        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(externglib.InternObject(plugin).Native()))

	_cret = C.webkit_plugin_get_mime_info_list(_arg0)
	runtime.KeepAlive(plugin)

	var _list []*MIMEInfo // out

	_list = make([]*MIMEInfo, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.WebKitMimeInfo)(v)
		var dst *MIMEInfo // out
		dst = (*MIMEInfo)(gextras.NewStructNative(unsafe.Pointer(src)))
		C.webkit_mime_info_ref(src)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.webkit_mime_info_unref((*C.WebKitMimeInfo)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// Name: deprecated: since version 2.32.
//
// The function returns the following values:
//
//    - utf8: name of the plugin.
//
func (plugin *Plugin) Name() string {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.gchar        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(externglib.InternObject(plugin).Native()))

	_cret = C.webkit_plugin_get_name(_arg0)
	runtime.KeepAlive(plugin)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Path: deprecated: since version 2.32.
//
// The function returns the following values:
//
//    - utf8: absolute path where the plugin is installed.
//
func (plugin *Plugin) Path() string {
	var _arg0 *C.WebKitPlugin // out
	var _cret *C.gchar        // in

	_arg0 = (*C.WebKitPlugin)(unsafe.Pointer(externglib.InternObject(plugin).Native()))

	_cret = C.webkit_plugin_get_path(_arg0)
	runtime.KeepAlive(plugin)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
