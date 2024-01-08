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
	GTypeDOMDOMTokenList = coreglib.Type(C.webkit_dom_dom_token_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMDOMTokenList, F: marshalDOMDOMTokenList},
	})
}

// DOMDOMTokenListOverrides contains methods that are overridable.
type DOMDOMTokenListOverrides struct {
}

func defaultDOMDOMTokenListOverrides(v *DOMDOMTokenList) DOMDOMTokenListOverrides {
	return DOMDOMTokenListOverrides{}
}

type DOMDOMTokenList struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMDOMTokenList)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMDOMTokenList, *DOMDOMTokenListClass, DOMDOMTokenListOverrides](
		GTypeDOMDOMTokenList,
		initDOMDOMTokenListClass,
		wrapDOMDOMTokenList,
		defaultDOMDOMTokenListOverrides,
	)
}

func initDOMDOMTokenListClass(gclass unsafe.Pointer, overrides DOMDOMTokenListOverrides, classInitFunc func(*DOMDOMTokenListClass)) {
	if classInitFunc != nil {
		class := (*DOMDOMTokenListClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMDOMTokenList(obj *coreglib.Object) *DOMDOMTokenList {
	return &DOMDOMTokenList{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMDOMTokenList(p uintptr) (interface{}, error) {
	return wrapDOMDOMTokenList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Contains: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - token: #gchar.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMDOMTokenList) Contains(token string) bool {
	var _arg0 *C.WebKitDOMDOMTokenList // out
	var _arg1 *C.gchar                 // out
	var _cret C.gboolean               // in

	_arg0 = (*C.WebKitDOMDOMTokenList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(token)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_dom_dom_token_list_contains(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(token)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Length: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMDOMTokenList) Length() uint32 {
	var _arg0 *C.WebKitDOMDOMTokenList // out
	var _cret C.gulong                 // in

	_arg0 = (*C.WebKitDOMDOMTokenList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_dom_token_list_get_length(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// Value: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMDOMTokenList) Value() string {
	var _arg0 *C.WebKitDOMDOMTokenList // out
	var _cret *C.gchar                 // in

	_arg0 = (*C.WebKitDOMDOMTokenList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_dom_token_list_get_value(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Item: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - index: #gulong.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMDOMTokenList) Item(index uint32) string {
	var _arg0 *C.WebKitDOMDOMTokenList // out
	var _arg1 C.gulong                 // out
	var _cret *C.gchar                 // in

	_arg0 = (*C.WebKitDOMDOMTokenList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	_cret = C.webkit_dom_dom_token_list_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Replace: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - token: #gchar.
//   - newToken: #gchar.
//
func (self *DOMDOMTokenList) Replace(token, newToken string) error {
	var _arg0 *C.WebKitDOMDOMTokenList // out
	var _arg1 *C.gchar                 // out
	var _arg2 *C.gchar                 // out
	var _cerr *C.GError                // in

	_arg0 = (*C.WebKitDOMDOMTokenList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(token)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(newToken)))
	defer C.free(unsafe.Pointer(_arg2))

	C.webkit_dom_dom_token_list_replace(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(token)
	runtime.KeepAlive(newToken)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetValue: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMDOMTokenList) SetValue(value string) {
	var _arg0 *C.WebKitDOMDOMTokenList // out
	var _arg1 *C.gchar                 // out

	_arg0 = (*C.WebKitDOMDOMTokenList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_dom_token_list_set_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// Toggle: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - token: #gchar.
//   - force: #gboolean.
//
func (self *DOMDOMTokenList) Toggle(token string, force bool) error {
	var _arg0 *C.WebKitDOMDOMTokenList // out
	var _arg1 *C.gchar                 // out
	var _arg2 C.gboolean               // out
	var _cerr *C.GError                // in

	_arg0 = (*C.WebKitDOMDOMTokenList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(token)))
	defer C.free(unsafe.Pointer(_arg1))
	if force {
		_arg2 = C.TRUE
	}

	C.webkit_dom_dom_token_list_toggle(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(token)
	runtime.KeepAlive(force)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DOMDOMTokenListClass: instance of this type is always passed by reference.
type DOMDOMTokenListClass struct {
	*domdomTokenListClass
}

// domdomTokenListClass is the struct that's finalized.
type domdomTokenListClass struct {
	native *C.WebKitDOMDOMTokenListClass
}

func (d *DOMDOMTokenListClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}