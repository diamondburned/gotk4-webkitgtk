// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
// SoupLoggerLogLevel _gotk4_soup2_LoggerFilter(SoupLogger*, SoupMessage*, gpointer);
// extern void callbackDelete(gpointer);
// void _gotk4_soup2_LoggerPrinter(SoupLogger*, SoupLoggerLogLevel, char, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_logger_log_level_get_type()), F: marshalLoggerLogLevel},
		{T: externglib.Type(C.soup_logger_get_type()), F: marshalLoggerer},
	})
}

// LOGGER_LEVEL alias for the Logger:level property, qv.
const LOGGER_LEVEL = "level"

// LOGGER_MAX_BODY_SIZE alias for the Logger:max-body-size property, qv.
const LOGGER_MAX_BODY_SIZE = "max-body-size"

// LoggerLogLevel describes the level of logging output to provide.
type LoggerLogLevel int

const (
	// LoggerLogNone: no logging.
	LoggerLogNone LoggerLogLevel = iota
	// LoggerLogMinimal: log the Request-Line or Status-Line and the Soup-Debug
	// pseudo-headers.
	LoggerLogMinimal
	// LoggerLogHeaders: log the full request/response headers.
	LoggerLogHeaders
	// LoggerLogBody: log the full headers and request/response bodies.
	LoggerLogBody
)

func marshalLoggerLogLevel(p uintptr) (interface{}, error) {
	return LoggerLogLevel(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for LoggerLogLevel.
func (l LoggerLogLevel) String() string {
	switch l {
	case LoggerLogNone:
		return "None"
	case LoggerLogMinimal:
		return "Minimal"
	case LoggerLogHeaders:
		return "Headers"
	case LoggerLogBody:
		return "Body"
	default:
		return fmt.Sprintf("LoggerLogLevel(%d)", l)
	}
}

// LoggerFilter: prototype for a logging filter. The filter callback will be
// invoked for each request or response, and should analyze it and return a
// LoggerLogLevel value indicating how much of the message to log. Eg, it might
// choose between SOUP_LOGGER_LOG_BODY and SOUP_LOGGER_LOG_HEADERS depending on
// the Content-Type.
type LoggerFilter func(logger *Logger, msg *Message) (loggerLogLevel LoggerLogLevel)

//export _gotk4_soup2_LoggerFilter
func _gotk4_soup2_LoggerFilter(arg0 *C.SoupLogger, arg1 *C.SoupMessage, arg2 C.gpointer) (cret C.SoupLoggerLogLevel) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var logger *Logger // out
	var msg *Message   // out

	logger = wrapLogger(externglib.Take(unsafe.Pointer(arg0)))
	msg = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))

	fn := v.(LoggerFilter)
	loggerLogLevel := fn(logger, msg)

	cret = C.SoupLoggerLogLevel(loggerLogLevel)

	return cret
}

// LoggerPrinter: prototype for a custom printing callback.
//
// level indicates what kind of information is being printed. Eg, it will be
// SOUP_LOGGER_LOG_HEADERS if data is header data.
//
// direction is either '<', '>', or ' ', and data is the single line to print;
// the printer is expected to add a terminating newline.
//
// To get the effect of the default printer, you would do:
//
// <informalexample><programlisting> printf ("c s\n", direction, data);
// </programlisting></informalexample>.
type LoggerPrinter func(logger *Logger, level LoggerLogLevel, direction byte, data string)

//export _gotk4_soup2_LoggerPrinter
func _gotk4_soup2_LoggerPrinter(arg0 *C.SoupLogger, arg1 C.SoupLoggerLogLevel, arg2 C.char, arg3 *C.char, arg4 C.gpointer) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var logger *Logger       // out
	var level LoggerLogLevel // out
	var direction byte       // out
	var data string          // out

	logger = wrapLogger(externglib.Take(unsafe.Pointer(arg0)))
	level = LoggerLogLevel(arg1)
	direction = byte(arg2)
	data = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	fn := v.(LoggerPrinter)
	fn(logger, level, direction, data)
}

type Logger struct {
	*externglib.Object

	SessionFeature
}

func wrapLogger(obj *externglib.Object) *Logger {
	return &Logger{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalLoggerer(p uintptr) (interface{}, error) {
	return wrapLogger(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLogger creates a new Logger with the given debug level. If level is
// SOUP_LOGGER_LOG_BODY, max_body_size gives the maximum number of bytes of the
// body that will be logged. (-1 means "no limit".)
//
// If you need finer control over what message parts are and aren't logged, use
// soup_logger_set_request_filter() and soup_logger_set_response_filter().
//
// The function takes the following parameters:
//
//    - level: debug level.
//    - maxBodySize: maximum body size to output, or -1.
//
func NewLogger(level LoggerLogLevel, maxBodySize int) *Logger {
	var _arg1 C.SoupLoggerLogLevel // out
	var _arg2 C.int                // out
	var _cret *C.SoupLogger        // in

	_arg1 = C.SoupLoggerLogLevel(level)
	_arg2 = C.int(maxBodySize)

	_cret = C.soup_logger_new(_arg1, _arg2)
	runtime.KeepAlive(level)
	runtime.KeepAlive(maxBodySize)

	var _logger *Logger // out

	_logger = wrapLogger(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _logger
}

// Attach sets logger to watch session and print debug information for its
// messages.
//
// (The session will take a reference on logger, which will be removed when you
// call soup_logger_detach(), or when the session is destroyed.)
//
// Deprecated: Use soup_session_add_feature() instead.
//
// The function takes the following parameters:
//
//    - session: Session.
//
func (logger *Logger) Attach(session *Session) {
	var _arg0 *C.SoupLogger  // out
	var _arg1 *C.SoupSession // out

	_arg0 = (*C.SoupLogger)(unsafe.Pointer(logger.Native()))
	_arg1 = (*C.SoupSession)(unsafe.Pointer(session.Native()))

	C.soup_logger_attach(_arg0, _arg1)
	runtime.KeepAlive(logger)
	runtime.KeepAlive(session)
}

// Detach stops logger from watching session.
//
// Deprecated: Use soup_session_remove_feature() instead.
//
// The function takes the following parameters:
//
//    - session: Session.
//
func (logger *Logger) Detach(session *Session) {
	var _arg0 *C.SoupLogger  // out
	var _arg1 *C.SoupSession // out

	_arg0 = (*C.SoupLogger)(unsafe.Pointer(logger.Native()))
	_arg1 = (*C.SoupSession)(unsafe.Pointer(session.Native()))

	C.soup_logger_detach(_arg0, _arg1)
	runtime.KeepAlive(logger)
	runtime.KeepAlive(session)
}

// SetPrinter sets up an alternate log printing routine, if you don't want the
// log to go to <literal>stdout</literal>.
//
// The function takes the following parameters:
//
//    - printer: callback for printing logging output.
//
func (logger *Logger) SetPrinter(printer LoggerPrinter) {
	var _arg0 *C.SoupLogger       // out
	var _arg1 C.SoupLoggerPrinter // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupLogger)(unsafe.Pointer(logger.Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup2_LoggerPrinter)
	_arg2 = C.gpointer(gbox.Assign(printer))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_logger_set_printer(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(logger)
	runtime.KeepAlive(printer)
}

// SetRequestFilter sets up a filter to determine the log level for a given
// request. For each HTTP request logger will invoke request_filter to determine
// how much (if any) of that request to log. (If you do not set a request
// filter, logger will just always log requests at the level passed to
// soup_logger_new().).
//
// The function takes the following parameters:
//
//    - requestFilter: callback for request debugging.
//
func (logger *Logger) SetRequestFilter(requestFilter LoggerFilter) {
	var _arg0 *C.SoupLogger      // out
	var _arg1 C.SoupLoggerFilter // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupLogger)(unsafe.Pointer(logger.Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup2_LoggerFilter)
	_arg2 = C.gpointer(gbox.Assign(requestFilter))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_logger_set_request_filter(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(logger)
	runtime.KeepAlive(requestFilter)
}

// SetResponseFilter sets up a filter to determine the log level for a given
// response. For each HTTP response logger will invoke response_filter to
// determine how much (if any) of that response to log. (If you do not set a
// response filter, logger will just always log responses at the level passed to
// soup_logger_new().).
//
// The function takes the following parameters:
//
//    - responseFilter: callback for response debugging.
//
func (logger *Logger) SetResponseFilter(responseFilter LoggerFilter) {
	var _arg0 *C.SoupLogger      // out
	var _arg1 C.SoupLoggerFilter // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.SoupLogger)(unsafe.Pointer(logger.Native()))
	_arg1 = (*[0]byte)(C._gotk4_soup2_LoggerFilter)
	_arg2 = C.gpointer(gbox.Assign(responseFilter))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.soup_logger_set_response_filter(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(logger)
	runtime.KeepAlive(responseFilter)
}
