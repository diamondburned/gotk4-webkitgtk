// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.soup_hsts_enforcer_get_type()), F: marshalHSTSEnforcerer},
	})
}

// HSTSEnforcerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type HSTSEnforcerOverrider interface {
	// The function takes the following parameters:
	//
	//    - oldPolicy
	//    - newPolicy
	//
	Changed(oldPolicy, newPolicy *HSTSPolicy)
	// HasValidPolicy gets whether hsts_enforcer has a currently valid policy
	// for domain.
	//
	// The function takes the following parameters:
	//
	//    - domain: domain.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if access to domain should happen over HTTPS, false
	//      otherwise.
	//
	HasValidPolicy(domain string) bool
	// The function takes the following parameters:
	//
	HstsEnforced(message *Message)
	// IsPersistent gets whether hsts_enforcer stores policies persistenly.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if hsts_enforcer storage is persistent or FALSE otherwise.
	//
	IsPersistent() bool
}

type HSTSEnforcer struct {
	_ [0]func() // equal guard
	*externglib.Object

	SessionFeature
}

var (
	_ externglib.Objector = (*HSTSEnforcer)(nil)
)

func wrapHSTSEnforcer(obj *externglib.Object) *HSTSEnforcer {
	return &HSTSEnforcer{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalHSTSEnforcerer(p uintptr) (interface{}, error) {
	return wrapHSTSEnforcer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged: emitted when hsts_enforcer changes. If a policy has been
// added, new_policy will contain the newly-added policy and old_policy will be
// NULL. If a policy has been deleted, old_policy will contain the to-be-deleted
// policy and new_policy will be NULL. If a policy has been changed, old_policy
// will contain its old value, and new_policy its new value.
//
// Note that you shouldn't modify the policies from a callback to this signal.
func (hstsEnforcer *HSTSEnforcer) ConnectChanged(f func(oldPolicy, newPolicy *HSTSPolicy)) externglib.SignalHandle {
	return hstsEnforcer.Connect("changed", f)
}

// ConnectHstsEnforced: emitted when hsts_enforcer has upgraded the protocol for
// message to HTTPS as a result of matching its domain with a HSTS policy.
func (hstsEnforcer *HSTSEnforcer) ConnectHstsEnforced(f func(message Message)) externglib.SignalHandle {
	return hstsEnforcer.Connect("hsts-enforced", f)
}

// NewHSTSEnforcer creates a new HSTSEnforcer. The base HSTSEnforcer class does
// not support persistent storage of HSTS policies, see HSTSEnforcerDB for that.
//
// The function returns the following values:
//
//    - hstsEnforcer: new HSTSEnforcer.
//
func NewHSTSEnforcer() *HSTSEnforcer {
	var _cret *C.SoupHSTSEnforcer // in

	_cret = C.soup_hsts_enforcer_new()

	var _hstsEnforcer *HSTSEnforcer // out

	_hstsEnforcer = wrapHSTSEnforcer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _hstsEnforcer
}

// Domains gets a list of domains for which there are policies in enforcer.
//
// The function takes the following parameters:
//
//    - sessionPolicies: whether to include session policies.
//
// The function returns the following values:
//
//    - list: newly allocated list of domains. Use g_list_free_full() and
//      g_free() to free the list.
//
func (hstsEnforcer *HSTSEnforcer) Domains(sessionPolicies bool) []string {
	var _arg0 *C.SoupHSTSEnforcer // out
	var _arg1 C.gboolean          // out
	var _cret *C.GList            // in

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(hstsEnforcer.Native()))
	if sessionPolicies {
		_arg1 = C.TRUE
	}

	_cret = C.soup_hsts_enforcer_get_domains(_arg0, _arg1)
	runtime.KeepAlive(hstsEnforcer)
	runtime.KeepAlive(sessionPolicies)

	var _list []string // out

	_list = make([]string, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(src)))
		defer C.free(unsafe.Pointer(src))
		_list = append(_list, dst)
	})

	return _list
}

// Policies gets a list with the policies in enforcer.
//
// The function takes the following parameters:
//
//    - sessionPolicies: whether to include session policies.
//
// The function returns the following values:
//
//    - list: newly allocated list of policies. Use g_list_free_full() and
//      soup_hsts_policy_free() to free the list.
//
func (hstsEnforcer *HSTSEnforcer) Policies(sessionPolicies bool) []*HSTSPolicy {
	var _arg0 *C.SoupHSTSEnforcer // out
	var _arg1 C.gboolean          // out
	var _cret *C.GList            // in

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(hstsEnforcer.Native()))
	if sessionPolicies {
		_arg1 = C.TRUE
	}

	_cret = C.soup_hsts_enforcer_get_policies(_arg0, _arg1)
	runtime.KeepAlive(hstsEnforcer)
	runtime.KeepAlive(sessionPolicies)

	var _list []*HSTSPolicy // out

	_list = make([]*HSTSPolicy, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.SoupHSTSPolicy)(v)
		var dst *HSTSPolicy // out
		dst = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.soup_hsts_policy_free((*C.SoupHSTSPolicy)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}

// HasValidPolicy gets whether hsts_enforcer has a currently valid policy for
// domain.
//
// The function takes the following parameters:
//
//    - domain: domain.
//
// The function returns the following values:
//
//    - ok: TRUE if access to domain should happen over HTTPS, false otherwise.
//
func (hstsEnforcer *HSTSEnforcer) HasValidPolicy(domain string) bool {
	var _arg0 *C.SoupHSTSEnforcer // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(hstsEnforcer.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.soup_hsts_enforcer_has_valid_policy(_arg0, _arg1)
	runtime.KeepAlive(hstsEnforcer)
	runtime.KeepAlive(domain)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsPersistent gets whether hsts_enforcer stores policies persistenly.
//
// The function returns the following values:
//
//    - ok: TRUE if hsts_enforcer storage is persistent or FALSE otherwise.
//
func (hstsEnforcer *HSTSEnforcer) IsPersistent() bool {
	var _arg0 *C.SoupHSTSEnforcer // out
	var _cret C.gboolean          // in

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(hstsEnforcer.Native()))

	_cret = C.soup_hsts_enforcer_is_persistent(_arg0)
	runtime.KeepAlive(hstsEnforcer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPolicy sets policy to hsts_enforcer. If policy is expired, any existing
// HSTS policy for its host will be removed instead. If a policy existed for
// this host, it will be replaced. Otherwise, the new policy will be inserted.
// If the policy is a session policy, that is, one created with
// soup_hsts_policy_new_session_policy(), the policy will not expire and will be
// enforced during the lifetime of hsts_enforcer's Session.
//
// The function takes the following parameters:
//
//    - policy of the HSTS host.
//
func (hstsEnforcer *HSTSEnforcer) SetPolicy(policy *HSTSPolicy) {
	var _arg0 *C.SoupHSTSEnforcer // out
	var _arg1 *C.SoupHSTSPolicy   // out

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(hstsEnforcer.Native()))
	_arg1 = (*C.SoupHSTSPolicy)(gextras.StructNative(unsafe.Pointer(policy)))

	C.soup_hsts_enforcer_set_policy(_arg0, _arg1)
	runtime.KeepAlive(hstsEnforcer)
	runtime.KeepAlive(policy)
}

// SetSessionPolicy sets a session policy for domain. A session policy is a
// policy that is permanent to the lifetime of hsts_enforcer's Session and
// doesn't expire.
//
// The function takes the following parameters:
//
//    - domain: policy domain or hostname.
//    - includeSubdomains: TRUE if the policy applies on sub domains.
//
func (hstsEnforcer *HSTSEnforcer) SetSessionPolicy(domain string, includeSubdomains bool) {
	var _arg0 *C.SoupHSTSEnforcer // out
	var _arg1 *C.char             // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(hstsEnforcer.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg1))
	if includeSubdomains {
		_arg2 = C.TRUE
	}

	C.soup_hsts_enforcer_set_session_policy(_arg0, _arg1, _arg2)
	runtime.KeepAlive(hstsEnforcer)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(includeSubdomains)
}
