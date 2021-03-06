// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// extern WebKitWebView* _gotk4_webkit24_AutomationSession_ConnectCreateWebView(gpointer, guintptr);
import "C"

// glib.Type values for WebKitAutomationSession.go.
var (
	GTypeAutomationBrowsingContextPresentation = externglib.Type(C.webkit_automation_browsing_context_presentation_get_type())
	GTypeAutomationSession                     = externglib.Type(C.webkit_automation_session_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeAutomationBrowsingContextPresentation, F: marshalAutomationBrowsingContextPresentation},
		{T: GTypeAutomationSession, F: marshalAutomationSession},
	})
}

// AutomationBrowsingContextPresentation: enum values used for determining the
// automation browsing context presentation.
type AutomationBrowsingContextPresentation C.gint

const (
	// AutomationBrowsingContextPresentationWindow: window.
	AutomationBrowsingContextPresentationWindow AutomationBrowsingContextPresentation = iota
	// AutomationBrowsingContextPresentationTab: tab.
	AutomationBrowsingContextPresentationTab
)

func marshalAutomationBrowsingContextPresentation(p uintptr) (interface{}, error) {
	return AutomationBrowsingContextPresentation(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AutomationBrowsingContextPresentation.
func (a AutomationBrowsingContextPresentation) String() string {
	switch a {
	case AutomationBrowsingContextPresentationWindow:
		return "Window"
	case AutomationBrowsingContextPresentationTab:
		return "Tab"
	default:
		return fmt.Sprintf("AutomationBrowsingContextPresentation(%d)", a)
	}
}

// AutomationSessionOverrider contains methods that are overridable.
type AutomationSessionOverrider interface {
}

type AutomationSession struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*AutomationSession)(nil)
)

func classInitAutomationSessioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapAutomationSession(obj *externglib.Object) *AutomationSession {
	return &AutomationSession{
		Object: obj,
	}
}

func marshalAutomationSession(p uintptr) (interface{}, error) {
	return wrapAutomationSession(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_webkit24_AutomationSession_ConnectCreateWebView
func _gotk4_webkit24_AutomationSession_ConnectCreateWebView(arg0 C.gpointer, arg1 C.guintptr) (cret *C.WebKitWebView) {
	var f func() (webView *WebView)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (webView *WebView))
	}

	webView := f()

	cret = (*C.WebKitWebView)(unsafe.Pointer(externglib.InternObject(webView).Native()))

	return cret
}

// ConnectCreateWebView: this signal is emitted when the automation client
// requests a new browsing context to interact with it. The callback handler
// should return a KitWebView created with
// KitWebView:is-controlled-by-automation construct property enabled and
// KitWebView:automation-presentation-type construct property set if needed.
//
// If the signal is emitted with "tab" detail, the returned KitWebView should be
// a new web view added to a new tab of the current browsing context window. If
// the signal is emitted with "window" detail, the returned KitWebView should be
// a new web view added to a new window. When creating a new web view and
// there's an active browsing context, the new window or tab shouldn't be
// focused.
func (session *AutomationSession) ConnectCreateWebView(f func() (webView *WebView)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(session, "create-web-view", false, unsafe.Pointer(C._gotk4_webkit24_AutomationSession_ConnectCreateWebView), f)
}

// ApplicationInfo: get the KitAutomationSession previously set with
// webkit_automation_session_set_application_info().
//
// The function returns the following values:
//
//    - applicationInfo of session, or NULL if no one has been set.
//
func (session *AutomationSession) ApplicationInfo() *ApplicationInfo {
	var _arg0 *C.WebKitAutomationSession // out
	var _cret *C.WebKitApplicationInfo   // in

	_arg0 = (*C.WebKitAutomationSession)(unsafe.Pointer(externglib.InternObject(session).Native()))

	_cret = C.webkit_automation_session_get_application_info(_arg0)
	runtime.KeepAlive(session)

	var _applicationInfo *ApplicationInfo // out

	_applicationInfo = (*ApplicationInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.webkit_application_info_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_applicationInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_application_info_unref((*C.WebKitApplicationInfo)(intern.C))
		},
	)

	return _applicationInfo
}

// ID: get the unique identifier of a KitAutomationSession.
//
// The function returns the following values:
//
//    - utf8: unique identifier of session.
//
func (session *AutomationSession) ID() string {
	var _arg0 *C.WebKitAutomationSession // out
	var _cret *C.char                    // in

	_arg0 = (*C.WebKitAutomationSession)(unsafe.Pointer(externglib.InternObject(session).Native()))

	_cret = C.webkit_automation_session_get_id(_arg0)
	runtime.KeepAlive(session)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetApplicationInfo: set the application information to session. This
// information will be used by the driver service to match the requested
// capabilities with the actual application information. If this information is
// not provided to the session when a new automation session is requested, the
// creation might fail if the client requested a specific browser name or
// version. This will not have any effect when called after the automation
// session has been fully created, so this must be called in the callback of
// KitWebContext::automation-started signal.
//
// The function takes the following parameters:
//
//    - info: KitApplicationInfo.
//
func (session *AutomationSession) SetApplicationInfo(info *ApplicationInfo) {
	var _arg0 *C.WebKitAutomationSession // out
	var _arg1 *C.WebKitApplicationInfo   // out

	_arg0 = (*C.WebKitAutomationSession)(unsafe.Pointer(externglib.InternObject(session).Native()))
	_arg1 = (*C.WebKitApplicationInfo)(gextras.StructNative(unsafe.Pointer(info)))

	C.webkit_automation_session_set_application_info(_arg0, _arg1)
	runtime.KeepAlive(session)
	runtime.KeepAlive(info)
}
