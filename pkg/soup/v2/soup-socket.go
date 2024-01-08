// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern void _gotk4_soup2_Socket_ConnectWritable(gpointer, guintptr);
// extern void _gotk4_soup2_Socket_ConnectReadable(gpointer, guintptr);
// extern void _gotk4_soup2_Socket_ConnectNewConnection(gpointer, SoupSocket*, guintptr);
// extern void _gotk4_soup2_Socket_ConnectEvent(gpointer, GSocketClientEvent, GIOStream*, guintptr);
// extern void _gotk4_soup2_Socket_ConnectDisconnected(gpointer, guintptr);
// extern void _gotk4_soup2_SocketClass_writable(SoupSocket*);
// extern void _gotk4_soup2_SocketClass_readable(SoupSocket*);
// extern void _gotk4_soup2_SocketClass_new_connection(SoupSocket*, SoupSocket*);
// extern void _gotk4_soup2_SocketClass_disconnected(SoupSocket*);
// extern void _gotk4_soup2_SocketCallback(SoupSocket*, guint, gpointer);
// void _gotk4_soup2_Socket_virtual_disconnected(void* fnptr, SoupSocket* arg0) {
//   ((void (*)(SoupSocket*))(fnptr))(arg0);
// };
// void _gotk4_soup2_Socket_virtual_new_connection(void* fnptr, SoupSocket* arg0, SoupSocket* arg1) {
//   ((void (*)(SoupSocket*, SoupSocket*))(fnptr))(arg0, arg1);
// };
// void _gotk4_soup2_Socket_virtual_readable(void* fnptr, SoupSocket* arg0) {
//   ((void (*)(SoupSocket*))(fnptr))(arg0);
// };
// void _gotk4_soup2_Socket_virtual_writable(void* fnptr, SoupSocket* arg0) {
//   ((void (*)(SoupSocket*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeSocketIOStatus = coreglib.Type(C.soup_socket_io_status_get_type())
	GTypeSocket         = coreglib.Type(C.soup_socket_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSocketIOStatus, F: marshalSocketIOStatus},
		coreglib.TypeMarshaler{T: GTypeSocket, F: marshalSocket},
	})
}

// SOCKET_ASYNC_CONTEXT alias for the Socket:async-context property. (The
// socket's Context.).
const SOCKET_ASYNC_CONTEXT = "async-context"

// SOCKET_FLAG_NONBLOCKING alias for the Socket:non-blocking property. (Whether
// or not the socket uses non-blocking I/O.).
const SOCKET_FLAG_NONBLOCKING = "non-blocking"

// SOCKET_IS_SERVER alias for the Socket:is-server property, qv.
const SOCKET_IS_SERVER = "is-server"

// SOCKET_LOCAL_ADDRESS alias for the Socket:local-address property. (Address of
// local end of socket.).
const SOCKET_LOCAL_ADDRESS = "local-address"

// SOCKET_REMOTE_ADDRESS alias for the Socket:remote-address property. (Address
// of remote end of socket.).
const SOCKET_REMOTE_ADDRESS = "remote-address"

// SOCKET_SSL_CREDENTIALS alias for the Socket:ssl-creds property. (SSL
// credential information.).
const SOCKET_SSL_CREDENTIALS = "ssl-creds"

// SOCKET_SSL_FALLBACK alias for the Socket:ssl-fallback property.
const SOCKET_SSL_FALLBACK = "ssl-fallback"

// SOCKET_SSL_STRICT alias for the Socket:ssl-strict property.
const SOCKET_SSL_STRICT = "ssl-strict"

// SOCKET_TIMEOUT alias for the Socket:timeout property. (The timeout in seconds
// for blocking socket I/O operations.).
const SOCKET_TIMEOUT = "timeout"

// SOCKET_TRUSTED_CERTIFICATE alias for the Socket:trusted-certificate property.
const SOCKET_TRUSTED_CERTIFICATE = "trusted-certificate"

// SocketIOStatus: return value from the Socket IO methods.
type SocketIOStatus C.gint

const (
	// SocketOK: success.
	SocketOK SocketIOStatus = iota
	// SocketWouldBlock: cannot read/write any more at this time.
	SocketWouldBlock
	// SocketEOF: end of file.
	SocketEOF
	// SocketError: other error.
	SocketError
)

func marshalSocketIOStatus(p uintptr) (interface{}, error) {
	return SocketIOStatus(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SocketIOStatus.
func (s SocketIOStatus) String() string {
	switch s {
	case SocketOK:
		return "OK"
	case SocketWouldBlock:
		return "WouldBlock"
	case SocketEOF:
		return "EOF"
	case SocketError:
		return "Error"
	default:
		return fmt.Sprintf("SocketIOStatus(%d)", s)
	}
}

// SocketCallback: callback function passed to soup_socket_connect_async().
type SocketCallback func(sock *Socket, status uint)

// SocketOverrides contains methods that are overridable.
type SocketOverrides struct {
	Disconnected func()
	// The function takes the following parameters:
	//
	NewConnection func(newSock *Socket)
	Readable      func()
	Writable      func()
}

func defaultSocketOverrides(v *Socket) SocketOverrides {
	return SocketOverrides{
		Disconnected:  v.disconnected,
		NewConnection: v.newConnection,
		Readable:      v.readable,
		Writable:      v.writable,
	}
}

type Socket struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.Initable
}

var (
	_ coreglib.Objector = (*Socket)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Socket, *SocketClass, SocketOverrides](
		GTypeSocket,
		initSocketClass,
		wrapSocket,
		defaultSocketOverrides,
	)
}

func initSocketClass(gclass unsafe.Pointer, overrides SocketOverrides, classInitFunc func(*SocketClass)) {
	pclass := (*C.SoupSocketClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeSocket))))

	if overrides.Disconnected != nil {
		pclass.disconnected = (*[0]byte)(C._gotk4_soup2_SocketClass_disconnected)
	}

	if overrides.NewConnection != nil {
		pclass.new_connection = (*[0]byte)(C._gotk4_soup2_SocketClass_new_connection)
	}

	if overrides.Readable != nil {
		pclass.readable = (*[0]byte)(C._gotk4_soup2_SocketClass_readable)
	}

	if overrides.Writable != nil {
		pclass.writable = (*[0]byte)(C._gotk4_soup2_SocketClass_writable)
	}

	if classInitFunc != nil {
		class := (*SocketClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSocket(obj *coreglib.Object) *Socket {
	return &Socket{
		Object: obj,
		Initable: gio.Initable{
			Object: obj,
		},
	}
}

func marshalSocket(p uintptr) (interface{}, error) {
	return wrapSocket(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectDisconnected is emitted when the socket is disconnected, for whatever
// reason.
func (sock *Socket) ConnectDisconnected(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sock, "disconnected", false, unsafe.Pointer(C._gotk4_soup2_Socket_ConnectDisconnected), f)
}

// ConnectEvent is emitted when a network-related event occurs. See
// Client::event for more details.
func (sock *Socket) ConnectEvent(f func(event gio.SocketClientEvent, connection gio.IOStreamer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sock, "event", false, unsafe.Pointer(C._gotk4_soup2_Socket_ConnectEvent), f)
}

// ConnectNewConnection is emitted when a listening socket (set up with
// soup_socket_listen()) receives a new connection.
//
// You must ref the new if you want to keep it; otherwise it will be destroyed
// after the signal is emitted.
func (sock *Socket) ConnectNewConnection(f func(new *Socket)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sock, "new-connection", false, unsafe.Pointer(C._gotk4_soup2_Socket_ConnectNewConnection), f)
}

// ConnectReadable is emitted when an async socket is readable. See
// soup_socket_read(), soup_socket_read_until() and Socket:non-blocking.
func (sock *Socket) ConnectReadable(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sock, "readable", false, unsafe.Pointer(C._gotk4_soup2_Socket_ConnectReadable), f)
}

// ConnectWritable is emitted when an async socket is writable. See
// soup_socket_write() and Socket:non-blocking.
func (sock *Socket) ConnectWritable(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sock, "writable", false, unsafe.Pointer(C._gotk4_soup2_Socket_ConnectWritable), f)
}

// ConnectAsync begins asynchronously connecting to sock's remote address. The
// socket will call callback when it succeeds or fails (but not before returning
// from this function).
//
// If cancellable is non-NULL, it can be used to cancel the connection. callback
// will still be invoked in this case, with a status of SOUP_STATUS_CANCELLED.
//
// The function takes the following parameters:
//
//   - ctx (optional) or NULL.
//   - callback to call after connecting.
//
func (sock *Socket) ConnectAsync(ctx context.Context, callback SocketCallback) {
	var _arg0 *C.SoupSocket        // out
	var _arg1 *C.GCancellable      // out
	var _arg2 C.SoupSocketCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (*[0]byte)(C._gotk4_soup2_SocketCallback)
	_arg3 = C.gpointer(gbox.AssignOnce(callback))

	C.soup_socket_connect_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// ConnectSync: attempt to synchronously connect sock to its remote address.
//
// If cancellable is non-NULL, it can be used to cancel the connection, in which
// case soup_socket_connect_sync() will return SOUP_STATUS_CANCELLED.
//
// The function takes the following parameters:
//
//   - ctx (optional) or NULL.
//
// The function returns the following values:
//
//   - guint success or failure code.
//
func (sock *Socket) ConnectSync(ctx context.Context) uint {
	var _arg0 *C.SoupSocket   // out
	var _arg1 *C.GCancellable // out
	var _cret C.guint         // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.soup_socket_connect_sync(_arg0, _arg1)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Disconnect disconnects sock. Any further read or write attempts on it will
// fail.
func (sock *Socket) Disconnect() {
	var _arg0 *C.SoupSocket // out

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	C.soup_socket_disconnect(_arg0)
	runtime.KeepAlive(sock)
}

// Fd gets sock's underlying file descriptor.
//
// Note that fiddling with the file descriptor may break the Socket.
//
// The function returns the following values:
//
//   - gint sock's file descriptor.
//
func (sock *Socket) Fd() int {
	var _arg0 *C.SoupSocket // out
	var _cret C.int         // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	_cret = C.soup_socket_get_fd(_arg0)
	runtime.KeepAlive(sock)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// LocalAddress returns the Address corresponding to the local end of sock.
//
// Calling this method on an unconnected socket is considered to be an error,
// and produces undefined results.
//
// The function returns the following values:
//
//   - address: Address.
//
func (sock *Socket) LocalAddress() *Address {
	var _arg0 *C.SoupSocket  // out
	var _cret *C.SoupAddress // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	_cret = C.soup_socket_get_local_address(_arg0)
	runtime.KeepAlive(sock)

	var _address *Address // out

	_address = wrapAddress(coreglib.Take(unsafe.Pointer(_cret)))

	return _address
}

// RemoteAddress returns the Address corresponding to the remote end of sock.
//
// Calling this method on an unconnected socket is considered to be an error,
// and produces undefined results.
//
// The function returns the following values:
//
//   - address: Address.
//
func (sock *Socket) RemoteAddress() *Address {
	var _arg0 *C.SoupSocket  // out
	var _cret *C.SoupAddress // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	_cret = C.soup_socket_get_remote_address(_arg0)
	runtime.KeepAlive(sock)

	var _address *Address // out

	_address = wrapAddress(coreglib.Take(unsafe.Pointer(_cret)))

	return _address
}

// IsConnected tests if sock is connected to another host.
//
// The function returns the following values:
//
//   - ok: TRUE or FALSE.
//
func (sock *Socket) IsConnected() bool {
	var _arg0 *C.SoupSocket // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	_cret = C.soup_socket_is_connected(_arg0)
	runtime.KeepAlive(sock)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSSL tests if sock is doing (or has attempted to do) SSL.
//
// The function returns the following values:
//
//   - ok: TRUE if sock has SSL credentials set.
//
func (sock *Socket) IsSSL() bool {
	var _arg0 *C.SoupSocket // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	_cret = C.soup_socket_is_ssl(_arg0)
	runtime.KeepAlive(sock)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Listen makes sock start listening on its local address. When connections come
// in, sock will emit Socket::new_connection.
//
// The function returns the following values:
//
//   - ok: whether or not sock is now listening.
//
func (sock *Socket) Listen() bool {
	var _arg0 *C.SoupSocket // out
	var _cret C.gboolean    // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	_cret = C.soup_socket_listen(_arg0)
	runtime.KeepAlive(sock)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Read attempts to read up to len bytes from sock into buffer. If some data is
// successfully read, soup_socket_read() will return SOUP_SOCKET_OK, and *nread
// will contain the number of bytes actually read (which may be less than len).
//
// If sock is non-blocking, and no data is available, the return value will
// be SOUP_SOCKET_WOULD_BLOCK. In this case, the caller can connect to the
// Socket::readable signal to know when there is more data to read. (NB:
// You MUST read all available data off the socket first. Socket::readable
// is only emitted after soup_socket_read() returns SOUP_SOCKET_WOULD_BLOCK,
// and it is only emitted once. See the documentation for Socket:non-blocking.).
//
// The function takes the following parameters:
//
//   - ctx (optional) or NULL.
//   - buffer to read into.
//
// The function returns the following values:
//
//   - nread: on return, the number of bytes read into buffer.
//   - socketIOStatus as described above (or SOUP_SOCKET_EOF if the socket is no
//     longer connected, or SOUP_SOCKET_ERROR on any other error, in which case
//     error will also be set).
//
func (sock *Socket) Read(ctx context.Context, buffer []byte) (uint, SocketIOStatus, error) {
	var _arg0 *C.SoupSocket   // out
	var _arg4 *C.GCancellable // out
	var _arg1 C.gpointer      // out
	var _arg2 C.gsize
	var _arg3 C.gsize              // in
	var _cret C.SoupSocketIOStatus // in
	var _cerr *C.GError            // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (C.gpointer)(unsafe.Pointer(&buffer[0]))
	}

	_cret = C.soup_socket_read(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _nread uint                    // out
	var _socketIOStatus SocketIOStatus // out
	var _goerr error                   // out

	_nread = uint(_arg3)
	_socketIOStatus = SocketIOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _nread, _socketIOStatus, _goerr
}

// ReadUntil: like soup_socket_read(), but reads no further than the first
// occurrence of boundary. (If the boundary is found, it will be included in
// the returned data, and *got_boundary will be set to TRUE.) Any data after the
// boundary will returned in future reads.
//
// soup_socket_read_until() will almost always return fewer than len bytes:
// if the boundary is found, then it will only return the bytes up until the
// end of the boundary, and if the boundary is not found, then it will leave
// the last <literal>(boundary_len - 1)</literal> bytes in its internal buffer,
// in case they form the start of the boundary string. Thus, len normally
// needs to be at least 1 byte longer than boundary_len if you want to make any
// progress at all.
//
// The function takes the following parameters:
//
//   - ctx (optional) or NULL.
//   - buffer to read into.
//   - boundary (optional) to read until.
//   - boundaryLen: length of boundary in bytes.
//
// The function returns the following values:
//
//   - nread: on return, the number of bytes read into buffer.
//   - gotBoundary: on return, whether or not the data in buffer ends with the
//     boundary string.
//   - socketIOStatus as for soup_socket_read().
//
func (sock *Socket) ReadUntil(ctx context.Context, buffer []byte, boundary unsafe.Pointer, boundaryLen uint) (uint, bool, SocketIOStatus, error) {
	var _arg0 *C.SoupSocket   // out
	var _arg7 *C.GCancellable // out
	var _arg1 C.gpointer      // out
	var _arg2 C.gsize
	var _arg3 C.gconstpointer      // out
	var _arg4 C.gsize              // out
	var _arg5 C.gsize              // in
	var _arg6 C.gboolean           // in
	var _cret C.SoupSocketIOStatus // in
	var _cerr *C.GError            // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (C.gpointer)(unsafe.Pointer(&buffer[0]))
	}
	_arg3 = (C.gconstpointer)(unsafe.Pointer(boundary))
	_arg4 = C.gsize(boundaryLen)

	_cret = C.soup_socket_read_until(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_arg6, _arg7, &_cerr)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(boundary)
	runtime.KeepAlive(boundaryLen)

	var _nread uint                    // out
	var _gotBoundary bool              // out
	var _socketIOStatus SocketIOStatus // out
	var _goerr error                   // out

	_nread = uint(_arg5)
	if _arg6 != 0 {
		_gotBoundary = true
	}
	_socketIOStatus = SocketIOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _nread, _gotBoundary, _socketIOStatus, _goerr
}

// StartProxySsl starts using SSL on socket, expecting to find a host named
// ssl_host.
//
// The function takes the following parameters:
//
//   - ctx (optional): #GCancellable.
//   - sslHost: hostname of the SSL server.
//
// The function returns the following values:
//
//   - ok success or failure.
//
func (sock *Socket) StartProxySsl(ctx context.Context, sslHost string) bool {
	var _arg0 *C.SoupSocket   // out
	var _arg2 *C.GCancellable // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(sslHost)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_socket_start_proxy_ssl(_arg0, _arg1, _arg2)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(sslHost)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartSSL starts using SSL on socket.
//
// The function takes the following parameters:
//
//   - ctx (optional): #GCancellable.
//
// The function returns the following values:
//
//   - ok success or failure.
//
func (sock *Socket) StartSSL(ctx context.Context) bool {
	var _arg0 *C.SoupSocket   // out
	var _arg1 *C.GCancellable // out
	var _cret C.gboolean      // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.soup_socket_start_ssl(_arg0, _arg1)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Write attempts to write len bytes from buffer to sock. If some data is
// successfully written, the return status will be SOUP_SOCKET_OK, and *nwrote
// will contain the number of bytes actually written (which may be less than
// len).
//
// If sock is non-blocking, and no data could be written right away, the return
// value will be SOUP_SOCKET_WOULD_BLOCK. In this case, the caller can connect
// to the Socket::writable signal to know when more data can be written.
// (NB: Socket::writable is only emitted after soup_socket_write() returns
// SOUP_SOCKET_WOULD_BLOCK, and it is only emitted once. See the documentation
// for Socket:non-blocking.).
//
// The function takes the following parameters:
//
//   - ctx (optional) or NULL.
//   - buffer: data to write.
//
// The function returns the following values:
//
//   - nwrote: on return, number of bytes written.
//   - socketIOStatus as described above (or SOUP_SOCKET_EOF or
//     SOUP_SOCKET_ERROR. error will be set if the return value is
//     SOUP_SOCKET_ERROR.).
//
func (sock *Socket) Write(ctx context.Context, buffer []byte) (uint, SocketIOStatus, error) {
	var _arg0 *C.SoupSocket   // out
	var _arg4 *C.GCancellable // out
	var _arg1 C.gconstpointer // out
	var _arg2 C.gsize
	var _arg3 C.gsize              // in
	var _cret C.SoupSocketIOStatus // in
	var _cerr *C.GError            // in

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg1 = (C.gconstpointer)(unsafe.Pointer(&buffer[0]))
	}

	_cret = C.soup_socket_write(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)
	runtime.KeepAlive(sock)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(buffer)

	var _nwrote uint                   // out
	var _socketIOStatus SocketIOStatus // out
	var _goerr error                   // out

	_nwrote = uint(_arg3)
	_socketIOStatus = SocketIOStatus(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _nwrote, _socketIOStatus, _goerr
}

func (sock *Socket) disconnected() {
	gclass := (*C.SoupSocketClass)(coreglib.PeekParentClass(sock))
	fnarg := gclass.disconnected

	var _arg0 *C.SoupSocket // out

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	C._gotk4_soup2_Socket_virtual_disconnected(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(sock)
}

// The function takes the following parameters:
//
func (listener *Socket) newConnection(newSock *Socket) {
	gclass := (*C.SoupSocketClass)(coreglib.PeekParentClass(listener))
	fnarg := gclass.new_connection

	var _arg0 *C.SoupSocket // out
	var _arg1 *C.SoupSocket // out

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(newSock).Native()))

	C._gotk4_soup2_Socket_virtual_new_connection(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(newSock)
}

func (sock *Socket) readable() {
	gclass := (*C.SoupSocketClass)(coreglib.PeekParentClass(sock))
	fnarg := gclass.readable

	var _arg0 *C.SoupSocket // out

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	C._gotk4_soup2_Socket_virtual_readable(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(sock)
}

func (sock *Socket) writable() {
	gclass := (*C.SoupSocketClass)(coreglib.PeekParentClass(sock))
	fnarg := gclass.writable

	var _arg0 *C.SoupSocket // out

	_arg0 = (*C.SoupSocket)(unsafe.Pointer(coreglib.InternObject(sock).Native()))

	C._gotk4_soup2_Socket_virtual_writable(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(sock)
}

// SocketClass: instance of this type is always passed by reference.
type SocketClass struct {
	*socketClass
}

// socketClass is the struct that's finalized.
type socketClass struct {
	native *C.SoupSocketClass
}
