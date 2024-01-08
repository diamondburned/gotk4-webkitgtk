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
	GTypeDOMNodeIterator = coreglib.Type(C.webkit_dom_node_iterator_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMNodeIterator, F: marshalDOMNodeIterator},
	})
}

// DOMNodeIteratorOverrides contains methods that are overridable.
type DOMNodeIteratorOverrides struct {
}

func defaultDOMNodeIteratorOverrides(v *DOMNodeIterator) DOMNodeIteratorOverrides {
	return DOMNodeIteratorOverrides{}
}

type DOMNodeIterator struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMNodeIterator)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMNodeIterator, *DOMNodeIteratorClass, DOMNodeIteratorOverrides](
		GTypeDOMNodeIterator,
		initDOMNodeIteratorClass,
		wrapDOMNodeIterator,
		defaultDOMNodeIteratorOverrides,
	)
}

func initDOMNodeIteratorClass(gclass unsafe.Pointer, overrides DOMNodeIteratorOverrides, classInitFunc func(*DOMNodeIteratorClass)) {
	if classInitFunc != nil {
		class := (*DOMNodeIteratorClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMNodeIterator(obj *coreglib.Object) *DOMNodeIterator {
	return &DOMNodeIterator{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMNodeIterator(p uintptr) (interface{}, error) {
	return wrapDOMNodeIterator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Detach: deprecated: Use JavaScriptCore API instead.
func (self *DOMNodeIterator) Detach() {
	var _arg0 *C.WebKitDOMNodeIterator // out

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.webkit_dom_node_iterator_detach(_arg0)
	runtime.KeepAlive(self)
}

// ExpandEntityReferences: this function has been removed from the DOM spec and
// it just returns FALSE.
//
// Deprecated: since version 2.12.
//
// The function returns the following values:
//
//   - ok: #gboolean *.
//
func (self *DOMNodeIterator) ExpandEntityReferences() bool {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret C.gboolean               // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_get_expand_entity_references(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Filter: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNodeFilter: KitDOMNodeFilter.
//
func (self *DOMNodeIterator) Filter() *DOMNodeFilter {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret *C.WebKitDOMNodeFilter   // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_get_filter(_arg0)
	runtime.KeepAlive(self)

	var _domNodeFilter *DOMNodeFilter // out

	_domNodeFilter = wrapDOMNodeFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domNodeFilter
}

// PointerBeforeReferenceNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMNodeIterator) PointerBeforeReferenceNode() bool {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret C.gboolean               // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_get_pointer_before_reference_node(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReferenceNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNodeIterator) ReferenceNode() *DOMNode {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret *C.WebKitDOMNode         // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_get_reference_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// Root: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNodeIterator) Root() *DOMNode {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret *C.WebKitDOMNode         // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_get_root(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// WhatToShow: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMNodeIterator) WhatToShow() uint32 {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret C.gulong                 // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_get_what_to_show(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// NextNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNodeIterator) NextNode() (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret *C.WebKitDOMNode         // in
	var _cerr *C.GError                // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_next_node(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// PreviousNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNodeIterator) PreviousNode() (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNodeIterator // out
	var _cret *C.WebKitDOMNode         // in
	var _cerr *C.GError                // in

	_arg0 = (*C.WebKitDOMNodeIterator)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_iterator_previous_node(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// DOMNodeIteratorClass: instance of this type is always passed by reference.
type DOMNodeIteratorClass struct {
	*domNodeIteratorClass
}

// domNodeIteratorClass is the struct that's finalized.
type domNodeIteratorClass struct {
	native *C.WebKitDOMNodeIteratorClass
}

func (d *DOMNodeIteratorClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
