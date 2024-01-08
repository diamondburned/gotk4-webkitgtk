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
// extern void _gotk4_webkit6_BackForwardList_ConnectChanged(gpointer, WebKitBackForwardListItem*, gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeBackForwardList = coreglib.Type(C.webkit_back_forward_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBackForwardList, F: marshalBackForwardList},
	})
}

// BackForwardListOverrides contains methods that are overridable.
type BackForwardListOverrides struct {
}

func defaultBackForwardListOverrides(v *BackForwardList) BackForwardListOverrides {
	return BackForwardListOverrides{}
}

// BackForwardList: list of visited pages.
//
// WebKitBackForwardList maintains a list of visited pages used to navigate to
// recent pages. Items are inserted in the list in the order they are visited.
//
// WebKitBackForwardList also maintains the notion of the current item (which
// is always at index 0), the preceding item (which is at index -1), and the
// following item (which is at index 1). Methods webkit_web_view_go_back()
// and webkit_web_view_go_forward() move the current item backward or
// forward by one. Method webkit_web_view_go_to_back_forward_list_item()
// sets the current item to the specified item. All other methods returning
// KitBackForwardListItem<!-- -->s do not change the value of the current item,
// they just return the requested item or items.
type BackForwardList struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*BackForwardList)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*BackForwardList, *BackForwardListClass, BackForwardListOverrides](
		GTypeBackForwardList,
		initBackForwardListClass,
		wrapBackForwardList,
		defaultBackForwardListOverrides,
	)
}

func initBackForwardListClass(gclass unsafe.Pointer, overrides BackForwardListOverrides, classInitFunc func(*BackForwardListClass)) {
	if classInitFunc != nil {
		class := (*BackForwardListClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBackForwardList(obj *coreglib.Object) *BackForwardList {
	return &BackForwardList{
		Object: obj,
	}
}

func marshalBackForwardList(p uintptr) (interface{}, error) {
	return wrapBackForwardList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged: this signal is emitted when back_forward_list changes.
// This happens when the current item is updated, a new item is added or one or
// more items are removed. Note that both item_added and items_removed can NULL
// when only the current item is updated. Items are only removed when the list
// is cleared or the maximum items limit is reached.
func (backForwardList *BackForwardList) ConnectChanged(f func(itemAdded *BackForwardListItem, itemsRemoved unsafe.Pointer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(backForwardList, "changed", false, unsafe.Pointer(C._gotk4_webkit6_BackForwardList_ConnectChanged), f)
}

// BackItem returns the item that precedes the current item.
//
// The function returns the following values:
//
//   - backForwardListItem (optional): KitBackForwardListItem preceding the
//     current item or NULL.
//
func (backForwardList *BackForwardList) BackItem() *BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList     // out
	var _cret *C.WebKitBackForwardListItem // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))

	_cret = C.webkit_back_forward_list_get_back_item(_arg0)
	runtime.KeepAlive(backForwardList)

	var _backForwardListItem *BackForwardListItem // out

	if _cret != nil {
		_backForwardListItem = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _backForwardListItem
}

// BackList: obtain the list of items preceding the current one.
//
// The function returns the following values:
//
//   - list of items preceding the current item.
//
func (backForwardList *BackForwardList) BackList() []*BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList // out
	var _cret *C.GList                 // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))

	_cret = C.webkit_back_forward_list_get_back_list(_arg0)
	runtime.KeepAlive(backForwardList)

	var _list []*BackForwardListItem // out

	_list = make([]*BackForwardListItem, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.WebKitBackForwardListItem)(v)
		var dst *BackForwardListItem // out
		dst = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// BackListWithLimit: obtain a list up to some number of items preceding the
// current one.
//
// The function takes the following parameters:
//
//   - limit: number of items to retrieve.
//
// The function returns the following values:
//
//   - list of items preceding the current item limited by limit.
//
func (backForwardList *BackForwardList) BackListWithLimit(limit uint) []*BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList // out
	var _arg1 C.guint                  // out
	var _cret *C.GList                 // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))
	_arg1 = C.guint(limit)

	_cret = C.webkit_back_forward_list_get_back_list_with_limit(_arg0, _arg1)
	runtime.KeepAlive(backForwardList)
	runtime.KeepAlive(limit)

	var _list []*BackForwardListItem // out

	_list = make([]*BackForwardListItem, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.WebKitBackForwardListItem)(v)
		var dst *BackForwardListItem // out
		dst = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// CurrentItem returns the current item in back_forward_list.
//
// The function returns the following values:
//
//   - backForwardListItem (optional): KitBackForwardListItem or NULL if
//     back_forward_list is empty.
//
func (backForwardList *BackForwardList) CurrentItem() *BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList     // out
	var _cret *C.WebKitBackForwardListItem // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))

	_cret = C.webkit_back_forward_list_get_current_item(_arg0)
	runtime.KeepAlive(backForwardList)

	var _backForwardListItem *BackForwardListItem // out

	if _cret != nil {
		_backForwardListItem = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _backForwardListItem
}

// ForwardItem returns the item that follows the current item.
//
// The function returns the following values:
//
//   - backForwardListItem (optional): KitBackForwardListItem following the
//     current item or NULL.
//
func (backForwardList *BackForwardList) ForwardItem() *BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList     // out
	var _cret *C.WebKitBackForwardListItem // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))

	_cret = C.webkit_back_forward_list_get_forward_item(_arg0)
	runtime.KeepAlive(backForwardList)

	var _backForwardListItem *BackForwardListItem // out

	if _cret != nil {
		_backForwardListItem = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _backForwardListItem
}

// ForwardList: obtain the list of items following the current one.
//
// The function returns the following values:
//
//   - list of items following the current item.
//
func (backForwardList *BackForwardList) ForwardList() []*BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList // out
	var _cret *C.GList                 // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))

	_cret = C.webkit_back_forward_list_get_forward_list(_arg0)
	runtime.KeepAlive(backForwardList)

	var _list []*BackForwardListItem // out

	_list = make([]*BackForwardListItem, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.WebKitBackForwardListItem)(v)
		var dst *BackForwardListItem // out
		dst = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// ForwardListWithLimit: obtain a list up to some number of items following the
// current one.
//
// The function takes the following parameters:
//
//   - limit: number of items to retrieve.
//
// The function returns the following values:
//
//   - list of items following the current item limited by limit.
//
func (backForwardList *BackForwardList) ForwardListWithLimit(limit uint) []*BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList // out
	var _arg1 C.guint                  // out
	var _cret *C.GList                 // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))
	_arg1 = C.guint(limit)

	_cret = C.webkit_back_forward_list_get_forward_list_with_limit(_arg0, _arg1)
	runtime.KeepAlive(backForwardList)
	runtime.KeepAlive(limit)

	var _list []*BackForwardListItem // out

	_list = make([]*BackForwardListItem, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.WebKitBackForwardListItem)(v)
		var dst *BackForwardListItem // out
		dst = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// Length: obtain the amount of items in the list.
//
// The function returns the following values:
//
//   - guint: length of back_forward_list.
//
func (backForwardList *BackForwardList) Length() uint {
	var _arg0 *C.WebKitBackForwardList // out
	var _cret C.guint                  // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))

	_cret = C.webkit_back_forward_list_get_length(_arg0)
	runtime.KeepAlive(backForwardList)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NthItem returns the item at a given index relative to the current item.
//
// The function takes the following parameters:
//
//   - index of the item.
//
// The function returns the following values:
//
//   - backForwardListItem (optional): KitBackForwardListItem located at the
//     specified index relative to the current item or NULL.
//
func (backForwardList *BackForwardList) NthItem(index int) *BackForwardListItem {
	var _arg0 *C.WebKitBackForwardList     // out
	var _arg1 C.gint                       // out
	var _cret *C.WebKitBackForwardListItem // in

	_arg0 = (*C.WebKitBackForwardList)(unsafe.Pointer(coreglib.InternObject(backForwardList).Native()))
	_arg1 = C.gint(index)

	_cret = C.webkit_back_forward_list_get_nth_item(_arg0, _arg1)
	runtime.KeepAlive(backForwardList)
	runtime.KeepAlive(index)

	var _backForwardListItem *BackForwardListItem // out

	if _cret != nil {
		_backForwardListItem = wrapBackForwardListItem(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _backForwardListItem
}

// BackForwardListClass: instance of this type is always passed by reference.
type BackForwardListClass struct {
	*backForwardListClass
}

// backForwardListClass is the struct that's finalized.
type backForwardListClass struct {
	native *C.WebKitBackForwardListClass
}
