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
	GTypeDOMHTMLStyleElement = coreglib.Type(C.webkit_dom_html_style_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLStyleElement, F: marshalDOMHTMLStyleElement},
	})
}

// DOMHTMLStyleElementOverrides contains methods that are overridable.
type DOMHTMLStyleElementOverrides struct {
}

func defaultDOMHTMLStyleElementOverrides(v *DOMHTMLStyleElement) DOMHTMLStyleElementOverrides {
	return DOMHTMLStyleElementOverrides{}
}

type DOMHTMLStyleElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLStyleElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLStyleElement, *DOMHTMLStyleElementClass, DOMHTMLStyleElementOverrides](
		GTypeDOMHTMLStyleElement,
		initDOMHTMLStyleElementClass,
		wrapDOMHTMLStyleElement,
		defaultDOMHTMLStyleElementOverrides,
	)
}

func initDOMHTMLStyleElementClass(gclass unsafe.Pointer, overrides DOMHTMLStyleElementOverrides, classInitFunc func(*DOMHTMLStyleElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLStyleElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLStyleElement(obj *coreglib.Object) *DOMHTMLStyleElement {
	return &DOMHTMLStyleElement{
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

func marshalDOMHTMLStyleElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLStyleElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Disabled: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLStyleElement) Disabled() bool {
	var _arg0 *C.WebKitDOMHTMLStyleElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLStyleElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_style_element_get_disabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Media: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLStyleElement) Media() string {
	var _arg0 *C.WebKitDOMHTMLStyleElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLStyleElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_style_element_get_media(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Sheet: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domStyleSheet: KitDOMStyleSheet.
//
func (self *DOMHTMLStyleElement) Sheet() *DOMStyleSheet {
	var _arg0 *C.WebKitDOMHTMLStyleElement // out
	var _cret *C.WebKitDOMStyleSheet       // in

	_arg0 = (*C.WebKitDOMHTMLStyleElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_style_element_get_sheet(_arg0)
	runtime.KeepAlive(self)

	var _domStyleSheet *DOMStyleSheet // out

	_domStyleSheet = wrapDOMStyleSheet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domStyleSheet
}

// TypeAttr: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLStyleElement) TypeAttr() string {
	var _arg0 *C.WebKitDOMHTMLStyleElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLStyleElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_style_element_get_type_attr(_arg0)
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
func (self *DOMHTMLStyleElement) SetDisabled(value bool) {
	var _arg0 *C.WebKitDOMHTMLStyleElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLStyleElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_style_element_set_disabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetMedia: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLStyleElement) SetMedia(value string) {
	var _arg0 *C.WebKitDOMHTMLStyleElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLStyleElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_style_element_set_media(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetTypeAttr: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLStyleElement) SetTypeAttr(value string) {
	var _arg0 *C.WebKitDOMHTMLStyleElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLStyleElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_style_element_set_type_attr(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLStyleElementClass: instance of this type is always passed by
// reference.
type DOMHTMLStyleElementClass struct {
	*domhtmlStyleElementClass
}

// domhtmlStyleElementClass is the struct that's finalized.
type domhtmlStyleElementClass struct {
	native *C.WebKitDOMHTMLStyleElementClass
}

func (d *DOMHTMLStyleElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
