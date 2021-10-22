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

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
// void _gotk4_soup2_ProxyURIResolverCallback(SoupProxyURIResolver*, guint, SoupURI*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_proxy_uri_resolver_get_type()), F: marshalProxyURIResolverer},
	})
}

// ProxyURIResolverCallback: callback for
// soup_proxy_uri_resolver_get_proxy_uri_async().
type ProxyURIResolverCallback func(resolver ProxyURIResolverer, status uint, proxyUri *URI)

//export _gotk4_soup2_ProxyURIResolverCallback
func _gotk4_soup2_ProxyURIResolverCallback(arg0 *C.SoupProxyURIResolver, arg1 C.guint, arg2 *C.SoupURI, arg3 C.gpointer) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var resolver ProxyURIResolverer // out
	var status uint                 // out
	var proxyUri *URI               // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type soup.ProxyURIResolverer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(ProxyURIResolverer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not soup.ProxyURIResolverer")
		}
		resolver = rv
	}
	status = uint(arg1)
	proxyUri = (*URI)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	fn := v.(ProxyURIResolverCallback)
	fn(resolver, status, proxyUri)
}

// ProxyURIResolverOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ProxyURIResolverOverrider interface {
	// ProxyUriAsync: asynchronously determines a proxy URI to use for msg and
	// calls callback.
	//
	// Deprecated: ProxyURIResolver is deprecated in favor of Resolver.
	ProxyUriAsync(ctx context.Context, uri *URI, asyncContext *glib.MainContext, callback ProxyURIResolverCallback)
	// ProxyUriSync: synchronously determines a proxy URI to use for uri. If uri
	// should be sent via proxy, *proxy_uri will be set to the URI of the proxy,
	// else it will be set to NULL.
	//
	// Deprecated: ProxyURIResolver is deprecated in favor of Resolver.
	ProxyUriSync(ctx context.Context, uri *URI) (*URI, uint)
}

type ProxyURIResolver struct {
	SessionFeature
}

// ProxyURIResolverer describes ProxyURIResolver's interface methods.
type ProxyURIResolverer interface {
	externglib.Objector

	// ProxyUriAsync: asynchronously determines a proxy URI to use for msg and
	// calls callback.
	ProxyUriAsync(ctx context.Context, uri *URI, asyncContext *glib.MainContext, callback ProxyURIResolverCallback)
	// ProxyUriSync: synchronously determines a proxy URI to use for uri.
	ProxyUriSync(ctx context.Context, uri *URI) (*URI, uint)
}

var _ ProxyURIResolverer = (*ProxyURIResolver)(nil)

func wrapProxyURIResolver(obj *externglib.Object) *ProxyURIResolver {
	return &ProxyURIResolver{
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalProxyURIResolverer(p uintptr) (interface{}, error) {
	return wrapProxyURIResolver(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ProxyUriAsync: asynchronously determines a proxy URI to use for msg and calls
// callback.
//
// Deprecated: ProxyURIResolver is deprecated in favor of Resolver.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - uri you want a proxy for.
//    - asyncContext to invoke callback in.
//    - callback to invoke with the proxy address.
//
func (proxyUriResolver *ProxyURIResolver) ProxyUriAsync(ctx context.Context, uri *URI, asyncContext *glib.MainContext, callback ProxyURIResolverCallback) {
	var _arg0 *C.SoupProxyURIResolver        // out
	var _arg3 *C.GCancellable                // out
	var _arg1 *C.SoupURI                     // out
	var _arg2 *C.GMainContext                // out
	var _arg4 C.SoupProxyURIResolverCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.SoupProxyURIResolver)(unsafe.Pointer(proxyUriResolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))
	if asyncContext != nil {
		_arg2 = (*C.GMainContext)(gextras.StructNative(unsafe.Pointer(asyncContext)))
	}
	_arg4 = (*[0]byte)(C._gotk4_soup2_ProxyURIResolverCallback)
	_arg5 = C.gpointer(gbox.AssignOnce(callback))

	C.soup_proxy_uri_resolver_get_proxy_uri_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(proxyUriResolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(asyncContext)
	runtime.KeepAlive(callback)
}

// ProxyUriSync: synchronously determines a proxy URI to use for uri. If uri
// should be sent via proxy, *proxy_uri will be set to the URI of the proxy,
// else it will be set to NULL.
//
// Deprecated: ProxyURIResolver is deprecated in favor of Resolver.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - uri you want a proxy for.
//
func (proxyUriResolver *ProxyURIResolver) ProxyUriSync(ctx context.Context, uri *URI) (*URI, uint) {
	var _arg0 *C.SoupProxyURIResolver // out
	var _arg2 *C.GCancellable         // out
	var _arg1 *C.SoupURI              // out
	var _arg3 *C.SoupURI              // in
	var _cret C.guint                 // in

	_arg0 = (*C.SoupProxyURIResolver)(unsafe.Pointer(proxyUriResolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.SoupURI)(gextras.StructNative(unsafe.Pointer(uri)))

	_cret = C.soup_proxy_uri_resolver_get_proxy_uri_sync(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(proxyUriResolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)

	var _proxyUri *URI // out
	var _guint uint    // out

	_proxyUri = (*URI)(gextras.NewStructNative(unsafe.Pointer(_arg3)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_proxyUri)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_uri_free((*C.SoupURI)(intern.C))
		},
	)
	_guint = uint(_cret)

	return _proxyUri, _guint
}
