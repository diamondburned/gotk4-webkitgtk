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

// glib.Type values for soup-websocket-extension-deflate.go.
var GTypeWebsocketExtensionDeflate = externglib.Type(C.soup_websocket_extension_deflate_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeWebsocketExtensionDeflate, F: marshalWebsocketExtensionDeflate},
	})
}

// WebsocketExtensionDeflateOverrider contains methods that are overridable.
type WebsocketExtensionDeflateOverrider interface {
}

type WebsocketExtensionDeflate struct {
	_ [0]func() // equal guard
	WebsocketExtension
}

var (
	_ WebsocketExtensioner = (*WebsocketExtensionDeflate)(nil)
)

func classInitWebsocketExtensionDeflater(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWebsocketExtensionDeflate(obj *externglib.Object) *WebsocketExtensionDeflate {
	return &WebsocketExtensionDeflate{
		WebsocketExtension: WebsocketExtension{
			Object: obj,
		},
	}
}

func marshalWebsocketExtensionDeflate(p uintptr) (interface{}, error) {
	return wrapWebsocketExtensionDeflate(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
