package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const webkitModule = "github.com/diamondburned/gotk4-webkitgtk/pkg"

var data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: webkitModule,
		Packages: []genmain.Package{
			{Name: "libsoup-3.0"},
			{Name: "webkit2gtk-4.1"},
			{Name: "webkitgtk-6.0"},
		},
		PkgGenerated: []string{
			"javascriptcore",
			"soup",
			"webkit",
			"webkit2",
			"webkit2webextension",
			"webkitwebprocessextension",
		},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
		},
		Filters: []types.FilterMatcher{
			// This requires some weird C import.
			types.AbsoluteFilter("C.sockaddr"),
			// These aren't found as well for some reason.
			types.AbsoluteFilter("Soup.PasswordManager"),
			types.AbsoluteFilter("Soup.Requester"),
			types.AbsoluteFilter("Soup.get_resource"),
			types.AbsoluteFilter("C.soup_requester_error_quark"),
			// gio.TLSProtocolVersion is too new.
			types.AbsoluteFilter("Soup-3.Message.get_tls_protocol_version"),
			// webkit.h only includes this header when ENABLE_POINTER_LOCK is
			// defined, which it never is for consumers of the library.
			types.AbsoluteFilter("WebKit-6.PointerLockPermissionRequest"),
			types.AbsoluteFilter("WebKit-6.PointerLockPermissionRequestClass"),
			types.AbsoluteFilter("WebKit2-4.PointerLockPermissionRequest"),
			types.AbsoluteFilter("WebKit2-4.PointerLockPermissionRequestClass"),
			// Deprecated in favor of the lowercase variant, which has the same
			// Go name.
			types.AbsoluteFilter("C.webkit_web_extension_match_pattern_register_custom_URL_scheme"),
			// The generated callback passes text as []byte while the override
			// takes a string, since the docs never call text UTF-8.
			types.AbsoluteFilter("WebKit-6.InputMethodContext.notify_surrounding"),
			types.AbsoluteFilter("WebKit2-4.InputMethodContext.notify_surrounding"),
		},
		Postprocessors: map[string][]girgen.Postprocessor{
			"WebKitWebProcessExtension-6": {
				func(n *girgen.NamespaceGenerator) error {
					// Fix the name colliding with an all-lowercase one.
					n.Rename("WebKitWebProcessExtension.go", "WebKitWebProcessExtension2.go")
					return nil
				},
			},
		},
	},
)

func main() {
	genmain.Run(data)
}
