// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

// GType values.
var (
	GTypeConsoleMessage = coreglib.Type(C.webkit_console_message_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeConsoleMessage, F: marshalConsoleMessage},
	})
}

// ConsoleMessage: instance of this type is always passed by reference.
type ConsoleMessage struct {
	*consoleMessage
}

// consoleMessage is the struct that's finalized.
type consoleMessage struct {
	native *C.WebKitConsoleMessage
}

func marshalConsoleMessage(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ConsoleMessage{&consoleMessage{(*C.WebKitConsoleMessage)(b)}}, nil
}

// Copy: make a copy of console_message.
//
// Deprecated: since version 2.40.
//
// The function returns the following values:
//
//   - consoleMessage: copy of passed in KitConsoleMessage.
//
func (consoleMessage *ConsoleMessage) Copy() *ConsoleMessage {
	var _arg0 *C.WebKitConsoleMessage // out
	var _cret *C.WebKitConsoleMessage // in

	_arg0 = (*C.WebKitConsoleMessage)(gextras.StructNative(unsafe.Pointer(consoleMessage)))

	_cret = C.webkit_console_message_copy(_arg0)
	runtime.KeepAlive(consoleMessage)

	var _consoleMessage *ConsoleMessage // out

	_consoleMessage = (*ConsoleMessage)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_consoleMessage)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_console_message_free((*C.WebKitConsoleMessage)(intern.C))
		},
	)

	return _consoleMessage
}

// Level gets the log level of a KitConsoleMessage
//
// Deprecated: since version 2.40.
//
// The function returns the following values:
//
//   - consoleMessageLevel indicating the log level of console_message.
//
func (consoleMessage *ConsoleMessage) Level() ConsoleMessageLevel {
	var _arg0 *C.WebKitConsoleMessage     // out
	var _cret C.WebKitConsoleMessageLevel // in

	_arg0 = (*C.WebKitConsoleMessage)(gextras.StructNative(unsafe.Pointer(consoleMessage)))

	_cret = C.webkit_console_message_get_level(_arg0)
	runtime.KeepAlive(consoleMessage)

	var _consoleMessageLevel ConsoleMessageLevel // out

	_consoleMessageLevel = ConsoleMessageLevel(_cret)

	return _consoleMessageLevel
}

// Line gets the line number of a KitConsoleMessage
//
// Deprecated: since version 2.40.
//
// The function returns the following values:
//
//   - guint: line number of console_message.
//
func (consoleMessage *ConsoleMessage) Line() uint {
	var _arg0 *C.WebKitConsoleMessage // out
	var _cret C.guint                 // in

	_arg0 = (*C.WebKitConsoleMessage)(gextras.StructNative(unsafe.Pointer(consoleMessage)))

	_cret = C.webkit_console_message_get_line(_arg0)
	runtime.KeepAlive(consoleMessage)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Source gets the source of a KitConsoleMessage
//
// Deprecated: since version 2.40.
//
// The function returns the following values:
//
//   - consoleMessageSource indicating the source of console_message.
//
func (consoleMessage *ConsoleMessage) Source() ConsoleMessageSource {
	var _arg0 *C.WebKitConsoleMessage      // out
	var _cret C.WebKitConsoleMessageSource // in

	_arg0 = (*C.WebKitConsoleMessage)(gextras.StructNative(unsafe.Pointer(consoleMessage)))

	_cret = C.webkit_console_message_get_source(_arg0)
	runtime.KeepAlive(consoleMessage)

	var _consoleMessageSource ConsoleMessageSource // out

	_consoleMessageSource = ConsoleMessageSource(_cret)

	return _consoleMessageSource
}

// SourceID gets the source identifier of a KitConsoleMessage
//
// Deprecated: since version 2.40.
//
// The function returns the following values:
//
//   - utf8: source identifier of console_message.
//
func (consoleMessage *ConsoleMessage) SourceID() string {
	var _arg0 *C.WebKitConsoleMessage // out
	var _cret *C.gchar                // in

	_arg0 = (*C.WebKitConsoleMessage)(gextras.StructNative(unsafe.Pointer(consoleMessage)))

	_cret = C.webkit_console_message_get_source_id(_arg0)
	runtime.KeepAlive(consoleMessage)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Text gets the text message of a KitConsoleMessage
//
// Deprecated: since version 2.40.
//
// The function returns the following values:
//
//   - utf8: text message of console_message.
//
func (consoleMessage *ConsoleMessage) Text() string {
	var _arg0 *C.WebKitConsoleMessage // out
	var _cret *C.gchar                // in

	_arg0 = (*C.WebKitConsoleMessage)(gextras.StructNative(unsafe.Pointer(consoleMessage)))

	_cret = C.webkit_console_message_get_text(_arg0)
	runtime.KeepAlive(consoleMessage)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}