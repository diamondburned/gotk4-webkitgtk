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
	GTypeDOMHTMLButtonElement = coreglib.Type(C.webkit_dom_html_button_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLButtonElement, F: marshalDOMHTMLButtonElement},
	})
}

// DOMHTMLButtonElementOverrides contains methods that are overridable.
type DOMHTMLButtonElementOverrides struct {
}

func defaultDOMHTMLButtonElementOverrides(v *DOMHTMLButtonElement) DOMHTMLButtonElementOverrides {
	return DOMHTMLButtonElementOverrides{}
}

type DOMHTMLButtonElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLButtonElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLButtonElement, *DOMHTMLButtonElementClass, DOMHTMLButtonElementOverrides](
		GTypeDOMHTMLButtonElement,
		initDOMHTMLButtonElementClass,
		wrapDOMHTMLButtonElement,
		defaultDOMHTMLButtonElementOverrides,
	)
}

func initDOMHTMLButtonElementClass(gclass unsafe.Pointer, overrides DOMHTMLButtonElementOverrides, classInitFunc func(*DOMHTMLButtonElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLButtonElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLButtonElement(obj *coreglib.Object) *DOMHTMLButtonElement {
	return &DOMHTMLButtonElement{
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

func marshalDOMHTMLButtonElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLButtonElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Autofocus: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLButtonElement) Autofocus() bool {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _cret C.gboolean                    // in

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_button_element_get_autofocus(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ButtonType: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLButtonElement) ButtonType() string {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _cret *C.gchar                      // in

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_button_element_get_button_type(_arg0)
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
func (self *DOMHTMLButtonElement) Disabled() bool {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _cret C.gboolean                    // in

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_button_element_get_disabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Form: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domhtmlFormElement: KitDOMHTMLFormElement.
//
func (self *DOMHTMLButtonElement) Form() *DOMHTMLFormElement {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _cret *C.WebKitDOMHTMLFormElement   // in

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_button_element_get_form(_arg0)
	runtime.KeepAlive(self)

	var _domhtmlFormElement *DOMHTMLFormElement // out

	_domhtmlFormElement = wrapDOMHTMLFormElement(coreglib.Take(unsafe.Pointer(_cret)))

	return _domhtmlFormElement
}

// Name: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLButtonElement) Name() string {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _cret *C.gchar                      // in

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_button_element_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Value: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLButtonElement) Value() string {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _cret *C.gchar                      // in

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_button_element_get_value(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// WillValidate: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLButtonElement) WillValidate() bool {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _cret C.gboolean                    // in

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_button_element_get_will_validate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAutofocus: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLButtonElement) SetAutofocus(value bool) {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _arg1 C.gboolean                    // out

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_button_element_set_autofocus(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetButtonType: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLButtonElement) SetButtonType(value string) {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _arg1 *C.gchar                      // out

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_button_element_set_button_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetDisabled: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLButtonElement) SetDisabled(value bool) {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _arg1 C.gboolean                    // out

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_button_element_set_disabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetName: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLButtonElement) SetName(value string) {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _arg1 *C.gchar                      // out

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_button_element_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetValue: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLButtonElement) SetValue(value string) {
	var _arg0 *C.WebKitDOMHTMLButtonElement // out
	var _arg1 *C.gchar                      // out

	_arg0 = (*C.WebKitDOMHTMLButtonElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_button_element_set_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLButtonElementClass: instance of this type is always passed by
// reference.
type DOMHTMLButtonElementClass struct {
	*domhtmlButtonElementClass
}

// domhtmlButtonElementClass is the struct that's finalized.
type domhtmlButtonElementClass struct {
	native *C.WebKitDOMHTMLButtonElementClass
}

func (d *DOMHTMLButtonElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
