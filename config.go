package main

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module     = "github.com/diamondburned/gotk4/pkg"
	webkitgtkModule = "github.com/diamondburned/gotk4-webkitgtk/pkg"
)

var packages = []gendata.Package{
	{PkgName: "libsoup-2.4", Namespaces: []string{"Soup-2"}},
	{PkgName: "libsoup-gnome-2.4", Namespaces: []string{"SoupGNOME-2"}},
	{PkgName: "webkit2gtk-4.0", Namespaces: []string{"JavaScriptCore-4", "WebKit2-4"}},
}

var pkgGenerated = []string{
	"javascriptcore",
	"soup",
	"soupgnome",
	"webkit2",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
}

var preprocessors = []types.Preprocessor{
	types.PreprocessorFunc(func(repos gir.Repositories) {
		res := repos.FindFullType("Soup-2.Auth")
		// Remove all virtual methods, since .Update() is invalid.
		res.Type.(*gir.Class).VirtualMethods = nil
	}),
}

var filters = []types.FilterMatcher{
	// This requires some weird C import.
	types.AbsoluteFilter("C.sockaddr"),
	// These aren't found as well for some reason.
	types.AbsoluteFilter("Soup.PasswordManager"),
	types.AbsoluteFilter("Soup.Requester"),
	types.AbsoluteFilter("Soup.get_resource"),
	types.AbsoluteFilter("C.soup_requester_error_quark"),
}
