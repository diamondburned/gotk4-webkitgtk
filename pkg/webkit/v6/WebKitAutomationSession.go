// Code generated by girgen. DO NOT EDIT.

package webkit

// #include <stdlib.h>
// #include <webkit/webkit.h>
import "C"

// AutomationSessionClass: instance of this type is always passed by reference.
type AutomationSessionClass struct {
	*automationSessionClass
}

// automationSessionClass is the struct that's finalized.
type automationSessionClass struct {
	native *C.WebKitAutomationSessionClass
}
