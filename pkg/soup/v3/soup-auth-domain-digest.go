// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern void callbackDelete(gpointer);
// extern char* _gotk4_soup3_AuthDomainDigestAuthCallback(SoupAuthDomain*, SoupServerMessage*, char*, gpointer);
import "C"

// GType values.
var (
	GTypeAuthDomainDigest = coreglib.Type(C.soup_auth_domain_digest_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAuthDomainDigest, F: marshalAuthDomainDigest},
	})
}

// AuthDomainDigestAuthCallback: callback used by AuthDomainDigest for
// authentication purposes.
//
// The application should look up username in its password database, and return
// the corresponding encoded password (see authdomaindigest.EncodePassword().
type AuthDomainDigestAuthCallback func(domain *AuthDomainDigest, msg *ServerMessage, username string) (utf8 string)

// AuthDomainDigestOverrides contains methods that are overridable.
type AuthDomainDigestOverrides struct {
}

func defaultAuthDomainDigestOverrides(v *AuthDomainDigest) AuthDomainDigestOverrides {
	return AuthDomainDigestOverrides{}
}

// AuthDomainDigest: server-side "Digest" authentication.
//
// AuthDomainDigest handles the server side of HTTP "Digest" authentication.
type AuthDomainDigest struct {
	_ [0]func() // equal guard
	AuthDomain
}

var (
	_ AuthDomainer = (*AuthDomainDigest)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*AuthDomainDigest, *AuthDomainDigestClass, AuthDomainDigestOverrides](
		GTypeAuthDomainDigest,
		initAuthDomainDigestClass,
		wrapAuthDomainDigest,
		defaultAuthDomainDigestOverrides,
	)
}

func initAuthDomainDigestClass(gclass unsafe.Pointer, overrides AuthDomainDigestOverrides, classInitFunc func(*AuthDomainDigestClass)) {
	if classInitFunc != nil {
		class := (*AuthDomainDigestClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAuthDomainDigest(obj *coreglib.Object) *AuthDomainDigest {
	return &AuthDomainDigest{
		AuthDomain: AuthDomain{
			Object: obj,
		},
	}
}

func marshalAuthDomainDigest(p uintptr) (interface{}, error) {
	return wrapAuthDomainDigest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SetAuthCallback sets the callback that domain will use to authenticate
// incoming requests.
//
// For each request containing authorization, domain will invoke the callback,
// and then either accept or reject the request based on callback's return
// value.
//
// You can also set the auth callback by setting the
// authdomaindigest:auth-callback and authdomaindigest:auth-data properties,
// which can also be used to set the callback at construct time.
//
// The function takes the following parameters:
//
//   - callback: callback.
//
func (domain *AuthDomainDigest) SetAuthCallback(callback AuthDomainDigestAuthCallback) {
	var _arg0 *C.SoupAuthDomain                  // out
	var _arg1 C.SoupAuthDomainDigestAuthCallback // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(coreglib.InternObject(domain).Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup3_AuthDomainDigestAuthCallback)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_auth_domain_digest_set_auth_callback(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(callback)
}

// AuthDomainDigestEncodePassword encodes the username/realm/password triplet
// for Digest authentication.
//
// That is, it returns a stringified MD5 hash of username, realm, and password
// concatenated together. This is the form that is needed as the return value of
// AuthDomainDigest's auth handler.
//
// For security reasons, you should store the encoded hash, rather than storing
// the cleartext password itself and calling this method only when you need to
// verify it. This way, if your server is compromised, the attackers will not
// gain access to cleartext passwords which might also be usable at other sites.
// (Note also that the encoded password returned by this method is identical to
// the encoded password stored in an Apache .htdigest file.).
//
// The function takes the following parameters:
//
//   - username: username.
//   - realm: auth realm name.
//   - password for username in realm.
//
// The function returns the following values:
//
//   - utf8: encoded password.
//
func AuthDomainDigestEncodePassword(username, realm, password string) string {
	var _arg1 *C.char // out
	var _arg2 *C.char // out
	var _arg3 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(realm)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.soup_auth_domain_digest_encode_password(_arg1, _arg2, _arg3)
	runtime.KeepAlive(username)
	runtime.KeepAlive(realm)
	runtime.KeepAlive(password)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AuthDomainDigestClass: instance of this type is always passed by reference.
type AuthDomainDigestClass struct {
	*authDomainDigestClass
}

// authDomainDigestClass is the struct that's finalized.
type authDomainDigestClass struct {
	native *C.SoupAuthDomainDigestClass
}

func (a *AuthDomainDigestClass) ParentClass() *AuthDomainClass {
	valptr := &a.native.parent_class
	var _v *AuthDomainClass // out
	_v = (*AuthDomainClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}