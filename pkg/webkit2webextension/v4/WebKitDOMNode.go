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
	GTypeDOMNode = coreglib.Type(C.webkit_dom_node_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMNode, F: marshalDOMNode},
	})
}

// DOM_NODE_ATTRIBUTE_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_ATTRIBUTE_NODE = 2

// DOM_NODE_CDATA_SECTION_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_CDATA_SECTION_NODE = 4

// DOM_NODE_COMMENT_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_COMMENT_NODE = 8

// DOM_NODE_DOCUMENT_FRAGMENT_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_DOCUMENT_FRAGMENT_NODE = 11

// DOM_NODE_DOCUMENT_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_DOCUMENT_NODE = 9

// DOM_NODE_DOCUMENT_POSITION_CONTAINED_BY: deprecated: Use JavaScriptCore API
// instead.
const DOM_NODE_DOCUMENT_POSITION_CONTAINED_BY = 16

// DOM_NODE_DOCUMENT_POSITION_CONTAINS: deprecated: Use JavaScriptCore API
// instead.
const DOM_NODE_DOCUMENT_POSITION_CONTAINS = 8

// DOM_NODE_DOCUMENT_POSITION_DISCONNECTED: deprecated: Use JavaScriptCore API
// instead.
const DOM_NODE_DOCUMENT_POSITION_DISCONNECTED = 1

// DOM_NODE_DOCUMENT_POSITION_FOLLOWING: deprecated: Use JavaScriptCore API
// instead.
const DOM_NODE_DOCUMENT_POSITION_FOLLOWING = 4

// DOM_NODE_DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC: deprecated: Use
// JavaScriptCore API instead.
const DOM_NODE_DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 32

// DOM_NODE_DOCUMENT_POSITION_PRECEDING: deprecated: Use JavaScriptCore API
// instead.
const DOM_NODE_DOCUMENT_POSITION_PRECEDING = 2

// DOM_NODE_DOCUMENT_TYPE_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_DOCUMENT_TYPE_NODE = 10

// DOM_NODE_ELEMENT_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_ELEMENT_NODE = 1

// DOM_NODE_ENTITY_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_ENTITY_NODE = 6

// DOM_NODE_ENTITY_REFERENCE_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_ENTITY_REFERENCE_NODE = 5

// DOM_NODE_PROCESSING_INSTRUCTION_NODE: deprecated: Use JavaScriptCore API
// instead.
const DOM_NODE_PROCESSING_INSTRUCTION_NODE = 7

// DOM_NODE_TEXT_NODE: deprecated: Use JavaScriptCore API instead.
const DOM_NODE_TEXT_NODE = 3

// DOMNodeOverrides contains methods that are overridable.
type DOMNodeOverrides struct {
}

func defaultDOMNodeOverrides(v *DOMNode) DOMNodeOverrides {
	return DOMNodeOverrides{}
}

type DOMNode struct {
	_ [0]func() // equal guard
	DOMObject

	*coreglib.Object
	DOMEventTarget
}

var (
	_ coreglib.Objector = (*DOMNode)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMNode, *DOMNodeClass, DOMNodeOverrides](
		GTypeDOMNode,
		initDOMNodeClass,
		wrapDOMNode,
		defaultDOMNodeOverrides,
	)
}

func initDOMNodeClass(gclass unsafe.Pointer, overrides DOMNodeOverrides, classInitFunc func(*DOMNodeClass)) {
	if classInitFunc != nil {
		class := (*DOMNodeClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMNode(obj *coreglib.Object) *DOMNode {
	return &DOMNode{
		DOMObject: DOMObject{
			Object: obj,
		},
		Object: obj,
		DOMEventTarget: DOMEventTarget{
			Object: obj,
		},
	}
}

func marshalDOMNode(p uintptr) (interface{}, error) {
	return wrapDOMNode(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AppendChild: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - newChild: KitDOMNode.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) AppendChild(newChild *DOMNode) (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(newChild).Native()))

	_cret = C.webkit_dom_node_append_child(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newChild)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// CloneNode: deprecated: Use webkit_dom_node_clone_node_with_error() instead.
//
// The function takes the following parameters:
//
//   - deep: #gboolean.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) CloneNode(deep bool) (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 C.gboolean       // out
	var _cret *C.WebKitDOMNode // in
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if deep {
		_arg1 = C.TRUE
	}

	_cret = C.webkit_dom_node_clone_node(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(deep)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// CloneNodeWithError: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - deep: #gboolean.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) CloneNodeWithError(deep bool) (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 C.gboolean       // out
	var _cret *C.WebKitDOMNode // in
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if deep {
		_arg1 = C.TRUE
	}

	_cret = C.webkit_dom_node_clone_node_with_error(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(deep)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// CompareDocumentPosition: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - other: KitDOMNode.
//
// The function returns the following values:
//
//   - gushort: #gushort.
//
func (self *DOMNode) CompareDocumentPosition(other *DOMNode) uint16 {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _cret C.gushort        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_cret = C.webkit_dom_node_compare_document_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(other)

	var _gushort uint16 // out

	_gushort = uint16(_cret)

	return _gushort
}

// Contains: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - other: KitDOMNode.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMNode) Contains(other *DOMNode) bool {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _cret C.gboolean       // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_cret = C.webkit_dom_node_contains(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BaseURI: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) BaseURI() string {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_base_uri(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ChildNodes: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNodeList: KitDOMNodeList.
//
func (self *DOMNode) ChildNodes() *DOMNodeList {
	var _arg0 *C.WebKitDOMNode     // out
	var _cret *C.WebKitDOMNodeList // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_child_nodes(_arg0)
	runtime.KeepAlive(self)

	var _domNodeList *DOMNodeList // out

	_domNodeList = wrapDOMNodeList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domNodeList
}

// FirstChild: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) FirstChild() *DOMNode {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_first_child(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// LastChild: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) LastChild() *DOMNode {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_last_child(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// LocalName: deprecated: Use webkit_dom_attr_get_local_name() or
// webkit_dom_element_get_local_name() instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) LocalName() string {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_local_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// NamespaceURI: deprecated: Use webkit_dom_attr_get_namespace_uri() or
// webkit_dom_element_get_namespace_uri() instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) NamespaceURI() string {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_namespace_uri(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// NextSibling: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) NextSibling() *DOMNode {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_next_sibling(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// NodeName: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) NodeName() string {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_node_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// NodeType: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gushort: #gushort.
//
func (self *DOMNode) NodeType() uint16 {
	var _arg0 *C.WebKitDOMNode // out
	var _cret C.gushort        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_node_type(_arg0)
	runtime.KeepAlive(self)

	var _gushort uint16 // out

	_gushort = uint16(_cret)

	return _gushort
}

// NodeValue: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) NodeValue() string {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_node_value(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// OwnerDocument: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domDocument: KitDOMDocument.
//
func (self *DOMNode) OwnerDocument() *DOMDocument {
	var _arg0 *C.WebKitDOMNode     // out
	var _cret *C.WebKitDOMDocument // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_owner_document(_arg0)
	runtime.KeepAlive(self)

	var _domDocument *DOMDocument // out

	_domDocument = wrapDOMDocument(coreglib.Take(unsafe.Pointer(_cret)))

	return _domDocument
}

// ParentElement: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domElement: KitDOMElement.
//
func (self *DOMNode) ParentElement() *DOMElement {
	var _arg0 *C.WebKitDOMNode    // out
	var _cret *C.WebKitDOMElement // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_parent_element(_arg0)
	runtime.KeepAlive(self)

	var _domElement *DOMElement // out

	_domElement = wrapDOMElement(coreglib.Take(unsafe.Pointer(_cret)))

	return _domElement
}

// ParentNode: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) ParentNode() *DOMNode {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_parent_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// Prefix: deprecated: Use webkit_dom_attr_get_prefix() or
// webkit_dom_element_get_prefix() instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) Prefix() string {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_prefix(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PreviousSibling: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) PreviousSibling() *DOMNode {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_previous_sibling(_arg0)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))

	return _domNode
}

// TextContent: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) TextContent() string {
	var _arg0 *C.WebKitDOMNode // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_get_text_content(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// HasChildNodes: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMNode) HasChildNodes() bool {
	var _arg0 *C.WebKitDOMNode // out
	var _cret C.gboolean       // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_node_has_child_nodes(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InsertBefore: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - newChild: KitDOMNode.
//   - refChild (optional): KitDOMNode.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) InsertBefore(newChild, refChild *DOMNode) (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _arg2 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(newChild).Native()))
	if refChild != nil {
		_arg2 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(refChild).Native()))
	}

	_cret = C.webkit_dom_node_insert_before(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newChild)
	runtime.KeepAlive(refChild)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// IsDefaultNamespace: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - namespaceURI: #gchar.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMNode) IsDefaultNamespace(namespaceURI string) bool {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(namespaceURI)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_dom_node_is_default_namespace(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(namespaceURI)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEqualNode: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - other: KitDOMNode.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMNode) IsEqualNode(other *DOMNode) bool {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _cret C.gboolean       // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_cret = C.webkit_dom_node_is_equal_node(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSameNode: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - other: KitDOMNode.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMNode) IsSameNode(other *DOMNode) bool {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _cret C.gboolean       // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_cret = C.webkit_dom_node_is_same_node(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSupported: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - feature: #gchar.
//   - version: #gchar.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMNode) IsSupported(feature, version string) bool {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(feature)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(version)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.webkit_dom_node_is_supported(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(version)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupNamespaceURI: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - prefix: #gchar.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) LookupNamespaceURI(prefix string) string {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.gchar         // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(prefix)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_dom_node_lookup_namespace_uri(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(prefix)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// LookupPrefix: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - namespaceURI: #gchar.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMNode) LookupPrefix(namespaceURI string) string {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.gchar         // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(namespaceURI)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_dom_node_lookup_prefix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(namespaceURI)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Normalize: deprecated: Use JavaScriptCore API instead.
func (self *DOMNode) Normalize() {
	var _arg0 *C.WebKitDOMNode // out

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.webkit_dom_node_normalize(_arg0)
	runtime.KeepAlive(self)
}

// RemoveChild: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - oldChild: KitDOMNode.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) RemoveChild(oldChild *DOMNode) (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(oldChild).Native()))

	_cret = C.webkit_dom_node_remove_child(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(oldChild)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// ReplaceChild: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - newChild: KitDOMNode.
//   - oldChild: KitDOMNode.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMNode) ReplaceChild(newChild, oldChild *DOMNode) (*DOMNode, error) {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.WebKitDOMNode // out
	var _arg2 *C.WebKitDOMNode // out
	var _cret *C.WebKitDOMNode // in
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(newChild).Native()))
	_arg2 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(oldChild).Native()))

	_cret = C.webkit_dom_node_replace_child(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newChild)
	runtime.KeepAlive(oldChild)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// SetNodeValue: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMNode) SetNodeValue(value string) error {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.gchar         // out
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_node_set_node_value(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetPrefix: deprecated: since version 2.14.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMNode) SetPrefix(value string) error {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.gchar         // out
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_node_set_prefix(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetTextContent: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMNode) SetTextContent(value string) error {
	var _arg0 *C.WebKitDOMNode // out
	var _arg1 *C.gchar         // out
	var _cerr *C.GError        // in

	_arg0 = (*C.WebKitDOMNode)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_node_set_text_content(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DOMNodeClass: instance of this type is always passed by reference.
type DOMNodeClass struct {
	*domNodeClass
}

// domNodeClass is the struct that's finalized.
type domNodeClass struct {
	native *C.WebKitDOMNodeClass
}

func (d *DOMNodeClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
