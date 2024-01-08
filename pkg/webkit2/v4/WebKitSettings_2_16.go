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
	GTypeHardwareAccelerationPolicy = coreglib.Type(C.webkit_hardware_acceleration_policy_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHardwareAccelerationPolicy, F: marshalHardwareAccelerationPolicy},
	})
}

// HardwareAccelerationPolicy: enum values used for determining the hardware
// acceleration policy.
type HardwareAccelerationPolicy C.gint

const (
	// HardwareAccelerationPolicyOnDemand: hardware acceleration is
	// enabled/disabled as request by web contents.
	HardwareAccelerationPolicyOnDemand HardwareAccelerationPolicy = iota
	// HardwareAccelerationPolicyAlways: hardware acceleration is always
	// enabled, even for websites not requesting it.
	HardwareAccelerationPolicyAlways
	// HardwareAccelerationPolicyNever: hardware acceleration is always
	// disabled, even for websites requesting it.
	HardwareAccelerationPolicyNever
)

func marshalHardwareAccelerationPolicy(p uintptr) (interface{}, error) {
	return HardwareAccelerationPolicy(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for HardwareAccelerationPolicy.
func (h HardwareAccelerationPolicy) String() string {
	switch h {
	case HardwareAccelerationPolicyOnDemand:
		return "OnDemand"
	case HardwareAccelerationPolicyAlways:
		return "Always"
	case HardwareAccelerationPolicyNever:
		return "Never"
	default:
		return fmt.Sprintf("HardwareAccelerationPolicy(%d)", h)
	}
}
