// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// GType values.
var (
	GTypeWebsocketExtensionManager = coreglib.Type(C.soup_websocket_extension_manager_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWebsocketExtensionManager, F: marshalWebsocketExtensionManager},
	})
}

// WebsocketExtensionManagerOverrides contains methods that are overridable.
type WebsocketExtensionManagerOverrides struct {
}

func defaultWebsocketExtensionManagerOverrides(v *WebsocketExtensionManager) WebsocketExtensionManagerOverrides {
	return WebsocketExtensionManagerOverrides{}
}

// WebsocketExtensionManager is the sessionfeature that handles WebSockets
// extensions for a session.
//
// A WebsocketExtensionManager is added to the session by default, and normally
// you don't need to worry about it at all. However, if you want to disable
// WebSocket extensions, you can remove the feature from the session with
// session.RemoveFeatureByType or disable it on individual requests with
// message.DisableFeature.
type WebsocketExtensionManager struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SessionFeature
}

var (
	_ coreglib.Objector = (*WebsocketExtensionManager)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WebsocketExtensionManager, *WebsocketExtensionManagerClass, WebsocketExtensionManagerOverrides](
		GTypeWebsocketExtensionManager,
		initWebsocketExtensionManagerClass,
		wrapWebsocketExtensionManager,
		defaultWebsocketExtensionManagerOverrides,
	)
}

func initWebsocketExtensionManagerClass(gclass unsafe.Pointer, overrides WebsocketExtensionManagerOverrides, classInitFunc func(*WebsocketExtensionManagerClass)) {
	if classInitFunc != nil {
		class := (*WebsocketExtensionManagerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWebsocketExtensionManager(obj *coreglib.Object) *WebsocketExtensionManager {
	return &WebsocketExtensionManager{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalWebsocketExtensionManager(p uintptr) (interface{}, error) {
	return wrapWebsocketExtensionManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// WebsocketExtensionManagerClass: instance of this type is always passed by
// reference.
type WebsocketExtensionManagerClass struct {
	*websocketExtensionManagerClass
}

// websocketExtensionManagerClass is the struct that's finalized.
type websocketExtensionManagerClass struct {
	native *C.SoupWebsocketExtensionManagerClass
}
