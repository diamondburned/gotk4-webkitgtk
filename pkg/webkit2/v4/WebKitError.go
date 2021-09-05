// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: webkit2gtk-4.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_download_error_get_type()), F: marshalDownloadError},
		{T: externglib.Type(C.webkit_javascript_error_get_type()), F: marshalJavascriptError},
		{T: externglib.Type(C.webkit_network_error_get_type()), F: marshalNetworkError},
		{T: externglib.Type(C.webkit_plugin_error_get_type()), F: marshalPluginError},
		{T: externglib.Type(C.webkit_policy_error_get_type()), F: marshalPolicyError},
		{T: externglib.Type(C.webkit_print_error_get_type()), F: marshalPrintError},
		{T: externglib.Type(C.webkit_snapshot_error_get_type()), F: marshalSnapshotError},
	})
}

// DownloadError: enum values used to denote the various download errors.
type DownloadError int

const (
	// DownloadErrorNetwork: download failure due to network error
	DownloadErrorNetwork DownloadError = 499
	// DownloadErrorCancelledByUser: download was cancelled by user
	DownloadErrorCancelledByUser DownloadError = 400
	// DownloadErrorDestination: download failure due to destination error
	DownloadErrorDestination DownloadError = 401
)

func marshalDownloadError(p uintptr) (interface{}, error) {
	return DownloadError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for DownloadError.
func (d DownloadError) String() string {
	switch d {
	case DownloadErrorNetwork:
		return "Network"
	case DownloadErrorCancelledByUser:
		return "CancelledByUser"
	case DownloadErrorDestination:
		return "Destination"
	default:
		return fmt.Sprintf("DownloadError(%d)", d)
	}
}

// JavascriptError: enum values used to denote errors happening when executing
// JavaScript
type JavascriptError int

const (
	// JavascriptErrorScriptFailed: exception was raised in JavaScript execution
	JavascriptErrorScriptFailed JavascriptError = 699
)

func marshalJavascriptError(p uintptr) (interface{}, error) {
	return JavascriptError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for JavascriptError.
func (j JavascriptError) String() string {
	switch j {
	case JavascriptErrorScriptFailed:
		return "Failed"
	default:
		return fmt.Sprintf("JavascriptError(%d)", j)
	}
}

// NetworkError: enum values used to denote the various network errors.
type NetworkError int

const (
	// NetworkErrorFailed: generic load failure
	NetworkErrorFailed NetworkError = 399
	// NetworkErrorTransport: load failure due to transport error
	NetworkErrorTransport NetworkError = 300
	// NetworkErrorUnknownProtocol: load failure due to unknown protocol
	NetworkErrorUnknownProtocol NetworkError = 301
	// NetworkErrorCancelled: load failure due to cancellation
	NetworkErrorCancelled NetworkError = 302
	// NetworkErrorFileDoesNotExist: load failure due to missing file
	NetworkErrorFileDoesNotExist NetworkError = 303
)

func marshalNetworkError(p uintptr) (interface{}, error) {
	return NetworkError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for NetworkError.
func (n NetworkError) String() string {
	switch n {
	case NetworkErrorFailed:
		return "Failed"
	case NetworkErrorTransport:
		return "Transport"
	case NetworkErrorUnknownProtocol:
		return "UnknownProtocol"
	case NetworkErrorCancelled:
		return "Cancelled"
	case NetworkErrorFileDoesNotExist:
		return "FileDoesNotExist"
	default:
		return fmt.Sprintf("NetworkError(%d)", n)
	}
}

// PluginError: enum values used to denote the various plugin and multimedia
// errors.
type PluginError int

const (
	// PluginErrorFailed: generic plugin load failure. Deprecated 2.32
	PluginErrorFailed PluginError = 299
	// PluginErrorCannotFindPlugin: load failure due to missing plugin.
	// Deprecated 2.32
	PluginErrorCannotFindPlugin PluginError = 200
	// PluginErrorCannotLoadPlugin: load failure due to inability to load
	// plugin. Deprecated 2.32
	PluginErrorCannotLoadPlugin PluginError = 201
	// PluginErrorJavaUnavailable: load failure due to missing Java support that
	// is required to load plugin. Deprecated 2.32
	PluginErrorJavaUnavailable PluginError = 202
	// PluginErrorConnectionCancelled: load failure due to connection
	// cancellation. Deprecated 2.32
	PluginErrorConnectionCancelled PluginError = 203
	// PluginErrorWillHandleLoad: preliminary load failure for media content
	// types. A new load will be started to perform the media load.
	PluginErrorWillHandleLoad PluginError = 204
)

func marshalPluginError(p uintptr) (interface{}, error) {
	return PluginError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PluginError.
func (p PluginError) String() string {
	switch p {
	case PluginErrorFailed:
		return "Failed"
	case PluginErrorCannotFindPlugin:
		return "CannotFindPlugin"
	case PluginErrorCannotLoadPlugin:
		return "CannotLoadPlugin"
	case PluginErrorJavaUnavailable:
		return "JavaUnavailable"
	case PluginErrorConnectionCancelled:
		return "ConnectionCancelled"
	case PluginErrorWillHandleLoad:
		return "WillHandleLoad"
	default:
		return fmt.Sprintf("PluginError(%d)", p)
	}
}

// PolicyError: enum values used to denote the various policy errors.
type PolicyError int

const (
	// PolicyErrorFailed: generic load failure due to policy error
	PolicyErrorFailed PolicyError = 199
	// PolicyErrorCannotShowMIMEType: load failure due to unsupported mime type
	PolicyErrorCannotShowMIMEType PolicyError = 100
	// PolicyErrorCannotShowURI: load failure due to URI that can not be shown
	PolicyErrorCannotShowURI PolicyError = 101
	// PolicyErrorFrameLoadInterruptedByPolicyChange: load failure due to frame
	// load interruption by policy change
	PolicyErrorFrameLoadInterruptedByPolicyChange PolicyError = 102
	// PolicyErrorCannotUseRestrictedPort: load failure due to port restriction
	PolicyErrorCannotUseRestrictedPort PolicyError = 103
)

func marshalPolicyError(p uintptr) (interface{}, error) {
	return PolicyError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PolicyError.
func (p PolicyError) String() string {
	switch p {
	case PolicyErrorFailed:
		return "Failed"
	case PolicyErrorCannotShowMIMEType:
		return "CannotShowMIMEType"
	case PolicyErrorCannotShowURI:
		return "CannotShowURI"
	case PolicyErrorFrameLoadInterruptedByPolicyChange:
		return "FrameLoadInterruptedByPolicyChange"
	case PolicyErrorCannotUseRestrictedPort:
		return "CannotUseRestrictedPort"
	default:
		return fmt.Sprintf("PolicyError(%d)", p)
	}
}

// PrintError: enum values used to denote the various print errors.
type PrintError int

const (
	// PrintErrorGeneral: unspecified error during a print operation
	PrintErrorGeneral PrintError = 599
	// PrintErrorPrinterNotFound: selected printer cannot be found
	PrintErrorPrinterNotFound PrintError = 500
	// PrintErrorInvalidPageRange: invalid page range
	PrintErrorInvalidPageRange PrintError = 501
)

func marshalPrintError(p uintptr) (interface{}, error) {
	return PrintError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PrintError.
func (p PrintError) String() string {
	switch p {
	case PrintErrorGeneral:
		return "General"
	case PrintErrorPrinterNotFound:
		return "PrinterNotFound"
	case PrintErrorInvalidPageRange:
		return "InvalidPageRange"
	default:
		return fmt.Sprintf("PrintError(%d)", p)
	}
}

// SnapshotError: enum values used to denote errors happening when creating
// snapshots of KitWebView
type SnapshotError int

const (
	// SnapshotErrorFailedToCreate: error occurred when creating a webpage
	// snapshot.
	SnapshotErrorFailedToCreate SnapshotError = 799
)

func marshalSnapshotError(p uintptr) (interface{}, error) {
	return SnapshotError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for SnapshotError.
func (s SnapshotError) String() string {
	switch s {
	case SnapshotErrorFailedToCreate:
		return "Create"
	default:
		return fmt.Sprintf("SnapshotError(%d)", s)
	}
}
