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
	GTypeBackForwardListItem = coreglib.Type(C.webkit_back_forward_list_item_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBackForwardListItem, F: marshalBackForwardListItem},
	})
}

// BackForwardListItemOverrides contains methods that are overridable.
type BackForwardListItemOverrides struct {
}

func defaultBackForwardListItemOverrides(v *BackForwardListItem) BackForwardListItemOverrides {
	return BackForwardListItemOverrides{}
}

// BackForwardListItem: one item of the KitBackForwardList.
//
// A history item is part of the KitBackForwardList and consists out of a title
// and a URI.
type BackForwardListItem struct {
	_ [0]func() // equal guard
	coreglib.InitiallyUnowned
}

var ()

func init() {
	coreglib.RegisterClassInfo[*BackForwardListItem, *BackForwardListItemClass, BackForwardListItemOverrides](
		GTypeBackForwardListItem,
		initBackForwardListItemClass,
		wrapBackForwardListItem,
		defaultBackForwardListItemOverrides,
	)
}

func initBackForwardListItemClass(gclass unsafe.Pointer, overrides BackForwardListItemOverrides, classInitFunc func(*BackForwardListItemClass)) {
	if classInitFunc != nil {
		class := (*BackForwardListItemClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBackForwardListItem(obj *coreglib.Object) *BackForwardListItem {
	return &BackForwardListItem{
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalBackForwardListItem(p uintptr) (interface{}, error) {
	return wrapBackForwardListItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// OriginalURI: obtain the original URI of the item.
//
// See also webkit_back_forward_list_item_get_uri().
//
// The function returns the following values:
//
//   - utf8: original URI of list_item or NULL when the original URI is empty.
//
func (listItem *BackForwardListItem) OriginalURI() string {
	var _arg0 *C.WebKitBackForwardListItem // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitBackForwardListItem)(unsafe.Pointer(coreglib.InternObject(listItem).Native()))

	_cret = C.webkit_back_forward_list_item_get_original_uri(_arg0)
	runtime.KeepAlive(listItem)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Title: obtain the title of the item.
//
// The function returns the following values:
//
//   - utf8: page title of list_item or NULL when the title is empty.
//
func (listItem *BackForwardListItem) Title() string {
	var _arg0 *C.WebKitBackForwardListItem // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitBackForwardListItem)(unsafe.Pointer(coreglib.InternObject(listItem).Native()))

	_cret = C.webkit_back_forward_list_item_get_title(_arg0)
	runtime.KeepAlive(listItem)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URI: obtain the URI of the item.
//
// This URI may differ from the original URI if the page was,
// for example, redirected to a new location. See also
// webkit_back_forward_list_item_get_original_uri().
//
// The function returns the following values:
//
//   - utf8: URI of list_item or NULL when the URI is empty.
//
func (listItem *BackForwardListItem) URI() string {
	var _arg0 *C.WebKitBackForwardListItem // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitBackForwardListItem)(unsafe.Pointer(coreglib.InternObject(listItem).Native()))

	_cret = C.webkit_back_forward_list_item_get_uri(_arg0)
	runtime.KeepAlive(listItem)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// BackForwardListItemClass: instance of this type is always passed by
// reference.
type BackForwardListItemClass struct {
	*backForwardListItemClass
}

// backForwardListItemClass is the struct that's finalized.
type backForwardListItemClass struct {
	native *C.WebKitBackForwardListItemClass
}
