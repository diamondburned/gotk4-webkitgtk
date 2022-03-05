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
import "C"

// glib.Type values for WebKitNavigationAction.go.
var (
	GTypeNavigationType   = externglib.Type(C.webkit_navigation_type_get_type())
	GTypeNavigationAction = externglib.Type(C.webkit_navigation_action_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeNavigationType, F: marshalNavigationType},
		{T: GTypeNavigationAction, F: marshalNavigationAction},
	})
}

// NavigationType: enum values used to denote the various navigation types.
type NavigationType C.gint

const (
	// NavigationTypeLinkClicked: navigation was triggered by clicking a link.
	NavigationTypeLinkClicked NavigationType = iota
	// NavigationTypeFormSubmitted: navigation was triggered by submitting a
	// form.
	NavigationTypeFormSubmitted
	// NavigationTypeBackForward: navigation was triggered by navigating forward
	// or backward.
	NavigationTypeBackForward
	// NavigationTypeReload: navigation was triggered by reloading.
	NavigationTypeReload
	// NavigationTypeFormResubmitted: navigation was triggered by resubmitting a
	// form.
	NavigationTypeFormResubmitted
	// NavigationTypeOther: navigation was triggered by some other action.
	NavigationTypeOther
)

func marshalNavigationType(p uintptr) (interface{}, error) {
	return NavigationType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for NavigationType.
func (n NavigationType) String() string {
	switch n {
	case NavigationTypeLinkClicked:
		return "LinkClicked"
	case NavigationTypeFormSubmitted:
		return "FormSubmitted"
	case NavigationTypeBackForward:
		return "BackForward"
	case NavigationTypeReload:
		return "Reload"
	case NavigationTypeFormResubmitted:
		return "FormResubmitted"
	case NavigationTypeOther:
		return "Other"
	default:
		return fmt.Sprintf("NavigationType(%d)", n)
	}
}

// NavigationAction: instance of this type is always passed by reference.
type NavigationAction struct {
	*navigationAction
}

// navigationAction is the struct that's finalized.
type navigationAction struct {
	native *C.WebKitNavigationAction
}

func marshalNavigationAction(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &NavigationAction{&navigationAction{(*C.WebKitNavigationAction)(b)}}, nil
}

// Copy: make a copy of navigation.
//
// The function returns the following values:
//
//    - navigationAction: copy of passed in KitNavigationAction.
//
func (navigation *NavigationAction) Copy() *NavigationAction {
	var _arg0 *C.WebKitNavigationAction // out
	var _cret *C.WebKitNavigationAction // in

	_arg0 = (*C.WebKitNavigationAction)(gextras.StructNative(unsafe.Pointer(navigation)))

	_cret = C.webkit_navigation_action_copy(_arg0)
	runtime.KeepAlive(navigation)

	var _navigationAction *NavigationAction // out

	_navigationAction = (*NavigationAction)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_navigationAction)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_navigation_action_free((*C.WebKitNavigationAction)(intern.C))
		},
	)

	return _navigationAction
}

// Modifiers: return a bitmask of ModifierType values describing the modifier
// keys that were in effect when the navigation was requested.
//
// The function returns the following values:
//
//    - guint: modifier keys.
//
func (navigation *NavigationAction) Modifiers() uint {
	var _arg0 *C.WebKitNavigationAction // out
	var _cret C.guint                   // in

	_arg0 = (*C.WebKitNavigationAction)(gextras.StructNative(unsafe.Pointer(navigation)))

	_cret = C.webkit_navigation_action_get_modifiers(_arg0)
	runtime.KeepAlive(navigation)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MouseButton: return the number of the mouse button that triggered the
// navigation, or 0 if the navigation was not started by a mouse event.
//
// The function returns the following values:
//
//    - guint: mouse button number or 0.
//
func (navigation *NavigationAction) MouseButton() uint {
	var _arg0 *C.WebKitNavigationAction // out
	var _cret C.guint                   // in

	_arg0 = (*C.WebKitNavigationAction)(gextras.StructNative(unsafe.Pointer(navigation)))

	_cret = C.webkit_navigation_action_get_mouse_button(_arg0)
	runtime.KeepAlive(navigation)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NavigationType: return the type of action that triggered the navigation.
//
// The function returns the following values:
//
//    - navigationType: KitNavigationType.
//
func (navigation *NavigationAction) NavigationType() NavigationType {
	var _arg0 *C.WebKitNavigationAction // out
	var _cret C.WebKitNavigationType    // in

	_arg0 = (*C.WebKitNavigationAction)(gextras.StructNative(unsafe.Pointer(navigation)))

	_cret = C.webkit_navigation_action_get_navigation_type(_arg0)
	runtime.KeepAlive(navigation)

	var _navigationType NavigationType // out

	_navigationType = NavigationType(_cret)

	return _navigationType
}

// Request: return the KitURIRequest associated with the navigation action.
// Modifications to the returned object are <emphasis>not</emphasis> taken into
// account when the request is sent over the network, and is intended only to
// aid in evaluating whether a navigation action should be taken or not. To
// modify requests before they are sent over the network the
// KitPage::send-request signal can be used instead.
//
// The function returns the following values:
//
//    - uriRequest: KitURIRequest.
//
func (navigation *NavigationAction) Request() *URIRequest {
	var _arg0 *C.WebKitNavigationAction // out
	var _cret *C.WebKitURIRequest       // in

	_arg0 = (*C.WebKitNavigationAction)(gextras.StructNative(unsafe.Pointer(navigation)))

	_cret = C.webkit_navigation_action_get_request(_arg0)
	runtime.KeepAlive(navigation)

	var _uriRequest *URIRequest // out

	_uriRequest = wrapURIRequest(externglib.Take(unsafe.Pointer(_cret)))

	return _uriRequest
}

// IsRedirect returns whether the navigation was redirected.
//
// The function returns the following values:
//
//    - ok: TRUE if the original navigation was redirected, FALSE otherwise.
//
func (navigation *NavigationAction) IsRedirect() bool {
	var _arg0 *C.WebKitNavigationAction // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitNavigationAction)(gextras.StructNative(unsafe.Pointer(navigation)))

	_cret = C.webkit_navigation_action_is_redirect(_arg0)
	runtime.KeepAlive(navigation)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsUserGesture: return whether the navigation was triggered by a user gesture
// like a mouse click.
//
// The function returns the following values:
//
//    - ok: whether navigation action is a user gesture.
//
func (navigation *NavigationAction) IsUserGesture() bool {
	var _arg0 *C.WebKitNavigationAction // out
	var _cret C.gboolean                // in

	_arg0 = (*C.WebKitNavigationAction)(gextras.StructNative(unsafe.Pointer(navigation)))

	_cret = C.webkit_navigation_action_is_user_gesture(_arg0)
	runtime.KeepAlive(navigation)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
