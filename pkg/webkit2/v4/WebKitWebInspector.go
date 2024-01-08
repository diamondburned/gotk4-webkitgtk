// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// extern void _gotk4_webkit24_WebInspector_ConnectClosed(gpointer, guintptr);
// extern gboolean _gotk4_webkit24_WebInspector_ConnectOpenWindow(gpointer, guintptr);
// extern gboolean _gotk4_webkit24_WebInspector_ConnectDetach(gpointer, guintptr);
// extern gboolean _gotk4_webkit24_WebInspector_ConnectBringToFront(gpointer, guintptr);
// extern gboolean _gotk4_webkit24_WebInspector_ConnectAttach(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeWebInspector = coreglib.Type(C.webkit_web_inspector_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWebInspector, F: marshalWebInspector},
	})
}

// WebInspectorOverrides contains methods that are overridable.
type WebInspectorOverrides struct {
}

func defaultWebInspectorOverrides(v *WebInspector) WebInspectorOverrides {
	return WebInspectorOverrides{}
}

// WebInspector access to the WebKit inspector.
//
// The WebKit Inspector is a graphical tool to inspect and change the content of
// a KitWebView. It also includes an interactive JavaScript debugger. Using this
// class one can get a Widget which can be embedded into an application to show
// the inspector.
//
// The inspector is available when the KitSettings of the KitWebView has set
// the KitSettings:enable-developer-extras to true, otherwise no inspector is
// available.
//
//    // Enable the developer extras
//    WebKitSettings *settings = webkit_web_view_get_settings (WEBKIT_WEB_VIEW(my_webview));
//    g_object_set (G_OBJECT(settings), "enable-developer-extras", TRUE, NULL);
//
//    // Load some data or reload to be able to inspect the page
//    webkit_web_view_load_uri (WEBKIT_WEB_VIEW(my_webview), "http://www.gnome.org");
//
//    // Show the inspector
//    WebKitWebInspector *inspector = webkit_web_view_get_inspector (WEBKIT_WEB_VIEW(my_webview));
//    webkit_web_inspector_show (WEBKIT_WEB_INSPECTOR(inspector));.
type WebInspector struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*WebInspector)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WebInspector, *WebInspectorClass, WebInspectorOverrides](
		GTypeWebInspector,
		initWebInspectorClass,
		wrapWebInspector,
		defaultWebInspectorOverrides,
	)
}

func initWebInspectorClass(gclass unsafe.Pointer, overrides WebInspectorOverrides, classInitFunc func(*WebInspectorClass)) {
	if classInitFunc != nil {
		class := (*WebInspectorClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWebInspector(obj *coreglib.Object) *WebInspector {
	return &WebInspector{
		Object: obj,
	}
}

func marshalWebInspector(p uintptr) (interface{}, error) {
	return wrapWebInspector(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAttach is emitted when the inspector is requested to be attached to
// the window where the inspected web view is. If this signal is not handled
// the inspector view will be automatically attached to the inspected view,
// so you only need to handle this signal if you want to attach the inspector
// view yourself (for example, to add the inspector view to a browser tab).
//
// To prevent the inspector view from being attached you can connect to this
// signal and simply return TRUE.
func (inspector *WebInspector) ConnectAttach(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(inspector, "attach", false, unsafe.Pointer(C._gotk4_webkit24_WebInspector_ConnectAttach), f)
}

// ConnectBringToFront is emitted when the inspector should be shown.
//
// If the inspector is not attached the inspector window should be shown on
// top of any other windows. If the inspector is attached the inspector view
// should be made visible. For example, if the inspector view is attached
// using a tab in a browser window, the browser window should be raised and
// the tab containing the inspector view should be the active one. In both
// cases, if this signal is not handled, the default implementation calls
// gtk_window_present() on the current toplevel Window of the inspector view.
func (inspector *WebInspector) ConnectBringToFront(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(inspector, "bring-to-front", false, unsafe.Pointer(C._gotk4_webkit24_WebInspector_ConnectBringToFront), f)
}

// ConnectClosed is emitted when the inspector page is closed. If you are using
// your own inspector window, you should connect to this signal and destroy your
// window.
func (inspector *WebInspector) ConnectClosed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(inspector, "closed", false, unsafe.Pointer(C._gotk4_webkit24_WebInspector_ConnectClosed), f)
}

// ConnectDetach is emitted when the inspector is requested to be detached from
// the window it is currently attached to. The inspector is detached when the
// inspector page is about to be closed, and this signal is emitted right before
// KitWebInspector::closed, or when the user clicks on the detach button in the
// inspector view to show the inspector in a separate window. In this case the
// signal KitWebInspector::open-window is emitted after this one.
//
// To prevent the inspector view from being detached you can connect to this
// signal and simply return TRUE.
func (inspector *WebInspector) ConnectDetach(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(inspector, "detach", false, unsafe.Pointer(C._gotk4_webkit24_WebInspector_ConnectDetach), f)
}

// ConnectOpenWindow is emitted when the inspector is requested to open in a
// separate window. If this signal is not handled, a Window with the inspector
// will be created and shown, so you only need to handle this signal if you want
// to use your own window. This signal is emitted after KitWebInspector::detach
// to show the inspector in a separate window after being detached.
//
// To prevent the inspector from being shown you can connect to this signal and
// simply return TRUE.
func (inspector *WebInspector) ConnectOpenWindow(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(inspector, "open-window", false, unsafe.Pointer(C._gotk4_webkit24_WebInspector_ConnectOpenWindow), f)
}

// Attach: request inspector to be attached.
//
// The signal KitWebInspector::attach will be emitted. If the inspector is
// already attached it does nothing.
func (inspector *WebInspector) Attach() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	C.webkit_web_inspector_attach(_arg0)
	runtime.KeepAlive(inspector)
}

// Close: request inspector to be closed.
func (inspector *WebInspector) Close() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	C.webkit_web_inspector_close(_arg0)
	runtime.KeepAlive(inspector)
}

// Detach: request inspector to be detached.
//
// The signal KitWebInspector::detach will be emitted. If the inspector is
// already detached it does nothing.
func (inspector *WebInspector) Detach() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	C.webkit_web_inspector_detach(_arg0)
	runtime.KeepAlive(inspector)
}

// AttachedHeight: get the height that the inspector view when attached.
//
// Get the height that the inspector view should have when it's attached.
// If the inspector view is not attached this returns 0.
//
// The function returns the following values:
//
//   - guint: height of the inspector view when attached.
//
func (inspector *WebInspector) AttachedHeight() uint {
	var _arg0 *C.WebKitWebInspector // out
	var _cret C.guint               // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	_cret = C.webkit_web_inspector_get_attached_height(_arg0)
	runtime.KeepAlive(inspector)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CanAttach: whether the inspector can be attached to the same window that
// contains the inspected view.
//
// The function returns the following values:
//
//   - ok: TRUE if there is enough room for the inspector view inside the window
//     that contains the inspected view, or FALSE otherwise.
//
func (inspector *WebInspector) CanAttach() bool {
	var _arg0 *C.WebKitWebInspector // out
	var _cret C.gboolean            // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	_cret = C.webkit_web_inspector_get_can_attach(_arg0)
	runtime.KeepAlive(inspector)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InspectedURI: get the URI that is currently being inspected.
//
// This can be NULL if nothing has been loaded yet in the inspected view,
// if the inspector has been closed or when inspected view was loaded from a
// HTML string instead of a URI.
//
// The function returns the following values:
//
//   - utf8: URI that is currently being inspected or NULL.
//
func (inspector *WebInspector) InspectedURI() string {
	var _arg0 *C.WebKitWebInspector // out
	var _cret *C.char               // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	_cret = C.webkit_web_inspector_get_inspected_uri(_arg0)
	runtime.KeepAlive(inspector)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WebView: get the KitWebViewBase used to display the inspector.
//
// This might be NULL if the inspector hasn't been loaded yet, or it has been
// closed.
//
// The function returns the following values:
//
//   - webViewBase used to display the inspector or NULL.
//
func (inspector *WebInspector) WebView() *WebViewBase {
	var _arg0 *C.WebKitWebInspector // out
	var _cret *C.WebKitWebViewBase  // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	_cret = C.webkit_web_inspector_get_web_view(_arg0)
	runtime.KeepAlive(inspector)

	var _webViewBase *WebViewBase // out

	_webViewBase = wrapWebViewBase(coreglib.Take(unsafe.Pointer(_cret)))

	return _webViewBase
}

// IsAttached: whether the inspector view is currently attached to the same
// window that contains the inspected view.
//
// The function returns the following values:
//
//   - ok: TRUE if inspector is currently attached or FALSE otherwise.
//
func (inspector *WebInspector) IsAttached() bool {
	var _arg0 *C.WebKitWebInspector // out
	var _cret C.gboolean            // in

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	_cret = C.webkit_web_inspector_is_attached(_arg0)
	runtime.KeepAlive(inspector)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Show: request inspector to be shown.
func (inspector *WebInspector) Show() {
	var _arg0 *C.WebKitWebInspector // out

	_arg0 = (*C.WebKitWebInspector)(unsafe.Pointer(coreglib.InternObject(inspector).Native()))

	C.webkit_web_inspector_show(_arg0)
	runtime.KeepAlive(inspector)
}

// WebInspectorClass: instance of this type is always passed by reference.
type WebInspectorClass struct {
	*webInspectorClass
}

// webInspectorClass is the struct that's finalized.
type webInspectorClass struct {
	native *C.WebKitWebInspectorClass
}
