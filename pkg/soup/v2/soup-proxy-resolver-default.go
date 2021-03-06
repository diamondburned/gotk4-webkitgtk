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

// glib.Type values for soup-proxy-resolver-default.go.
var GTypeProxyResolverDefault = externglib.Type(C.soup_proxy_resolver_default_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeProxyResolverDefault, F: marshalProxyResolverDefault},
	})
}

// ProxyResolverDefaultOverrider contains methods that are overridable.
type ProxyResolverDefaultOverrider interface {
}

type ProxyResolverDefault struct {
	_ [0]func() // equal guard
	*externglib.Object

	ProxyURIResolver
}

var (
	_ externglib.Objector = (*ProxyResolverDefault)(nil)
)

func classInitProxyResolverDefaulter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

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

func marshalProxyResolverDefault(p uintptr) (interface{}, error) {
	return wrapProxyResolverDefault(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
