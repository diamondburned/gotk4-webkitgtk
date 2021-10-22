// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_auth_get_type()), F: marshalAuther},
	})
}

// AUTH_HOST alias for the Auth:host property. (The host being authenticated
// to.).
const AUTH_HOST = "host"

// AUTH_IS_AUTHENTICATED alias for the Auth:is-authenticated property. (Whether
// or not the auth has been authenticated.).
const AUTH_IS_AUTHENTICATED = "is-authenticated"

// AUTH_IS_FOR_PROXY alias for the Auth:is-for-proxy property. (Whether or not
// the auth is for a proxy server.).
const AUTH_IS_FOR_PROXY = "is-for-proxy"

// AUTH_REALM alias for the Auth:realm property. (The authentication realm.).
const AUTH_REALM = "realm"

// AUTH_SCHEME_NAME alias for the Auth:scheme-name property. (The authentication
// scheme name.).
const AUTH_SCHEME_NAME = "scheme-name"

// Auth: abstract base class for handling authentication. Specific HTTP
// Authentication mechanisms are implemented by its subclasses, but applications
// never need to be aware of the specific subclasses being used.
type Auth struct {
	*externglib.Object
}

// Auther describes types inherited from class Auth.
// To get the original type, the caller must assert this to an interface or
// another type.
type Auther interface {
	externglib.Objector

	// BaseAuth returns the underlying base class.
	BaseAuth() *Auth
}

var _ Auther = (*Auth)(nil)

func wrapAuth(obj *externglib.Object) *Auth {
	return &Auth{
		Object: obj,
	}
}

func marshalAuther(p uintptr) (interface{}, error) {
	return wrapAuth(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAuth creates a new Auth of type type with the information from msg and
// auth_header.
//
// This is called by Session; you will normally not create auths yourself.
//
// The function takes the following parameters:
//
//    - typ: type of auth to create (a subtype of Auth).
//    - msg the auth is being created for.
//    - authHeader: WWW-Authenticate/Proxy-Authenticate header.
//
func NewAuth(typ externglib.Type, msg *Message, authHeader string) *Auth {
	var _arg1 C.GType        // out
	var _arg2 *C.SoupMessage // out
	var _arg3 *C.char        // out
	var _cret *C.SoupAuth    // in

	_arg1 = C.GType(typ)
	_arg2 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(authHeader)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.soup_auth_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(authHeader)

	var _auth *Auth // out

	if _cret != nil {
		_auth = wrapAuth(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _auth
}

// Authenticate: call this on an auth to authenticate it; normally this will
// cause the auth's message to be requeued with the new authentication info.
//
// The function takes the following parameters:
//
//    - username provided by the user or client.
//    - password provided by the user or client.
//
func (auth *Auth) Authenticate(username, password string) {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _arg2 *C.char     // out

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg2))

	C.soup_auth_authenticate(_arg0, _arg1, _arg2)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)
}

// CanAuthenticate tests if auth is able to authenticate by providing
// credentials to the soup_auth_authenticate().
func (auth *Auth) CanAuthenticate() bool {
	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_can_authenticate(_arg0)
	runtime.KeepAlive(auth)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Authorization generates an appropriate "Authorization" header for msg. (The
// session will only call this if soup_auth_is_authenticated() returned TRUE.).
//
// The function takes the following parameters:
//
//    - msg to be authorized.
//
func (auth *Auth) Authorization(msg *Message) string {
	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _cret *C.char        // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	_cret = C.soup_auth_get_authorization(_arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Host returns the host that auth is associated with.
func (auth *Auth) Host() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_get_host(_arg0)
	runtime.KeepAlive(auth)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Info gets an opaque identifier for auth, for use as a hash key or the like.
// Auth objects from the same server with the same identifier refer to the same
// authentication domain (eg, the URLs associated with them take the same
// usernames and passwords).
func (auth *Auth) Info() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_get_info(_arg0)
	runtime.KeepAlive(auth)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ProtectionSpace returns a list of paths on the server which auth extends
// over. (All subdirectories of these paths are also assumed to be part of
// auth's protection space, unless otherwise discovered not to be.).
//
// The function takes the following parameters:
//
//    - sourceUri: URI of the request that auth was generated in response to.
//
func (auth *Auth) ProtectionSpace(sourceUri *URI) []string {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.SoupURI  // out
	var _cret *C.GSList   // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(sourceUri)))

	_cret = C.soup_auth_get_protection_space(_arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(sourceUri)

	var _sList []string // out

	_sList = make([]string, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		defer C.free(unsafe.Pointer(src))
		_sList = append(_sList, dst)
	})

	return _sList
}

// Realm returns auth's realm. This is an identifier that distinguishes separate
// authentication spaces on a given server, and may be some string that is
// meaningful to the user. (Although it is probably not localized.).
func (auth *Auth) Realm() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_get_realm(_arg0)
	runtime.KeepAlive(auth)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

//
// The function takes the following parameters:
//

//
func (auth *Auth) SavedPassword(user string) string {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(user)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_auth_get_saved_password(_arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(user)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

func (auth *Auth) SavedUsers() []string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.GSList   // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_get_saved_users(_arg0)
	runtime.KeepAlive(auth)

	var _sList []string // out

	_sList = make([]string, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		defer C.free(unsafe.Pointer(src))
		_sList = append(_sList, dst)
	})

	return _sList
}

// SchemeName returns auth's scheme name. (Eg, "Basic", "Digest", or "NTLM").
func (auth *Auth) SchemeName() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_get_scheme_name(_arg0)
	runtime.KeepAlive(auth)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

//
// The function takes the following parameters:
//

//
func (auth *Auth) HasSavedPassword(username, password string) {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _arg2 *C.char     // out

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg2))

	C.soup_auth_has_saved_password(_arg0, _arg1, _arg2)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)
}

// IsAuthenticated tests if auth has been given a username and password.
func (auth *Auth) IsAuthenticated() bool {
	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_is_authenticated(_arg0)
	runtime.KeepAlive(auth)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsForProxy tests whether or not auth is associated with a proxy server rather
// than an "origin" server.
func (auth *Auth) IsForProxy() bool {
	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))

	_cret = C.soup_auth_is_for_proxy(_arg0)
	runtime.KeepAlive(auth)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsReady tests if auth is ready to make a request for msg with. For most
// auths, this is equivalent to soup_auth_is_authenticated(), but for some auth
// types (eg, NTLM), the auth may be sendable (eg, as an authentication request)
// even before it is authenticated.
//
// The function takes the following parameters:
//
//    - msg: Message.
//
func (auth *Auth) IsReady(msg *Message) bool {
	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _cret C.gboolean     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	_cret = C.soup_auth_is_ready(_arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

//
// The function takes the following parameters:
//

//
func (auth *Auth) SavePassword(username, password string) {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _arg2 *C.char     // out

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg2))

	C.soup_auth_save_password(_arg0, _arg1, _arg2)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)
}

// Update updates auth with the information from msg and auth_header, possibly
// un-authenticating it. As with soup_auth_new(), this is normally only used by
// Session.
//
// The function takes the following parameters:
//
//    - msg auth is being updated for.
//    - authHeader: WWW-Authenticate/Proxy-Authenticate header.
//
func (auth *Auth) Update(msg *Message, authHeader string) bool {
	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _arg2 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(auth.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(authHeader)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.soup_auth_update(_arg0, _arg1, _arg2)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(authHeader)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BaseAuth returns auth.
func (auth *Auth) BaseAuth() *Auth {
	return auth
}

// AuthNegotiateSupported indicates whether libsoup was built with GSSAPI
// support. If this is FALSE, SOUP_TYPE_AUTH_NEGOTIATE will still be defined and
// can still be added to a Session, but libsoup will never attempt to actually
// use this auth type.
func AuthNegotiateSupported() bool {
	var _cret C.gboolean // in

	_cret = C.soup_auth_negotiate_supported()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
