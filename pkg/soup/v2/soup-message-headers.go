// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// void _gotk4_soup2_MessageHeadersForEachFunc(char*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_encoding_get_type()), F: marshalEncoding},
		{T: externglib.Type(C.soup_message_headers_type_get_type()), F: marshalMessageHeadersType},
		{T: externglib.Type(C.soup_expectation_get_type()), F: marshalExpectation},
		{T: externglib.Type(C.soup_message_headers_get_type()), F: marshalMessageHeaders},
	})
}

// Encoding: how a message body is encoded for transport.
type Encoding C.gint

const (
	// EncodingUnrecognized: unknown / error.
	EncodingUnrecognized Encoding = iota
	// EncodingNone: no body is present (which is not the same as a 0-length
	// body, and only occurs in certain places).
	EncodingNone
	// EncodingContentLength: content-Length encoding.
	EncodingContentLength
	// EncodingEOF: response body ends when the connection is closed.
	EncodingEOF
	// EncodingChunked: chunked encoding (currently only supported for
	// response).
	EncodingChunked
	// EncodingByteranges multipart/byteranges (Reserved for future use: NOT
	// CURRENTLY IMPLEMENTED).
	EncodingByteranges
)

func marshalEncoding(p uintptr) (interface{}, error) {
	return Encoding(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Encoding.
func (e Encoding) String() string {
	switch e {
	case EncodingUnrecognized:
		return "Unrecognized"
	case EncodingNone:
		return "None"
	case EncodingContentLength:
		return "ContentLength"
	case EncodingEOF:
		return "EOF"
	case EncodingChunked:
		return "Chunked"
	case EncodingByteranges:
		return "Byteranges"
	default:
		return fmt.Sprintf("Encoding(%d)", e)
	}
}

// MessageHeadersType: value passed to soup_message_headers_new() to set certain
// default behaviors.
type MessageHeadersType C.gint

const (
	// MessageHeadersRequest: request headers.
	MessageHeadersRequest MessageHeadersType = iota
	// MessageHeadersResponse: response headers.
	MessageHeadersResponse
	// MessageHeadersMultipart: multipart body part headers.
	MessageHeadersMultipart
)

func marshalMessageHeadersType(p uintptr) (interface{}, error) {
	return MessageHeadersType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for MessageHeadersType.
func (m MessageHeadersType) String() string {
	switch m {
	case MessageHeadersRequest:
		return "Request"
	case MessageHeadersResponse:
		return "Response"
	case MessageHeadersMultipart:
		return "Multipart"
	default:
		return fmt.Sprintf("MessageHeadersType(%d)", m)
	}
}

// Expectation represents the parsed value of the "Expect" header.
type Expectation C.guint

const (
	// ExpectationUnrecognized: any unrecognized expectation.
	ExpectationUnrecognized Expectation = 0b1
	// ExpectationContinue: "100-continue".
	ExpectationContinue Expectation = 0b10
)

func marshalExpectation(p uintptr) (interface{}, error) {
	return Expectation(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for Expectation.
func (e Expectation) String() string {
	if e == 0 {
		return "Expectation(0)"
	}

	var builder strings.Builder
	builder.Grow(43)

	for e != 0 {
		next := e & (e - 1)
		bit := e - next

		switch bit {
		case ExpectationUnrecognized:
			builder.WriteString("Unrecognized|")
		case ExpectationContinue:
			builder.WriteString("Continue|")
		default:
			builder.WriteString(fmt.Sprintf("Expectation(0b%b)|", bit))
		}

		e = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if e contains other.
func (e Expectation) Has(other Expectation) bool {
	return (e & other) == other
}

// MessageHeadersForEachFunc: callback passed to soup_message_headers_foreach().
type MessageHeadersForEachFunc func(name, value string)

//export _gotk4_soup2_MessageHeadersForEachFunc
func _gotk4_soup2_MessageHeadersForEachFunc(arg0 *C.char, arg1 *C.char, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var name string  // out
	var value string // out

	name = C.GoString((*C.gchar)(unsafe.Pointer(arg0)))
	value = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	fn := v.(MessageHeadersForEachFunc)
	fn(name, value)
}

// MessageHeaders: HTTP message headers associated with a request or response.
//
// An instance of this type is always passed by reference.
type MessageHeaders struct {
	*messageHeaders
}

// messageHeaders is the struct that's finalized.
type messageHeaders struct {
	native *C.SoupMessageHeaders
}

func marshalMessageHeaders(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &MessageHeaders{&messageHeaders{(*C.SoupMessageHeaders)(b)}}, nil
}

// NewMessageHeaders constructs a struct MessageHeaders.
func NewMessageHeaders(typ MessageHeadersType) *MessageHeaders {
	var _arg1 C.SoupMessageHeadersType // out
	var _cret *C.SoupMessageHeaders    // in

	_arg1 = C.SoupMessageHeadersType(typ)

	_cret = C.soup_message_headers_new(_arg1)
	runtime.KeepAlive(typ)

	var _messageHeaders *MessageHeaders // out

	_messageHeaders = (*MessageHeaders)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_messageHeaders)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_message_headers_free((*C.SoupMessageHeaders)(intern.C))
		},
	)

	return _messageHeaders
}

// Append appends a new header with name name and value value to hdrs. (If there
// is an existing header with name name, then this creates a second one, which
// is only allowed for list-valued headers; see also
// soup_message_headers_replace().)
//
// The caller is expected to make sure that name and value are syntactically
// correct.
//
// The function takes the following parameters:
//
//    - name: header name to add.
//    - value: new value of name.
//
func (hdrs *MessageHeaders) Append(name string, value string) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _arg2 *C.char               // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))

	C.soup_message_headers_append(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// CleanConnectionHeaders removes all the headers listed in the Connection
// header.
func (hdrs *MessageHeaders) CleanConnectionHeaders() {
	var _arg0 *C.SoupMessageHeaders // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	C.soup_message_headers_clean_connection_headers(_arg0)
	runtime.KeepAlive(hdrs)
}

// Clear clears hdrs.
func (hdrs *MessageHeaders) Clear() {
	var _arg0 *C.SoupMessageHeaders // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	C.soup_message_headers_clear(_arg0)
	runtime.KeepAlive(hdrs)
}

// ForEach calls func once for each header value in hdrs.
//
// Beware that unlike soup_message_headers_get(), this processes the headers in
// exactly the way they were added, rather than concatenating multiple
// same-named headers into a single value. (This is intentional; it ensures that
// if you call soup_message_headers_append() multiple times with the same name,
// then the I/O code will output multiple copies of the header when sending the
// message to the remote implementation, which may be required for
// interoperability in some cases.)
//
// You may not modify the headers from func.
//
// The function takes the following parameters:
//
//    - fn: callback function to run for each header.
//
func (hdrs *MessageHeaders) ForEach(fn MessageHeadersForEachFunc) {
	var _arg0 *C.SoupMessageHeaders           // out
	var _arg1 C.SoupMessageHeadersForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*[0]byte)(C._gotk4_soup2_MessageHeadersForEachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.soup_message_headers_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(fn)
}

// FreeRanges frees the array of ranges returned from
// soup_message_headers_get_ranges().
//
// The function takes the following parameters:
//
//    - ranges: array of Range.
//
func (hdrs *MessageHeaders) FreeRanges(ranges *Range) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.SoupRange          // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.SoupRange)(gextras.StructNative(unsafe.Pointer(ranges)))

	C.soup_message_headers_free_ranges(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(ranges)
}

// Get gets the value of header name in hdrs.
//
// This method was supposed to work correctly for both single-valued and
// list-valued headers, but because some HTTP clients/servers mistakenly send
// multiple copies of headers that are supposed to be single-valued, it
// sometimes returns incorrect results. To fix this, the methods
// soup_message_headers_get_one() and soup_message_headers_get_list() were
// introduced, so callers can explicitly state which behavior they are
// expecting.
//
// Deprecated: Use soup_message_headers_get_one() or
// soup_message_headers_get_list() instead.
//
// The function takes the following parameters:
//
//    - name: header name.
//
// The function returns the following values:
//
//    - utf8 (optional) as with soup_message_headers_get_list().
//
func (hdrs *MessageHeaders) Get(name string) string {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_message_headers_get(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ContentDisposition looks up the "Content-Disposition" header in hdrs, parses
// it, and returns its value in *disposition and *params. params can be NULL if
// you are only interested in the disposition-type.
//
// In HTTP, the most common use of this header is to set a disposition-type of
// "attachment", to suggest to the browser that a response should be saved to
// disk rather than displayed in the browser. If params contains a "filename"
// parameter, this is a suggestion of a filename to use. (If the parameter value
// in the header contains an absolute or relative path, libsoup will truncate it
// down to just the final path component, so you do not need to test this
// yourself.)
//
// Content-Disposition is also used in "multipart/form-data", however this is
// handled automatically by Multipart and the associated form methods.
//
// The function returns the following values:
//
//    - disposition: return location for the disposition-type, or NULL.
//    - params: return location for the Content-Disposition parameters, or NULL.
//    - ok: TRUE if hdrs contains a "Content-Disposition" header, FALSE if not
//      (in which case *disposition and *params will be unchanged).
//
func (hdrs *MessageHeaders) ContentDisposition() (string, map[string]string, bool) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // in
	var _arg2 *C.GHashTable         // in
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	_cret = C.soup_message_headers_get_content_disposition(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(hdrs)

	var _disposition string       // out
	var _params map[string]string // out
	var _ok bool                  // out

	_disposition = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	defer C.free(unsafe.Pointer(_arg1))
	_params = make(map[string]string, gextras.HashTableSize(unsafe.Pointer(_arg2)))
	gextras.MoveHashTable(unsafe.Pointer(_arg2), true, func(k, v unsafe.Pointer) {
		ksrc := *(**C.gchar)(k)
		vsrc := *(**C.gchar)(v)
		var kdst string // out
		var vdst string // out
		kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
		vdst = C.GoString((*C.gchar)(unsafe.Pointer(vsrc)))
		_params[kdst] = vdst
	})
	if _cret != 0 {
		_ok = true
	}

	return _disposition, _params, _ok
}

// ContentLength gets the message body length that hdrs declare. This will only
// be non-0 if soup_message_headers_get_encoding() returns
// SOUP_ENCODING_CONTENT_LENGTH.
//
// The function returns the following values:
//
//    - gint64: message body length declared by hdrs.
//
func (hdrs *MessageHeaders) ContentLength() int64 {
	var _arg0 *C.SoupMessageHeaders // out
	var _cret C.goffset             // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	_cret = C.soup_message_headers_get_content_length(_arg0)
	runtime.KeepAlive(hdrs)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// ContentRange parses hdrs's Content-Range header and returns it in start, end,
// and total_length. If the total length field in the header was specified as
// "*", then total_length will be set to -1.
//
// The function returns the following values:
//
//    - start: return value for the start of the range.
//    - end: return value for the end of the range.
//    - totalLength (optional): return value for the total length of the
//      resource, or NULL if you don't care.
//    - ok: TRUE if hdrs contained a "Content-Range" header containing a byte
//      range which could be parsed, FALSE otherwise.
//
func (hdrs *MessageHeaders) ContentRange() (start int64, end int64, totalLength int64, ok bool) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 C.goffset             // in
	var _arg2 C.goffset             // in
	var _arg3 C.goffset             // in
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	_cret = C.soup_message_headers_get_content_range(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(hdrs)

	var _start int64       // out
	var _end int64         // out
	var _totalLength int64 // out
	var _ok bool           // out

	_start = int64(_arg1)
	_end = int64(_arg2)
	_totalLength = int64(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _start, _end, _totalLength, _ok
}

// ContentType looks up the "Content-Type" header in hdrs, parses it, and
// returns its value in *content_type and *params. params can be NULL if you are
// only interested in the content type itself.
//
// The function returns the following values:
//
//    - params (optional): return location for the Content-Type parameters (eg,
//      "charset"), or NULL.
//    - utf8 (optional): string with the value of the "Content-Type" header or
//      NULL if hdrs does not contain that header or it cannot be parsed (in
//      which case *params will be unchanged).
//
func (hdrs *MessageHeaders) ContentType() (map[string]string, string) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.GHashTable         // in
	var _cret *C.char               // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	_cret = C.soup_message_headers_get_content_type(_arg0, &_arg1)
	runtime.KeepAlive(hdrs)

	var _params map[string]string // out
	var _utf8 string              // out

	if _arg1 != nil {
		_params = make(map[string]string, gextras.HashTableSize(unsafe.Pointer(_arg1)))
		gextras.MoveHashTable(unsafe.Pointer(_arg1), true, func(k, v unsafe.Pointer) {
			ksrc := *(**C.gchar)(k)
			vsrc := *(**C.gchar)(v)
			var kdst string // out
			var vdst string // out
			kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
			vdst = C.GoString((*C.gchar)(unsafe.Pointer(vsrc)))
			_params[kdst] = vdst
		})
	}
	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _params, _utf8
}

// Encoding gets the message body encoding that hdrs declare. This may not
// always correspond to the encoding used on the wire; eg, a HEAD response may
// declare a Content-Length or Transfer-Encoding, but it will never actually
// include a body.
//
// The function returns the following values:
//
//    - encoding declared by hdrs.
//
func (hdrs *MessageHeaders) Encoding() Encoding {
	var _arg0 *C.SoupMessageHeaders // out
	var _cret C.SoupEncoding        // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	_cret = C.soup_message_headers_get_encoding(_arg0)
	runtime.KeepAlive(hdrs)

	var _encoding Encoding // out

	_encoding = Encoding(_cret)

	return _encoding
}

// Expectations gets the expectations declared by hdrs's "Expect" header.
// Currently this will either be SOUP_EXPECTATION_CONTINUE or
// SOUP_EXPECTATION_UNRECOGNIZED.
//
// The function returns the following values:
//
//    - expectation contents of hdrs's "Expect" header.
//
func (hdrs *MessageHeaders) Expectations() Expectation {
	var _arg0 *C.SoupMessageHeaders // out
	var _cret C.SoupExpectation     // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	_cret = C.soup_message_headers_get_expectations(_arg0)
	runtime.KeepAlive(hdrs)

	var _expectation Expectation // out

	_expectation = Expectation(_cret)

	return _expectation
}

// HeadersType gets the type of headers.
//
// The function returns the following values:
//
//    - messageHeadersType header's type.
//
func (hdrs *MessageHeaders) HeadersType() MessageHeadersType {
	var _arg0 *C.SoupMessageHeaders    // out
	var _cret C.SoupMessageHeadersType // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	_cret = C.soup_message_headers_get_headers_type(_arg0)
	runtime.KeepAlive(hdrs)

	var _messageHeadersType MessageHeadersType // out

	_messageHeadersType = MessageHeadersType(_cret)

	return _messageHeadersType
}

// List gets the value of header name in hdrs. Use this for headers whose values
// are comma-delimited lists, and which are therefore allowed to appear multiple
// times in the headers. For non-list-valued headers, use
// soup_message_headers_get_one().
//
// If name appears multiple times in hdrs, soup_message_headers_get_list() will
// concatenate all of the values together, separated by commas. This is
// sometimes awkward to parse (eg, WWW-Authenticate, Set-Cookie), but you have
// to be able to deal with it anyway, because the HTTP spec explicitly states
// that this transformation is allowed, and so an upstream proxy could do the
// same thing.
//
// The function takes the following parameters:
//
//    - name: header name.
//
// The function returns the following values:
//
//    - utf8 (optional) header's value or NULL if not found.
//
func (hdrs *MessageHeaders) List(name string) string {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_message_headers_get_list(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// One gets the value of header name in hdrs. Use this for headers whose values
// are <emphasis>not</emphasis> comma-delimited lists, and which therefore can
// only appear at most once in the headers. For list-valued headers, use
// soup_message_headers_get_list().
//
// If hdrs does erroneously contain multiple copies of the header, it is not
// defined which one will be returned. (Ideally, it will return whichever one
// makes libsoup most compatible with other HTTP implementations.).
//
// The function takes the following parameters:
//
//    - name: header name.
//
// The function returns the following values:
//
//    - utf8 (optional) header's value or NULL if not found.
//
func (hdrs *MessageHeaders) One(name string) string {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_message_headers_get_one(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Ranges parses hdrs's Range header and returns an array of the requested byte
// ranges. The returned array must be freed with
// soup_message_headers_free_ranges().
//
// If total_length is non-0, its value will be used to adjust the returned
// ranges to have explicit start and end values, and the returned ranges will be
// sorted and non-overlapping. If total_length is 0, then some ranges may have
// an end value of -1, as described under Range, and some of the ranges may be
// redundant.
//
// Beware that even if given a total_length, this function does not check that
// the ranges are satisfiable.
//
// <note><para> Server has built-in handling for range requests. If your server
// handler returns a SOUP_STATUS_OK response containing the complete response
// body (rather than pausing the message and returning some of the response body
// later), and there is a Range header in the request, then libsoup will
// automatically convert the response to a SOUP_STATUS_PARTIAL_CONTENT response
// containing only the range(s) requested by the client.
//
// The only time you need to process the Range header yourself is if either you
// need to stream the response body rather than returning it all at once, or you
// do not already have the complete response body available, and only want to
// generate the parts that were actually requested by the client.
// </para></note>.
//
// The function takes the following parameters:
//
//    - totalLength: total_length of the response body.
//
// The function returns the following values:
//
//    - ranges: return location for an array of Range.
//    - ok: TRUE if hdrs contained a syntactically-valid "Range" header, FALSE
//      otherwise (in which case range and length will not be set).
//
func (hdrs *MessageHeaders) Ranges(totalLength int64) ([]Range, bool) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 C.goffset             // out
	var _arg2 *C.SoupRange          // in
	var _arg3 C.int                 // in
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = C.goffset(totalLength)

	_cret = C.soup_message_headers_get_ranges(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(totalLength)

	var _ranges []Range // out
	var _ok bool        // out

	defer C.free(unsafe.Pointer(_arg2))
	{
		src := unsafe.Slice(_arg2, _arg3)
		_ranges = make([]Range, _arg3)
		for i := 0; i < int(_arg3); i++ {
			_ranges[i] = *(*Range)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_ranges[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _ranges, _ok
}

// HeaderContains checks whether the list-valued header name is present in hdrs,
// and contains a case-insensitive match for token.
//
// (If name is present in hdrs, then this is equivalent to calling
// soup_header_contains() on its value.).
//
// The function takes the following parameters:
//
//    - name: header name.
//    - token to look for.
//
// The function returns the following values:
//
//    - ok: TRUE if the header is present and contains token, FALSE otherwise.
//
func (hdrs *MessageHeaders) HeaderContains(name string, token string) bool {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _arg2 *C.char               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(token)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.soup_message_headers_header_contains(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)
	runtime.KeepAlive(token)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HeaderEquals checks whether the header name is present in hdrs and is
// (case-insensitively) equal to value.
//
// The function takes the following parameters:
//
//    - name: header name.
//    - value: expected value.
//
// The function returns the following values:
//
//    - ok: TRUE if the header is present and its value is value, FALSE
//      otherwise.
//
func (hdrs *MessageHeaders) HeaderEquals(name string, value string) bool {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _arg2 *C.char               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.soup_message_headers_header_equals(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remove removes name from hdrs. If there are multiple values for name, they
// are all removed.
//
// The function takes the following parameters:
//
//    - name: header name to remove.
//
func (hdrs *MessageHeaders) Remove(name string) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.soup_message_headers_remove(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)
}

// Replace replaces the value of the header name in hdrs with value. (See also
// soup_message_headers_append().)
//
// The caller is expected to make sure that name and value are syntactically
// correct.
//
// The function takes the following parameters:
//
//    - name: header name to replace.
//    - value: new value of name.
//
func (hdrs *MessageHeaders) Replace(name string, value string) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _arg2 *C.char               // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))

	C.soup_message_headers_replace(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// SetContentDisposition sets the "Content-Disposition" header in hdrs to
// disposition, optionally with additional parameters specified in params.
//
// See soup_message_headers_get_content_disposition() for a discussion of how
// Content-Disposition is used in HTTP.
//
// The function takes the following parameters:
//
//    - disposition: disposition-type.
//    - params (optional): additional parameters, or NULL.
//
func (hdrs *MessageHeaders) SetContentDisposition(disposition string, params map[string]string) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _arg2 *C.GHashTable         // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(disposition)))
	defer C.free(unsafe.Pointer(_arg1))
	if params != nil {
		_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
		for ksrc, vsrc := range params {
			var kdst *C.gchar // out
			var vdst *C.gchar // out
			kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
			defer C.free(unsafe.Pointer(kdst))
			vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
			defer C.free(unsafe.Pointer(vdst))
			C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
		}
		defer C.g_hash_table_unref(_arg2)
	}

	C.soup_message_headers_set_content_disposition(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(disposition)
	runtime.KeepAlive(params)
}

// SetContentLength sets the message body length that hdrs will declare, and
// sets hdrs's encoding to SOUP_ENCODING_CONTENT_LENGTH.
//
// You do not normally need to call this; if hdrs is set to use Content-Length
// encoding, libsoup will automatically set its Content-Length header for you
// immediately before sending the headers. One situation in which this method is
// useful is when generating the response to a HEAD request; Calling
// soup_message_headers_set_content_length() allows you to put the correct
// content length into the response without needing to waste memory by filling
// in a response body which won't actually be sent.
//
// The function takes the following parameters:
//
//    - contentLength: message body length.
//
func (hdrs *MessageHeaders) SetContentLength(contentLength int64) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 C.goffset             // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = C.goffset(contentLength)

	C.soup_message_headers_set_content_length(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(contentLength)
}

// SetContentRange sets hdrs's Content-Range header according to the given
// values. (Note that total_length is the total length of the entire resource
// that this is a range of, not simply end - start + 1.)
//
// <note><para> Server has built-in handling for range requests, and you do not
// normally need to call this function youself. See
// soup_message_headers_get_ranges() for more details. </para></note>.
//
// The function takes the following parameters:
//
//    - start of the range.
//    - end of the range.
//    - totalLength: total length of the resource, or -1 if unknown.
//
func (hdrs *MessageHeaders) SetContentRange(start int64, end int64, totalLength int64) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 C.goffset             // out
	var _arg2 C.goffset             // out
	var _arg3 C.goffset             // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = C.goffset(start)
	_arg2 = C.goffset(end)
	_arg3 = C.goffset(totalLength)

	C.soup_message_headers_set_content_range(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
	runtime.KeepAlive(totalLength)
}

// SetContentType sets the "Content-Type" header in hdrs to content_type,
// optionally with additional parameters specified in params.
//
// The function takes the following parameters:
//
//    - contentType: MIME type.
//    - params (optional): additional parameters, or NULL.
//
func (hdrs *MessageHeaders) SetContentType(contentType string, params map[string]string) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.char               // out
	var _arg2 *C.GHashTable         // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg1))
	if params != nil {
		_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
		for ksrc, vsrc := range params {
			var kdst *C.gchar // out
			var vdst *C.gchar // out
			kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
			defer C.free(unsafe.Pointer(kdst))
			vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
			defer C.free(unsafe.Pointer(vdst))
			C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
		}
		defer C.g_hash_table_unref(_arg2)
	}

	C.soup_message_headers_set_content_type(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(contentType)
	runtime.KeepAlive(params)
}

// SetEncoding sets the message body encoding that hdrs will declare. In
// particular, you should use this if you are going to send a request or
// response in chunked encoding.
//
// The function takes the following parameters:
//
//    - encoding: Encoding.
//
func (hdrs *MessageHeaders) SetEncoding(encoding Encoding) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 C.SoupEncoding        // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = C.SoupEncoding(encoding)

	C.soup_message_headers_set_encoding(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(encoding)
}

// SetExpectations sets hdrs's "Expect" header according to expectations.
//
// Currently SOUP_EXPECTATION_CONTINUE is the only known expectation value. You
// should set this value on a request if you are sending a large message body
// (eg, via POST or PUT), and want to give the server a chance to reject the
// request after seeing just the headers (eg, because it will require
// authentication before allowing you to post, or because you're POSTing to a
// URL that doesn't exist). This saves you from having to transmit the large
// request body when the server is just going to ignore it anyway.
//
// The function takes the following parameters:
//
//    - expectations to set.
//
func (hdrs *MessageHeaders) SetExpectations(expectations Expectation) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 C.SoupExpectation     // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = C.SoupExpectation(expectations)

	C.soup_message_headers_set_expectations(_arg0, _arg1)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(expectations)
}

// SetRange sets hdrs's Range header to request the indicated range. start and
// end are interpreted as in a Range.
//
// If you need to request multiple ranges, use
// soup_message_headers_set_ranges().
//
// The function takes the following parameters:
//
//    - start of the range to request.
//    - end of the range to request.
//
func (hdrs *MessageHeaders) SetRange(start int64, end int64) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 C.goffset             // out
	var _arg2 C.goffset             // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = C.goffset(start)
	_arg2 = C.goffset(end)

	C.soup_message_headers_set_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)
}

// SetRanges sets hdrs's Range header to request the indicated ranges. (If you
// only want to request a single range, you can use
// soup_message_headers_set_range().).
//
// The function takes the following parameters:
//
//    - ranges: array of Range.
//    - length of range.
//
func (hdrs *MessageHeaders) SetRanges(ranges *Range, length int) {
	var _arg0 *C.SoupMessageHeaders // out
	var _arg1 *C.SoupRange          // out
	var _arg2 C.int                 // out

	_arg0 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))
	_arg1 = (*C.SoupRange)(gextras.StructNative(unsafe.Pointer(ranges)))
	_arg2 = C.int(length)

	C.soup_message_headers_set_ranges(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hdrs)
	runtime.KeepAlive(ranges)
	runtime.KeepAlive(length)
}

// MessageHeadersIter: opaque type used to iterate over a SoupMessageHeaders
// structure.
//
// After intializing the iterator with soup_message_headers_iter_init(), call
// soup_message_headers_iter_next() to fetch data from it.
//
// You may not modify the headers while iterating over them.
//
// An instance of this type is always passed by reference.
type MessageHeadersIter struct {
	*messageHeadersIter
}

// messageHeadersIter is the struct that's finalized.
type messageHeadersIter struct {
	native *C.SoupMessageHeadersIter
}

// Next yields the next name/value pair in the SoupMessageHeaders being iterated
// by iter. If iter has already yielded the last header, then
// soup_message_headers_iter_next() will return FALSE and name and value will be
// unchanged.
//
// The function returns the following values:
//
//    - name: pointer to a variable to return the header name in.
//    - value: pointer to a variable to return the header value in.
//    - ok: TRUE if another name and value were returned, FALSE if the end of the
//      headers has been reached.
//
func (iter *MessageHeadersIter) Next() (name string, value string, ok bool) {
	var _arg0 *C.SoupMessageHeadersIter // out
	var _arg1 *C.char                   // in
	var _arg2 *C.char                   // in
	var _cret C.gboolean                // in

	_arg0 = (*C.SoupMessageHeadersIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.soup_message_headers_iter_next(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(iter)

	var _name string  // out
	var _value string // out
	var _ok bool      // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	_value = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	if _cret != 0 {
		_ok = true
	}

	return _name, _value, _ok
}

// MessageHeadersIterInit initializes iter for iterating hdrs.
//
// The function takes the following parameters:
//
//    - hdrs: SoupMessageHeaders.
//
// The function returns the following values:
//
//    - iter: pointer to a SoupMessageHeadersIter structure.
//
func MessageHeadersIterInit(hdrs *MessageHeaders) *MessageHeadersIter {
	var _arg1 C.SoupMessageHeadersIter // in
	var _arg2 *C.SoupMessageHeaders    // out

	_arg2 = (*C.SoupMessageHeaders)(gextras.StructNative(unsafe.Pointer(hdrs)))

	C.soup_message_headers_iter_init(&_arg1, _arg2)
	runtime.KeepAlive(hdrs)

	var _iter *MessageHeadersIter // out

	_iter = (*MessageHeadersIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Range represents a byte range as used in the Range header.
//
// If end is non-negative, then start and end represent the bounds of of the
// range, counting from 0. (Eg, the first 500 bytes would be represented as
// start = 0 and end = 499.)
//
// If end is -1 and start is non-negative, then this represents a range starting
// at start and ending with the last byte of the requested resource body. (Eg,
// all but the first 500 bytes would be start = 500, and end = -1.)
//
// If end is -1 and start is negative, then it represents a "suffix range",
// referring to the last -start bytes of the resource body. (Eg, the last 500
// bytes would be start = -500 and end = -1.)
//
// An instance of this type is always passed by reference.
type Range struct {
	*_range
}

// _range is the struct that's finalized.
type _range struct {
	native *C.SoupRange
}

// NewRange creates a new Range instance from the given
// fields.
func NewRange(start, end int64) Range {
	var f0 C.goffset // out
	f0 = C.goffset(start)
	var f1 C.goffset // out
	f1 = C.goffset(end)

	v := C.SoupRange{
		start: f0,
		end:   f1,
	}

	return *(*Range)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// Start: start of the range.
func (r *Range) Start() int64 {
	var v int64 // out
	v = int64(r.native.start)
	return v
}

// End: end of the range.
func (r *Range) End() int64 {
	var v int64 // out
	v = int64(r.native.end)
	return v
}
