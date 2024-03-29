// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// GType values.
var (
	GTypeContentSniffer = coreglib.Type(C.soup_content_sniffer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContentSniffer, F: marshalContentSniffer},
	})
}

// ContentSnifferOverrides contains methods that are overridable.
type ContentSnifferOverrides struct {
}

func defaultContentSnifferOverrides(v *ContentSniffer) ContentSnifferOverrides {
	return ContentSnifferOverrides{}
}

// ContentSniffer sniffs the mime type of messages.
//
// A ContentSniffer tries to detect the actual content type of the files that
// are being downloaded by looking at some of the data before the message emits
// its message::got-headers signal. ContentSniffer implements sessionfeature,
// so you can add content sniffing to a session with session.AddFeature or
// session.AddFeatureByType.
type ContentSniffer struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SessionFeature
}

var (
	_ coreglib.Objector = (*ContentSniffer)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ContentSniffer, *ContentSnifferClass, ContentSnifferOverrides](
		GTypeContentSniffer,
		initContentSnifferClass,
		wrapContentSniffer,
		defaultContentSnifferOverrides,
	)
}

func initContentSnifferClass(gclass unsafe.Pointer, overrides ContentSnifferOverrides, classInitFunc func(*ContentSnifferClass)) {
	if classInitFunc != nil {
		class := (*ContentSnifferClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapContentSniffer(obj *coreglib.Object) *ContentSniffer {
	return &ContentSniffer{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalContentSniffer(p uintptr) (interface{}, error) {
	return wrapContentSniffer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewContentSniffer creates a new ContentSniffer.
//
// The function returns the following values:
//
//   - contentSniffer: new ContentSniffer.
//
func NewContentSniffer() *ContentSniffer {
	var _cret *C.SoupContentSniffer // in

	_cret = C.soup_content_sniffer_new()

	var _contentSniffer *ContentSniffer // out

	_contentSniffer = wrapContentSniffer(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _contentSniffer
}

// Sniff sniffs buffer to determine its Content-Type.
//
// The result may also be influenced by the Content-Type declared in msg's
// response headers.
//
// The function takes the following parameters:
//
//   - msg: message to sniff.
//   - buffer containing the start of msg's response body.
//
// The function returns the following values:
//
//   - params (optional): return location for Content-Type parameters (eg,
//     "charset"), or NULL.
//   - utf8: sniffed Content-Type of buffer; this will never be NULL, but may be
//     application/octet-stream.
//
func (sniffer *ContentSniffer) Sniff(msg *Message, buffer *glib.Bytes) (map[string]string, string) {
	var _arg0 *C.SoupContentSniffer // out
	var _arg1 *C.SoupMessage        // out
	var _arg2 *C.GBytes             // out
	var _arg3 *C.GHashTable         // in
	var _cret *C.char               // in

	_arg0 = (*C.SoupContentSniffer)(unsafe.Pointer(coreglib.InternObject(sniffer).Native()))
	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))
	_arg2 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(buffer)))

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

// ContentSnifferClass: instance of this type is always passed by reference.
type ContentSnifferClass struct {
	*contentSnifferClass
}

// contentSnifferClass is the struct that's finalized.
type contentSnifferClass struct {
	native *C.SoupContentSnifferClass
}
