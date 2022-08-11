package util

import (
	"k8s.io/utils/pointer"
	gatewayv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// -----------------------------------------------------------------------------
// Type conversion Utilities
// -----------------------------------------------------------------------------

// StringToGatewayHostname converts a string to a gatewayv1alpha2.Hostname.
func StringToGatewayHostname(hostname string) gatewayv1alpha2.Hostname {
	return (gatewayv1alpha2.Hostname)(hostname)
}

// StringToGatewayHostnamePtr converts a string to a *gatewayv1alpha2.Hostname.
func StringToGatewayHostnamePtr(hostname string) *gatewayv1alpha2.Hostname {
	return (*gatewayv1alpha2.Hostname)(pointer.StringPtr(hostname))
}

// StringToKind converts a string to a gatewayv1alpha2.Kind.
func StringToKind(kind string) gatewayv1alpha2.Kind {
	return (gatewayv1alpha2.Kind)(kind)
}

// StringToKindPtr converts a string to a *gatewayv1alpha2.Kind.
func StringToKindPtr(kind string) *gatewayv1alpha2.Kind {
	return (*gatewayv1alpha2.Kind)(pointer.StringPtr(kind))
}
