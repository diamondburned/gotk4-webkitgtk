// Code generated by girgen. DO NOT EDIT.

package javascriptcore

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: javascriptcoregtk-4.0 webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <jsc/jsc.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.jsc_virtual_machine_get_type()), F: marshalVirtualMachiner},
	})
}

type VirtualMachine struct {
	*externglib.Object
}

func wrapVirtualMachine(obj *externglib.Object) *VirtualMachine {
	return &VirtualMachine{
		Object: obj,
	}
}

func marshalVirtualMachiner(p uintptr) (interface{}, error) {
	return wrapVirtualMachine(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVirtualMachine: create a new CVirtualMachine.
func NewVirtualMachine() *VirtualMachine {
	var _cret *C.JSCVirtualMachine // in

	_cret = C.jsc_virtual_machine_new()

	var _virtualMachine *VirtualMachine // out

	_virtualMachine = wrapVirtualMachine(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _virtualMachine
}
