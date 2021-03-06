// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern SoupCacheability _gotk4_soup2_CacheClass_get_cacheability(SoupCache*, SoupMessage*);
import "C"

// glib.Type values for soup-cache.go.
var (
	GTypeCacheType = externglib.Type(C.soup_cache_type_get_type())
	GTypeCache     = externglib.Type(C.soup_cache_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCacheType, F: marshalCacheType},
		{T: GTypeCache, F: marshalCache},
	})
}

// CacheType: type of cache; this affects what kinds of responses will be saved.
type CacheType C.gint

const (
	// CacheSingleUser: single-user cache.
	CacheSingleUser CacheType = iota
	// CacheShared: shared cache.
	CacheShared
)

func marshalCacheType(p uintptr) (interface{}, error) {
	return CacheType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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
type CacheOverrider interface {
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Cacheability(msg *Message) Cacheability
}

type Cache struct {
	_ [0]func() // equal guard
	*externglib.Object

	SessionFeature
}

var (
	_ externglib.Objector = (*Cache)(nil)
)

func classInitCacher(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.SoupCacheClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.SoupCacheClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Cacheability(msg *Message) Cacheability
	}); ok {
		pclass.get_cacheability = (*[0]byte)(C._gotk4_soup2_CacheClass_get_cacheability)
	}
}

//export _gotk4_soup2_CacheClass_get_cacheability
func _gotk4_soup2_CacheClass_get_cacheability(arg0 *C.SoupCache, arg1 *C.SoupMessage) (cret C.SoupCacheability) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Cacheability(msg *Message) Cacheability
	})

	var _msg *Message // out

	_msg = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))

	cacheability := iface.Cacheability(_msg)

	cret = C.SoupCacheability(cacheability)

	return cret
}

func wrapCache(obj *externglib.Object) *Cache {
	return &Cache{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalCache(p uintptr) (interface{}, error) {
	return wrapCache(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCache creates a new Cache.
//
// The function takes the following parameters:
//
//    - cacheDir (optional): directory to store the cached data, or NULL to use
//      the default one. Note that since the cache isn't safe to access for
//      multiple processes at once, and the default directory isn't namespaced by
//      process, clients are strongly discouraged from passing NULL.
//    - cacheType of the cache.
//
// The function returns the following values:
//
//    - cache: new Cache.
//
func NewCache(cacheDir string, cacheType CacheType) *Cache {
	var _arg1 *C.char         // out
	var _arg2 C.SoupCacheType // out
	var _cret *C.SoupCache    // in

	if cacheDir != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(cacheDir)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.SoupCacheType(cacheType)

	_cret = C.soup_cache_new(_arg1, _arg2)
	runtime.KeepAlive(cacheDir)
	runtime.KeepAlive(cacheType)

	var _cache *Cache // out

	_cache = wrapCache(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cache
}

// Clear will remove all entries in the cache plus all the cache files.
func (cache *Cache) Clear() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(externglib.InternObject(cache).Native()))

	C.soup_cache_clear(_arg0)
	runtime.KeepAlive(cache)
}

// Dump: synchronously writes the cache index out to disk. Contrast with
// soup_cache_flush(), which writes pending cache <emphasis>entries</emphasis>
// to disk.
//
// You must call this before exiting if you want your cache data to persist
// between sessions.
func (cache *Cache) Dump() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(externglib.InternObject(cache).Native()))

	C.soup_cache_dump(_arg0)
	runtime.KeepAlive(cache)
}

// Flush: this function will force all pending writes in the cache to be
// committed to disk. For doing so it will iterate the Context associated with
// cache's session as long as needed.
//
// Contrast with soup_cache_dump(), which writes out the cache index file.
func (cache *Cache) Flush() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(externglib.InternObject(cache).Native()))

	C.soup_cache_flush(_arg0)
	runtime.KeepAlive(cache)
}

// MaxSize gets the maximum size of the cache.
//
// The function returns the following values:
//
//    - guint: maximum size of the cache, in bytes.
//
func (cache *Cache) MaxSize() uint {
	var _arg0 *C.SoupCache // out
	var _cret C.guint      // in

	_arg0 = (*C.SoupCache)(unsafe.Pointer(externglib.InternObject(cache).Native()))

	_cret = C.soup_cache_get_max_size(_arg0)
	runtime.KeepAlive(cache)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Load loads the contents of cache's index into memory.
func (cache *Cache) Load() {
	var _arg0 *C.SoupCache // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(externglib.InternObject(cache).Native()))

	C.soup_cache_load(_arg0)
	runtime.KeepAlive(cache)
}

// SetMaxSize sets the maximum size of the cache.
//
// The function takes the following parameters:
//
//    - maxSize: maximum size of the cache, in bytes.
//
func (cache *Cache) SetMaxSize(maxSize uint) {
	var _arg0 *C.SoupCache // out
	var _arg1 C.guint      // out

	_arg0 = (*C.SoupCache)(unsafe.Pointer(externglib.InternObject(cache).Native()))
	_arg1 = C.guint(maxSize)

	C.soup_cache_set_max_size(_arg0, _arg1)
	runtime.KeepAlive(cache)
	runtime.KeepAlive(maxSize)
}
