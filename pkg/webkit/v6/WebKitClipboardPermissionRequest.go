// Code generated by girgen. DO NOT EDIT.

package webkit

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

// ClipboardPermissionRequestClass: instance of this type is always passed by
// reference.
type ClipboardPermissionRequestClass struct {
	*clipboardPermissionRequestClass
}

// clipboardPermissionRequestClass is the struct that's finalized.
type clipboardPermissionRequestClass struct {
	native *C.WebKitClipboardPermissionRequestClass
}
