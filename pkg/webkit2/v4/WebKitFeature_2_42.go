// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
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
	GTypeFeatureStatus = coreglib.Type(C.webkit_feature_status_get_type())
	GTypeFeature       = coreglib.Type(C.webkit_feature_get_type())
	GTypeFeatureList   = coreglib.Type(C.webkit_feature_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFeatureStatus, F: marshalFeatureStatus},
		coreglib.TypeMarshaler{T: GTypeFeature, F: marshalFeature},
		coreglib.TypeMarshaler{T: GTypeFeatureList, F: marshalFeatureList},
	})
}

// FeatureStatus describes the status of a webkitfeature.
//
// The status for a given feature can be obtained with
// webkit_feature_get_status.
type FeatureStatus C.gint

const (
	// FeatureStatusEmbedder: feature that adjust behaviour for specific
	// application needs. The feature is not part of a Web platform feature,
	// not a mature feature intended to be always on.
	FeatureStatusEmbedder FeatureStatus = iota
	// FeatureStatusUnstable: feature in development. The feature may be
	// unfinished, and there are no guarantees about its safety and stability.
	FeatureStatusUnstable
	// FeatureStatusInternal: feature for debugging the WebKit engine. The
	// feature is not generally useful for user or web developers, and always
	// disabled by default.
	FeatureStatusInternal
	// FeatureStatusDeveloper: feature for web developers. The feature is not
	// generally useful for end users, and always disabled by default.
	FeatureStatusDeveloper
	// FeatureStatusTestable: feature in active development and complete enough
	// for testing. The feature may not be yet ready to ship and is disabled by
	// default.
	FeatureStatusTestable
	// FeatureStatusPreview: feature ready to be tested by users. The feature
	// is disabled by default, but may be enabled by applications automatically
	// e.g. in their “technology preview” or “beta” versions.
	FeatureStatusPreview
	// FeatureStatusStable: feature ready for general use. The feature is
	// enabled by default, but it may still be toggled to support debugging and
	// testing.
	FeatureStatusStable
	// FeatureStatusMature: feature in general use. The feature is always
	// enabled and in general there should be no user-facing interface to toggle
	// it.
	FeatureStatusMature
)

func marshalFeatureStatus(p uintptr) (interface{}, error) {
	return FeatureStatus(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FeatureStatus.
func (f FeatureStatus) String() string {
	switch f {
	case FeatureStatusEmbedder:
		return "Embedder"
	case FeatureStatusUnstable:
		return "Unstable"
	case FeatureStatusInternal:
		return "Internal"
	case FeatureStatusDeveloper:
		return "Developer"
	case FeatureStatusTestable:
		return "Testable"
	case FeatureStatusPreview:
		return "Preview"
	case FeatureStatusStable:
		return "Stable"
	case FeatureStatusMature:
		return "Mature"
	default:
		return fmt.Sprintf("FeatureStatus(%d)", f)
	}
}

// Feature describes a web engine feature that may be toggled at runtime.
//
// The WebKit web engine includes a set of features which may be toggled
// programmatically, each one represented by a KitFeature that provides
// information about it:
//
// - A unique “identifier”: feature.GetIdentifier.
//
// - A “default value”, which indicates whether the option is enabled
// automatically: feature.GetDefaultValue.
//
// - Its “status”, which determines whether it should be considered
// user-settable and its development stage (see featurestatus for details):
// feature.GetStatus.
//
// - A category, which may be used to group features together:
// feature.GetCategory.
//
// - An optional short “name” which can be presented to an user:
// feature.GetName.
//
// - An optional longer “detailed” description: feature.GetDetails.
//
// The lists of available features can be obtained with
// settings.GetAllFeatures(), settings.GetExperimentalFeatures(), and
// settings.GetDevelopmentFeatures()). As a rule of thumb, applications which
// may want to allow users (i.e. web developers) to test WebKit features should
// use the list of experimental features. Additionally, applications might want
// to expose development features *when targeting technically inclined users*
// for early testing of in-development features (i.e. in “technology preview” or
// “canary” builds).
//
// Applications **must not** expose the list of all features to end users
// because they often lack descriptions and control parts of the web engine
// which are either intended to be used during development of WebKit itself,
// or in specific scenarios to tweak how WebKit integrates with the application.
//
// An instance of this type is always passed by reference.
type Feature struct {
	*feature
}

// feature is the struct that's finalized.
type feature struct {
	native *C.WebKitFeature
}

func marshalFeature(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Feature{&feature{(*C.WebKitFeature)(b)}}, nil
}

// Category gets the category of the feature.
//
// Applications which include user interface to toggle features may want to use
// the category to group related features together.
//
// The function returns the following values:
//
//   - utf8: feature category.
//
func (feature *Feature) Category() string {
	var _arg0 *C.WebKitFeature // out
	var _cret *C.char          // in

	_arg0 = (*C.WebKitFeature)(gextras.StructNative(unsafe.Pointer(feature)))

	_cret = C.webkit_feature_get_category(_arg0)
	runtime.KeepAlive(feature)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// DefaultValue gets whether the feature is enabled by default.
//
// The default value may be used by applications which include user
// interface to toggle features to restore its settings to their defaults.
// Note that whether a feature is actually enabled must be checked with
// settings.GetFeatureEnabled.
//
// The function returns the following values:
//
//   - ok: whether the feature is enabled by default.
//
func (feature *Feature) DefaultValue() bool {
	var _arg0 *C.WebKitFeature // out
	var _cret C.gboolean       // in

	_arg0 = (*C.WebKitFeature)(gextras.StructNative(unsafe.Pointer(feature)))

	_cret = C.webkit_feature_get_default_value(_arg0)
	runtime.KeepAlive(feature)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Details gets a description for the feature.
//
// The detailed description should be considered an additional clarification on
// the purpose of the feature, to be used as complementary aid to be displayed
// along the feature name returned by feature.GetName. The returned string is
// suitable to be displayed to end users, but it should not be relied upon being
// localized.
//
// Note that some *features may not* have a detailed description, and NULL is
// returned in this case.
//
// The function returns the following values:
//
//   - utf8 (optional): feature description.
//
func (feature *Feature) Details() string {
	var _arg0 *C.WebKitFeature // out
	var _cret *C.char          // in

	_arg0 = (*C.WebKitFeature)(gextras.StructNative(unsafe.Pointer(feature)))

	_cret = C.webkit_feature_get_details(_arg0)
	runtime.KeepAlive(feature)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Identifier gets a string that uniquely identifies the feature.
//
// The function returns the following values:
//
//   - utf8: identifier string for the feature.
//
func (feature *Feature) Identifier() string {
	var _arg0 *C.WebKitFeature // out
	var _cret *C.char          // in

	_arg0 = (*C.WebKitFeature)(gextras.StructNative(unsafe.Pointer(feature)))

	_cret = C.webkit_feature_get_identifier(_arg0)
	runtime.KeepAlive(feature)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name gets a short name for the feature.
//
// The returned string is suitable to be displayed to end users, but it should
// not be relied upon being localized.
//
// Note that some *features may not* have a short name, and NULL is returned in
// this case.
//
// The function returns the following values:
//
//   - utf8 (optional): short feature name.
//
func (feature *Feature) Name() string {
	var _arg0 *C.WebKitFeature // out
	var _cret *C.char          // in

	_arg0 = (*C.WebKitFeature)(gextras.StructNative(unsafe.Pointer(feature)))

	_cret = C.webkit_feature_get_name(_arg0)
	runtime.KeepAlive(feature)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Status gets the status of the feature.
//
// The function returns the following values:
//
//   - featureStatus: feature status.
//
func (feature *Feature) Status() FeatureStatus {
	var _arg0 *C.WebKitFeature      // out
	var _cret C.WebKitFeatureStatus // in

	_arg0 = (*C.WebKitFeature)(gextras.StructNative(unsafe.Pointer(feature)))

	_cret = C.webkit_feature_get_status(_arg0)
	runtime.KeepAlive(feature)

	var _featureStatus FeatureStatus // out

	_featureStatus = FeatureStatus(_cret)

	return _featureStatus
}

// FeatureList contains a set of toggle-able web engine features.
//
// The list supports passing around a set of feature objects and iterating over
// them:
//
//    g_autoptr(WebKitFeatureList) list = webkit_settings_get_experimental_features();
//    for (gsize i = 0; i < webkit_feature_list_get_length(list): i++) {
//        WebKitFeature *feature = webkit_feature_list_get(list, i);
//        // Do something with "feature".
//    }
//
// Lists of features can be obtained with settings.GetExperimentalFeatures(),
// settings.GetDevelopmentFeatures(), and settings.GetAllFeatures().
//
// An instance of this type is always passed by reference.
type FeatureList struct {
	*featureList
}

// featureList is the struct that's finalized.
type featureList struct {
	native *C.WebKitFeatureList
}

func marshalFeatureList(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &FeatureList{&featureList{(*C.WebKitFeatureList)(b)}}, nil
}

// Get gets a feature given its index.
//
// The function takes the following parameters:
//
//   - index of the feature.
//
// The function returns the following values:
//
//   - feature at index.
//
func (featureList *FeatureList) Get(index uint) *Feature {
	var _arg0 *C.WebKitFeatureList // out
	var _arg1 C.gsize              // out
	var _cret *C.WebKitFeature     // in

	_arg0 = (*C.WebKitFeatureList)(gextras.StructNative(unsafe.Pointer(featureList)))
	_arg1 = C.gsize(index)

	_cret = C.webkit_feature_list_get(_arg0, _arg1)
	runtime.KeepAlive(featureList)
	runtime.KeepAlive(index)

	var _feature *Feature // out

	_feature = (*Feature)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.webkit_feature_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_feature)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_feature_unref((*C.WebKitFeature)(intern.C))
		},
	)

	return _feature
}

// Length gets the number of elements in the feature list.
//
// The function returns the following values:
//
//   - gsize: number of elements.
//
//     Since 2.42.
//
func (featureList *FeatureList) Length() uint {
	var _arg0 *C.WebKitFeatureList // out
	var _cret C.gsize              // in

	_arg0 = (*C.WebKitFeatureList)(gextras.StructNative(unsafe.Pointer(featureList)))

	_cret = C.webkit_feature_list_get_length(_arg0)
	runtime.KeepAlive(featureList)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}