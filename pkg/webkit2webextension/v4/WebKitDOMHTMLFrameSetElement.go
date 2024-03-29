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
	GTypeDOMHTMLFrameSetElement = coreglib.Type(C.webkit_dom_html_frame_set_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLFrameSetElement, F: marshalDOMHTMLFrameSetElement},
	})
}

// DOMHTMLFrameSetElementOverrides contains methods that are overridable.
type DOMHTMLFrameSetElementOverrides struct {
}

func defaultDOMHTMLFrameSetElementOverrides(v *DOMHTMLFrameSetElement) DOMHTMLFrameSetElementOverrides {
	return DOMHTMLFrameSetElementOverrides{}
}

type DOMHTMLFrameSetElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLFrameSetElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLFrameSetElement, *DOMHTMLFrameSetElementClass, DOMHTMLFrameSetElementOverrides](
		GTypeDOMHTMLFrameSetElement,
		initDOMHTMLFrameSetElementClass,
		wrapDOMHTMLFrameSetElement,
		defaultDOMHTMLFrameSetElementOverrides,
	)
}

func initDOMHTMLFrameSetElementClass(gclass unsafe.Pointer, overrides DOMHTMLFrameSetElementOverrides, classInitFunc func(*DOMHTMLFrameSetElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLFrameSetElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLFrameSetElement(obj *coreglib.Object) *DOMHTMLFrameSetElement {
	return &DOMHTMLFrameSetElement{
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

func marshalDOMHTMLFrameSetElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLFrameSetElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Cols: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLFrameSetElement) Cols() string {
	var _arg0 *C.WebKitDOMHTMLFrameSetElement // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitDOMHTMLFrameSetElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_frame_set_element_get_cols(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Rows: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLFrameSetElement) Rows() string {
	var _arg0 *C.WebKitDOMHTMLFrameSetElement // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitDOMHTMLFrameSetElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_frame_set_element_get_rows(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetCols: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLFrameSetElement) SetCols(value string) {
	var _arg0 *C.WebKitDOMHTMLFrameSetElement // out
	var _arg1 *C.gchar                        // out

	_arg0 = (*C.WebKitDOMHTMLFrameSetElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_frame_set_element_set_cols(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetRows: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLFrameSetElement) SetRows(value string) {
	var _arg0 *C.WebKitDOMHTMLFrameSetElement // out
	var _arg1 *C.gchar                        // out

	_arg0 = (*C.WebKitDOMHTMLFrameSetElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_frame_set_element_set_rows(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLFrameSetElementClass: instance of this type is always passed by
// reference.
type DOMHTMLFrameSetElementClass struct {
	*domhtmlFrameSetElementClass
}

// domhtmlFrameSetElementClass is the struct that's finalized.
type domhtmlFrameSetElementClass struct {
	native *C.WebKitDOMHTMLFrameSetElementClass
}

func (d *DOMHTMLFrameSetElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
