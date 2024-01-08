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
	GTypeDOMHTMLTableColElement = coreglib.Type(C.webkit_dom_html_table_col_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLTableColElement, F: marshalDOMHTMLTableColElement},
	})
}

// DOMHTMLTableColElementOverrides contains methods that are overridable.
type DOMHTMLTableColElementOverrides struct {
}

func defaultDOMHTMLTableColElementOverrides(v *DOMHTMLTableColElement) DOMHTMLTableColElementOverrides {
	return DOMHTMLTableColElementOverrides{}
}

type DOMHTMLTableColElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLTableColElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLTableColElement, *DOMHTMLTableColElementClass, DOMHTMLTableColElementOverrides](
		GTypeDOMHTMLTableColElement,
		initDOMHTMLTableColElementClass,
		wrapDOMHTMLTableColElement,
		defaultDOMHTMLTableColElementOverrides,
	)
}

func initDOMHTMLTableColElementClass(gclass unsafe.Pointer, overrides DOMHTMLTableColElementOverrides, classInitFunc func(*DOMHTMLTableColElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLTableColElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLTableColElement(obj *coreglib.Object) *DOMHTMLTableColElement {
	return &DOMHTMLTableColElement{
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

func marshalDOMHTMLTableColElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLTableColElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Align: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLTableColElement) Align() string {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_table_col_element_get_align(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Ch: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLTableColElement) Ch() string {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_table_col_element_get_ch(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ChOff: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLTableColElement) ChOff() string {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_table_col_element_get_ch_off(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Span: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - glong: #glong.
//
func (self *DOMHTMLTableColElement) Span() int32 {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _cret C.glong                         // in

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_table_col_element_get_span(_arg0)
	runtime.KeepAlive(self)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// VAlign: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLTableColElement) VAlign() string {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_table_col_element_get_v_align(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Width: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLTableColElement) Width() string {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _cret *C.gchar                        // in

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_table_col_element_get_width(_arg0)
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
func (self *DOMHTMLTableColElement) SetAlign(value string) {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _arg1 *C.gchar                        // out

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_table_col_element_set_align(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetCh: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLTableColElement) SetCh(value string) {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _arg1 *C.gchar                        // out

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_table_col_element_set_ch(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetChOff: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLTableColElement) SetChOff(value string) {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _arg1 *C.gchar                        // out

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_table_col_element_set_ch_off(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetSpan: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #glong.
//
func (self *DOMHTMLTableColElement) SetSpan(value int32) {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _arg1 C.glong                         // out

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.glong(value)

	C.webkit_dom_html_table_col_element_set_span(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetVAlign: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLTableColElement) SetVAlign(value string) {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _arg1 *C.gchar                        // out

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_table_col_element_set_v_align(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetWidth: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLTableColElement) SetWidth(value string) {
	var _arg0 *C.WebKitDOMHTMLTableColElement // out
	var _arg1 *C.gchar                        // out

	_arg0 = (*C.WebKitDOMHTMLTableColElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_table_col_element_set_width(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLTableColElementClass: instance of this type is always passed by
// reference.
type DOMHTMLTableColElementClass struct {
	*domhtmlTableColElementClass
}

// domhtmlTableColElementClass is the struct that's finalized.
type domhtmlTableColElementClass struct {
	native *C.WebKitDOMHTMLTableColElementClass
}

func (d *DOMHTMLTableColElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}