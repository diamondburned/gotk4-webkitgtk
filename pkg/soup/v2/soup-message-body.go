// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_memory_use_get_type()), F: marshalMemoryUse},
		{T: externglib.Type(C.soup_buffer_get_type()), F: marshalBuffer},
		{T: externglib.Type(C.soup_message_body_get_type()), F: marshalMessageBody},
	})
}

// MemoryUse describes how Buffer should use the data passed in by the caller.
//
// See also soup_buffer_new_with_owner(), which allows to you create a buffer
// containing data which is owned by another object.
type MemoryUse C.gint

const (
	// MemoryStatic: memory is statically allocated and constant; libsoup can
	// use the passed-in buffer directly and not need to worry about it being
	// modified or freed.
	MemoryStatic MemoryUse = iota
	// MemoryTake: caller has allocated the memory for the Buffer's use; libsoup
	// will assume ownership of it and free it (with g_free()) when it is done
	// with it.
	MemoryTake
	// MemoryCopy: passed-in data belongs to the caller; the Buffer will copy it
	// into new memory, leaving the caller free to reuse the original memory.
	MemoryCopy
	// MemoryTemporary: passed-in data belongs to the caller, but will remain
	// valid for the lifetime of the Buffer. The difference between this and
	// SOUP_MEMORY_STATIC is that if you copy a SOUP_MEMORY_TEMPORARY buffer, it
	// will make a copy of the memory as well, rather than reusing the original
	// memory.
	MemoryTemporary
)

func marshalMemoryUse(p uintptr) (interface{}, error) {
	return MemoryUse(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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
	case MemoryTemporary:
		return "Temporary"
	default:
		return fmt.Sprintf("MemoryUse(%d)", m)
	}
}

// Buffer: data buffer, generally used to represent a chunk of a MessageBody.
//
// data is a #char because that's generally convenient; in some situations you
// may need to cast it to #guchar or another type.
//
// An instance of this type is always passed by reference.
type Buffer struct {
	*buffer
}

// buffer is the struct that's finalized.
type buffer struct {
	native *C.SoupBuffer
}

func marshalBuffer(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Buffer{&buffer{(*C.SoupBuffer)(b)}}, nil
}

// NewBuffer constructs a struct Buffer.
func NewBuffer(data []byte) *Buffer {
	var _arg1 *C.guchar // out
	var _arg2 C.gsize
	var _cret *C.SoupBuffer // in

	_arg2 = (C.gsize)(len(data))
	_arg1 = (*C.guchar)(C.calloc(C.size_t(len(data)), C.size_t(C.sizeof_guchar)))
	copy(unsafe.Slice((*byte)(_arg1), len(data)), data)

	_cret = C.soup_buffer_new_take(_arg1, _arg2)
	runtime.KeepAlive(data)

	var _buffer *Buffer // out

	_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_buffer_free((*C.SoupBuffer)(intern.C))
		},
	)

	return _buffer
}

// Length: length of data.
func (b *Buffer) Length() uint {
	var v uint // out
	v = uint(b.native.length)
	return v
}

// Copy makes a copy of buffer. In reality, Buffer is a refcounted type, and
// calling soup_buffer_copy() will normally just increment the refcount on
// buffer and return it. However, if buffer was created with UP_MEMORY_TEMPORARY
// memory, then soup_buffer_copy() will actually return a copy of it, so that
// the data in the copy will remain valid after the temporary buffer is freed.
//
// The function returns the following values:
//
//    - ret: new (or newly-reffed) buffer.
//
func (buffer *Buffer) Copy() *Buffer {
	var _arg0 *C.SoupBuffer // out
	var _cret *C.SoupBuffer // in

	_arg0 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	_cret = C.soup_buffer_copy(_arg0)
	runtime.KeepAlive(buffer)

	var _ret *Buffer // out

	_ret = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_buffer_free((*C.SoupBuffer)(intern.C))
		},
	)

	return _ret
}

// AsBytes creates a #GBytes pointing to the same memory as buffer. The #GBytes
// will hold a reference on buffer to ensure that it is not freed while the
// #GBytes is still valid.
//
// The function returns the following values:
//
//    - bytes: new #GBytes which has the same content as the Buffer.
//
func (buffer *Buffer) AsBytes() *glib.Bytes {
	var _arg0 *C.SoupBuffer // out
	var _cret *C.GBytes     // in

	_arg0 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	_cret = C.soup_buffer_get_as_bytes(_arg0)
	runtime.KeepAlive(buffer)

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

// Data: this function exists for use by language bindings, because it's not
// currently possible to get the right effect by annotating the fields of
// Buffer.
//
// The function returns the following values:
//
//    - data: pointer to the buffer data is stored here.
//
func (buffer *Buffer) Data() []byte {
	var _arg0 *C.SoupBuffer // out
	var _arg1 *C.guint8     // in
	var _arg2 C.gsize       // in

	_arg0 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	C.soup_buffer_get_data(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(buffer)

	var _data []byte // out

	_data = make([]byte, _arg2)
	copy(_data, unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2))

	return _data
}

// Owner gets the "owner" object for a buffer created with
// soup_buffer_new_with_owner().
//
// The function returns the following values:
//
//    - gpointer (optional): owner pointer.
//
func (buffer *Buffer) Owner() cgo.Handle {
	var _arg0 *C.SoupBuffer // out
	var _cret C.gpointer    // in

	_arg0 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	_cret = C.soup_buffer_get_owner(_arg0)
	runtime.KeepAlive(buffer)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// NewSubbuffer creates a new Buffer containing length bytes "copied" from
// parent starting at offset. (Normally this will not actually copy any data,
// but will instead simply reference the same data as parent does.).
//
// The function takes the following parameters:
//
//    - offset within parent to start at.
//    - length: number of bytes to copy from parent.
//
// The function returns the following values:
//
//    - buffer: new Buffer.
//
func (parent *Buffer) NewSubbuffer(offset uint, length uint) *Buffer {
	var _arg0 *C.SoupBuffer // out
	var _arg1 C.gsize       // out
	var _arg2 C.gsize       // out
	var _cret *C.SoupBuffer // in

	_arg0 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(parent)))
	_arg1 = C.gsize(offset)
	_arg2 = C.gsize(length)

	_cret = C.soup_buffer_new_subbuffer(_arg0, _arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(length)

	var _buffer *Buffer // out

	_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_buffer_free((*C.SoupBuffer)(intern.C))
		},
	)

	return _buffer
}

// MessageBody request or response body.
//
// Note that while length always reflects the full length of the message body,
// data is normally NULL, and will only be filled in after
// soup_message_body_flatten() is called. For client-side messages, this
// automatically happens for the response body after it has been fully read,
// unless you set the SOUP_MESSAGE_OVERWRITE_CHUNKS flags. Likewise, for
// server-side messages, the request body is automatically filled in after being
// read.
//
// As an added bonus, when data is filled in, it is always terminated with a
// '\0' byte (which is not reflected in length).
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
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
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
			C.soup_message_body_free((*C.SoupMessageBody)(intern.C))
		},
	)

	return _messageBody
}

// Data: data.
func (m *MessageBody) Data() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(m.native.data)))
	return v
}

// Length: length of data.
func (m *MessageBody) Length() int64 {
	var v int64 // out
	v = int64(m.native.length)
	return v
}

// AppendBuffer appends the data from buffer to body. (MessageBody uses Buffers
// internally, so this is normally a constant-time operation that doesn't
// actually require copying the data in buffer.).
//
// The function takes the following parameters:
//
//    - buffer: Buffer.
//
func (body *MessageBody) AppendBuffer(buffer *Buffer) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 *C.SoupBuffer      // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	C.soup_message_body_append_buffer(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(buffer)
}

// Append appends length bytes from data to body.
//
// This function is exactly equivalent to soup_message_body_append() with
// SOUP_MEMORY_TAKE as second argument; it exists mainly for convenience and
// simplifying language bindings.
//
// The function takes the following parameters:
//
//    - data to append.
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

// Complete tags body as being complete; Call this when using chunked encoding
// after you have appended the last chunk.
func (body *MessageBody) Complete() {
	var _arg0 *C.SoupMessageBody // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))

	C.soup_message_body_complete(_arg0)
	runtime.KeepAlive(body)
}

// Flatten fills in body's data field with a buffer containing all of the data
// in body (plus an additional '\0' byte not counted by body's length field).
//
// The function returns the following values:
//
//    - buffer containing the same data as body. (You must free this buffer if
//      you do not want it.).
//
func (body *MessageBody) Flatten() *Buffer {
	var _arg0 *C.SoupMessageBody // out
	var _cret *C.SoupBuffer      // in

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))

	_cret = C.soup_message_body_flatten(_arg0)
	runtime.KeepAlive(body)

	var _buffer *Buffer // out

	_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_buffer_free((*C.SoupBuffer)(intern.C))
		},
	)

	return _buffer
}

// Accumulate gets the accumulate flag on body; see
// soup_message_body_set_accumulate() for details.
//
// The function returns the following values:
//
//    - ok: accumulate flag for body.
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

// Chunk gets a Buffer containing data from body starting at offset. The size of
// the returned chunk is unspecified. You can iterate through the entire body by
// first calling soup_message_body_get_chunk() with an offset of 0, and then on
// each successive call, increment the offset by the length of the
// previously-returned chunk.
//
// If offset is greater than or equal to the total length of body, then the
// return value depends on whether or not soup_message_body_complete() has been
// called or not; if it has, then soup_message_body_get_chunk() will return a
// 0-length chunk (indicating the end of body). If it has not, then
// soup_message_body_get_chunk() will return NULL (indicating that body may
// still potentially have more data, but that data is not currently available).
//
// The function takes the following parameters:
//
//    - offset: offset.
//
// The function returns the following values:
//
//    - buffer (optional) or NULL.
//
func (body *MessageBody) Chunk(offset int64) *Buffer {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 C.goffset          // out
	var _cret *C.SoupBuffer      // in

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = C.goffset(offset)

	_cret = C.soup_message_body_get_chunk(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(offset)

	var _buffer *Buffer // out

	if _cret != nil {
		_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_buffer)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_buffer_free((*C.SoupBuffer)(intern.C))
			},
		)
	}

	return _buffer
}

// GotChunk handles the MessageBody part of receiving a chunk of data from the
// network. Normally this means appending chunk to body, exactly as with
// soup_message_body_append_buffer(), but if you have set body's accumulate flag
// to FALSE, then that will not happen.
//
// This is a low-level method which you should not normally need to use.
//
// The function takes the following parameters:
//
//    - chunk received from the network.
//
func (body *MessageBody) GotChunk(chunk *Buffer) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 *C.SoupBuffer      // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(chunk)))

	C.soup_message_body_got_chunk(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(chunk)
}

// SetAccumulate sets or clears the accumulate flag on body. (The default value
// is TRUE.) If set to FALSE, body's data field will not be filled in after the
// body is fully sent/received, and the chunks that make up body may be
// discarded when they are no longer needed.
//
// In particular, if you set this flag to FALSE on an "incoming" message body
// (that is, the Message:response_body of a client-side message, or
// Message:request_body of a server-side message), this will cause each chunk of
// the body to be discarded after its corresponding Message::got_chunk signal is
// emitted. (This is equivalent to setting the deprecated
// SOUP_MESSAGE_OVERWRITE_CHUNKS flag on the message.)
//
// If you set this flag to FALSE on the Message:response_body of a server-side
// message, it will cause each chunk of the body to be discarded after its
// corresponding Message::wrote_chunk signal is emitted.
//
// If you set the flag to FALSE on the Message:request_body of a client-side
// message, it will block the accumulation of chunks into body's data field, but
// it will not normally cause the chunks to be discarded after being written
// like in the server-side Message:response_body case, because the request body
// needs to be kept around in case the request needs to be sent a second time
// due to redirection or authentication. However, if you set the
// SOUP_MESSAGE_CAN_REBUILD flag on the message, then the chunks will be
// discarded, and you will be responsible for recreating the request body after
// the Message::restarted signal is emitted.
//
// The function takes the following parameters:
//
//    - accumulate: whether or not to accumulate body chunks in body.
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
// network. Normally this is a no-op, but if you have set body's accumulate flag
// to FALSE, then this will cause chunk to be discarded to free up memory.
//
// This is a low-level method which you should not need to use, and there are
// further restrictions on its proper use which are not documented here.
//
// The function takes the following parameters:
//
//    - chunk returned from soup_message_body_get_chunk().
//
func (body *MessageBody) WroteChunk(chunk *Buffer) {
	var _arg0 *C.SoupMessageBody // out
	var _arg1 *C.SoupBuffer      // out

	_arg0 = (*C.SoupMessageBody)(gextras.StructNative(unsafe.Pointer(body)))
	_arg1 = (*C.SoupBuffer)(gextras.StructNative(unsafe.Pointer(chunk)))

	C.soup_message_body_wrote_chunk(_arg0, _arg1)
	runtime.KeepAlive(body)
	runtime.KeepAlive(chunk)
}
