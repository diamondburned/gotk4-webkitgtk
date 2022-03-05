// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// glib.Type values for WebKitWebsiteDataManager.go.
var (
	GTypeTLSErrorsPolicy    = externglib.Type(C.webkit_tls_errors_policy_get_type())
	GTypeWebsiteDataManager = externglib.Type(C.webkit_website_data_manager_get_type())
	GTypeITPFirstParty      = externglib.Type(C.webkit_itp_first_party_get_type())
	GTypeITPThirdParty      = externglib.Type(C.webkit_itp_third_party_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTLSErrorsPolicy, F: marshalTLSErrorsPolicy},
		{T: GTypeWebsiteDataManager, F: marshalWebsiteDataManager},
		{T: GTypeITPFirstParty, F: marshalITPFirstParty},
		{T: GTypeITPThirdParty, F: marshalITPThirdParty},
	})
}

// TLSErrorsPolicy: enum values used to denote the TLS errors policy.
type TLSErrorsPolicy C.gint

const (
	// TLSErrorsPolicyIgnore: ignore TLS errors.
	TLSErrorsPolicyIgnore TLSErrorsPolicy = iota
	// TLSErrorsPolicyFail: TLS errors will emit
	// KitWebView::load-failed-with-tls-errors and, if the signal is handled,
	// finish the load. In case the signal is not handled,
	// KitWebView::load-failed is emitted before the load finishes.
	TLSErrorsPolicyFail
)

func marshalTLSErrorsPolicy(p uintptr) (interface{}, error) {
	return TLSErrorsPolicy(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for TLSErrorsPolicy.
func (t TLSErrorsPolicy) String() string {
	switch t {
	case TLSErrorsPolicyIgnore:
		return "Ignore"
	case TLSErrorsPolicyFail:
		return "Fail"
	default:
		return fmt.Sprintf("TLSErrorsPolicy(%d)", t)
	}
}

// WebsiteDataManagerOverrider contains methods that are overridable.
type WebsiteDataManagerOverrider interface {
}

type WebsiteDataManager struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*WebsiteDataManager)(nil)
)

func classInitWebsiteDataManagerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWebsiteDataManager(obj *externglib.Object) *WebsiteDataManager {
	return &WebsiteDataManager{
		Object: obj,
	}
}

func marshalWebsiteDataManager(p uintptr) (interface{}, error) {
	return wrapWebsiteDataManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWebsiteDataManagerEphemeral creates an ephemeral KitWebsiteDataManager.
// See KitWebsiteDataManager:is-ephemeral for more details.
//
// The function returns the following values:
//
//    - websiteDataManager: new ephemeral KitWebsiteDataManager.
//
func NewWebsiteDataManagerEphemeral() *WebsiteDataManager {
	var _cret *C.WebKitWebsiteDataManager // in

	_cret = C.webkit_website_data_manager_new_ephemeral()

	var _websiteDataManager *WebsiteDataManager // out

	_websiteDataManager = wrapWebsiteDataManager(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _websiteDataManager
}

// Clear: asynchronously clear the website data of the given types modified in
// the past timespan. If timespan is 0, all website data will be removed.
//
// When the operation is finished, callback will be called. You can then call
// webkit_website_data_manager_clear_finish() to get the result of the
// operation.
//
// Due to implementation limitations, this function does not currently delete
// any stored cookies if timespan is nonzero. This behavior may change in the
// future.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - types: KitWebsiteDataTypes.
//    - timespan: Span.
//    - callback (optional) to call when the request is satisfied.
//
func (manager *WebsiteDataManager) Clear(ctx context.Context, types WebsiteDataTypes, timespan glib.TimeSpan, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg3 *C.GCancellable             // out
	var _arg1 C.WebKitWebsiteDataTypes    // out
	var _arg2 C.GTimeSpan                 // out
	var _arg4 C.GAsyncReadyCallback       // out
	var _arg5 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.WebKitWebsiteDataTypes(types)
	_arg2 = C.gint64(timespan)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_website_data_manager_clear(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(types)
	runtime.KeepAlive(timespan)
	runtime.KeepAlive(callback)
}

// ClearFinish: finish an asynchronous operation started with
// webkit_website_data_manager_clear().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (manager *WebsiteDataManager) ClearFinish(result gio.AsyncResulter) error {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GAsyncResult             // out
	var _cerr *C.GError                   // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	C.webkit_website_data_manager_clear_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Fetch: asynchronously get the list of KitWebsiteData for the given types.
//
// When the operation is finished, callback will be called. You can then call
// webkit_website_data_manager_fetch_finish() to get the result of the
// operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - types: KitWebsiteDataTypes.
//    - callback (optional) to call when the request is satisfied.
//
func (manager *WebsiteDataManager) Fetch(ctx context.Context, types WebsiteDataTypes, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg2 *C.GCancellable             // out
	var _arg1 C.WebKitWebsiteDataTypes    // out
	var _arg3 C.GAsyncReadyCallback       // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.WebKitWebsiteDataTypes(types)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_website_data_manager_fetch(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(types)
	runtime.KeepAlive(callback)
}

// FetchFinish: finish an asynchronous operation started with
// webkit_website_data_manager_fetch().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - list of KitWebsiteData. You must free the #GList with g_list_free() and
//      unref the KitWebsiteData<!-- -->s with webkit_website_data_unref() when
//      you're done with them.
//
func (manager *WebsiteDataManager) FetchFinish(result gio.AsyncResulter) ([]*WebsiteData, error) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GList                    // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.webkit_website_data_manager_fetch_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(result)

	var _list []*WebsiteData // out
	var _goerr error         // out

	_list = make([]*WebsiteData, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.WebKitWebsiteData)(v)
		var dst *WebsiteData // out
		dst = (*WebsiteData)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.webkit_website_data_unref((*C.WebKitWebsiteData)(intern.C))
			},
		)
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// BaseCacheDirectory: get the KitWebsiteDataManager:base-cache-directory
// property.
//
// The function returns the following values:
//
//    - utf8 (optional): base directory for Website cache, or NULL if
//      KitWebsiteDataManager:base-cache-directory was not provided or manager is
//      ephemeral.
//
func (manager *WebsiteDataManager) BaseCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_base_cache_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// BaseDataDirectory: get the KitWebsiteDataManager:base-data-directory
// property.
//
// The function returns the following values:
//
//    - utf8 (optional): base directory for Website data, or NULL if
//      KitWebsiteDataManager:base-data-directory was not provided or manager is
//      ephemeral.
//
func (manager *WebsiteDataManager) BaseDataDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_base_data_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CookieManager: get the KitCookieManager of manager.
//
// The function returns the following values:
//
//    - cookieManager: KitCookieManager.
//
func (manager *WebsiteDataManager) CookieManager() *CookieManager {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.WebKitCookieManager      // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_cookie_manager(_arg0)
	runtime.KeepAlive(manager)

	var _cookieManager *CookieManager // out

	_cookieManager = wrapCookieManager(externglib.Take(unsafe.Pointer(_cret)))

	return _cookieManager
}

// DiskCacheDirectory: get the KitWebsiteDataManager:disk-cache-directory
// property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where HTTP disk cache is stored or NULL if
//      manager is ephemeral.
//
func (manager *WebsiteDataManager) DiskCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_disk_cache_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// DomCacheDirectory: get the KitWebsiteDataManager:dom-cache-directory
// property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where DOM cache is stored or NULL if manager
//      is ephemeral.
//
func (manager *WebsiteDataManager) DomCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_dom_cache_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// HstsCacheDirectory: get the KitWebsiteDataManager:hsts-cache-directory
// property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where the HSTS cache is stored or NULL if
//      manager is ephemeral.
//
func (manager *WebsiteDataManager) HstsCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_hsts_cache_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IndexeddbDirectory: get the KitWebsiteDataManager:indexeddb-directory
// property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where IndexedDB databases are stored or NULL
//      if manager is ephemeral.
//
func (manager *WebsiteDataManager) IndexeddbDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_indexeddb_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ITPDirectory: get the KitWebsiteDataManager:itp-directory property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where Intelligent Tracking Prevention data is
//      stored or NULL if manager is ephemeral.
//
func (manager *WebsiteDataManager) ITPDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_itp_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ITPEnabled: get whether Intelligent Tracking Prevention (ITP) is enabled or
// not.
//
// The function returns the following values:
//
//    - ok: TRUE if ITP is enabled, or FALSE otherwise.
//
func (manager *WebsiteDataManager) ITPEnabled() bool {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_itp_enabled(_arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ITPSummary: asynchronously get the list of KitITPThirdParty seen for manager.
// Every KitITPThirdParty contains the list of KitITPFirstParty under which it
// has been seen.
//
// When the operation is finished, callback will be called. You can then call
// webkit_website_data_manager_get_itp_summary_finish() to get the result of the
// operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - callback (optional) to call when the request is satisfied.
//
func (manager *WebsiteDataManager) ITPSummary(ctx context.Context, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GCancellable             // out
	var _arg2 C.GAsyncReadyCallback       // out
	var _arg3 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_website_data_manager_get_itp_summary(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// ITPSummaryFinish: finish an asynchronous operation started with
// webkit_website_data_manager_get_itp_summary().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - list of KitITPThirdParty. You must free the #GList with g_list_free() and
//      unref the KitITPThirdParty<!-- -->s with webkit_itp_third_party_unref()
//      when you're done with them.
//
func (manager *WebsiteDataManager) ITPSummaryFinish(result gio.AsyncResulter) ([]*ITPThirdParty, error) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GList                    // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.webkit_website_data_manager_get_itp_summary_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(result)

	var _list []*ITPThirdParty // out
	var _goerr error           // out

	_list = make([]*ITPThirdParty, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.WebKitITPThirdParty)(v)
		var dst *ITPThirdParty // out
		dst = (*ITPThirdParty)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.webkit_itp_third_party_unref((*C.WebKitITPThirdParty)(intern.C))
			},
		)
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LocalStorageDirectory: get the KitWebsiteDataManager:local-storage-directory
// property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where local storage data is stored or NULL if
//      manager is ephemeral.
//
func (manager *WebsiteDataManager) LocalStorageDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_local_storage_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// OfflineApplicationCacheDirectory: get the
// KitWebsiteDataManager:offline-application-cache-directory property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where offline web application cache is stored
//      or NULL if manager is ephemeral.
//
func (manager *WebsiteDataManager) OfflineApplicationCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_offline_application_cache_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// PersistentCredentialStorageEnabled: get whether persistent credential storage
// is enabled or not. See also
// webkit_website_data_manager_set_persistent_credential_storage_enabled().
//
// The function returns the following values:
//
//    - ok: TRUE if persistent credential storage is enabled, or FALSE otherwise.
//
func (manager *WebsiteDataManager) PersistentCredentialStorageEnabled() bool {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_persistent_credential_storage_enabled(_arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ServiceWorkerRegistrationsDirectory: get the
// KitWebsiteDataManager:service-worker-registrations-directory property.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where service worker registrations are stored
//      or NULL if manager is ephemeral.
//
func (manager *WebsiteDataManager) ServiceWorkerRegistrationsDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_service_worker_registrations_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TLSErrorsPolicy: get the TLS errors policy of manager.
//
// The function returns the following values:
//
//    - tlsErrorsPolicy: KitTLSErrorsPolicy.
//
func (manager *WebsiteDataManager) TLSErrorsPolicy() TLSErrorsPolicy {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.WebKitTLSErrorsPolicy     // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_tls_errors_policy(_arg0)
	runtime.KeepAlive(manager)

	var _tlsErrorsPolicy TLSErrorsPolicy // out

	_tlsErrorsPolicy = TLSErrorsPolicy(_cret)

	return _tlsErrorsPolicy
}

// WebsqlDirectory: get the KitWebsiteDataManager:websql-directory property.
//
// Deprecated: WebSQL is no longer supported. Use IndexedDB instead.
//
// The function returns the following values:
//
//    - utf8 (optional): directory where WebSQL databases are stored or NULL if
//      manager is ephemeral.
//
func (manager *WebsiteDataManager) WebsqlDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_get_websql_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IsEphemeral: get whether a KitWebsiteDataManager is ephemeral. See
// KitWebsiteDataManager:is-ephemeral for more details.
//
// The function returns the following values:
//
//    - ok: TRUE if manager is ephemeral or FALSE otherwise.
//
func (manager *WebsiteDataManager) IsEphemeral() bool {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))

	_cret = C.webkit_website_data_manager_is_ephemeral(_arg0)
	runtime.KeepAlive(manager)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remove: asynchronously removes the website data of the for the given types
// for websites in the given website_data list. Use
// webkit_website_data_manager_clear() if you want to remove the website data
// for all sites.
//
// When the operation is finished, callback will be called. You can then call
// webkit_website_data_manager_remove_finish() to get the result of the
// operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - types: KitWebsiteDataTypes.
//    - websiteData of KitWebsiteData.
//    - callback (optional) to call when the request is satisfied.
//
func (manager *WebsiteDataManager) Remove(ctx context.Context, types WebsiteDataTypes, websiteData []*WebsiteData, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg3 *C.GCancellable             // out
	var _arg1 C.WebKitWebsiteDataTypes    // out
	var _arg2 *C.GList                    // out
	var _arg4 C.GAsyncReadyCallback       // out
	var _arg5 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.WebKitWebsiteDataTypes(types)
	for i := len(websiteData) - 1; i >= 0; i-- {
		src := websiteData[i]
		var dst *C.WebKitWebsiteData // out
		dst = (*C.WebKitWebsiteData)(gextras.StructNative(unsafe.Pointer(src)))
		_arg2 = C.g_list_prepend(_arg2, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_list_free(_arg2)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_website_data_manager_remove(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(types)
	runtime.KeepAlive(websiteData)
	runtime.KeepAlive(callback)
}

// RemoveFinish: finish an asynchronous operation started with
// webkit_website_data_manager_remove().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (manager *WebsiteDataManager) RemoveFinish(result gio.AsyncResulter) error {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GAsyncResult             // out
	var _cerr *C.GError                   // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	C.webkit_website_data_manager_remove_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetITPEnabled: enable or disable Intelligent Tracking Prevention (ITP). When
// ITP is enabled resource load statistics are collected and used to decide
// whether to allow or block third-party cookies and prevent user tracking. Note
// that while ITP is enabled the accept policy
// WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY is ignored and
// WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS is used instead. See also
// webkit_cookie_manager_set_accept_policy().
//
// The function takes the following parameters:
//
//    - enabled: value to set.
//
func (manager *WebsiteDataManager) SetITPEnabled(enabled bool) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 C.gboolean                  // out

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.webkit_website_data_manager_set_itp_enabled(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(enabled)
}

// SetNetworkProxySettings: set the network proxy settings to be used by
// connections started in manager session. By default
// WEBKIT_NETWORK_PROXY_MODE_DEFAULT is used, which means that the system
// settings will be used (g_proxy_resolver_get_default()). If you want to
// override the system default settings, you can either use
// WEBKIT_NETWORK_PROXY_MODE_NO_PROXY to make sure no proxies are used at all,
// or WEBKIT_NETWORK_PROXY_MODE_CUSTOM to provide your own proxy settings. When
// proxy_mode is WEBKIT_NETWORK_PROXY_MODE_CUSTOM proxy_settings must be a valid
// KitNetworkProxySettings; otherwise, proxy_settings must be NULL.
//
// The function takes the following parameters:
//
//    - proxyMode: KitNetworkProxyMode.
//    - proxySettings (optional) or NULL.
//
func (manager *WebsiteDataManager) SetNetworkProxySettings(proxyMode NetworkProxyMode, proxySettings *NetworkProxySettings) {
	var _arg0 *C.WebKitWebsiteDataManager   // out
	var _arg1 C.WebKitNetworkProxyMode      // out
	var _arg2 *C.WebKitNetworkProxySettings // out

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = C.WebKitNetworkProxyMode(proxyMode)
	if proxySettings != nil {
		_arg2 = (*C.WebKitNetworkProxySettings)(gextras.StructNative(unsafe.Pointer(proxySettings)))
	}

	C.webkit_website_data_manager_set_network_proxy_settings(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(proxyMode)
	runtime.KeepAlive(proxySettings)
}

// SetPersistentCredentialStorageEnabled: enable or disable persistent
// credential storage. When enabled, which is the default for non-ephemeral
// sessions, the network process will try to read and write HTTP authentiacation
// credentials from persistent storage.
//
// The function takes the following parameters:
//
//    - enabled: value to set.
//
func (manager *WebsiteDataManager) SetPersistentCredentialStorageEnabled(enabled bool) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 C.gboolean                  // out

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.webkit_website_data_manager_set_persistent_credential_storage_enabled(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(enabled)
}

// SetTLSErrorsPolicy: set the TLS errors policy of manager as policy.
//
// The function takes the following parameters:
//
//    - policy: KitTLSErrorsPolicy.
//
func (manager *WebsiteDataManager) SetTLSErrorsPolicy(policy TLSErrorsPolicy) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 C.WebKitTLSErrorsPolicy     // out

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(externglib.InternObject(manager).Native()))
	_arg1 = C.WebKitTLSErrorsPolicy(policy)

	C.webkit_website_data_manager_set_tls_errors_policy(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(policy)
}

// ITPFirstParty: instance of this type is always passed by reference.
type ITPFirstParty struct {
	*itpFirstParty
}

// itpFirstParty is the struct that's finalized.
type itpFirstParty struct {
	native *C.WebKitITPFirstParty
}

func marshalITPFirstParty(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ITPFirstParty{&itpFirstParty{(*C.WebKitITPFirstParty)(b)}}, nil
}

// Domain: get the domain name of itp_first_party.
//
// The function returns the following values:
//
//    - utf8: domain name.
//
func (itpFirstParty *ITPFirstParty) Domain() string {
	var _arg0 *C.WebKitITPFirstParty // out
	var _cret *C.char                // in

	_arg0 = (*C.WebKitITPFirstParty)(gextras.StructNative(unsafe.Pointer(itpFirstParty)))

	_cret = C.webkit_itp_first_party_get_domain(_arg0)
	runtime.KeepAlive(itpFirstParty)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LastUpdateTime: get the last time a KitITPThirdParty has been seen under
// itp_first_party. Each WebKitITPFirstParty is created by
// webkit_itp_third_party_get_first_parties() and therefore corresponds to
// exactly one KitITPThirdParty.
//
// The function returns the following values:
//
//    - dateTime: last update time as a Time.
//
func (itpFirstParty *ITPFirstParty) LastUpdateTime() *glib.DateTime {
	var _arg0 *C.WebKitITPFirstParty // out
	var _cret *C.GDateTime           // in

	_arg0 = (*C.WebKitITPFirstParty)(gextras.StructNative(unsafe.Pointer(itpFirstParty)))

	_cret = C.webkit_itp_first_party_get_last_update_time(_arg0)
	runtime.KeepAlive(itpFirstParty)

	var _dateTime *glib.DateTime // out

	_dateTime = (*glib.DateTime)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_date_time_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dateTime)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_date_time_unref((*C.GDateTime)(intern.C))
		},
	)

	return _dateTime
}

// WebsiteDataAccessAllowed: get whether itp_first_party has granted website
// data access to its KitITPThirdParty. Each WebKitITPFirstParty is created by
// webkit_itp_third_party_get_first_parties() and therefore corresponds to
// exactly one KitITPThirdParty.
//
// The function returns the following values:
//
//    - ok: TRUE if website data access has been granted, or FALSE otherwise.
//
func (itpFirstParty *ITPFirstParty) WebsiteDataAccessAllowed() bool {
	var _arg0 *C.WebKitITPFirstParty // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitITPFirstParty)(gextras.StructNative(unsafe.Pointer(itpFirstParty)))

	_cret = C.webkit_itp_first_party_get_website_data_access_allowed(_arg0)
	runtime.KeepAlive(itpFirstParty)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ITPThirdParty: instance of this type is always passed by reference.
type ITPThirdParty struct {
	*itpThirdParty
}

// itpThirdParty is the struct that's finalized.
type itpThirdParty struct {
	native *C.WebKitITPThirdParty
}

func marshalITPThirdParty(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ITPThirdParty{&itpThirdParty{(*C.WebKitITPThirdParty)(b)}}, nil
}

// Domain: get the domain name of itp_third_party.
//
// The function returns the following values:
//
//    - utf8: domain name.
//
func (itpThirdParty *ITPThirdParty) Domain() string {
	var _arg0 *C.WebKitITPThirdParty // out
	var _cret *C.char                // in

	_arg0 = (*C.WebKitITPThirdParty)(gextras.StructNative(unsafe.Pointer(itpThirdParty)))

	_cret = C.webkit_itp_third_party_get_domain(_arg0)
	runtime.KeepAlive(itpThirdParty)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// FirstParties: get the list of KitITPFirstParty under which itp_third_party
// has been seen.
//
// The function returns the following values:
//
//    - list of KitITPFirstParty.
//
func (itpThirdParty *ITPThirdParty) FirstParties() []*ITPFirstParty {
	var _arg0 *C.WebKitITPThirdParty // out
	var _cret *C.GList               // in

	_arg0 = (*C.WebKitITPThirdParty)(gextras.StructNative(unsafe.Pointer(itpThirdParty)))

	_cret = C.webkit_itp_third_party_get_first_parties(_arg0)
	runtime.KeepAlive(itpThirdParty)

	var _list []*ITPFirstParty // out

	_list = make([]*ITPFirstParty, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.WebKitITPFirstParty)(v)
		var dst *ITPFirstParty // out
		dst = (*ITPFirstParty)(gextras.NewStructNative(unsafe.Pointer(src)))
		C.webkit_itp_first_party_ref(src)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.webkit_itp_first_party_unref((*C.WebKitITPFirstParty)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}
