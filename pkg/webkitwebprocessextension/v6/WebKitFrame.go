// Code generated by girgen. DO NOT EDIT.

package webkitwebprocessextension

// #include <stdlib.h>
// #include <webkit/webkit-web-process-extension.h>
import "C"

// FrameClass: instance of this type is always passed by reference.
type FrameClass struct {
	*frameClass
}

// frameClass is the struct that's finalized.
type frameClass struct {
	native *C.WebKitFrameClass
}
