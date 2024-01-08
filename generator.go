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
			{Name: "libsoup-2.4"},
			{Name: "libsoup-3.0"},
			{Name: "webkit2gtk-4.1"},
			{Name: "webkitgtk-6.0"},
		},
		PkgGenerated: []string{
			"javascriptcore",
			"soup",
			"soupgnome",
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
			// Not sure what this type is, but it's not used.
			types.AbsoluteFilter("Soup-2.PasswordManagerInterface"),
			types.AbsoluteFilter("Soup-2.RequesterClass"),
		},
		Postprocessors: map[string][]girgen.Postprocessor{
			"WebKitWebProcessExtension-6": {
				func(n *girgen.NamespaceGenerator) error {
					// Fix the name colliding with an all-lowercase one.
					n.Rename("WebKitWebProcessExtension.go", "WebKitWebProcessExtension2.go")
					return nil
				},
			},
			"Soup-2": {
				func(n *girgen.NamespaceGenerator) error {
					f := n.Files["soup-message.go"]
					h := f.Header()
					// This import is not needed. Must be a gotk4 bug.
					delete(h.Imports, "github.com/diamondburned/gotk4/pkg/glib/v2")
					return nil
				},
			},
		},
	},
)

func main() {
	genmain.Run(data)
}
