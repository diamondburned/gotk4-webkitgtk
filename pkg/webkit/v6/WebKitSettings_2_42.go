// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

// SettingsGetAllFeatures gets the list of all available WebKit features.
//
// Features can be toggled with settings.SetFeatureEnabled, and their current
// state determined with settings.GetFeatureEnabled.
//
// Note that most applications should use settings.GetDevelopmentFeatures() and
// settings.GetExperimentalFeatures() instead.
//
// The function returns the following values:
//
//   - featureList: list of all features.
//
func SettingsGetAllFeatures() *FeatureList {
	var _cret *C.WebKitFeatureList // in

	_cret = C.webkit_settings_get_all_features()

	var _featureList *FeatureList // out

	_featureList = (*FeatureList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_featureList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_feature_list_unref((*C.WebKitFeatureList)(intern.C))
		},
	)

	return _featureList
}

// SettingsGetDevelopmentFeatures gets the list of available development WebKit
// features.
//
// The returned features are a subset of those returned by
// settings.GetAllFeatures(), and includes those which web and WebKit developers
// might find useful, but in general should *not* be exposed to end users;
// see featurestatus for more details.
//
// The function returns the following values:
//
//   - featureList: list of development features.
//
func SettingsGetDevelopmentFeatures() *FeatureList {
	var _cret *C.WebKitFeatureList // in

	_cret = C.webkit_settings_get_development_features()

	var _featureList *FeatureList // out

	_featureList = (*FeatureList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_featureList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_feature_list_unref((*C.WebKitFeatureList)(intern.C))
		},
	)

	return _featureList
}

// SettingsGetExperimentalFeatures gets the list of available experimental
// WebKit features.
//
// The returned features are a subset of those returned by
// settings.GetAllFeatures(), and includes those which certain applications may
// want to expose to end users; see featurestatus for more details.
//
// The function returns the following values:
//
//   - featureList: list of experimental features.
//
func SettingsGetExperimentalFeatures() *FeatureList {
	var _cret *C.WebKitFeatureList // in

	_cret = C.webkit_settings_get_experimental_features()

	var _featureList *FeatureList // out

	_featureList = (*FeatureList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_featureList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_feature_list_unref((*C.WebKitFeatureList)(intern.C))
		},
	)

	return _featureList
}
