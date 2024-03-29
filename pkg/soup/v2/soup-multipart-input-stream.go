// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// GType values.
var (
	GTypeMultipartInputStream = coreglib.Type(C.soup_multipart_input_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMultipartInputStream, F: marshalMultipartInputStream},
	})
}

// MultipartInputStreamOverrides contains methods that are overridable.
type MultipartInputStreamOverrides struct {
}

func defaultMultipartInputStreamOverrides(v *MultipartInputStream) MultipartInputStreamOverrides {
	return MultipartInputStreamOverrides{}
}

type MultipartInputStream struct {
	_ [0]func() // equal guard
	gio.FilterInputStream

	*coreglib.Object
	gio.InputStream
	gio.PollableInputStream
}

var (
	_ gio.FilterInputStreamer = (*MultipartInputStream)(nil)
	_ coreglib.Objector       = (*MultipartInputStream)(nil)
	_ gio.InputStreamer       = (*MultipartInputStream)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*MultipartInputStream, *MultipartInputStreamClass, MultipartInputStreamOverrides](
		GTypeMultipartInputStream,
		initMultipartInputStreamClass,
		wrapMultipartInputStream,
		defaultMultipartInputStreamOverrides,
	)
}

func initMultipartInputStreamClass(gclass unsafe.Pointer, overrides MultipartInputStreamOverrides, classInitFunc func(*MultipartInputStreamClass)) {
	if classInitFunc != nil {
		class := (*MultipartInputStreamClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMultipartInputStream(obj *coreglib.Object) *MultipartInputStream {
	return &MultipartInputStream{
		FilterInputStream: gio.FilterInputStream{
			InputStream: gio.InputStream{
				Object: obj,
			},
		},
		Object: obj,
		InputStream: gio.InputStream{
			Object: obj,
		},
		PollableInputStream: gio.PollableInputStream{
			InputStream: gio.InputStream{
				Object: obj,
			},
		},
	}
}

func marshalMultipartInputStream(p uintptr) (interface{}, error) {
	return wrapMultipartInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMultipartInputStream creates a new MultipartInputStream that wraps
// the Stream obtained by sending the Request. Reads should not be done
// directly through this object, use the input streams returned by
// soup_multipart_input_stream_next_part() or its async counterpart instead.
//
// The function takes the following parameters:
//
//   - msg the response is related to.
//   - baseStream returned by sending the request.
//
// The function returns the following values:
//
//   - multipartInputStream: new MultipartInputStream.
//
func NewMultipartInputStream(msg *Message, baseStream gio.InputStreamer) *MultipartInputStream {
	var _arg1 *C.SoupMessage              // out
	var _arg2 *C.GInputStream             // out
	var _cret *C.SoupMultipartInputStream // in

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))
	_arg2 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))

	_cret = C.soup_multipart_input_stream_new(_arg1, _arg2)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(baseStream)

	var _multipartInputStream *MultipartInputStream // out

	_multipartInputStream = wrapMultipartInputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _multipartInputStream
}

// Headers obtains the headers for the part currently being processed.
// Note that the MessageHeaders that are returned are owned by the
// MultipartInputStream and will be replaced when a call is made to
// soup_multipart_input_stream_next_part() or its async counterpart, so if
// keeping the headers is required, a copy must be made.
//
// Note that if a part had no headers at all an empty MessageHeaders will be
// returned.
//
// The function returns the following values:
//
//   - messageHeaders (optional) the headers for the part currently being
//     processed or NULL if the headers failed to parse.
//
func (multipart *MultipartInputStream) Headers() *MessageHeaders {
	var _arg0 *C.SoupMultipartInputStream // out
	var _cret *C.SoupMessageHeaders       // in

	_arg0 = (*C.SoupMultipartInputStream)(unsafe.Pointer(coreglib.InternObject(multipart).Native()))

	_cret = C.soup_multipart_input_stream_get_headers(_arg0)
	runtime.KeepAlive(multipart)

	var _messageHeaders *MessageHeaders // out

	if _cret != nil {
		_messageHeaders = (*MessageHeaders)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _messageHeaders
}

// NextPart obtains an input stream for the next part. When dealing
// with a multipart response the input stream needs to be wrapped in a
// MultipartInputStream and this function or its async counterpart need to be
// called to obtain the first part for reading.
//
// After calling this function, soup_multipart_input_stream_get_headers() can be
// used to obtain the headers for the first part. A read of 0 bytes indicates
// the end of the part; a new call to this function should be done at that
// point, to obtain the next part.
//
// The function takes the following parameters:
//
//   - ctx (optional): #GCancellable.
//
// The function returns the following values:
//
//   - inputStream (optional): new Stream, or NULL if there are no more parts.
//
func (multipart *MultipartInputStream) NextPart(ctx context.Context) (gio.InputStreamer, error) {
	var _arg0 *C.SoupMultipartInputStream // out
	var _arg1 *C.GCancellable             // out
	var _cret *C.GInputStream             // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.SoupMultipartInputStream)(unsafe.Pointer(coreglib.InternObject(multipart).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.soup_multipart_input_stream_next_part(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(ctx)

	var _inputStream gio.InputStreamer // out
	var _goerr error                   // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.InputStreamer)
				return ok
			})
			rv, ok := casted.(gio.InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _inputStream, _goerr
}

// NextPartFinish finishes an asynchronous request for the next part.
//
// The function takes the following parameters:
//
//   - result: Result.
//
// The function returns the following values:
//
//   - inputStream (optional): newly created Stream for reading the next part or
//     NULL if there are no more parts.
//
func (multipart *MultipartInputStream) NextPartFinish(result gio.AsyncResulter) (gio.InputStreamer, error) {
	var _arg0 *C.SoupMultipartInputStream // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GInputStream             // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.SoupMultipartInputStream)(unsafe.Pointer(coreglib.InternObject(multipart).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.soup_multipart_input_stream_next_part_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(multipart)
	runtime.KeepAlive(result)

	var _inputStream gio.InputStreamer // out
	var _goerr error                   // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.InputStreamer)
				return ok
			})
			rv, ok := casted.(gio.InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _inputStream, _goerr
}

// MultipartInputStreamClass: instance of this type is always passed by
// reference.
type MultipartInputStreamClass struct {
	*multipartInputStreamClass
}

// multipartInputStreamClass is the struct that's finalized.
type multipartInputStreamClass struct {
	native *C.SoupMultipartInputStreamClass
}

func (m *MultipartInputStreamClass) ParentClass() *gio.FilterInputStreamClass {
	valptr := &m.native.parent_class
	var _v *gio.FilterInputStreamClass // out
	_v = (*gio.FilterInputStreamClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
