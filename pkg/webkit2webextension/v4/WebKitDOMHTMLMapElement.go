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
	GTypeDOMHTMLMapElement = coreglib.Type(C.webkit_dom_html_map_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLMapElement, F: marshalDOMHTMLMapElement},
	})
}

// DOMHTMLMapElementOverrides contains methods that are overridable.
type DOMHTMLMapElementOverrides struct {
}

func defaultDOMHTMLMapElementOverrides(v *DOMHTMLMapElement) DOMHTMLMapElementOverrides {
	return DOMHTMLMapElementOverrides{}
}

type DOMHTMLMapElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLMapElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLMapElement, *DOMHTMLMapElementClass, DOMHTMLMapElementOverrides](
		GTypeDOMHTMLMapElement,
		initDOMHTMLMapElementClass,
		wrapDOMHTMLMapElement,
		defaultDOMHTMLMapElementOverrides,
	)
}

func initDOMHTMLMapElementClass(gclass unsafe.Pointer, overrides DOMHTMLMapElementOverrides, classInitFunc func(*DOMHTMLMapElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLMapElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLMapElement(obj *coreglib.Object) *DOMHTMLMapElement {
	return &DOMHTMLMapElement{
		DOMHTMLElement: DOMHTMLElement{
			DOMElement: DOMElement{
				DOMNode: DOMNode{
					DOMObject: DOMObject{
						Object: obj,
					},
					Object: obj,
					DOMEventTarget: DOMEventTarget{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalDOMHTMLMapElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLMapElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Areas: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domhtmlCollection: KitDOMHTMLCollection.
//
func (self *DOMHTMLMapElement) Areas() *DOMHTMLCollection {
	var _arg0 *C.WebKitDOMHTMLMapElement // out
	var _cret *C.WebKitDOMHTMLCollection // in

	_arg0 = (*C.WebKitDOMHTMLMapElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_map_element_get_areas(_arg0)
	runtime.KeepAlive(self)

	var _domhtmlCollection *DOMHTMLCollection // out

	_domhtmlCollection = wrapDOMHTMLCollection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domhtmlCollection
}

// Name: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLMapElement) Name() string {
	var _arg0 *C.WebKitDOMHTMLMapElement // out
	var _cret *C.gchar                   // in

	_arg0 = (*C.WebKitDOMHTMLMapElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_map_element_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetName: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLMapElement) SetName(value string) {
	var _arg0 *C.WebKitDOMHTMLMapElement // out
	var _arg1 *C.gchar                   // out

	_arg0 = (*C.WebKitDOMHTMLMapElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_map_element_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLMapElementClass: instance of this type is always passed by reference.
type DOMHTMLMapElementClass struct {
	*domhtmlMapElementClass
}

// domhtmlMapElementClass is the struct that's finalized.
type domhtmlMapElementClass struct {
	native *C.WebKitDOMHTMLMapElementClass
}

func (d *DOMHTMLMapElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}