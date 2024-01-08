// Code generated by girgen. DO NOT EDIT.

package webkit

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
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypeEditorTypingAttributes = coreglib.Type(C.webkit_editor_typing_attributes_get_type())
	GTypeEditorState            = coreglib.Type(C.webkit_editor_state_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEditorTypingAttributes, F: marshalEditorTypingAttributes},
		coreglib.TypeMarshaler{T: GTypeEditorState, F: marshalEditorState},
	})
}

// EditorTypingAttributes: enum values with flags representing typing
// attributes.
type EditorTypingAttributes C.guint

const (
	// EditorTypingAttributeNone: no typing attributes.
	EditorTypingAttributeNone EditorTypingAttributes = 0b10
	// EditorTypingAttributeBold: bold typing attribute.
	EditorTypingAttributeBold EditorTypingAttributes = 0b100
	// EditorTypingAttributeItalic: italic typing attribute.
	EditorTypingAttributeItalic EditorTypingAttributes = 0b1000
	// EditorTypingAttributeUnderline: underline typing attribute.
	EditorTypingAttributeUnderline EditorTypingAttributes = 0b10000
	// EditorTypingAttributeStrikethrough: strikethrough typing attribute.
	EditorTypingAttributeStrikethrough EditorTypingAttributes = 0b100000
)

func marshalEditorTypingAttributes(p uintptr) (interface{}, error) {
	return EditorTypingAttributes(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for EditorTypingAttributes.
func (e EditorTypingAttributes) String() string {
	if e == 0 {
		return "EditorTypingAttributes(0)"
	}

	var builder strings.Builder
	builder.Grow(145)

	for e != 0 {
		next := e & (e - 1)
		bit := e - next

		switch bit {
		case EditorTypingAttributeNone:
			builder.WriteString("None|")
		case EditorTypingAttributeBold:
			builder.WriteString("Bold|")
		case EditorTypingAttributeItalic:
			builder.WriteString("Italic|")
		case EditorTypingAttributeUnderline:
			builder.WriteString("Underline|")
		case EditorTypingAttributeStrikethrough:
			builder.WriteString("Strikethrough|")
		default:
			builder.WriteString(fmt.Sprintf("EditorTypingAttributes(0b%b)|", bit))
		}

		e = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if e contains other.
func (e EditorTypingAttributes) Has(other EditorTypingAttributes) bool {
	return (e & other) == other
}

// EditorStateOverrides contains methods that are overridable.
type EditorStateOverrides struct {
}

func defaultEditorStateOverrides(v *EditorState) EditorStateOverrides {
	return EditorStateOverrides{}
}

// EditorState: web editor state.
//
// WebKitEditorState represents the state of a KitWebView editor.
// Use webkit_web_view_get_editor_state() to get the WebKitEditorState of a
// KitWebView.
type EditorState struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*EditorState)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EditorState, *EditorStateClass, EditorStateOverrides](
		GTypeEditorState,
		initEditorStateClass,
		wrapEditorState,
		defaultEditorStateOverrides,
	)
}

func initEditorStateClass(gclass unsafe.Pointer, overrides EditorStateOverrides, classInitFunc func(*EditorStateClass)) {
	if classInitFunc != nil {
		class := (*EditorStateClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEditorState(obj *coreglib.Object) *EditorState {
	return &EditorState{
		Object: obj,
	}
}

func marshalEditorState(p uintptr) (interface{}, error) {
	return wrapEditorState(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TypingAttributes gets the typing attributes at the current cursor position.
//
// If there is a selection, this returns the typing attributes of the selected
// text. Note that in case of a selection, typing attributes are considered
// active only when they are present throughout the selection.
//
// The function returns the following values:
//
//   - guint: bitmask of KitEditorTypingAttributes flags.
//
func (editorState *EditorState) TypingAttributes() uint {
	var _arg0 *C.WebKitEditorState // out
	var _cret C.guint              // in

	_arg0 = (*C.WebKitEditorState)(unsafe.Pointer(coreglib.InternObject(editorState).Native()))

	_cret = C.webkit_editor_state_get_typing_attributes(_arg0)
	runtime.KeepAlive(editorState)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IsCopyAvailable gets whether a copy command can be issued.
//
// The function returns the following values:
//
//   - ok: TRUE if copy is currently available.
//
func (editorState *EditorState) IsCopyAvailable() bool {
	var _arg0 *C.WebKitEditorState // out
	var _cret C.gboolean           // in

	_arg0 = (*C.WebKitEditorState)(unsafe.Pointer(coreglib.InternObject(editorState).Native()))

	_cret = C.webkit_editor_state_is_copy_available(_arg0)
	runtime.KeepAlive(editorState)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsCutAvailable gets whether a cut command can be issued.
//
// The function returns the following values:
//
//   - ok: TRUE if cut is currently available.
//
func (editorState *EditorState) IsCutAvailable() bool {
	var _arg0 *C.WebKitEditorState // out
	var _cret C.gboolean           // in

	_arg0 = (*C.WebKitEditorState)(unsafe.Pointer(coreglib.InternObject(editorState).Native()))

	_cret = C.webkit_editor_state_is_cut_available(_arg0)
	runtime.KeepAlive(editorState)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsPasteAvailable gets whether a paste command can be issued.
//
// The function returns the following values:
//
//   - ok: TRUE if paste is currently available.
//
func (editorState *EditorState) IsPasteAvailable() bool {
	var _arg0 *C.WebKitEditorState // out
	var _cret C.gboolean           // in

	_arg0 = (*C.WebKitEditorState)(unsafe.Pointer(coreglib.InternObject(editorState).Native()))

	_cret = C.webkit_editor_state_is_paste_available(_arg0)
	runtime.KeepAlive(editorState)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRedoAvailable gets whether a redo command can be issued.
//
// The function returns the following values:
//
//   - ok: TRUE if redo is currently available.
//
func (editorState *EditorState) IsRedoAvailable() bool {
	var _arg0 *C.WebKitEditorState // out
	var _cret C.gboolean           // in

	_arg0 = (*C.WebKitEditorState)(unsafe.Pointer(coreglib.InternObject(editorState).Native()))

	_cret = C.webkit_editor_state_is_redo_available(_arg0)
	runtime.KeepAlive(editorState)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsUndoAvailable gets whether an undo command can be issued.
//
// The function returns the following values:
//
//   - ok: TRUE if undo is currently available.
//
func (editorState *EditorState) IsUndoAvailable() bool {
	var _arg0 *C.WebKitEditorState // out
	var _cret C.gboolean           // in

	_arg0 = (*C.WebKitEditorState)(unsafe.Pointer(coreglib.InternObject(editorState).Native()))

	_cret = C.webkit_editor_state_is_undo_available(_arg0)
	runtime.KeepAlive(editorState)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
