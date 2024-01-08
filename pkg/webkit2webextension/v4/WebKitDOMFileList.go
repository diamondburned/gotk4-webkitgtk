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
	GTypeDOMFileList = coreglib.Type(C.webkit_dom_file_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMFileList, F: marshalDOMFileList},
	})
}

// DOMFileListOverrides contains methods that are overridable.
type DOMFileListOverrides struct {
}

func defaultDOMFileListOverrides(v *DOMFileList) DOMFileListOverrides {
	return DOMFileListOverrides{}
}

type DOMFileList struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMFileList)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMFileList, *DOMFileListClass, DOMFileListOverrides](
		GTypeDOMFileList,
		initDOMFileListClass,
		wrapDOMFileList,
		defaultDOMFileListOverrides,
	)
}

func initDOMFileListClass(gclass unsafe.Pointer, overrides DOMFileListOverrides, classInitFunc func(*DOMFileListClass)) {
	if classInitFunc != nil {
		class := (*DOMFileListClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMFileList(obj *coreglib.Object) *DOMFileList {
	return &DOMFileList{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMFileList(p uintptr) (interface{}, error) {
	return wrapDOMFileList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Length: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMFileList) Length() uint32 {
	var _arg0 *C.WebKitDOMFileList // out
	var _cret C.gulong             // in

	_arg0 = (*C.WebKitDOMFileList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_file_list_get_length(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// Item: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - index: #gulong.
//
// The function returns the following values:
//
//   - domFile: KitDOMFile.
//
func (self *DOMFileList) Item(index uint32) *DOMFile {
	var _arg0 *C.WebKitDOMFileList // out
	var _arg1 C.gulong             // out
	var _cret *C.WebKitDOMFile     // in

	_arg0 = (*C.WebKitDOMFileList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	_cret = C.webkit_dom_file_list_item(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _domFile *DOMFile // out

	_domFile = wrapDOMFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domFile
}

// DOMFileListClass: instance of this type is always passed by reference.
type DOMFileListClass struct {
	*domFileListClass
}

// domFileListClass is the struct that's finalized.
type domFileListClass struct {
	native *C.WebKitDOMFileListClass
}

func (d *DOMFileListClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
