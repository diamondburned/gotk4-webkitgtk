// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern guint _gotk4_soup2_ProxyResolverInterface_get_proxy_sync(SoupProxyResolver*, SoupMessage*, GCancellable*, SoupAddress**);
// extern void _gotk4_soup2_ProxyResolverCallback(SoupProxyResolver*, SoupMessage*, guint, SoupAddress*, gpointer);
import "C"

// glib.Type values for soup-proxy-resolver.go.
var GTypeProxyResolver = externglib.Type(C.soup_proxy_resolver_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeProxyResolver, F: marshalProxyResolver},
	})
}

// ProxyResolverCallback: deprecated: Use SoupProxyURIResolver instead.
type ProxyResolverCallback func(proxyResolver ProxyResolverer, msg *Message, arg uint, addr *Address)

//export _gotk4_soup2_ProxyResolverCallback
func _gotk4_soup2_ProxyResolverCallback(arg1 *C.SoupProxyResolver, arg2 *C.SoupMessage, arg3 C.guint, arg4 *C.SoupAddress, arg5 C.gpointer) {
	var fn ProxyResolverCallback
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ProxyResolverCallback)
	}

	var _proxyResolver ProxyResolverer // out
	var _msg *Message                  // out
	var _arg uint                      // out
	var _addr *Address                 // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type soup.ProxyResolverer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(ProxyResolverer)
			return ok
		})
		rv, ok := casted.(ProxyResolverer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching soup.ProxyResolverer")
		}
		_proxyResolver = rv
	}
	_msg = wrapMessage(externglib.Take(unsafe.Pointer(arg2)))
	_arg = uint(arg3)
	_addr = wrapAddress(externglib.Take(unsafe.Pointer(arg4)))

	fn(_proxyResolver, _msg, _arg, _addr)
}

//
// ProxyResolver wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ProxyResolver struct {
	_ [0]func() // equal guard
	SessionFeature
}

var ()

// ProxyResolverer describes ProxyResolver's interface methods.
type ProxyResolverer interface {
	externglib.Objector

	// ProxyAsync: deprecated: Use SoupProxyURIResolver.get_proxy_uri_async
	// instead.
	ProxyAsync(ctx context.Context, msg *Message, asyncContext *glib.MainContext, callback ProxyResolverCallback)
	// ProxySync: deprecated: Use SoupProxyURIResolver.get_proxy_uri_sync()
	// instead.
	ProxySync(ctx context.Context, msg *Message) (*Address, uint)
}

var _ ProxyResolverer = (*ProxyResolver)(nil)

func wrapProxyResolver(obj *externglib.Object) *ProxyResolver {
	return &ProxyResolver{
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalProxyResolver(p uintptr) (interface{}, error) {
	return wrapProxyResolver(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ProxyAsync: deprecated: Use SoupProxyURIResolver.get_proxy_uri_async instead.
//
// The function takes the following parameters:
//
//    - ctx (optional)
//    - msg
//    - asyncContext
//    - callback
//
func (proxyResolver *ProxyResolver) ProxyAsync(ctx context.Context, msg *Message, asyncContext *glib.MainContext, callback ProxyResolverCallback) {
	var _arg0 *C.SoupProxyResolver        // out
	var _arg3 *C.GCancellable             // out
	var _arg1 *C.SoupMessage              // out
	var _arg2 *C.GMainContext             // out
	var _arg4 C.SoupProxyResolverCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.SoupProxyResolver)(unsafe.Pointer(externglib.InternObject(proxyResolver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(externglib.InternObject(msg).Native()))
	_arg2 = (*C.GMainContext)(gextras.StructNative(unsafe.Pointer(asyncContext)))
	_arg4 = (*[0]byte)(C._gotk4_soup2_ProxyResolverCallback)
	_arg5 = C.gpointer(gbox.AssignOnce(callback))

	C.soup_proxy_resolver_get_proxy_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(proxyResolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(asyncContext)
	runtime.KeepAlive(callback)
}

// ProxySync: deprecated: Use SoupProxyURIResolver.get_proxy_uri_sync() instead.
//
// The function takes the following parameters:
//
//    - ctx (optional)
//    - msg
//
// The function returns the following values:
//
//    - addr
//    - guint
//
func (proxyResolver *ProxyResolver) ProxySync(ctx context.Context, msg *Message) (*Address, uint) {
	var _arg0 *C.SoupProxyResolver // out
	var _arg2 *C.GCancellable      // out
	var _arg1 *C.SoupMessage       // out
	var _arg3 *C.SoupAddress       // in
	var _cret C.guint              // in

	_arg0 = (*C.SoupProxyResolver)(unsafe.Pointer(externglib.InternObject(proxyResolver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(externglib.InternObject(msg).Native()))

	_cret = C.soup_proxy_resolver_get_proxy_sync(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(proxyResolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(msg)

	var _addr *Address // out
	var _guint uint    // out

	_addr = wrapAddress(externglib.Take(unsafe.Pointer(_arg3)))
	_guint = uint(_cret)

	return _addr, _guint
}
