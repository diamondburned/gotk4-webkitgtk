// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_multipart_get_type()), F: marshalMultipart},
	})
}

// Multipart represents a multipart HTTP message body, parsed according to the
// syntax of RFC 2046. Of particular interest to HTTP are
// <literal>multipart/byte-ranges</literal> and
// <literal>multipart/form-data</literal>.
//
// Although the headers of a Multipart body part will contain the full headers
// from that body part, libsoup does not interpret them according to MIME rules.
// For example, each body part is assumed to have "binary"
// Content-Transfer-Encoding, even if its headers explicitly state otherwise. In
// other words, don't try to use Multipart for handling real MIME multiparts.
//
// An instance of this type is always passed by reference.
type Multipart struct {
	*multipart
}

// multipart is the struct that's finalized.
type multipart struct {
	native *C.SoupMultipart
}

func marshalMultipart(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Multipart{&multipart{(*C.SoupMultipart)(b)}}, nil
}

// NewMultipart constructs a struct Multipart.
func NewMultipart(mimeType string) *Multipart {
	var _arg1 *C.char          // out
	var _cret *C.SoupMultipart // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_multipart_new(_arg1)
	runtime.KeepAlive(mimeType)

	var _multipart *Multipart // out

	_multipart = (*Multipart)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_multipart)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_multipart_free((*C.SoupMultipart)(intern.C))
		},
	)

	return _multipart
}

// NewMultipartFromMessage constructs a struct Multipart.
func NewMultipartFromMessage(headers *MessageHeaders, body *MessageBody) *Multipart {
	var _arg1 *C.SoupMessageHeaders // out
	var _arg2 *C.SoupMessageBody    // out
	var _cret *C.SoupMultipart      // in

	_arg1 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(headers)))
	_arg2 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))

	_cret = C.soup_multipart_new_from_message(_arg1, _arg2)
	runtime.KeepAlive(headers)
	runtime.KeepAlive(body)

	var _multipart *Multipart // out

	if _cret != nil {
		_multipart = (*Multipart)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_multipart)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_multipart_free((*C.SoupMultipart)(intern.C))
			},
		)
	}

	return _multipart
}

// AppendFormFile adds a new MIME part containing body to multipart, using
// "Content-Disposition: form-data", as per the HTML forms specification. See
// soup_form_request_new_from_multipart() for more details.
func (multipart *Multipart) AppendFormFile(controlName string, filename string, contentType string, body *Buffer) {
	var _arg0 *C.SoupMultipart // out
	var _arg1 *C.char          // out
	var _arg2 *C.char          // out
	var _arg3 *C.char          // out
	var _arg4 *C.SoupBuffer    // out

	_arg0 = (*C.SoupMultipart)(gextras.StructNative(unsafe.Pointer(multipart)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(controlName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(body)))

	C.soup_multipart_append_form_file(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(controlName)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(contentType)
	runtime.KeepAlive(body)
}

// AppendFormString adds a new MIME part containing data to multipart, using
// "Content-Disposition: form-data", as per the HTML forms specification. See
// soup_form_request_new_from_multipart() for more details.
func (multipart *Multipart) AppendFormString(controlName string, data string) {
	var _arg0 *C.SoupMultipart // out
	var _arg1 *C.char          // out
	var _arg2 *C.char          // out

	_arg0 = (*C.SoupMultipart)(gextras.StructNative(unsafe.Pointer(multipart)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(controlName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(data)))
	defer C.free(unsafe.Pointer(_arg2))

	C.soup_multipart_append_form_string(_arg0, _arg1, _arg2)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(controlName)
	runtime.KeepAlive(data)
}

// AppendPart adds a new MIME part to multipart with the given headers and body.
// (The multipart will make its own copies of headers and body, so you should
// free your copies if you are not using them for anything else.).
func (multipart *Multipart) AppendPart(headers *MessageHeaders, body *Buffer) {
	var _arg0 *C.SoupMultipart      // out
	var _arg1 *C.SoupMessageHeaders // out
	var _arg2 *C.SoupBuffer         // out

	_arg0 = (*C.SoupMultipart)(gextras.StructNative(unsafe.Pointer(multipart)))
	_arg1 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(headers)))
	_arg2 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(body)))

	C.soup_multipart_append_part(_arg0, _arg1, _arg2)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(headers)
	runtime.KeepAlive(body)
}

// Length gets the number of body parts in multipart.
func (multipart *Multipart) Length() int {
	var _arg0 *C.SoupMultipart // out
	var _cret C.int            // in

	_arg0 = (*C.SoupMultipart)(gextras.StructNative(unsafe.Pointer(multipart)))

	_cret = C.soup_multipart_get_length(_arg0)
	runtime.KeepAlive(multipart)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Part gets the indicated body part from multipart.
func (multipart *Multipart) Part(part int) (*MessageHeaders, *Buffer, bool) {
	var _arg0 *C.SoupMultipart      // out
	var _arg1 C.int                 // out
	var _arg2 *C.SoupMessageHeaders // in
	var _arg3 *C.SoupBuffer         // in
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupMultipart)(gextras.StructNative(unsafe.Pointer(multipart)))
	_arg1 = C.int(part)

	_cret = C.soup_multipart_get_part(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(part)

	var _headers *MessageHeaders // out
	var _body *Buffer            // out
	var _ok bool                 // out

	_headers = (*MessageHeaders)(gextras.NewStructNative(unsafe.Pointer(_arg2)))
	_body = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(_arg3)))
	if _cret != 0 {
		_ok = true
	}

	return _headers, _body, _ok
}

// ToMessage serializes multipart to dest_headers and dest_body.
func (multipart *Multipart) ToMessage(destHeaders *MessageHeaders, destBody *MessageBody) {
	var _arg0 *C.SoupMultipart      // out
	var _arg1 *C.SoupMessageHeaders // out
	var _arg2 *C.SoupMessageBody    // out

	_arg0 = (*C.SoupMultipart)(gextras.StructNative(unsafe.Pointer(multipart)))
	_arg1 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(destHeaders)))
	_arg2 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(destBody)))

	C.soup_multipart_to_message(_arg0, _arg1, _arg2)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(destHeaders)
	runtime.KeepAlive(destBody)
}
