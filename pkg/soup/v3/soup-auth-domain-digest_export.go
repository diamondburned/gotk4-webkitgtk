// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

//export _gotk4_soup3_AuthDomainDigestAuthCallback
func _gotk4_soup3_AuthDomainDigestAuthCallback(arg1 *C.SoupAuthDomain, arg2 *C.SoupServerMessage, arg3 *C.char, arg4 C.gpointer) (cret *C.char) {
	var fn AuthDomainDigestAuthCallback
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AuthDomainDigestAuthCallback)
	}

	var _domain *AuthDomainDigest // out
	var _msg *ServerMessage       // out
	var _username string          // out

	_domain = wrapAuthDomainDigest(coreglib.Take(unsafe.Pointer(arg1)))
	_msg = wrapServerMessage(coreglib.Take(unsafe.Pointer(arg2)))
	_username = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	utf8 := fn(_domain, _msg, _username)

	var _ string

	if utf8 != "" {
		cret = (*C.char)(unsafe.Pointer(C.CString(utf8)))
	}

	return cret
}
