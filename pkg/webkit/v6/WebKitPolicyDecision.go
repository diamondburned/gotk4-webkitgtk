// Code generated by girgen. DO NOT EDIT.

package webkit

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <webkit/webkit.h>
import "C"

// GType values.
var (
	GTypePolicyDecision = coreglib.Type(C.webkit_policy_decision_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePolicyDecision, F: marshalPolicyDecision},
	})
}

// PolicyDecisionOverrides contains methods that are overridable.
type PolicyDecisionOverrides struct {
}

func defaultPolicyDecisionOverrides(v *PolicyDecision) PolicyDecisionOverrides {
	return PolicyDecisionOverrides{}
}

// PolicyDecision: pending policy decision.
//
// Often WebKit allows the client to decide the policy for certain operations.
// For instance, a client may want to open a link in a new tab, block a
// navigation entirely, query the user or trigger a download instead of a
// navigation. In these cases WebKit will fire the KitWebView::decide-policy
// signal with a KitPolicyDecision object. If the signal handler does nothing,
// WebKit will act as if webkit_policy_decision_use() was called as soon as
// signal handling completes. To make a policy decision asynchronously, simply
// increment the reference count of the KitPolicyDecision object.
type PolicyDecision struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PolicyDecision)(nil)
)

// PolicyDecisioner describes types inherited from class PolicyDecision.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type PolicyDecisioner interface {
	coreglib.Objector
	basePolicyDecision() *PolicyDecision
}

var _ PolicyDecisioner = (*PolicyDecision)(nil)

func init() {
	coreglib.RegisterClassInfo[*PolicyDecision, *PolicyDecisionClass, PolicyDecisionOverrides](
		GTypePolicyDecision,
		initPolicyDecisionClass,
		wrapPolicyDecision,
		defaultPolicyDecisionOverrides,
	)
}

func initPolicyDecisionClass(gclass unsafe.Pointer, overrides PolicyDecisionOverrides, classInitFunc func(*PolicyDecisionClass)) {
	if classInitFunc != nil {
		class := (*PolicyDecisionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPolicyDecision(obj *coreglib.Object) *PolicyDecision {
	return &PolicyDecision{
		Object: obj,
	}
}

func marshalPolicyDecision(p uintptr) (interface{}, error) {
	return wrapPolicyDecision(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (decision *PolicyDecision) basePolicyDecision() *PolicyDecision {
	return decision
}

// BasePolicyDecision returns the underlying base object.
func BasePolicyDecision(obj PolicyDecisioner) *PolicyDecision {
	return obj.basePolicyDecision()
}

// Download: spawn a download from this decision.
func (decision *PolicyDecision) Download() {
	var _arg0 *C.WebKitPolicyDecision // out

	_arg0 = (*C.WebKitPolicyDecision)(unsafe.Pointer(coreglib.InternObject(decision).Native()))

	C.webkit_policy_decision_download(_arg0)
	runtime.KeepAlive(decision)
}

// Ignore this would cancel the request.
//
// Ignore the action which triggered this decision. For instance, for a
// KitResponsePolicyDecision, this would cancel the request.
func (decision *PolicyDecision) Ignore() {
	var _arg0 *C.WebKitPolicyDecision // out

	_arg0 = (*C.WebKitPolicyDecision)(unsafe.Pointer(coreglib.InternObject(decision).Native()))

	C.webkit_policy_decision_ignore(_arg0)
	runtime.KeepAlive(decision)
}

// Use: accept the action which triggered this decision.
func (decision *PolicyDecision) Use() {
	var _arg0 *C.WebKitPolicyDecision // out

	_arg0 = (*C.WebKitPolicyDecision)(unsafe.Pointer(coreglib.InternObject(decision).Native()))

	C.webkit_policy_decision_use(_arg0)
	runtime.KeepAlive(decision)
}

// UseWithPolicies: accept the navigation action and continue with provided
// policies.
//
// Accept the navigation action which triggered this decision, and continue with
// policies affecting all subsequent loads of resources in the origin associated
// with the accepted navigation action.
//
// For example, a navigation decision to a video sharing website may be accepted
// under the priviso no movies are allowed to autoplay. The autoplay policy in
// this case would be set in the policies.
//
// The function takes the following parameters:
//
//   - policies: KitWebsitePolicies.
//
func (decision *PolicyDecision) UseWithPolicies(policies *WebsitePolicies) {
	var _arg0 *C.WebKitPolicyDecision  // out
	var _arg1 *C.WebKitWebsitePolicies // out

	_arg0 = (*C.WebKitPolicyDecision)(unsafe.Pointer(coreglib.InternObject(decision).Native()))
	_arg1 = (*C.WebKitWebsitePolicies)(unsafe.Pointer(coreglib.InternObject(policies).Native()))

	C.webkit_policy_decision_use_with_policies(_arg0, _arg1)
	runtime.KeepAlive(decision)
	runtime.KeepAlive(policies)
}

// PolicyDecisionClass: instance of this type is always passed by reference.
type PolicyDecisionClass struct {
	*policyDecisionClass
}

// policyDecisionClass is the struct that's finalized.
type policyDecisionClass struct {
	native *C.WebKitPolicyDecisionClass
}