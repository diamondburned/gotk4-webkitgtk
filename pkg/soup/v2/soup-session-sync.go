// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// glib.Type values for soup-session-sync.go.
var GTypeSessionSync = externglib.Type(C.soup_session_sync_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSessionSync, F: marshalSessionSync},
	})
}

// SessionSyncOverrider contains methods that are overridable.
type SessionSyncOverrider interface {
}

type SessionSync struct {
	_ [0]func() // equal guard
	Session
}

var (
	_ externglib.Objector = (*SessionSync)(nil)
)

func classInitSessionSyncer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSessionSync(obj *externglib.Object) *SessionSync {
	return &SessionSync{
		Session: Session{
			Object: obj,
		},
	}
}

func marshalSessionSync(p uintptr) (interface{}, error) {
	return wrapSessionSync(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSessionSync creates an synchronous Session with the default options.
//
// Deprecated: SessionSync is deprecated; use a plain Session, created with
// soup_session_new(). See the <link linkend="libsoup-session-porting">porting
// guide</link>.
//
// The function returns the following values:
//
//    - sessionSync: new session.
//
func NewSessionSync() *SessionSync {
	var _cret *C.SoupSession // in

	_cret = C.soup_session_sync_new()

	var _sessionSync *SessionSync // out

	_sessionSync = wrapSessionSync(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sessionSync
}
