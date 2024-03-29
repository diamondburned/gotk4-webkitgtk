// Code generated by girgen. DO NOT EDIT.

package javascriptcore

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <jsc/jsc.h>
// extern void _gotk4_javascriptcore4_WeakValue_ConnectCleared(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeWeakValue = coreglib.Type(C.jsc_weak_value_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWeakValue, F: marshalWeakValue},
	})
}

// WeakValueOverrides contains methods that are overridable.
type WeakValueOverrides struct {
}

func defaultWeakValueOverrides(v *WeakValue) WeakValueOverrides {
	return WeakValueOverrides{}
}

// WeakValue represents a weak reference to a value in a CContext. It can be
// used to keep a reference to a JavaScript value without protecting it from
// being garbage collected and without referencing the CContext either.
type WeakValue struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*WeakValue)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WeakValue, *WeakValueClass, WeakValueOverrides](
		GTypeWeakValue,
		initWeakValueClass,
		wrapWeakValue,
		defaultWeakValueOverrides,
	)
}

func initWeakValueClass(gclass unsafe.Pointer, overrides WeakValueOverrides, classInitFunc func(*WeakValueClass)) {
	if classInitFunc != nil {
		class := (*WeakValueClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWeakValue(obj *coreglib.Object) *WeakValue {
	return &WeakValue{
		Object: obj,
	}
}

func marshalWeakValue(p uintptr) (interface{}, error) {
	return wrapWeakValue(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectCleared: this signal is emitted when the JavaScript value is
// destroyed.
func (weakValue *WeakValue) ConnectCleared(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(weakValue, "cleared", false, unsafe.Pointer(C._gotk4_javascriptcore4_WeakValue_ConnectCleared), f)
}

// NewWeakValue: create a new CWeakValue for the JavaScript value referenced by
// value.
//
// The function takes the following parameters:
//
//   - value: CValue.
//
// The function returns the following values:
//
//   - weakValue: new CWeakValue.
//
func NewWeakValue(value *Value) *WeakValue {
	var _arg1 *C.JSCValue     // out
	var _cret *C.JSCWeakValue // in

	_arg1 = (*C.JSCValue)(unsafe.Pointer(coreglib.InternObject(value).Native()))

	_cret = C.jsc_weak_value_new(_arg1)
	runtime.KeepAlive(value)

	var _weakValue *WeakValue // out

	_weakValue = wrapWeakValue(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _weakValue
}

// Value: get a CValue referencing the JavaScript value of weak_value.
//
// The function returns the following values:
//
//   - value: new CValue or NULL if weak_value was cleared.
//
func (weakValue *WeakValue) Value() *Value {
	var _arg0 *C.JSCWeakValue // out
	var _cret *C.JSCValue     // in

	_arg0 = (*C.JSCWeakValue)(unsafe.Pointer(coreglib.InternObject(weakValue).Native()))

	_cret = C.jsc_weak_value_get_value(_arg0)
	runtime.KeepAlive(weakValue)

	var _value *Value // out

	_value = wrapValue(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _value
}

// WeakValueClass: instance of this type is always passed by reference.
type WeakValueClass struct {
	*weakValueClass
}

// weakValueClass is the struct that's finalized.
type weakValueClass struct {
	native *C.JSCWeakValueClass
}
