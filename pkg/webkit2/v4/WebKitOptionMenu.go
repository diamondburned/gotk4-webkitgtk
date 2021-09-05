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
		{T: externglib.Type(C.webkit_option_menu_get_type()), F: marshalOptionMenuer},
	})
}

type OptionMenu struct {
	*externglib.Object
}

func wrapOptionMenu(obj *externglib.Object) *OptionMenu {
	return &OptionMenu{
		Object: obj,
	}
}

func marshalOptionMenuer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapOptionMenu(obj), nil
}

// ActivateItem activates the KitOptionMenuItem at index in menu. Activating an
// item changes the value of the element making the item the active one. You are
// expected to close the menu with webkit_option_menu_close() after activating
// an item, calling this function again will have no effect.
func (menu *OptionMenu) ActivateItem(index uint32) {
	var _arg0 *C.WebKitOptionMenu // out
	var _arg1 C.guint             // out

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = C.guint(index)

	C.webkit_option_menu_activate_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(index)
}

// Close: request to close a KitOptionMenu. This emits WebKitOptionMenu::close
// signal. This function should always be called to notify WebKit that the
// associated menu has been closed. If the menu is closed and neither
// webkit_option_menu_select_item() nor webkit_option_menu_activate_item() have
// been called, the element value remains unchanged.
func (menu *OptionMenu) Close() {
	var _arg0 *C.WebKitOptionMenu // out

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(menu.Native()))

	C.webkit_option_menu_close(_arg0)
	runtime.KeepAlive(menu)
}

// Item returns the KitOptionMenuItem at index in menu.
func (menu *OptionMenu) Item(index uint32) *OptionMenuItem {
	var _arg0 *C.WebKitOptionMenu     // out
	var _arg1 C.guint                 // out
	var _cret *C.WebKitOptionMenuItem // in

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = C.guint(index)

	_cret = C.webkit_option_menu_get_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(index)

	var _optionMenuItem *OptionMenuItem // out

	_optionMenuItem = (*OptionMenuItem)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _optionMenuItem
}

// NItems gets the length of the menu.
func (menu *OptionMenu) NItems() uint32 {
	var _arg0 *C.WebKitOptionMenu // out
	var _cret C.guint             // in

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.webkit_option_menu_get_n_items(_arg0)
	runtime.KeepAlive(menu)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// SelectItem selects the KitOptionMenuItem at index in menu. Selecting an item
// changes the text shown by the combo button, but it doesn't change the value
// of the element. You need to explicitly activate the item with
// webkit_option_menu_select_item() or close the menu with
// webkit_option_menu_close() in which case the currently selected item will be
// activated.
func (menu *OptionMenu) SelectItem(index uint32) {
	var _arg0 *C.WebKitOptionMenu // out
	var _arg1 C.guint             // out

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = C.guint(index)

	C.webkit_option_menu_select_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(index)
}
