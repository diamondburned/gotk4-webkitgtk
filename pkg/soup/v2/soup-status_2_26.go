// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
)

// #include <stdlib.h>
// #include <libsoup/soup.h>
import "C"

// StatusProxify turns SOUP_STATUS_CANT_RESOLVE into
// SOUP_STATUS_CANT_RESOLVE_PROXY and SOUP_STATUS_CANT_CONNECT into
// SOUP_STATUS_CANT_CONNECT_PROXY. Other status codes are passed through
// unchanged.
//
// The function takes the following parameters:
//
//   - statusCode status code.
//
// The function returns the following values:
//
//   - guint: "proxified" equivalent of status_code.
//
func StatusProxify(statusCode uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(statusCode)

	_cret = C.soup_status_proxify(_arg1)
	runtime.KeepAlive(statusCode)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}
