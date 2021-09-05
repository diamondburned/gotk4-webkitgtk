// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_option_menu_item_get_type()), F: marshalOptionMenuItem},
	})
}

// OptionMenuItem: instance of this type is always passed by reference.
type OptionMenuItem struct {
	*optionMenuItem
}

// optionMenuItem is the struct that's finalized.
type optionMenuItem struct {
	native *C.WebKitOptionMenuItem
}

func marshalOptionMenuItem(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &OptionMenuItem{&optionMenuItem{(*C.WebKitOptionMenuItem)(unsafe.Pointer(b))}}, nil
}

// Copy: make a copy of the KitOptionMenuItem.
func (item *OptionMenuItem) Copy() *OptionMenuItem {
	var _arg0 *C.WebKitOptionMenuItem // out
	var _cret *C.WebKitOptionMenuItem // in

	_arg0 = (*C.WebKitOptionMenuItem)(gextras.StructNative(unsafe.Pointer(item)))

	_cret = C.webkit_option_menu_item_copy(_arg0)
	runtime.KeepAlive(item)

	var _optionMenuItem *OptionMenuItem // out

	_optionMenuItem = (*OptionMenuItem)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_optionMenuItem)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_option_menu_item_free((*C.WebKitOptionMenuItem)(intern.C))
		},
	)

	return _optionMenuItem
}

// Label: get the label of a KitOptionMenuItem.
func (item *OptionMenuItem) Label() string {
	var _arg0 *C.WebKitOptionMenuItem // out
	var _cret *C.gchar                // in

	_arg0 = (*C.WebKitOptionMenuItem)(gextras.StructNative(unsafe.Pointer(item)))

	_cret = C.webkit_option_menu_item_get_label(_arg0)
	runtime.KeepAlive(item)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Tooltip: get the tooltip of a KitOptionMenuItem.
func (item *OptionMenuItem) Tooltip() string {
	var _arg0 *C.WebKitOptionMenuItem // out
	var _cret *C.gchar                // in

	_arg0 = (*C.WebKitOptionMenuItem)(gextras.StructNative(unsafe.Pointer(item)))

	_cret = C.webkit_option_menu_item_get_tooltip(_arg0)
	runtime.KeepAlive(item)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IsEnabled: whether a KitOptionMenuItem is enabled.
func (item *OptionMenuItem) IsEnabled() bool {
	var _arg0 *C.WebKitOptionMenuItem // out
	var _cret C.gboolean              // in

	_arg0 = (*C.WebKitOptionMenuItem)(gextras.StructNative(unsafe.Pointer(item)))

	_cret = C.webkit_option_menu_item_is_enabled(_arg0)
	runtime.KeepAlive(item)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsGroupChild: whether a KitOptionMenuItem is a group child.
func (item *OptionMenuItem) IsGroupChild() bool {
	var _arg0 *C.WebKitOptionMenuItem // out
	var _cret C.gboolean              // in

	_arg0 = (*C.WebKitOptionMenuItem)(gextras.StructNative(unsafe.Pointer(item)))

	_cret = C.webkit_option_menu_item_is_group_child(_arg0)
	runtime.KeepAlive(item)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsGroupLabel: whether a KitOptionMenuItem is a group label.
func (item *OptionMenuItem) IsGroupLabel() bool {
	var _arg0 *C.WebKitOptionMenuItem // out
	var _cret C.gboolean              // in

	_arg0 = (*C.WebKitOptionMenuItem)(gextras.StructNative(unsafe.Pointer(item)))

	_cret = C.webkit_option_menu_item_is_group_label(_arg0)
	runtime.KeepAlive(item)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelected: whether a KitOptionMenuItem is the currently selected one.
func (item *OptionMenuItem) IsSelected() bool {
	var _arg0 *C.WebKitOptionMenuItem // out
	var _cret C.gboolean              // in

	_arg0 = (*C.WebKitOptionMenuItem)(gextras.StructNative(unsafe.Pointer(item)))

	_cret = C.webkit_option_menu_item_is_selected(_arg0)
	runtime.KeepAlive(item)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
