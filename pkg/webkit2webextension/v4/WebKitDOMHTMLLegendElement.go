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
	GTypeDOMHTMLLegendElement = coreglib.Type(C.webkit_dom_html_legend_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLLegendElement, F: marshalDOMHTMLLegendElement},
	})
}

// DOMHTMLLegendElementOverrides contains methods that are overridable.
type DOMHTMLLegendElementOverrides struct {
}

func defaultDOMHTMLLegendElementOverrides(v *DOMHTMLLegendElement) DOMHTMLLegendElementOverrides {
	return DOMHTMLLegendElementOverrides{}
}

type DOMHTMLLegendElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLLegendElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLLegendElement, *DOMHTMLLegendElementClass, DOMHTMLLegendElementOverrides](
		GTypeDOMHTMLLegendElement,
		initDOMHTMLLegendElementClass,
		wrapDOMHTMLLegendElement,
		defaultDOMHTMLLegendElementOverrides,
	)
}

func initDOMHTMLLegendElementClass(gclass unsafe.Pointer, overrides DOMHTMLLegendElementOverrides, classInitFunc func(*DOMHTMLLegendElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLLegendElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLLegendElement(obj *coreglib.Object) *DOMHTMLLegendElement {
	return &DOMHTMLLegendElement{
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

func marshalDOMHTMLLegendElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLLegendElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Align: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLLegendElement) Align() string {
	var _arg0 *C.WebKitDOMHTMLLegendElement // out
	var _cret *C.gchar                      // in

	_arg0 = (*C.WebKitDOMHTMLLegendElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_legend_element_get_align(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Form: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domhtmlFormElement: KitDOMHTMLFormElement.
//
func (self *DOMHTMLLegendElement) Form() *DOMHTMLFormElement {
	var _arg0 *C.WebKitDOMHTMLLegendElement // out
	var _cret *C.WebKitDOMHTMLFormElement   // in

	_arg0 = (*C.WebKitDOMHTMLLegendElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_legend_element_get_form(_arg0)
	runtime.KeepAlive(self)

	var _domhtmlFormElement *DOMHTMLFormElement // out

	_domhtmlFormElement = wrapDOMHTMLFormElement(coreglib.Take(unsafe.Pointer(_cret)))

	return _domhtmlFormElement
}

// SetAlign: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLLegendElement) SetAlign(value string) {
	var _arg0 *C.WebKitDOMHTMLLegendElement // out
	var _arg1 *C.gchar                      // out

	_arg0 = (*C.WebKitDOMHTMLLegendElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_legend_element_set_align(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLLegendElementClass: instance of this type is always passed by
// reference.
type DOMHTMLLegendElementClass struct {
	*domhtmlLegendElementClass
}

// domhtmlLegendElementClass is the struct that's finalized.
type domhtmlLegendElementClass struct {
	native *C.WebKitDOMHTMLLegendElementClass
}

func (d *DOMHTMLLegendElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}