// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_soup2_AuthDomainBasicAuthCallback(SoupAuthDomain*, SoupMessage*, char*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_auth_domain_basic_get_type()), F: marshalAuthDomainBasiccer},
	})
}

// AUTH_DOMAIN_BASIC_AUTH_CALLBACK alias for the AuthDomainBasic:auth-callback
// property. (The AuthDomainBasicAuthCallback.).
const AUTH_DOMAIN_BASIC_AUTH_CALLBACK = "auth-callback"

// AUTH_DOMAIN_BASIC_AUTH_DATA alias for the AuthDomainBasic:auth-data property.
// (The data to pass to the AuthDomainBasicAuthCallback.).
const AUTH_DOMAIN_BASIC_AUTH_DATA = "auth-data"

// AuthDomainBasicAuthCallback: callback used by AuthDomainBasic for
// authentication purposes. The application should verify that username and
// password and valid and return TRUE or FALSE.
//
// If you are maintaining your own password database (rather than using the
// password to authenticate against some other system like PAM or a remote
// server), you should make sure you know what you are doing. In particular,
// don't store cleartext passwords, or easily-computed hashes of cleartext
// passwords, even if you don't care that much about the security of your
// server, because users will frequently use the same password for multiple
// sites, and so compromising any site with a cleartext (or easily-cracked)
// password database may give attackers access to other more-interesting sites
// as well.
type AuthDomainBasicAuthCallback func(domain *AuthDomainBasic, msg *Message, username, password string) (ok bool)

//export _gotk4_soup2_AuthDomainBasicAuthCallback
func _gotk4_soup2_AuthDomainBasicAuthCallback(arg0 *C.SoupAuthDomain, arg1 *C.SoupMessage, arg2 *C.char, arg3 *C.char, arg4 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var domain *AuthDomainBasic // out
	var msg *Message            // out
	var username string         // out
	var password string         // out

	domain = wrapAuthDomainBasic(externglib.Take(unsafe.Pointer(arg0)))
	msg = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))
	username = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	password = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	fn := v.(AuthDomainBasicAuthCallback)
	ok := fn(domain, msg, username, password)

	if ok {
		cret = C.TRUE
	}

	return cret
}

type AuthDomainBasic struct {
	AuthDomain
}

func wrapAuthDomainBasic(obj *externglib.Object) *AuthDomainBasic {
	return &AuthDomainBasic{
		AuthDomain: AuthDomain{
			Object: obj,
		},
	}
}

func marshalAuthDomainBasiccer(p uintptr) (interface{}, error) {
	return wrapAuthDomainBasic(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SetAuthCallback sets the callback that domain will use to authenticate
// incoming requests. For each request containing authorization, domain will
// invoke the callback, and then either accept or reject the request based on
// callback's return value.
//
// You can also set the auth callback by setting the
// SOUP_AUTH_DOMAIN_BASIC_AUTH_CALLBACK and SOUP_AUTH_DOMAIN_BASIC_AUTH_DATA
// properties, which can also be used to set the callback at construct time.
//
// The function takes the following parameters:
//
//    - callback: callback.
//
func (domain *AuthDomainBasic) SetAuthCallback(callback AuthDomainBasicAuthCallback) {
	var _arg0 *C.SoupAuthDomain                 // out
	var _arg1 C.SoupAuthDomainBasicAuthCallback // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup2_AuthDomainBasicAuthCallback)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_auth_domain_basic_set_auth_callback(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(callback)
}
