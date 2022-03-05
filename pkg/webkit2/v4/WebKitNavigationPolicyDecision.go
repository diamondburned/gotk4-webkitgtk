// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// glib.Type values for WebKitNavigationPolicyDecision.go.
var GTypeNavigationPolicyDecision = externglib.Type(C.webkit_navigation_policy_decision_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeNavigationPolicyDecision, F: marshalNavigationPolicyDecision},
	})
}

// NavigationPolicyDecisionOverrider contains methods that are overridable.
type NavigationPolicyDecisionOverrider interface {
}

type NavigationPolicyDecision struct {
	_ [0]func() // equal guard
	PolicyDecision
}

var (
	_ PolicyDecisioner = (*NavigationPolicyDecision)(nil)
)

func classInitNavigationPolicyDecisioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNavigationPolicyDecision(obj *externglib.Object) *NavigationPolicyDecision {
	return &NavigationPolicyDecision{
		PolicyDecision: PolicyDecision{
			Object: obj,
		},
	}
}

func marshalNavigationPolicyDecision(p uintptr) (interface{}, error) {
	return wrapNavigationPolicyDecision(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// FrameName gets the value of the KitNavigationPolicyDecision:frame-name
// property.
//
// The function returns the following values:
//
//    - utf8: name of the new frame this navigation action targets or NULL.
//
func (decision *NavigationPolicyDecision) FrameName() string {
	var _arg0 *C.WebKitNavigationPolicyDecision // out
	var _cret *C.gchar                          // in

	_arg0 = (*C.WebKitNavigationPolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_navigation_policy_decision_get_frame_name(_arg0)
	runtime.KeepAlive(decision)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Modifiers gets the value of the KitNavigationPolicyDecision:modifiers
// property.
//
// Deprecated: Use webkit_navigation_policy_decision_get_navigation_action()
// instead.
//
// The function returns the following values:
//
//    - guint modifiers active if this decision was triggered by a mouse event.
//
func (decision *NavigationPolicyDecision) Modifiers() uint {
	var _arg0 *C.WebKitNavigationPolicyDecision // out
	var _cret C.guint                           // in

	_arg0 = (*C.WebKitNavigationPolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_navigation_policy_decision_get_modifiers(_arg0)
	runtime.KeepAlive(decision)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MouseButton gets the value of the KitNavigationPolicyDecision:mouse-button
// property.
//
// Deprecated: Use webkit_navigation_policy_decision_get_navigation_action()
// instead.
//
// The function returns the following values:
//
//    - guint: mouse button used if this decision was triggered by a mouse event
//      or 0 otherwise.
//
func (decision *NavigationPolicyDecision) MouseButton() uint {
	var _arg0 *C.WebKitNavigationPolicyDecision // out
	var _cret C.guint                           // in

	_arg0 = (*C.WebKitNavigationPolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_navigation_policy_decision_get_mouse_button(_arg0)
	runtime.KeepAlive(decision)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NavigationAction gets the value of the
// KitNavigationPolicyDecision:navigation-action property.
//
// The function returns the following values:
//
//    - navigationAction triggering this policy decision.
//
func (decision *NavigationPolicyDecision) NavigationAction() *NavigationAction {
	var _arg0 *C.WebKitNavigationPolicyDecision // out
	var _cret *C.WebKitNavigationAction         // in

	_arg0 = (*C.WebKitNavigationPolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_navigation_policy_decision_get_navigation_action(_arg0)
	runtime.KeepAlive(decision)

	var _navigationAction *NavigationAction // out

	_navigationAction = (*NavigationAction)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _navigationAction
}

// NavigationType gets the value of the
// KitNavigationPolicyDecision:navigation-type property.
//
// Deprecated: Use webkit_navigation_policy_decision_get_navigation_action()
// instead.
//
// The function returns the following values:
//
//    - navigationType: type of navigation triggering this policy decision.
//
func (decision *NavigationPolicyDecision) NavigationType() NavigationType {
	var _arg0 *C.WebKitNavigationPolicyDecision // out
	var _cret C.WebKitNavigationType            // in

	_arg0 = (*C.WebKitNavigationPolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_navigation_policy_decision_get_navigation_type(_arg0)
	runtime.KeepAlive(decision)

	var _navigationType NavigationType // out

	_navigationType = NavigationType(_cret)

	return _navigationType
}

// Request gets the value of the KitNavigationPolicyDecision:request property.
//
// Deprecated: Use webkit_navigation_policy_decision_get_navigation_action()
// instead.
//
// The function returns the following values:
//
//    - uriRequest: URI request that is associated with this navigation.
//
func (decision *NavigationPolicyDecision) Request() *URIRequest {
	var _arg0 *C.WebKitNavigationPolicyDecision // out
	var _cret *C.WebKitURIRequest               // in

	_arg0 = (*C.WebKitNavigationPolicyDecision)(unsafe.Pointer(externglib.InternObject(decision).Native()))

	_cret = C.webkit_navigation_policy_decision_get_request(_arg0)
	runtime.KeepAlive(decision)

	var _uriRequest *URIRequest // out

	_uriRequest = wrapURIRequest(externglib.Take(unsafe.Pointer(_cret)))

	return _uriRequest
}
