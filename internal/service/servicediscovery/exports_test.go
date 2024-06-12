// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicediscovery

// Exports for use in tests only.
var (
	ResourceHTTPNamespace       = resourceHTTPNamespace
	ResourcePrivateDNSNamespace = resourcePrivateDNSNamespace
	ResourcePublicDNSNamespace  = resourcePublicDNSNamespace

	FindNamespaceByID  = findNamespaceByID
	ValidNamespaceName = validNamespaceName
)
