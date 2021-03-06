// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern void _gotk4_soup2_WebsocketConnectionClass_closed(SoupWebsocketConnection*);
// extern void _gotk4_soup2_WebsocketConnectionClass_closing(SoupWebsocketConnection*);
// extern void _gotk4_soup2_WebsocketConnectionClass_error(SoupWebsocketConnection*, GError*);
// extern void _gotk4_soup2_WebsocketConnectionClass_message(SoupWebsocketConnection*, SoupWebsocketDataType, GBytes*);
// extern void _gotk4_soup2_WebsocketConnectionClass_pong(SoupWebsocketConnection*, GBytes*);
// extern void _gotk4_soup2_WebsocketConnection_ConnectClosed(gpointer, guintptr);
// extern void _gotk4_soup2_WebsocketConnection_ConnectClosing(gpointer, guintptr);
// extern void _gotk4_soup2_WebsocketConnection_ConnectError(gpointer, GError*, guintptr);
// extern void _gotk4_soup2_WebsocketConnection_ConnectMessage(gpointer, gint, GBytes*, guintptr);
// extern void _gotk4_soup2_WebsocketConnection_ConnectPong(gpointer, GBytes*, guintptr);
import "C"

// glib.Type values for soup-websocket-connection.go.
var GTypeWebsocketConnection = externglib.Type(C.soup_websocket_connection_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeWebsocketConnection, F: marshalWebsocketConnection},
	})
}

// WebsocketConnectionOverrider contains methods that are overridable.
type WebsocketConnectionOverrider interface {
	Closed()
	Closing()
	// The function takes the following parameters:
	//
	Error(err error)
	// The function takes the following parameters:
	//
	//    - typ
	//    - message
	//
	Message(typ WebsocketDataType, message *glib.Bytes)
	// The function takes the following parameters:
	//
	Pong(message *glib.Bytes)
}

// WebsocketConnection class representing a WebSocket connection.
type WebsocketConnection struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*WebsocketConnection)(nil)
)

func classInitWebsocketConnectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.SoupWebsocketConnectionClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.SoupWebsocketConnectionClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Closed() }); ok {
		pclass.closed = (*[0]byte)(C._gotk4_soup2_WebsocketConnectionClass_closed)
	}

	if _, ok := goval.(interface{ Closing() }); ok {
		pclass.closing = (*[0]byte)(C._gotk4_soup2_WebsocketConnectionClass_closing)
	}

	if _, ok := goval.(interface{ Error(err error) }); ok {
		pclass.error = (*[0]byte)(C._gotk4_soup2_WebsocketConnectionClass_error)
	}

	if _, ok := goval.(interface {
		Message(typ WebsocketDataType, message *glib.Bytes)
	}); ok {
		pclass.message = (*[0]byte)(C._gotk4_soup2_WebsocketConnectionClass_message)
	}

	if _, ok := goval.(interface{ Pong(message *glib.Bytes) }); ok {
		pclass.pong = (*[0]byte)(C._gotk4_soup2_WebsocketConnectionClass_pong)
	}
}

//export _gotk4_soup2_WebsocketConnectionClass_closed
func _gotk4_soup2_WebsocketConnectionClass_closed(arg0 *C.SoupWebsocketConnection) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Closed() })

	iface.Closed()
}

//export _gotk4_soup2_WebsocketConnectionClass_closing
func _gotk4_soup2_WebsocketConnectionClass_closing(arg0 *C.SoupWebsocketConnection) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Closing() })

	iface.Closing()
}

//export _gotk4_soup2_WebsocketConnectionClass_error
func _gotk4_soup2_WebsocketConnectionClass_error(arg0 *C.SoupWebsocketConnection, arg1 *C.GError) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Error(err error) })

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(arg1))

	iface.Error(_err)
}

//export _gotk4_soup2_WebsocketConnectionClass_message
func _gotk4_soup2_WebsocketConnectionClass_message(arg0 *C.SoupWebsocketConnection, arg1 C.SoupWebsocketDataType, arg2 *C.GBytes) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Message(typ WebsocketDataType, message *glib.Bytes)
	})

	var _typ WebsocketDataType // out
	var _message *glib.Bytes   // out

	_typ = WebsocketDataType(arg1)
	_message = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.g_bytes_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	iface.Message(_typ, _message)
}

//export _gotk4_soup2_WebsocketConnectionClass_pong
func _gotk4_soup2_WebsocketConnectionClass_pong(arg0 *C.SoupWebsocketConnection, arg1 *C.GBytes) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Pong(message *glib.Bytes) })

	var _message *glib.Bytes // out

	_message = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.g_bytes_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	iface.Pong(_message)
}

func wrapWebsocketConnection(obj *externglib.Object) *WebsocketConnection {
	return &WebsocketConnection{
		Object: obj,
	}
}

func marshalWebsocketConnection(p uintptr) (interface{}, error) {
	return wrapWebsocketConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_soup2_WebsocketConnection_ConnectClosed
func _gotk4_soup2_WebsocketConnection_ConnectClosed(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClosed is emitted when the connection has completely closed, either
// due to an orderly close from the peer, one initiated via
// soup_websocket_connection_close() or a fatal error condition that caused a
// close.
//
// This signal will be emitted once.
func (self *WebsocketConnection) ConnectClosed(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "closed", false, unsafe.Pointer(C._gotk4_soup2_WebsocketConnection_ConnectClosed), f)
}

//export _gotk4_soup2_WebsocketConnection_ConnectClosing
func _gotk4_soup2_WebsocketConnection_ConnectClosing(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClosing: this signal will be emitted during an orderly close.
func (self *WebsocketConnection) ConnectClosing(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "closing", false, unsafe.Pointer(C._gotk4_soup2_WebsocketConnection_ConnectClosing), f)
}

//export _gotk4_soup2_WebsocketConnection_ConnectError
func _gotk4_soup2_WebsocketConnection_ConnectError(arg0 C.gpointer, arg1 *C.GError, arg2 C.guintptr) {
	var f func(err error)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(err error))
	}

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(arg1))

	f(_err)
}

// ConnectError is emitted when an error occurred on the WebSocket. This may be
// fired multiple times. Fatal errors will be followed by the
// WebsocketConnection::closed signal being emitted.
func (self *WebsocketConnection) ConnectError(f func(err error)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "error", false, unsafe.Pointer(C._gotk4_soup2_WebsocketConnection_ConnectError), f)
}

//export _gotk4_soup2_WebsocketConnection_ConnectMessage
func _gotk4_soup2_WebsocketConnection_ConnectMessage(arg0 C.gpointer, arg1 C.gint, arg2 *C.GBytes, arg3 C.guintptr) {
	var f func(typ int, message *glib.Bytes)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(typ int, message *glib.Bytes))
	}

	var _typ int             // out
	var _message *glib.Bytes // out

	_typ = int(arg1)
	_message = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.g_bytes_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	f(_typ, _message)
}

// ConnectMessage is emitted when we receive a message from the peer.
//
// As a convenience, the message data will always be NUL-terminated, but the NUL
// byte will not be included in the length count.
func (self *WebsocketConnection) ConnectMessage(f func(typ int, message *glib.Bytes)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "message", false, unsafe.Pointer(C._gotk4_soup2_WebsocketConnection_ConnectMessage), f)
}

//export _gotk4_soup2_WebsocketConnection_ConnectPong
func _gotk4_soup2_WebsocketConnection_ConnectPong(arg0 C.gpointer, arg1 *C.GBytes, arg2 C.guintptr) {
	var f func(message *glib.Bytes)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(message *glib.Bytes))
	}

	var _message *glib.Bytes // out

	_message = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.g_bytes_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	f(_message)
}

// ConnectPong is emitted when we receive a Pong frame (solicited or
// unsolicited) from the peer.
//
// As a convenience, the message data will always be NUL-terminated, but the NUL
// byte will not be included in the length count.
func (self *WebsocketConnection) ConnectPong(f func(message *glib.Bytes)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "pong", false, unsafe.Pointer(C._gotk4_soup2_WebsocketConnection_ConnectPong), f)
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
//    - origin (optional): origin of the client.
//    - protocol (optional) in use.
//
// The function returns the following values:
//
//    - websocketConnection: new WebsocketConnection.
//
func NewWebsocketConnection(stream gio.IOStreamer, uri *URI, typ WebsocketConnectionType, origin, protocol string) *WebsocketConnection {
	var _arg1 *C.GIOStream                  // out
	var _arg2 *C.SoupURI                    // out
	var _arg3 C.SoupWebsocketConnectionType // out
	var _arg4 *C.char                       // out
	var _arg5 *C.char                       // out
	var _cret *C.SoupWebsocketConnection    // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))
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
//    - origin (optional): origin of the client.
//    - protocol (optional) in use.
//    - extensions of WebsocketExtension objects.
//
// The function returns the following values:
//
//    - websocketConnection: new WebsocketConnection.
//
func NewWebsocketConnectionWithExtensions(stream gio.IOStreamer, uri *URI, typ WebsocketConnectionType, origin, protocol string, extensions []WebsocketExtensioner) *WebsocketConnection {
	var _arg1 *C.GIOStream                  // out
	var _arg2 *C.SoupURI                    // out
	var _arg3 C.SoupWebsocketConnectionType // out
	var _arg4 *C.char                       // out
	var _arg5 *C.char                       // out
	var _arg6 *C.GList                      // out
	var _cret *C.SoupWebsocketConnection    // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))
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
		dst = (*C.SoupWebsocketExtension)(unsafe.Pointer(externglib.InternObject(src).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(src).Native()))
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
//    - data (optional): close data.
//
func (self *WebsocketConnection) Close(code uint16, data string) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 C.gushort                  // out
	var _arg2 *C.char                    // out

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))
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
//
// The function returns the following values:
//
//    - gushort: close code or zero.
//
func (self *WebsocketConnection) CloseCode() uint16 {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.gushort                  // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

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
//
// The function returns the following values:
//
//    - utf8: close data or NULL.
//
func (self *WebsocketConnection) CloseData() string {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.char                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.soup_websocket_connection_get_close_data(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ConnectionType: get the connection type (client/server) of the connection.
//
// The function returns the following values:
//
//    - websocketConnectionType: connection type.
//
func (self *WebsocketConnection) ConnectionType() WebsocketConnectionType {
	var _arg0 *C.SoupWebsocketConnection    // out
	var _cret C.SoupWebsocketConnectionType // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.soup_websocket_connection_get_connection_type(_arg0)
	runtime.KeepAlive(self)

	var _websocketConnectionType WebsocketConnectionType // out

	_websocketConnectionType = WebsocketConnectionType(_cret)

	return _websocketConnectionType
}

// Extensions: get the extensions chosen via negotiation with the peer.
//
// The function returns the following values:
//
//    - list of WebsocketExtension objects.
//
func (self *WebsocketConnection) Extensions() []WebsocketExtensioner {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.GList                   // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

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
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(WebsocketExtensioner)
				return ok
			})
			rv, ok := casted.(WebsocketExtensioner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching soup.WebsocketExtensioner")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// IOStream: get the I/O stream the WebSocket is communicating over.
//
// The function returns the following values:
//
//    - ioStream webSocket's I/O stream.
//
func (self *WebsocketConnection) IOStream() gio.IOStreamer {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.GIOStream               // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.soup_websocket_connection_get_io_stream(_arg0)
	runtime.KeepAlive(self)

	var _ioStream gio.IOStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.IOStreamer)
			return ok
		})
		rv, ok := casted.(gio.IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}

	return _ioStream
}

// KeepaliveInterval gets the keepalive interval in seconds or 0 if disabled.
//
// The function returns the following values:
//
//    - guint: keepalive interval.
//
func (self *WebsocketConnection) KeepaliveInterval() uint {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.guint                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.soup_websocket_connection_get_keepalive_interval(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MaxIncomingPayloadSize gets the maximum payload size allowed for incoming
// packets.
//
// The function returns the following values:
//
//    - guint64: maximum payload size.
//
func (self *WebsocketConnection) MaxIncomingPayloadSize() uint64 {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.guint64                  // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.soup_websocket_connection_get_max_incoming_payload_size(_arg0)
	runtime.KeepAlive(self)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Origin: get the origin of the WebSocket.
//
// The function returns the following values:
//
//    - utf8 (optional): origin, or NULL.
//
func (self *WebsocketConnection) Origin() string {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.char                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.soup_websocket_connection_get_origin(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Protocol: get the protocol chosen via negotiation with the peer.
//
// The function returns the following values:
//
//    - utf8 (optional): chosen protocol, or NULL.
//
func (self *WebsocketConnection) Protocol() string {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.char                    // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.soup_websocket_connection_get_protocol(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// State: get the current state of the WebSocket.
//
// The function returns the following values:
//
//    - websocketState: state.
//
func (self *WebsocketConnection) State() WebsocketState {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret C.SoupWebsocketState       // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

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
//
// The function returns the following values:
//
//    - urI: URI.
//
func (self *WebsocketConnection) URI() *URI {
	var _arg0 *C.SoupWebsocketConnection // out
	var _cret *C.SoupURI                 // in

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))

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
//    - data (optional): message contents.
//
func (self *WebsocketConnection) SendBinary(data []byte) {
	var _arg0 *C.SoupWebsocketConnection // out
	var _arg1 C.gconstpointer            // out
	var _arg2 C.gsize

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))
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

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))
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

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))
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

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))
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

	_arg0 = (*C.SoupWebsocketConnection)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.guint64(maxIncomingPayloadSize)

	C.soup_websocket_connection_set_max_incoming_payload_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(maxIncomingPayloadSize)
}
