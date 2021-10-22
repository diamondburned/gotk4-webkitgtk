// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_print_custom_widget_get_type()), F: marshalPrintCustomWidgetter},
	})
}

// PrintCustomWidgetOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PrintCustomWidgetOverrider interface {
	Apply(widget gtk.Widgetter)
	Update(widget gtk.Widgetter, pageSetup *gtk.PageSetup, printSettings *gtk.PrintSettings)
}

type PrintCustomWidget struct {
	*externglib.Object
}

func wrapPrintCustomWidget(obj *externglib.Object) *PrintCustomWidget {
	return &PrintCustomWidget{
		Object: obj,
	}
}

func marshalPrintCustomWidgetter(p uintptr) (interface{}, error) {
	return wrapPrintCustomWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPrintCustomWidget: create a new KitPrintCustomWidget with given widget and
// title. The widget ownership is taken and it is destroyed together with the
// dialog even if this object could still be alive at that point. You typically
// want to pass a container widget with multiple widgets in it.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//    - title widget's title.
//
func NewPrintCustomWidget(widget gtk.Widgetter, title string) *PrintCustomWidget {
	var _arg1 *C.GtkWidget               // out
	var _arg2 *C.char                    // out
	var _cret *C.WebKitPrintCustomWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.webkit_print_custom_widget_new(_arg1, _arg2)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(title)

	var _printCustomWidget *PrintCustomWidget // out

	_printCustomWidget = wrapPrintCustomWidget(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printCustomWidget
}

// Title: return the value of KitPrintCustomWidget:title property for the given
// print_custom_widget object.
func (printCustomWidget *PrintCustomWidget) Title() string {
	var _arg0 *C.WebKitPrintCustomWidget // out
	var _cret *C.gchar                   // in

	_arg0 = (*C.WebKitPrintCustomWidget)(unsafe.Pointer(printCustomWidget.Native()))

	_cret = C.webkit_print_custom_widget_get_title(_arg0)
	runtime.KeepAlive(printCustomWidget)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Widget: return the value of KitPrintCustomWidget:widget property for the
// given print_custom_widget object. The returned value will always be valid if
// called from KitPrintCustomWidget::apply or KitPrintCustomWidget::update
// callbacks, but it will be NULL if called after the
// KitPrintCustomWidget::apply signal is emitted.
func (printCustomWidget *PrintCustomWidget) Widget() gtk.Widgetter {
	var _arg0 *C.WebKitPrintCustomWidget // out
	var _cret *C.GtkWidget               // in

	_arg0 = (*C.WebKitPrintCustomWidget)(unsafe.Pointer(printCustomWidget.Native()))

	_cret = C.webkit_print_custom_widget_get_widget(_arg0)
	runtime.KeepAlive(printCustomWidget)

	var _widget gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// ConnectApply: emitted right before the printing will start. You should read
// the information from the widget and update the content based on it if
// necessary. The widget is not guaranteed to be valid at a later time.
func (printCustomWidget *PrintCustomWidget) ConnectApply(f func()) externglib.SignalHandle {
	return printCustomWidget.Connect("apply", f)
}

// ConnectUpdate: emitted after change of selected printer in the dialog. The
// actual page setup and print settings are available and the custom widget can
// actualize itself according to their values.
func (printCustomWidget *PrintCustomWidget) ConnectUpdate(f func(pageSetup gtk.PageSetup, printSettings gtk.PrintSettings)) externglib.SignalHandle {
	return printCustomWidget.Connect("update", f)
}
