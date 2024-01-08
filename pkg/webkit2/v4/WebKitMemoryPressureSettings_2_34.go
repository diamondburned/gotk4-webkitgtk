// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// GType values.
var (
	GTypeMemoryPressureSettings = coreglib.Type(C.webkit_memory_pressure_settings_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMemoryPressureSettings, F: marshalMemoryPressureSettings},
	})
}

// MemoryPressureSettings: boxed type representing the settings for the memory
// pressure handler
//
// KitMemoryPressureSettings is a boxed type that can be used to provide some
// custom settings to control how the memory pressure situations are handled by
// the different processes.
//
// The memory pressure system implemented inside the different process will try
// to keep the memory usage under the defined memory limit. In order to do that,
// it will check the used memory with a user defined frequency and decide
// whether it should try to release memory. The thresholds passed will define
// how urgent is to release that memory.
//
// Take into account that badly defined parameters can greatly reduce the
// performance of the engine. For example, setting memory limit too low with a
// fast poll interval can cause the process to constantly be trying to release
// memory.
//
// A KitMemoryPressureSettings can be passed to a KitWebContext constructor,
// and the settings will be applied to all the web processes created by that
// context.
//
// A KitMemoryPressureSettings can be passed to
// webkit_website_data_manager_set_memory_pressure_settings(), and the settings
// will be applied to all the network processes created after that call by any
// instance of KitWebsiteDataManager.
//
// An instance of this type is always passed by reference.
type MemoryPressureSettings struct {
	*memoryPressureSettings
}

// memoryPressureSettings is the struct that's finalized.
type memoryPressureSettings struct {
	native *C.WebKitMemoryPressureSettings
}

func marshalMemoryPressureSettings(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &MemoryPressureSettings{&memoryPressureSettings{(*C.WebKitMemoryPressureSettings)(b)}}, nil
}

// NewMemoryPressureSettings constructs a struct MemoryPressureSettings.
func NewMemoryPressureSettings() *MemoryPressureSettings {
	var _cret *C.WebKitMemoryPressureSettings // in

	_cret = C.webkit_memory_pressure_settings_new()

	var _memoryPressureSettings *MemoryPressureSettings // out

	_memoryPressureSettings = (*MemoryPressureSettings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_memoryPressureSettings)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_memory_pressure_settings_free((*C.WebKitMemoryPressureSettings)(intern.C))
		},
	)

	return _memoryPressureSettings
}

// Copy: make a copy of settings.
//
// The function returns the following values:
//
//   - memoryPressureSettings: copy of of the passed KitMemoryPressureSettings.
//
func (settings *MemoryPressureSettings) Copy() *MemoryPressureSettings {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _cret *C.WebKitMemoryPressureSettings // in

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))

	_cret = C.webkit_memory_pressure_settings_copy(_arg0)
	runtime.KeepAlive(settings)

	var _memoryPressureSettings *MemoryPressureSettings // out

	_memoryPressureSettings = (*MemoryPressureSettings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_memoryPressureSettings)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_memory_pressure_settings_free((*C.WebKitMemoryPressureSettings)(intern.C))
		},
	)

	return _memoryPressureSettings
}

// ConservativeThreshold gets the conservative memory usage threshold.
//
// The function returns the following values:
//
//   - gdouble: value in the (0, 1) range.
//
func (settings *MemoryPressureSettings) ConservativeThreshold() float64 {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _cret C.gdouble                       // in

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))

	_cret = C.webkit_memory_pressure_settings_get_conservative_threshold(_arg0)
	runtime.KeepAlive(settings)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// KillThreshold gets the kill memory usage threshold.
//
// The function returns the following values:
//
//   - gdouble: positive value, can be zero.
//
func (settings *MemoryPressureSettings) KillThreshold() float64 {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _cret C.gdouble                       // in

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))

	_cret = C.webkit_memory_pressure_settings_get_kill_threshold(_arg0)
	runtime.KeepAlive(settings)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MemoryLimit gets the memory usage limit.
//
// The function returns the following values:
//
//   - guint: current value, in megabytes.
//
func (settings *MemoryPressureSettings) MemoryLimit() uint {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _cret C.guint                         // in

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))

	_cret = C.webkit_memory_pressure_settings_get_memory_limit(_arg0)
	runtime.KeepAlive(settings)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// PollInterval gets the interval at which memory usage is checked.
//
// The function returns the following values:
//
//   - gdouble: current interval value, in seconds.
//
func (settings *MemoryPressureSettings) PollInterval() float64 {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _cret C.gdouble                       // in

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))

	_cret = C.webkit_memory_pressure_settings_get_poll_interval(_arg0)
	runtime.KeepAlive(settings)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// StrictThreshold gets the strict memory usage threshold.
//
// The function returns the following values:
//
//   - gdouble: value in the (0, 1) range.
//
func (settings *MemoryPressureSettings) StrictThreshold() float64 {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _cret C.gdouble                       // in

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))

	_cret = C.webkit_memory_pressure_settings_get_strict_threshold(_arg0)
	runtime.KeepAlive(settings)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetConservativeThreshold sets the memory limit for the conservative policy to
// start working.
//
// Sets value as the fraction of the defined memory limit where the conservative
// policy starts working. This policy will try to reduce the memory footprint by
// releasing non critical memory.
//
// The threshold must be bigger than 0 and smaller than 1, and it must be
// smaller than the strict threshold defined in settings. The default value is
// 0.33.
//
// The function takes the following parameters:
//
//   - value: fraction of the memory limit where the conservative policy starts
//     working.
//
func (settings *MemoryPressureSettings) SetConservativeThreshold(value float64) {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _arg1 C.gdouble                       // out

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))
	_arg1 = C.gdouble(value)

	C.webkit_memory_pressure_settings_set_conservative_threshold(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(value)
}

// SetKillThreshold sets value as the fraction of the defined memory limit where
// the process will be killed.
//
// The threshold must be a value bigger or equal to 0. A value of 0 means that
// the process is never killed. If the threshold is not 0, then it must be
// bigger than the strict threshold defined in settings. The threshold can also
// have values bigger than 1. The default value is 0.
//
// The function takes the following parameters:
//
//   - value: fraction of the memory limit where the process will be killed
//     because of excessive memory usage.
//
func (settings *MemoryPressureSettings) SetKillThreshold(value float64) {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _arg1 C.gdouble                       // out

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))
	_arg1 = C.gdouble(value)

	C.webkit_memory_pressure_settings_set_kill_threshold(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(value)
}

// SetMemoryLimit sets memory_limit the memory limit value to settings.
//
// The default value is the system's RAM size with a maximum of 3GB.
//
// The function takes the following parameters:
//
//   - memoryLimit: amount of memory (in MB) that the process is allowed to use.
//
func (settings *MemoryPressureSettings) SetMemoryLimit(memoryLimit uint) {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _arg1 C.guint                         // out

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))
	_arg1 = C.guint(memoryLimit)

	C.webkit_memory_pressure_settings_set_memory_limit(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(memoryLimit)
}

// SetPollInterval sets value as the poll interval used by settings.
//
// The poll interval value must be bigger than 0. The default value is 30
// seconds.
//
// The function takes the following parameters:
//
//   - value: period (in seconds) between memory usage measurements.
//
func (settings *MemoryPressureSettings) SetPollInterval(value float64) {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _arg1 C.gdouble                       // out

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))
	_arg1 = C.gdouble(value)

	C.webkit_memory_pressure_settings_set_poll_interval(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(value)
}

// SetStrictThreshold sets the memory limit for the strict policy to start
// working.
//
// Sets value as the fraction of the defined memory limit where the strict
// policy starts working. This policy will try to reduce the memory footprint by
// releasing critical memory.
//
// The threshold must be bigger than 0 and smaller than 1. Also, it must be
// bigger than the conservative threshold defined in settings, and smaller than
// the kill threshold if the latter is not 0. The default value is 0.5.
//
// The function takes the following parameters:
//
//   - value: fraction of the memory limit where the strict policy starts
//     working.
//
func (settings *MemoryPressureSettings) SetStrictThreshold(value float64) {
	var _arg0 *C.WebKitMemoryPressureSettings // out
	var _arg1 C.gdouble                       // out

	_arg0 = (*C.WebKitMemoryPressureSettings)(gextras.StructNative(unsafe.Pointer(settings)))
	_arg1 = C.gdouble(value)

	C.webkit_memory_pressure_settings_set_strict_threshold(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(value)
}