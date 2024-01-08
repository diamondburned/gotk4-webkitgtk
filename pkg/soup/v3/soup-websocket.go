// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// GType values.
var (
	GTypeWebsocketCloseCode      = coreglib.Type(C.soup_websocket_close_code_get_type())
	GTypeWebsocketConnectionType = coreglib.Type(C.soup_websocket_connection_type_get_type())
	GTypeWebsocketDataType       = coreglib.Type(C.soup_websocket_data_type_get_type())
	GTypeWebsocketError          = coreglib.Type(C.soup_websocket_error_get_type())
	GTypeWebsocketState          = coreglib.Type(C.soup_websocket_state_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWebsocketCloseCode, F: marshalWebsocketCloseCode},
		coreglib.TypeMarshaler{T: GTypeWebsocketConnectionType, F: marshalWebsocketConnectionType},
		coreglib.TypeMarshaler{T: GTypeWebsocketDataType, F: marshalWebsocketDataType},
		coreglib.TypeMarshaler{T: GTypeWebsocketError, F: marshalWebsocketError},
		coreglib.TypeMarshaler{T: GTypeWebsocketState, F: marshalWebsocketState},
	})
}

// WebsocketCloseCode: pre-defined close codes that can be passed to
// websocketconnection.Close or received from websocketconnection.GetCloseCode.
//
// However, other codes are also allowed.
type WebsocketCloseCode C.gint

const (
	// WebsocketCloseNormal: normal, non-error close.
	WebsocketCloseNormal WebsocketCloseCode = 1000
	// WebsocketCloseGoingAway: client/server is going away.
	WebsocketCloseGoingAway WebsocketCloseCode = 1001
	// WebsocketCloseProtocolError: protocol error occurred.
	WebsocketCloseProtocolError WebsocketCloseCode = 1002
	// WebsocketCloseUnsupportedData: endpoint received data of a type that it
	// does not support.
	WebsocketCloseUnsupportedData WebsocketCloseCode = 1003
	// WebsocketCloseNoStatus: reserved value indicating that no close code was
	// present; must not be sent.
	WebsocketCloseNoStatus WebsocketCloseCode = 1005
	// WebsocketCloseAbnormal: reserved value indicating that the connection was
	// closed abnormally; must not be sent.
	WebsocketCloseAbnormal WebsocketCloseCode = 1006
	// WebsocketCloseBadData: endpoint received data that was invalid (eg,
	// non-UTF-8 data in a text message).
	WebsocketCloseBadData WebsocketCloseCode = 1007
	// WebsocketClosePolicyViolation: generic error code indicating some sort of
	// policy violation.
	WebsocketClosePolicyViolation WebsocketCloseCode = 1008
	// WebsocketCloseTooBig: endpoint received a message that is too big to
	// process.
	WebsocketCloseTooBig WebsocketCloseCode = 1009
	// WebsocketCloseNoExtension: client is closing the connection because the
	// server failed to negotiate a required extension.
	WebsocketCloseNoExtension WebsocketCloseCode = 1010
	// WebsocketCloseServerError: server is closing the connection because it
	// was unable to fulfill the request.
	WebsocketCloseServerError WebsocketCloseCode = 1011
	// WebsocketCloseTLSHandshake: reserved value indicating that the TLS
	// handshake failed; must not be sent.
	WebsocketCloseTLSHandshake WebsocketCloseCode = 1015
)

func marshalWebsocketCloseCode(p uintptr) (interface{}, error) {
	return WebsocketCloseCode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketCloseCode.
func (w WebsocketCloseCode) String() string {
	switch w {
	case WebsocketCloseNormal:
		return "Normal"
	case WebsocketCloseGoingAway:
		return "GoingAway"
	case WebsocketCloseProtocolError:
		return "ProtocolError"
	case WebsocketCloseUnsupportedData:
		return "UnsupportedData"
	case WebsocketCloseNoStatus:
		return "NoStatus"
	case WebsocketCloseAbnormal:
		return "Abnormal"
	case WebsocketCloseBadData:
		return "BadData"
	case WebsocketClosePolicyViolation:
		return "PolicyViolation"
	case WebsocketCloseTooBig:
		return "TooBig"
	case WebsocketCloseNoExtension:
		return "NoExtension"
	case WebsocketCloseServerError:
		return "ServerError"
	case WebsocketCloseTLSHandshake:
		return "TLSHandshake"
	default:
		return fmt.Sprintf("WebsocketCloseCode(%d)", w)
	}
}

// WebsocketConnectionType: type of a websocketconnection.
type WebsocketConnectionType C.gint

const (
	// WebsocketConnectionUnknown: unknown/invalid connection.
	WebsocketConnectionUnknown WebsocketConnectionType = iota
	// WebsocketConnectionClient: client-side connection.
	WebsocketConnectionClient
	// WebsocketConnectionServer: server-side connection.
	WebsocketConnectionServer
)

func marshalWebsocketConnectionType(p uintptr) (interface{}, error) {
	return WebsocketConnectionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketConnectionType.
func (w WebsocketConnectionType) String() string {
	switch w {
	case WebsocketConnectionUnknown:
		return "Unknown"
	case WebsocketConnectionClient:
		return "Client"
	case WebsocketConnectionServer:
		return "Server"
	default:
		return fmt.Sprintf("WebsocketConnectionType(%d)", w)
	}
}

// WebsocketDataType: type of data contained in a websocketconnection::message
// signal.
type WebsocketDataType C.gint

const (
	// WebsocketDataText: UTF-8 text.
	WebsocketDataText WebsocketDataType = 1
	// WebsocketDataBinary: binary data.
	WebsocketDataBinary WebsocketDataType = 2
)

func marshalWebsocketDataType(p uintptr) (interface{}, error) {
	return WebsocketDataType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketDataType.
func (w WebsocketDataType) String() string {
	switch w {
	case WebsocketDataText:
		return "Text"
	case WebsocketDataBinary:
		return "Binary"
	default:
		return fmt.Sprintf("WebsocketDataType(%d)", w)
	}
}

// WebsocketError: webSocket-related errors.
type WebsocketError C.gint

const (
	// WebsocketErrorFailed: generic error.
	WebsocketErrorFailed WebsocketError = iota
	// WebsocketErrorNotWebsocket: attempted to handshake with a server that
	// does not appear to understand WebSockets.
	WebsocketErrorNotWebsocket
	// WebsocketErrorBadHandshake: webSocket handshake failed because some
	// detail was invalid (eg, incorrect accept key).
	WebsocketErrorBadHandshake
	// WebsocketErrorBadOrigin: webSocket handshake failed because the "Origin"
	// header was not an allowed value.
	WebsocketErrorBadOrigin
)

func marshalWebsocketError(p uintptr) (interface{}, error) {
	return WebsocketError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketError.
func (w WebsocketError) String() string {
	switch w {
	case WebsocketErrorFailed:
		return "Failed"
	case WebsocketErrorNotWebsocket:
		return "NotWebsocket"
	case WebsocketErrorBadHandshake:
		return "BadHandshake"
	case WebsocketErrorBadOrigin:
		return "BadOrigin"
	default:
		return fmt.Sprintf("WebsocketError(%d)", w)
	}
}

// WebsocketErrorQuark registers error quark for SoupWebsocket if needed.
//
// The function returns the following values:
//
//   - quark: error quark for SoupWebsocket.
//
func WebsocketErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.soup_websocket_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

// WebsocketState: state of the WebSocket connection.
type WebsocketState C.gint

const (
	// WebsocketStateOpen: connection is ready to send messages.
	WebsocketStateOpen WebsocketState = 1
	// WebsocketStateClosing: connection is in the process of closing down;
	// messages may be received, but not sent.
	WebsocketStateClosing WebsocketState = 2
	// WebsocketStateClosed: connection is completely closed down.
	WebsocketStateClosed WebsocketState = 3
)

func marshalWebsocketState(p uintptr) (interface{}, error) {
	return WebsocketState(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebsocketState.
func (w WebsocketState) String() string {
	switch w {
	case WebsocketStateOpen:
		return "Open"
	case WebsocketStateClosing:
		return "Closing"
	case WebsocketStateClosed:
		return "Closed"
	default:
		return fmt.Sprintf("WebsocketState(%d)", w)
	}
}
