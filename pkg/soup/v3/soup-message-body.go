// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
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
	GTypeMemoryUse   = coreglib.Type(C.soup_memory_use_get_type())
	GTypeMessageBody = coreglib.Type(C.soup_message_body_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMemoryUse, F: marshalMemoryUse},
		coreglib.TypeMarshaler{T: GTypeMessageBody, F: marshalMessageBody},
	})
}

// MemoryUse: lifetime of the memory being passed.
type MemoryUse C.gint

const (
	// MemoryStatic: memory is statically allocated and constant; libsoup can
	// use the passed-in buffer directly and not need to worry about it being
	// modified or freed.
	MemoryStatic MemoryUse = iota
	// MemoryTake: caller has allocated the memory and libsoup will assume
	// ownership of it and free it with glib.Free().
	MemoryTake
	// MemoryCopy: passed-in data belongs to the caller and libsoup will copy it
	// into new memory leaving the caller free to reuse the original memory.
	MemoryCopy
)

func marshalMemoryUse(p uintptr) (interface{}, error) {
	return MemoryUse(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for MemoryUse.
func (m MemoryUse) String() string {
	switch m {
	case MemoryStatic:
		return "Static"
	case MemoryTake:
		return "Take"
	case MemoryCopy:
		return "Copy"
	default:
		return fmt.Sprintf("MemoryUse(%d)", m)
	}
}

// MessageBody represents the request or response body of a message.
//
// Note that while length always reflects the full length of the message body,
// data is normally NULL, and will only be filled in after messagebody.Flatten
// is called. For client-side messages, this automatically happens for the
// response body after it has been fully read. Likewise, for server-side
// messages, the request body is automatically filled in after being read.
//
// As an added bonus, when data is filled in, it is always terminated with a \0
// byte (which is not reflected in length).
//
// An instance of this type is always passed by reference.
type MessageBody struct {
	*messageBody
}

// messageBody is the struct that's finalized.
type messageBody struct {
	native *C.SoupMessageBody
}

func marshalMessageBody(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &MessageBody{&messageBody{(*C.SoupMessageBody)(b)}}, nil
}

// NewMessageBody constructs a struct MessageBody.
func NewMessageBody() *MessageBody {
	var _cret *C.SoupMessageBody // in

	_cret = C.soup_message_body_new()

	var _messageBody *MessageBody // out

	_messageBody = (*MessageBody)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_messageBody)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_message_body_unref((*C.SoupMessageBody)(intern.C))
		},
	)

	return _messageBody
}

// AppendBytes appends the data from buffer to body.
//
// The function takes the following parameters:
//
//   - buffer: #GBytes.
//
func (body *MessageBody) AppendBytes(buffer *glib.Bytes) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 *C.GBytes          // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(buffer)))

	C.soup_message_body_append_bytes(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(buffer)
}

// Append appends length bytes from data to body.
//
// This function is exactly equivalent to messagebody.Append with
// SOUP_MEMORY_TAKE as second argument; it exists mainly for convenience and
// simplifying language bindings.
//
// The function takes the following parameters:
//
//   - data to append.
//
func (body *MessageBody) Append(data []byte) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 *C.guchar          // out
	var _arg2 C.gsize

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg2 = (C.gsize)(len(data))
	_arg1 = (*C.guchar)(C.calloc(C.size_t(len(data)), C.size_t(C.sizeof_guchar)))
	copy(unsafe.Slice((*byte)(_arg1), len(data)), data)

	C.soup_message_body_append_take(_arg0, _arg1, _arg2)
	runtime.KeepAlive(body)
	runtime.KeepAlive(data)
}

// Complete tags body as being complete.
//
// Call this when using chunked encoding after you have appended the last chunk.
func (body *MessageBody) Complete() {
	var _arg0 *C.SoupMessageBody // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))

	C.soup_message_body_complete(_arg0)
	runtime.KeepAlive(body)
}

// Flatten fills in body's data field with a buffer containing all of the data
// in body.
//
// Adds an additional \0 byte not counted by body's length field.
//
// The function returns the following values:
//
//   - bytes containing the same data as body. (You must glib.Bytes.Unref() this
//     if you do not want it.).
//
func (body *MessageBody) Flatten() *glib.Bytes {
	var _arg0 *C.SoupMessageBody // out
	var _cret *C.GBytes          // in

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))

	_cret = C.soup_message_body_flatten(_arg0)
	runtime.KeepAlive(body)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}

// Accumulate gets the accumulate flag on body.
//
// See [methodMessageBody.set_accumulate. for details.
//
// The function returns the following values:
//
//   - ok: accumulate flag for body.
//
func (body *MessageBody) Accumulate() bool {
	var _arg0 *C.SoupMessageBody // out
	var _cret C.gboolean         // in

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))

	_cret = C.soup_message_body_get_accumulate(_arg0)
	runtime.KeepAlive(body)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Chunk gets a glib.Bytes containing data from body starting at offset.
//
// The size of the returned chunk is unspecified. You can iterate through the
// entire body by first calling messagebody.GetChunk with an offset of 0,
// and then on each successive call, increment the offset by the length of the
// previously-returned chunk.
//
// If offset is greater than or equal to the total length of body, then the
// return value depends on whether or not messagebody.Complete has been called
// or not; if it has, then messagebody.GetChunk will return a 0-length chunk
// (indicating the end of body). If it has not, then messagebody.GetChunk will
// return NULL (indicating that body may still potentially have more data,
// but that data is not currently available).
//
// The function takes the following parameters:
//
//   - offset: offset.
//
// The function returns the following values:
//
//   - bytes (optional): #GBytes.
//
func (body *MessageBody) Chunk(offset int64) *glib.Bytes {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 C.goffset          // out
	var _cret *C.GBytes          // in

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = C.goffset(offset)

	_cret = C.soup_message_body_get_chunk(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(offset)

	var _bytes *glib.Bytes // out

	if _cret != nil {
		_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_bytes)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_bytes_unref((*C.GBytes)(intern.C))
			},
		)
	}

	return _bytes
}

// GotChunk handles the MessageBody part of receiving a chunk of data from the
// network.
//
// Normally this means appending chunk to body, exactly as with
// messagebody.AppendBytes, but if you have set body's accumulate flag to FALSE,
// then that will not happen.
//
// This is a low-level method which you should not normally need to use.
//
// The function takes the following parameters:
//
//   - chunk received from the network.
//
func (body *MessageBody) GotChunk(chunk *glib.Bytes) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 *C.GBytes          // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(chunk)))

	C.soup_message_body_got_chunk(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(chunk)
}

// SetAccumulate sets or clears the accumulate flag on body.
//
// (The default value is TRUE.) If set to FALSE, body's data field will not be
// filled in after the body is fully sent/received, and the chunks that make up
// body may be discarded when they are no longer needed.
//
// If you set the flag to FALSE on the message request_body of a client-side
// message, it will block the accumulation of chunks into body's data field,
// but it will not normally cause the chunks to be discarded after being written
// like in the server-side message response_body case, because the request body
// needs to be kept around in case the request needs to be sent a second time
// due to redirection or authentication.
//
// The function takes the following parameters:
//
//   - accumulate: whether or not to accumulate body chunks in body.
//
func (body *MessageBody) SetAccumulate(accumulate bool) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	if accumulate {
		_arg1 = C.TRUE
	}

	C.soup_message_body_set_accumulate(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(accumulate)
}

// Truncate deletes all of the data in body.
func (body *MessageBody) Truncate() {
	var _arg0 *C.SoupMessageBody // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))

	C.soup_message_body_truncate(_arg0)
	runtime.KeepAlive(body)
}

// WroteChunk handles the MessageBody part of writing a chunk of data to the
// network.
//
// Normally this is a no-op, but if you have set body's accumulate flag to
// FALSE, then this will cause chunk to be discarded to free up memory.
//
// This is a low-level method which you should not need to use, and there are
// further restrictions on its proper use which are not documented here.
//
// The function takes the following parameters:
//
//   - chunk returned from messagebody.GetChunk.
//
func (body *MessageBody) WroteChunk(chunk *glib.Bytes) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 *C.GBytes          // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(chunk)))

	C.soup_message_body_wrote_chunk(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(chunk)
}
