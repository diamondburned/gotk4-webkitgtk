// Code generated by girgen. DO NOT EDIT.

package webkitwebprocessextension

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit-web-process-extension.h>
import "C"

// GType values.
var (
	GTypeHitTestResult = coreglib.Type(C.webkit_hit_test_result_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHitTestResult, F: marshalHitTestResult},
	})
}

// HitTestResultContext: enum values with flags representing the context of a
// KitHitTestResult.
type HitTestResultContext C.guint

const (
	// HitTestResultContextDocument: anywhere in the document.
	HitTestResultContextDocument HitTestResultContext = 0b10
	// HitTestResultContextLink: hyperlink element.
	HitTestResultContextLink HitTestResultContext = 0b100
	// HitTestResultContextImage: image element.
	HitTestResultContextImage HitTestResultContext = 0b1000
	// HitTestResultContextMedia: video or audio element.
	HitTestResultContextMedia HitTestResultContext = 0b10000
	// HitTestResultContextEditable: editable element.
	HitTestResultContextEditable HitTestResultContext = 0b100000
	// HitTestResultContextScrollbar: scrollbar element.
	HitTestResultContextScrollbar HitTestResultContext = 0b1000000
	// HitTestResultContextSelection: selected element. Since 2.8.
	HitTestResultContextSelection HitTestResultContext = 0b10000000
)

// String returns the names in string for HitTestResultContext.
func (h HitTestResultContext) String() string {
	if h == 0 {
		return "HitTestResultContext(0)"
	}

	var builder strings.Builder
	builder.Grow(194)

	for h != 0 {
		next := h & (h - 1)
		bit := h - next

		switch bit {
		case HitTestResultContextDocument:
			builder.WriteString("Document|")
		case HitTestResultContextLink:
			builder.WriteString("Link|")
		case HitTestResultContextImage:
			builder.WriteString("Image|")
		case HitTestResultContextMedia:
			builder.WriteString("Media|")
		case HitTestResultContextEditable:
			builder.WriteString("Editable|")
		case HitTestResultContextScrollbar:
			builder.WriteString("Scrollbar|")
		case HitTestResultContextSelection:
			builder.WriteString("Selection|")
		default:
			builder.WriteString(fmt.Sprintf("HitTestResultContext(0b%b)|", bit))
		}

		h = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if h contains other.
func (h HitTestResultContext) Has(other HitTestResultContext) bool {
	return (h & other) == other
}

// HitTestResultOverrides contains methods that are overridable.
type HitTestResultOverrides struct {
}

func defaultHitTestResultOverrides(v *HitTestResult) HitTestResultOverrides {
	return HitTestResultOverrides{}
}

// HitTestResult: result of a Hit Test.
//
// A Hit Test is an operation to get context information about a given point
// in a KitWebView. KitHitTestResult represents the result of a Hit Test.
// It provides context information about what is at the coordinates of the Hit
// Test, such as if there's a link, an image or a media.
//
// You can get the context of the HitTestResult with
// webkit_hit_test_result_get_context() that returns a
// bitmask of KitHitTestResultContext flags. You can
// also use webkit_hit_test_result_context_is_link(),
// webkit_hit_test_result_context_is_image() and
// webkit_hit_test_result_context_is_media() to determine whether there's a
// link, image or a media element at the coordinates of the Hit Test. Note that
// it's possible that several KitHitTestResultContext flags are active at the
// same time, for example if there's a link containing an image.
//
// When the mouse is moved over a KitWebView a Hit Test is performed for the
// mouse coordinates and KitWebView::mouse-target-changed signal is emitted with
// a KitHitTestResult.
type HitTestResult struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*HitTestResult)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*HitTestResult, *HitTestResultClass, HitTestResultOverrides](
		GTypeHitTestResult,
		initHitTestResultClass,
		wrapHitTestResult,
		defaultHitTestResultOverrides,
	)
}

func initHitTestResultClass(gclass unsafe.Pointer, overrides HitTestResultOverrides, classInitFunc func(*HitTestResultClass)) {
	if classInitFunc != nil {
		class := (*HitTestResultClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapHitTestResult(obj *coreglib.Object) *HitTestResult {
	return &HitTestResult{
		Object: obj,
	}
}

func marshalHitTestResult(p uintptr) (interface{}, error) {
	return wrapHitTestResult(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ContextIsEditable gets whether WEBKIT_HIT_TEST_RESULT_CONTEXT_EDITABLE flag
// is present in KitHitTestResult:context.
//
// The function returns the following values:
//
//   - ok: TRUE if there's an editable element at the coordinates of the
//     hit_test_result, or FALSE otherwise.
//
func (hitTestResult *HitTestResult) ContextIsEditable() bool {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_context_is_editable(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContextIsImage gets whether WEBKIT_HIT_TEST_RESULT_CONTEXT_IMAGE flag is
// present in KitHitTestResult:context.
//
// The function returns the following values:
//
//   - ok: TRUE if there's an image element in the coordinates of the Hit Test,
//     or FALSE otherwise.
//
func (hitTestResult *HitTestResult) ContextIsImage() bool {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_context_is_image(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContextIsLink gets whether WEBKIT_HIT_TEST_RESULT_CONTEXT_LINK flag is
// present in KitHitTestResult:context.
//
// The function returns the following values:
//
//   - ok: TRUE if there's a link element in the coordinates of the Hit Test,
//     or FALSE otherwise.
//
func (hitTestResult *HitTestResult) ContextIsLink() bool {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_context_is_link(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContextIsMedia gets whether WEBKIT_HIT_TEST_RESULT_CONTEXT_MEDIA flag is
// present in KitHitTestResult:context.
//
// The function returns the following values:
//
//   - ok: TRUE if there's a media element in the coordinates of the Hit Test,
//     or FALSE otherwise.
//
func (hitTestResult *HitTestResult) ContextIsMedia() bool {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_context_is_media(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContextIsScrollbar gets whether WEBKIT_HIT_TEST_RESULT_CONTEXT_SCROLLBAR flag
// is present in KitHitTestResult:context.
//
// The function returns the following values:
//
//   - ok: TRUE if there's a scrollbar element at the coordinates of the
//     hit_test_result, or FALSE otherwise.
//
func (hitTestResult *HitTestResult) ContextIsScrollbar() bool {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_context_is_scrollbar(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContextIsSelection gets whether WEBKIT_HIT_TEST_RESULT_CONTEXT_SELECTION flag
// is present in KitHitTestResult:context.
//
// The function returns the following values:
//
//   - ok: TRUE if there's a selected element at the coordinates of the
//     hit_test_result, or FALSE otherwise.
//
func (hitTestResult *HitTestResult) ContextIsSelection() bool {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_context_is_selection(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Context gets the value of the KitHitTestResult:context property.
//
// The function returns the following values:
//
//   - guint: bitmask of KitHitTestResultContext flags.
//
func (hitTestResult *HitTestResult) Context() uint {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret C.guint                // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_get_context(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ImageURI gets the value of the KitHitTestResult:image-uri property.
//
// The function returns the following values:
//
//   - utf8: URI of the image element in the coordinates of the Hit Test,
//     or NULL if there isn't an image element in hit_test_result context.
//
func (hitTestResult *HitTestResult) ImageURI() string {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_get_image_uri(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LinkLabel gets the value of the KitHitTestResult:link-label property.
//
// The function returns the following values:
//
//   - utf8: label of the link element in the coordinates of the Hit Test,
//     or NULL if there isn't a link element in hit_test_result context or the
//     link element doesn't have a label.
//
func (hitTestResult *HitTestResult) LinkLabel() string {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_get_link_label(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LinkTitle gets the value of the KitHitTestResult:link-title property.
//
// The function returns the following values:
//
//   - utf8: title of the link element in the coordinates of the Hit Test,
//     or NULL if there isn't a link element in hit_test_result context or the
//     link element doesn't have a title.
//
func (hitTestResult *HitTestResult) LinkTitle() string {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_get_link_title(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LinkURI gets the value of the KitHitTestResult:link-uri property.
//
// The function returns the following values:
//
//   - utf8: URI of the link element in the coordinates of the Hit Test,
//     or NULL if there isn't a link element in hit_test_result context.
//
func (hitTestResult *HitTestResult) LinkURI() string {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_get_link_uri(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// MediaURI gets the value of the KitHitTestResult:media-uri property.
//
// The function returns the following values:
//
//   - utf8: URI of the media element in the coordinates of the Hit Test,
//     or NULL if there isn't a media element in hit_test_result context.
//
func (hitTestResult *HitTestResult) MediaURI() string {
	var _arg0 *C.WebKitHitTestResult // out
	var _cret *C.gchar               // in

	_arg0 = (*C.WebKitHitTestResult)(unsafe.Pointer(coreglib.InternObject(hitTestResult).Native()))

	_cret = C.webkit_hit_test_result_get_media_uri(_arg0)
	runtime.KeepAlive(hitTestResult)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HitTestResultClass: instance of this type is always passed by reference.
type HitTestResultClass struct {
	*hitTestResultClass
}

// hitTestResultClass is the struct that's finalized.
type hitTestResultClass struct {
	native *C.WebKitHitTestResultClass
}
