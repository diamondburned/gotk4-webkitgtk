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
		{T: externglib.Type(C.soup_uri_get_type()), F: marshalURI},
	})
}

// URI represents a (parsed) URI. URI supports RFC 3986 (URI Generic Syntax),
// and can parse any valid URI. However, libsoup only uses "http" and "https"
// URIs internally; You can use SOUP_URI_VALID_FOR_HTTP() to test if a URI is a
// valid HTTP URI.
//
// scheme will always be set in any URI. It is an interned string and is always
// all lowercase. (If you parse a URI with a non-lowercase scheme, it will be
// converted to lowercase.) The macros SOUP_URI_SCHEME_HTTP and
// SOUP_URI_SCHEME_HTTPS provide the interned values for "http" and "https" and
// can be compared against URI scheme values.
//
// user and password are parsed as defined in the older URI specs (ie, separated
// by a colon; RFC 3986 only talks about a single "userinfo" field). Note that
// password is not included in the output of soup_uri_to_string(). libsoup does
// not normally use these fields; authentication is handled via Session signals.
//
// host contains the hostname, and port the port specified in the URI. If the
// URI doesn't contain a hostname, host will be NULL, and if it doesn't specify
// a port, port may be 0. However, for "http" and "https" URIs, host is
// guaranteed to be non-NULL (trying to parse an http URI with no host will
// return NULL), and port will always be non-0 (because libsoup knows the
// default value to use when it is not specified in the URI).
//
// path is always non-NULL. For http/https URIs, path will never be an empty
// string either; if the input URI has no path, the parsed URI will have a path
// of "/".
//
// query and fragment are optional for all URI types. soup_form_decode() may be
// useful for parsing query.
//
// Note that path, query, and fragment may contain %<!-- -->-encoded characters.
// soup_uri_new() calls soup_uri_normalize() on them, but not soup_uri_decode().
// This is necessary to ensure that soup_uri_to_string() will generate a URI
// that has exactly the same meaning as the original. (In theory, URI should
// leave user, password, and host partially-encoded as well, but this would be
// more annoying than useful.)
//
// An instance of this type is always passed by reference.
type URI struct {
	*urI
}

// urI is the struct that's finalized.
type urI struct {
	native *C.SoupURI
}

func marshalURI(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &URI{&urI{(*C.SoupURI)(unsafe.Pointer(b))}}, nil
}

// NewURI constructs a struct URI.
func NewURI(uriString string) *URI {
	var _arg1 *C.char    // out
	var _cret *C.SoupURI // in

	if uriString != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(uriString)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.soup_uri_new(_arg1)
	runtime.KeepAlive(uriString)

	var _urI *URI // out

	if _cret != nil {
		_urI = (*URI)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_urI)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_uri_free((*C.SoupURI)(intern.C))
			},
		)
	}

	return _urI
}

// NewURIWithBase constructs a struct URI.
func NewURIWithBase(base *URI, uriString string) *URI {
	var _arg1 *C.SoupURI // out
	var _arg2 *C.char    // out
	var _cret *C.SoupURI // in

	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(base)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(uriString)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.soup_uri_new_with_base(_arg1, _arg2)
	runtime.KeepAlive(base)
	runtime.KeepAlive(uriString)

	var _urI *URI // out

	_urI = (*URI)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_urI)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_uri_free((*C.SoupURI)(intern.C))
		},
	)

	return _urI
}

// Copy copies uri
func (uri *URI) Copy() *URI {
	var _arg0 *C.SoupURI // out
	var _cret *C.SoupURI // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_copy(_arg0)
	runtime.KeepAlive(uri)

	var _urI *URI // out

	_urI = (*URI)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_urI)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_uri_free((*C.SoupURI)(intern.C))
		},
	)

	return _urI
}

// CopyHost makes a copy of uri, considering only the protocol, host, and port
func (uri *URI) CopyHost() *URI {
	var _arg0 *C.SoupURI // out
	var _cret *C.SoupURI // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_copy_host(_arg0)
	runtime.KeepAlive(uri)

	var _urI *URI // out

	_urI = (*URI)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_urI)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_uri_free((*C.SoupURI)(intern.C))
		},
	)

	return _urI
}

// Equal tests whether or not uri1 and uri2 are equal in all parts
func (uri1 *URI) Equal(uri2 *URI) bool {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.SoupURI // out
	var _cret C.gboolean // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri1)))
	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri2)))

	_cret = C.soup_uri_equal(_arg0, _arg1)
	runtime.KeepAlive(uri1)
	runtime.KeepAlive(uri2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Fragment gets uri's fragment.
func (uri *URI) Fragment() string {
	var _arg0 *C.SoupURI // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_fragment(_arg0)
	runtime.KeepAlive(uri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Host gets uri's host.
func (uri *URI) Host() string {
	var _arg0 *C.SoupURI // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_host(_arg0)
	runtime.KeepAlive(uri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Password gets uri's password.
func (uri *URI) Password() string {
	var _arg0 *C.SoupURI // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_password(_arg0)
	runtime.KeepAlive(uri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Path gets uri's path.
func (uri *URI) Path() string {
	var _arg0 *C.SoupURI // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_path(_arg0)
	runtime.KeepAlive(uri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Port gets uri's port.
func (uri *URI) Port() uint32 {
	var _arg0 *C.SoupURI // out
	var _cret C.guint    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_port(_arg0)
	runtime.KeepAlive(uri)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// Query gets uri's query.
func (uri *URI) Query() string {
	var _arg0 *C.SoupURI // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_query(_arg0)
	runtime.KeepAlive(uri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Scheme gets uri's scheme.
func (uri *URI) Scheme() string {
	var _arg0 *C.SoupURI // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_scheme(_arg0)
	runtime.KeepAlive(uri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// User gets uri's user.
func (uri *URI) User() string {
	var _arg0 *C.SoupURI // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_get_user(_arg0)
	runtime.KeepAlive(uri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HostEqual compares v1 and v2, considering only the scheme, host, and port.
func (v1 *URI) HostEqual(v2 *URI) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(v1)))
	_arg1 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(v2)))

	_cret = C.soup_uri_host_equal(_arg0, _arg1)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HostHash hashes key, considering only the scheme, host, and port.
func (key *URI) HostHash() uint32 {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(key)))

	_cret = C.soup_uri_host_hash(_arg0)
	runtime.KeepAlive(key)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// SetFragment sets uri's fragment to fragment.
func (uri *URI) SetFragment(fragment string) {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.char    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	if fragment != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(fragment)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.soup_uri_set_fragment(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(fragment)
}

// SetHost sets uri's host to host.
//
// If host is an IPv6 IP address, it should not include the brackets required by
// the URI syntax; they will be added automatically when converting uri to a
// string.
//
// http and https URIs should not have a NULL host.
func (uri *URI) SetHost(host string) {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.char    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	if host != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(host)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.soup_uri_set_host(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(host)
}

// SetPassword sets uri's password to password.
func (uri *URI) SetPassword(password string) {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.char    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	if password != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(password)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.soup_uri_set_password(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(password)
}

// SetPath sets uri's path to path.
func (uri *URI) SetPath(path string) {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.char    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_uri_set_path(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(path)
}

// SetPort sets uri's port to port. If port is 0, uri will not have an
// explicitly-specified port.
func (uri *URI) SetPort(port uint32) {
	var _arg0 *C.SoupURI // out
	var _arg1 C.guint    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	_arg1 = C.guint(port)

	C.soup_uri_set_port(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(port)
}

// SetQuery sets uri's query to query.
func (uri *URI) SetQuery(query string) {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.char    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	if query != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(query)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.soup_uri_set_query(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(query)
}

// SetQueryFromForm sets uri's query to the result of encoding form according to
// the HTML form rules. See soup_form_encode_hash() for more information.
func (uri *URI) SetQueryFromForm(form map[string]string) {
	var _arg0 *C.SoupURI    // out
	var _arg1 *C.GHashTable // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range form {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		defer C.free(unsafe.Pointer(vdst))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)

	C.soup_uri_set_query_from_form(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(form)
}

// SetScheme sets uri's scheme to scheme. This will also set uri's port to the
// default port for scheme, if known.
func (uri *URI) SetScheme(scheme string) {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.char    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_uri_set_scheme(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(scheme)
}

// SetUser sets uri's user to user.
func (uri *URI) SetUser(user string) {
	var _arg0 *C.SoupURI // out
	var _arg1 *C.char    // out

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	if user != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(user)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.soup_uri_set_user(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(user)
}

// String returns a string representing uri.
//
// If just_path_and_query is TRUE, this concatenates the path and query
// together. That is, it constructs the string that would be needed in the
// Request-Line of an HTTP request for uri.
//
// Note that the output will never contain a password, even if uri does.
func (uri *URI) String(justPathAndQuery bool) string {
	var _arg0 *C.SoupURI // out
	var _arg1 C.gboolean // out
	var _cret *C.char    // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	if justPathAndQuery {
		_arg1 = C.TRUE
	}

	_cret = C.soup_uri_to_string(_arg0, _arg1)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(justPathAndQuery)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// UsesDefaultPort tests if uri uses the default port for its scheme. (Eg, 80
// for http.) (This only works for http, https and ftp; libsoup does not know
// the default ports of other protocols.)
func (uri *URI) UsesDefaultPort() bool {
	var _arg0 *C.SoupURI // out
	var _cret C.gboolean // in

	_arg0 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_uri_uses_default_port(_arg0)
	runtime.KeepAlive(uri)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// URIDecode: fully %<!-- -->-decodes part.
//
// In the past, this would return NULL if part contained invalid
// percent-encoding, but now it just ignores the problem (as soup_uri_new()
// already did).
func URIDecode(part string) string {
	var _arg1 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(part)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_uri_decode(_arg1)
	runtime.KeepAlive(part)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// URIEncode: this %<!-- -->-encodes the given URI part and returns the escaped
// version in allocated memory, which the caller must free when it is done.
func URIEncode(part string, escapeExtra string) string {
	var _arg1 *C.char // out
	var _arg2 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(part)))
	defer C.free(unsafe.Pointer(_arg1))
	if escapeExtra != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(escapeExtra)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.soup_uri_encode(_arg1, _arg2)
	runtime.KeepAlive(part)
	runtime.KeepAlive(escapeExtra)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// URINormalize: %<!-- -->-decodes any "unreserved" characters (or characters in
// unescape_extra) in part, and %<!-- -->-encodes any non-ASCII characters,
// spaces, and non-printing characters in part.
//
// "Unreserved" characters are those that are not allowed to be used for
// punctuation according to the URI spec. For example, letters are unreserved,
// so soup_uri_normalize() will turn <literal>http://example.com/foo/b%<!--
// -->61r</literal> into <literal>http://example.com/foo/bar</literal>, which is
// guaranteed to mean the same thing. However, "/" is "reserved", so
// <literal>http://example.com/foo%<!-- -->2Fbar</literal> would not be changed,
// because it might mean something different to the server.
//
// In the past, this would return NULL if part contained invalid
// percent-encoding, but now it just ignores the problem (as soup_uri_new()
// already did).
func URINormalize(part string, unescapeExtra string) string {
	var _arg1 *C.char // out
	var _arg2 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(part)))
	defer C.free(unsafe.Pointer(_arg1))
	if unescapeExtra != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(unescapeExtra)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.soup_uri_normalize(_arg1, _arg2)
	runtime.KeepAlive(part)
	runtime.KeepAlive(unescapeExtra)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
