// Code generated by girgen. DO NOT EDIT.

package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"

// WebsiteDataAccessPermissionRequestClass: instance of this type is always
// passed by reference.
type WebsiteDataAccessPermissionRequestClass struct {
	*websiteDataAccessPermissionRequestClass
}

// websiteDataAccessPermissionRequestClass is the struct that's finalized.
type websiteDataAccessPermissionRequestClass struct {
	native *C.WebKitWebsiteDataAccessPermissionRequestClass
}
