// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

// #include <stdlib.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

// DOM_NODE_FILTER_ACCEPT: accept the node. Use this macro as return value
// of webkit_dom_node_filter_accept_node() implementation to accept the given
// KitDOMNode
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_ACCEPT = 1

// DOM_NODE_FILTER_REJECT: reject the node. Use this macro as return value
// of webkit_dom_node_filter_accept_node() implementation to reject the given
// KitDOMNode. The children of the given node will be rejected too.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_REJECT = 2

// DOM_NODE_FILTER_SHOW_ALL: show all nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_ALL = 4294967295

// DOM_NODE_FILTER_SHOW_ATTRIBUTE: show KitDOMAttr nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_ATTRIBUTE = 2

// DOM_NODE_FILTER_SHOW_CDATA_SECTION: show KitDOMCDataSection nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_CDATA_SECTION = 8

// DOM_NODE_FILTER_SHOW_COMMENT: show KitDOMComment nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_COMMENT = 128

// DOM_NODE_FILTER_SHOW_DOCUMENT: show KitDOMDocument nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_DOCUMENT = 256

// DOM_NODE_FILTER_SHOW_DOCUMENT_FRAGMENT: show KitDOMDocumentFragment nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_DOCUMENT_FRAGMENT = 1024

// DOM_NODE_FILTER_SHOW_DOCUMENT_TYPE: show KitDOMDocumentType nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_DOCUMENT_TYPE = 512

// DOM_NODE_FILTER_SHOW_ELEMENT: show KitDOMElement nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_ELEMENT = 1

// DOM_NODE_FILTER_SHOW_ENTITY: show KitDOMEntity nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_ENTITY = 32

// DOM_NODE_FILTER_SHOW_ENTITY_REFERENCE: show KitDOMEntityReference nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_ENTITY_REFERENCE = 16

// DOM_NODE_FILTER_SHOW_NOTATION: show KitDOMNotation nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_NOTATION = 2048

// DOM_NODE_FILTER_SHOW_PROCESSING_INSTRUCTION: show KitDOMProcessingInstruction
// nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_PROCESSING_INSTRUCTION = 64

// DOM_NODE_FILTER_SHOW_TEXT: show KitDOMText nodes.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SHOW_TEXT = 4

// DOM_NODE_FILTER_SKIP: skip the node. Use this macro as return value of
// webkit_dom_node_filter_accept_node() implementation to skip the given
// KitDOMNode. The children of the given node will not be skipped.
//
// Deprecated: Use JavaScriptCore API instead.
const DOM_NODE_FILTER_SKIP = 3