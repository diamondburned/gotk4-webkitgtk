// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// GType values.
var (
	GTypeAuthenticationScheme = coreglib.Type(C.webkit_authentication_scheme_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAuthenticationScheme, F: marshalAuthenticationScheme},
	})
}

// AuthenticationScheme: enum values representing the authentication scheme.
type AuthenticationScheme C.gint

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
	// AuthenticationSchemeClientCertificatePINRequested: client certificate PIN
	// required for use. Since: 2.34.
	AuthenticationSchemeClientCertificatePINRequested AuthenticationScheme = 9
	// AuthenticationSchemeUnknown: authentication scheme unknown.
	AuthenticationSchemeUnknown AuthenticationScheme = 100
)

func marshalAuthenticationScheme(p uintptr) (interface{}, error) {
	return AuthenticationScheme(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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
	case AuthenticationSchemeClientCertificatePINRequested:
		return "ClientCertificatePINRequested"
	case AuthenticationSchemeUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("AuthenticationScheme(%d)", a)
	}
}
