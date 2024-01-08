// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit-web-extension.h>
// gshort _gotk4_webkit2webextension4_DOMNodeFilter_virtual_accept_node(void* fnptr, WebKitDOMNodeFilter* arg0, WebKitDOMNode* arg1) {
//   return ((gshort (*)(WebKitDOMNodeFilter*, WebKitDOMNode*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeDOMNodeFilter = coreglib.Type(C.webkit_dom_node_filter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMNodeFilter, F: marshalDOMNodeFilter},
	})
}

//
// DOMNodeFilter wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DOMNodeFilter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DOMNodeFilter)(nil)
)

// DOMNodeFilterer describes DOMNodeFilter's interface methods.
type DOMNodeFilterer interface {
	coreglib.Objector

	// AcceptNode: deprecated: Use JavaScriptCore API instead.
	AcceptNode(node *DOMNode) int16
}

var _ DOMNodeFilterer = (*DOMNodeFilter)(nil)

func wrapDOMNodeFilter(obj *coreglib.Object) *DOMNodeFilter {
	return &DOMNodeFilter{
		Object: obj,
	}
}

func marshalDOMNodeFilter(p uintptr) (interface{}, error) {
	return wrapDOMNodeFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AcceptNode: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - node: KitDOMNode.
//
// The function returns the following values:
//
//   - gshort: #gshort.
//
func (filter *DOMNodeFilter) AcceptNode(node *DOMNode) int16 {
	var _arg0 *C.WebKitDOMNodeFilter // out
	var _arg1 *C.WebKitDOMNode       // out
	var _cret C.gshort               // in

	_arg0 = (*C.WebKitDOMNodeFilter)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(node).Native()))

	_cret = C.webkit_dom_node_filter_accept_node(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(node)

	var _gshort int16 // out

	_gshort = int16(_cret)

	return _gshort
}

// acceptNode: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - node: KitDOMNode.
//
// The function returns the following values:
//
//   - gshort: #gshort.
//
func (filter *DOMNodeFilter) acceptNode(node *DOMNode) int16 {
	gclass := (*C.WebKitDOMNodeFilterIface)(coreglib.PeekParentClass(filter))
	fnarg := gclass.accept_node

	var _arg0 *C.WebKitDOMNodeFilter // out
	var _arg1 *C.WebKitDOMNode       // out
	var _cret C.gshort               // in

	_arg0 = (*C.WebKitDOMNodeFilter)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(node).Native()))

	_cret = C._gotk4_webkit2webextension4_DOMNodeFilter_virtual_accept_node(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(node)

	var _gshort int16 // out

	_gshort = int16(_cret)

	return _gshort
}

// DOMNodeFilterIface: instance of this type is always passed by reference.
type DOMNodeFilterIface struct {
	*domNodeFilterIface
}

// domNodeFilterIface is the struct that's finalized.
type domNodeFilterIface struct {
	native *C.WebKitDOMNodeFilterIface
}