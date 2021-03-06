// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern gboolean _gotk4_soup2_AuthDomainBasicAuthCallback(SoupAuthDomain*, SoupMessage*, char*, char*, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for soup-auth-domain-basic.go.
var GTypeAuthDomainBasic = externglib.Type(C.soup_auth_domain_basic_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeAuthDomainBasic, F: marshalAuthDomainBasic},
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
func _gotk4_soup2_AuthDomainBasicAuthCallback(arg1 *C.SoupAuthDomain, arg2 *C.SoupMessage, arg3 *C.char, arg4 *C.char, arg5 C.gpointer) (cret C.gboolean) {
	var fn AuthDomainBasicAuthCallback
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AuthDomainBasicAuthCallback)
	}

	var _domain *AuthDomainBasic // out
	var _msg *Message            // out
	var _username string         // out
	var _password string         // out

	_domain = wrapAuthDomainBasic(externglib.Take(unsafe.Pointer(arg1)))
	_msg = wrapMessage(externglib.Take(unsafe.Pointer(arg2)))
	_username = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_password = C.GoString((*C.gchar)(unsafe.Pointer(arg4)))

	ok := fn(_domain, _msg, _username, _password)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// AuthDomainBasicOverrider contains methods that are overridable.
type AuthDomainBasicOverrider interface {
}

type AuthDomainBasic struct {
	_ [0]func() // equal guard
	AuthDomain
}

var (
	_ AuthDomainer = (*AuthDomainBasic)(nil)
)

func classInitAuthDomainBasiccer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapAuthDomainBasic(obj *externglib.Object) *AuthDomainBasic {
	return &AuthDomainBasic{
		AuthDomain: AuthDomain{
			Object: obj,
		},
	}
}

func marshalAuthDomainBasic(p uintptr) (interface{}, error) {
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

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(externglib.InternObject(domain).Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup2_AuthDomainBasicAuthCallback)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_auth_domain_basic_set_auth_callback(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(callback)
}
