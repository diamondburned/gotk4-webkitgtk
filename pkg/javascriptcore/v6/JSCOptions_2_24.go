// Code generated by girgen. DO NOT EDIT.

package javascriptcore

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <jsc/jsc.h>
// extern gboolean _gotk4_javascriptcore6_OptionsFunc(char*, JSCOptionType, char*, gpointer);
import "C"

// OPTIONS_USE_DFG allows the DFG JIT to be used if TRUE. Option type:
// JSC_OPTION_BOOLEAN Default value: TRUE.
const OPTIONS_USE_DFG = "useDFGJIT"

// OPTIONS_USE_FTL allows the FTL JIT to be used if TRUE. Option type:
// JSC_OPTION_BOOLEAN Default value: TRUE.
const OPTIONS_USE_FTL = "useFTLJIT"

// OPTIONS_USE_JIT allows the executable pages to be allocated for JIT and
// thunks if TRUE. Option type: JSC_OPTION_BOOLEAN Default value: TRUE.
const OPTIONS_USE_JIT = "useJIT"

// OPTIONS_USE_LLINT allows the LLINT to be used if TRUE. Option type:
// JSC_OPTION_BOOLEAN Default value: TRUE.
const OPTIONS_USE_LLINT = "useLLInt"

// OptionType: enum values for options types.
type OptionType C.gint

const (
	// OptionBoolean option type.
	OptionBoolean OptionType = iota
	// OptionInt option type.
	OptionInt
	// OptionUint option type.
	OptionUint
	// OptionSize options type.
	OptionSize
	// OptionDouble options type.
	OptionDouble
	// OptionString: string option type.
	OptionString
	// OptionRangeString: range string option type.
	OptionRangeString
)

// String returns the name in string for OptionType.
func (o OptionType) String() string {
	switch o {
	case OptionBoolean:
		return "Boolean"
	case OptionInt:
		return "Int"
	case OptionUint:
		return "Uint"
	case OptionSize:
		return "Size"
	case OptionDouble:
		return "Double"
	case OptionString:
		return "String"
	case OptionRangeString:
		return "RangeString"
	default:
		return fmt.Sprintf("OptionType(%d)", o)
	}
}

// OptionsFunc: function used to iterate options.
//
// Not that description string is not localized.
type OptionsFunc func(option string, typ OptionType, description string) (ok bool)

// OptionsForEach iterates all available options calling function for each one.
// Iteration can stop early if function returns FALSE.
//
// The function takes the following parameters:
//
//   - function: COptionsFunc callback.
//
func OptionsForEach(function OptionsFunc) {
	var _arg1 C.JSCOptionsFunc // out
	var _arg2 C.gpointer

	_arg1 = (*[0]byte)(C._gotk4_javascriptcore6_OptionsFunc)
	_arg2 = C.gpointer(gbox.Assign(function))
	defer gbox.Delete(uintptr(_arg2))

	C.jsc_options_foreach(_arg1, _arg2)
	runtime.KeepAlive(function)
}

// OptionsGetBoolean: get option as a #gboolean value.
//
// The function takes the following parameters:
//
//   - option identifier.
//
// The function returns the following values:
//
//   - value: return location for the option value.
//   - ok: TRUE if value has been set or FALSE if the option doesn't exist.
//
func OptionsGetBoolean(option string) (value, ok bool) {
	var _arg1 *C.char    // out
	var _arg2 C.gboolean // in
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.jsc_options_get_boolean(_arg1, &_arg2)
	runtime.KeepAlive(option)

	var _value bool // out
	var _ok bool    // out

	if _arg2 != 0 {
		_value = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// OptionsGetDouble: get option as a #gdouble value.
//
// The function takes the following parameters:
//
//   - option identifier.
//
// The function returns the following values:
//
//   - value: return location for the option value.
//   - ok: TRUE if value has been set or FALSE if the option doesn't exist.
//
func OptionsGetDouble(option string) (float64, bool) {
	var _arg1 *C.char    // out
	var _arg2 C.gdouble  // in
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.jsc_options_get_double(_arg1, &_arg2)
	runtime.KeepAlive(option)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// OptionsGetInt: get option as a #gint value.
//
// The function takes the following parameters:
//
//   - option identifier.
//
// The function returns the following values:
//
//   - value: return location for the option value.
//   - ok: TRUE if value has been set or FALSE if the option doesn't exist.
//
func OptionsGetInt(option string) (int, bool) {
	var _arg1 *C.char    // out
	var _arg2 C.gint     // in
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.jsc_options_get_int(_arg1, &_arg2)
	runtime.KeepAlive(option)

	var _value int // out
	var _ok bool   // out

	_value = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// OptionsGetOptionGroup: create a Group to handle JSCOptions as command line
// arguments. The options will be exposed as command line arguments with the
// form <emphasis>--jsc-&lt;option&gt;=&lt;value&gt;</emphasis>. Each entry
// in the returned Group is configured to apply the corresponding option
// during command line parsing. Applications only need to pass the returned
// group to g_option_context_add_group(), and the rest will be taken care for
// automatically.
//
// The function returns the following values:
//
//   - optionGroup for the JSCOptions.
//
func OptionsGetOptionGroup() *glib.OptionGroup {
	var _cret *C.GOptionGroup // in

	_cret = C.jsc_options_get_option_group()

	var _optionGroup *glib.OptionGroup // out

	_optionGroup = (*glib.OptionGroup)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_optionGroup)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_option_group_unref((*C.GOptionGroup)(intern.C))
		},
	)

	return _optionGroup
}

// OptionsGetRangeString: get option as a range string. The string must be in
// the format <emphasis>[!]&lt;low&gt;[:&lt;high&gt;]</emphasis> where low and
// high are #guint values. Values between low and high (both included) will be
// considered in the range, unless <emphasis>!</emphasis> is used to invert the
// range.
//
// The function takes the following parameters:
//
//   - option identifier.
//
// The function returns the following values:
//
//   - value: return location for the option value.
//   - ok: TRUE if value has been set or FALSE if the option doesn't exist.
//
func OptionsGetRangeString(option string) (string, bool) {
	var _arg1 *C.char    // out
	var _arg2 *C.char    // in
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.jsc_options_get_range_string(_arg1, &_arg2)
	runtime.KeepAlive(option)

	var _value string // out
	var _ok bool      // out

	_value = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// OptionsGetSize: get option as a #gsize value.
//
// The function takes the following parameters:
//
//   - option identifier.
//
// The function returns the following values:
//
//   - value: return location for the option value.
//   - ok: TRUE if value has been set or FALSE if the option doesn't exist.
//
func OptionsGetSize(option string) (uint, bool) {
	var _arg1 *C.char    // out
	var _arg2 C.gsize    // in
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.jsc_options_get_size(_arg1, &_arg2)
	runtime.KeepAlive(option)

	var _value uint // out
	var _ok bool    // out

	_value = uint(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// OptionsGetString: get option as a string.
//
// The function takes the following parameters:
//
//   - option identifier.
//
// The function returns the following values:
//
//   - value: return location for the option value.
//   - ok: TRUE if value has been set or FALSE if the option doesn't exist.
//
func OptionsGetString(option string) (string, bool) {
	var _arg1 *C.char    // out
	var _arg2 *C.char    // in
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.jsc_options_get_string(_arg1, &_arg2)
	runtime.KeepAlive(option)

	var _value string // out
	var _ok bool      // out

	_value = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// OptionsGetUint: get option as a #guint value.
//
// The function takes the following parameters:
//
//   - option identifier.
//
// The function returns the following values:
//
//   - value: return location for the option value.
//   - ok: TRUE if value has been set or FALSE if the option doesn't exist.
//
func OptionsGetUint(option string) (uint, bool) {
	var _arg1 *C.char    // out
	var _arg2 C.guint    // in
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.jsc_options_get_uint(_arg1, &_arg2)
	runtime.KeepAlive(option)

	var _value uint // out
	var _ok bool    // out

	_value = uint(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// OptionsSetBoolean: set option as a #gboolean value.
//
// The function takes the following parameters:
//
//   - option identifier.
//   - value to set.
//
// The function returns the following values:
//
//   - ok: TRUE if option was correctly set or FALSE otherwise.
//
func OptionsSetBoolean(option string, value bool) bool {
	var _arg1 *C.char    // out
	var _arg2 C.gboolean // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))
	if value {
		_arg2 = C.TRUE
	}

	_cret = C.jsc_options_set_boolean(_arg1, _arg2)
	runtime.KeepAlive(option)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OptionsSetDouble: set option as a #gdouble value.
//
// The function takes the following parameters:
//
//   - option identifier.
//   - value to set.
//
// The function returns the following values:
//
//   - ok: TRUE if option was correctly set or FALSE otherwise.
//
func OptionsSetDouble(option string, value float64) bool {
	var _arg1 *C.char    // out
	var _arg2 C.gdouble  // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(value)

	_cret = C.jsc_options_set_double(_arg1, _arg2)
	runtime.KeepAlive(option)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OptionsSetInt: set option as a #gint value.
//
// The function takes the following parameters:
//
//   - option identifier.
//   - value to set.
//
// The function returns the following values:
//
//   - ok: TRUE if option was correctly set or FALSE otherwise.
//
func OptionsSetInt(option string, value int) bool {
	var _arg1 *C.char    // out
	var _arg2 C.gint     // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(value)

	_cret = C.jsc_options_set_int(_arg1, _arg2)
	runtime.KeepAlive(option)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OptionsSetRangeString: set option as a range string. The string must be in
// the format <emphasis>[!]&lt;low&gt;[:&lt;high&gt;]</emphasis> where low and
// high are #guint values. Values between low and high (both included) will be
// considered in the range, unless <emphasis>!</emphasis> is used to invert the
// range.
//
// The function takes the following parameters:
//
//   - option identifier.
//   - value to set.
//
// The function returns the following values:
//
//   - ok: TRUE if option was correctly set or FALSE otherwise.
//
func OptionsSetRangeString(option, value string) bool {
	var _arg1 *C.char    // out
	var _arg2 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.jsc_options_set_range_string(_arg1, _arg2)
	runtime.KeepAlive(option)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OptionsSetSize: set option as a #gsize value.
//
// The function takes the following parameters:
//
//   - option identifier.
//   - value to set.
//
// The function returns the following values:
//
//   - ok: TRUE if option was correctly set or FALSE otherwise.
//
func OptionsSetSize(option string, value uint) bool {
	var _arg1 *C.char    // out
	var _arg2 C.gsize    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gsize(value)

	_cret = C.jsc_options_set_size(_arg1, _arg2)
	runtime.KeepAlive(option)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OptionsSetString: set option as a string.
//
// The function takes the following parameters:
//
//   - option identifier.
//   - value to set.
//
// The function returns the following values:
//
//   - ok: TRUE if option was correctly set or FALSE otherwise.
//
func OptionsSetString(option, value string) bool {
	var _arg1 *C.char    // out
	var _arg2 *C.char    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.jsc_options_set_string(_arg1, _arg2)
	runtime.KeepAlive(option)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OptionsSetUint: set option as a #guint value.
//
// The function takes the following parameters:
//
//   - option identifier.
//   - value to set.
//
// The function returns the following values:
//
//   - ok: TRUE if option was correctly set or FALSE otherwise.
//
func OptionsSetUint(option string, value uint) bool {
	var _arg1 *C.char    // out
	var _arg2 C.guint    // out
	var _cret C.gboolean // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(value)

	_cret = C.jsc_options_set_uint(_arg1, _arg2)
	runtime.KeepAlive(option)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}