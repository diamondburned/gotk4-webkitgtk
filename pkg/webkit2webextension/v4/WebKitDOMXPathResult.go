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
	GTypeDOMXPathResult = coreglib.Type(C.webkit_dom_xpath_result_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMXPathResult, F: marshalDOMXPathResult},
	})
}

// DOM_XPATH_RESULT_ANY_TYPE: deprecated: Use JavaScriptCore API instead.
const DOM_XPATH_RESULT_ANY_TYPE = 0

// DOM_XPATH_RESULT_ANY_UNORDERED_NODE_TYPE: deprecated: Use JavaScriptCore API
// instead.
const DOM_XPATH_RESULT_ANY_UNORDERED_NODE_TYPE = 8

// DOM_XPATH_RESULT_BOOLEAN_TYPE: deprecated: Use JavaScriptCore API instead.
const DOM_XPATH_RESULT_BOOLEAN_TYPE = 3

// DOM_XPATH_RESULT_FIRST_ORDERED_NODE_TYPE: deprecated: Use JavaScriptCore API
// instead.
const DOM_XPATH_RESULT_FIRST_ORDERED_NODE_TYPE = 9

// DOM_XPATH_RESULT_NUMBER_TYPE: deprecated: Use JavaScriptCore API instead.
const DOM_XPATH_RESULT_NUMBER_TYPE = 1

// DOM_XPATH_RESULT_ORDERED_NODE_ITERATOR_TYPE: deprecated: Use JavaScriptCore
// API instead.
const DOM_XPATH_RESULT_ORDERED_NODE_ITERATOR_TYPE = 5

// DOM_XPATH_RESULT_ORDERED_NODE_SNAPSHOT_TYPE: deprecated: Use JavaScriptCore
// API instead.
const DOM_XPATH_RESULT_ORDERED_NODE_SNAPSHOT_TYPE = 7

// DOM_XPATH_RESULT_STRING_TYPE: deprecated: Use JavaScriptCore API instead.
const DOM_XPATH_RESULT_STRING_TYPE = 2

// DOM_XPATH_RESULT_UNORDERED_NODE_ITERATOR_TYPE: deprecated: Use JavaScriptCore
// API instead.
const DOM_XPATH_RESULT_UNORDERED_NODE_ITERATOR_TYPE = 4

// DOM_XPATH_RESULT_UNORDERED_NODE_SNAPSHOT_TYPE: deprecated: Use JavaScriptCore
// API instead.
const DOM_XPATH_RESULT_UNORDERED_NODE_SNAPSHOT_TYPE = 6

// DOMXPathResultOverrides contains methods that are overridable.
type DOMXPathResultOverrides struct {
}

func defaultDOMXPathResultOverrides(v *DOMXPathResult) DOMXPathResultOverrides {
	return DOMXPathResultOverrides{}
}

type DOMXPathResult struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMXPathResult)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMXPathResult, *DOMXPathResultClass, DOMXPathResultOverrides](
		GTypeDOMXPathResult,
		initDOMXPathResultClass,
		wrapDOMXPathResult,
		defaultDOMXPathResultOverrides,
	)
}

func initDOMXPathResultClass(gclass unsafe.Pointer, overrides DOMXPathResultOverrides, classInitFunc func(*DOMXPathResultClass)) {
	if classInitFunc != nil {
		class := (*DOMXPathResultClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMXPathResult(obj *coreglib.Object) *DOMXPathResult {
	return &DOMXPathResult{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMXPathResult(p uintptr) (interface{}, error) {
	return wrapDOMXPathResult(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BooleanValue: deprecated: Use JavaScriptCore API instead.
func (self *DOMXPathResult) BooleanValue() error {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cerr *C.GError               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.webkit_dom_xpath_result_get_boolean_value(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// InvalidIteratorState: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - ok: #gboolean.
//
func (self *DOMXPathResult) InvalidIteratorState() bool {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cret C.gboolean              // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_xpath_result_get_invalid_iterator_state(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NumberValue: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gdouble: #gdouble.
//
func (self *DOMXPathResult) NumberValue() (float64, error) {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cret C.gdouble               // in
	var _cerr *C.GError               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_xpath_result_get_number_value(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _gdouble float64 // out
	var _goerr error     // out

	_gdouble = float64(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gdouble, _goerr
}

// ResultType: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gushort: #gushort.
//
func (self *DOMXPathResult) ResultType() uint16 {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cret C.gushort               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_xpath_result_get_result_type(_arg0)
	runtime.KeepAlive(self)

	var _gushort uint16 // out

	_gushort = uint16(_cret)

	return _gushort
}

// SingleNodeValue: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMXPathResult) SingleNodeValue() (*DOMNode, error) {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cret *C.WebKitDOMNode        // in
	var _cerr *C.GError               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_xpath_result_get_single_node_value(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// SnapshotLength: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMXPathResult) SnapshotLength() (uint32, error) {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cret C.gulong                // in
	var _cerr *C.GError               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_xpath_result_get_snapshot_length(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _gulong uint32 // out
	var _goerr error   // out

	_gulong = uint32(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gulong, _goerr
}

// StringValue: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMXPathResult) StringValue() (string, error) {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cret *C.gchar                // in
	var _cerr *C.GError               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_xpath_result_get_string_value(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}

// IterateNext: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMXPathResult) IterateNext() (*DOMNode, error) {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _cret *C.WebKitDOMNode        // in
	var _cerr *C.GError               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_xpath_result_iterate_next(_arg0, &_cerr)
	runtime.KeepAlive(self)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// SnapshotItem: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - index: #gulong.
//
// The function returns the following values:
//
//   - domNode: KitDOMNode.
//
func (self *DOMXPathResult) SnapshotItem(index uint32) (*DOMNode, error) {
	var _arg0 *C.WebKitDOMXPathResult // out
	var _arg1 C.gulong                // out
	var _cret *C.WebKitDOMNode        // in
	var _cerr *C.GError               // in

	_arg0 = (*C.WebKitDOMXPathResult)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	_cret = C.webkit_dom_xpath_result_snapshot_item(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _domNode *DOMNode // out
	var _goerr error      // out

	_domNode = wrapDOMNode(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNode, _goerr
}

// DOMXPathResultClass: instance of this type is always passed by reference.
type DOMXPathResultClass struct {
	*domxPathResultClass
}

// domxPathResultClass is the struct that's finalized.
type domxPathResultClass struct {
	native *C.WebKitDOMXPathResultClass
}

func (d *DOMXPathResultClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}