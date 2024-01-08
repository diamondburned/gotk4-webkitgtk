// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypeWebViewBase = coreglib.Type(C.webkit_web_view_base_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWebViewBase, F: marshalWebViewBase},
	})
}

// WebViewBaseOverrides contains methods that are overridable.
type WebViewBaseOverrides struct {
}

func defaultWebViewBaseOverrides(v *WebViewBase) WebViewBaseOverrides {
	return WebViewBaseOverrides{}
}

type WebViewBase struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*WebViewBase)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WebViewBase, *WebViewBaseClass, WebViewBaseOverrides](
		GTypeWebViewBase,
		initWebViewBaseClass,
		wrapWebViewBase,
		defaultWebViewBaseOverrides,
	)
}

func initWebViewBaseClass(gclass unsafe.Pointer, overrides WebViewBaseOverrides, classInitFunc func(*WebViewBaseClass)) {
	if classInitFunc != nil {
		class := (*WebViewBaseClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWebViewBase(obj *coreglib.Object) *WebViewBase {
	return &WebViewBase{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalWebViewBase(p uintptr) (interface{}, error) {
	return wrapWebViewBase(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WebViewBaseClass: instance of this type is always passed by reference.
type WebViewBaseClass struct {
	*webViewBaseClass
}

// webViewBaseClass is the struct that's finalized.
type webViewBaseClass struct {
	native *C.WebKitWebViewBaseClass
}

func (w *WebViewBaseClass) ParentClass() *gtk.WidgetClass {
	valptr := &w.native.parentClass
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}