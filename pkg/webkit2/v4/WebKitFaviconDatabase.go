// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_favicon_database_error_get_type()), F: marshalFaviconDatabaseError},
		{T: externglib.Type(C.webkit_favicon_database_get_type()), F: marshalFaviconDatabaser},
	})
}

// FaviconDatabaseError: enum values used to denote the various errors related
// to the KitFaviconDatabase.
type FaviconDatabaseError int

const (
	// FaviconDatabaseErrorNotInitialized has not been initialized yet.
	FaviconDatabaseErrorNotInitialized FaviconDatabaseError = iota
	// FaviconDatabaseErrorFaviconNotFound: there is not an icon available for
	// the requested URL.
	FaviconDatabaseErrorFaviconNotFound
	// FaviconDatabaseErrorFaviconUnknown: there might be an icon for the
	// requested URL, but its data is unknown at the moment.
	FaviconDatabaseErrorFaviconUnknown
)

func marshalFaviconDatabaseError(p uintptr) (interface{}, error) {
	return FaviconDatabaseError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FaviconDatabaseError.
func (f FaviconDatabaseError) String() string {
	switch f {
	case FaviconDatabaseErrorNotInitialized:
		return "NotInitialized"
	case FaviconDatabaseErrorFaviconNotFound:
		return "FaviconNotFound"
	case FaviconDatabaseErrorFaviconUnknown:
		return "FaviconUnknown"
	default:
		return fmt.Sprintf("FaviconDatabaseError(%d)", f)
	}
}

type FaviconDatabase struct {
	*externglib.Object
}

func wrapFaviconDatabase(obj *externglib.Object) *FaviconDatabase {
	return &FaviconDatabase{
		Object: obj,
	}
}

func marshalFaviconDatabaser(p uintptr) (interface{}, error) {
	return wrapFaviconDatabase(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Clear clears all icons from the database.
func (database *FaviconDatabase) Clear() {
	var _arg0 *C.WebKitFaviconDatabase // out

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(database.Native()))

	C.webkit_favicon_database_clear(_arg0)
	runtime.KeepAlive(database)
}

// Favicon: asynchronously obtains a #cairo_surface_t of the favicon for the
// given page URI. It returns the cached icon if it's in the database
// asynchronously waiting for the icon to be read from the database.
//
// This is an asynchronous method. When the operation is finished, callback will
// be invoked. You can then call webkit_favicon_database_get_favicon_finish() to
// get the result of the operation.
//
// You must call webkit_web_context_set_favicon_database_directory() for the
// KitWebContext associated with this KitFaviconDatabase before attempting to
// use this function; otherwise, webkit_favicon_database_get_favicon_finish()
// will return WEBKIT_FAVICON_DATABASE_ERROR_NOT_INITIALIZED.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - pageUri: URI of the page for which we want to retrieve the favicon.
//    - callback to call when the request is satisfied or NULL if you don't
//    care about the result.
//
func (database *FaviconDatabase) Favicon(ctx context.Context, pageUri string, callback gio.AsyncReadyCallback) {
	var _arg0 *C.WebKitFaviconDatabase // out
	var _arg2 *C.GCancellable          // out
	var _arg1 *C.gchar                 // out
	var _arg3 C.GAsyncReadyCallback    // out
	var _arg4 C.gpointer

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(database.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pageUri)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.webkit_favicon_database_get_favicon(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(database)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(pageUri)
	runtime.KeepAlive(callback)
}

// FaviconFinish finishes an operation started with
// webkit_favicon_database_get_favicon().
//
// The function takes the following parameters:
//
//    - result obtained from the ReadyCallback passed to
//    webkit_favicon_database_get_favicon().
//
func (database *FaviconDatabase) FaviconFinish(result gio.AsyncResulter) (*cairo.Surface, error) {
	var _arg0 *C.WebKitFaviconDatabase // out
	var _arg1 *C.GAsyncResult          // out
	var _cret *C.cairo_surface_t       // in
	var _cerr *C.GError                // in

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(database.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.webkit_favicon_database_get_favicon_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(database)
	runtime.KeepAlive(result)

	var _surface *cairo.Surface // out
	var _goerr error            // out

	_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
		C.cairo_surface_destroy((*C.cairo_surface_t)(unsafe.Pointer(v.Native())))
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _surface, _goerr
}

// FaviconURI obtains the URI of the favicon for the given page_uri.
//
// The function takes the following parameters:
//
//    - pageUri: URI of the page containing the icon.
//
func (database *FaviconDatabase) FaviconURI(pageUri string) string {
	var _arg0 *C.WebKitFaviconDatabase // out
	var _arg1 *C.gchar                 // out
	var _cret *C.gchar                 // in

	_arg0 = (*C.WebKitFaviconDatabase)(unsafe.Pointer(database.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pageUri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.webkit_favicon_database_get_favicon_uri(_arg0, _arg1)
	runtime.KeepAlive(database)
	runtime.KeepAlive(pageUri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ConnectFaviconChanged: this signal is emitted when the favicon URI of
// page_uri has been changed to favicon_uri in the database. You can connect to
// this signal and call webkit_favicon_database_get_favicon() to get the
// favicon. If you are interested in the favicon of a KitWebView it's easier to
// use the KitWebView:favicon property. See webkit_web_view_get_favicon() for
// more details.
func (database *FaviconDatabase) ConnectFaviconChanged(f func(pageUri, faviconUri string)) externglib.SignalHandle {
	return database.Connect("favicon-changed", f)
}
