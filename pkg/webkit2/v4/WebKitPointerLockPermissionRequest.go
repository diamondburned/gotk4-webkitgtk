// Code generated by girgen. DO NOT EDIT.

package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"

// PointerLockPermissionRequestClass: instance of this type is always passed by
// reference.
type PointerLockPermissionRequestClass struct {
	*pointerLockPermissionRequestClass
}

// pointerLockPermissionRequestClass is the struct that's finalized.
type pointerLockPermissionRequestClass struct {
	native *C.WebKitPointerLockPermissionRequestClass
}
