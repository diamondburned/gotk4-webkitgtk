// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// GType values.
var (
	GTypePermissionState = coreglib.Type(C.webkit_permission_state_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePermissionState, F: marshalPermissionState},
	})
}

// PermissionState: enum values representing query permission results.
type PermissionState C.gint

const (
	// PermissionStateGranted access to the feature is granted.
	PermissionStateGranted PermissionState = iota
	// PermissionStateDenied access to the feature is denied.
	PermissionStateDenied
	// PermissionStatePrompt access to the feature has to be requested via user
	// prompt.
	PermissionStatePrompt
)

func marshalPermissionState(p uintptr) (interface{}, error) {
	return PermissionState(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PermissionState.
func (p PermissionState) String() string {
	switch p {
	case PermissionStateGranted:
		return "Granted"
	case PermissionStateDenied:
		return "Denied"
	case PermissionStatePrompt:
		return "Prompt"
	default:
		return fmt.Sprintf("PermissionState(%d)", p)
	}
}
