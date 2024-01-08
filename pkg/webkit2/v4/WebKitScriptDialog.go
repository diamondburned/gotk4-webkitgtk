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
	GTypeScriptDialogType = coreglib.Type(C.webkit_script_dialog_type_get_type())
	GTypeScriptDialog     = coreglib.Type(C.webkit_script_dialog_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScriptDialogType, F: marshalScriptDialogType},
		coreglib.TypeMarshaler{T: GTypeScriptDialog, F: marshalScriptDialog},
	})
}

// ScriptDialogType: enum values used for determining the type of
// KitScriptDialog.
type ScriptDialogType C.gint

const (
	// ScriptDialogAlert: alert script dialog, used to show a message to the
	// user.
	ScriptDialogAlert ScriptDialogType = iota
	// ScriptDialogConfirm: confirm script dialog, used to ask confirmation to
	// the user.
	ScriptDialogConfirm
	// ScriptDialogPrompt: prompt script dialog, used to ask information to the
	// user.
	ScriptDialogPrompt
	// ScriptDialogBeforeUnloadConfirm: before unload confirm dialog, used to
	// ask confirmation to leave the current page to the user. Since 2.12.
	ScriptDialogBeforeUnloadConfirm
)

func marshalScriptDialogType(p uintptr) (interface{}, error) {
	return ScriptDialogType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ScriptDialogType.
func (s ScriptDialogType) String() string {
	switch s {
	case ScriptDialogAlert:
		return "Alert"
	case ScriptDialogConfirm:
		return "Confirm"
	case ScriptDialogPrompt:
		return "Prompt"
	case ScriptDialogBeforeUnloadConfirm:
		return "BeforeUnloadConfirm"
	default:
		return fmt.Sprintf("ScriptDialogType(%d)", s)
	}
}

// ScriptDialog carries details to be shown in user-facing dialogs.
//
// An instance of this type is always passed by reference.
type ScriptDialog struct {
	*scriptDialog
}

// scriptDialog is the struct that's finalized.
type scriptDialog struct {
	native *C.WebKitScriptDialog
}

func marshalScriptDialog(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ScriptDialog{&scriptDialog{(*C.WebKitScriptDialog)(b)}}, nil
}

// Close dialog.
//
// When handling a KitScriptDialog asynchronously (webkit_script_dialog_ref()
// was called in KitWebView::script-dialog callback), this function needs to be
// called to notify that we are done with the script dialog. The dialog will be
// closed on destruction if this function hasn't been called before.
func (dialog *ScriptDialog) Close() {
	var _arg0 *C.WebKitScriptDialog // out

	_arg0 = (*C.WebKitScriptDialog)(gextras.StructNative(unsafe.Pointer(dialog)))

	C.webkit_script_dialog_close(_arg0)
	runtime.KeepAlive(dialog)
}

// ConfirmSetConfirmed: set whether the user confirmed the dialog.
//
// This method is used for WEBKIT_SCRIPT_DIALOG_CONFIRM and
// WEBKIT_SCRIPT_DIALOG_BEFORE_UNLOAD_CONFIRM dialogs when
// KitWebView::script-dialog signal is emitted to set whether the user confirmed
// the dialog or not. The default implementation of KitWebView::script-dialog
// signal sets TRUE when the OK or Stay buttons are clicked and FALSE otherwise.
// It's an error to use this method with a KitScriptDialog that is not of type
// WEBKIT_SCRIPT_DIALOG_CONFIRM or WEBKIT_SCRIPT_DIALOG_BEFORE_UNLOAD_CONFIRM.
//
// The function takes the following parameters:
//
//   - confirmed: whether user confirmed the dialog.
//
func (dialog *ScriptDialog) ConfirmSetConfirmed(confirmed bool) {
	var _arg0 *C.WebKitScriptDialog // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.WebKitScriptDialog)(gextras.StructNative(unsafe.Pointer(dialog)))
	if confirmed {
		_arg1 = C.TRUE
	}

	C.webkit_script_dialog_confirm_set_confirmed(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(confirmed)
}

// DialogType: get the dialog type of a KitScriptDialog.
//
// The function returns the following values:
//
//   - scriptDialogType of dialog.
//
func (dialog *ScriptDialog) DialogType() ScriptDialogType {
	var _arg0 *C.WebKitScriptDialog    // out
	var _cret C.WebKitScriptDialogType // in

	_arg0 = (*C.WebKitScriptDialog)(gextras.StructNative(unsafe.Pointer(dialog)))

	_cret = C.webkit_script_dialog_get_dialog_type(_arg0)
	runtime.KeepAlive(dialog)

	var _scriptDialogType ScriptDialogType // out

	_scriptDialogType = ScriptDialogType(_cret)

	return _scriptDialogType
}

// Message: get the message of a KitScriptDialog.
//
// The function returns the following values:
//
//   - utf8: message of dialog.
//
func (dialog *ScriptDialog) Message() string {
	var _arg0 *C.WebKitScriptDialog // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitScriptDialog)(gextras.StructNative(unsafe.Pointer(dialog)))

	_cret = C.webkit_script_dialog_get_message(_arg0)
	runtime.KeepAlive(dialog)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PromptGetDefaultText: get the default text of a KitScriptDialog of type
// WEBKIT_SCRIPT_DIALOG_PROMPT.
//
// It's an error to use this method with a KitScriptDialog that is not of type
// WEBKIT_SCRIPT_DIALOG_PROMPT.
//
// The function returns the following values:
//
//   - utf8: default text of dialog.
//
func (dialog *ScriptDialog) PromptGetDefaultText() string {
	var _arg0 *C.WebKitScriptDialog // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitScriptDialog)(gextras.StructNative(unsafe.Pointer(dialog)))

	_cret = C.webkit_script_dialog_prompt_get_default_text(_arg0)
	runtime.KeepAlive(dialog)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PromptSetText: set the text entered by the user in the dialog.
//
// This method is used for WEBKIT_SCRIPT_DIALOG_PROMPT dialogs when
// KitWebView::script-dialog signal is emitted to set the text entered by the
// user. The default implementation of KitWebView::script-dialog signal sets
// the text of the entry form when OK button is clicked, otherwise NULL is set.
// It's an error to use this method with a KitScriptDialog that is not of type
// WEBKIT_SCRIPT_DIALOG_PROMPT.
//
// The function takes the following parameters:
//
//   - text to set.
//
func (dialog *ScriptDialog) PromptSetText(text string) {
	var _arg0 *C.WebKitScriptDialog // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.WebKitScriptDialog)(gextras.StructNative(unsafe.Pointer(dialog)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_script_dialog_prompt_set_text(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(text)
}
