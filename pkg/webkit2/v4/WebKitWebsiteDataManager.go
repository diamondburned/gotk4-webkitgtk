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

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_tls_errors_policy_get_type()), F: marshalTLSErrorsPolicy},
		{T: externglib.Type(C.webkit_website_data_manager_get_type()), F: marshalWebsiteDataManagerer},
		{T: externglib.Type(C.webkit_itp_first_party_get_type()), F: marshalITPFirstParty},
		{T: externglib.Type(C.webkit_itp_third_party_get_type()), F: marshalITPThirdParty},
	})
}

// TLSErrorsPolicy: enum values used to denote the TLS errors policy.
type TLSErrorsPolicy int

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

type WebsiteDataManager struct {
	*externglib.Object
}

func wrapWebsiteDataManager(obj *externglib.Object) *WebsiteDataManager {
	return &WebsiteDataManager{
		Object: obj,
	}
}

func marshalWebsiteDataManagerer(p uintptr) (interface{}, error) {
	return wrapWebsiteDataManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWebsiteDataManagerEphemeral creates an ephemeral KitWebsiteDataManager.
// See KitWebsiteDataManager:is-ephemeral for more details.
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
//    - ctx or NULL to ignore.
//    - types: KitWebsiteDataTypes.
//    - timespan: Span.
//    - callback to call when the request is satisfied.
//
func (manager *WebsiteDataManager) Clear(ctx context.Context, types WebsiteDataTypes, timespan glib.TimeSpan, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg3 *C.GCancellable             // out
	var _arg1 C.WebKitWebsiteDataTypes    // out
	var _arg2 C.GTimeSpan                 // out
	var _arg4 C.GAsyncReadyCallback       // out
	var _arg5 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

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
//    - ctx or NULL to ignore.
//    - types: KitWebsiteDataTypes.
//    - callback to call when the request is satisfied.
//
func (manager *WebsiteDataManager) Fetch(ctx context.Context, types WebsiteDataTypes, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg2 *C.GCancellable             // out
	var _arg1 C.WebKitWebsiteDataTypes    // out
	var _arg3 C.GAsyncReadyCallback       // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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
func (manager *WebsiteDataManager) FetchFinish(result gio.AsyncResulter) ([]*WebsiteData, error) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GList                    // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

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
func (manager *WebsiteDataManager) BaseCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) BaseDataDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

	_cret = C.webkit_website_data_manager_get_base_data_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CookieManager: get the KitCookieManager of manager.
func (manager *WebsiteDataManager) CookieManager() *CookieManager {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.WebKitCookieManager      // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

	_cret = C.webkit_website_data_manager_get_cookie_manager(_arg0)
	runtime.KeepAlive(manager)

	var _cookieManager *CookieManager // out

	_cookieManager = wrapCookieManager(externglib.Take(unsafe.Pointer(_cret)))

	return _cookieManager
}

// DiskCacheDirectory: get the KitWebsiteDataManager:disk-cache-directory
// property.
func (manager *WebsiteDataManager) DiskCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) DomCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) HstsCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) IndexeddbDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

	_cret = C.webkit_website_data_manager_get_indexeddb_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ITPDirectory: get the KitWebsiteDataManager:itp-directory property.
func (manager *WebsiteDataManager) ITPDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) ITPEnabled() bool {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
//    - ctx or NULL to ignore.
//    - callback to call when the request is satisfied.
//
func (manager *WebsiteDataManager) ITPSummary(ctx context.Context, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GCancellable             // out
	var _arg2 C.GAsyncReadyCallback       // out
	var _arg3 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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
func (manager *WebsiteDataManager) ITPSummaryFinish(result gio.AsyncResulter) ([]*ITPThirdParty, error) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GList                    // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

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
func (manager *WebsiteDataManager) LocalStorageDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) OfflineApplicationCacheDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) PersistentCredentialStorageEnabled() bool {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) ServiceWorkerRegistrationsDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

	_cret = C.webkit_website_data_manager_get_service_worker_registrations_directory(_arg0)
	runtime.KeepAlive(manager)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TLSErrorsPolicy: get the TLS errors policy of manager.
func (manager *WebsiteDataManager) TLSErrorsPolicy() TLSErrorsPolicy {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.WebKitTLSErrorsPolicy     // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

	_cret = C.webkit_website_data_manager_get_tls_errors_policy(_arg0)
	runtime.KeepAlive(manager)

	var _tlsErrorsPolicy TLSErrorsPolicy // out

	_tlsErrorsPolicy = TLSErrorsPolicy(_cret)

	return _tlsErrorsPolicy
}

// WebsqlDirectory: get the KitWebsiteDataManager:websql-directory property.
//
// Deprecated: WebSQL is no longer supported. Use IndexedDB instead.
func (manager *WebsiteDataManager) WebsqlDirectory() string {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
func (manager *WebsiteDataManager) IsEphemeral() bool {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))

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
//    - ctx or NULL to ignore.
//    - types: KitWebsiteDataTypes.
//    - websiteData of KitWebsiteData.
//    - callback to call when the request is satisfied.
//
func (manager *WebsiteDataManager) Remove(ctx context.Context, types WebsiteDataTypes, websiteData []*WebsiteData, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitWebsiteDataManager // out
	var _arg3 *C.GCancellable             // out
	var _arg1 C.WebKitWebsiteDataTypes    // out
	var _arg2 *C.GList                    // out
	var _arg4 C.GAsyncReadyCallback       // out
	var _arg5 C.gpointer

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

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

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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
//    - proxySettings or NULL.
//
func (manager *WebsiteDataManager) SetNetworkProxySettings(proxyMode NetworkProxyMode, proxySettings *NetworkProxySettings) {
	var _arg0 *C.WebKitWebsiteDataManager   // out
	var _arg1 C.WebKitNetworkProxyMode      // out
	var _arg2 *C.WebKitNetworkProxySettings // out

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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

	_arg0 = (*C.WebKitWebsiteDataManager)(unsafe.Pointer(manager.Native()))
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
