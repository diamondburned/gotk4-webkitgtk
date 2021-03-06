// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern char* _gotk4_soup2_ContentSnifferClass_sniff(SoupContentSniffer*, SoupMessage*, SoupBuffer*, GHashTable**);
// extern gsize _gotk4_soup2_ContentSnifferClass_get_buffer_size(SoupContentSniffer*);
import "C"

// glib.Type values for soup-content-sniffer.go.
var GTypeContentSniffer = externglib.Type(C.soup_content_sniffer_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeContentSniffer, F: marshalContentSniffer},
	})
}

// ContentSnifferOverrider contains methods that are overridable.
type ContentSnifferOverrider interface {
	// BufferSize gets the number of bytes sniffer needs in order to properly
	// sniff a buffer.
	//
	// The function returns the following values:
	//
	//    - gsize: number of bytes to sniff.
	//
	BufferSize() uint
	// Sniff sniffs buffer to determine its Content-Type. The result may also be
	// influenced by the Content-Type declared in msg's response headers.
	//
	// The function takes the following parameters:
	//
	//    - msg: message to sniff.
	//    - buffer containing the start of msg's response body.
	//
	// The function returns the following values:
	//
	//    - params (optional): return location for Content-Type parameters (eg,
	//      "charset"), or NULL.
	//    - utf8: sniffed Content-Type of buffer; this will never be NULL, but
	//      may be "application/octet-stream".
	//
	Sniff(msg *Message, buffer *Buffer) (map[string]string, string)
}

type ContentSniffer struct {
	_ [0]func() // equal guard
	*externglib.Object

	SessionFeature
}

var (
	_ externglib.Objector = (*ContentSniffer)(nil)
)

func classInitContentSnifferer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.SoupContentSnifferClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.SoupContentSnifferClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ BufferSize() uint }); ok {
		pclass.get_buffer_size = (*[0]byte)(C._gotk4_soup2_ContentSnifferClass_get_buffer_size)
	}

	if _, ok := goval.(interface {
		Sniff(msg *Message, buffer *Buffer) (map[string]string, string)
	}); ok {
		pclass.sniff = (*[0]byte)(C._gotk4_soup2_ContentSnifferClass_sniff)
	}
}

//export _gotk4_soup2_ContentSnifferClass_get_buffer_size
func _gotk4_soup2_ContentSnifferClass_get_buffer_size(arg0 *C.SoupContentSniffer) (cret C.gsize) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ BufferSize() uint })

	gsize := iface.BufferSize()

	cret = C.gsize(gsize)

	return cret
}

//export _gotk4_soup2_ContentSnifferClass_sniff
func _gotk4_soup2_ContentSnifferClass_sniff(arg0 *C.SoupContentSniffer, arg1 *C.SoupMessage, arg2 *C.SoupBuffer, arg3 **C.GHashTable) (cret *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Sniff(msg *Message, buffer *Buffer) (map[string]string, string)
	})

	var _msg *Message   // out
	var _buffer *Buffer // out

	_msg = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))
	_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	params, utf8 := iface.Sniff(_msg, _buffer)

	if params != nil {
		*arg3 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
		for ksrc, vsrc := range params {
			var kdst *C.gchar // out
			var vdst *C.gchar // out
			kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
			defer C.free(unsafe.Pointer(kdst))
			vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
			defer C.free(unsafe.Pointer(vdst))
			C.g_hash_table_insert(*arg3, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
		}
	}
	cret = (*C.char)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

func wrapContentSniffer(obj *externglib.Object) *ContentSniffer {
	return &ContentSniffer{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalContentSniffer(p uintptr) (interface{}, error) {
	return wrapContentSniffer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewContentSniffer creates a new ContentSniffer.
//
// The function returns the following values:
//
//    - contentSniffer: new ContentSniffer.
//
func NewContentSniffer() *ContentSniffer {
	var _cret *C.SoupContentSniffer // in

	_cret = C.soup_content_sniffer_new()

	var _contentSniffer *ContentSniffer // out

	_contentSniffer = wrapContentSniffer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _contentSniffer
}

// BufferSize gets the number of bytes sniffer needs in order to properly sniff
// a buffer.
//
// The function returns the following values:
//
//    - gsize: number of bytes to sniff.
//
func (sniffer *ContentSniffer) BufferSize() uint {
	var _arg0 *C.SoupContentSniffer // out
	var _cret C.gsize               // in

	_arg0 = (*C.SoupContentSniffer)(unsafe.Pointer(externglib.InternObject(sniffer).Native()))

	_cret = C.soup_content_sniffer_get_buffer_size(_arg0)
	runtime.KeepAlive(sniffer)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Sniff sniffs buffer to determine its Content-Type. The result may also be
// influenced by the Content-Type declared in msg's response headers.
//
// The function takes the following parameters:
//
//    - msg: message to sniff.
//    - buffer containing the start of msg's response body.
//
// The function returns the following values:
//
//    - params (optional): return location for Content-Type parameters (eg,
//      "charset"), or NULL.
//    - utf8: sniffed Content-Type of buffer; this will never be NULL, but may be
//      "application/octet-stream".
//
func (sniffer *ContentSniffer) Sniff(msg *Message, buffer *Buffer) (map[string]string, string) {
	var _arg0 *C.SoupContentSniffer // out
	var _arg1 *C.SoupMessage        // out
	var _arg2 *C.SoupBuffer         // out
	var _arg3 *C.GHashTable         // in
	var _cret *C.char               // in

	_arg0 = (*C.SoupContentSniffer)(unsafe.Pointer(externglib.InternObject(sniffer).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(externglib.InternObject(msg).Native()))
	_arg2 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	_cret = C.soup_content_sniffer_sniff(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(sniffer)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(buffer)

	var _params map[string]string // out
	var _utf8 string              // out

	if _arg3 != nil {
		_params = make(map[string]string, gextras.HashTableSize(unsafe.Pointer(_arg3)))
		gextras.MoveHashTable(unsafe.Pointer(_arg3), true, func(k, v unsafe.Pointer) {
			ksrc := *(**C.gchar)(k)
			vsrc := *(**C.gchar)(v)
			var kdst string // out
			var vdst string // out
			kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
			vdst = C.GoString((*C.gchar)(unsafe.Pointer(vsrc)))
			_params[kdst] = vdst
		})
	}
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _params, _utf8
}
