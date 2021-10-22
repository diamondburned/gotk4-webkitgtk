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
		{T: externglib.Type(C.webkit_authentication_scheme_get_type()), F: marshalAuthenticationScheme},
		{T: externglib.Type(C.webkit_authentication_request_get_type()), F: marshalAuthenticationRequester},
	})
}

// AuthenticationScheme: enum values representing the authentication scheme.
type AuthenticationScheme int

const (
	// AuthenticationSchemeDefault: default authentication scheme of WebKit.
	AuthenticationSchemeDefault AuthenticationScheme = 1
	// AuthenticationSchemeHTTPBasic: basic authentication scheme as defined in
	// RFC 2617.
	AuthenticationSchemeHTTPBasic AuthenticationScheme = 2
	// AuthenticationSchemeHTTPDigest: digest authentication scheme as defined
	// in RFC 2617.
	AuthenticationSchemeHTTPDigest AuthenticationScheme = 3
	// AuthenticationSchemeHtmlForm: HTML Form authentication.
	AuthenticationSchemeHtmlForm AuthenticationScheme = 4
	// AuthenticationSchemeNtlm: NTLM Microsoft proprietary authentication
	// scheme.
	AuthenticationSchemeNtlm AuthenticationScheme = 5
	// AuthenticationSchemeNegotiate: negotiate (or SPNEGO) authentication
	// scheme as defined in RFC 4559.
	AuthenticationSchemeNegotiate AuthenticationScheme = 6
	// AuthenticationSchemeClientCertificateRequested: client Certificate
	// Authentication (see RFC 2246).
	AuthenticationSchemeClientCertificateRequested AuthenticationScheme = 7
	// AuthenticationSchemeServerTrustEvaluationRequested: server Trust
	// Authentication.
	AuthenticationSchemeServerTrustEvaluationRequested AuthenticationScheme = 8
	// AuthenticationSchemeUnknown: authentication scheme unknown.
	AuthenticationSchemeUnknown AuthenticationScheme = 100
)

func marshalAuthenticationScheme(p uintptr) (interface{}, error) {
	return AuthenticationScheme(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AuthenticationScheme.
func (a AuthenticationScheme) String() string {
	switch a {
	case AuthenticationSchemeDefault:
		return "Default"
	case AuthenticationSchemeHTTPBasic:
		return "HTTPBasic"
	case AuthenticationSchemeHTTPDigest:
		return "HTTPDigest"
	case AuthenticationSchemeHtmlForm:
		return "HtmlForm"
	case AuthenticationSchemeNtlm:
		return "Ntlm"
	case AuthenticationSchemeNegotiate:
		return "Negotiate"
	case AuthenticationSchemeClientCertificateRequested:
		return "ClientCertificateRequested"
	case AuthenticationSchemeServerTrustEvaluationRequested:
		return "ServerTrustEvaluationRequested"
	case AuthenticationSchemeUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("AuthenticationScheme(%d)", a)
	}
}

type AuthenticationRequest struct {
	*externglib.Object
}

func wrapAuthenticationRequest(obj *externglib.Object) *AuthenticationRequest {
	return &AuthenticationRequest{
		Object: obj,
	}
}

func marshalAuthenticationRequester(p uintptr) (interface{}, error) {
	return wrapAuthenticationRequest(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Authenticate the KitAuthenticationRequest using the KitCredential supplied.
// To continue without credentials, pass NULL as credential.
//
// The function takes the following parameters:
//
//    - credential or NULL.
//
func (request *AuthenticationRequest) Authenticate(credential *Credential) {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _arg1 *C.WebKitCredential            // out

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))
	if credential != nil {
		_arg1 = (*C.WebKitCredential)(gextras.StructNative(unsafe.Pointer(credential)))
	}

	C.webkit_authentication_request_authenticate(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(credential)
}

// CanSaveCredentials: determine whether the authentication method associated
// with this KitAuthenticationRequest should allow the storage of credentials.
// This will return FALSE if WebKit doesn't support credential storing, if
// private browsing is enabled, or if persistent credential storage has been
// disabled in KitWebsiteDataManager, unless credentials saving has been
// explicitly enabled with
// webkit_authentication_request_set_can_save_credentials().
func (request *AuthenticationRequest) CanSaveCredentials() bool {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_can_save_credentials(_arg0)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Cancel the authentication challenge. This will also cancel the page loading
// and result in a KitWebView::load-failed signal with a KitNetworkError of type
// WEBKIT_NETWORK_ERROR_CANCELLED being emitted.
func (request *AuthenticationRequest) Cancel() {
	var _arg0 *C.WebKitAuthenticationRequest // out

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	C.webkit_authentication_request_cancel(_arg0)
	runtime.KeepAlive(request)
}

// Host: get the host that this authentication challenge is applicable to.
func (request *AuthenticationRequest) Host() string {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_get_host(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Port: get the port that this authentication challenge is applicable to.
func (request *AuthenticationRequest) Port() uint {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret C.guint                        // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_get_port(_arg0)
	runtime.KeepAlive(request)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ProposedCredential: get the KitCredential of the proposed authentication
// challenge that was stored from a previous session. The client can use this
// directly for authentication or construct their own KitCredential.
func (request *AuthenticationRequest) ProposedCredential() *Credential {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret *C.WebKitCredential            // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_get_proposed_credential(_arg0)
	runtime.KeepAlive(request)

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

// Realm: get the realm that this authentication challenge is applicable to.
func (request *AuthenticationRequest) Realm() string {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret *C.gchar                       // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_get_realm(_arg0)
	runtime.KeepAlive(request)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Scheme: get the authentication scheme of the authentication challenge.
func (request *AuthenticationRequest) Scheme() AuthenticationScheme {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret C.WebKitAuthenticationScheme   // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_get_scheme(_arg0)
	runtime.KeepAlive(request)

	var _authenticationScheme AuthenticationScheme // out

	_authenticationScheme = AuthenticationScheme(_cret)

	return _authenticationScheme
}

// SecurityOrigin: get the KitSecurityOrigin that this authentication challenge
// is applicable to.
func (request *AuthenticationRequest) SecurityOrigin() *SecurityOrigin {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret *C.WebKitSecurityOrigin        // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_get_security_origin(_arg0)
	runtime.KeepAlive(request)

	var _securityOrigin *SecurityOrigin // out

	_securityOrigin = (*SecurityOrigin)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_securityOrigin)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_security_origin_unref((*C.WebKitSecurityOrigin)(intern.C))
		},
	)

	return _securityOrigin
}

// IsForProxy: determine whether the authentication challenge is associated with
// a proxy server rather than an "origin" server.
func (request *AuthenticationRequest) IsForProxy() bool {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_is_for_proxy(_arg0)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRetry: determine whether this this is a first attempt or a retry for this
// authentication challenge.
func (request *AuthenticationRequest) IsRetry() bool {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _cret C.gboolean                     // in

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))

	_cret = C.webkit_authentication_request_is_retry(_arg0)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCanSaveCredentials: set whether the authentication method associated with
// request should allow the storage of credentials. This should be used by
// applications handling their own credentials storage to indicate that it
// should be supported even when internal credential storage is disabled or
// unsupported. Note that storing of credentials will not be allowed on
// ephemeral sessions in any case.
//
// The function takes the following parameters:
//
//    - enabled: value to set.
//
func (request *AuthenticationRequest) SetCanSaveCredentials(enabled bool) {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _arg1 C.gboolean                     // out

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.webkit_authentication_request_set_can_save_credentials(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(enabled)
}

// SetProposedCredential: set the KitCredential of the proposed authentication
// challenge that was stored from a previous session. This should only be used
// by applications handling their own credential storage. (When using the
// default WebKit credential storage,
// webkit_authentication_request_get_proposed_credential() already contains
// previously-stored credentials.) Passing a NULL credential will clear the
// proposed credential.
//
// The function takes the following parameters:
//
//    - credential or NULL.
//
func (request *AuthenticationRequest) SetProposedCredential(credential *Credential) {
	var _arg0 *C.WebKitAuthenticationRequest // out
	var _arg1 *C.WebKitCredential            // out

	_arg0 = (*C.WebKitAuthenticationRequest)(unsafe.Pointer(request.Native()))
	_arg1 = (*C.WebKitCredential)(gextras.StructNative(unsafe.Pointer(credential)))

	C.webkit_authentication_request_set_proposed_credential(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(credential)
}

// ConnectAuthenticated: this signal is emitted when the user authentication
// request succeeded. Applications handling their own credential storage should
// connect to this signal to save the credentials.
func (request *AuthenticationRequest) ConnectAuthenticated(f func(credential *Credential)) externglib.SignalHandle {
	return request.Connect("authenticated", f)
}

// ConnectCancelled: this signal is emitted when the user authentication request
// is cancelled. It allows the application to dismiss its authentication dialog
// in case of page load failure for example.
func (request *AuthenticationRequest) ConnectCancelled(f func()) externglib.SignalHandle {
	return request.Connect("cancelled", f)
}
