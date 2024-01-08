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
	GTypeDOMNodeList = coreglib.Type(C.webkit_dom_node_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMNodeList, F: marshalDOMNodeList},
	})
}

// DOMNodeListOverrides contains methods that are overridable.
type DOMNodeListOverrides struct {
}

func defaultDOMNodeListOverrides(v *DOMNodeList) DOMNodeListOverrides {
	return DOMNodeListOverrides{}
}

type DOMNodeList struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMNodeList)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMNodeList, *DOMNodeListClass, DOMNodeListOverrides](
		GTypeDOMNodeList,
		initDOMNodeListClass,
		wrapDOMNodeList,
		defaultDOMNodeListOverrides,
	)
}

func initDOMNodeListClass(gclass unsafe.Pointer, overrides DOMNodeListOverrides, classInitFunc func(*DOMNodeListClass)) {
	if classInitFunc != nil {
		class := (*DOMNodeListClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMNodeList(obj *coreglib.Object) *DOMNodeList {
	return &DOMNodeList{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMNodeList(p uintptr) (interface{}, error) {
	return wrapDOMNodeList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Length: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMNodeList) Length() uint32 {
	var _arg0 *C.WebKitDOMNodeList // out
	var _cret C.gulong             // in

	_arg0 = (*C.WebKitDOMNodeList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_list_get_length(_arg0)
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
//   - domNode: KitDOMNode.
//
func (self *DOMNodeList) Item(index uint32) *DOMNode {
	var _arg0 *C.WebKitDOMNodeList // out
	var _arg1 C.gulong             // out
	var _cret *C.WebKitDOMNode     // in

	_arg0 = (*C.WebKitDOMNodeList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	_cret = C.webkit_dom_node_list_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// DOMNodeListClass: instance of this type is always passed by reference.
type DOMNodeListClass struct {
	*domNodeListClass
}

// domNodeListClass is the struct that's finalized.
type domNodeListClass struct {
	native *C.WebKitDOMNodeListClass
}

func (d *DOMNodeListClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}