// Code generated by girgen. DO NOT EDIT.

package soupgnome

import (
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup-gnome.h>
import "C"

// The function returns the following values:
//
func GnomeFeatures226_GetType() externglib.Type {
	var _cret C.GType // in

	_cret = C.soup_gnome_features_2_26_get_type()

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}
