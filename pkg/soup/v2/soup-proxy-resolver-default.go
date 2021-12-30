// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_proxy_resolver_default_get_type()), F: marshalProxyResolverDefaulter},
	})
}

type ProxyResolverDefault struct {
	_ [0]func() // equal guard
	*externglib.Object

	ProxyURIResolver
}

var (
	_ externglib.Objector = (*ProxyResolverDefault)(nil)
)

func wrapProxyResolverDefault(obj *externglib.Object) *ProxyResolverDefault {
	return &ProxyResolverDefault{
		Object: obj,
		ProxyURIResolver: ProxyURIResolver{
			SessionFeature: SessionFeature{
				Object: obj,
			},
		},
	}
}

func marshalProxyResolverDefaulter(p uintptr) (interface{}, error) {
	return wrapProxyResolverDefault(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
