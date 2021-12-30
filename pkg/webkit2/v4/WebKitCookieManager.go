// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4-webkitgtk/pkg/soup/v2"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_cookie_accept_policy_get_type()), F: marshalCookieAcceptPolicy},
		{T: externglib.Type(C.webkit_cookie_persistent_storage_get_type()), F: marshalCookiePersistentStorage},
		{T: externglib.Type(C.webkit_cookie_manager_get_type()), F: marshalCookieManagerer},
	})
}

// CookieAcceptPolicy: enum values used to denote the cookie acceptance
// policies.
type CookieAcceptPolicy C.gint

const (
	// CookiePolicyAcceptAlways: accept all cookies unconditionally.
	CookiePolicyAcceptAlways CookieAcceptPolicy = iota
	// CookiePolicyAcceptNever: reject all cookies unconditionally.
	CookiePolicyAcceptNever
	// CookiePolicyAcceptNoThirdParty: accept only cookies set by the main
	// document loaded.
	CookiePolicyAcceptNoThirdParty
)

func marshalCookieAcceptPolicy(p uintptr) (interface{}, error) {
	return CookieAcceptPolicy(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CookieAcceptPolicy.
func (c CookieAcceptPolicy) String() string {
	switch c {
	case CookiePolicyAcceptAlways:
		return "Always"
	case CookiePolicyAcceptNever:
		return "Never"
	case CookiePolicyAcceptNoThirdParty:
		return "NoThirdParty"
	default:
		return fmt.Sprintf("CookieAcceptPolicy(%d)", c)
	}
}

// CookiePersistentStorage: enum values used to denote the cookie persistent
// storage types.
type CookiePersistentStorage C.gint

const (
	// CookiePersistentStorageText cookies are stored in a text file in the
	// Mozilla "cookies.txt" format.
	CookiePersistentStorageText CookiePersistentStorage = iota
	// CookiePersistentStorageSqlite cookies are stored in a SQLite file in the
	// current Mozilla format.
	CookiePersistentStorageSqlite
)

func marshalCookiePersistentStorage(p uintptr) (interface{}, error) {
	return CookiePersistentStorage(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CookiePersistentStorage.
func (c CookiePersistentStorage) String() string {
	switch c {
	case CookiePersistentStorageText:
		return "Text"
	case CookiePersistentStorageSqlite:
		return "Sqlite"
	default:
		return fmt.Sprintf("CookiePersistentStorage(%d)", c)
	}
}

type CookieManager struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*CookieManager)(nil)
)

func wrapCookieManager(obj *externglib.Object) *CookieManager {
	return &CookieManager{
		Object: obj,
	}
}

func marshalCookieManagerer(p uintptr) (interface{}, error) {
	return wrapCookieManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged: this signal is emitted when cookies are added, removed or
// modified.
func (cookieManager *CookieManager) ConnectChanged(f func()) externglib.SignalHandle {
	return cookieManager.Connect("changed", f)
}

// AddCookie: asynchronously add a Cookie to the underlying storage.
//
// When the operation is finished, callback will be called. You can then call
// webkit_cookie_manager_add_cookie_finish() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - cookie to be added.
//    - callback (optional) to call when the request is satisfied.
//
func (cookieManager *CookieManager) AddCookie(ctx context.Context, cookie *soup.Cookie, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg2 *C.GCancellable        // out
	var _arg1 *C.SoupCookie          // out
	var _arg3 C.GAsyncReadyCallback  // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_cookie_manager_add_cookie(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(callback)
}

// AddCookieFinish: finish an asynchronous operation started with
// webkit_cookie_manager_add_cookie().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (cookieManager *CookieManager) AddCookieFinish(result gio.AsyncResulter) error {
	var _arg0 *C.WebKitCookieManager // out
	var _arg1 *C.GAsyncResult        // out
	var _cerr *C.GError              // in

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.webkit_cookie_manager_add_cookie_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DeleteAllCookies: delete all cookies of cookie_manager
//
// Deprecated: Use webkit_website_data_manager_clear() instead.
func (cookieManager *CookieManager) DeleteAllCookies() {
	var _arg0 *C.WebKitCookieManager // out

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))

	C.webkit_cookie_manager_delete_all_cookies(_arg0)
	runtime.KeepAlive(cookieManager)
}

// DeleteCookie: asynchronously delete a Cookie from the current session.
//
// When the operation is finished, callback will be called. You can then call
// webkit_cookie_manager_delete_cookie_finish() to get the result of the
// operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - cookie to be deleted.
//    - callback (optional) to call when the request is satisfied.
//
func (cookieManager *CookieManager) DeleteCookie(ctx context.Context, cookie *soup.Cookie, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg2 *C.GCancellable        // out
	var _arg1 *C.SoupCookie          // out
	var _arg3 C.GAsyncReadyCallback  // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.SoupCookie)(gextras.StructNative(unsafe.Pointer(cookie)))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_cookie_manager_delete_cookie(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(cookie)
	runtime.KeepAlive(callback)
}

// DeleteCookieFinish: finish an asynchronous operation started with
// webkit_cookie_manager_delete_cookie().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (cookieManager *CookieManager) DeleteCookieFinish(result gio.AsyncResulter) error {
	var _arg0 *C.WebKitCookieManager // out
	var _arg1 *C.GAsyncResult        // out
	var _cerr *C.GError              // in

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.webkit_cookie_manager_delete_cookie_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DeleteCookiesForDomain: remove all cookies of cookie_manager for the given
// domain.
//
// Deprecated: Use webkit_website_data_manager_remove() instead.
//
// The function takes the following parameters:
//
//    - domain name.
//
func (cookieManager *CookieManager) DeleteCookiesForDomain(domain string) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))

	C.webkit_cookie_manager_delete_cookies_for_domain(_arg0, _arg1)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(domain)
}

// AcceptPolicy: asynchronously get the cookie acceptance policy of
// cookie_manager. Note that when policy was set to
// WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY and ITP is enabled, this will
// return WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS. See also
// webkit_website_data_manager_set_itp_enabled().
//
// When the operation is finished, callback will be called. You can then call
// webkit_cookie_manager_get_accept_policy_finish() to get the result of the
// operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - callback (optional) to call when the request is satisfied.
//
func (cookieManager *CookieManager) AcceptPolicy(ctx context.Context, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg1 *C.GCancellable        // out
	var _arg2 C.GAsyncReadyCallback  // out
	var _arg3 C.gpointer

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_cookie_manager_get_accept_policy(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// AcceptPolicyFinish: finish an asynchronous operation started with
// webkit_cookie_manager_get_accept_policy().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - cookieAcceptPolicy: cookie acceptance policy of cookie_manager as a
//      KitCookieAcceptPolicy.
//
func (cookieManager *CookieManager) AcceptPolicyFinish(result gio.AsyncResulter) (CookieAcceptPolicy, error) {
	var _arg0 *C.WebKitCookieManager     // out
	var _arg1 *C.GAsyncResult            // out
	var _cret C.WebKitCookieAcceptPolicy // in
	var _cerr *C.GError                  // in

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.webkit_cookie_manager_get_accept_policy_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(result)

	var _cookieAcceptPolicy CookieAcceptPolicy // out
	var _goerr error                           // out

	_cookieAcceptPolicy = CookieAcceptPolicy(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _cookieAcceptPolicy, _goerr
}

// Cookies: asynchronously get a list of Cookie from cookie_manager associated
// with uri, which must be either an HTTP or an HTTPS URL.
//
// When the operation is finished, callback will be called. You can then call
// webkit_cookie_manager_get_cookies_finish() to get the result of the
// operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - uri: URI associated to the cookies to be retrieved.
//    - callback (optional) to call when the request is satisfied.
//
func (cookieManager *CookieManager) Cookies(ctx context.Context, uri string, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg2 *C.GCancellable        // out
	var _arg1 *C.gchar               // out
	var _arg3 C.GAsyncReadyCallback  // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_cookie_manager_get_cookies(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(callback)
}

// CookiesFinish: finish an asynchronous operation started with
// webkit_cookie_manager_get_cookies(). The return value is a List of Cookie
// instances which should be released with g_list_free_full() and
// soup_cookie_free().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - list of Cookie instances.
//
func (cookieManager *CookieManager) CookiesFinish(result gio.AsyncResulter) ([]*soup.Cookie, error) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg1 *C.GAsyncResult        // out
	var _cret *C.GList               // in
	var _cerr *C.GError              // in

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.webkit_cookie_manager_get_cookies_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(result)

	var _list []*soup.Cookie // out
	var _goerr error         // out

	_list = make([]*soup.Cookie, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.SoupCookie)(v)
		var dst *soup.Cookie // out
		dst = (*soup.Cookie)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_cookie_free((*C.SoupCookie)(intern.C))
			},
		)
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// DomainsWithCookies: asynchronously get the list of domains for which
// cookie_manager contains cookies.
//
// When the operation is finished, callback will be called. You can then call
// webkit_cookie_manager_get_domains_with_cookies_finish() to get the result of
// the operation.
//
// Deprecated: Use webkit_website_data_manager_fetch() instead.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL to ignore.
//    - callback (optional) to call when the request is satisfied.
//
func (cookieManager *CookieManager) DomainsWithCookies(ctx context.Context, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg1 *C.GCancellable        // out
	var _arg2 C.GAsyncReadyCallback  // out
	var _arg3 C.gpointer

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_cookie_manager_get_domains_with_cookies(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// DomainsWithCookiesFinish: finish an asynchronous operation started with
// webkit_cookie_manager_get_domains_with_cookies(). The return value is a NULL
// terminated list of strings which should be released with g_strfreev().
//
// Deprecated: Use webkit_website_data_manager_fetch_finish() instead.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - utf8s: NULL terminated array of domain names or NULL in case of error.
//
func (cookieManager *CookieManager) DomainsWithCookiesFinish(result gio.AsyncResulter) ([]string, error) {
	var _arg0 *C.WebKitCookieManager // out
	var _arg1 *C.GAsyncResult        // out
	var _cret **C.gchar              // in
	var _cerr *C.GError              // in

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.webkit_cookie_manager_get_domains_with_cookies_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(result)

	var _utf8s []string // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8s, _goerr
}

// SetAcceptPolicy: set the cookie acceptance policy of cookie_manager as
// policy. Note that ITP has its own way to handle third-party cookies, so when
// it's enabled, and policy is set to
// WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY,
// WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS will be used instead. Once disabled, the
// policy will be set back to WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY. See
// also webkit_website_data_manager_set_itp_enabled().
//
// The function takes the following parameters:
//
//    - policy: KitCookieAcceptPolicy.
//
func (cookieManager *CookieManager) SetAcceptPolicy(policy CookieAcceptPolicy) {
	var _arg0 *C.WebKitCookieManager     // out
	var _arg1 C.WebKitCookieAcceptPolicy // out

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = C.WebKitCookieAcceptPolicy(policy)

	C.webkit_cookie_manager_set_accept_policy(_arg0, _arg1)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(policy)
}

// SetPersistentStorage: set the filename where non-session cookies are stored
// persistently using storage as the format to read/write the cookies. Cookies
// are initially read from filename to create an initial set of cookies. Then,
// non-session cookies will be written to filename when the
// WebKitCookieManager::changed signal is emitted. By default, cookie_manager
// doesn't store the cookies persistently, so you need to call this method to
// keep cookies saved across sessions.
//
// This method should never be called on a KitCookieManager associated to an
// ephemeral KitWebsiteDataManager.
//
// The function takes the following parameters:
//
//    - filename to read to/write from.
//    - storage: KitCookiePersistentStorage.
//
func (cookieManager *CookieManager) SetPersistentStorage(filename string, storage CookiePersistentStorage) {
	var _arg0 *C.WebKitCookieManager          // out
	var _arg1 *C.gchar                        // out
	var _arg2 C.WebKitCookiePersistentStorage // out

	_arg0 = (*C.WebKitCookieManager)(unsafe.Pointer(cookieManager.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.WebKitCookiePersistentStorage(storage)

	C.webkit_cookie_manager_set_persistent_storage(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cookieManager)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(storage)
}
