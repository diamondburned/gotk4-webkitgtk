// Code generated by girgen. DO NOT EDIT.

package webkit2

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit2/webkit2.h>
import "C"

// GType values.
var (
	GTypeITPFirstParty = coreglib.Type(C.webkit_itp_first_party_get_type())
	GTypeITPThirdParty = coreglib.Type(C.webkit_itp_third_party_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeITPFirstParty, F: marshalITPFirstParty},
		coreglib.TypeMarshaler{T: GTypeITPThirdParty, F: marshalITPThirdParty},
	})
}

// ITPFirstParty describes a first party origin.
//
// An instance of this type is always passed by reference.
type ITPFirstParty struct {
	*itpFirstParty
}

// itpFirstParty is the struct that's finalized.
type itpFirstParty struct {
	native *C.WebKitITPFirstParty
}

func marshalITPFirstParty(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ITPFirstParty{&itpFirstParty{(*C.WebKitITPFirstParty)(b)}}, nil
}

// Domain: get the domain name of itp_first_party.
//
// The function returns the following values:
//
//   - utf8: domain name.
//
func (itpFirstParty *ITPFirstParty) Domain() string {
	var _arg0 *C.WebKitITPFirstParty // out
	var _cret *C.char                // in

	_arg0 = (*C.WebKitITPFirstParty)(gextras.StructNative(unsafe.Pointer(itpFirstParty)))

	_cret = C.webkit_itp_first_party_get_domain(_arg0)
	runtime.KeepAlive(itpFirstParty)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LastUpdateTime: get the last time a KitITPThirdParty has been seen under
// itp_first_party.
//
// Each WebKitITPFirstParty is created by
// webkit_itp_third_party_get_first_parties() and therefore corresponds to
// exactly one KitITPThirdParty.
//
// The function returns the following values:
//
//   - dateTime: last update time as a Time.
//
func (itpFirstParty *ITPFirstParty) LastUpdateTime() *glib.DateTime {
	var _arg0 *C.WebKitITPFirstParty // out
	var _cret *C.GDateTime           // in

	_arg0 = (*C.WebKitITPFirstParty)(gextras.StructNative(unsafe.Pointer(itpFirstParty)))

	_cret = C.webkit_itp_first_party_get_last_update_time(_arg0)
	runtime.KeepAlive(itpFirstParty)

	var _dateTime *glib.DateTime // out

	_dateTime = (*glib.DateTime)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_date_time_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dateTime)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_date_time_unref((*C.GDateTime)(intern.C))
		},
	)

	return _dateTime
}

// WebsiteDataAccessAllowed: get whether itp_first_party has granted website
// data access to its KitITPThirdParty.
//
// Each WebKitITPFirstParty is created by
// webkit_itp_third_party_get_first_parties() and therefore corresponds to
// exactly one KitITPThirdParty.
//
// The function returns the following values:
//
//   - ok: TRUE if website data access has been granted, or FALSE otherwise.
//
func (itpFirstParty *ITPFirstParty) WebsiteDataAccessAllowed() bool {
	var _arg0 *C.WebKitITPFirstParty // out
	var _cret C.gboolean             // in

	_arg0 = (*C.WebKitITPFirstParty)(gextras.StructNative(unsafe.Pointer(itpFirstParty)))

	_cret = C.webkit_itp_first_party_get_website_data_access_allowed(_arg0)
	runtime.KeepAlive(itpFirstParty)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ITPThirdParty describes a third party origin.
//
// An instance of this type is always passed by reference.
type ITPThirdParty struct {
	*itpThirdParty
}

// itpThirdParty is the struct that's finalized.
type itpThirdParty struct {
	native *C.WebKitITPThirdParty
}

func marshalITPThirdParty(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ITPThirdParty{&itpThirdParty{(*C.WebKitITPThirdParty)(b)}}, nil
}

// Domain: get the domain name of itp_third_party.
//
// The function returns the following values:
//
//   - utf8: domain name.
//
func (itpThirdParty *ITPThirdParty) Domain() string {
	var _arg0 *C.WebKitITPThirdParty // out
	var _cret *C.char                // in

	_arg0 = (*C.WebKitITPThirdParty)(gextras.StructNative(unsafe.Pointer(itpThirdParty)))

	_cret = C.webkit_itp_third_party_get_domain(_arg0)
	runtime.KeepAlive(itpThirdParty)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// FirstParties: get the list of KitITPFirstParty under which itp_third_party
// has been seen.
//
// The function returns the following values:
//
//   - list of KitITPFirstParty.
//
func (itpThirdParty *ITPThirdParty) FirstParties() []*ITPFirstParty {
	var _arg0 *C.WebKitITPThirdParty // out
	var _cret *C.GList               // in

	_arg0 = (*C.WebKitITPThirdParty)(gextras.StructNative(unsafe.Pointer(itpThirdParty)))

	_cret = C.webkit_itp_third_party_get_first_parties(_arg0)
	runtime.KeepAlive(itpThirdParty)

	var _list []*ITPFirstParty // out

	_list = make([]*ITPFirstParty, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.WebKitITPFirstParty)(v)
		var dst *ITPFirstParty // out
		dst = (*ITPFirstParty)(gextras.NewStructNative(unsafe.Pointer(src)))
		C.webkit_itp_first_party_ref(src)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.webkit_itp_first_party_unref((*C.WebKitITPFirstParty)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}
