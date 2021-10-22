// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_websocket_connection_get_type()), F: marshalWebsocketConnectioner},
	})
}

// WebsocketConnectionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type WebsocketConnectionOverrider interface {
	Closed()
	Closing()
	Error(err error)
	Message(typ WebsocketDataType, message *glib.Bytes)
	Pong(message *glib.Bytes)
}

// WebsocketConnection class representing a WebSocket connection.
type WebsocketConnection struct {
	*externglib.Object
}

func wrapWebsocketConnection(obj *externglib.Object) *WebsocketConnection {
	return &WebsocketConnection{
		Object: obj,
	}
}

func marshalWebsocketConnectioner(p uintptr) (interface{}, error) {
	return wrapWebsocketConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWebsocketConnection creates a WebsocketConnection on stream. This should
// be called after completing the handshake to begin using the WebSocket
// protocol.
//
// The function takes the following parameters:
//
//    - stream connected to the WebSocket server.
//    - uri: URI of the connection.
//    - typ: type of connection (client/side).
//    - origin: origin of the client.
//    - protocol in use.
//
func NewWebsocketConnection(stream gio.IOStreamer, uri *URI, typ WebsocketConnectionType, origin, protocol string) *WebsocketConnection {
	var _arg1 *C.GIOStream                  // out
	var _arg2 *C.SoupURI                    // out
	var _arg3 C.SoupWebsocketConnectionType // out
	var _arg4 *C.char                       // out
	var _arg5 *C.char                       // out
	var _cret *C.SoupWebsocketConnection    // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	_arg3 = C.SoupWebsocketConnectionType(typ)
	if origin != "" {
		_arg4 = (*C.char)(unsafe.Pointer(C.CString(origin)))
		defer C.free(unsafe.Pointer(_arg4))
	}
	if protocol != "" {
		_arg5 = (*C.char)(unsafe.Pointer(C.CString(protocol)))
		defer C.free(unsafe.Pointer(_arg5))
	}

	_cret = C.soup_websocket_connection_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(origin)
	runtime.KeepAlive(protocol)

	var _websocketConnection *WebsocketConnection // out

	_websocketConnection = wrapWebsocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _websocketConnection
}

// NewWebsocketConnectionWithExtensions creates a WebsocketConnection on stream
// with the given active extensions. This should be called after completing the
// handshake to begin using the WebSocket protocol.
//
// The function takes the following parameters:
//
//    - stream connected to the WebSocket server.
//    - uri: URI of the connection.
//    - typ: type of connection (client/side).
//    - origin: origin of the client.
//    - protocol in use.
//    - extensions of WebsocketExtension objects.
//
func NewWebsocketConnectionWithExtensions(stream gio.IOStreamer, uri *URI, typ WebsocketConnectionType, origin, protocol string, extensions []WebsocketExtensioner) *WebsocketConnection {
	var _arg1 *C.GIOStream                  // out
	var _arg2 *C.SoupURI                    // out
	var _arg3 C.SoupWebsocketConnectionType // out
	var _arg4 *C.char                       // out
	var _arg5 *C.char                       // out
	var _arg6 *C.GList                      // out
	var _cret *C.SoupWebsocketConnection    // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	_arg3 = C.SoupWebsocketConnectionType(typ)
	if origin != "" {
		_arg4 = (*C.char)(unsafe.Pointer(C.CString(origin)))
		defer C.free(unsafe.Pointer(_arg4))
	}
	if protocol != "" {
		_arg5 = (*C.char)(unsafe.Pointer(C.CString(protocol)))
		defer C.free(unsafe.Pointer(_arg5))
	}
	for i := len(extensions) - 1; i >= 0; i-- {
		src := extensions[i]
		var dst *C.SoupWebsocketExtension // out
		dst = (*C.SoupWebsocketExtension)(unsafe.Pointer(src.Native()))
		C.g_object_ref(C.gpointer(src.Native()))
		_arg6 = C.g_list_prepend(_arg6, C.gpointer(unsafe.Pointer(dst)))
	}

	_cret = C.soup_websocket_connection_new_with_extensions(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(origin)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(extensions)

	var _websocketConnection *WebsocketConnection // out

	_websocketConnection = wrapWebsocketConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _websocketConnection
}

// Close the connection in an orderly fashion.
//
// Note that until the WebsocketConnection::closed signal fires, the connection
// is not yet completely closed. The close message is not even sent until the
// main loop runs.
//
// The code and data are sent to the peer along with the close request. If code
// is SOUP_WEBSOCKET_CLOSE_NO_STATUS a close message with no body (without code
// and data) is sent. Note that the data must be UTF-8 valid.
//
// The function takes the following parameters:
//
//    - code: close code.
//    - data: close data.
//
func (self *WebsocketConnection) Close(code uint16, data string) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 C.gushort                  // out
	var _arg2 *C.char                    // out

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))
	_arg1 = C.gushort(code)
	if data != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(data)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.soup_websocket_connection_close(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(code)
	runtime.KeepAlive(data)
}

// CloseCode: get the close code received from the WebSocket peer.
//
// This only becomes valid once the WebSocket is in the
// SOUP_WEBSOCKET_STATE_CLOSED state. The value will often be in the
// WebsocketCloseCode enumeration, but may also be an application defined close
// code.
func (self *WebsocketConnection) CloseCode() uint16 {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.gushort                  // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_close_code(_arg0)
	runtime.KeepAlive(self)

	var _gushort uint16 // out

	_gushort = uint16(_cret)

	return _gushort
}

// CloseData: get the close data received from the WebSocket peer.
//
// This only becomes valid once the WebSocket is in the
// SOUP_WEBSOCKET_STATE_CLOSED state. The data may be freed once the main loop
// is run, so copy it if you need to keep it around.
func (self *WebsocketConnection) CloseData() string {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.char                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_close_data(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ConnectionType: get the connection type (client/server) of the connection.
func (self *WebsocketConnection) ConnectionType() WebsocketConnectionType {
	var _arg0 *C.SoupWebsocketConnection    // out
	var _cret C.SoupWebsocketConnectionType // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_connection_type(_arg0)
	runtime.KeepAlive(self)

	var _websocketConnectionType WebsocketConnectionType // out

	_websocketConnectionType = WebsocketConnectionType(_cret)

	return _websocketConnectionType
}

// Extensions: get the extensions chosen via negotiation with the peer.
func (self *WebsocketConnection) Extensions() []WebsocketExtensioner {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.GList                   // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_extensions(_arg0)
	runtime.KeepAlive(self)

	var _list []WebsocketExtensioner // out

	_list = make([]WebsocketExtensioner, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.SoupWebsocketExtension)(v)
		var dst WebsocketExtensioner // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type soup.WebsocketExtensioner is nil")
			}

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(WebsocketExtensioner)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not soup.WebsocketExtensioner")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// IOStream: get the I/O stream the WebSocket is communicating over.
func (self *WebsocketConnection) IOStream() gio.IOStreamer {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.GIOStream               // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_io_stream(_arg0)
	runtime.KeepAlive(self)

	var _ioStream gio.IOStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gio.IOStreamer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.IOStreamer")
		}
		_ioStream = rv
	}

	return _ioStream
}

// KeepaliveInterval gets the keepalive interval in seconds or 0 if disabled.
func (self *WebsocketConnection) KeepaliveInterval() uint {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.guint                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_keepalive_interval(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MaxIncomingPayloadSize gets the maximum payload size allowed for incoming
// packets.
func (self *WebsocketConnection) MaxIncomingPayloadSize() uint64 {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.guint64                  // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_max_incoming_payload_size(_arg0)
	runtime.KeepAlive(self)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Origin: get the origin of the WebSocket.
func (self *WebsocketConnection) Origin() string {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.char                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_origin(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Protocol: get the protocol chosen via negotiation with the peer.
func (self *WebsocketConnection) Protocol() string {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.char                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_protocol(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// State: get the current state of the WebSocket.
func (self *WebsocketConnection) State() WebsocketState {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.SoupWebsocketState       // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_state(_arg0)
	runtime.KeepAlive(self)

	var _websocketState WebsocketState // out

	_websocketState = WebsocketState(_cret)

	return _websocketState
}

// URI: get the URI of the WebSocket.
//
// For servers this represents the address of the WebSocket, and for clients it
// is the address connected to.
func (self *WebsocketConnection) URI() *URI {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.SoupURI                 // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))

	_cret = C.soup_websocket_connection_get_uri(_arg0)
	runtime.KeepAlive(self)

	var _urI *URI // out

	_urI = (*URI)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _urI
}

// SendBinary: send a binary message to the peer. If length is 0, data may be
// NULL.
//
// The message is queued to be sent and will be sent when the main loop is run.
//
// The function takes the following parameters:
//
//    - data: message contents.
//
func (self *WebsocketConnection) SendBinary(data []byte) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 C.gconstpointer            // out
	var _arg2 C.gsize

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))
	_arg2 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg1 = (C.gconstpointer)(unsafe.Pointer(&data[0]))
	}

	C.soup_websocket_connection_send_binary(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(data)
}

// SendMessage: send a message of the given type to the peer. Note that this
// method, allows to send text messages containing NULL characters.
//
// The message is queued to be sent and will be sent when the main loop is run.
//
// The function takes the following parameters:
//
//    - typ: type of message contents.
//    - message data as #GBytes.
//
func (self *WebsocketConnection) SendMessage(typ WebsocketDataType, message *glib.Bytes) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 C.SoupWebsocketDataType    // out
	var _arg2 *C.GBytes                  // out

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))
	_arg1 = C.SoupWebsocketDataType(typ)
	_arg2 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(message)))

	C.soup_websocket_connection_send_message(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(message)
}

// SendText: send a NULL-terminated text (UTF-8) message to the peer. If you
// need to send text messages containing NULL characters use
// soup_websocket_connection_send_message() instead.
//
// The message is queued to be sent and will be sent when the main loop is run.
//
// The function takes the following parameters:
//
//    - text: message contents.
//
func (self *WebsocketConnection) SendText(text string) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 *C.char                    // out

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_websocket_connection_send_text(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(text)
}

// SetKeepaliveInterval sets the interval in seconds on when to send a ping
// message which will serve as a keepalive message. If set to 0 the keepalive
// message is disabled.
//
// The function takes the following parameters:
//
//    - interval to send a ping message or 0 to disable it.
//
func (self *WebsocketConnection) SetKeepaliveInterval(interval uint) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 C.guint                    // out

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(interval)

	C.soup_websocket_connection_set_keepalive_interval(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(interval)
}

// SetMaxIncomingPayloadSize sets the maximum payload size allowed for incoming
// packets. It does not limit the outgoing packet size.
//
// The function takes the following parameters:
//
//    - maxIncomingPayloadSize: maximum payload size.
//
func (self *WebsocketConnection) SetMaxIncomingPayloadSize(maxIncomingPayloadSize uint64) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 C.guint64                  // out

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint64(maxIncomingPayloadSize)

	C.soup_websocket_connection_set_max_incoming_payload_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(maxIncomingPayloadSize)
}

// ConnectClosed: emitted when the connection has completely closed, either due
// to an orderly close from the peer, one initiated via
// soup_websocket_connection_close() or a fatal error condition that caused a
// close.
//
// This signal will be emitted once.
func (self *WebsocketConnection) ConnectClosed(f func()) externglib.SignalHandle {
	return self.Connect("closed", f)
}

// ConnectClosing: this signal will be emitted during an orderly close.
func (self *WebsocketConnection) ConnectClosing(f func()) externglib.SignalHandle {
	return self.Connect("closing", f)
}

// ConnectMessage: emitted when we receive a message from the peer.
//
// As a convenience, the message data will always be NUL-terminated, but the NUL
// byte will not be included in the length count.
func (self *WebsocketConnection) ConnectMessage(f func(typ int, message *glib.Bytes)) externglib.SignalHandle {
	return self.Connect("message", f)
}

// ConnectPong: emitted when we receive a Pong frame (solicited or unsolicited)
// from the peer.
//
// As a convenience, the message data will always be NUL-terminated, but the NUL
// byte will not be included in the length count.
func (self *WebsocketConnection) ConnectPong(f func(message *glib.Bytes)) externglib.SignalHandle {
	return self.Connect("pong", f)
}
