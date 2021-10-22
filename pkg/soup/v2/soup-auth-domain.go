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
// gboolean _gotk4_soup2_AuthDomainFilter(SoupAuthDomain*, SoupMessage*, gpointer);
// gboolean _gotk4_soup2_AuthDomainGenericAuthCallback(SoupAuthDomain*, SoupMessage*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_auth_domain_get_type()), F: marshalAuthDomainer},
	})
}

// AUTH_DOMAIN_ADD_PATH alias for the AuthDomain:add-path property. (Shortcut
// for calling soup_auth_domain_add_path().).
const AUTH_DOMAIN_ADD_PATH = "add-path"

// AUTH_DOMAIN_FILTER alias for the AuthDomain:filter property. (The
// AuthDomainFilter for the domain.).
const AUTH_DOMAIN_FILTER = "filter"

// AUTH_DOMAIN_FILTER_DATA alias for the AuthDomain:filter-data property. (Data
// to pass to the AuthDomainFilter.).
const AUTH_DOMAIN_FILTER_DATA = "filter-data"

// AUTH_DOMAIN_GENERIC_AUTH_CALLBACK alias for the
// AuthDomain:generic-auth-callback property. (The
// AuthDomainGenericAuthCallback.).
const AUTH_DOMAIN_GENERIC_AUTH_CALLBACK = "generic-auth-callback"

// AUTH_DOMAIN_GENERIC_AUTH_DATA alias for the AuthDomain:generic-auth-data
// property. (The data to pass to the AuthDomainGenericAuthCallback.).
const AUTH_DOMAIN_GENERIC_AUTH_DATA = "generic-auth-data"

// AUTH_DOMAIN_PROXY alias for the AuthDomain:proxy property. (Whether or not
// this is a proxy auth domain.).
const AUTH_DOMAIN_PROXY = "proxy"

// AUTH_DOMAIN_REALM alias for the AuthDomain:realm property. (The realm of this
// auth domain.).
const AUTH_DOMAIN_REALM = "realm"

// AUTH_DOMAIN_REMOVE_PATH alias for the AuthDomain:remove-path property.
// (Shortcut for calling soup_auth_domain_remove_path().).
const AUTH_DOMAIN_REMOVE_PATH = "remove-path"

// AuthDomainFilter: prototype for a AuthDomain filter; see
// soup_auth_domain_set_filter() for details.
type AuthDomainFilter func(domain AuthDomainer, msg *Message) (ok bool)

//export _gotk4_soup2_AuthDomainFilter
func _gotk4_soup2_AuthDomainFilter(arg0 *C.SoupAuthDomain, arg1 *C.SoupMessage, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var domain AuthDomainer // out
	var msg *Message        // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type soup.AuthDomainer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(AuthDomainer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not soup.AuthDomainer")
		}
		domain = rv
	}
	msg = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))

	fn := v.(AuthDomainFilter)
	ok := fn(domain, msg)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// AuthDomainGenericAuthCallback: prototype for a AuthDomain generic
// authentication callback.
//
// The callback should look up the user's password, call
// soup_auth_domain_check_password(), and use the return value from that method
// as its own return value.
//
// In general, for security reasons, it is preferable to use the
// auth-domain-specific auth callbacks (eg, AuthDomainBasicAuthCallback and
// AuthDomainDigestAuthCallback), because they don't require keeping a cleartext
// password database. Most users will use the same password for many different
// sites, meaning if any site with a cleartext password database is compromised,
// accounts on other servers might be compromised as well. For many of the cases
// where Server is used, this is not really relevant, but it may still be worth
// considering.
type AuthDomainGenericAuthCallback func(domain AuthDomainer, msg *Message, username string) (ok bool)

//export _gotk4_soup2_AuthDomainGenericAuthCallback
func _gotk4_soup2_AuthDomainGenericAuthCallback(arg0 *C.SoupAuthDomain, arg1 *C.SoupMessage, arg2 *C.char, arg3 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var domain AuthDomainer // out
	var msg *Message        // out
	var username string     // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type soup.AuthDomainer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(AuthDomainer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not soup.AuthDomainer")
		}
		domain = rv
	}
	msg = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))
	username = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	fn := v.(AuthDomainGenericAuthCallback)
	ok := fn(domain, msg, username)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// AuthDomainOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AuthDomainOverrider interface {
	Accepts(msg *Message, header string) string
	// Challenge adds a "WWW-Authenticate" or "Proxy-Authenticate" header to
	// msg, requesting that the client authenticate, and sets msg's status
	// accordingly.
	//
	// This is used by Server internally and is probably of no use to anyone
	// else.
	Challenge(msg *Message) string
	// CheckPassword checks if msg authenticates to domain via username and
	// password. This would normally be called from a
	// AuthDomainGenericAuthCallback.
	CheckPassword(msg *Message, username, password string) bool
}

type AuthDomain struct {
	*externglib.Object
}

// AuthDomainer describes types inherited from class AuthDomain.
// To get the original type, the caller must assert this to an interface or
// another type.
type AuthDomainer interface {
	externglib.Objector

	// BaseAuthDomain returns the underlying base class.
	BaseAuthDomain() *AuthDomain
}

var _ AuthDomainer = (*AuthDomain)(nil)

func wrapAuthDomain(obj *externglib.Object) *AuthDomain {
	return &AuthDomain{
		Object: obj,
	}
}

func marshalAuthDomainer(p uintptr) (interface{}, error) {
	return wrapAuthDomain(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Accepts checks if msg contains appropriate authorization for domain to accept
// it. Mirroring soup_auth_domain_covers(), this does not check whether or not
// domain <emphasis>cares</emphasis> if msg is authorized.
//
// This is used by Server internally and is probably of no use to anyone else.
//
// The function takes the following parameters:
//
//    - msg: Message.
//
func (domain *AuthDomain) Accepts(msg *Message) string {
	var _arg0 *C.SoupAuthDomain // out
	var _arg1 *C.SoupMessage    // out
	var _cret *C.char           // in

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	_cret = C.soup_auth_domain_accepts(_arg0, _arg1)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msg)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// AddPath adds path to domain, such that requests under path on domain's server
// will require authentication (unless overridden by
// soup_auth_domain_remove_path() or soup_auth_domain_set_filter()).
//
// You can also add paths by setting the SOUP_AUTH_DOMAIN_ADD_PATH property,
// which can also be used to add one or more paths at construct time.
//
// The function takes the following parameters:
//
//    - path to add to domain.
//
func (domain *AuthDomain) AddPath(path string) {
	var _arg0 *C.SoupAuthDomain // out
	var _arg1 *C.char           // out

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_auth_domain_add_path(_arg0, _arg1)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(path)
}

// Challenge adds a "WWW-Authenticate" or "Proxy-Authenticate" header to msg,
// requesting that the client authenticate, and sets msg's status accordingly.
//
// This is used by Server internally and is probably of no use to anyone else.
//
// The function takes the following parameters:
//
//    - msg: Message.
//
func (domain *AuthDomain) Challenge(msg *Message) {
	var _arg0 *C.SoupAuthDomain // out
	var _arg1 *C.SoupMessage    // out

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	C.soup_auth_domain_challenge(_arg0, _arg1)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msg)
}

// CheckPassword checks if msg authenticates to domain via username and
// password. This would normally be called from a AuthDomainGenericAuthCallback.
//
// The function takes the following parameters:
//
//    - msg: Message.
//    - username: username.
//    - password: password.
//
func (domain *AuthDomain) CheckPassword(msg *Message, username, password string) bool {
	var _arg0 *C.SoupAuthDomain // out
	var _arg1 *C.SoupMessage    // out
	var _arg2 *C.char           // out
	var _arg3 *C.char           // out
	var _cret C.gboolean        // in

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.soup_auth_domain_check_password(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Covers checks if domain requires msg to be authenticated (according to its
// paths and filter function). This does not actually look at whether msg
// <emphasis>is</emphasis> authenticated, merely whether or not it needs to be.
//
// This is used by Server internally and is probably of no use to anyone else.
//
// The function takes the following parameters:
//
//    - msg: Message.
//
func (domain *AuthDomain) Covers(msg *Message) bool {
	var _arg0 *C.SoupAuthDomain // out
	var _arg1 *C.SoupMessage    // out
	var _cret C.gboolean        // in

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	_cret = C.soup_auth_domain_covers(_arg0, _arg1)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msg)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Realm gets the realm name associated with domain.
func (domain *AuthDomain) Realm() string {
	var _arg0 *C.SoupAuthDomain // out
	var _cret *C.char           // in

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))

	_cret = C.soup_auth_domain_get_realm(_arg0)
	runtime.KeepAlive(domain)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// RemovePath removes path from domain, such that requests under path on
// domain's server will NOT require authentication.
//
// This is not simply an undo-er for soup_auth_domain_add_path(); it can be used
// to "carve out" a subtree that does not require authentication inside a
// hierarchy that does. Note also that unlike with soup_auth_domain_add_path(),
// this cannot be overridden by adding a filter, as filters can only bypass
// authentication that would otherwise be required, not require it where it
// would otherwise be unnecessary.
//
// You can also remove paths by setting the SOUP_AUTH_DOMAIN_REMOVE_PATH
// property, which can also be used to remove one or more paths at construct
// time.
//
// The function takes the following parameters:
//
//    - path to remove from domain.
//
func (domain *AuthDomain) RemovePath(path string) {
	var _arg0 *C.SoupAuthDomain // out
	var _arg1 *C.char           // out

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_auth_domain_remove_path(_arg0, _arg1)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(path)
}

// SetFilter adds filter as an authentication filter to domain. The filter gets
// a chance to bypass authentication for certain requests that would otherwise
// require it. Eg, it might check the message's path in some way that is too
// complicated to do via the other methods, or it might check the message's
// method, and allow GETs but not PUTs.
//
// The filter function returns TRUE if the request should still require
// authentication, or FALSE if authentication is unnecessary for this request.
//
// To help prevent security holes, your filter should return TRUE by default,
// and only return FALSE under specifically-tested circumstances, rather than
// the other way around. Eg, in the example above, where you want to
// authenticate PUTs but not GETs, you should check if the method is GET and
// return FALSE in that case, and then return TRUE for all other methods (rather
// than returning TRUE for PUT and FALSE for all other methods). This way if it
// turned out (now or later) that some paths supported additional methods
// besides GET and PUT, those methods would default to being NOT allowed for
// unauthenticated users.
//
// You can also set the filter by setting the SOUP_AUTH_DOMAIN_FILTER and
// SOUP_AUTH_DOMAIN_FILTER_DATA properties, which can also be used to set the
// filter at construct time.
//
// The function takes the following parameters:
//
//    - filter: auth filter for domain.
//
func (domain *AuthDomain) SetFilter(filter AuthDomainFilter) {
	var _arg0 *C.SoupAuthDomain      // out
	var _arg1 C.SoupAuthDomainFilter // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup2_AuthDomainFilter)
	_arg2 = C.gpointer(gbox.Assign(filter))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_auth_domain_set_filter(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(filter)
}

// SetGenericAuthCallback sets auth_callback as an authentication-handling
// callback for domain. Whenever a request comes in to domain which cannot be
// authenticated via a domain-specific auth callback (eg,
// AuthDomainDigestAuthCallback), the generic auth callback will be invoked. See
// AuthDomainGenericAuthCallback for information on what the callback should do.
//
// The function takes the following parameters:
//
//    - authCallback: auth callback.
//
func (domain *AuthDomain) SetGenericAuthCallback(authCallback AuthDomainGenericAuthCallback) {
	var _arg0 *C.SoupAuthDomain                   // out
	var _arg1 C.SoupAuthDomainGenericAuthCallback // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup2_AuthDomainGenericAuthCallback)
	_arg2 = C.gpointer(gbox.Assign(authCallback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_auth_domain_set_generic_auth_callback(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(authCallback)
}

//
// The function takes the following parameters:
//

//
func (domain *AuthDomain) TryGenericAuthCallback(msg *Message, username string) bool {
	var _arg0 *C.SoupAuthDomain // out
	var _arg1 *C.SoupMessage    // out
	var _arg2 *C.char           // out
	var _cret C.gboolean        // in

	_arg0 = (*C.SoupAuthDomain)(unsafe.Pointer(domain.Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.soup_auth_domain_try_generic_auth_callback(_arg0, _arg1, _arg2)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(username)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BaseAuthDomain returns domain.
func (domain *AuthDomain) BaseAuthDomain() *AuthDomain {
	return domain
}
