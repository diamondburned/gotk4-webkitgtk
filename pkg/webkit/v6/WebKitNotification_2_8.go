// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
// extern void _gotk4_webkit6_Notification_ConnectClosed(gpointer, guintptr);
// extern void _gotk4_webkit6_Notification_ConnectClicked(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeNotification = coreglib.Type(C.webkit_notification_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNotification, F: marshalNotification},
	})
}

// NotificationOverrides contains methods that are overridable.
type NotificationOverrides struct {
}

func defaultNotificationOverrides(v *Notification) NotificationOverrides {
	return NotificationOverrides{}
}

// Notification holds information about a notification that should be shown to
// the user.
type Notification struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Notification)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Notification, *NotificationClass, NotificationOverrides](
		GTypeNotification,
		initNotificationClass,
		wrapNotification,
		defaultNotificationOverrides,
	)
}

func initNotificationClass(gclass unsafe.Pointer, overrides NotificationOverrides, classInitFunc func(*NotificationClass)) {
	if classInitFunc != nil {
		class := (*NotificationClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapNotification(obj *coreglib.Object) *Notification {
	return &Notification{
		Object: obj,
	}
}

func marshalNotification(p uintptr) (interface{}, error) {
	return wrapNotification(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClicked is emitted when a notification has been clicked. See
// webkit_notification_clicked().
func (notification *Notification) ConnectClicked(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(notification, "clicked", false, unsafe.Pointer(C._gotk4_webkit6_Notification_ConnectClicked), f)
}

// ConnectClosed is emitted when a notification has been withdrawn.
//
// The default handler will close the notification using libnotify, if built
// with support for it.
func (notification *Notification) ConnectClosed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(notification, "closed", false, unsafe.Pointer(C._gotk4_webkit6_Notification_ConnectClosed), f)
}

// Clicked tells WebKit the notification has been clicked.
//
// This will emit the KitNotification::clicked signal.
func (notification *Notification) Clicked() {
	var _arg0 *C.WebKitNotification // out

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))

	C.webkit_notification_clicked(_arg0)
	runtime.KeepAlive(notification)
}

// Close closes the notification.
func (notification *Notification) Close() {
	var _arg0 *C.WebKitNotification // out

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))

	C.webkit_notification_close(_arg0)
	runtime.KeepAlive(notification)
}

// Body obtains the body for the notification.
//
// The function returns the following values:
//
//   - utf8: body for the notification.
//
func (notification *Notification) Body() string {
	var _arg0 *C.WebKitNotification // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))

	_cret = C.webkit_notification_get_body(_arg0)
	runtime.KeepAlive(notification)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ID obtains the unique id for the notification.
//
// The function returns the following values:
//
//   - guint64: unique id for the notification.
//
func (notification *Notification) ID() uint64 {
	var _arg0 *C.WebKitNotification // out
	var _cret C.guint64             // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))

	_cret = C.webkit_notification_get_id(_arg0)
	runtime.KeepAlive(notification)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Tag obtains the tag identifier for the notification.
//
// The function returns the following values:
//
//   - utf8 (optional): tag for the notification.
//
func (notification *Notification) Tag() string {
	var _arg0 *C.WebKitNotification // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))

	_cret = C.webkit_notification_get_tag(_arg0)
	runtime.KeepAlive(notification)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title obtains the title for the notification.
//
// The function returns the following values:
//
//   - utf8: title for the notification.
//
func (notification *Notification) Title() string {
	var _arg0 *C.WebKitNotification // out
	var _cret *C.gchar              // in

	_arg0 = (*C.WebKitNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))

	_cret = C.webkit_notification_get_title(_arg0)
	runtime.KeepAlive(notification)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
