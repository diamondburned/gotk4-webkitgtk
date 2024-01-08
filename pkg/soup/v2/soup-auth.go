// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern void _gotk4_soup2_AuthClass_authenticate(SoupAuth*, char*, char*);
// extern gboolean _gotk4_soup2_AuthClass_update(SoupAuth*, SoupMessage*, GHashTable*);
// extern gboolean _gotk4_soup2_AuthClass_is_ready(SoupAuth*, SoupMessage*);
// extern gboolean _gotk4_soup2_AuthClass_is_authenticated(SoupAuth*);
// extern gboolean _gotk4_soup2_AuthClass_can_authenticate(SoupAuth*);
// extern char* _gotk4_soup2_AuthClass_get_authorization(SoupAuth*, SoupMessage*);
// extern GSList* _gotk4_soup2_AuthClass_get_protection_space(SoupAuth*, SoupURI*);
// GSList* _gotk4_soup2_Auth_virtual_get_protection_space(void* fnptr, SoupAuth* arg0, SoupURI* arg1) {
//   return ((GSList* (*)(SoupAuth*, SoupURI*))(fnptr))(arg0, arg1);
// };
// char* _gotk4_soup2_Auth_virtual_get_authorization(void* fnptr, SoupAuth* arg0, SoupMessage* arg1) {
//   return ((char* (*)(SoupAuth*, SoupMessage*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_soup2_Auth_virtual_can_authenticate(void* fnptr, SoupAuth* arg0) {
//   return ((gboolean (*)(SoupAuth*))(fnptr))(arg0);
// };
// gboolean _gotk4_soup2_Auth_virtual_is_authenticated(void* fnptr, SoupAuth* arg0) {
//   return ((gboolean (*)(SoupAuth*))(fnptr))(arg0);
// };
// gboolean _gotk4_soup2_Auth_virtual_is_ready(void* fnptr, SoupAuth* arg0, SoupMessage* arg1) {
//   return ((gboolean (*)(SoupAuth*, SoupMessage*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_soup2_Auth_virtual_update(void* fnptr, SoupAuth* arg0, SoupMessage* arg1, GHashTable* arg2) {
//   return ((gboolean (*)(SoupAuth*, SoupMessage*, GHashTable*))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_soup2_Auth_virtual_authenticate(void* fnptr, SoupAuth* arg0, char* arg1, char* arg2) {
//   ((void (*)(SoupAuth*, char*, char*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeAuth = coreglib.Type(C.soup_auth_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAuth, F: marshalAuth},
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

// AuthOverrides contains methods that are overridable.
type AuthOverrides struct {
	// Authenticate: call this on an auth to authenticate it; normally this will
	// cause the auth's message to be requeued with the new authentication info.
	//
	// The function takes the following parameters:
	//
	//   - username provided by the user or client.
	//   - password provided by the user or client.
	//
	Authenticate func(username, password string)
	// CanAuthenticate tests if auth is able to authenticate by providing
	// credentials to the soup_auth_authenticate().
	//
	// The function returns the following values:
	//
	//   - ok: TRUE if auth is able to accept credentials.
	//
	CanAuthenticate func() bool
	// Authorization generates an appropriate "Authorization" header for msg.
	// (The session will only call this if soup_auth_is_authenticated() returned
	// TRUE.).
	//
	// The function takes the following parameters:
	//
	//   - msg to be authorized.
	//
	// The function returns the following values:
	//
	//   - utf8: "Authorization" header, which must be freed.
	//
	Authorization func(msg *Message) string
	// ProtectionSpace returns a list of paths on the server which auth extends
	// over. (All subdirectories of these paths are also assumed to be part of
	// auth's protection space, unless otherwise discovered not to be.).
	//
	// The function takes the following parameters:
	//
	//   - sourceUri: URI of the request that auth was generated in response to.
	//
	// The function returns the following values:
	//
	//   - sList: list of paths, which can be freed with
	//     soup_auth_free_protection_space().
	//
	ProtectionSpace func(sourceUri *URI) []string
	// IsAuthenticated tests if auth has been given a username and password.
	//
	// The function returns the following values:
	//
	//   - ok: TRUE if auth has been given a username and password.
	//
	IsAuthenticated func() bool
	// IsReady tests if auth is ready to make a request for msg with. For most
	// auths, this is equivalent to soup_auth_is_authenticated(), but for some
	// auth types (eg, NTLM), the auth may be sendable (eg, as an authentication
	// request) even before it is authenticated.
	//
	// The function takes the following parameters:
	//
	//   - msg: Message.
	//
	// The function returns the following values:
	//
	//   - ok: TRUE if auth is ready to make a request with.
	//
	IsReady func(msg *Message) bool
	// Update updates auth with the information from msg and auth_header,
	// possibly un-authenticating it. As with soup_auth_new(), this is normally
	// only used by Session.
	//
	// The function takes the following parameters:
	//
	//   - msg auth is being updated for.
	//   - authHeader: WWW-Authenticate/Proxy-Authenticate header.
	//
	// The function returns the following values:
	//
	//   - ok: TRUE if auth is still a valid (but potentially unauthenticated)
	//     Auth. FALSE if something about auth_params could not be parsed or
	//     incorporated into auth at all.
	//
	Update func(msg *Message, authHeader map[unsafe.Pointer]unsafe.Pointer) bool
}

func defaultAuthOverrides(v *Auth) AuthOverrides {
	return AuthOverrides{
		Authenticate:    v.authenticate,
		CanAuthenticate: v.canAuthenticate,
		Authorization:   v.authorization,
		ProtectionSpace: v.protectionSpace,
		IsAuthenticated: v.isAuthenticated,
		IsReady:         v.isReady,
		Update:          v.update,
	}
}

// Auth: abstract base class for handling authentication. Specific HTTP
// Authentication mechanisms are implemented by its subclasses, but applications
// never need to be aware of the specific subclasses being used.
type Auth struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Auth)(nil)
)

// Auther describes types inherited from class Auth.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Auther interface {
	coreglib.Objector
	baseAuth() *Auth
}

var _ Auther = (*Auth)(nil)

func init() {
	coreglib.RegisterClassInfo[*Auth, *AuthClass, AuthOverrides](
		GTypeAuth,
		initAuthClass,
		wrapAuth,
		defaultAuthOverrides,
	)
}

func initAuthClass(gclass unsafe.Pointer, overrides AuthOverrides, classInitFunc func(*AuthClass)) {
	pclass := (*C.SoupAuthClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeAuth))))

	if overrides.Authenticate != nil {
		pclass.authenticate = (*[0]byte)(C._gotk4_soup2_AuthClass_authenticate)
	}

	if overrides.CanAuthenticate != nil {
		pclass.can_authenticate = (*[0]byte)(C._gotk4_soup2_AuthClass_can_authenticate)
	}

	if overrides.Authorization != nil {
		pclass.get_authorization = (*[0]byte)(C._gotk4_soup2_AuthClass_get_authorization)
	}

	if overrides.ProtectionSpace != nil {
		pclass.get_protection_space = (*[0]byte)(C._gotk4_soup2_AuthClass_get_protection_space)
	}

	if overrides.IsAuthenticated != nil {
		pclass.is_authenticated = (*[0]byte)(C._gotk4_soup2_AuthClass_is_authenticated)
	}

	if overrides.IsReady != nil {
		pclass.is_ready = (*[0]byte)(C._gotk4_soup2_AuthClass_is_ready)
	}

	if overrides.Update != nil {
		pclass.update = (*[0]byte)(C._gotk4_soup2_AuthClass_update)
	}

	if classInitFunc != nil {
		class := (*AuthClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAuth(obj *coreglib.Object) *Auth {
	return &Auth{
		Object: obj,
	}
}

func marshalAuth(p uintptr) (interface{}, error) {
	return wrapAuth(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (auth *Auth) baseAuth() *Auth {
	return auth
}

// BaseAuth returns the underlying base object.
func BaseAuth(obj Auther) *Auth {
	return obj.baseAuth()
}

// NewAuth creates a new Auth of type type with the information from msg and
// auth_header.
//
// This is called by Session; you will normally not create auths yourself.
//
// The function takes the following parameters:
//
//   - typ: type of auth to create (a subtype of Auth).
//   - msg the auth is being created for.
//   - authHeader: WWW-Authenticate/Proxy-Authenticate header.
//
// The function returns the following values:
//
//   - auth (optional): new Auth, or NULL if it could not be created.
//
func NewAuth(typ coreglib.Type, msg *Message, authHeader string) *Auth {
	var _arg1 C.GType        // out
	var _arg2 *C.SoupMessage // out
	var _arg3 *C.char        // out
	var _cret *C.SoupAuth    // in

	_arg1 = C.GType(typ)
	_arg2 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(authHeader)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.soup_auth_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(authHeader)

	var _auth *Auth // out

	if _cret != nil {
		_auth = wrapAuth(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _auth
}

// Authenticate: call this on an auth to authenticate it; normally this will
// cause the auth's message to be requeued with the new authentication info.
//
// The function takes the following parameters:
//
//   - username provided by the user or client.
//   - password provided by the user or client.
//
func (auth *Auth) Authenticate(username, password string) {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _arg2 *C.char     // out

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
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
//
// The function returns the following values:
//
//   - ok: TRUE if auth is able to accept credentials.
//
func (auth *Auth) CanAuthenticate() bool {
	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

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
//   - msg to be authorized.
//
// The function returns the following values:
//
//   - utf8: "Authorization" header, which must be freed.
//
func (auth *Auth) Authorization(msg *Message) string {
	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _cret *C.char        // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))

	_cret = C.soup_auth_get_authorization(_arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Host returns the host that auth is associated with.
//
// The function returns the following values:
//
//   - utf8: hostname.
//
func (auth *Auth) Host() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

	_cret = C.soup_auth_get_host(_arg0)
	runtime.KeepAlive(auth)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Info gets an opaque identifier for auth, for use as a hash key or the like.
// Auth objects from the same server with the same identifier refer to the
// same authentication domain (eg, the URLs associated with them take the same
// usernames and passwords).
//
// The function returns the following values:
//
//   - utf8: identifier.
//
func (auth *Auth) Info() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

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
//   - sourceUri: URI of the request that auth was generated in response to.
//
// The function returns the following values:
//
//   - sList: list of paths, which can be freed with
//     soup_auth_free_protection_space().
//
func (auth *Auth) ProtectionSpace(sourceUri *URI) []string {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.SoupURI  // out
	var _cret *C.GSList   // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
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
//
// The function returns the following values:
//
//   - utf8: realm name.
//
func (auth *Auth) Realm() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

	_cret = C.soup_auth_get_realm(_arg0)
	runtime.KeepAlive(auth)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (auth *Auth) SavedPassword(user string) string {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(user)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_auth_get_saved_password(_arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(user)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function returns the following values:
//
func (auth *Auth) SavedUsers() []string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.GSList   // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

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
//
// The function returns the following values:
//
//   - utf8: scheme name.
//
func (auth *Auth) SchemeName() string {
	var _arg0 *C.SoupAuth // out
	var _cret *C.char     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

	_cret = C.soup_auth_get_scheme_name(_arg0)
	runtime.KeepAlive(auth)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
//   - username
//   - password
//
func (auth *Auth) HasSavedPassword(username, password string) {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _arg2 *C.char     // out

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
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
//
// The function returns the following values:
//
//   - ok: TRUE if auth has been given a username and password.
//
func (auth *Auth) IsAuthenticated() bool {
	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

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
//
// The function returns the following values:
//
//   - ok: TRUE or FALSE.
//
func (auth *Auth) IsForProxy() bool {
	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

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
//   - msg: Message.
//
// The function returns the following values:
//
//   - ok: TRUE if auth is ready to make a request with.
//
func (auth *Auth) IsReady(msg *Message) bool {
	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _cret C.gboolean     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))

	_cret = C.soup_auth_is_ready(_arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//   - username
//   - password
//
func (auth *Auth) SavePassword(username, password string) {
	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _arg2 *C.char     // out

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
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
//   - msg auth is being updated for.
//   - authHeader: WWW-Authenticate/Proxy-Authenticate header.
//
// The function returns the following values:
//
//   - ok: TRUE if auth is still a valid (but potentially unauthenticated) Auth.
//     FALSE if something about auth_params could not be parsed or incorporated
//     into auth at all.
//
func (auth *Auth) Update(msg *Message, authHeader string) bool {
	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _arg2 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))
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

// Authenticate: call this on an auth to authenticate it; normally this will
// cause the auth's message to be requeued with the new authentication info.
//
// The function takes the following parameters:
//
//   - username provided by the user or client.
//   - password provided by the user or client.
//
func (auth *Auth) authenticate(username, password string) {
	gclass := (*C.SoupAuthClass)(coreglib.PeekParentClass(auth))
	fnarg := gclass.authenticate

	var _arg0 *C.SoupAuth // out
	var _arg1 *C.char     // out
	var _arg2 *C.char     // out

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg2))

	C._gotk4_soup2_Auth_virtual_authenticate(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)
}

// canAuthenticate tests if auth is able to authenticate by providing
// credentials to the soup_auth_authenticate().
//
// The function returns the following values:
//
//   - ok: TRUE if auth is able to accept credentials.
//
func (auth *Auth) canAuthenticate() bool {
	gclass := (*C.SoupAuthClass)(coreglib.PeekParentClass(auth))
	fnarg := gclass.can_authenticate

	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

	_cret = C._gotk4_soup2_Auth_virtual_can_authenticate(unsafe.Pointer(fnarg), _arg0)
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
//   - msg to be authorized.
//
// The function returns the following values:
//
//   - utf8: "Authorization" header, which must be freed.
//
func (auth *Auth) authorization(msg *Message) string {
	gclass := (*C.SoupAuthClass)(coreglib.PeekParentClass(auth))
	fnarg := gclass.get_authorization

	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _cret *C.char        // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))

	_cret = C._gotk4_soup2_Auth_virtual_get_authorization(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// protectionSpace returns a list of paths on the server which auth extends
// over. (All subdirectories of these paths are also assumed to be part of
// auth's protection space, unless otherwise discovered not to be.).
//
// The function takes the following parameters:
//
//   - sourceUri: URI of the request that auth was generated in response to.
//
// The function returns the following values:
//
//   - sList: list of paths, which can be freed with
//     soup_auth_free_protection_space().
//
func (auth *Auth) protectionSpace(sourceUri *URI) []string {
	gclass := (*C.SoupAuthClass)(coreglib.PeekParentClass(auth))
	fnarg := gclass.get_protection_space

	var _arg0 *C.SoupAuth // out
	var _arg1 *C.SoupURI  // out
	var _cret *C.GSList   // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(sourceUri)))

	_cret = C._gotk4_soup2_Auth_virtual_get_protection_space(unsafe.Pointer(fnarg), _arg0, _arg1)
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

// isAuthenticated tests if auth has been given a username and password.
//
// The function returns the following values:
//
//   - ok: TRUE if auth has been given a username and password.
//
func (auth *Auth) isAuthenticated() bool {
	gclass := (*C.SoupAuthClass)(coreglib.PeekParentClass(auth))
	fnarg := gclass.is_authenticated

	var _arg0 *C.SoupAuth // out
	var _cret C.gboolean  // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))

	_cret = C._gotk4_soup2_Auth_virtual_is_authenticated(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(auth)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// isReady tests if auth is ready to make a request for msg with. For most
// auths, this is equivalent to soup_auth_is_authenticated(), but for some auth
// types (eg, NTLM), the auth may be sendable (eg, as an authentication request)
// even before it is authenticated.
//
// The function takes the following parameters:
//
//   - msg: Message.
//
// The function returns the following values:
//
//   - ok: TRUE if auth is ready to make a request with.
//
func (auth *Auth) isReady(msg *Message) bool {
	gclass := (*C.SoupAuthClass)(coreglib.PeekParentClass(auth))
	fnarg := gclass.is_ready

	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _cret C.gboolean     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))

	_cret = C._gotk4_soup2_Auth_virtual_is_ready(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Update updates auth with the information from msg and auth_header, possibly
// un-authenticating it. As with soup_auth_new(), this is normally only used by
// Session.
//
// The function takes the following parameters:
//
//   - msg auth is being updated for.
//   - authHeader: WWW-Authenticate/Proxy-Authenticate header.
//
// The function returns the following values:
//
//   - ok: TRUE if auth is still a valid (but potentially unauthenticated) Auth.
//     FALSE if something about auth_params could not be parsed or incorporated
//     into auth at all.
//
func (auth *Auth) update(msg *Message, authHeader map[unsafe.Pointer]unsafe.Pointer) bool {
	gclass := (*C.SoupAuthClass)(coreglib.PeekParentClass(auth))
	fnarg := gclass.update

	var _arg0 *C.SoupAuth    // out
	var _arg1 *C.SoupMessage // out
	var _arg2 *C.GHashTable  // out
	var _cret C.gboolean     // in

	_arg0 = (*C.SoupAuth)(unsafe.Pointer(coreglib.InternObject(auth).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range authHeader {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg2)

	_cret = C._gotk4_soup2_Auth_virtual_update(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(auth)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(authHeader)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AuthClass: instance of this type is always passed by reference.
type AuthClass struct {
	*authClass
}

// authClass is the struct that's finalized.
type authClass struct {
	native *C.SoupAuthClass
}

func (a *AuthClass) SchemeName() string {
	valptr := &a.native.scheme_name
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

func (a *AuthClass) Strength() uint {
	valptr := &a.native.strength
	var _v uint // out
	_v = uint(*valptr)
	return _v
}
