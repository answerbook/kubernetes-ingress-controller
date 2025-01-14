package util

import (
	"k8s.io/utils/pointer"
	gatewayv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// -----------------------------------------------------------------------------
// Type conversion Utilities
// -----------------------------------------------------------------------------

// StringToGatewayAPIHostname converts a string to a gatewayv1alpha2.Hostname.
func StringToGatewayAPIHostname(hostname string) gatewayv1alpha2.Hostname {
	return (gatewayv1alpha2.Hostname)(hostname)
}

// StringToGatewayAPIHostnamePtr converts a string to a *gatewayv1alpha2.Hostname.
func StringToGatewayAPIHostnamePtr(hostname string) *gatewayv1alpha2.Hostname {
	return (*gatewayv1alpha2.Hostname)(pointer.StringPtr(hostname))
}

// StringToGatewayAPIKind converts a string to a gatewayv1alpha2.Kind.
func StringToGatewayAPIKind(kind string) gatewayv1alpha2.Kind {
	return (gatewayv1alpha2.Kind)(kind)
}

// StringToGatewayAPIKindPtr converts a string to a *gatewayv1alpha2.Kind.
func StringToGatewayAPIKindPtr(kind string) *gatewayv1alpha2.Kind {
	return (*gatewayv1alpha2.Kind)(pointer.StringPtr(kind))
}
