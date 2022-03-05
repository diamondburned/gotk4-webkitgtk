// Code generated by girgen. DO NOT EDIT.

package javascriptcore

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <jsc/jsc.h>
import "C"

// glib.Type values for JSCVirtualMachine.go.
var GTypeVirtualMachine = externglib.Type(C.jsc_virtual_machine_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeVirtualMachine, F: marshalVirtualMachine},
	})
}

// VirtualMachineOverrider contains methods that are overridable.
type VirtualMachineOverrider interface {
}

type VirtualMachine struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*VirtualMachine)(nil)
)

func classInitVirtualMachiner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapVirtualMachine(obj *externglib.Object) *VirtualMachine {
	return &VirtualMachine{
		Object: obj,
	}
}

func marshalVirtualMachine(p uintptr) (interface{}, error) {
	return wrapVirtualMachine(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVirtualMachine: create a new CVirtualMachine.
//
// The function returns the following values:
//
//    - virtualMachine: newly created CVirtualMachine.
//
func NewVirtualMachine() *VirtualMachine {
	var _cret *C.JSCVirtualMachine // in

	_cret = C.jsc_virtual_machine_new()

	var _virtualMachine *VirtualMachine // out

	_virtualMachine = wrapVirtualMachine(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _virtualMachine
}
