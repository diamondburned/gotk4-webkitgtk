// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

// GType values.
var (
	GTypeHSTSPolicy = coreglib.Type(C.soup_hsts_policy_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHSTSPolicy, F: marshalHSTSPolicy},
	})
}

// HSTSPolicy: HTTP Strict Transport Security policy.
//
// domain represents the host that this policy applies to. The domain must be
// IDNA-canonicalized. soup_hsts_policy_new() and related methods will do this
// for you.
//
// max_age contains the 'max-age' value from the Strict Transport Security
// header and indicates the time to live of this policy, in seconds.
//
// expires will be non-NULL if the policy has been set by the host and hence
// has an expiry time. If expires is NULL, it indicates that the policy is a
// permanent session policy set by the user agent.
//
// If include_subdomains is TRUE, the Strict Transport Security policy must also
// be enforced on subdomains of domain.
//
// An instance of this type is always passed by reference.
type HSTSPolicy struct {
	*hstsPolicy
}

// hstsPolicy is the struct that's finalized.
type hstsPolicy struct {
	native *C.SoupHSTSPolicy
}

func marshalHSTSPolicy(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &HSTSPolicy{&hstsPolicy{(*C.SoupHSTSPolicy)(b)}}, nil
}

// NewHSTSPolicy constructs a struct HSTSPolicy.
func NewHSTSPolicy(domain string, maxAge uint32, includeSubdomains bool) *HSTSPolicy {
	var _arg1 *C.char           // out
	var _arg2 C.ulong           // out
	var _arg3 C.gboolean        // out
	var _cret *C.SoupHSTSPolicy // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.ulong(maxAge)
	if includeSubdomains {
		_arg3 = C.TRUE
	}

	_cret = C.soup_hsts_policy_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(maxAge)
	runtime.KeepAlive(includeSubdomains)

	var _hstsPolicy *HSTSPolicy // out

	_hstsPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_hstsPolicy)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_hsts_policy_free((*C.SoupHSTSPolicy)(intern.C))
		},
	)

	return _hstsPolicy
}

// NewHSTSPolicyFromResponse constructs a struct HSTSPolicy.
func NewHSTSPolicyFromResponse(msg *Message) *HSTSPolicy {
	var _arg1 *C.SoupMessage    // out
	var _cret *C.SoupHSTSPolicy // in

	_arg1 = (*C.SoupMessage)(unsafe.Pointer(coreglib.InternObject(msg).Native()))

	_cret = C.soup_hsts_policy_new_from_response(_arg1)
	runtime.KeepAlive(msg)

	var _hstsPolicy *HSTSPolicy // out

	if _cret != nil {
		_hstsPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_hstsPolicy)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_hsts_policy_free((*C.SoupHSTSPolicy)(intern.C))
			},
		)
	}

	return _hstsPolicy
}

// NewHSTSPolicyFull constructs a struct HSTSPolicy.
func NewHSTSPolicyFull(domain string, maxAge uint32, expires *Date, includeSubdomains bool) *HSTSPolicy {
	var _arg1 *C.char           // out
	var _arg2 C.ulong           // out
	var _arg3 *C.SoupDate       // out
	var _arg4 C.gboolean        // out
	var _cret *C.SoupHSTSPolicy // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.ulong(maxAge)
	_arg3 = (*C.SoupDate)(gextras.StructNative(unsafe.Pointer(expires)))
	if includeSubdomains {
		_arg4 = C.TRUE
	}

	_cret = C.soup_hsts_policy_new_full(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(maxAge)
	runtime.KeepAlive(expires)
	runtime.KeepAlive(includeSubdomains)

	var _hstsPolicy *HSTSPolicy // out

	_hstsPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_hstsPolicy)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_hsts_policy_free((*C.SoupHSTSPolicy)(intern.C))
		},
	)

	return _hstsPolicy
}

// NewHSTSPolicySessionPolicy constructs a struct HSTSPolicy.
func NewHSTSPolicySessionPolicy(domain string, includeSubdomains bool) *HSTSPolicy {
	var _arg1 *C.char           // out
	var _arg2 C.gboolean        // out
	var _cret *C.SoupHSTSPolicy // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))
	if includeSubdomains {
		_arg2 = C.TRUE
	}

	_cret = C.soup_hsts_policy_new_session_policy(_arg1, _arg2)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(includeSubdomains)

	var _hstsPolicy *HSTSPolicy // out

	_hstsPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_hstsPolicy)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_hsts_policy_free((*C.SoupHSTSPolicy)(intern.C))
		},
	)

	return _hstsPolicy
}

// MaxAge: maximum age, in seconds, that the policy is valid.
func (h *HSTSPolicy) MaxAge() uint32 {
	valptr := &h.native.max_age
	var _v uint32 // out
	_v = uint32(*valptr)
	return _v
}

// Expires: policy expiration time, or NULL for a permanent session policy.
func (h *HSTSPolicy) Expires() *Date {
	valptr := &h.native.expires
	var _v *Date // out
	_v = (*Date)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// IncludeSubdomains: TRUE if the policy applies on subdomains.
func (h *HSTSPolicy) IncludeSubdomains() bool {
	valptr := &h.native.include_subdomains
	var _v bool // out
	if *valptr != 0 {
		_v = true
	}
	return _v
}

// MaxAge: maximum age, in seconds, that the policy is valid.
func (h *HSTSPolicy) SetMaxAge(maxAge uint32) {
	valptr := &h.native.max_age
	*valptr = C.ulong(maxAge)
}

// IncludeSubdomains: TRUE if the policy applies on subdomains.
func (h *HSTSPolicy) SetIncludeSubdomains(includeSubdomains bool) {
	valptr := &h.native.include_subdomains
	if includeSubdomains {
		*valptr = C.TRUE
	}
}

// Copy copies policy.
//
// The function returns the following values:
//
//   - hstsPolicy: copy of policy.
//
func (policy *HSTSPolicy) Copy() *HSTSPolicy {
	var _arg0 *C.SoupHSTSPolicy // out
	var _cret *C.SoupHSTSPolicy // in

	_arg0 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy)))

	_cret = C.soup_hsts_policy_copy(_arg0)
	runtime.KeepAlive(policy)

	var _hstsPolicy *HSTSPolicy // out

	_hstsPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_hstsPolicy)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.soup_hsts_policy_free((*C.SoupHSTSPolicy)(intern.C))
		},
	)

	return _hstsPolicy
}

// Equal tests if policy1 and policy2 are equal.
//
// The function takes the following parameters:
//
//   - policy2: HSTSPolicy.
//
// The function returns the following values:
//
//   - ok: whether the policies are equal.
//
func (policy1 *HSTSPolicy) Equal(policy2 *HSTSPolicy) bool {
	var _arg0 *C.SoupHSTSPolicy // out
	var _arg1 *C.SoupHSTSPolicy // out
	var _cret C.gboolean        // in

	_arg0 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy1)))
	_arg1 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy2)))

	_cret = C.soup_hsts_policy_equal(_arg0, _arg1)
	runtime.KeepAlive(policy1)
	runtime.KeepAlive(policy2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Domain gets policy's domain.
//
// The function returns the following values:
//
//   - utf8 policy's domain.
//
func (policy *HSTSPolicy) Domain() string {
	var _arg0 *C.SoupHSTSPolicy // out
	var _cret *C.char           // in

	_arg0 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy)))

	_cret = C.soup_hsts_policy_get_domain(_arg0)
	runtime.KeepAlive(policy)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IncludesSubdomains gets whether policy include its subdomains.
//
// The function returns the following values:
//
//   - ok: TRUE if policy includes subdomains, FALSE otherwise.
//
func (policy *HSTSPolicy) IncludesSubdomains() bool {
	var _arg0 *C.SoupHSTSPolicy // out
	var _cret C.gboolean        // in

	_arg0 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy)))

	_cret = C.soup_hsts_policy_includes_subdomains(_arg0)
	runtime.KeepAlive(policy)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsExpired gets whether policy is expired. Permanent policies never expire.
//
// The function returns the following values:
//
//   - ok: TRUE if policy is expired, FALSE otherwise.
//
func (policy *HSTSPolicy) IsExpired() bool {
	var _arg0 *C.SoupHSTSPolicy // out
	var _cret C.gboolean        // in

	_arg0 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy)))

	_cret = C.soup_hsts_policy_is_expired(_arg0)
	runtime.KeepAlive(policy)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSessionPolicy gets whether policy is a non-permanent, non-expirable session
// policy. see soup_hsts_policy_new_session_policy() for details.
//
// The function returns the following values:
//
//   - ok: TRUE if policy is permanent, FALSE otherwise.
//
func (policy *HSTSPolicy) IsSessionPolicy() bool {
	var _arg0 *C.SoupHSTSPolicy // out
	var _cret C.gboolean        // in

	_arg0 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy)))

	_cret = C.soup_hsts_policy_is_session_policy(_arg0)
	runtime.KeepAlive(policy)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
