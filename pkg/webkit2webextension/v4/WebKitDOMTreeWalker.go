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
	GTypeDOMTreeWalker = coreglib.Type(C.webkit_dom_tree_walker_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMTreeWalker, F: marshalDOMTreeWalker},
	})
}

// DOMTreeWalkerOverrides contains methods that are overridable.
type DOMTreeWalkerOverrides struct {
}

func defaultDOMTreeWalkerOverrides(v *DOMTreeWalker) DOMTreeWalkerOverrides {
	return DOMTreeWalkerOverrides{}
}

type DOMTreeWalker struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMTreeWalker)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMTreeWalker, *DOMTreeWalkerClass, DOMTreeWalkerOverrides](
		GTypeDOMTreeWalker,
		initDOMTreeWalkerClass,
		wrapDOMTreeWalker,
		defaultDOMTreeWalkerOverrides,
	)
}

func initDOMTreeWalkerClass(gclass unsafe.Pointer, overrides DOMTreeWalkerOverrides, classInitFunc func(*DOMTreeWalkerClass)) {
	if classInitFunc != nil {
		class := (*DOMTreeWalkerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMTreeWalker(obj *coreglib.Object) *DOMTreeWalker {
	return &DOMTreeWalker{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMTreeWalker(p uintptr) (interface{}, error) {
	return wrapDOMTreeWalker(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// FirstChild: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) FirstChild() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_first_child(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// CurrentNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) CurrentNode() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_get_current_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// ExpandEntityReferences: this function has been removed from the DOM spec and
// it just returns FALSE.
//
// Deprecated: since version 2.12.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMTreeWalker) ExpandEntityReferences() bool {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_get_expand_entity_references(_arg0)
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
func (self *DOMTreeWalker) Filter() *DOMNodeFilter {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNodeFilter // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_get_filter(_arg0)
	runtime.KeepAlive(self)

	var _domNodeFilter *DOMNodeFilter // out

	_domNodeFilter = wrapDOMNodeFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domNodeFilter
}

// Root: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) Root() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_get_root(_arg0)
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
func (self *DOMTreeWalker) WhatToShow() uint32 {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret C.gulong               // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_get_what_to_show(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// LastChild: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) LastChild() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_last_child(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// NextNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) NextNode() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_next_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// NextSibling: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) NextSibling() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_next_sibling(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// ParentNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) ParentNode() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_parent_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// PreviousNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) PreviousNode() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_previous_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// PreviousSibling: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMTreeWalker) PreviousSibling() *DOMNode {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _cret *C.WebKitDOMNode       // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_tree_walker_previous_sibling(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// SetCurrentNode: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: KitDOMNode.
//
func (self *DOMTreeWalker) SetCurrentNode(value *DOMNode) error {
	var _arg0 *C.WebKitDOMTreeWalker // out
	var _arg1 *C.WebKitDOMNode       // out
	var _cerr *C.GError              // in

	_arg0 = (*C.WebKitDOMTreeWalker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(value).Native()))

	C.webkit_dom_tree_walker_set_current_node(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DOMTreeWalkerClass: instance of this type is always passed by reference.
type DOMTreeWalkerClass struct {
	*domTreeWalkerClass
}

// domTreeWalkerClass is the struct that's finalized.
type domTreeWalkerClass struct {
	native *C.WebKitDOMTreeWalkerClass
}

func (d *DOMTreeWalkerClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
