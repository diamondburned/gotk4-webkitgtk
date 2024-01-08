// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit-web-extension.h>
// extern void _gotk4_webkit2webextension4_WebPage_ConnectDocumentLoaded(gpointer, guintptr);
// extern void _gotk4_webkit2webextension4_WebPage_ConnectConsoleMessageSent(gpointer, WebKitConsoleMessage*, guintptr);
// extern gboolean _gotk4_webkit2webextension4_WebPage_ConnectUserMessageReceived(gpointer, WebKitUserMessage*, guintptr);
// extern gboolean _gotk4_webkit2webextension4_WebPage_ConnectSendRequest(gpointer, WebKitURIRequest*, WebKitURIResponse*, guintptr);
// extern gboolean _gotk4_webkit2webextension4_WebPage_ConnectContextMenu(gpointer, WebKitContextMenu*, WebKitWebHitTestResult*, guintptr);
import "C"

// GType values.
var (
	GTypeWebPage = coreglib.Type(C.webkit_web_page_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWebPage, F: marshalWebPage},
	})
}

// WebPageOverrides contains methods that are overridable.
type WebPageOverrides struct {
}

func defaultWebPageOverrides(v *WebPage) WebPageOverrides {
	return WebPageOverrides{}
}

// WebPage: loaded web page.
type WebPage struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*WebPage)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*WebPage, *WebPageClass, WebPageOverrides](
		GTypeWebPage,
		initWebPageClass,
		wrapWebPage,
		defaultWebPageOverrides,
	)
}

func initWebPageClass(gclass unsafe.Pointer, overrides WebPageOverrides, classInitFunc func(*WebPageClass)) {
	if classInitFunc != nil {
		class := (*WebPageClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapWebPage(obj *coreglib.Object) *WebPage {
	return &WebPage{
		Object: obj,
	}
}

func marshalWebPage(p uintptr) (interface{}, error) {
	return wrapWebPage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectConsoleMessageSent is emitted when a message is sent to the console.
// This can be a message produced by the use of JavaScript console API,
// a JavaScript exception, a security error or other errors, warnings, debug or
// log messages. The console_message contains information of the message.
func (webPage *WebPage) ConnectConsoleMessageSent(f func(consoleMessage *ConsoleMessage)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(webPage, "console-message-sent", false, unsafe.Pointer(C._gotk4_webkit2webextension4_WebPage_ConnectConsoleMessageSent), f)
}

// ConnectContextMenu is emitted before a context menu is displayed in the UI
// Process to give the application a chance to customize the proposed menu,
// build its own context menu or pass user data to the UI Process. This signal
// is useful when the information available in the UI Process is not enough
// to build or customize the context menu, for example, to add menu entries
// depending on the node at the coordinates of the hit_test_result. Otherwise,
// it's recommended to use KitWebView::context-menu signal instead.
func (webPage *WebPage) ConnectContextMenu(f func(contextMenu *ContextMenu, hitTestResult *WebHitTestResult) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(webPage, "context-menu", false, unsafe.Pointer(C._gotk4_webkit2webextension4_WebPage_ConnectContextMenu), f)
}

// ConnectDocumentLoaded: this signal is emitted when the DOM document of a
// KitWebPage has been loaded.
//
// You can wait for this signal to get the DOM document.
func (webPage *WebPage) ConnectDocumentLoaded(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(webPage, "document-loaded", false, unsafe.Pointer(C._gotk4_webkit2webextension4_WebPage_ConnectDocumentLoaded), f)
}

// ConnectSendRequest: this signal is emitted when request is about to be sent
// to the server. This signal can be used to modify the KitURIRequest that will
// be sent to the server. You can also cancel the resource load operation by
// connecting to this signal and returning TRUE.
//
// In case of a server redirection this signal is emitted again with the request
// argument containing the new request to be sent to the server due to the
// redirection and the redirected_response parameter containing the response
// received by the server for the initial request.
//
// Modifications to the KitURIRequest and its associated MessageHeaders will be
// taken into account when the request is sent over the network.
func (webPage *WebPage) ConnectSendRequest(f func(request *URIRequest, redirectedResponse *URIResponse) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(webPage, "send-request", false, unsafe.Pointer(C._gotk4_webkit2webextension4_WebPage_ConnectSendRequest), f)
}

// ConnectUserMessageReceived: this signal is emitted when a KitUserMessage is
// received from the KitWebView corresponding to web_page. You can reply to the
// message using webkit_user_message_send_reply().
//
// You can handle the user message asynchronously by calling g_object_ref() on
// message and returning TRUE. If the last reference of message is removed and
// the message has been replied, the operation in the KitWebView will finish
// with error WEBKIT_USER_MESSAGE_UNHANDLED_MESSAGE.
func (webPage *WebPage) ConnectUserMessageReceived(f func(message *UserMessage) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(webPage, "user-message-received", false, unsafe.Pointer(C._gotk4_webkit2webextension4_WebPage_ConnectUserMessageReceived), f)
}

// DomDocument: get the KitDOMDocument currently loaded in web_page
//
// Deprecated: Use JavaScriptCore API instead.
//
// The function returns the following values:
//
//   - domDocument currently loaded, or NULL if no document is currently loaded.
//
func (webPage *WebPage) DomDocument() *DOMDocument {
	var _arg0 *C.WebKitWebPage     // out
	var _cret *C.WebKitDOMDocument // in

	_arg0 = (*C.WebKitWebPage)(unsafe.Pointer(coreglib.InternObject(webPage).Native()))

	_cret = C.webkit_web_page_get_dom_document(_arg0)
	runtime.KeepAlive(webPage)

	var _domDocument *DOMDocument // out

	_domDocument = wrapDOMDocument(coreglib.Take(unsafe.Pointer(_cret)))

	return _domDocument
}

// Editor gets the KitWebEditor of a KitWebPage.
//
// The function returns the following values:
//
//   - webEditor: KitWebEditor.
//
func (webPage *WebPage) Editor() *WebEditor {
	var _arg0 *C.WebKitWebPage   // out
	var _cret *C.WebKitWebEditor // in

	_arg0 = (*C.WebKitWebPage)(unsafe.Pointer(coreglib.InternObject(webPage).Native()))

	_cret = C.webkit_web_page_get_editor(_arg0)
	runtime.KeepAlive(webPage)

	var _webEditor *WebEditor // out

	_webEditor = wrapWebEditor(coreglib.Take(unsafe.Pointer(_cret)))

	return _webEditor
}

// FormManager: get the KitWebFormManager of web_page in world.
//
// The function takes the following parameters:
//
//   - world (optional): KitScriptWorld.
//
// The function returns the following values:
//
//   - webFormManager: KitWebFormManager.
//
func (webPage *WebPage) FormManager(world *ScriptWorld) *WebFormManager {
	var _arg0 *C.WebKitWebPage        // out
	var _arg1 *C.WebKitScriptWorld    // out
	var _cret *C.WebKitWebFormManager // in

	_arg0 = (*C.WebKitWebPage)(unsafe.Pointer(coreglib.InternObject(webPage).Native()))
	if world != nil {
		_arg1 = (*C.WebKitScriptWorld)(unsafe.Pointer(coreglib.InternObject(world).Native()))
	}

	_cret = C.webkit_web_page_get_form_manager(_arg0, _arg1)
	runtime.KeepAlive(webPage)
	runtime.KeepAlive(world)

	var _webFormManager *WebFormManager // out

	_webFormManager = wrapWebFormManager(coreglib.Take(unsafe.Pointer(_cret)))

	return _webFormManager
}

// ID: get the identifier of the KitWebPage.
//
// The function returns the following values:
//
//   - guint64: identifier of web_page.
//
func (webPage *WebPage) ID() uint64 {
	var _arg0 *C.WebKitWebPage // out
	var _cret C.guint64        // in

	_arg0 = (*C.WebKitWebPage)(unsafe.Pointer(coreglib.InternObject(webPage).Native()))

	_cret = C.webkit_web_page_get_id(_arg0)
	runtime.KeepAlive(webPage)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// MainFrame returns the main frame of a KitWebPage.
//
// The function returns the following values:
//
//   - frame that is the main frame of web_page.
//
func (webPage *WebPage) MainFrame() *Frame {
	var _arg0 *C.WebKitWebPage // out
	var _cret *C.WebKitFrame   // in

	_arg0 = (*C.WebKitWebPage)(unsafe.Pointer(coreglib.InternObject(webPage).Native()))

	_cret = C.webkit_web_page_get_main_frame(_arg0)
	runtime.KeepAlive(webPage)

	var _frame *Frame // out

	_frame = wrapFrame(coreglib.Take(unsafe.Pointer(_cret)))

	return _frame
}

// URI returns the current active URI of web_page.
//
// You can monitor the active URI by connecting to the notify::uri signal of
// web_page.
//
// The function returns the following values:
//
//   - utf8: current active URI of web_view or NULL if nothing has been loaded
//     yet.
//
func (webPage *WebPage) URI() string {
	var _arg0 *C.WebKitWebPage // out
	var _cret *C.gchar         // in

	_arg0 = (*C.WebKitWebPage)(unsafe.Pointer(coreglib.InternObject(webPage).Native()))

	_cret = C.webkit_web_page_get_uri(_arg0)
	runtime.KeepAlive(webPage)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SendMessageToViewFinish: finish an asynchronous operation started with
// webkit_web_page_send_message_to_view().
//
// The function takes the following parameters:
//
//   - result: Result.
//
// The function returns the following values:
//
//   - userMessage with the reply or NULL in case of error.
//
func (webPage *WebPage) SendMessageToViewFinish(result gio.AsyncResulter) (*UserMessage, error) {
	var _arg0 *C.WebKitWebPage     // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.WebKitUserMessage // in
	var _cerr *C.GError            // in

	_arg0 = (*C.WebKitWebPage)(unsafe.Pointer(coreglib.InternObject(webPage).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.webkit_web_page_send_message_to_view_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(webPage)
	runtime.KeepAlive(result)

	var _userMessage *UserMessage // out
	var _goerr error              // out

	_userMessage = wrapUserMessage(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _userMessage, _goerr
}

// WebPageClass: instance of this type is always passed by reference.
type WebPageClass struct {
	*webPageClass
}

// webPageClass is the struct that's finalized.
type webPageClass struct {
	native *C.WebKitWebPageClass
}
