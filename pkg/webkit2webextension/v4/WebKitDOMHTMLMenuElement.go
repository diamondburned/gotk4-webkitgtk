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
	GTypeDOMHTMLMenuElement = coreglib.Type(C.webkit_dom_html_menu_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLMenuElement, F: marshalDOMHTMLMenuElement},
	})
}

// DOMHTMLMenuElementOverrides contains methods that are overridable.
type DOMHTMLMenuElementOverrides struct {
}

func defaultDOMHTMLMenuElementOverrides(v *DOMHTMLMenuElement) DOMHTMLMenuElementOverrides {
	return DOMHTMLMenuElementOverrides{}
}

type DOMHTMLMenuElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLMenuElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLMenuElement, *DOMHTMLMenuElementClass, DOMHTMLMenuElementOverrides](
		GTypeDOMHTMLMenuElement,
		initDOMHTMLMenuElementClass,
		wrapDOMHTMLMenuElement,
		defaultDOMHTMLMenuElementOverrides,
	)
}

func initDOMHTMLMenuElementClass(gclass unsafe.Pointer, overrides DOMHTMLMenuElementOverrides, classInitFunc func(*DOMHTMLMenuElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLMenuElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLMenuElement(obj *coreglib.Object) *DOMHTMLMenuElement {
	return &DOMHTMLMenuElement{
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

func marshalDOMHTMLMenuElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLMenuElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Compact: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLMenuElement) Compact() bool {
	var _arg0 *C.WebKitDOMHTMLMenuElement // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitDOMHTMLMenuElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_menu_element_get_compact(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCompact: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLMenuElement) SetCompact(value bool) {
	var _arg0 *C.WebKitDOMHTMLMenuElement // out
	var _arg1 C.gboolean                  // out

	_arg0 = (*C.WebKitDOMHTMLMenuElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_menu_element_set_compact(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLMenuElementClass: instance of this type is always passed by reference.
type DOMHTMLMenuElementClass struct {
	*domhtmlMenuElementClass
}

// domhtmlMenuElementClass is the struct that's finalized.
type domhtmlMenuElementClass struct {
	native *C.WebKitDOMHTMLMenuElementClass
}

func (d *DOMHTMLMenuElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}