// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	_ "runtime/cgo"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// GType values.
var (
	GTypeCacheResponse   = coreglib.Type(C.soup_cache_response_get_type())
	GTypeConnectionState = coreglib.Type(C.soup_connection_state_get_type())
	GTypeKnownStatusCode = coreglib.Type(C.soup_known_status_code_get_type())
	GTypeRequesterError  = coreglib.Type(C.soup_requester_error_get_type())
	GTypeSameSitePolicy  = coreglib.Type(C.soup_same_site_policy_get_type())
	GTypeXMLRPCError     = coreglib.Type(C.soup_xmlrpc_error_get_type())
	GTypeCacheability    = coreglib.Type(C.soup_cacheability_get_type())
	GTypeAuthBasic       = coreglib.Type(C.soup_auth_basic_get_type())
	GTypeAuthDigest      = coreglib.Type(C.soup_auth_digest_get_type())
	GTypeAuthNTLM        = coreglib.Type(C.soup_auth_ntlm_get_type())
	GTypeAuthNegotiate   = coreglib.Type(C.soup_auth_negotiate_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCacheResponse, F: marshalCacheResponse},
		coreglib.TypeMarshaler{T: GTypeConnectionState, F: marshalConnectionState},
		coreglib.TypeMarshaler{T: GTypeKnownStatusCode, F: marshalKnownStatusCode},
		coreglib.TypeMarshaler{T: GTypeRequesterError, F: marshalRequesterError},
		coreglib.TypeMarshaler{T: GTypeSameSitePolicy, F: marshalSameSitePolicy},
		coreglib.TypeMarshaler{T: GTypeXMLRPCError, F: marshalXMLRPCError},
		coreglib.TypeMarshaler{T: GTypeCacheability, F: marshalCacheability},
		coreglib.TypeMarshaler{T: GTypeAuthBasic, F: marshalAuthBasic},
		coreglib.TypeMarshaler{T: GTypeAuthDigest, F: marshalAuthDigest},
		coreglib.TypeMarshaler{T: GTypeAuthNTLM, F: marshalAuthNTLM},
		coreglib.TypeMarshaler{T: GTypeAuthNegotiate, F: marshalAuthNegotiate},
	})
}

type CacheResponse C.gint

const (
	CacheResponseFresh CacheResponse = iota
	CacheResponseNeedsValidation
	CacheResponseStale
)

func marshalCacheResponse(p uintptr) (interface{}, error) {
	return CacheResponse(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CacheResponse.
func (c CacheResponse) String() string {
	switch c {
	case CacheResponseFresh:
		return "Fresh"
	case CacheResponseNeedsValidation:
		return "NeedsValidation"
	case CacheResponseStale:
		return "Stale"
	default:
		return fmt.Sprintf("CacheResponse(%d)", c)
	}
}

type ConnectionState C.gint

const (
	NewConnection ConnectionState = iota
	ConnectionConnecting
	ConnectionIdle
	ConnectionInUse
	ConnectionRemoteDisconnected
	ConnectionDisconnected
)

func marshalConnectionState(p uintptr) (interface{}, error) {
	return ConnectionState(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ConnectionState.
func (c ConnectionState) String() string {
	switch c {
	case NewConnection:
		return "New"
	case ConnectionConnecting:
		return "Connecting"
	case ConnectionIdle:
		return "Idle"
	case ConnectionInUse:
		return "InUse"
	case ConnectionRemoteDisconnected:
		return "RemoteDisconnected"
	case ConnectionDisconnected:
		return "Disconnected"
	default:
		return fmt.Sprintf("ConnectionState(%d)", c)
	}
}

type KnownStatusCode C.gint

const (
	KnownStatusCodeNone                         KnownStatusCode = 0
	KnownStatusCodeCancelled                    KnownStatusCode = 1
	KnownStatusCodeCantResolve                  KnownStatusCode = 2
	KnownStatusCodeCantResolveProxy             KnownStatusCode = 3
	KnownStatusCodeCantConnect                  KnownStatusCode = 4
	KnownStatusCodeCantConnectProxy             KnownStatusCode = 5
	KnownStatusCodeSSLFailed                    KnownStatusCode = 6
	KnownStatusCodeIOError                      KnownStatusCode = 7
	KnownStatusCodeMalformed                    KnownStatusCode = 8
	KnownStatusCodeTryAgain                     KnownStatusCode = 9
	KnownStatusCodeTooManyRedirects             KnownStatusCode = 10
	KnownStatusCodeTLSFailed                    KnownStatusCode = 11
	KnownStatusCodeContinue                     KnownStatusCode = 100
	KnownStatusCodeSwitchingProtocols           KnownStatusCode = 101
	KnownStatusCodeProcessing                   KnownStatusCode = 102
	KnownStatusCodeOK                           KnownStatusCode = 200
	KnownStatusCodeCreated                      KnownStatusCode = 201
	KnownStatusCodeAccepted                     KnownStatusCode = 202
	KnownStatusCodeNonAuthoritative             KnownStatusCode = 203
	KnownStatusCodeNoContent                    KnownStatusCode = 204
	KnownStatusCodeResetContent                 KnownStatusCode = 205
	KnownStatusCodePartialContent               KnownStatusCode = 206
	KnownStatusCodeMultiStatus                  KnownStatusCode = 207
	KnownStatusCodeMultipleChoices              KnownStatusCode = 300
	KnownStatusCodeMovedPermanently             KnownStatusCode = 301
	KnownStatusCodeFound                        KnownStatusCode = 302
	KnownStatusCodeMovedTemporarily             KnownStatusCode = 302
	KnownStatusCodeSeeOther                     KnownStatusCode = 303
	KnownStatusCodeNotModified                  KnownStatusCode = 304
	KnownStatusCodeUseProxy                     KnownStatusCode = 305
	KnownStatusCodeNotAppearingInThisProtocol   KnownStatusCode = 306
	KnownStatusCodeTemporaryRedirect            KnownStatusCode = 307
	KnownStatusCodeBadRequest                   KnownStatusCode = 400
	KnownStatusCodeUnauthorized                 KnownStatusCode = 401
	KnownStatusCodePaymentRequired              KnownStatusCode = 402
	KnownStatusCodeForbidden                    KnownStatusCode = 403
	KnownStatusCodeNotFound                     KnownStatusCode = 404
	KnownStatusCodeMethodNotAllowed             KnownStatusCode = 405
	KnownStatusCodeNotAcceptable                KnownStatusCode = 406
	KnownStatusCodeProxyAuthenticationRequired  KnownStatusCode = 407
	KnownStatusCodeProxyUnauthorized            KnownStatusCode = 407
	KnownStatusCodeRequestTimeout               KnownStatusCode = 408
	KnownStatusCodeConflict                     KnownStatusCode = 409
	KnownStatusCodeGone                         KnownStatusCode = 410
	KnownStatusCodeLengthRequired               KnownStatusCode = 411
	KnownStatusCodePreconditionFailed           KnownStatusCode = 412
	KnownStatusCodeRequestEntityTooLarge        KnownStatusCode = 413
	KnownStatusCodeRequestURITooLong            KnownStatusCode = 414
	KnownStatusCodeUnsupportedMediaType         KnownStatusCode = 415
	KnownStatusCodeRequestedRangeNotSatisfiable KnownStatusCode = 416
	KnownStatusCodeInvalidRange                 KnownStatusCode = 416
	KnownStatusCodeExpectationFailed            KnownStatusCode = 417
	KnownStatusCodeUnprocessableEntity          KnownStatusCode = 422
	KnownStatusCodeLocked                       KnownStatusCode = 423
	KnownStatusCodeFailedDependency             KnownStatusCode = 424
	KnownStatusCodeInternalServerError          KnownStatusCode = 500
	KnownStatusCodeNotImplemented               KnownStatusCode = 501
	KnownStatusCodeBadGateway                   KnownStatusCode = 502
	KnownStatusCodeServiceUnavailable           KnownStatusCode = 503
	KnownStatusCodeGatewayTimeout               KnownStatusCode = 504
	KnownStatusCodeHTTPVersionNotSupported      KnownStatusCode = 505
	KnownStatusCodeInsufficientStorage          KnownStatusCode = 507
	KnownStatusCodeNotExtended                  KnownStatusCode = 510
)

func marshalKnownStatusCode(p uintptr) (interface{}, error) {
	return KnownStatusCode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for KnownStatusCode.
func (k KnownStatusCode) String() string {
	switch k {
	case KnownStatusCodeNone:
		return "None"
	case KnownStatusCodeCancelled:
		return "Cancelled"
	case KnownStatusCodeCantResolve:
		return "CantResolve"
	case KnownStatusCodeCantResolveProxy:
		return "CantResolveProxy"
	case KnownStatusCodeCantConnect:
		return "CantConnect"
	case KnownStatusCodeCantConnectProxy:
		return "CantConnectProxy"
	case KnownStatusCodeSSLFailed:
		return "SSLFailed"
	case KnownStatusCodeIOError:
		return "IOError"
	case KnownStatusCodeMalformed:
		return "Malformed"
	case KnownStatusCodeTryAgain:
		return "TryAgain"
	case KnownStatusCodeTooManyRedirects:
		return "TooManyRedirects"
	case KnownStatusCodeTLSFailed:
		return "TLSFailed"
	case KnownStatusCodeContinue:
		return "Continue"
	case KnownStatusCodeSwitchingProtocols:
		return "SwitchingProtocols"
	case KnownStatusCodeProcessing:
		return "Processing"
	case KnownStatusCodeOK:
		return "OK"
	case KnownStatusCodeCreated:
		return "Created"
	case KnownStatusCodeAccepted:
		return "Accepted"
	case KnownStatusCodeNonAuthoritative:
		return "NonAuthoritative"
	case KnownStatusCodeNoContent:
		return "NoContent"
	case KnownStatusCodeResetContent:
		return "ResetContent"
	case KnownStatusCodePartialContent:
		return "PartialContent"
	case KnownStatusCodeMultiStatus:
		return "MultiStatus"
	case KnownStatusCodeMultipleChoices:
		return "MultipleChoices"
	case KnownStatusCodeMovedPermanently:
		return "MovedPermanently"
	case KnownStatusCodeFound:
		return "Found"
	case KnownStatusCodeSeeOther:
		return "SeeOther"
	case KnownStatusCodeNotModified:
		return "NotModified"
	case KnownStatusCodeUseProxy:
		return "UseProxy"
	case KnownStatusCodeNotAppearingInThisProtocol:
		return "NotAppearingInThisProtocol"
	case KnownStatusCodeTemporaryRedirect:
		return "TemporaryRedirect"
	case KnownStatusCodeBadRequest:
		return "BadRequest"
	case KnownStatusCodeUnauthorized:
		return "Unauthorized"
	case KnownStatusCodePaymentRequired:
		return "PaymentRequired"
	case KnownStatusCodeForbidden:
		return "Forbidden"
	case KnownStatusCodeNotFound:
		return "NotFound"
	case KnownStatusCodeMethodNotAllowed:
		return "MethodNotAllowed"
	case KnownStatusCodeNotAcceptable:
		return "NotAcceptable"
	case KnownStatusCodeProxyAuthenticationRequired:
		return "ProxyAuthenticationRequired"
	case KnownStatusCodeRequestTimeout:
		return "RequestTimeout"
	case KnownStatusCodeConflict:
		return "Conflict"
	case KnownStatusCodeGone:
		return "Gone"
	case KnownStatusCodeLengthRequired:
		return "LengthRequired"
	case KnownStatusCodePreconditionFailed:
		return "PreconditionFailed"
	case KnownStatusCodeRequestEntityTooLarge:
		return "RequestEntityTooLarge"
	case KnownStatusCodeRequestURITooLong:
		return "RequestURITooLong"
	case KnownStatusCodeUnsupportedMediaType:
		return "UnsupportedMediaType"
	case KnownStatusCodeRequestedRangeNotSatisfiable:
		return "RequestedRangeNotSatisfiable"
	case KnownStatusCodeExpectationFailed:
		return "ExpectationFailed"
	case KnownStatusCodeUnprocessableEntity:
		return "UnprocessableEntity"
	case KnownStatusCodeLocked:
		return "Locked"
	case KnownStatusCodeFailedDependency:
		return "FailedDependency"
	case KnownStatusCodeInternalServerError:
		return "InternalServerError"
	case KnownStatusCodeNotImplemented:
		return "NotImplemented"
	case KnownStatusCodeBadGateway:
		return "BadGateway"
	case KnownStatusCodeServiceUnavailable:
		return "ServiceUnavailable"
	case KnownStatusCodeGatewayTimeout:
		return "GatewayTimeout"
	case KnownStatusCodeHTTPVersionNotSupported:
		return "HTTPVersionNotSupported"
	case KnownStatusCodeInsufficientStorage:
		return "InsufficientStorage"
	case KnownStatusCodeNotExtended:
		return "NotExtended"
	default:
		return fmt.Sprintf("KnownStatusCode(%d)", k)
	}
}

// The function returns the following values:
//
func RequestErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.soup_request_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type RequesterError C.gint

const (
	RequesterErrorBadURI RequesterError = iota
	RequesterErrorUnsupportedURIScheme
)

func marshalRequesterError(p uintptr) (interface{}, error) {
	return RequesterError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RequesterError.
func (r RequesterError) String() string {
	switch r {
	case RequesterErrorBadURI:
		return "BadURI"
	case RequesterErrorUnsupportedURIScheme:
		return "UnsupportedURIScheme"
	default:
		return fmt.Sprintf("RequesterError(%d)", r)
	}
}

type SameSitePolicy C.gint

const (
	// SameSitePolicyNone: cookie is exposed with both cross-site and same-site
	// requests.
	SameSitePolicyNone SameSitePolicy = iota
	// SameSitePolicyLax: cookie is withheld on cross-site requests but exposed
	// on cross-site navigations.
	SameSitePolicyLax
	// SameSitePolicyStrict: cookie is only exposed for same-site requests.
	SameSitePolicyStrict
)

func marshalSameSitePolicy(p uintptr) (interface{}, error) {
	return SameSitePolicy(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SameSitePolicy.
func (s SameSitePolicy) String() string {
	switch s {
	case SameSitePolicyNone:
		return "None"
	case SameSitePolicyLax:
		return "Lax"
	case SameSitePolicyStrict:
		return "Strict"
	default:
		return fmt.Sprintf("SameSitePolicy(%d)", s)
	}
}

// The function returns the following values:
//
func TLDErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.soup_tld_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type XMLRPCError C.gint

const (
	XmlrpcErrorArguments XMLRPCError = iota
	XmlrpcErrorRetval
)

func marshalXMLRPCError(p uintptr) (interface{}, error) {
	return XMLRPCError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for XMLRPCError.
func (x XMLRPCError) String() string {
	switch x {
	case XmlrpcErrorArguments:
		return "Arguments"
	case XmlrpcErrorRetval:
		return "Retval"
	default:
		return fmt.Sprintf("XMLRPCError(%d)", x)
	}
}

// The function returns the following values:
//
func XMLRPCErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.soup_xmlrpc_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type Cacheability C.guint

const (
	CacheCacheable   Cacheability = 0b1
	CacheUncacheable Cacheability = 0b10
	CacheInvalidates Cacheability = 0b100
	CacheValidates   Cacheability = 0b1000
)

func marshalCacheability(p uintptr) (interface{}, error) {
	return Cacheability(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for Cacheability.
func (c Cacheability) String() string {
	if c == 0 {
		return "Cacheability(0)"
	}

	var builder strings.Builder
	builder.Grow(63)

	for c != 0 {
		next := c & (c - 1)
		bit := c - next

		switch bit {
		case CacheCacheable:
			builder.WriteString("Cacheable|")
		case CacheUncacheable:
			builder.WriteString("Uncacheable|")
		case CacheInvalidates:
			builder.WriteString("Invalidates|")
		case CacheValidates:
			builder.WriteString("Validates|")
		default:
			builder.WriteString(fmt.Sprintf("Cacheability(0b%b)|", bit))
		}

		c = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if c contains other.
func (c Cacheability) Has(other Cacheability) bool {
	return (c & other) == other
}

// The function returns the following values:
//
func HTTPErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.soup_http_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type AuthBasic struct {
	_ [0]func() // equal guard
	Auth
}

var (
	_ Auther = (*AuthBasic)(nil)
)

func wrapAuthBasic(obj *coreglib.Object) *AuthBasic {
	return &AuthBasic{
		Auth: Auth{
			Object: obj,
		},
	}
}

func marshalAuthBasic(p uintptr) (interface{}, error) {
	return wrapAuthBasic(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

type AuthDigest struct {
	_ [0]func() // equal guard
	Auth
}

var (
	_ Auther = (*AuthDigest)(nil)
)

func wrapAuthDigest(obj *coreglib.Object) *AuthDigest {
	return &AuthDigest{
		Auth: Auth{
			Object: obj,
		},
	}
}

func marshalAuthDigest(p uintptr) (interface{}, error) {
	return wrapAuthDigest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

type AuthNTLM struct {
	_ [0]func() // equal guard
	Auth
}

var (
	_ Auther = (*AuthNTLM)(nil)
)

func wrapAuthNTLM(obj *coreglib.Object) *AuthNTLM {
	return &AuthNTLM{
		Auth: Auth{
			Object: obj,
		},
	}
}

func marshalAuthNTLM(p uintptr) (interface{}, error) {
	return wrapAuthNTLM(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

type AuthNegotiate struct {
	_ [0]func() // equal guard
	Auth
}

var (
	_ Auther = (*AuthNegotiate)(nil)
)

func wrapAuthNegotiate(obj *coreglib.Object) *AuthNegotiate {
	return &AuthNegotiate{
		Auth: Auth{
			Object: obj,
		},
	}
}

func marshalAuthNegotiate(p uintptr) (interface{}, error) {
	return wrapAuthNegotiate(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
