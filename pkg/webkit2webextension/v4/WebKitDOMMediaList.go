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
	GTypeDOMMediaList = coreglib.Type(C.webkit_dom_media_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMMediaList, F: marshalDOMMediaList},
	})
}

// DOMMediaListOverrides contains methods that are overridable.
type DOMMediaListOverrides struct {
}

func defaultDOMMediaListOverrides(v *DOMMediaList) DOMMediaListOverrides {
	return DOMMediaListOverrides{}
}

type DOMMediaList struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMMediaList)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMMediaList, *DOMMediaListClass, DOMMediaListOverrides](
		GTypeDOMMediaList,
		initDOMMediaListClass,
		wrapDOMMediaList,
		defaultDOMMediaListOverrides,
	)
}

func initDOMMediaListClass(gclass unsafe.Pointer, overrides DOMMediaListOverrides, classInitFunc func(*DOMMediaListClass)) {
	if classInitFunc != nil {
		class := (*DOMMediaListClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMMediaList(obj *coreglib.Object) *DOMMediaList {
	return &DOMMediaList{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMMediaList(p uintptr) (interface{}, error) {
	return wrapDOMMediaList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AppendMedium: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - newMedium: #gchar.
//
func (self *DOMMediaList) AppendMedium(newMedium string) error {
	var _arg0 *C.WebKitDOMMediaList // out
	var _arg1 *C.gchar              // out
	var _cerr *C.GError             // in

	_arg0 = (*C.WebKitDOMMediaList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(newMedium)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_media_list_append_medium(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newMedium)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DeleteMedium: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - oldMedium: #gchar.
//
func (self *DOMMediaList) DeleteMedium(oldMedium string) error {
	var _arg0 *C.WebKitDOMMediaList // out
	var _arg1 *C.gchar              // out
	var _cerr *C.GError             // in

	_arg0 = (*C.WebKitDOMMediaList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(oldMedium)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_media_list_delete_medium(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(oldMedium)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Length: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMMediaList) Length() uint32 {
	var _arg0 *C.WebKitDOMMediaList // out
	var _cret C.gulong              // in

	_arg0 = (*C.WebKitDOMMediaList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_media_list_get_length(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// MediaText: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMMediaList) MediaText() string {
	var _arg0 *C.WebKitDOMMediaList // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitDOMMediaList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_media_list_get_media_text(_arg0)
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
func (self *DOMMediaList) Item(index uint32) string {
	var _arg0 *C.WebKitDOMMediaList // out
	var _arg1 C.gulong              // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitDOMMediaList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	_cret = C.webkit_dom_media_list_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetMediaText: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMMediaList) SetMediaText(value string) error {
	var _arg0 *C.WebKitDOMMediaList // out
	var _arg1 *C.gchar              // out
	var _cerr *C.GError             // in

	_arg0 = (*C.WebKitDOMMediaList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_media_list_set_media_text(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DOMMediaListClass: instance of this type is always passed by reference.
type DOMMediaListClass struct {
	*domMediaListClass
}

// domMediaListClass is the struct that's finalized.
type domMediaListClass struct {
	native *C.WebKitDOMMediaListClass
}

func (d *DOMMediaListClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
