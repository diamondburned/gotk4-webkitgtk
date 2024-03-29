// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

// WebsocketExtensionClass class structure for the SoupWebsocketExtension.
//
// An instance of this type is always passed by reference.
type WebsocketExtensionClass struct {
	*websocketExtensionClass
}

// websocketExtensionClass is the struct that's finalized.
type websocketExtensionClass struct {
	native *C.SoupWebsocketExtensionClass
}

func (w *WebsocketExtensionClass) Name() string {
	valptr := &w.native.name
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}
