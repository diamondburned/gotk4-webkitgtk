// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

//export _gotk4_webkit2webextension4_WebPage_ConnectConsoleMessageSent
func _gotk4_webkit2webextension4_WebPage_ConnectConsoleMessageSent(arg0 C.gpointer, arg1 *C.WebKitConsoleMessage, arg2 C.guintptr) {
	var f func(consoleMessage *ConsoleMessage)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(consoleMessage *ConsoleMessage))
	}

	var _consoleMessage *ConsoleMessage // out

	_consoleMessage = (*ConsoleMessage)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_consoleMessage)
}

//export _gotk4_webkit2webextension4_WebPage_ConnectContextMenu
func _gotk4_webkit2webextension4_WebPage_ConnectContextMenu(arg0 C.gpointer, arg1 *C.WebKitContextMenu, arg2 *C.WebKitWebHitTestResult, arg3 C.guintptr) (cret C.gboolean) {
	var f func(contextMenu *ContextMenu, hitTestResult *WebHitTestResult) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(contextMenu *ContextMenu, hitTestResult *WebHitTestResult) (ok bool))
	}

	var _contextMenu *ContextMenu        // out
	var _hitTestResult *WebHitTestResult // out

	_contextMenu = wrapContextMenu(coreglib.Take(unsafe.Pointer(arg1)))
	_hitTestResult = wrapWebHitTestResult(coreglib.Take(unsafe.Pointer(arg2)))

	ok := f(_contextMenu, _hitTestResult)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_webkit2webextension4_WebPage_ConnectDocumentLoaded
func _gotk4_webkit2webextension4_WebPage_ConnectDocumentLoaded(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_webkit2webextension4_WebPage_ConnectSendRequest
func _gotk4_webkit2webextension4_WebPage_ConnectSendRequest(arg0 C.gpointer, arg1 *C.WebKitURIRequest, arg2 *C.WebKitURIResponse, arg3 C.guintptr) (cret C.gboolean) {
	var f func(request *URIRequest, redirectedResponse *URIResponse) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(request *URIRequest, redirectedResponse *URIResponse) (ok bool))
	}

	var _request *URIRequest             // out
	var _redirectedResponse *URIResponse // out

	_request = wrapURIRequest(coreglib.Take(unsafe.Pointer(arg1)))
	_redirectedResponse = wrapURIResponse(coreglib.Take(unsafe.Pointer(arg2)))

	ok := f(_request, _redirectedResponse)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_webkit2webextension4_WebPage_ConnectUserMessageReceived
func _gotk4_webkit2webextension4_WebPage_ConnectUserMessageReceived(arg0 C.gpointer, arg1 *C.WebKitUserMessage, arg2 C.guintptr) (cret C.gboolean) {
	var f func(message *UserMessage) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(message *UserMessage) (ok bool))
	}

	var _message *UserMessage // out

	_message = wrapUserMessage(coreglib.Take(unsafe.Pointer(arg1)))

	ok := f(_message)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
