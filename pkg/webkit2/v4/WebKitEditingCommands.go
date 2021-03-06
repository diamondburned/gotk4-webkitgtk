// Code generated by girgen. DO NOT EDIT.

package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"

// EDITING_COMMAND_COPY: copy clipboard command. Copies the current selection
// inside a KitWebView to the clipboard. You can check whether it's possible to
// execute the command with webkit_web_view_can_execute_editing_command(). In
// general it's possible to copy to the clipboard when there is an active
// selection inside the KitWebView.
const EDITING_COMMAND_COPY = "Copy"

// EDITING_COMMAND_CREATE_LINK: create link command. Creates a link element that
// is inserted at the current cursor position. If there's a selection, the
// selected text will be used as the link text, otherwise the URL itself will be
// used. It receives the link URL as argument. This command should be executed
// with webkit_web_view_execute_editing_command_with_argument().
const EDITING_COMMAND_CREATE_LINK = "CreateLink"

// EDITING_COMMAND_CUT: cut clipboard command. Copies the current selection
// inside a KitWebView to the clipboard and deletes the selected content. You
// can check whether it's possible to execute the command with
// webkit_web_view_can_execute_editing_command(). In general it's possible to
// cut to the clipboard when the KitWebView content is editable and there is an
// active selection.
const EDITING_COMMAND_CUT = "Cut"

// EDITING_COMMAND_INSERT_IMAGE: insert image command. Creates an image element
// that is inserted at the current cursor position. It receives an URI as
// argument, that is used as the image source. This command should be executed
// with webkit_web_view_execute_editing_command_with_argument().
const EDITING_COMMAND_INSERT_IMAGE = "InsertImage"

// EDITING_COMMAND_PASTE: paste clipboard command. Pastes the contents of the
// clipboard to a KitWebView. You can check whether it's possible to execute the
// command with webkit_web_view_can_execute_editing_command(). In general it's
// possible to paste from the clipboard when the KitWebView content is editable
// and clipboard is not empty.
const EDITING_COMMAND_PASTE = "Paste"

// EDITING_COMMAND_PASTE_AS_PLAIN_TEXT: paste as plaintext clipboard command.
// Pastes the contents of the clipboard to a KitWebView, with formatting
// removed. You can check whether it's possible to execute the command with
// webkit_web_view_can_execute_editing_command(). In general it's possible to
// paste from the clipboard when the KitWebView content is editable and
// clipboard is not empty.
const EDITING_COMMAND_PASTE_AS_PLAIN_TEXT = "PasteAsPlainText"

// EDITING_COMMAND_REDO: redo command. Redoes a previously undone editing
// command in a KitWebView. You can check whether it's possible to execute the
// command with webkit_web_view_can_execute_editing_command(). It's only
// possible to redo a command when it has been previously undone.
const EDITING_COMMAND_REDO = "Redo"

// EDITING_COMMAND_SELECT_ALL: select all command. Selects all the content of
// the current text field in a KitWebView. It is always possible to select all
// text, no matter whether the KitWebView content is editable or not. You can
// still check it with webkit_web_view_can_execute_editing_command().
const EDITING_COMMAND_SELECT_ALL = "SelectAll"

// EDITING_COMMAND_UNDO: undo command. Undoes the last editing command in a
// KitWebView. You can check whether it's possible to execute the command with
// webkit_web_view_can_execute_editing_command(). It's only possible to undo a
// command after a previously executed editing operation.
const EDITING_COMMAND_UNDO = "Undo"
