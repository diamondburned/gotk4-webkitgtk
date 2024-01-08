// Code generated by girgen. DO NOT EDIT.

package webkit2webextension

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <webkit2/webkit-web-extension.h>
import "C"

// WebHitTestResultClass: instance of this type is always passed by reference.
type WebHitTestResultClass struct {
	*webHitTestResultClass
}

// webHitTestResultClass is the struct that's finalized.
type webHitTestResultClass struct {
	native *C.WebKitWebHitTestResultClass
}

func (w *WebHitTestResultClass) ParentClass() *HitTestResultClass {
	valptr := &w.native.parent_class
	var _v *HitTestResultClass // out
	_v = (*HitTestResultClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
