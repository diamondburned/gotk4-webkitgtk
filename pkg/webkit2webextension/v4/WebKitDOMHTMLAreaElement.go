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
	GTypeDOMHTMLAreaElement = coreglib.Type(C.webkit_dom_html_area_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLAreaElement, F: marshalDOMHTMLAreaElement},
	})
}

// DOMHTMLAreaElementOverrides contains methods that are overridable.
type DOMHTMLAreaElementOverrides struct {
}

func defaultDOMHTMLAreaElementOverrides(v *DOMHTMLAreaElement) DOMHTMLAreaElementOverrides {
	return DOMHTMLAreaElementOverrides{}
}

type DOMHTMLAreaElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLAreaElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLAreaElement, *DOMHTMLAreaElementClass, DOMHTMLAreaElementOverrides](
		GTypeDOMHTMLAreaElement,
		initDOMHTMLAreaElementClass,
		wrapDOMHTMLAreaElement,
		defaultDOMHTMLAreaElementOverrides,
	)
}

func initDOMHTMLAreaElementClass(gclass unsafe.Pointer, overrides DOMHTMLAreaElementOverrides, classInitFunc func(*DOMHTMLAreaElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLAreaElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLAreaElement(obj *coreglib.Object) *DOMHTMLAreaElement {
	return &DOMHTMLAreaElement{
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

func marshalDOMHTMLAreaElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLAreaElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Alt: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Alt() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_alt(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Coords: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Coords() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_coords(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Hash: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Hash() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_hash(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Host: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Host() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_host(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Hostname: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Hostname() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_hostname(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Href: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Href() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_href(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// NoHref: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLAreaElement) NoHref() bool {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_no_href(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Pathname: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Pathname() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_pathname(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Port: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Port() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_port(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Protocol: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Protocol() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_protocol(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Search: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Search() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_search(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Shape: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Shape() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_shape(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Target: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLAreaElement) Target() string {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_area_element_get_target(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetAlt: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetAlt(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_alt(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetCoords: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetCoords(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_coords(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetHash: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetHash(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_hash(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetHost: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetHost(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_host(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetHostname: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetHostname(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_hostname(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetHref: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetHref(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_href(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetNoHref: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLAreaElement) SetNoHref(value bool) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 C.gboolean                  // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_area_element_set_no_href(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetPathname: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetPathname(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_pathname(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetPort: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetPort(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_port(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetProtocol: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetProtocol(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_protocol(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetSearch: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetSearch(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_search(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetShape: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetShape(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_shape(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetTarget: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLAreaElement) SetTarget(value string) {
	var _arg0 *C.WebKitDOMHTMLAreaElement // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitDOMHTMLAreaElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_area_element_set_target(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLAreaElementClass: instance of this type is always passed by reference.
type DOMHTMLAreaElementClass struct {
	*domhtmlAreaElementClass
}

// domhtmlAreaElementClass is the struct that's finalized.
type domhtmlAreaElementClass struct {
	native *C.WebKitDOMHTMLAreaElementClass
}

func (d *DOMHTMLAreaElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
