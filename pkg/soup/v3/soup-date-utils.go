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
	GTypeDateFormat = coreglib.Type(C.soup_date_format_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDateFormat, F: marshalDateFormat},
	})
}

// DateFormat: date formats that date_time_to_string can use.
//
// SOUP_DATE_HTTP and SOUP_DATE_COOKIE always coerce the time to UTC.
//
// This enum may be extended with more values in future releases.
type DateFormat C.gint

const (
	// DateHTTP: RFC 1123 format, used by the HTTP "Date" header. Eg "Sun,
	// 06 Nov 1994 08:49:37 GMT".
	DateHTTP DateFormat = 1
	// DateCookie: format for the "Expires" timestamp in the Netscape cookie
	// specification. Eg, "Sun, 06-Nov-1994 08:49:37 GMT".
	DateCookie DateFormat = 2
)

func marshalDateFormat(p uintptr) (interface{}, error) {
	return DateFormat(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DateFormat.
func (d DateFormat) String() string {
	switch d {
	case DateHTTP:
		return "HTTP"
	case DateCookie:
		return "Cookie"
	default:
		return fmt.Sprintf("DateFormat(%d)", d)
	}
}

// DateTimeNewFromHTTPString parses date_string and tries to extract a date from
// it.
//
// This recognizes all of the "HTTP-date" formats from RFC 2616, RFC 2822 dates,
// and reasonable approximations thereof. (Eg, it is lenient about whitespace,
// leading "0"s, etc.).
//
// The function takes the following parameters:
//
//   - dateString: date as a string.
//
// The function returns the following values:
//
//   - dateTime (optional): new Time, or NULL if date_string could not be
//     parsed.
//
func DateTimeNewFromHTTPString(dateString string) *glib.DateTime {
	var _arg1 *C.char      // out
	var _cret *C.GDateTime // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(dateString)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_date_time_new_from_http_string(_arg1)
	runtime.KeepAlive(dateString)

	var _dateTime *glib.DateTime // out

	if _cret != nil {
		_dateTime = (*glib.DateTime)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dateTime)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_date_time_unref((*C.GDateTime)(intern.C))
			},
		)
	}

	return _dateTime
}

// DateTimeToString converts date to a string in the format described by format.
//
// The function takes the following parameters:
//
//   - date: Time.
//   - format to generate the date in.
//
// The function returns the following values:
//
//   - utf8: date as a string or NULL.
//
func DateTimeToString(date *glib.DateTime, format DateFormat) string {
	var _arg1 *C.GDateTime     // out
	var _arg2 C.SoupDateFormat // out
	var _cret *C.char          // in

	_arg1 = (*C.GDateTime)(gextras.StructNative(unsafe.Pointer(date)))
	_arg2 = C.SoupDateFormat(format)

	_cret = C.soup_date_time_to_string(_arg1, _arg2)
	runtime.KeepAlive(date)
	runtime.KeepAlive(format)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}