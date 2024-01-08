// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypeWebExtensionMode = coreglib.Type(C.webkit_web_extension_mode_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWebExtensionMode, F: marshalWebExtensionMode},
	})
}

// WebExtensionMode: enum values used for setting if a KitWebView is intended
// for WebExtensions.
type WebExtensionMode C.gint

const (
	// WebExtensionModeNone: not for an extension.
	WebExtensionModeNone WebExtensionMode = iota
	// WebExtensionModeManifestv2: for a ManifestV2 extension.
	WebExtensionModeManifestv2
	// WebExtensionModeManifestv3: for a ManifestV3 extension.
	WebExtensionModeManifestv3
)

func marshalWebExtensionMode(p uintptr) (interface{}, error) {
	return WebExtensionMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for WebExtensionMode.
func (w WebExtensionMode) String() string {
	switch w {
	case WebExtensionModeNone:
		return "None"
	case WebExtensionModeManifestv2:
		return "Manifestv2"
	case WebExtensionModeManifestv3:
		return "Manifestv3"
	default:
		return fmt.Sprintf("WebExtensionMode(%d)", w)
	}
}