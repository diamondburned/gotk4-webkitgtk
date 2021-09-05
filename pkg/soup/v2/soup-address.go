// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
// void _gotk4_soup2_AddressCallback(SoupAddress*, guint, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_address_family_get_type()), F: marshalAddressFamily},
		{T: externglib.Type(C.soup_address_get_type()), F: marshalAddresser},
	})
}

// ADDRESS_ANY_PORT: this can be passed to any Address method that expects a
// port, to indicate that you don't care what port is used.
const ADDRESS_ANY_PORT = 0

// ADDRESS_FAMILY alias for the Address:family property. (The AddressFamily for
// this address.)
const ADDRESS_FAMILY = "family"

// ADDRESS_NAME alias for the Address:name property. (The hostname for this
// address.)
const ADDRESS_NAME = "name"

// ADDRESS_PHYSICAL alias for the Address:physical property. (The stringified IP
// address for this address.)
const ADDRESS_PHYSICAL = "physical"

// ADDRESS_PORT alias for the Address:port property. (The port for this
// address.)
const ADDRESS_PORT = "port"

// ADDRESS_PROTOCOL alias for the Address:protocol property. (The URI scheme
// used with this address.)
const ADDRESS_PROTOCOL = "protocol"

// ADDRESS_SOCKADDR alias for the Address:sockaddr property. (A pointer to the
// struct sockaddr for this address.)
const ADDRESS_SOCKADDR = "sockaddr"

// AddressFamily: supported address families.
type AddressFamily int

const (
	// AddressFamilyInvalid SoupAddress
	AddressFamilyInvalid AddressFamily = -1
	// AddressFamilyIPv4: IPv4 address
	AddressFamilyIPv4 AddressFamily = 2
	// AddressFamilyIPv6: IPv6 address
	AddressFamilyIPv6 AddressFamily = 10
)

func marshalAddressFamily(p uintptr) (interface{}, error) {
	return AddressFamily(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for AddressFamily.
func (a AddressFamily) String() string {
	switch a {
	case AddressFamilyInvalid:
		return "Invalid"
	case AddressFamilyIPv4:
		return "IPv4"
	case AddressFamilyIPv6:
		return "IPv6"
	default:
		return fmt.Sprintf("AddressFamily(%d)", a)
	}
}

// AddressCallback: callback function passed to soup_address_resolve_async().
type AddressCallback func(addr *Address, status uint32)

//export _gotk4_soup2_AddressCallback
func _gotk4_soup2_AddressCallback(arg0 *C.SoupAddress, arg1 C.guint, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var addr *Address // out
	var status uint32 // out

	addr = wrapAddress(externglib.Take(unsafe.Pointer(arg0)))
	status = uint32(arg1)

	fn := v.(AddressCallback)
	fn(addr, status)
}

type Address struct {
	*externglib.Object

	gio.SocketConnectable
}

func wrapAddress(obj *externglib.Object) *Address {
	return &Address{
		Object: obj,
		SocketConnectable: gio.SocketConnectable{
			Object: obj,
		},
	}
}

func marshalAddresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAddress(obj), nil
}

// NewAddress creates a Address from name and port. The Address's IP address may
// not be available right away; the caller can call soup_address_resolve_async()
// or soup_address_resolve_sync() to force a DNS resolution.
func NewAddress(name string, port uint32) *Address {
	var _arg1 *C.char        // out
	var _arg2 C.guint        // out
	var _cret *C.SoupAddress // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(port)

	_cret = C.soup_address_new(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(port)

	var _address *Address // out

	_address = wrapAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _address
}

// NewAddressAny returns a Address corresponding to the "any" address for family
// (or NULL if family isn't supported), suitable for using as a listening
// Socket.
func NewAddressAny(family AddressFamily, port uint32) *Address {
	var _arg1 C.SoupAddressFamily // out
	var _arg2 C.guint             // out
	var _cret *C.SoupAddress      // in

	_arg1 = C.SoupAddressFamily(family)
	_arg2 = C.guint(port)

	_cret = C.soup_address_new_any(_arg1, _arg2)
	runtime.KeepAlive(family)
	runtime.KeepAlive(port)

	var _address *Address // out

	if _cret != nil {
		_address = wrapAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _address
}

// EqualByIP tests if addr1 and addr2 have the same IP address. This method can
// be used with soup_address_hash_by_ip() to create a Table that hashes on IP
// address.
//
// This would be used to distinguish hosts in situations where different virtual
// hosts on the same IP address should be considered the same. Eg, if
// "www.example.com" and "www.example.net" have the same IP address, then a
// single connection can be used to talk to either of them.
//
// See also soup_address_equal_by_name(), which compares by name rather than by
// IP address.
func (addr1 *Address) EqualByIP(addr2 *Address) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = C.gconstpointer(unsafe.Pointer(addr1.Native()))
	_arg1 = C.gconstpointer(unsafe.Pointer(addr2.Native()))

	_cret = C.soup_address_equal_by_ip(_arg0, _arg1)
	runtime.KeepAlive(addr1)
	runtime.KeepAlive(addr2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EqualByName tests if addr1 and addr2 have the same "name". This method can be
// used with soup_address_hash_by_name() to create a Table that hashes on
// address "names".
//
// Comparing by name normally means comparing the addresses by their hostnames.
// But if the address was originally created using an IP address literal, then
// it will be compared by that instead.
//
// In particular, if "www.example.com" has the IP address 10.0.0.1, and addr1
// was created with the name "www.example.com" and addr2 was created with the
// name "10.0.0.1", then they will compare as unequal for purposes of
// soup_address_equal_by_name().
//
// This would be used to distinguish hosts in situations where different virtual
// hosts on the same IP address should be considered different. Eg, for purposes
// of HTTP authentication or cookies, two hosts with the same IP address but
// different names are considered to be different hosts.
//
// See also soup_address_equal_by_ip(), which compares by IP address rather than
// by name.
func (addr1 *Address) EqualByName(addr2 *Address) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = C.gconstpointer(unsafe.Pointer(addr1.Native()))
	_arg1 = C.gconstpointer(unsafe.Pointer(addr2.Native()))

	_cret = C.soup_address_equal_by_name(_arg0, _arg1)
	runtime.KeepAlive(addr1)
	runtime.KeepAlive(addr2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Gsockaddr creates a new Address corresponding to addr (which is assumed to
// only have one socket address associated with it).
func (addr *Address) Gsockaddr() gio.SocketAddresser {
	var _arg0 *C.SoupAddress    // out
	var _cret *C.GSocketAddress // in

	_arg0 = (*C.SoupAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.soup_address_get_gsockaddr(_arg0)
	runtime.KeepAlive(addr)

	var _socketAddress gio.SocketAddresser // out

	_socketAddress = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.SocketAddresser)

	return _socketAddress
}

// Name returns the hostname associated with addr.
//
// This method is not thread-safe; if you call it while addr is being resolved
// in another thread, it may return garbage. You can use
// soup_address_is_resolved() to safely test whether or not an address is
// resolved before fetching its name or address.
func (addr *Address) Name() string {
	var _arg0 *C.SoupAddress // out
	var _cret *C.char        // in

	_arg0 = (*C.SoupAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.soup_address_get_name(_arg0)
	runtime.KeepAlive(addr)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Physical returns the physical address associated with addr as a string. (Eg,
// "127.0.0.1"). If the address is not yet known, returns NULL.
//
// This method is not thread-safe; if you call it while addr is being resolved
// in another thread, it may return garbage. You can use
// soup_address_is_resolved() to safely test whether or not an address is
// resolved before fetching its name or address.
func (addr *Address) Physical() string {
	var _arg0 *C.SoupAddress // out
	var _cret *C.char        // in

	_arg0 = (*C.SoupAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.soup_address_get_physical(_arg0)
	runtime.KeepAlive(addr)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Port returns the port associated with addr.
func (addr *Address) Port() uint32 {
	var _arg0 *C.SoupAddress // out
	var _cret C.guint        // in

	_arg0 = (*C.SoupAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.soup_address_get_port(_arg0)
	runtime.KeepAlive(addr)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// HashByIP: hash function (for Table) that corresponds to
// soup_address_equal_by_ip(), qv
func (addr *Address) HashByIP() uint32 {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = C.gconstpointer(unsafe.Pointer(addr.Native()))

	_cret = C.soup_address_hash_by_ip(_arg0)
	runtime.KeepAlive(addr)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// HashByName: hash function (for Table) that corresponds to
// soup_address_equal_by_name(), qv
func (addr *Address) HashByName() uint32 {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = C.gconstpointer(unsafe.Pointer(addr.Native()))

	_cret = C.soup_address_hash_by_name(_arg0)
	runtime.KeepAlive(addr)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// IsResolved tests if addr has already been resolved. Unlike the other Address
// "get" methods, this is safe to call when addr might be being resolved in
// another thread.
func (addr *Address) IsResolved() bool {
	var _arg0 *C.SoupAddress // out
	var _cret C.gboolean     // in

	_arg0 = (*C.SoupAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.soup_address_is_resolved(_arg0)
	runtime.KeepAlive(addr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResolveAsync: asynchronously resolves the missing half of addr (its IP
// address if it was created with soup_address_new(), or its hostname if it was
// created with soup_address_new_from_sockaddr() or soup_address_new_any().)
//
// If cancellable is non-NULL, it can be used to cancel the resolution. callback
// will still be invoked in this case, with a status of SOUP_STATUS_CANCELLED.
//
// It is safe to call this more than once on a given address, from the same
// thread, with the same async_context (and doing so will not result in
// redundant DNS queries being made). But it is not safe to call from multiple
// threads, or with different async_contexts, or mixed with calls to
// soup_address_resolve_sync().
func (addr *Address) ResolveAsync(ctx context.Context, asyncContext *glib.MainContext, callback AddressCallback) {
	var _arg0 *C.SoupAddress        // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GMainContext       // out
	var _arg3 C.SoupAddressCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.SoupAddress)(unsafe.Pointer(addr.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if asyncContext != nil {
		_arg1 = (*C.GMainContext)(gextras.StructNative(unsafe.Pointer(asyncContext)))
	}
	_arg3 = (*[0]byte)(C._gotk4_soup2_AddressCallback)
	_arg4 = C.gpointer(gbox.AssignOnce(callback))

	C.soup_address_resolve_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(addr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(asyncContext)
	runtime.KeepAlive(callback)
}

// ResolveSync: synchronously resolves the missing half of addr, as with
// soup_address_resolve_async().
//
// If cancellable is non-NULL, it can be used to cancel the resolution.
// soup_address_resolve_sync() will then return a status of
// SOUP_STATUS_CANCELLED.
//
// It is safe to call this more than once, even from different threads, but it
// is not safe to mix calls to soup_address_resolve_sync() with calls to
// soup_address_resolve_async() on the same address.
func (addr *Address) ResolveSync(ctx context.Context) uint32 {
	var _arg0 *C.SoupAddress  // out
	var _arg1 *C.GCancellable // out
	var _cret C.guint         // in

	_arg0 = (*C.SoupAddress)(unsafe.Pointer(addr.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.soup_address_resolve_sync(_arg0, _arg1)
	runtime.KeepAlive(addr)
	runtime.KeepAlive(ctx)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}
