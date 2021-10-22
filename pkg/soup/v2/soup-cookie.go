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
		{T: externglib.Type(C.soup_cookie_get_type()), F: marshalCookie},
	})
}

// COOKIE_MAX_AGE_ONE_DAY: constant corresponding to 1 day, for use with
// soup_cookie_new() and soup_cookie_set_max_age().
const COOKIE_MAX_AGE_ONE_DAY = 0

// COOKIE_MAX_AGE_ONE_HOUR: constant corresponding to 1 hour, for use with
// soup_cookie_new() and soup_cookie_set_max_age().
const COOKIE_MAX_AGE_ONE_HOUR = 3600

// COOKIE_MAX_AGE_ONE_WEEK: constant corresponding to 1 week, for use with
// soup_cookie_new() and soup_cookie_set_max_age().
const COOKIE_MAX_AGE_ONE_WEEK = 0

// COOKIE_MAX_AGE_ONE_YEAR: constant corresponding to 1 year, for use with
// soup_cookie_new() and soup_cookie_set_max_age().
const COOKIE_MAX_AGE_ONE_YEAR = 0

// CookiesFromRequest parses msg's Cookie request header and returns a List of
// Cookie<!-- -->s. As the "Cookie" header, unlike "Set-Cookie", only contains
// cookie names and values, none of the other Cookie fields will be filled in.
// (Thus, you can't generally pass a cookie returned from this method directly
// to soup_cookies_to_response().).
//
// The function takes the following parameters:
//
//    - msg containing a "Cookie" request header.
//
func CookiesFromRequest(msg *Message) []Cookie {
	var _arg1 *C.SoupMessage // out
	var _cret *C.GSList      // in

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	_cret = C.soup_cookies_from_request(_arg1)
	runtime.KeepAlive(msg)

	var _sList []Cookie // out

	_sList = make([]Cookie, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.SoupCookie)(v)
		var dst Cookie // out
		dst = *(*Cookie)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(&dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_cookie_free((*C.SoupCookie)(intern.C))
			},
		)
		_sList = append(_sList, dst)
	})

	return _sList
}

// CookiesFromResponse parses msg's Set-Cookie response headers and returns a
// List of Cookie<!-- -->s. Cookies that do not specify "path" or "domain"
// attributes will have their values defaulted from msg.
//
// The function takes the following parameters:
//
//    - msg containing a "Set-Cookie" response header.
//
func CookiesFromResponse(msg *Message) []Cookie {
	var _arg1 *C.SoupMessage // out
	var _cret *C.GSList      // in

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	_cret = C.soup_cookies_from_response(_arg1)
	runtime.KeepAlive(msg)

	var _sList []Cookie // out

	_sList = make([]Cookie, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.SoupCookie)(v)
		var dst Cookie // out
		dst = *(*Cookie)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(&dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_cookie_free((*C.SoupCookie)(intern.C))
			},
		)
		_sList = append(_sList, dst)
	})

	return _sList
}

// CookiesToCookieHeader serializes a List of Cookie into a string suitable for
// setting as the value of the "Cookie" header.
//
// The function takes the following parameters:
//
//    - cookies of Cookie.
//
func CookiesToCookieHeader(cookies []Cookie) string {
	var _arg1 *C.GSList // out
	var _cret *C.char   // in

	for i := len(cookies) - 1; i >= 0; i-- {
		src := cookies[i]
		var dst *C.SoupCookie // out
		dst = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer((&src))))
		_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_slist_free(_arg1)

	_cret = C.soup_cookies_to_cookie_header(_arg1)
	runtime.KeepAlive(cookies)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// CookiesToRequest adds the name and value of each cookie in cookies to msg's
// "Cookie" request. (If msg already has a "Cookie" request header, these
// cookies will be appended to the cookies already present. Be careful that you
// do not append the same cookies twice, eg, when requeuing a message.).
//
// The function takes the following parameters:
//
//    - cookies of Cookie.
//    - msg: Message.
//
func CookiesToRequest(cookies []Cookie, msg *Message) {
	var _arg1 *C.GSList      // out
	var _arg2 *C.SoupMessage // out

	for i := len(cookies) - 1; i >= 0; i-- {
		src := cookies[i]
		var dst *C.SoupCookie // out
		dst = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer((&src))))
		_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_slist_free(_arg1)
	_arg2 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	C.soup_cookies_to_request(_arg1, _arg2)
	runtime.KeepAlive(cookies)
	runtime.KeepAlive(msg)
}

// CookiesToResponse appends a "Set-Cookie" response header to msg for each
// cookie in cookies. (This is in addition to any other "Set-Cookie" headers msg
// may already have.).
//
// The function takes the following parameters:
//
//    - cookies of Cookie.
//    - msg: Message.
//
func CookiesToResponse(cookies []Cookie, msg *Message) {
	var _arg1 *C.GSList      // out
	var _arg2 *C.SoupMessage // out

	for i := len(cookies) - 1; i >= 0; i-- {
		src := cookies[i]
		var dst *C.SoupCookie // out
		dst = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer((&src))))
		_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_slist_free(_arg1)
	_arg2 = (*C.SoupMessage)(unsafe.Pointer(msg.Native()))

	C.soup_cookies_to_response(_arg1, _arg2)
	runtime.KeepAlive(cookies)
	runtime.KeepAlive(msg)
}

// Cookie: HTTP cookie.
//
// name and value will be set for all cookies. If the cookie is generated from a
// string that appears to have no name, then name will be the empty string.
//
// domain and path give the host or domain, and path within that host/domain, to
// restrict this cookie to. If domain starts with ".", that indicates a domain
// (which matches the string after the ".", or any hostname that has domain as a
// suffix). Otherwise, it is a hostname and must match exactly.
//
// expires will be non-NULL if the cookie uses either the original "expires"
// attribute, or the newer "max-age" attribute. If expires is NULL, it indicates
// that neither "expires" nor "max-age" was specified, and the cookie expires at
// the end of the session.
//
// If http_only is set, the cookie should not be exposed to untrusted code (eg,
// javascript), so as to minimize the danger posed by cross-site scripting
// attacks.
//
// An instance of this type is always passed by reference.
type Cookie struct {
	*cookie
}

// cookie is the struct that's finalized.
type cookie struct {
	native *C.SoupCookie
}

func marshalCookie(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Cookie{&cookie{(*C.SoupCookie)(b)}}, nil
}

// NewCookie constructs a struct Cookie.
func NewCookie(name string, value string, domain string, path string, maxAge int) *Cookie {
	var _arg1 *C.char       // out
	var _arg2 *C.char       // out
	var _arg3 *C.char       // out
	var _arg4 *C.char       // out
	var _arg5 C.int         // out
	var _cret *C.SoupCookie // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.int(maxAge)

	_cret = C.soup_cookie_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(path)
	runtime.KeepAlive(maxAge)

	var _cookie *Cookie // out

	_cookie = (*Cookie)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_cookie)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_cookie_free((*C.SoupCookie)(intern.C))
		},
	)

	return _cookie
}

// AppliesToURI tests if cookie should be sent to uri.
//
// (At the moment, this does not check that cookie's domain matches uri, because
// it assumes that the caller has already done that. But don't rely on that; it
// may change in the future.).
func (cookie *Cookie) AppliesToURI(uri *URI) bool {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.SoupURI    // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_cookie_applies_to_uri(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(uri)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Copy copies cookie.
func (cookie *Cookie) Copy() *Cookie {
	var _arg0 *C.SoupCookie // out
	var _cret *C.SoupCookie // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_copy(_arg0)
	runtime.KeepAlive(cookie)

	var _ret *Cookie // out

	_ret = (*Cookie)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_cookie_free((*C.SoupCookie)(intern.C))
		},
	)

	return _ret
}

// DomainMatches checks if the cookie's domain and host match in the sense that
// cookie should be sent when making a request to host, or that cookie should be
// accepted when receiving a response from host.
func (cookie *Cookie) DomainMatches(host string) bool {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(host)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_cookie_domain_matches(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(host)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Equal tests if cookie1 and cookie2 are equal.
//
// Note that currently, this does not check that the cookie domains match. This
// may change in the future.
func (cookie1 *Cookie) Equal(cookie2 *Cookie) bool {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.SoupCookie // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie1)))
	_arg1 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie2)))

	_cret = C.soup_cookie_equal(_arg0, _arg1)
	runtime.KeepAlive(cookie1)
	runtime.KeepAlive(cookie2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Domain gets cookie's domain.
func (cookie *Cookie) Domain() string {
	var _arg0 *C.SoupCookie // out
	var _cret *C.char       // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_domain(_arg0)
	runtime.KeepAlive(cookie)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Expires gets cookie's expiration time.
func (cookie *Cookie) Expires() *Date {
	var _arg0 *C.SoupCookie // out
	var _cret *C.SoupDate   // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_expires(_arg0)
	runtime.KeepAlive(cookie)

	var _date *Date // out

	if _cret != nil {
		_date = (*Date)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _date
}

// HTTPOnly gets cookie's HttpOnly attribute.
func (cookie *Cookie) HTTPOnly() bool {
	var _arg0 *C.SoupCookie // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_http_only(_arg0)
	runtime.KeepAlive(cookie)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name gets cookie's name.
func (cookie *Cookie) Name() string {
	var _arg0 *C.SoupCookie // out
	var _cret *C.char       // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_name(_arg0)
	runtime.KeepAlive(cookie)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Path gets cookie's path.
func (cookie *Cookie) Path() string {
	var _arg0 *C.SoupCookie // out
	var _cret *C.char       // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_path(_arg0)
	runtime.KeepAlive(cookie)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

func (cookie *Cookie) SameSitePolicy() SameSitePolicy {
	var _arg0 *C.SoupCookie        // out
	var _cret C.SoupSameSitePolicy // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_same_site_policy(_arg0)
	runtime.KeepAlive(cookie)

	var _sameSitePolicy SameSitePolicy // out

	_sameSitePolicy = SameSitePolicy(_cret)

	return _sameSitePolicy
}

// Secure gets cookie's secure attribute.
func (cookie *Cookie) Secure() bool {
	var _arg0 *C.SoupCookie // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_secure(_arg0)
	runtime.KeepAlive(cookie)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Value gets cookie's value.
func (cookie *Cookie) Value() string {
	var _arg0 *C.SoupCookie // out
	var _cret *C.char       // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_get_value(_arg0)
	runtime.KeepAlive(cookie)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetDomain sets cookie's domain to domain.
func (cookie *Cookie) SetDomain(domain string) {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.char       // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_cookie_set_domain(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(domain)
}

// SetExpires sets cookie's expiration time to expires. If expires is NULL,
// cookie will be a session cookie and will expire at the end of the client's
// session.
//
// (This sets the same property as soup_cookie_set_max_age().).
func (cookie *Cookie) SetExpires(expires *Date) {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.SoupDate   // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = (*C.SoupDate)(gextras.StructNative(unsafe.Pointer(expires)))

	C.soup_cookie_set_expires(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(expires)
}

// SetHTTPOnly sets cookie's HttpOnly attribute to http_only. If TRUE, cookie
// will be marked as "http only", meaning it should not be exposed to web page
// scripts or other untrusted code.
func (cookie *Cookie) SetHTTPOnly(httpOnly bool) {
	var _arg0 *C.SoupCookie // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	if httpOnly {
		_arg1 = C.TRUE
	}

	C.soup_cookie_set_http_only(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(httpOnly)
}

// SetMaxAge sets cookie's max age to max_age. If max_age is -1, the cookie is a
// session cookie, and will expire at the end of the client's session.
// Otherwise, it is the number of seconds until the cookie expires. You can use
// the constants SOUP_COOKIE_MAX_AGE_ONE_HOUR, SOUP_COOKIE_MAX_AGE_ONE_DAY,
// SOUP_COOKIE_MAX_AGE_ONE_WEEK and SOUP_COOKIE_MAX_AGE_ONE_YEAR (or multiples
// thereof) to calculate this value. (A value of 0 indicates that the cookie
// should be considered already-expired.)
//
// (This sets the same property as soup_cookie_set_expires().).
func (cookie *Cookie) SetMaxAge(maxAge int) {
	var _arg0 *C.SoupCookie // out
	var _arg1 C.int         // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = C.int(maxAge)

	C.soup_cookie_set_max_age(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(maxAge)
}

// SetName sets cookie's name to name.
func (cookie *Cookie) SetName(name string) {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.char       // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_cookie_set_name(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(name)
}

// SetPath sets cookie's path to path.
func (cookie *Cookie) SetPath(path string) {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.char       // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_cookie_set_path(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(path)
}

// SetSameSitePolicy: when used in conjunction with
// soup_cookie_jar_get_cookie_list_with_same_site_info() this sets the policy of
// when this cookie should be exposed.
func (cookie *Cookie) SetSameSitePolicy(policy SameSitePolicy) {
	var _arg0 *C.SoupCookie        // out
	var _arg1 C.SoupSameSitePolicy // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = C.SoupSameSitePolicy(policy)

	C.soup_cookie_set_same_site_policy(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(policy)
}

// SetSecure sets cookie's secure attribute to secure. If TRUE, cookie will only
// be transmitted from the client to the server over secure (https) connections.
func (cookie *Cookie) SetSecure(secure bool) {
	var _arg0 *C.SoupCookie // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	if secure {
		_arg1 = C.TRUE
	}

	C.soup_cookie_set_secure(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(secure)
}

// SetValue sets cookie's value to value.
func (cookie *Cookie) SetValue(value string) {
	var _arg0 *C.SoupCookie // out
	var _arg1 *C.char       // out

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_cookie_set_value(_arg0, _arg1)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(value)
}

// ToCookieHeader serializes cookie in the format used by the Cookie header (ie,
// for returning a cookie from a Session to a server).
func (cookie *Cookie) ToCookieHeader() string {
	var _arg0 *C.SoupCookie // out
	var _cret *C.char       // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_to_cookie_header(_arg0)
	runtime.KeepAlive(cookie)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ToSetCookieHeader serializes cookie in the format used by the Set-Cookie
// header (ie, for sending a cookie from a Server to a client).
func (cookie *Cookie) ToSetCookieHeader() string {
	var _arg0 *C.SoupCookie // out
	var _cret *C.char       // in

	_arg0 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))

	_cret = C.soup_cookie_to_set_cookie_header(_arg0)
	runtime.KeepAlive(cookie)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// CookieParse parses header and returns a Cookie. (If header contains multiple
// cookies, only the first one will be parsed.)
//
// If header does not have "path" or "domain" attributes, they will be defaulted
// from origin. If origin is NULL, path will default to "/", but domain will be
// left as NULL. Note that this is not a valid state for a Cookie, and you will
// need to fill in some appropriate string for the domain if you want to
// actually make use of the cookie.
//
// The function takes the following parameters:
//
//    - header: cookie string (eg, the value of a Set-Cookie header).
//    - origin of the cookie, or NULL.
//
func CookieParse(header string, origin *URI) *Cookie {
	var _arg1 *C.char       // out
	var _arg2 *C.SoupURI    // out
	var _cret *C.SoupCookie // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(header)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(origin)))

	_cret = C.soup_cookie_parse(_arg1, _arg2)
	runtime.KeepAlive(header)
	runtime.KeepAlive(origin)

	var _cookie *Cookie // out

	if _cret != nil {
		_cookie = (*Cookie)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_cookie)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_cookie_free((*C.SoupCookie)(intern.C))
			},
		)
	}

	return _cookie
}
