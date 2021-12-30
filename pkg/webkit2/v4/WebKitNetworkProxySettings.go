// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.webkit_network_proxy_mode_get_type()), F: marshalNetworkProxyMode},
		{T: externglib.Type(C.webkit_network_proxy_settings_get_type()), F: marshalNetworkProxySettings},
	})
}

// NetworkProxyMode: enum values used to set the network proxy mode.
type NetworkProxyMode C.gint

const (
	// NetworkProxyModeDefault: use the default proxy of the system.
	NetworkProxyModeDefault NetworkProxyMode = iota
	// NetworkProxyModeNoProxy: do not use any proxy.
	NetworkProxyModeNoProxy
	// NetworkProxyModeCustom: use custom proxy settings.
	NetworkProxyModeCustom
)

func marshalNetworkProxyMode(p uintptr) (interface{}, error) {
	return NetworkProxyMode(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for NetworkProxyMode.
func (n NetworkProxyMode) String() string {
	switch n {
	case NetworkProxyModeDefault:
		return "Default"
	case NetworkProxyModeNoProxy:
		return "NoProxy"
	case NetworkProxyModeCustom:
		return "Custom"
	default:
		return fmt.Sprintf("NetworkProxyMode(%d)", n)
	}
}

// NetworkProxySettings: instance of this type is always passed by reference.
type NetworkProxySettings struct {
	*networkProxySettings
}

// networkProxySettings is the struct that's finalized.
type networkProxySettings struct {
	native *C.WebKitNetworkProxySettings
}

func marshalNetworkProxySettings(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &NetworkProxySettings{&networkProxySettings{(*C.WebKitNetworkProxySettings)(b)}}, nil
}

// NewNetworkProxySettings constructs a struct NetworkProxySettings.
func NewNetworkProxySettings(defaultProxyUri string, ignoreHosts []string) *NetworkProxySettings {
	var _arg1 *C.gchar                      // out
	var _arg2 **C.gchar                     // out
	var _cret *C.WebKitNetworkProxySettings // in

	if defaultProxyUri != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(defaultProxyUri)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	{
		_arg2 = (**C.gchar)(C.calloc(C.size_t((len(ignoreHosts) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(ignoreHosts)+1)
			var zero *C.gchar
			out[len(ignoreHosts)] = zero
			for i := range ignoreHosts {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(ignoreHosts[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.webkit_network_proxy_settings_new(_arg1, _arg2)
	runtime.KeepAlive(defaultProxyUri)
	runtime.KeepAlive(ignoreHosts)

	var _networkProxySettings *NetworkProxySettings // out

	_networkProxySettings = (*NetworkProxySettings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_networkProxySettings)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_network_proxy_settings_free((*C.WebKitNetworkProxySettings)(intern.C))
		},
	)

	return _networkProxySettings
}

// AddProxyForScheme adds a URI-scheme-specific proxy. URIs whose scheme matches
// uri_scheme will be proxied via proxy_uri. As with the default proxy URI, if
// proxy_uri starts with "socks://", it will be treated as referring to all
// three of the socks5, socks4a, and socks4 proxy types.
//
// The function takes the following parameters:
//
//    - scheme: URI scheme to add a proxy for.
//    - proxyUri: proxy URI to use for uri_scheme.
//
func (proxySettings *NetworkProxySettings) AddProxyForScheme(scheme string, proxyUri string) {
	var _arg0 *C.WebKitNetworkProxySettings // out
	var _arg1 *C.gchar                      // out
	var _arg2 *C.gchar                      // out

	_arg0 = (*C.WebKitNetworkProxySettings)(gextras.StructNative(unsafe.Pointer(proxySettings)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(proxyUri)))
	defer C.free(unsafe.Pointer(_arg2))

	C.webkit_network_proxy_settings_add_proxy_for_scheme(_arg0, _arg1, _arg2)
	runtime.KeepAlive(proxySettings)
	runtime.KeepAlive(scheme)
	runtime.KeepAlive(proxyUri)
}

// Copy: make a copy of the KitNetworkProxySettings.
//
// The function returns the following values:
//
//    - networkProxySettings: copy of passed in KitNetworkProxySettings.
//
func (proxySettings *NetworkProxySettings) Copy() *NetworkProxySettings {
	var _arg0 *C.WebKitNetworkProxySettings // out
	var _cret *C.WebKitNetworkProxySettings // in

	_arg0 = (*C.WebKitNetworkProxySettings)(gextras.StructNative(unsafe.Pointer(proxySettings)))

	_cret = C.webkit_network_proxy_settings_copy(_arg0)
	runtime.KeepAlive(proxySettings)

	var _networkProxySettings *NetworkProxySettings // out

	_networkProxySettings = (*NetworkProxySettings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_networkProxySettings)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.webkit_network_proxy_settings_free((*C.WebKitNetworkProxySettings)(intern.C))
		},
	)

	return _networkProxySettings
}
