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
	GTypeDOMStyleSheetList = coreglib.Type(C.webkit_dom_style_sheet_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMStyleSheetList, F: marshalDOMStyleSheetList},
	})
}

// DOMStyleSheetListOverrides contains methods that are overridable.
type DOMStyleSheetListOverrides struct {
}

func defaultDOMStyleSheetListOverrides(v *DOMStyleSheetList) DOMStyleSheetListOverrides {
	return DOMStyleSheetListOverrides{}
}

type DOMStyleSheetList struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMStyleSheetList)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMStyleSheetList, *DOMStyleSheetListClass, DOMStyleSheetListOverrides](
		GTypeDOMStyleSheetList,
		initDOMStyleSheetListClass,
		wrapDOMStyleSheetList,
		defaultDOMStyleSheetListOverrides,
	)
}

func initDOMStyleSheetListClass(gclass unsafe.Pointer, overrides DOMStyleSheetListOverrides, classInitFunc func(*DOMStyleSheetListClass)) {
	if classInitFunc != nil {
		class := (*DOMStyleSheetListClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMStyleSheetList(obj *coreglib.Object) *DOMStyleSheetList {
	return &DOMStyleSheetList{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMStyleSheetList(p uintptr) (interface{}, error) {
	return wrapDOMStyleSheetList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Length: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMStyleSheetList) Length() uint32 {
	var _arg0 *C.WebKitDOMStyleSheetList // out
	var _cret C.gulong                   // in

	_arg0 = (*C.WebKitDOMStyleSheetList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_style_sheet_list_get_length(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// Item: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - index: #gulong.
//
// The function returns the following values:
//
//   - domStyleSheet: KitDOMStyleSheet.
//
func (self *DOMStyleSheetList) Item(index uint32) *DOMStyleSheet {
	var _arg0 *C.WebKitDOMStyleSheetList // out
	var _arg1 C.gulong                   // out
	var _cret *C.WebKitDOMStyleSheet     // in

	_arg0 = (*C.WebKitDOMStyleSheetList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	_cret = C.webkit_dom_style_sheet_list_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _domStyleSheet *DOMStyleSheet // out

	_domStyleSheet = wrapDOMStyleSheet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domStyleSheet
}

// DOMStyleSheetListClass: instance of this type is always passed by reference.
type DOMStyleSheetListClass struct {
	*domStyleSheetListClass
}

// domStyleSheetListClass is the struct that's finalized.
type domStyleSheetListClass struct {
	native *C.WebKitDOMStyleSheetListClass
}

func (d *DOMStyleSheetListClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
