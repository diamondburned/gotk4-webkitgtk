// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_hsts_enforcer_db_get_type()), F: marshalHSTSEnforcerDBer},
	})
}

const HSTS_ENFORCER_DB_FILENAME = "filename"

type HSTSEnforcerDB struct {
	_ [0]func() // equal guard
	HSTSEnforcer
}

var (
	_ externglib.Objector = (*HSTSEnforcerDB)(nil)
)

func wrapHSTSEnforcerDB(obj *externglib.Object) *HSTSEnforcerDB {
	return &HSTSEnforcerDB{
		HSTSEnforcer: HSTSEnforcer{
			Object: obj,
			SessionFeature: SessionFeature{
				Object: obj,
			},
		},
	}
}

func marshalHSTSEnforcerDBer(p uintptr) (interface{}, error) {
	return wrapHSTSEnforcerDB(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHSTSEnforcerDB creates a HSTSEnforcerDB.
//
// filename will be read in during the initialization of a HSTSEnforcerDB, in
// order to create an initial set of HSTS policies. If the file doesn't exist, a
// new database will be created and initialized. Changes to the policies during
// the lifetime of a HSTSEnforcerDB will be written to filename when
// HSTSEnforcer::changed is emitted.
//
// The function takes the following parameters:
//
//    - filename of the database to read/write from.
//
// The function returns the following values:
//
//    - hstsEnforcerDB: new HSTSEnforcer.
//
func NewHSTSEnforcerDB(filename string) *HSTSEnforcerDB {
	var _arg1 *C.char             // out
	var _cret *C.SoupHSTSEnforcer // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_hsts_enforcer_db_new(_arg1)
	runtime.KeepAlive(filename)

	var _hstsEnforcerDB *HSTSEnforcerDB // out

	_hstsEnforcerDB = wrapHSTSEnforcerDB(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _hstsEnforcerDB
}
