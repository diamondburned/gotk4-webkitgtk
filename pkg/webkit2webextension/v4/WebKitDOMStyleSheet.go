// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

// GType values.
var (
	GTypeDOMStyleSheet = coreglib.Type(C.webkit_dom_style_sheet_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMStyleSheet, F: marshalDOMStyleSheet},
	})
}

// DOMStyleSheetOverrides contains methods that are overridable.
type DOMStyleSheetOverrides struct {
}

func defaultDOMStyleSheetOverrides(v *DOMStyleSheet) DOMStyleSheetOverrides {
	return DOMStyleSheetOverrides{}
}

type DOMStyleSheet struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMStyleSheet)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMStyleSheet, *DOMStyleSheetClass, DOMStyleSheetOverrides](
		GTypeDOMStyleSheet,
		initDOMStyleSheetClass,
		wrapDOMStyleSheet,
		defaultDOMStyleSheetOverrides,
	)
}

func initDOMStyleSheetClass(gclass unsafe.Pointer, overrides DOMStyleSheetOverrides, classInitFunc func(*DOMStyleSheetClass)) {
	if classInitFunc != nil {
		class := (*DOMStyleSheetClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMStyleSheet(obj *coreglib.Object) *DOMStyleSheet {
	return &DOMStyleSheet{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMStyleSheet(p uintptr) (interface{}, error) {
	return wrapDOMStyleSheet(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ContentType: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMStyleSheet) ContentType() string {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_get_content_type(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Disabled: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMStyleSheet) Disabled() bool {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_get_disabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Href: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMStyleSheet) Href() string {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_get_href(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Media: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domMediaList: KitDOMMediaList.
//
func (self *DOMStyleSheet) Media() *DOMMediaList {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _cret *C.WebKitDOMMediaList  // in

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_get_media(_arg0)
	runtime.KeepAlive(self)

	var _domMediaList *DOMMediaList // out

	_domMediaList = wrapDOMMediaList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domMediaList
}

// OwnerNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMStyleSheet) OwnerNode() *DOMNode {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_get_owner_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// ParentStyleSheet: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domStyleSheet: KitDOMStyleSheet.
//
func (self *DOMStyleSheet) ParentStyleSheet() *DOMStyleSheet {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _cret *C.WebKitDOMStyleSheet // in

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_get_parent_style_sheet(_arg0)
	runtime.KeepAlive(self)

	var _domStyleSheet *DOMStyleSheet // out

	_domStyleSheet = wrapDOMStyleSheet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domStyleSheet
}

// Title: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMStyleSheet) Title() string {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetDisabled: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMStyleSheet) SetDisabled(value bool) {
	var _arg0 *C.WebKitDOMStyleSheet // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.WebKitDOMStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_style_sheet_set_disabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMStyleSheetClass: instance of this type is always passed by reference.
type DOMStyleSheetClass struct {
	*domStyleSheetClass
}

// domStyleSheetClass is the struct that's finalized.
type domStyleSheetClass struct {
	native *C.WebKitDOMStyleSheetClass
}

func (d *DOMStyleSheetClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}