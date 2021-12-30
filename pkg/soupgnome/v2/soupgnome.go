// Code generated by girgen. DO NOT EDIT.

package soupgnome

import (
	_ "runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4-webkitgtk/pkg/soup/v2"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: libsoup-gnome-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup-gnome.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_password_manager_gnome_get_type()), F: marshalPasswordManagerGNOMEer},
		{T: externglib.Type(C.soup_proxy_resolver_gnome_get_type()), F: marshalProxyResolverGNOMEer},
	})
}

type PasswordManagerGNOME struct {
	_ [0]func() // equal guard
	*externglib.Object

	soup.SessionFeature
}

var (
	_ externglib.Objector = (*PasswordManagerGNOME)(nil)
)

func wrapPasswordManagerGNOME(obj *externglib.Object) *PasswordManagerGNOME {
	return &PasswordManagerGNOME{
		Object: obj,
		SessionFeature: soup.SessionFeature{
			Object: obj,
		},
	}
}

func marshalPasswordManagerGNOMEer(p uintptr) (interface{}, error) {
	return wrapPasswordManagerGNOME(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

type ProxyResolverGNOME struct {
	_ [0]func() // equal guard
	soup.ProxyResolverDefault
}

var (
	_ externglib.Objector = (*ProxyResolverGNOME)(nil)
)

func wrapProxyResolverGNOME(obj *externglib.Object) *ProxyResolverGNOME {
	return &ProxyResolverGNOME{
		ProxyResolverDefault: soup.ProxyResolverDefault{
			Object: obj,
			ProxyURIResolver: soup.ProxyURIResolver{
				SessionFeature: soup.SessionFeature{
					Object: obj,
				},
			},
		},
	}
}

func marshalProxyResolverGNOMEer(p uintptr) (interface{}, error) {
	return wrapProxyResolverGNOME(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
