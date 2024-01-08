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
	GTypeDOMCSSRule = coreglib.Type(C.webkit_dom_css_rule_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDOMCSSRule, F: marshalDOMCSSRule},
	})
}

// DOM_CSS_RULE_CHARSET_RULE: deprecated: Use JavaScriptCore API instead.
const DOM_CSS_RULE_CHARSET_RULE = 2

// DOM_CSS_RULE_FONT_FACE_RULE: deprecated: Use JavaScriptCore API instead.
const DOM_CSS_RULE_FONT_FACE_RULE = 5

// DOM_CSS_RULE_IMPORT_RULE: deprecated: Use JavaScriptCore API instead.
const DOM_CSS_RULE_IMPORT_RULE = 3

// DOM_CSS_RULE_MEDIA_RULE: deprecated: Use JavaScriptCore API instead.
const DOM_CSS_RULE_MEDIA_RULE = 4

// DOM_CSS_RULE_PAGE_RULE: deprecated: Use JavaScriptCore API instead.
const DOM_CSS_RULE_PAGE_RULE = 6

// DOM_CSS_RULE_STYLE_RULE: deprecated: Use JavaScriptCore API instead.
const DOM_CSS_RULE_STYLE_RULE = 1

// DOM_CSS_RULE_UNKNOWN_RULE: deprecated: Use JavaScriptCore API instead.
const DOM_CSS_RULE_UNKNOWN_RULE = 0

// DOMCSSRuleOverrides contains methods that are overridable.
type DOMCSSRuleOverrides struct {
}

func defaultDOMCSSRuleOverrides(v *DOMCSSRule) DOMCSSRuleOverrides {
	return DOMCSSRuleOverrides{}
}

type DOMCSSRule struct {
	_ [0]func() // equal guard
	DOMObject
}

var (
	_ coreglib.Objector = (*DOMCSSRule)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DOMCSSRule, *DOMCSSRuleClass, DOMCSSRuleOverrides](
		GTypeDOMCSSRule,
		initDOMCSSRuleClass,
		wrapDOMCSSRule,
		defaultDOMCSSRuleOverrides,
	)
}

func initDOMCSSRuleClass(gclass unsafe.Pointer, overrides DOMCSSRuleOverrides, classInitFunc func(*DOMCSSRuleClass)) {
	if classInitFunc != nil {
		class := (*DOMCSSRuleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDOMCSSRule(obj *coreglib.Object) *DOMCSSRule {
	return &DOMCSSRule{
		DOMObject: DOMObject{
			Object: obj,
		},
	}
}

func marshalDOMCSSRule(p uintptr) (interface{}, error) {
	return wrapDOMCSSRule(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CSSText: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - utf8: #gchar.
//
func (self *DOMCSSRule) CSSText() string {
	var _arg0 *C.WebKitDOMCSSRule // out
	var _cret *C.gchar            // in

	_arg0 = (*C.WebKitDOMCSSRule)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_css_rule_get_css_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ParentRule: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domcssRule: KitDOMCSSRule.
//
func (self *DOMCSSRule) ParentRule() *DOMCSSRule {
	var _arg0 *C.WebKitDOMCSSRule // out
	var _cret *C.WebKitDOMCSSRule // in

	_arg0 = (*C.WebKitDOMCSSRule)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_css_rule_get_parent_rule(_arg0)
	runtime.KeepAlive(self)

	var _domcssRule *DOMCSSRule // out

	_domcssRule = wrapDOMCSSRule(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domcssRule
}

// ParentStyleSheet: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domcssStyleSheet: KitDOMCSSStyleSheet.
//
func (self *DOMCSSRule) ParentStyleSheet() *DOMCSSStyleSheet {
	var _arg0 *C.WebKitDOMCSSRule       // out
	var _cret *C.WebKitDOMCSSStyleSheet // in

	_arg0 = (*C.WebKitDOMCSSRule)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_css_rule_get_parent_style_sheet(_arg0)
	runtime.KeepAlive(self)

	var _domcssStyleSheet *DOMCSSStyleSheet // out

	_domcssStyleSheet = wrapDOMCSSStyleSheet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domcssStyleSheet
}

// RuleType: deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - gushort: #gushort.
//
func (self *DOMCSSRule) RuleType() uint16 {
	var _arg0 *C.WebKitDOMCSSRule // out
	var _cret C.gushort           // in

	_arg0 = (*C.WebKitDOMCSSRule)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.webkit_dom_css_rule_get_rule_type(_arg0)
	runtime.KeepAlive(self)

	var _gushort uint16 // out

	_gushort = uint16(_cret)

	return _gushort
}

// SetCSSText: deprecated: Use JavaScriptCore API instead.
//
// The function takes the following parameters:
//
//   - value: #gchar.
//
func (self *DOMCSSRule) SetCSSText(value string) error {
	var _arg0 *C.WebKitDOMCSSRule // out
	var _arg1 *C.gchar            // out
	var _cerr *C.GError           // in

	_arg0 = (*C.WebKitDOMCSSRule)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_dom_css_rule_set_css_text(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DOMCSSRuleClass: instance of this type is always passed by reference.
type DOMCSSRuleClass struct {
	*domcssRuleClass
}

// domcssRuleClass is the struct that's finalized.
type domcssRuleClass struct {
	native *C.WebKitDOMCSSRuleClass
}

func (d *DOMCSSRuleClass) ParentClass() *DOMObjectClass {
	valptr := &d.native.parent_class
	var _v *DOMObjectClass // out
	_v = (*DOMObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}