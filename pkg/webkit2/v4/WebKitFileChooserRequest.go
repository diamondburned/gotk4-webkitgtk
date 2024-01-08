// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// GType values.
var (
	GTypeFileChooserRequest = coreglib.Type(C.webkit_file_chooser_request_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileChooserRequest, F: marshalFileChooserRequest},
	})
}

// FileChooserRequestOverrides contains methods that are overridable.
type FileChooserRequestOverrides struct {
}

func defaultFileChooserRequestOverrides(v *FileChooserRequest) FileChooserRequestOverrides {
	return FileChooserRequestOverrides{}
}

// FileChooserRequest: request to open a file chooser.
//
// Whenever the user interacts with an HTML input element with file type,
// WebKit will need to show a dialog to choose one or more files to be uploaded
// to the server along with the rest of the form data. For that to happen in a
// general way, instead of just opening a FileChooserDialog (which might be not
// desirable in some cases, which could prefer to use their own file chooser
// dialog), WebKit will fire the KitWebView::run-file-chooser signal with a
// KitFileChooserRequest object, which will allow the client application to
// specify the files to be selected, to inspect the details of the request (e.g.
// if multiple selection should be allowed) and to cancel the request, in case
// nothing was selected.
//
// In case the client application does not wish to handle this signal,
// WebKit will provide a default handler which will asynchronously run a regular
// FileChooserDialog for the user to interact with.
type FileChooserRequest struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FileChooserRequest)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FileChooserRequest, *FileChooserRequestClass, FileChooserRequestOverrides](
		GTypeFileChooserRequest,
		initFileChooserRequestClass,
		wrapFileChooserRequest,
		defaultFileChooserRequestOverrides,
	)
}

func initFileChooserRequestClass(gclass unsafe.Pointer, overrides FileChooserRequestOverrides, classInitFunc func(*FileChooserRequestClass)) {
	if classInitFunc != nil {
		class := (*FileChooserRequestClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFileChooserRequest(obj *coreglib.Object) *FileChooserRequest {
	return &FileChooserRequest{
		Object: obj,
	}
}

func marshalFileChooserRequest(p uintptr) (interface{}, error) {
	return wrapFileChooserRequest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Cancel: ask WebKit to cancel the request.
//
// It's important to do this in case no selection has been made in the client,
// otherwise the request won't be properly completed and the browser will keep
// the request pending forever, which might cause the browser to hang.
func (request *FileChooserRequest) Cancel() {
	var _arg0 *C.WebKitFileChooserRequest // out

	_arg0 = (*C.WebKitFileChooserRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	C.webkit_file_chooser_request_cancel(_arg0)
	runtime.KeepAlive(request)
}

// MIMETypes: get the list of MIME types the file chooser dialog should handle.
//
// Get the list of MIME types the file chooser dialog should handle, in the
// format specified in RFC 2046 for "media types". Its contents depend on the
// value of the 'accept' attribute for HTML input elements. This function should
// normally be called before presenting the file chooser dialog to the user,
// to decide whether to allow the user to select multiple files at once or only
// one.
//
// The function returns the following values:
//
//   - utf8s: a NULL-terminated array of strings if a list of accepted MIME
//     types is defined or NULL otherwise, meaning that any MIME type should be
//     accepted. This array and its contents are owned by WebKit and should not
//     be modified or freed.
//
func (request *FileChooserRequest) MIMETypes() []string {
	var _arg0 *C.WebKitFileChooserRequest // out
	var _cret **C.gchar                   // in

	_arg0 = (*C.WebKitFileChooserRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_file_chooser_request_get_mime_types(_arg0)
	runtime.KeepAlive(request)

	var _utf8s []string // out

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
		}
	}

	return _utf8s
}

// MIMETypesFilter: get the filter currently associated with the request.
//
// Get the filter currently associated with the request, ready to be used by
// FileChooser. This function should normally be called before presenting the
// file chooser dialog to the user, to decide whether to apply a filter so the
// user would not be allowed to select files with other MIME types.
//
// See webkit_file_chooser_request_get_mime_types() if you are interested in
// getting the list of accepted MIME types.
//
// The function returns the following values:
//
//   - fileFilter if a list of accepted MIME types is defined or NULL otherwise.
//     The returned object is owned by WebKit should not be modified or freed.
//
func (request *FileChooserRequest) MIMETypesFilter() *gtk.FileFilter {
	var _arg0 *C.WebKitFileChooserRequest // out
	var _cret *C.GtkFileFilter            // in

	_arg0 = (*C.WebKitFileChooserRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_file_chooser_request_get_mime_types_filter(_arg0)
	runtime.KeepAlive(request)

	var _fileFilter *gtk.FileFilter // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_fileFilter = &gtk.FileFilter{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Buildable: gtk.Buildable{
				Object: obj,
			},
		}
	}

	return _fileFilter
}

// SelectMultiple: whether the file chooser should allow selecting multiple
// files.
//
// Determine whether the file chooser associated to this KitFileChooserRequest
// should allow selecting multiple files, which depends on the HTML input
// element having a 'multiple' attribute defined.
//
// The function returns the following values:
//
//   - ok: TRUE if the file chooser should allow selecting multiple files or
//     FALSE otherwise.
//
func (request *FileChooserRequest) SelectMultiple() bool {
	var _arg0 *C.WebKitFileChooserRequest // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.WebKitFileChooserRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_file_chooser_request_get_select_multiple(_arg0)
	runtime.KeepAlive(request)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectedFiles: get the list of selected files associated to the request.
//
// Get the list of selected files currently associated to the request.
// Initially, the return value of this method contains any files selected
// in previous file chooser requests for this HTML input element. Once
// webkit_file_chooser_request_select_files, the value will reflect whatever
// files are given.
//
// This function should normally be called only before presenting the file
// chooser dialog to the user, to decide whether to perform some extra action,
// like pre-selecting the files from a previous request.
//
// The function returns the following values:
//
//   - utf8s: a NULL-terminated array of strings if there are selected files
//     associated with the request or NULL otherwise. This array and its
//     contents are owned by WebKit and should not be modified or freed.
//
func (request *FileChooserRequest) SelectedFiles() []string {
	var _arg0 *C.WebKitFileChooserRequest // out
	var _cret **C.gchar                   // in

	_arg0 = (*C.WebKitFileChooserRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))

	_cret = C.webkit_file_chooser_request_get_selected_files(_arg0)
	runtime.KeepAlive(request)

	var _utf8s []string // out

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
		}
	}

	return _utf8s
}

// SelectFiles: ask WebKit to select local files for upload and complete the
// request.
//
// The function takes the following parameters:
//
//   - files: a NULL-terminated array of strings, containing paths to local
//     files.
//
func (request *FileChooserRequest) SelectFiles(files []string) {
	var _arg0 *C.WebKitFileChooserRequest // out
	var _arg1 **C.gchar                   // out

	_arg0 = (*C.WebKitFileChooserRequest)(unsafe.Pointer(coreglib.InternObject(request).Native()))
	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(files) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(files)+1)
			var zero *C.gchar
			out[len(files)] = zero
			for i := range files {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(files[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.webkit_file_chooser_request_select_files(_arg0, _arg1)
	runtime.KeepAlive(request)
	runtime.KeepAlive(files)
}

// FileChooserRequestClass: instance of this type is always passed by reference.
type FileChooserRequestClass struct {
	*fileChooserRequestClass
}

// fileChooserRequestClass is the struct that's finalized.
type fileChooserRequestClass struct {
	native *C.WebKitFileChooserRequestClass
}
