// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// GType values.
var (
	GTypeContextMenu = coreglib.Type(C.webkit_context_menu_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContextMenu, F: marshalContextMenu},
	})
}

// ContextMenuOverrides contains methods that are overridable.
type ContextMenuOverrides struct {
}

func defaultContextMenuOverrides(v *ContextMenu) ContextMenuOverrides {
	return ContextMenuOverrides{}
}

// ContextMenu represents the context menu in a KitWebView.
//
// KitContextMenu represents a context menu containing KitContextMenuItem<!--
// -->s in a KitWebView.
//
// When a KitWebView is about to display the context menu, it emits the
// KitWebView::context-menu signal, which has the KitContextMenu as an
// argument. You can modify it, adding new submenus that you can create
// with webkit_context_menu_new(), adding new KitContextMenuItem<!-- -->s
// with webkit_context_menu_prepend(), webkit_context_menu_append() or
// webkit_context_menu_insert(), maybe after having removed the existing ones
// with webkit_context_menu_remove_all().
type ContextMenu struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ContextMenu)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ContextMenu, *ContextMenuClass, ContextMenuOverrides](
		GTypeContextMenu,
		initContextMenuClass,
		wrapContextMenu,
		defaultContextMenuOverrides,
	)
}

func initContextMenuClass(gclass unsafe.Pointer, overrides ContextMenuOverrides, classInitFunc func(*ContextMenuClass)) {
	if classInitFunc != nil {
		class := (*ContextMenuClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapContextMenu(obj *coreglib.Object) *ContextMenu {
	return &ContextMenu{
		Object: obj,
	}
}

func marshalContextMenu(p uintptr) (interface{}, error) {
	return wrapContextMenu(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewContextMenu creates a new KitContextMenu object.
//
// Creates a new KitContextMenu object to be used as a submenu of an existing
// KitContextMenu. The context menu of a KitWebView is created by the view and
// passed as an argument of KitWebView::context-menu signal. To add items to
// the menu use webkit_context_menu_prepend(), webkit_context_menu_append() or
// webkit_context_menu_insert(). See also webkit_context_menu_new_with_items()
// to create a KitContextMenu with a list of initial items.
//
// The function returns the following values:
//
//   - contextMenu: newly created KitContextMenu object.
//
func NewContextMenu() *ContextMenu {
	var _cret *C.WebKitContextMenu // in

	_cret = C.webkit_context_menu_new()

	var _contextMenu *ContextMenu // out

	_contextMenu = wrapContextMenu(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _contextMenu
}

// NewContextMenuWithItems creates a new KitContextMenu object with the given
// items.
//
// Creates a new KitContextMenu object to be used as a submenu of
// an existing KitContextMenu with the given initial items. See also
// webkit_context_menu_new().
//
// The function takes the following parameters:
//
//   - items of KitContextMenuItem.
//
// The function returns the following values:
//
//   - contextMenu: newly created KitContextMenu object.
//
func NewContextMenuWithItems(items []*ContextMenuItem) *ContextMenu {
	var _arg1 *C.GList             // out
	var _cret *C.WebKitContextMenu // in

	for i := len(items) - 1; i >= 0; i-- {
		src := items[i]
		var dst *C.WebKitContextMenuItem // out
		dst = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(src).Native()))
		_arg1 = C.g_list_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_list_free(_arg1)

	_cret = C.webkit_context_menu_new_with_items(_arg1)
	runtime.KeepAlive(items)

	var _contextMenu *ContextMenu // out

	_contextMenu = wrapContextMenu(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _contextMenu
}

// Append adds item at the end of the menu.
//
// The function takes the following parameters:
//
//   - item to add.
//
func (menu *ContextMenu) Append(item *ContextMenuItem) {
	var _arg0 *C.WebKitContextMenu     // out
	var _arg1 *C.WebKitContextMenuItem // out

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	C.webkit_context_menu_append(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(item)
}

// First gets the first item in the menu.
//
// The function returns the following values:
//
//   - contextMenuItem: first KitContextMenuItem of menu, or NULL if the
//     KitContextMenu is empty.
//
func (menu *ContextMenu) First() *ContextMenuItem {
	var _arg0 *C.WebKitContextMenu     // out
	var _cret *C.WebKitContextMenuItem // in

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_context_menu_first(_arg0)
	runtime.KeepAlive(menu)

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// Event gets the Event that triggered the context menu. This function
// only returns a valid Event when called for a KitContextMenu passed to
// KitWebView::context-menu signal; in all other cases, NULL is returned.
//
// The returned Event is expected to be one of the following types:
// <itemizedlist> <listitem><para> a EventButton of type GDK_BUTTON_PRESS
// when the context menu was triggered with mouse. </para></listitem>
// <listitem><para> a EventKey of type GDK_KEY_PRESS if the keyboard was used
// to show the menu. </para></listitem> <listitem><para> a generic Event of type
// GDK_NOTHING when the Widget::popup-menu signal was used to show the context
// menu. </para></listitem> </itemizedlist>.
//
// The function returns the following values:
//
//   - event: menu event or NULL.
//
func (menu *ContextMenu) Event() *gdk.Event {
	var _arg0 *C.WebKitContextMenu // out
	var _cret *C.GdkEvent          // in

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_context_menu_get_event(_arg0)
	runtime.KeepAlive(menu)

	var _event *gdk.Event // out

	{
		v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		v = gdk.CopyEventer(v)
		_event = v
	}

	return _event
}

// ItemAtPosition gets the item at the given position in the menu.
//
// The function takes the following parameters:
//
//   - position of the item, counting from 0.
//
// The function returns the following values:
//
//   - contextMenuItem at position position in menu, or NULL if the position is
//     off the end of the menu.
//
func (menu *ContextMenu) ItemAtPosition(position uint) *ContextMenuItem {
	var _arg0 *C.WebKitContextMenu     // out
	var _arg1 C.guint                  // out
	var _cret *C.WebKitContextMenuItem // in

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.guint(position)

	_cret = C.webkit_context_menu_get_item_at_position(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(position)

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// Items returns the item list of menu.
//
// The function returns the following values:
//
//   - list of KitContextMenuItem<!-- -->s.
//
func (menu *ContextMenu) Items() []*ContextMenuItem {
	var _arg0 *C.WebKitContextMenu // out
	var _cret *C.GList             // in

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_context_menu_get_items(_arg0)
	runtime.KeepAlive(menu)

	var _list []*ContextMenuItem // out

	_list = make([]*ContextMenuItem, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.WebKitContextMenuItem)(v)
		var dst *ContextMenuItem // out
		dst = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// NItems gets the length of the menu.
//
// The function returns the following values:
//
//   - guint: number of KitContextMenuItem<!-- -->s in menu.
//
func (menu *ContextMenu) NItems() uint {
	var _arg0 *C.WebKitContextMenu // out
	var _cret C.guint              // in

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_context_menu_get_n_items(_arg0)
	runtime.KeepAlive(menu)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// UserData gets the user data of menu.
//
// This function can be used from the UI Process to get user data previously set
// from the Web Process with webkit_context_menu_set_user_data().
//
// The function returns the following values:
//
//   - variant: user data of menu, or NULL if menu doesn't have user data.
//
func (menu *ContextMenu) UserData() *glib.Variant {
	var _arg0 *C.WebKitContextMenu // out
	var _cret *C.GVariant          // in

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_context_menu_get_user_data(_arg0)
	runtime.KeepAlive(menu)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// Insert inserts item into the menu at the given position.
//
// If position is negative, or is larger than the number of items in the
// KitContextMenu, the item is added on to the end of the menu. The first
// position is 0.
//
// The function takes the following parameters:
//
//   - item to add.
//   - position to insert the item.
//
func (menu *ContextMenu) Insert(item *ContextMenuItem, position int) {
	var _arg0 *C.WebKitContextMenu     // out
	var _arg1 *C.WebKitContextMenuItem // out
	var _arg2 C.gint                   // out

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))
	_arg2 = C.gint(position)

	C.webkit_context_menu_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(item)
	runtime.KeepAlive(position)
}

// Last gets the last item in the menu.
//
// The function returns the following values:
//
//   - contextMenuItem: last KitContextMenuItem of menu, or NULL if the
//     KitContextMenu is empty.
//
func (menu *ContextMenu) Last() *ContextMenuItem {
	var _arg0 *C.WebKitContextMenu     // out
	var _cret *C.WebKitContextMenuItem // in

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.webkit_context_menu_last(_arg0)
	runtime.KeepAlive(menu)

	var _contextMenuItem *ContextMenuItem // out

	_contextMenuItem = wrapContextMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _contextMenuItem
}

// MoveItem moves item to the given position in the menu.
//
// If position is negative, or is larger than the number of items in the
// KitContextMenu, the item is added on to the end of the menu. The first
// position is 0.
//
// The function takes the following parameters:
//
//   - item to add.
//   - position: new position to move the item.
//
func (menu *ContextMenu) MoveItem(item *ContextMenuItem, position int) {
	var _arg0 *C.WebKitContextMenu     // out
	var _arg1 *C.WebKitContextMenuItem // out
	var _arg2 C.gint                   // out

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))
	_arg2 = C.gint(position)

	C.webkit_context_menu_move_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(item)
	runtime.KeepAlive(position)
}

// Prepend adds item at the beginning of the menu.
//
// The function takes the following parameters:
//
//   - item to add.
//
func (menu *ContextMenu) Prepend(item *ContextMenuItem) {
	var _arg0 *C.WebKitContextMenu     // out
	var _arg1 *C.WebKitContextMenuItem // out

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	C.webkit_context_menu_prepend(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(item)
}

// Remove removes item from the menu.
//
// See also webkit_context_menu_remove_all() to remove all items.
//
// The function takes the following parameters:
//
//   - item to remove.
//
func (menu *ContextMenu) Remove(item *ContextMenuItem) {
	var _arg0 *C.WebKitContextMenu     // out
	var _arg1 *C.WebKitContextMenuItem // out

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.WebKitContextMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	C.webkit_context_menu_remove(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(item)
}

// RemoveAll removes all items of the menu.
func (menu *ContextMenu) RemoveAll() {
	var _arg0 *C.WebKitContextMenu // out

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	C.webkit_context_menu_remove_all(_arg0)
	runtime.KeepAlive(menu)
}

// SetUserData sets user data to menu.
//
// This function can be used from a Web Process extension to
// set user data that can be retrieved from the UI Process using
// webkit_context_menu_get_user_data(). If the user_data #GVariant is floating,
// it is consumed.
//
// The function takes the following parameters:
//
//   - userData: #GVariant.
//
func (menu *ContextMenu) SetUserData(userData *glib.Variant) {
	var _arg0 *C.WebKitContextMenu // out
	var _arg1 *C.GVariant          // out

	_arg0 = (*C.WebKitContextMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(userData)))

	C.webkit_context_menu_set_user_data(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(userData)
}

// ContextMenuClass: instance of this type is always passed by reference.
type ContextMenuClass struct {
	*contextMenuClass
}

// contextMenuClass is the struct that's finalized.
type contextMenuClass struct {
	native *C.WebKitContextMenuClass
}
