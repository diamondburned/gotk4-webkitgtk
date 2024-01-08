// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

// GType values.
var (
	GTypeDOMHTMLInputElement = coreglib.Type(C.webkit_dom_html_input_element_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMHTMLInputElement, F: marshalDOMHTMLInputElement},
	})
}

// DOMHTMLInputElementOverrides contains methods that are overridable.
type DOMHTMLInputElementOverrides struct {
}

func defaultDOMHTMLInputElementOverrides(v *DOMHTMLInputElement) DOMHTMLInputElementOverrides {
	return DOMHTMLInputElementOverrides{}
}

type DOMHTMLInputElement struct {
	_ [0]func() // equal guard
	DOMHTMLElement
}

var (
	_ coreglib.Objector = (*DOMHTMLInputElement)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMHTMLInputElement, *DOMHTMLInputElementClass, DOMHTMLInputElementOverrides](
		GTypeDOMHTMLInputElement,
		initDOMHTMLInputElementClass,
		wrapDOMHTMLInputElement,
		defaultDOMHTMLInputElementOverrides,
	)
}

func initDOMHTMLInputElementClass(gclass unsafe.Pointer, overrides DOMHTMLInputElementOverrides, classInitFunc func(*DOMHTMLInputElementClass)) {
	if classInitFunc != nil {
		class := (*DOMHTMLInputElementClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMHTMLInputElement(obj *coreglib.Object) *DOMHTMLInputElement {
	return &DOMHTMLInputElement{
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

func marshalDOMHTMLInputElement(p uintptr) (interface{}, error) {
	return wrapDOMHTMLInputElement(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Accept: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) Accept() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_accept(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Align: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) Align() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_align(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Alt: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) Alt() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_alt(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AutoFilled: deprecated: Use
// webkit_dom_element_html_input_element_get_auto_filled() instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) AutoFilled() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_auto_filled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Autofocus: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) Autofocus() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_autofocus(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Capture: deprecated: Use webkit_dom_html_input_element_get_capture_type()
// instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) Capture() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_capture(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CaptureType: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) CaptureType() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_capture_type(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Checked: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) Checked() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_checked(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DefaultChecked: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) DefaultChecked() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_default_checked(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DefaultValue: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) DefaultValue() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_default_value(_arg0)
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
func (self *DOMHTMLInputElement) Disabled() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_disabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Files: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domFileList: KitDOMFileList.
//
func (self *DOMHTMLInputElement) Files() *DOMFileList {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.WebKitDOMFileList         // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_files(_arg0)
	runtime.KeepAlive(self)

	var _domFileList *DOMFileList // out

	_domFileList = wrapDOMFileList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domFileList
}

// Form: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domhtmlFormElement: KitDOMHTMLFormElement.
//
func (self *DOMHTMLInputElement) Form() *DOMHTMLFormElement {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.WebKitDOMHTMLFormElement  // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_form(_arg0)
	runtime.KeepAlive(self)

	var _domhtmlFormElement *DOMHTMLFormElement // out

	_domhtmlFormElement = wrapDOMHTMLFormElement(coreglib.Take(unsafe.Pointer(_cret)))

	return _domhtmlFormElement
}

// Height: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMHTMLInputElement) Height() uint32 {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gulong                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_height(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// Indeterminate: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) Indeterminate() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_indeterminate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InputType: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) InputType() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_input_type(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// MaxLength: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - glong: #glong.
//
func (self *DOMHTMLInputElement) MaxLength() int32 {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.glong                      // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_max_length(_arg0)
	runtime.KeepAlive(self)

	var _glong int32 // out

	_glong = int32(_cret)

	return _glong
}

// Multiple: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) Multiple() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_multiple(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) Name() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ReadOnly: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) ReadOnly() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_read_only(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Size: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMHTMLInputElement) Size() uint32 {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gulong                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_size(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// Src: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) Src() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_src(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// UseMap: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMHTMLInputElement) UseMap() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_use_map(_arg0)
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
func (self *DOMHTMLInputElement) Value() string {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret *C.gchar                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_value(_arg0)
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
//   - gulong: #gulong.
//
func (self *DOMHTMLInputElement) Width() uint32 {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gulong                     // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_width(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// WillValidate: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMHTMLInputElement) WillValidate() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_html_input_element_get_will_validate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEdited: deprecated: Use
// webkit_dom_element_html_input_element_is_user_edited() instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (input *DOMHTMLInputElement) IsEdited() bool {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _cret C.gboolean                   // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(input).Native()))

	_cret = C.webkit_dom_html_input_element_is_edited(_arg0)
	runtime.KeepAlive(input)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Select: deprecated: Use JavaScriptCore API instead.
func (self *DOMHTMLInputElement) Select() {
	var _arg0 *C.WebKitDOMHTMLInputElement // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.webkit_dom_html_input_element_select(_arg0)
	runtime.KeepAlive(self)
}

// SetAccept: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetAccept(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_accept(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetAlign: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetAlign(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_align(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetAlt: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetAlt(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_alt(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetAutoFilled: deprecated: Use
// webkit_dom_element_html_input_element_set_auto_filled() instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetAutoFilled(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_auto_filled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetAutofocus: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetAutofocus(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_autofocus(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetCaptureType: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetCaptureType(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_capture_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetChecked: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetChecked(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_checked(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetDefaultChecked: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetDefaultChecked(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_default_checked(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetDefaultValue: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetDefaultValue(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_default_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetDisabled: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetDisabled(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_disabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetEditingValue: deprecated: Use
// webkit_dom_element_html_input_element_set_editing_value() instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetEditingValue(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_editing_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetFiles: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: KitDOMFileList.
//
func (self *DOMHTMLInputElement) SetFiles(value *DOMFileList) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.WebKitDOMFileList         // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMFileList)(unsafe.Pointer(coreglib.InternObject(value).Native()))

	C.webkit_dom_html_input_element_set_files(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetHeight: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gulong.
//
func (self *DOMHTMLInputElement) SetHeight(value uint32) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gulong                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(value)

	C.webkit_dom_html_input_element_set_height(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetIndeterminate: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetIndeterminate(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_indeterminate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetInputType: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetInputType(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_input_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetMaxLength: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #glong.
//
func (self *DOMHTMLInputElement) SetMaxLength(value int32) error {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.glong                      // out
	var _cerr *C.GError                    // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.glong(value)

	C.webkit_dom_html_input_element_set_max_length(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetMultiple: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetMultiple(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_multiple(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetName: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetName(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetReadOnly: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gboolean.
//
func (self *DOMHTMLInputElement) SetReadOnly(value bool) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gboolean                   // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.webkit_dom_html_input_element_set_read_only(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetSize: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gulong.
//
func (self *DOMHTMLInputElement) SetSize(value uint32) error {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gulong                     // out
	var _cerr *C.GError                    // in

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(value)

	C.webkit_dom_html_input_element_set_size(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetSrc: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetSrc(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_src(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetUseMap: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetUseMap(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_use_map(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetValue: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMHTMLInputElement) SetValue(value string) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 *C.gchar                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_html_input_element_set_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetWidth: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gulong.
//
func (self *DOMHTMLInputElement) SetWidth(value uint32) {
	var _arg0 *C.WebKitDOMHTMLInputElement // out
	var _arg1 C.gulong                     // out

	_arg0 = (*C.WebKitDOMHTMLInputElement)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(value)

	C.webkit_dom_html_input_element_set_width(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// DOMHTMLInputElementClass: instance of this type is always passed by
// reference.
type DOMHTMLInputElementClass struct {
	*domhtmlInputElementClass
}

// domhtmlInputElementClass is the struct that's finalized.
type domhtmlInputElementClass struct {
	native *C.WebKitDOMHTMLInputElementClass
}

func (d *DOMHTMLInputElementClass) ParentClass() *DOMHTMLElementClass {
	valptr := &d.native.parent_class
	var _v *DOMHTMLElementClass // out
	_v = (*DOMHTMLElementClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}