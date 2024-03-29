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
	GTypeDOMCSSStyleSheet = coreglib.Type(C.webkit_dom_css_style_sheet_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMCSSStyleSheet, F: marshalDOMCSSStyleSheet},
	})
}

// DOMCSSStyleSheetOverrides contains methods that are overridable.
type DOMCSSStyleSheetOverrides struct {
}

func defaultDOMCSSStyleSheetOverrides(v *DOMCSSStyleSheet) DOMCSSStyleSheetOverrides {
	return DOMCSSStyleSheetOverrides{}
}

type DOMCSSStyleSheet struct {
	_ [0]func() // equal guard
	DOMStyleSheet
}

var (
	_ coreglib.Objector = (*DOMCSSStyleSheet)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMCSSStyleSheet, *DOMCSSStyleSheetClass, DOMCSSStyleSheetOverrides](
		GTypeDOMCSSStyleSheet,
		initDOMCSSStyleSheetClass,
		wrapDOMCSSStyleSheet,
		defaultDOMCSSStyleSheetOverrides,
	)
}

func initDOMCSSStyleSheetClass(gclass unsafe.Pointer, overrides DOMCSSStyleSheetOverrides, classInitFunc func(*DOMCSSStyleSheetClass)) {
	if classInitFunc != nil {
		class := (*DOMCSSStyleSheetClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMCSSStyleSheet(obj *coreglib.Object) *DOMCSSStyleSheet {
	return &DOMCSSStyleSheet{
		DOMStyleSheet: DOMStyleSheet{
			DOMObject: DOMObject{
				Object: obj,
			},
		},
	}
}

func marshalDOMCSSStyleSheet(p uintptr) (interface{}, error) {
	return wrapDOMCSSStyleSheet(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddRule: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - selector: #gchar.
//   - style: #gchar.
//   - index: #gulong.
//
// The function returns the following values:
//
//   - glong: #glong.
//
func (self *DOMCSSStyleSheet) AddRule(selector, style string, index uint32) (int32, error) {
	var _arg0 *C.WebKitDOMCSSStyleSheet // out
	var _arg1 *C.gchar                  // out
	var _arg2 *C.gchar                  // out
	var _arg3 C.gulong                  // out
	var _cret C.glong                   // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.WebKitDOMCSSStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(selector)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(style)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gulong(index)

	_cret = C.webkit_dom_css_style_sheet_add_rule(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(selector)
	runtime.KeepAlive(style)
	runtime.KeepAlive(index)

	var _glong int32 // out
	var _goerr error // out

	_glong = int32(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _glong, _goerr
}

// DeleteRule: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - index: #gulong.
//
func (self *DOMCSSStyleSheet) DeleteRule(index uint32) error {
	var _arg0 *C.WebKitDOMCSSStyleSheet // out
	var _arg1 C.gulong                  // out
	var _cerr *C.GError                 // in

	_arg0 = (*C.WebKitDOMCSSStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	C.webkit_dom_css_style_sheet_delete_rule(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CSSRules: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domcssRuleList: KitDOMCSSRuleList.
//
func (self *DOMCSSStyleSheet) CSSRules() *DOMCSSRuleList {
	var _arg0 *C.WebKitDOMCSSStyleSheet // out
	var _cret *C.WebKitDOMCSSRuleList   // in

	_arg0 = (*C.WebKitDOMCSSStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_css_style_sheet_get_css_rules(_arg0)
	runtime.KeepAlive(self)

	var _domcssRuleList *DOMCSSRuleList // out

	_domcssRuleList = wrapDOMCSSRuleList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domcssRuleList
}

// OwnerRule: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domcssRule: KitDOMCSSRule.
//
func (self *DOMCSSStyleSheet) OwnerRule() *DOMCSSRule {
	var _arg0 *C.WebKitDOMCSSStyleSheet // out
	var _cret *C.WebKitDOMCSSRule       // in

	_arg0 = (*C.WebKitDOMCSSStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_css_style_sheet_get_owner_rule(_arg0)
	runtime.KeepAlive(self)

	var _domcssRule *DOMCSSRule // out

	_domcssRule = wrapDOMCSSRule(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domcssRule
}

// Rules: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domcssRuleList: KitDOMCSSRuleList.
//
func (self *DOMCSSStyleSheet) Rules() *DOMCSSRuleList {
	var _arg0 *C.WebKitDOMCSSStyleSheet // out
	var _cret *C.WebKitDOMCSSRuleList   // in

	_arg0 = (*C.WebKitDOMCSSStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_css_style_sheet_get_rules(_arg0)
	runtime.KeepAlive(self)

	var _domcssRuleList *DOMCSSRuleList // out

	_domcssRuleList = wrapDOMCSSRuleList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domcssRuleList
}

// InsertRule: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - rule: #gchar.
//   - index: #gulong.
//
// The function returns the following values:
//
//   - gulong: #gulong.
//
func (self *DOMCSSStyleSheet) InsertRule(rule string, index uint32) (uint32, error) {
	var _arg0 *C.WebKitDOMCSSStyleSheet // out
	var _arg1 *C.gchar                  // out
	var _arg2 C.gulong                  // out
	var _cret C.gulong                  // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.WebKitDOMCSSStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(rule)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gulong(index)

	_cret = C.webkit_dom_css_style_sheet_insert_rule(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(rule)
	runtime.KeepAlive(index)

	var _gulong uint32 // out
	var _goerr error   // out

	_gulong = uint32(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gulong, _goerr
}

// RemoveRule: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - index: #gulong.
//
func (self *DOMCSSStyleSheet) RemoveRule(index uint32) error {
	var _arg0 *C.WebKitDOMCSSStyleSheet // out
	var _arg1 C.gulong                  // out
	var _cerr *C.GError                 // in

	_arg0 = (*C.WebKitDOMCSSStyleSheet)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gulong(index)

	C.webkit_dom_css_style_sheet_remove_rule(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(index)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DOMCSSStyleSheetClass: instance of this type is always passed by reference.
type DOMCSSStyleSheetClass struct {
	*domcssStyleSheetClass
}

// domcssStyleSheetClass is the struct that's finalized.
type domcssStyleSheetClass struct {
	native *C.WebKitDOMCSSStyleSheetClass
}

func (d *DOMCSSStyleSheetClass) ParentClass() *DOMStyleSheetClass {
	valptr := &d.native.parent_class
	var _v *DOMStyleSheetClass // out
	_v = (*DOMStyleSheetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
