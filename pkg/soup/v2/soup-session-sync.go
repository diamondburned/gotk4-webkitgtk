// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// GType values.
var (
	GTypeSessionSync = coreglib.Type(C.soup_session_sync_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSessionSync, F: marshalSessionSync},
	})
}

// SessionSyncOverrides contains methods that are overridable.
type SessionSyncOverrides struct {
}

func defaultSessionSyncOverrides(v *SessionSync) SessionSyncOverrides {
	return SessionSyncOverrides{}
}

type SessionSync struct {
	_ [0]func() // equal guard
	Session
}

var (
	_ coreglib.Objector = (*SessionSync)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SessionSync, *SessionSyncClass, SessionSyncOverrides](
		GTypeSessionSync,
		initSessionSyncClass,
		wrapSessionSync,
		defaultSessionSyncOverrides,
	)
}

func initSessionSyncClass(gclass unsafe.Pointer, overrides SessionSyncOverrides, classInitFunc func(*SessionSyncClass)) {
	if classInitFunc != nil {
		class := (*SessionSyncClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSessionSync(obj *coreglib.Object) *SessionSync {
	return &SessionSync{
		Session: Session{
			Object: obj,
		},
	}
}

func marshalSessionSync(p uintptr) (interface{}, error) {
	return wrapSessionSync(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSessionSync creates an synchronous Session with the default options.
//
// Deprecated: SessionSync is deprecated; use a plain Session, created with
// soup_session_new(). See the <link linkend="libsoup-session-porting">porting
// guide</link>.
//
// The function returns the following values:
//
//   - sessionSync: new session.
//
func NewSessionSync() *SessionSync {
	var _cret *C.SoupSession // in

	_cret = C.soup_session_sync_new()

	var _sessionSync *SessionSync // out

	_sessionSync = wrapSessionSync(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sessionSync
}

// SessionSyncClass: instance of this type is always passed by reference.
type SessionSyncClass struct {
	*sessionSyncClass
}

// sessionSyncClass is the struct that's finalized.
type sessionSyncClass struct {
	native *C.SoupSessionSyncClass
}

func (s *SessionSyncClass) ParentClass() *SessionClass {
	valptr := &s.native.parent_class
	var _v *SessionClass // out
	_v = (*SessionClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
