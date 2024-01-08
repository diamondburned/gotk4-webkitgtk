// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// extern void _gotk4_webkit24_OptionMenu_ConnectClose(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeOptionMenu = coreglib.Type(C.webkit_option_menu_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOptionMenu, F: marshalOptionMenu},
	})
}

// OptionMenuOverrides contains methods that are overridable.
type OptionMenuOverrides struct {
}

func defaultOptionMenuOverrides(v *OptionMenu) OptionMenuOverrides {
	return OptionMenuOverrides{}
}

// OptionMenu represents the dropdown menu of a select element in a KitWebView.
//
// When a select element in a KitWebView needs to display a dropdown menu, the
// signal KitWebView::show-option-menu is emitted, providing a WebKitOptionMenu
// with the KitOptionMenuItem<!-- -->s that should be displayed.
type OptionMenu struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*OptionMenu)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*OptionMenu, *OptionMenuClass, OptionMenuOverrides](
		GTypeOptionMenu,
		initOptionMenuClass,
		wrapOptionMenu,
		defaultOptionMenuOverrides,
	)
}

func initOptionMenuClass(gclass unsafe.Pointer, overrides OptionMenuOverrides, classInitFunc func(*OptionMenuClass)) {
	if classInitFunc != nil {
		class := (*OptionMenuClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapOptionMenu(obj *coreglib.Object) *OptionMenu {
	return &OptionMenu{
		Object: obj,
	}
}

func marshalOptionMenu(p uintptr) (interface{}, error) {
	return wrapOptionMenu(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClose is emitted when closing a KitOptionMenu is requested. This can
// happen when the user explicitly calls webkit_option_menu_close() or when the
// element is detached from the current page.
func (menu *OptionMenu) ConnectClose(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(menu, "close", false, unsafe.Pointer(C._gotk4_webkit24_OptionMenu_ConnectClose), f)
}

// ActivateItem activates the KitOptionMenuItem at index in menu.
//
// Activating an item changes the value of the element making the
// item the active one. You are expected to close the menu with
// webkit_option_menu_close() after activating an item, calling this function
// again will have no effect.
//
// The function takes the following parameters:
//
//   - index of the item.
//
func (menu *OptionMenu) ActivateItem(index uint) {
	var _arg0 *C.WebKitOptionMenu // out
	var _arg1 C.guint             // out

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.guint(index)

	C.webkit_option_menu_activate_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(index)
}

// Close: request to close a KitOptionMenu.
//
// This emits WebKitOptionMenu::close signal. This function should always
// be called to notify WebKit that the associated menu has been closed.
// If the menu is closed and neither webkit_option_menu_select_item() nor
// webkit_option_menu_activate_item() have been called, the element value
// remains unchanged.
func (menu *OptionMenu) Close() {
	var _arg0 *C.WebKitOptionMenu // out

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	C.webkit_option_menu_close(_arg0)
	runtime.KeepAlive(menu)
}

// Event gets the Event that triggered the dropdown menu. If menu was not
// triggered by a user interaction, like a mouse click, NULL is returned.
//
// The function returns the following values:
//
//   - event: menu event or NULL.
//
func (menu *OptionMenu) Event() *gdk.Event {
	var _arg0 *C.WebKitOptionMenu // out
	var _cret *C.GdkEvent         // in

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_option_menu_get_event(_arg0)
	runtime.KeepAlive(menu)

	var _event *gdk.Event // out

	{
		v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		v = gdk.CopyEventer(v)
		_event = v
	}

	return _event
}

// Item returns the KitOptionMenuItem at index in menu.
//
// The function takes the following parameters:
//
//   - index of the item.
//
// The function returns the following values:
//
//   - optionMenuItem of menu.
//
func (menu *OptionMenu) Item(index uint) *OptionMenuItem {
	var _arg0 *C.WebKitOptionMenu     // out
	var _arg1 C.guint                 // out
	var _cret *C.WebKitOptionMenuItem // in

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.guint(index)

	_cret = C.webkit_option_menu_get_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(index)

	var _optionMenuItem *OptionMenuItem // out

	_optionMenuItem = (*OptionMenuItem)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _optionMenuItem
}

// NItems gets the length of the menu.
//
// The function returns the following values:
//
//   - guint: number of KitOptionMenuItem<!-- -->s in menu.
//
func (menu *OptionMenu) NItems() uint {
	var _arg0 *C.WebKitOptionMenu // out
	var _cret C.guint             // in

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_option_menu_get_n_items(_arg0)
	runtime.KeepAlive(menu)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SelectItem selects the KitOptionMenuItem at index in menu.
//
// Selecting an item changes the text shown by the combo button, but it
// doesn't change the value of the element. You need to explicitly activate
// the item with webkit_option_menu_select_item() or close the menu with
// webkit_option_menu_close() in which case the currently selected item will be
// activated.
//
// The function takes the following parameters:
//
//   - index of the item.
//
func (menu *OptionMenu) SelectItem(index uint) {
	var _arg0 *C.WebKitOptionMenu // out
	var _arg1 C.guint             // out

	_arg0 = (*C.WebKitOptionMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.guint(index)

	C.webkit_option_menu_select_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(index)
}
