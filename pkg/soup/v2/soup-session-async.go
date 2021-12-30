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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_session_async_get_type()), F: marshalSessionAsyncer},
	})
}

type SessionAsync struct {
	_ [0]func() // equal guard
	Session
}

var (
	_ externglib.Objector = (*SessionAsync)(nil)
)

func wrapSessionAsync(obj *externglib.Object) *SessionAsync {
	return &SessionAsync{
		Session: Session{
			Object: obj,
		},
	}
}

func marshalSessionAsyncer(p uintptr) (interface{}, error) {
	return wrapSessionAsync(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSessionAsync creates an asynchronous Session with the default options.
//
// Deprecated: SessionAsync is deprecated; use a plain Session, created with
// soup_session_new(). See the <link linkend="libsoup-session-porting">porting
// guide</link>.
//
// The function returns the following values:
//
//    - sessionAsync: new session.
//
func NewSessionAsync() *SessionAsync {
	var _cret *C.SoupSession // in

	_cret = C.soup_session_async_new()

	var _sessionAsync *SessionAsync // out

	_sessionAsync = wrapSessionAsync(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sessionAsync
}
