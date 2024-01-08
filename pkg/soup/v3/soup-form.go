// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

// FORM_MIME_TYPE_MULTIPART: macro containing the value multipart/form-data;
// the MIME type used for posting form data that contains files to be uploaded.
const FORM_MIME_TYPE_MULTIPART = "multipart/form-data"

// FORM_MIME_TYPE_URLENCODED: macro containing the value
// application/x-www-form-urlencoded; the default MIME type for POSTing HTML
// form data.
const FORM_MIME_TYPE_URLENCODED = "application/x-www-form-urlencoded"

// FormDecode decodes form.
//
// which is an urlencoded dataset as defined in the HTML 4.01 spec.
//
// The function takes the following parameters:
//
//   - encodedForm: data of type "application/x-www-form-urlencoded".
//
// The function returns the following values:
//
//   - hashTable: hash table containing the name/value pairs from encoded_form,
//     which you can free with glib.HashTable().Destroy.
//
func FormDecode(encodedForm string) map[string]string {
	var _arg1 *C.char       // out
	var _cret *C.GHashTable // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(encodedForm)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_form_decode(_arg1)
	runtime.KeepAlive(encodedForm)

	var _hashTable map[string]string // out

	_hashTable = make(map[string]string, gextras.HashTableSize(unsafe.Pointer(_cret)))
	gextras.MoveHashTable(unsafe.Pointer(_cret), true, func(k, v unsafe.Pointer) {
		ksrc := *(**C.gchar)(k)
		vsrc := *(**C.gchar)(v)
		var kdst string // out
		var vdst string // out
		kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
		vdst = C.GoString((*C.gchar)(unsafe.Pointer(vsrc)))
		_hashTable[kdst] = vdst
	})

	return _hashTable
}

// FormDecodeMultipart decodes the "multipart/form-data" request in multipart.
//
// this is a convenience method for the case when you have a single file
// upload control in a form. (Or when you don't have any file upload controls,
// but are still using "multipart/form-data" anyway.) Pass the name of the file
// upload control in file_control_name, and form_decode_multipart will extract
// the uploaded file data into filename, content_type, and file. All of the
// other form control data will be returned (as strings, as with form_decode in
// the returned glib.HashTable.
//
// You may pass NULL for filename, content_type and/or file if you do not care
// about those fields. form_decode_multipart may also return NULL in those
// fields if the client did not provide that information. You must free the
// returned filename and content-type with glib.Free(), and the returned file
// data with glib.Bytes.Unref().
//
// If you have a form with more than one file upload control, you will need to
// decode it manually, using multipart.NewFromMessage and multipart.GetPart.
//
// The function takes the following parameters:
//
//   - multipart: Multipart.
//   - fileControlName (optional): name of the HTML file upload control.
//
// The function returns the following values:
//
//   - filename (optional): return location for the name of the uploaded file.
//   - contentType (optional): return location for the MIME type of the uploaded
//     file.
//   - file (optional): return location for the uploaded file data.
//   - hashTable (optional): a hash table containing the name/value pairs
//     (other than file_control_name) from msg, which you can free with
//     glib.HashTable().Destroy. On error, it will return NULL.
//
func FormDecodeMultipart(multipart *Multipart, fileControlName string) (filename, contentType string, file *glib.Bytes, hashTable map[string]string) {
	var _arg1 *C.SoupMultipart // out
	var _arg2 *C.char          // out
	var _arg3 *C.char          // in
	var _arg4 *C.char          // in
	var _arg5 *C.GBytes        // in
	var _cret *C.GHashTable    // in

	_arg1 = (*C.SoupMultipart)(gextras.StructNative(unsafe.Pointer(multipart)))
	if fileControlName != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(fileControlName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.soup_form_decode_multipart(_arg1, _arg2, &_arg3, &_arg4, &_arg5)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(fileControlName)

	var _filename string             // out
	var _contentType string          // out
	var _file *glib.Bytes            // out
	var _hashTable map[string]string // out

	if _arg3 != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_arg3)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	if _arg4 != nil {
		_contentType = C.GoString((*C.gchar)(unsafe.Pointer(_arg4)))
		defer C.free(unsafe.Pointer(_arg4))
	}
	if _arg5 != nil {
		_file = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_arg5)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_file)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_bytes_unref((*C.GBytes)(intern.C))
			},
		)
	}
	if _cret != nil {
		_hashTable = make(map[string]string, gextras.HashTableSize(unsafe.Pointer(_cret)))
		gextras.MoveHashTable(unsafe.Pointer(_cret), true, func(k, v unsafe.Pointer) {
			ksrc := *(**C.gchar)(k)
			vsrc := *(**C.gchar)(v)
			var kdst string // out
			var vdst string // out
			kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
			vdst = C.GoString((*C.gchar)(unsafe.Pointer(vsrc)))
			_hashTable[kdst] = vdst
		})
	}

	return _filename, _contentType, _file, _hashTable
}

// FormEncodeHash encodes form_data_set into a value of type
// "application/x-www-form-urlencoded".
//
// Encodes as defined in the HTML 4.01 spec.
//
// Note that the HTML spec states that "The control names/values are listed in
// the order they appear in the document." Since this method takes a hash table,
// it cannot enforce that; if you care about the ordering of the form fields,
// use form_encode_datalist.
//
// See also: message.NewFromEncodedForm.
//
// The function takes the following parameters:
//
//   - formDataSet: hash table containing name/value pairs (as strings).
//
// The function returns the following values:
//
//   - utf8: encoded form.
//
func FormEncodeHash(formDataSet map[string]string) string {
	var _arg1 *C.GHashTable // out
	var _cret *C.char       // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range formDataSet {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		defer C.free(unsafe.Pointer(vdst))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)

	_cret = C.soup_form_encode_hash(_arg1)
	runtime.KeepAlive(formDataSet)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}