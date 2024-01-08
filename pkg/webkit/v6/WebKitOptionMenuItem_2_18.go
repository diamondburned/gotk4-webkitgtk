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
	GTypeOptionMenuItem = coreglib.Type(C.webkit_option_menu_item_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOptionMenuItem, F: marshalOptionMenuItem},
	})
}

// OptionMenuItem: one item of a KitOptionMenu.
//
// The KitOptionMenu is composed of WebKitOptionMenuItem<!-- -->s.
// A WebKitOptionMenuItem always has a label and can contain a tooltip text. You
// can use the WebKitOptionMenuItem of a KitOptionMenu to build your own menus.
//
// An instance of this type is always passed by reference.
type OptionMenuItem struct {
	*optionMenuItem
}

// optionMenuItem is the struct that's finalized.
type optionMenuItem struct {
	native *C.WebKitOptionMenuItem
}

func marshalOptionMenuItem(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &OptionMenuItem{&optionMenuItem{(*C.WebKitOptionMenuItem)(b)}}, nil
}

// Copy: make a copy of the KitOptionMenuItem.
//
// The function returns the following values:
//
//   - optionMenuItem: copy of passed in KitOptionMenuItem.
//
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
//
// The function returns the following values:
//
//   - utf8: label of item.
//
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
//
// The function returns the following values:
//
//   - utf8: tooltip of item, or NULL.
//
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
//
// The function returns the following values:
//
//   - ok: TRUE if the item is enabled or FALSE otherwise.
//
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
//
// The function returns the following values:
//
//   - ok: TRUE if the item is a group child or FALSE otherwise.
//
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
//
// The function returns the following values:
//
//   - ok: TRUE if the item is a group label or FALSE otherwise.
//
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
//
// The function returns the following values:
//
//   - ok: TRUE if the item is selected or FALSE otherwise.
//
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
