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
	GTypeDOMDocumentFragment = coreglib.Type(C.webkit_dom_document_fragment_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMDocumentFragment, F: marshalDOMDocumentFragment},
	})
}

// DOMDocumentFragmentOverrides contains methods that are overridable.
type DOMDocumentFragmentOverrides struct {
}

func defaultDOMDocumentFragmentOverrides(v *DOMDocumentFragment) DOMDocumentFragmentOverrides {
	return DOMDocumentFragmentOverrides{}
}

type DOMDocumentFragment struct {
	_ [0]func() // equal guard
	DOMNode
}

var (
	_ coreglib.Objector = (*DOMDocumentFragment)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMDocumentFragment, *DOMDocumentFragmentClass, DOMDocumentFragmentOverrides](
		GTypeDOMDocumentFragment,
		initDOMDocumentFragmentClass,
		wrapDOMDocumentFragment,
		defaultDOMDocumentFragmentOverrides,
	)
}

func initDOMDocumentFragmentClass(gclass unsafe.Pointer, overrides DOMDocumentFragmentOverrides, classInitFunc func(*DOMDocumentFragmentClass)) {
	if classInitFunc != nil {
		class := (*DOMDocumentFragmentClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMDocumentFragment(obj *coreglib.Object) *DOMDocumentFragment {
	return &DOMDocumentFragment{
		DOMNode: DOMNode{
			DOMObject: DOMObject{
				Object: obj,
			},
			Object: obj,
			DOMEventTarget: DOMEventTarget{
				Object: obj,
			},
		},
	}
}

func marshalDOMDocumentFragment(p uintptr) (interface{}, error) {
	return wrapDOMDocumentFragment(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ChildElementCount: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMDocumentFragment) ChildElementCount() uint32 {
	var _arg0 *C.WebKitDOMDocumentFragment // out
	var _cret C.gulong                     // in

	_arg0 = (*C.WebKitDOMDocumentFragment)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_document_fragment_get_child_element_count(_arg0)
	runtime.KeepAlive(self)

	var _gulong uint32 // out

	_gulong = uint32(_cret)

	return _gulong
}

// Children: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domhtmlCollection: KitDOMHTMLCollection.
//
func (self *DOMDocumentFragment) Children() *DOMHTMLCollection {
	var _arg0 *C.WebKitDOMDocumentFragment // out
	var _cret *C.WebKitDOMHTMLCollection   // in

	_arg0 = (*C.WebKitDOMDocumentFragment)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_document_fragment_get_children(_arg0)
	runtime.KeepAlive(self)

	var _domhtmlCollection *DOMHTMLCollection // out

	_domhtmlCollection = wrapDOMHTMLCollection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domhtmlCollection
}

// ElementByID: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - elementId: #gchar.
//
// The function returns the following values:
//
//   - domElement: KitDOMElement.
//
func (self *DOMDocumentFragment) ElementByID(elementId string) *DOMElement {
	var _arg0 *C.WebKitDOMDocumentFragment // out
	var _arg1 *C.gchar                     // out
	var _cret *C.WebKitDOMElement          // in

	_arg0 = (*C.WebKitDOMDocumentFragment)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(elementId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_dom_document_fragment_get_element_by_id(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(elementId)

	var _domElement *DOMElement // out

	_domElement = wrapDOMElement(coreglib.Take(unsafe.Pointer(_cret)))

	return _domElement
}

// FirstElementChild: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domElement: KitDOMElement.
//
func (self *DOMDocumentFragment) FirstElementChild() *DOMElement {
	var _arg0 *C.WebKitDOMDocumentFragment // out
	var _cret *C.WebKitDOMElement          // in

	_arg0 = (*C.WebKitDOMDocumentFragment)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_document_fragment_get_first_element_child(_arg0)
	runtime.KeepAlive(self)

	var _domElement *DOMElement // out

	_domElement = wrapDOMElement(coreglib.Take(unsafe.Pointer(_cret)))

	return _domElement
}

// LastElementChild: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domElement: KitDOMElement.
//
func (self *DOMDocumentFragment) LastElementChild() *DOMElement {
	var _arg0 *C.WebKitDOMDocumentFragment // out
	var _cret *C.WebKitDOMElement          // in

	_arg0 = (*C.WebKitDOMDocumentFragment)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_document_fragment_get_last_element_child(_arg0)
	runtime.KeepAlive(self)

	var _domElement *DOMElement // out

	_domElement = wrapDOMElement(coreglib.Take(unsafe.Pointer(_cret)))

	return _domElement
}

// QuerySelector: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - selectors: #gchar.
//
// The function returns the following values:
//
//   - domElement: KitDOMElement.
//
func (self *DOMDocumentFragment) QuerySelector(selectors string) (*DOMElement, error) {
	var _arg0 *C.WebKitDOMDocumentFragment // out
	var _arg1 *C.gchar                     // out
	var _cret *C.WebKitDOMElement          // in
	var _cerr *C.GError                    // in

	_arg0 = (*C.WebKitDOMDocumentFragment)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(selectors)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_dom_document_fragment_query_selector(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(selectors)

	var _domElement *DOMElement // out
	var _goerr error            // out

	_domElement = wrapDOMElement(coreglib.Take(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domElement, _goerr
}

// QuerySelectorAll: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - selectors: #gchar.
//
// The function returns the following values:
//
//   - domNodeList: KitDOMNodeList.
//
func (self *DOMDocumentFragment) QuerySelectorAll(selectors string) (*DOMNodeList, error) {
	var _arg0 *C.WebKitDOMDocumentFragment // out
	var _arg1 *C.gchar                     // out
	var _cret *C.WebKitDOMNodeList         // in
	var _cerr *C.GError                    // in

	_arg0 = (*C.WebKitDOMDocumentFragment)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(selectors)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_dom_document_fragment_query_selector_all(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(selectors)

	var _domNodeList *DOMNodeList // out
	var _goerr error              // out

	_domNodeList = wrapDOMNodeList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domNodeList, _goerr
}

// DOMDocumentFragmentClass: instance of this type is always passed by
// reference.
type DOMDocumentFragmentClass struct {
	*domDocumentFragmentClass
}

// domDocumentFragmentClass is the struct that's finalized.
type domDocumentFragmentClass struct {
	native *C.WebKitDOMDocumentFragmentClass
}

func (d *DOMDocumentFragmentClass) ParentClass() *DOMNodeClass {
	valptr := &d.native.parent_class
	var _v *DOMNodeClass // out
	_v = (*DOMNodeClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
