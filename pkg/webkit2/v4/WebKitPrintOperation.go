// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
// extern void _gotk4_webkit24_PrintOperation_ConnectFinished(gpointer, guintptr);
// extern void _gotk4_webkit24_PrintOperation_ConnectFailed(gpointer, GError*, guintptr);
// extern WebKitPrintCustomWidget* _gotk4_webkit24_PrintOperation_ConnectCreateCustomWidget(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypePrintOperationResponse = coreglib.Type(C.webkit_print_operation_response_get_type())
	GTypePrintOperation         = coreglib.Type(C.webkit_print_operation_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePrintOperationResponse, F: marshalPrintOperationResponse},
		coreglib.TypeMarshaler{T: GTypePrintOperation, F: marshalPrintOperation},
	})
}

// PrintOperationResponse: enum values representing the response of the print
// dialog shown with webkit_print_operation_run_dialog().
type PrintOperationResponse C.gint

const (
	// PrintOperationResponsePrint: print button was clicked in print dialog.
	PrintOperationResponsePrint PrintOperationResponse = iota
	// PrintOperationResponseCancel: print dialog was cancelled.
	PrintOperationResponseCancel
)

func marshalPrintOperationResponse(p uintptr) (interface{}, error) {
	return PrintOperationResponse(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PrintOperationResponse.
func (p PrintOperationResponse) String() string {
	switch p {
	case PrintOperationResponsePrint:
		return "Print"
	case PrintOperationResponseCancel:
		return "Cancel"
	default:
		return fmt.Sprintf("PrintOperationResponse(%d)", p)
	}
}

// PrintOperationOverrides contains methods that are overridable.
type PrintOperationOverrides struct {
}

func defaultPrintOperationOverrides(v *PrintOperation) PrintOperationOverrides {
	return PrintOperationOverrides{}
}

// PrintOperation controls a print operation.
//
// A KitPrintOperation controls a print operation in WebKit. With a
// similar API to PrintOperation, it lets you set the print settings with
// webkit_print_operation_set_print_settings() or display the print dialog with
// webkit_print_operation_run_dialog().
type PrintOperation struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PrintOperation)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*PrintOperation, *PrintOperationClass, PrintOperationOverrides](
		GTypePrintOperation,
		initPrintOperationClass,
		wrapPrintOperation,
		defaultPrintOperationOverrides,
	)
}

func initPrintOperationClass(gclass unsafe.Pointer, overrides PrintOperationOverrides, classInitFunc func(*PrintOperationClass)) {
	if classInitFunc != nil {
		class := (*PrintOperationClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPrintOperation(obj *coreglib.Object) *PrintOperation {
	return &PrintOperation{
		Object: obj,
	}
}

func marshalPrintOperation(p uintptr) (interface{}, error) {
	return wrapPrintOperation(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectCreateCustomWidget is emitted when displaying the print dialog with
// webkit_print_operation_run_dialog(). The returned KitPrintCustomWidget
// will be added to the print dialog and it will be owned by the
// print_operation. However, the object is guaranteed to be alive until the
// KitPrintCustomWidget::apply is emitted.
func (printOperation *PrintOperation) ConnectCreateCustomWidget(f func() (printCustomWidget *PrintCustomWidget)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(printOperation, "create-custom-widget", false, unsafe.Pointer(C._gotk4_webkit24_PrintOperation_ConnectCreateCustomWidget), f)
}

// ConnectFailed is emitted when an error occurs while printing. The given
// error, of the domain WEBKIT_PRINT_ERROR, contains further details of the
// failure. The KitPrintOperation::finished signal is emitted after this one.
func (printOperation *PrintOperation) ConnectFailed(f func(err error)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(printOperation, "failed", false, unsafe.Pointer(C._gotk4_webkit24_PrintOperation_ConnectFailed), f)
}

// ConnectFinished is emitted when the print operation has finished doing
// everything required for printing.
func (printOperation *PrintOperation) ConnectFinished(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(printOperation, "finished", false, unsafe.Pointer(C._gotk4_webkit24_PrintOperation_ConnectFinished), f)
}

// NewPrintOperation: create a new KitPrintOperation to print web_view contents.
//
// The function takes the following parameters:
//
//   - webView: KitWebView.
//
// The function returns the following values:
//
//   - printOperation: new KitPrintOperation.
//
func NewPrintOperation(webView *WebView) *PrintOperation {
	var _arg1 *C.WebKitWebView        // out
	var _cret *C.WebKitPrintOperation // in

	_arg1 = (*C.WebKitWebView)(unsafe.Pointer(coreglib.InternObject(webView).Native()))

	_cret = C.webkit_print_operation_new(_arg1)
	runtime.KeepAlive(webView)

	var _printOperation *PrintOperation // out

	_printOperation = wrapPrintOperation(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printOperation
}

// PageSetup: return the current page setup of print_operation.
//
// It returns NULL until either webkit_print_operation_set_page_setup() or
// webkit_print_operation_run_dialog() have been called.
//
// The function returns the following values:
//
//   - pageSetup: current PageSetup of print_operation.
//
func (printOperation *PrintOperation) PageSetup() *gtk.PageSetup {
	var _arg0 *C.WebKitPrintOperation // out
	var _cret *C.GtkPageSetup         // in

	_arg0 = (*C.WebKitPrintOperation)(unsafe.Pointer(coreglib.InternObject(printOperation).Native()))

	_cret = C.webkit_print_operation_get_page_setup(_arg0)
	runtime.KeepAlive(printOperation)

	var _pageSetup *gtk.PageSetup // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_pageSetup = &gtk.PageSetup{
			Object: obj,
		}
	}

	return _pageSetup
}

// PrintSettings: return the current print settings of print_operation.
//
// It returns NULL until either webkit_print_operation_set_print_settings() or
// webkit_print_operation_run_dialog() have been called.
//
// The function returns the following values:
//
//   - printSettings: current PrintSettings of print_operation.
//
func (printOperation *PrintOperation) PrintSettings() *gtk.PrintSettings {
	var _arg0 *C.WebKitPrintOperation // out
	var _cret *C.GtkPrintSettings     // in

	_arg0 = (*C.WebKitPrintOperation)(unsafe.Pointer(coreglib.InternObject(printOperation).Native()))

	_cret = C.webkit_print_operation_get_print_settings(_arg0)
	runtime.KeepAlive(printOperation)

	var _printSettings *gtk.PrintSettings // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_printSettings = &gtk.PrintSettings{
			Object: obj,
		}
	}

	return _printSettings
}

// Print: start a print operation using current print settings and page setup.
//
// Start a print operation using current print settings and page setup
// without showing the print dialog. If either print settings or page
// setup are not set with webkit_print_operation_set_print_settings()
// and webkit_print_operation_set_page_setup(), the default options
// will be used and the print job will be sent to the default printer.
// The KitPrintOperation::finished signal is emitted when the printing
// operation finishes. If an error occurs while printing the signal
// KitPrintOperation::failed is emitted before KitPrintOperation::finished.
func (printOperation *PrintOperation) Print() {
	var _arg0 *C.WebKitPrintOperation // out

	_arg0 = (*C.WebKitPrintOperation)(unsafe.Pointer(coreglib.InternObject(printOperation).Native()))

	C.webkit_print_operation_print(_arg0)
	runtime.KeepAlive(printOperation)
}

// RunDialog: run the print dialog and start printing.
//
// Run the print dialog and start printing using the options selected by the
// user. This method returns when the print dialog is closed. If the print
// dialog is cancelled WEBKIT_PRINT_OPERATION_RESPONSE_CANCEL is returned.
// If the user clicks on the print button, WEBKIT_PRINT_OPERATION_RESPONSE_PRINT
// is returned and the print operation starts. In this case, the
// KitPrintOperation::finished signal is emitted when the operation finishes.
// If an error occurs while printing, the signal KitPrintOperation::failed
// is emitted before KitPrintOperation::finished. If the print dialog is
// not cancelled current print settings and page setup of print_operation
// are updated with options selected by the user when Print button is
// pressed in print dialog. You can get the updated print settings and
// page setup by calling webkit_print_operation_get_print_settings() and
// webkit_print_operation_get_page_setup() after this method.
//
// The function takes the following parameters:
//
//   - parent (optional): transient parent of the print dialog.
//
// The function returns the following values:
//
//   - printOperationResponse of the print dialog.
//
func (printOperation *PrintOperation) RunDialog(parent *gtk.Window) PrintOperationResponse {
	var _arg0 *C.WebKitPrintOperation        // out
	var _arg1 *C.GtkWindow                   // out
	var _cret C.WebKitPrintOperationResponse // in

	_arg0 = (*C.WebKitPrintOperation)(unsafe.Pointer(coreglib.InternObject(printOperation).Native()))
	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_cret = C.webkit_print_operation_run_dialog(_arg0, _arg1)
	runtime.KeepAlive(printOperation)
	runtime.KeepAlive(parent)

	var _printOperationResponse PrintOperationResponse // out

	_printOperationResponse = PrintOperationResponse(_cret)

	return _printOperationResponse
}

// SetPageSetup: set the current page setup of print_operation.
//
// Current page setup is used for the initial values of the print dialog when
// webkit_print_operation_run_dialog() is called.
//
// The function takes the following parameters:
//
//   - pageSetup to set.
//
func (printOperation *PrintOperation) SetPageSetup(pageSetup *gtk.PageSetup) {
	var _arg0 *C.WebKitPrintOperation // out
	var _arg1 *C.GtkPageSetup         // out

	_arg0 = (*C.WebKitPrintOperation)(unsafe.Pointer(coreglib.InternObject(printOperation).Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(pageSetup).Native()))

	C.webkit_print_operation_set_page_setup(_arg0, _arg1)
	runtime.KeepAlive(printOperation)
	runtime.KeepAlive(pageSetup)
}

// SetPrintSettings: set the current print settings of print_operation.
//
// Set the current print settings of print_operation. Current print
// settings are used for the initial values of the print dialog when
// webkit_print_operation_run_dialog() is called.
//
// The function takes the following parameters:
//
//   - printSettings to set.
//
func (printOperation *PrintOperation) SetPrintSettings(printSettings *gtk.PrintSettings) {
	var _arg0 *C.WebKitPrintOperation // out
	var _arg1 *C.GtkPrintSettings     // out

	_arg0 = (*C.WebKitPrintOperation)(unsafe.Pointer(coreglib.InternObject(printOperation).Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer(coreglib.InternObject(printSettings).Native()))

	C.webkit_print_operation_set_print_settings(_arg0, _arg1)
	runtime.KeepAlive(printOperation)
	runtime.KeepAlive(printSettings)
}

// PrintOperationClass: instance of this type is always passed by reference.
type PrintOperationClass struct {
	*printOperationClass
}

// printOperationClass is the struct that's finalized.
type printOperationClass struct {
	native *C.WebKitPrintOperationClass
}
