/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane-contrib/provider-mongodbatlas/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"mongodbatlas_access_list_api_key": config.ParameterAsIdentifier("org_id"),
	"mongodbatlas_api_key":             config.ParameterAsIdentifier("org_id"),
	"mongodbatlas_auditing":            config.IdentifierFromProvider,
	"mongodbatlas_advanced_cluster": config.TemplatedStringAsIdentifier(
		"name",
		"{{ .parameters.project_id }}-{{ .externalName }}",
	),
	"mongodbatlas_cluster": config.TemplatedStringAsIdentifier(
		"name",
		"{{ .parameters.project_id }}-{{ .externalName }}",
	),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
			// Potential breaking changes, bumping API version to v1alpha3.
			r.Version = common.VersionV1Alpha3
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
