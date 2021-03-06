// Code generated by girgen. DO NOT EDIT.

package soup

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsoup/soup.h>
// extern gboolean _gotk4_soup2_HSTSEnforcerClass_has_valid_policy(SoupHSTSEnforcer*, char*);
// extern gboolean _gotk4_soup2_HSTSEnforcerClass_is_persistent(SoupHSTSEnforcer*);
// extern void _gotk4_soup2_HSTSEnforcerClass_changed(SoupHSTSEnforcer*, SoupHSTSPolicy*, SoupHSTSPolicy*);
// extern void _gotk4_soup2_HSTSEnforcerClass_hsts_enforced(SoupHSTSEnforcer*, SoupMessage*);
// extern void _gotk4_soup2_HSTSEnforcer_ConnectChanged(gpointer, SoupHSTSPolicy*, SoupHSTSPolicy*, guintptr);
// extern void _gotk4_soup2_HSTSEnforcer_ConnectHstsEnforced(gpointer, SoupMessage*, guintptr);
import "C"

// glib.Type values for soup-hsts-enforcer.go.
var GTypeHSTSEnforcer = externglib.Type(C.soup_hsts_enforcer_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeHSTSEnforcer, F: marshalHSTSEnforcer},
	})
}

// HSTSEnforcerOverrider contains methods that are overridable.
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

func classInitHSTSEnforcerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.SoupHSTSEnforcerClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.SoupHSTSEnforcerClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Changed(oldPolicy, newPolicy *HSTSPolicy)
	}); ok {
		pclass.changed = (*[0]byte)(C._gotk4_soup2_HSTSEnforcerClass_changed)
	}

	if _, ok := goval.(interface{ HasValidPolicy(domain string) bool }); ok {
		pclass.has_valid_policy = (*[0]byte)(C._gotk4_soup2_HSTSEnforcerClass_has_valid_policy)
	}

	if _, ok := goval.(interface{ HstsEnforced(message *Message) }); ok {
		pclass.hsts_enforced = (*[0]byte)(C._gotk4_soup2_HSTSEnforcerClass_hsts_enforced)
	}

	if _, ok := goval.(interface{ IsPersistent() bool }); ok {
		pclass.is_persistent = (*[0]byte)(C._gotk4_soup2_HSTSEnforcerClass_is_persistent)
	}
}

//export _gotk4_soup2_HSTSEnforcerClass_changed
func _gotk4_soup2_HSTSEnforcerClass_changed(arg0 *C.SoupHSTSEnforcer, arg1 *C.SoupHSTSPolicy, arg2 *C.SoupHSTSPolicy) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Changed(oldPolicy, newPolicy *HSTSPolicy)
	})

	var _oldPolicy *HSTSPolicy // out
	var _newPolicy *HSTSPolicy // out

	_oldPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_newPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	iface.Changed(_oldPolicy, _newPolicy)
}

//export _gotk4_soup2_HSTSEnforcerClass_has_valid_policy
func _gotk4_soup2_HSTSEnforcerClass_has_valid_policy(arg0 *C.SoupHSTSEnforcer, arg1 *C.char) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ HasValidPolicy(domain string) bool })

	var _domain string // out

	_domain = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := iface.HasValidPolicy(_domain)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_soup2_HSTSEnforcerClass_hsts_enforced
func _gotk4_soup2_HSTSEnforcerClass_hsts_enforced(arg0 *C.SoupHSTSEnforcer, arg1 *C.SoupMessage) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ HstsEnforced(message *Message) })

	var _message *Message // out

	_message = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))

	iface.HstsEnforced(_message)
}

//export _gotk4_soup2_HSTSEnforcerClass_is_persistent
func _gotk4_soup2_HSTSEnforcerClass_is_persistent(arg0 *C.SoupHSTSEnforcer) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ IsPersistent() bool })

	ok := iface.IsPersistent()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapHSTSEnforcer(obj *externglib.Object) *HSTSEnforcer {
	return &HSTSEnforcer{
		Object: obj,
		SessionFeature: SessionFeature{
			Object: obj,
		},
	}
}

func marshalHSTSEnforcer(p uintptr) (interface{}, error) {
	return wrapHSTSEnforcer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_soup2_HSTSEnforcer_ConnectChanged
func _gotk4_soup2_HSTSEnforcer_ConnectChanged(arg0 C.gpointer, arg1 *C.SoupHSTSPolicy, arg2 *C.SoupHSTSPolicy, arg3 C.guintptr) {
	var f func(oldPolicy, newPolicy *HSTSPolicy)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(oldPolicy, newPolicy *HSTSPolicy))
	}

	var _oldPolicy *HSTSPolicy // out
	var _newPolicy *HSTSPolicy // out

	_oldPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_newPolicy = (*HSTSPolicy)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	f(_oldPolicy, _newPolicy)
}

// ConnectChanged is emitted when hsts_enforcer changes. If a policy has been
// added, new_policy will contain the newly-added policy and old_policy will be
// NULL. If a policy has been deleted, old_policy will contain the to-be-deleted
// policy and new_policy will be NULL. If a policy has been changed, old_policy
// will contain its old value, and new_policy its new value.
//
// Note that you shouldn't modify the policies from a callback to this signal.
func (hstsEnforcer *HSTSEnforcer) ConnectChanged(f func(oldPolicy, newPolicy *HSTSPolicy)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(hstsEnforcer, "changed", false, unsafe.Pointer(C._gotk4_soup2_HSTSEnforcer_ConnectChanged), f)
}

//export _gotk4_soup2_HSTSEnforcer_ConnectHstsEnforced
func _gotk4_soup2_HSTSEnforcer_ConnectHstsEnforced(arg0 C.gpointer, arg1 *C.SoupMessage, arg2 C.guintptr) {
	var f func(message *Message)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(message *Message))
	}

	var _message *Message // out

	_message = wrapMessage(externglib.Take(unsafe.Pointer(arg1)))

	f(_message)
}

// ConnectHstsEnforced is emitted when hsts_enforcer has upgraded the protocol
// for message to HTTPS as a result of matching its domain with a HSTS policy.
func (hstsEnforcer *HSTSEnforcer) ConnectHstsEnforced(f func(message *Message)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(hstsEnforcer, "hsts-enforced", false, unsafe.Pointer(C._gotk4_soup2_HSTSEnforcer_ConnectHstsEnforced), f)
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

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(externglib.InternObject(hstsEnforcer).Native()))
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

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(externglib.InternObject(hstsEnforcer).Native()))
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

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(externglib.InternObject(hstsEnforcer).Native()))
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

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(externglib.InternObject(hstsEnforcer).Native()))

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

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(externglib.InternObject(hstsEnforcer).Native()))
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

	_arg0 = (*C.SoupHSTSEnforcer)(unsafe.Pointer(externglib.InternObject(hstsEnforcer).Native()))
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
