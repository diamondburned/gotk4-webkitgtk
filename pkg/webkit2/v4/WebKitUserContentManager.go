// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// extern void _gotk4_webkit24_UserContentManager_ConnectScriptMessageReceived(gpointer, WebKitJavascriptResult*, guintptr);
import "C"

// glib.Type values for WebKitUserContentManager.go.
var GTypeUserContentManager = externglib.Type(C.webkit_user_content_manager_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeUserContentManager, F: marshalUserContentManager},
	})
}

// UserContentManagerOverrider contains methods that are overridable.
type UserContentManagerOverrider interface {
}

type UserContentManager struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*UserContentManager)(nil)
)

func classInitUserContentManagerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapUserContentManager(obj *externglib.Object) *UserContentManager {
	return &UserContentManager{
		Object: obj,
	}
}

func marshalUserContentManager(p uintptr) (interface{}, error) {
	return wrapUserContentManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_webkit24_UserContentManager_ConnectScriptMessageReceived
func _gotk4_webkit24_UserContentManager_ConnectScriptMessageReceived(arg0 C.gpointer, arg1 *C.WebKitJavascriptResult, arg2 C.guintptr) {
	var f func(jsResult *JavascriptResult)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(jsResult *JavascriptResult))
	}

	var _jsResult *JavascriptResult // out

	_jsResult = (*JavascriptResult)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.webkit_javascript_result_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_jsResult)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_javascript_result_unref((*C.WebKitJavascriptResult)(intern.C))
		},
	)

	f(_jsResult)
}

// ConnectScriptMessageReceived: this signal is emitted when JavaScript in a web
// view calls
// <code>window.webkit.messageHandlers.&lt;name&gt;.postMessage()</code>, after
// registering <code>&lt;name&gt;</code> using
// webkit_user_content_manager_register_script_message_handler().
func (manager *UserContentManager) ConnectScriptMessageReceived(f func(jsResult *JavascriptResult)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(manager, "script-message-received", false, unsafe.Pointer(C._gotk4_webkit24_UserContentManager_ConnectScriptMessageReceived), f)
}

// NewUserContentManager creates a new user content manager.
//
// The function returns the following values:
//
//    - userContentManager: KitUserContentManager.
//
func NewUserContentManager() *UserContentManager {
	var _cret *C.WebKitUserContentManager // in

	_cret = C.webkit_user_content_manager_new()

	var _userContentManager *UserContentManager // out

	_userContentManager = wrapUserContentManager(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _userContentManager
}

// AddFilter adds a KitUserContentFilter to the given KitUserContentManager. The
// same KitUserContentFilter can be reused with multiple KitUserContentManager
// instances.
//
// Filters need to be saved and loaded from KitUserContentFilterStore.
//
// The function takes the following parameters:
//
//    - filter: KitUserContentFilter.
//
func (manager *UserContentManager) AddFilter(filter *UserContentFilter) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.WebKitUserContentFilter  // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.WebKitUserContentFilter)(gextras.StructNative(unsafe.Pointer(filter)))

	C.webkit_user_content_manager_add_filter(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(filter)
}

// AddScript adds a KitUserScript to the given KitUserContentManager. The same
// KitUserScript can be reused with multiple KitUserContentManager instances.
//
// The function takes the following parameters:
//
//    - script: KitUserScript.
//
func (manager *UserContentManager) AddScript(script *UserScript) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.WebKitUserScript         // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.WebKitUserScript)(gextras.StructNative(unsafe.Pointer(script)))

	C.webkit_user_content_manager_add_script(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(script)
}

// AddStyleSheet adds a KitUserStyleSheet to the given KitUserContentManager.
// The same KitUserStyleSheet can be reused with multiple KitUserContentManager
// instances.
//
// The function takes the following parameters:
//
//    - stylesheet: KitUserStyleSheet.
//
func (manager *UserContentManager) AddStyleSheet(stylesheet *UserStyleSheet) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.WebKitUserStyleSheet     // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.WebKitUserStyleSheet)(gextras.StructNative(unsafe.Pointer(stylesheet)))

	C.webkit_user_content_manager_add_style_sheet(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(stylesheet)
}

// RegisterScriptMessageHandler registers a new user script message handler.
// After it is registered, scripts can use
// window.webkit.messageHandlers.&lt;name&gt;.postMessage(value) to send
// messages. Those messages are received by connecting handlers to the
// KitUserContentManager::script-message-received signal. The handler name is
// used as the detail of the signal. To avoid race conditions between
// registering the handler name, and starting to receive the signals, it is
// recommended to connect to the signal *before* registering the handler name:
//
// <informalexample><programlisting> WebKitWebView *view = webkit_web_view_new
// (); WebKitUserContentManager *manager =
// webkit_web_view_get_user_content_manager (); g_signal_connect (manager,
// "script-message-received::foobar", G_CALLBACK (handle_script_message), NULL);
// webkit_user_content_manager_register_script_message_handler (manager,
// "foobar"); </programlisting></informalexample>
//
// Registering a script message handler will fail if the requested name has been
// already registered before.
//
// The function takes the following parameters:
//
//    - name: name of the script message channel.
//
// The function returns the following values:
//
//    - ok: TRUE if message handler was registered successfully, or FALSE
//      otherwise.
//
func (manager *UserContentManager) RegisterScriptMessageHandler(name string) bool {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.gchar                    // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_user_content_manager_register_script_message_handler(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RegisterScriptMessageHandlerInWorld registers a new user script message
// handler in script world with name world_name. See
// webkit_user_content_manager_register_script_message_handler() for full
// description.
//
// Registering a script message handler will fail if the requested name has been
// already registered before.
//
// The function takes the following parameters:
//
//    - name: name of the script message channel.
//    - worldName: name of a KitScriptWorld.
//
// The function returns the following values:
//
//    - ok: TRUE if message handler was registered successfully, or FALSE
//      otherwise.
//
func (manager *UserContentManager) RegisterScriptMessageHandlerInWorld(name, worldName string) bool {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.gchar                    // out
	var _arg2 *C.gchar                    // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(worldName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.webkit_user_content_manager_register_script_message_handler_in_world(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(name)
	runtime.KeepAlive(worldName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveAllFilters removes all content filters from the given
// KitUserContentManager.
func (manager *UserContentManager) RemoveAllFilters() {
	var _arg0 *C.WebKitUserContentManager // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	C.webkit_user_content_manager_remove_all_filters(_arg0)
	runtime.KeepAlive(manager)
}

// RemoveAllScripts removes all user scripts from the given
// KitUserContentManager
//
// See also webkit_user_content_manager_remove_script().
func (manager *UserContentManager) RemoveAllScripts() {
	var _arg0 *C.WebKitUserContentManager // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	C.webkit_user_content_manager_remove_all_scripts(_arg0)
	runtime.KeepAlive(manager)
}

// RemoveAllStyleSheets removes all user style sheets from the given
// KitUserContentManager.
func (manager *UserContentManager) RemoveAllStyleSheets() {
	var _arg0 *C.WebKitUserContentManager // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	C.webkit_user_content_manager_remove_all_style_sheets(_arg0)
	runtime.KeepAlive(manager)
}

// RemoveFilter removes a filter from the given KitUserContentManager.
//
// Since 2.24.
//
// The function takes the following parameters:
//
//    - filter: KitUserContentFilter.
//
func (manager *UserContentManager) RemoveFilter(filter *UserContentFilter) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.WebKitUserContentFilter  // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.WebKitUserContentFilter)(gextras.StructNative(unsafe.Pointer(filter)))

	C.webkit_user_content_manager_remove_filter(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(filter)
}

// RemoveFilterByID removes a filter from the given KitUserContentManager given
// the identifier of a KitUserContentFilter as returned by
// webkit_user_content_filter_get_identifier().
//
// The function takes the following parameters:
//
//    - filterId: filter identifier.
//
func (manager *UserContentManager) RemoveFilterByID(filterId string) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.char                     // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filterId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_user_content_manager_remove_filter_by_id(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(filterId)
}

// RemoveScript removes a KitUserScript from the given KitUserContentManager.
//
// See also webkit_user_content_manager_remove_all_scripts().
//
// The function takes the following parameters:
//
//    - script: KitUserScript.
//
func (manager *UserContentManager) RemoveScript(script *UserScript) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.WebKitUserScript         // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.WebKitUserScript)(gextras.StructNative(unsafe.Pointer(script)))

	C.webkit_user_content_manager_remove_script(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(script)
}

// RemoveStyleSheet removes a KitUserStyleSheet from the given
// KitUserContentManager.
//
// See also webkit_user_content_manager_remove_all_style_sheets().
//
// The function takes the following parameters:
//
//    - stylesheet: KitUserStyleSheet.
//
func (manager *UserContentManager) RemoveStyleSheet(stylesheet *UserStyleSheet) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.WebKitUserStyleSheet     // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.WebKitUserStyleSheet)(gextras.StructNative(unsafe.Pointer(stylesheet)))

	C.webkit_user_content_manager_remove_style_sheet(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(stylesheet)
}

// UnregisterScriptMessageHandler unregisters a previously registered message
// handler.
//
// Note that this does *not* disconnect handlers for the
// KitUserContentManager::script-message-received signal; they will be kept
// connected, but the signal will not be emitted unless the handler name is
// registered again.
//
// See also webkit_user_content_manager_register_script_message_handler().
//
// The function takes the following parameters:
//
//    - name: name of the script message channel.
//
func (manager *UserContentManager) UnregisterScriptMessageHandler(name string) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_user_content_manager_unregister_script_message_handler(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(name)
}

// UnregisterScriptMessageHandlerInWorld unregisters a previously registered
// message handler in script world with name world_name.
//
// Note that this does *not* disconnect handlers for the
// KitUserContentManager::script-message-received signal; they will be kept
// connected, but the signal will not be emitted unless the handler name is
// registered again.
//
// See also
// webkit_user_content_manager_register_script_message_handler_in_world().
//
// The function takes the following parameters:
//
//    - name: name of the script message channel.
//    - worldName: name of a KitScriptWorld.
//
func (manager *UserContentManager) UnregisterScriptMessageHandlerInWorld(name, worldName string) {
	var _arg0 *C.WebKitUserContentManager // out
	var _arg1 *C.gchar                    // out
	var _arg2 *C.gchar                    // out

	_arg0 = (*C.WebKitUserContentManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(worldName)))
	defer C.free(unsafe.Pointer(_arg2))

	C.webkit_user_content_manager_unregister_script_message_handler_in_world(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(name)
	runtime.KeepAlive(worldName)
}
