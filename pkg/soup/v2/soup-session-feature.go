// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern gboolean _gotk4_soup2_SessionFeatureInterface_add_feature(SoupSessionFeature*, GType);
// extern gboolean _gotk4_soup2_SessionFeatureInterface_has_feature(SoupSessionFeature*, GType);
// extern gboolean _gotk4_soup2_SessionFeatureInterface_remove_feature(SoupSessionFeature*, GType);
// extern void _gotk4_soup2_SessionFeatureInterface_attach(SoupSessionFeature*, SoupSession*);
// extern void _gotk4_soup2_SessionFeatureInterface_detach(SoupSessionFeature*, SoupSession*);
// extern void _gotk4_soup2_SessionFeatureInterface_request_queued(SoupSessionFeature*, SoupSession*, SoupMessage*);
// extern void _gotk4_soup2_SessionFeatureInterface_request_started(SoupSessionFeature*, SoupSession*, SoupMessage*, SoupSocket*);
// extern void _gotk4_soup2_SessionFeatureInterface_request_unqueued(SoupSessionFeature*, SoupSession*, SoupMessage*);
import "C"

// glib.Type values for soup-session-feature.go.
var GTypeSessionFeature = externglib.Type(C.soup_session_feature_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSessionFeature, F: marshalSessionFeature},
	})
}

// SessionFeatureOverrider contains methods that are overridable.
type SessionFeatureOverrider interface {
	// AddFeature adds a "sub-feature" of type type to the base feature feature.
	// This is used for features that can be extended with multiple different
	// types. Eg, the authentication manager can be extended with subtypes of
	// Auth.
	//
	// The function takes the following parameters:
	//
	//    - typ of a "sub-feature".
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if feature accepted type as a subfeature.
	//
	AddFeature(typ externglib.Type) bool
	// The function takes the following parameters:
	//
	Attach(session *Session)
	// The function takes the following parameters:
	//
	Detach(session *Session)
	// HasFeature tests if feature has a "sub-feature" of type type. See
	// soup_session_feature_add_feature().
	//
	// The function takes the following parameters:
	//
	//    - typ of a "sub-feature".
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if feature has a subfeature of type type.
	//
	HasFeature(typ externglib.Type) bool
	// RemoveFeature removes the "sub-feature" of type type from the base
	// feature feature. See soup_session_feature_add_feature().
	//
	// The function takes the following parameters:
	//
	//    - typ of a "sub-feature".
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if type was removed from feature.
	//
	RemoveFeature(typ externglib.Type) bool
	// The function takes the following parameters:
	//
	//    - session
	//    - msg
	//
	RequestQueued(session *Session, msg *Message)
	// The function takes the following parameters:
	//
	//    - session
	//    - msg
	//    - socket
	//
	RequestStarted(session *Session, msg *Message, socket *Socket)
	// The function takes the following parameters:
	//
	//    - session
	//    - msg
	//
	RequestUnqueued(session *Session, msg *Message)
}

// SessionFeature: object that implement some sort of optional feature for
// Session.
//
// SessionFeature wraps an interface. This means the user can get the
// underlying type by calling Cast().
type SessionFeature struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*SessionFeature)(nil)
)

// SessionFeaturer describes SessionFeature's interface methods.
type SessionFeaturer interface {
	externglib.Objector

	// AddFeature adds a "sub-feature" of type type to the base feature feature.
	AddFeature(typ externglib.Type) bool
	Attach(session *Session)
	Detach(session *Session)
	// HasFeature tests if feature has a "sub-feature" of type type.
	HasFeature(typ externglib.Type) bool
	// RemoveFeature removes the "sub-feature" of type type from the base
	// feature feature.
	RemoveFeature(typ externglib.Type) bool
}

var _ SessionFeaturer = (*SessionFeature)(nil)

func ifaceInitSessionFeaturer(gifacePtr, data C.gpointer) {
	iface := (*C.SoupSessionFeatureInterface)(unsafe.Pointer(gifacePtr))
	iface.add_feature = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_add_feature)
	iface.attach = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_attach)
	iface.detach = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_detach)
	iface.has_feature = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_has_feature)
	iface.remove_feature = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_remove_feature)
	iface.request_queued = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_request_queued)
	iface.request_started = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_request_started)
	iface.request_unqueued = (*[0]byte)(C._gotk4_soup2_SessionFeatureInterface_request_unqueued)
}

//export _gotk4_soup2_SessionFeatureInterface_add_feature
func _gotk4_soup2_SessionFeatureInterface_add_feature(arg0 *C.SoupSessionFeature, arg1 C.GType) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _typ externglib.Type // out

	_typ = externglib.Type(arg1)

	ok := iface.AddFeature(_typ)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_soup2_SessionFeatureInterface_attach
func _gotk4_soup2_SessionFeatureInterface_attach(arg0 *C.SoupSessionFeature, arg1 *C.SoupSession) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _session *Session // out

	_session = wrapSession(externglib.Take(unsafe.Pointer(arg1)))

	iface.Attach(_session)
}

//export _gotk4_soup2_SessionFeatureInterface_detach
func _gotk4_soup2_SessionFeatureInterface_detach(arg0 *C.SoupSessionFeature, arg1 *C.SoupSession) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _session *Session // out

	_session = wrapSession(externglib.Take(unsafe.Pointer(arg1)))

	iface.Detach(_session)
}

//export _gotk4_soup2_SessionFeatureInterface_has_feature
func _gotk4_soup2_SessionFeatureInterface_has_feature(arg0 *C.SoupSessionFeature, arg1 C.GType) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _typ externglib.Type // out

	_typ = externglib.Type(arg1)

	ok := iface.HasFeature(_typ)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_soup2_SessionFeatureInterface_remove_feature
func _gotk4_soup2_SessionFeatureInterface_remove_feature(arg0 *C.SoupSessionFeature, arg1 C.GType) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _typ externglib.Type // out

	_typ = externglib.Type(arg1)

	ok := iface.RemoveFeature(_typ)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_soup2_SessionFeatureInterface_request_queued
func _gotk4_soup2_SessionFeatureInterface_request_queued(arg0 *C.SoupSessionFeature, arg1 *C.SoupSession, arg2 *C.SoupMessage) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _session *Session // out
	var _msg *Message     // out

	_session = wrapSession(externglib.Take(unsafe.Pointer(arg1)))
	_msg = wrapMessage(externglib.Take(unsafe.Pointer(arg2)))

	iface.RequestQueued(_session, _msg)
}

//export _gotk4_soup2_SessionFeatureInterface_request_started
func _gotk4_soup2_SessionFeatureInterface_request_started(arg0 *C.SoupSessionFeature, arg1 *C.SoupSession, arg2 *C.SoupMessage, arg3 *C.SoupSocket) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _session *Session // out
	var _msg *Message     // out
	var _socket *Socket   // out

	_session = wrapSession(externglib.Take(unsafe.Pointer(arg1)))
	_msg = wrapMessage(externglib.Take(unsafe.Pointer(arg2)))
	_socket = wrapSocket(externglib.Take(unsafe.Pointer(arg3)))

	iface.RequestStarted(_session, _msg, _socket)
}

//export _gotk4_soup2_SessionFeatureInterface_request_unqueued
func _gotk4_soup2_SessionFeatureInterface_request_unqueued(arg0 *C.SoupSessionFeature, arg1 *C.SoupSession, arg2 *C.SoupMessage) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(SessionFeatureOverrider)

	var _session *Session // out
	var _msg *Message     // out

	_session = wrapSession(externglib.Take(unsafe.Pointer(arg1)))
	_msg = wrapMessage(externglib.Take(unsafe.Pointer(arg2)))

	iface.RequestUnqueued(_session, _msg)
}

func wrapSessionFeature(obj *externglib.Object) *SessionFeature {
	return &SessionFeature{
		Object: obj,
	}
}

func marshalSessionFeature(p uintptr) (interface{}, error) {
	return wrapSessionFeature(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddFeature adds a "sub-feature" of type type to the base feature feature.
// This is used for features that can be extended with multiple different types.
// Eg, the authentication manager can be extended with subtypes of Auth.
//
// The function takes the following parameters:
//
//    - typ of a "sub-feature".
//
// The function returns the following values:
//
//    - ok: TRUE if feature accepted type as a subfeature.
//
func (feature *SessionFeature) AddFeature(typ externglib.Type) bool {
	var _arg0 *C.SoupSessionFeature // out
	var _arg1 C.GType               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupSessionFeature)(unsafe.Pointer(externglib.InternObject(feature).Native()))
	_arg1 = C.GType(typ)

	_cret = C.soup_session_feature_add_feature(_arg0, _arg1)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
func (feature *SessionFeature) Attach(session *Session) {
	var _arg0 *C.SoupSessionFeature // out
	var _arg1 *C.SoupSession        // out

	_arg0 = (*C.SoupSessionFeature)(unsafe.Pointer(externglib.InternObject(feature).Native()))
	_arg1 = (*C.SoupSession)(unsafe.Pointer(externglib.InternObject(session).Native()))

	C.soup_session_feature_attach(_arg0, _arg1)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(session)
}

// The function takes the following parameters:
//
func (feature *SessionFeature) Detach(session *Session) {
	var _arg0 *C.SoupSessionFeature // out
	var _arg1 *C.SoupSession        // out

	_arg0 = (*C.SoupSessionFeature)(unsafe.Pointer(externglib.InternObject(feature).Native()))
	_arg1 = (*C.SoupSession)(unsafe.Pointer(externglib.InternObject(session).Native()))

	C.soup_session_feature_detach(_arg0, _arg1)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(session)
}

// HasFeature tests if feature has a "sub-feature" of type type. See
// soup_session_feature_add_feature().
//
// The function takes the following parameters:
//
//    - typ of a "sub-feature".
//
// The function returns the following values:
//
//    - ok: TRUE if feature has a subfeature of type type.
//
func (feature *SessionFeature) HasFeature(typ externglib.Type) bool {
	var _arg0 *C.SoupSessionFeature // out
	var _arg1 C.GType               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupSessionFeature)(unsafe.Pointer(externglib.InternObject(feature).Native()))
	_arg1 = C.GType(typ)

	_cret = C.soup_session_feature_has_feature(_arg0, _arg1)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveFeature removes the "sub-feature" of type type from the base feature
// feature. See soup_session_feature_add_feature().
//
// The function takes the following parameters:
//
//    - typ of a "sub-feature".
//
// The function returns the following values:
//
//    - ok: TRUE if type was removed from feature.
//
func (feature *SessionFeature) RemoveFeature(typ externglib.Type) bool {
	var _arg0 *C.SoupSessionFeature // out
	var _arg1 C.GType               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupSessionFeature)(unsafe.Pointer(externglib.InternObject(feature).Native()))
	_arg1 = C.GType(typ)

	_cret = C.soup_session_feature_remove_feature(_arg0, _arg1)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
