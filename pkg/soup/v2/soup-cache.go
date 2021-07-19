// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libsoup-2.4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_cache_type_get_type()), F: marshalCacheType},
		{T: externglib.Type(C.soup_cache_get_type()), F: marshalCacher},
	})
}

// CacheType: type of cache; this affects what kinds of responses will be saved.
type CacheType int

const (
	// CacheSingleUser: single-user cache
	CacheSingleUser CacheType = iota
	// CacheShared cache
	CacheShared
)

func marshalCacheType(p uintptr) (interface{}, error) {
	return CacheType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for CacheType.
func (c CacheType) String() string {
	switch c {
	case CacheSingleUser:
		return "SingleUser"
	case CacheShared:
		return "Shared"
	default:
		return fmt.Sprintf("CacheType(%d)", c)
	}
}

// CacheOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CacheOverrider interface {
	Cacheability(msg *Message) Cacheability
}

type Cache struct {
	*externglib.Object

	SessionFeature
}

var _ gextras.Nativer = (*Cache)(nil)

func wrapCache(obj *externglib.Object) *Cache {
	return &Cache{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalCacher(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCache(obj), nil
}

// NewCache creates a new Cache.
func NewCache(cacheDir string, cacheType CacheType) *Cache {
	var _arg1 *C.char         // out
	var _arg2 C.SoupCacheType // out
	var _cret *C.SoupCache    // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(cacheDir)))
	_arg2 = C.SoupCacheType(cacheType)

	_cret = C.soup_cache_new(_arg1, _arg2)

	var _cache *Cache // out

	_cache = wrapCache(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cache
}

// Clear will remove all entries in the cache plus all the cache files.
func (cache *Cache) Clear() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(cache.Native()))

	C.soup_cache_clear(_arg0)
}

// Dump: synchronously writes the cache index out to disk. Contrast with
// soup_cache_flush(), which writes pending cache <emphasis>entries</emphasis>
// to disk.
//
// You must call this before exiting if you want your cache data to persist
// between sessions.
func (cache *Cache) Dump() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(cache.Native()))

	C.soup_cache_dump(_arg0)
}

// Flush: this function will force all pending writes in the cache to be
// committed to disk. For doing so it will iterate the Context associated with
// cache's session as long as needed.
//
// Contrast with soup_cache_dump(), which writes out the cache index file.
func (cache *Cache) Flush() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(cache.Native()))

	C.soup_cache_flush(_arg0)
}

// MaxSize gets the maximum size of the cache.
func (cache *Cache) MaxSize() uint {
	var _arg0 *C.SoupCache // out
	var _cret C.guint      // in

	_arg0 = (*C.SoupCache)(unsafe.Pointer(cache.Native()))

	_cret = C.soup_cache_get_max_size(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Load loads the contents of cache's index into memory.
func (cache *Cache) Load() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(cache.Native()))

	C.soup_cache_load(_arg0)
}

// SetMaxSize sets the maximum size of the cache.
func (cache *Cache) SetMaxSize(maxSize uint) {
	var _arg0 *C.SoupCache // out
	var _arg1 C.guint      // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(cache.Native()))
	_arg1 = C.guint(maxSize)

	C.soup_cache_set_max_size(_arg0, _arg1)
}
