// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_request_file_get_type()), F: marshalRequestFiler},
	})
}

type RequestFile struct {
	Request
}

func wrapRequestFile(obj *externglib.Object) *RequestFile {
	return &RequestFile{
		Request: Request{
			Object: obj,
			Initable: gio.Initable{
				Object: obj,
			},
		},
	}
}

func marshalRequestFiler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRequestFile(obj), nil
}

// File gets a #GFile corresponding to file's URI
func (file *RequestFile) File() gio.Filer {
	var _arg0 *C.SoupRequestFile // out
	var _cret *C.GFile           // in

	_arg0 = (*C.SoupRequestFile)(unsafe.Pointer(file.Native()))

	_cret = C.soup_request_file_get_file(_arg0)
	runtime.KeepAlive(file)

	var _ret gio.Filer // out

	_ret = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.Filer)

	return _ret
}
