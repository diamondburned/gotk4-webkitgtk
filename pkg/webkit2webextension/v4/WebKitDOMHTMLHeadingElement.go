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
	GTypeDOMHTMLHeadingElement = coreglib.Type(C.webkit_dom_html_heading_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLHeadingElement, F: marshalDOMHTMLHeadingElement},
	})
}

// DOMHTMLHeadingElementOverrides contains methods that are overridable.
type DOMHTMLHeadingElementOverrides struct {
}

func defaultDOMHTMLHeadingElementOverrides(v *DOMHTMLHeadingElement) DOMHTMLHeadingElementOverrides {
	return DOMHTMLHeadingElementOverrides{}
}

type DOMHTMLHeadingElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLHeadingElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLHeadingElement, *DOMHTMLHeadingElementClass, DOMHTMLHeadingElementOverrides](
		GTypeDOMHTMLHeadingElement,
		initDOMHTMLHeadingElementClass,
		wrapDOMHTMLHeadingElement,
		defaultDOMHTMLHeadingElementOverrides,
	)
}

func initDOMHTMLHeadingElementClass(gclass unsafe.Pointer, overrides DOMHTMLHeadingElementOverrides, classInitFunc func(*DOMHTMLHeadingElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLHeadingElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLHeadingElement(obj *coreglib.Object) *DOMHTMLHeadingElement {
	return &DOMHTMLHeadingElement{
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

func marshalDOMHTMLHeadingElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLHeadingElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Align: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLHeadingElement) Align() string {
	var _arg0 *C.WebKitDOMHTMLHeadingElement // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.WebKitDOMHTMLHeadingElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_heading_element_get_align(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetAlign: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLHeadingElement) SetAlign(value string) {
	var _arg0 *C.WebKitDOMHTMLHeadingElement // out
	var _arg1 *C.gchar                       // out

	_arg0 = (*C.WebKitDOMHTMLHeadingElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_heading_element_set_align(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLHeadingElementClass: instance of this type is always passed by
// reference.
type DOMHTMLHeadingElementClass struct {
	*domhtmlHeadingElementClass
}

// domhtmlHeadingElementClass is the struct that's finalized.
type domhtmlHeadingElementClass struct {
	native *C.WebKitDOMHTMLHeadingElementClass
}

func (d *DOMHTMLHeadingElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
