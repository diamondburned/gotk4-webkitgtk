// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_credential_persistence_get_type()), F: marshalCredentialPersistence},
		{T: externglib.Type(C.webkit_credential_get_type()), F: marshalCredential},
	})
}

// CredentialPersistence: enum values representing the duration for which a
// credential persists.
type CredentialPersistence int

const (
	// CredentialPersistenceNone: credential does not persist.
	CredentialPersistenceNone CredentialPersistence = iota
	// CredentialPersistenceForSession: credential persists for session only.
	CredentialPersistenceForSession
	// CredentialPersistencePermanent: credential persists permanently.
	CredentialPersistencePermanent
)

func marshalCredentialPersistence(p uintptr) (interface{}, error) {
	return CredentialPersistence(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CredentialPersistence.
func (c CredentialPersistence) String() string {
	switch c {
	case CredentialPersistenceNone:
		return "None"
	case CredentialPersistenceForSession:
		return "ForSession"
	case CredentialPersistencePermanent:
		return "Permanent"
	default:
		return fmt.Sprintf("CredentialPersistence(%d)", c)
	}
}

// Credential: instance of this type is always passed by reference.
type Credential struct {
	*credential
}

// credential is the struct that's finalized.
type credential struct {
	native *C.WebKitCredential
}

func marshalCredential(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Credential{&credential{(*C.WebKitCredential)(b)}}, nil
}

// NewCredential constructs a struct Credential.
func NewCredential(username string, password string, persistence CredentialPersistence) *Credential {
	var _arg1 *C.gchar                      // out
	var _arg2 *C.gchar                      // out
	var _arg3 C.WebKitCredentialPersistence // out
	var _cret *C.WebKitCredential           // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.WebKitCredentialPersistence(persistence)

	_cret = C.webkit_credential_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)
	runtime.KeepAlive(persistence)

	var _credential *Credential // out

	_credential = (*Credential)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_credential)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_credential_free((*C.WebKitCredential)(intern.C))
		},
	)

	return _credential
}

// Copy: make a copy of the KitCredential.
func (credential *Credential) Copy() *Credential {
	var _arg0 *C.WebKitCredential // out
	var _cret *C.WebKitCredential // in

	_arg0 = (*C.WebKitCredential)(gextras.StructNative(unsafe.Pointer(credential)))

	_cret = C.webkit_credential_copy(_arg0)
	runtime.KeepAlive(credential)

	var _ret *Credential // out

	_ret = (*Credential)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_credential_free((*C.WebKitCredential)(intern.C))
		},
	)

	return _ret
}

// Password: get the password currently held by this KitCredential.
func (credential *Credential) Password() string {
	var _arg0 *C.WebKitCredential // out
	var _cret *C.gchar            // in

	_arg0 = (*C.WebKitCredential)(gextras.StructNative(unsafe.Pointer(credential)))

	_cret = C.webkit_credential_get_password(_arg0)
	runtime.KeepAlive(credential)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Persistence: get the persistence mode currently held by this KitCredential.
func (credential *Credential) Persistence() CredentialPersistence {
	var _arg0 *C.WebKitCredential           // out
	var _cret C.WebKitCredentialPersistence // in

	_arg0 = (*C.WebKitCredential)(gextras.StructNative(unsafe.Pointer(credential)))

	_cret = C.webkit_credential_get_persistence(_arg0)
	runtime.KeepAlive(credential)

	var _credentialPersistence CredentialPersistence // out

	_credentialPersistence = CredentialPersistence(_cret)

	return _credentialPersistence
}

// Username: get the username currently held by this KitCredential.
func (credential *Credential) Username() string {
	var _arg0 *C.WebKitCredential // out
	var _cret *C.gchar            // in

	_arg0 = (*C.WebKitCredential)(gextras.StructNative(unsafe.Pointer(credential)))

	_cret = C.webkit_credential_get_username(_arg0)
	runtime.KeepAlive(credential)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasPassword: determine whether this credential has a password stored.
func (credential *Credential) HasPassword() bool {
	var _arg0 *C.WebKitCredential // out
	var _cret C.gboolean          // in

	_arg0 = (*C.WebKitCredential)(gextras.StructNative(unsafe.Pointer(credential)))

	_cret = C.webkit_credential_has_password(_arg0)
	runtime.KeepAlive(credential)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
