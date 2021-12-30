// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_context_menu_action_get_type()), F: marshalContextMenuAction},
	})
}

// ContextMenuAction: enum values used to denote the stock actions for
// KitContextMenuItem<!-- -->s.
type ContextMenuAction C.gint

const (
	// ContextMenuActionNoAction: no action, used by separator menu items.
	ContextMenuActionNoAction ContextMenuAction = 0
	// ContextMenuActionOpenLink: open current link.
	ContextMenuActionOpenLink ContextMenuAction = 1
	// ContextMenuActionOpenLinkInNewWindow: open current link in a new window.
	ContextMenuActionOpenLinkInNewWindow ContextMenuAction = 2
	// ContextMenuActionDownloadLinkToDisk: download link destination.
	ContextMenuActionDownloadLinkToDisk ContextMenuAction = 3
	// ContextMenuActionCopyLinkToClipboard: copy link location to the
	// clipboard.
	ContextMenuActionCopyLinkToClipboard ContextMenuAction = 4
	// ContextMenuActionOpenImageInNewWindow: open current image in a new
	// window.
	ContextMenuActionOpenImageInNewWindow ContextMenuAction = 5
	// ContextMenuActionDownloadImageToDisk: download current image.
	ContextMenuActionDownloadImageToDisk ContextMenuAction = 6
	// ContextMenuActionCopyImageToClipboard: copy current image to the
	// clipboard.
	ContextMenuActionCopyImageToClipboard ContextMenuAction = 7
	// ContextMenuActionCopyImageURLToClipboard: copy current image location to
	// the clipboard.
	ContextMenuActionCopyImageURLToClipboard ContextMenuAction = 8
	// ContextMenuActionOpenFrameInNewWindow: open current frame in a new
	// window.
	ContextMenuActionOpenFrameInNewWindow ContextMenuAction = 9
	// ContextMenuActionGoBack: load the previous history item.
	ContextMenuActionGoBack ContextMenuAction = 10
	// ContextMenuActionGoForward: load the next history item.
	ContextMenuActionGoForward ContextMenuAction = 11
	// ContextMenuActionStop: stop any ongoing loading operation.
	ContextMenuActionStop ContextMenuAction = 12
	// ContextMenuActionReload: reload the contents of current view.
	ContextMenuActionReload ContextMenuAction = 13
	// ContextMenuActionCopy: copy current selection the clipboard.
	ContextMenuActionCopy ContextMenuAction = 14
	// ContextMenuActionCut: cut current selection to the clipboard.
	ContextMenuActionCut ContextMenuAction = 15
	// ContextMenuActionPaste: paste clipboard contents.
	ContextMenuActionPaste ContextMenuAction = 16
	// ContextMenuActionDelete: delete current selection.
	ContextMenuActionDelete ContextMenuAction = 17
	// ContextMenuActionSelectAll: select all text.
	ContextMenuActionSelectAll ContextMenuAction = 18
	// ContextMenuActionInputMethods: input methods menu.
	ContextMenuActionInputMethods ContextMenuAction = 19
	// ContextMenuActionUnicode: unicode menu.
	ContextMenuActionUnicode ContextMenuAction = 20
	// ContextMenuActionSpellingGuess: proposed replacement for a misspelled
	// word.
	ContextMenuActionSpellingGuess ContextMenuAction = 21
	// ContextMenuActionNoGuessesFound: indicator that spellchecking found no
	// proposed replacements.
	ContextMenuActionNoGuessesFound ContextMenuAction = 22
	// ContextMenuActionIgnoreSpelling causes the spellchecker to ignore the
	// word for this session.
	ContextMenuActionIgnoreSpelling ContextMenuAction = 23
	// ContextMenuActionLearnSpelling causes the spellchecker to add the word to
	// the dictionary.
	ContextMenuActionLearnSpelling ContextMenuAction = 24
	// ContextMenuActionIgnoreGrammar: ignore grammar.
	ContextMenuActionIgnoreGrammar ContextMenuAction = 25
	// ContextMenuActionFontMenu: font options menu.
	ContextMenuActionFontMenu ContextMenuAction = 26
	// ContextMenuActionBold: bold.
	ContextMenuActionBold ContextMenuAction = 27
	// ContextMenuActionItalic: italic.
	ContextMenuActionItalic ContextMenuAction = 28
	// ContextMenuActionUnderline: underline.
	ContextMenuActionUnderline ContextMenuAction = 29
	// ContextMenuActionOutline: outline.
	ContextMenuActionOutline ContextMenuAction = 30
	// ContextMenuActionInspectElement: open current element in the inspector.
	ContextMenuActionInspectElement ContextMenuAction = 31
	// ContextMenuActionOpenVideoInNewWindow: open current video element in a
	// new window.
	ContextMenuActionOpenVideoInNewWindow ContextMenuAction = 32
	// ContextMenuActionOpenAudioInNewWindow: open current audio element in a
	// new window.
	ContextMenuActionOpenAudioInNewWindow ContextMenuAction = 33
	// ContextMenuActionCopyVideoLinkToClipboard: copy video link location in to
	// the clipboard.
	ContextMenuActionCopyVideoLinkToClipboard ContextMenuAction = 34
	// ContextMenuActionCopyAudioLinkToClipboard: copy audio link location in to
	// the clipboard.
	ContextMenuActionCopyAudioLinkToClipboard ContextMenuAction = 35
	// ContextMenuActionToggleMediaControls: enable or disable media controls.
	ContextMenuActionToggleMediaControls ContextMenuAction = 36
	// ContextMenuActionToggleMediaLoop: enable or disable media loop.
	ContextMenuActionToggleMediaLoop ContextMenuAction = 37
	// ContextMenuActionEnterVideoFullscreen: show current video element in
	// fullscreen mode.
	ContextMenuActionEnterVideoFullscreen ContextMenuAction = 38
	// ContextMenuActionMediaPlay: play current media element.
	ContextMenuActionMediaPlay ContextMenuAction = 39
	// ContextMenuActionMediaPause: pause current media element.
	ContextMenuActionMediaPause ContextMenuAction = 40
	// ContextMenuActionMediaMute: mute current media element.
	ContextMenuActionMediaMute ContextMenuAction = 41
	// ContextMenuActionDownloadVideoToDisk: download video to disk. Since 2.2.
	ContextMenuActionDownloadVideoToDisk ContextMenuAction = 42
	// ContextMenuActionDownloadAudioToDisk: download audio to disk. Since 2.2.
	ContextMenuActionDownloadAudioToDisk ContextMenuAction = 43
	// ContextMenuActionInsertEmoji: insert an emoji. Since 2.26.
	ContextMenuActionInsertEmoji ContextMenuAction = 44
	// ContextMenuActionPasteAsPlainText: paste clipboard contents as plain
	// text. Since 2.30.
	ContextMenuActionPasteAsPlainText ContextMenuAction = 45
	// ContextMenuActionCustom: custom action defined by applications.
	ContextMenuActionCustom ContextMenuAction = 10000
)

func marshalContextMenuAction(p uintptr) (interface{}, error) {
	return ContextMenuAction(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ContextMenuAction.
func (c ContextMenuAction) String() string {
	switch c {
	case ContextMenuActionNoAction:
		return "NoAction"
	case ContextMenuActionOpenLink:
		return "OpenLink"
	case ContextMenuActionOpenLinkInNewWindow:
		return "OpenLinkInNewWindow"
	case ContextMenuActionDownloadLinkToDisk:
		return "DownloadLinkToDisk"
	case ContextMenuActionCopyLinkToClipboard:
		return "CopyLinkToClipboard"
	case ContextMenuActionOpenImageInNewWindow:
		return "OpenImageInNewWindow"
	case ContextMenuActionDownloadImageToDisk:
		return "DownloadImageToDisk"
	case ContextMenuActionCopyImageToClipboard:
		return "CopyImageToClipboard"
	case ContextMenuActionCopyImageURLToClipboard:
		return "CopyImageURLToClipboard"
	case ContextMenuActionOpenFrameInNewWindow:
		return "OpenFrameInNewWindow"
	case ContextMenuActionGoBack:
		return "GoBack"
	case ContextMenuActionGoForward:
		return "GoForward"
	case ContextMenuActionStop:
		return "Stop"
	case ContextMenuActionReload:
		return "Reload"
	case ContextMenuActionCopy:
		return "Copy"
	case ContextMenuActionCut:
		return "Cut"
	case ContextMenuActionPaste:
		return "Paste"
	case ContextMenuActionDelete:
		return "Delete"
	case ContextMenuActionSelectAll:
		return "SelectAll"
	case ContextMenuActionInputMethods:
		return "InputMethods"
	case ContextMenuActionUnicode:
		return "Unicode"
	case ContextMenuActionSpellingGuess:
		return "SpellingGuess"
	case ContextMenuActionNoGuessesFound:
		return "NoGuessesFound"
	case ContextMenuActionIgnoreSpelling:
		return "IgnoreSpelling"
	case ContextMenuActionLearnSpelling:
		return "LearnSpelling"
	case ContextMenuActionIgnoreGrammar:
		return "IgnoreGrammar"
	case ContextMenuActionFontMenu:
		return "FontMenu"
	case ContextMenuActionBold:
		return "Bold"
	case ContextMenuActionItalic:
		return "Italic"
	case ContextMenuActionUnderline:
		return "Underline"
	case ContextMenuActionOutline:
		return "Outline"
	case ContextMenuActionInspectElement:
		return "InspectElement"
	case ContextMenuActionOpenVideoInNewWindow:
		return "OpenVideoInNewWindow"
	case ContextMenuActionOpenAudioInNewWindow:
		return "OpenAudioInNewWindow"
	case ContextMenuActionCopyVideoLinkToClipboard:
		return "CopyVideoLinkToClipboard"
	case ContextMenuActionCopyAudioLinkToClipboard:
		return "CopyAudioLinkToClipboard"
	case ContextMenuActionToggleMediaControls:
		return "ToggleMediaControls"
	case ContextMenuActionToggleMediaLoop:
		return "ToggleMediaLoop"
	case ContextMenuActionEnterVideoFullscreen:
		return "EnterVideoFullscreen"
	case ContextMenuActionMediaPlay:
		return "MediaPlay"
	case ContextMenuActionMediaPause:
		return "MediaPause"
	case ContextMenuActionMediaMute:
		return "MediaMute"
	case ContextMenuActionDownloadVideoToDisk:
		return "DownloadVideoToDisk"
	case ContextMenuActionDownloadAudioToDisk:
		return "DownloadAudioToDisk"
	case ContextMenuActionInsertEmoji:
		return "InsertEmoji"
	case ContextMenuActionPasteAsPlainText:
		return "PasteAsPlainText"
	case ContextMenuActionCustom:
		return "Custom"
	default:
		return fmt.Sprintf("ContextMenuAction(%d)", c)
	}
}
