/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-mongodbatlas/config/advancedcluster"
	"github.com/crossplane-contrib/provider-mongodbatlas/config/auditing"
	"github.com/crossplane-contrib/provider-mongodbatlas/config/cluster"
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "mongodbatlas"
	modulePath     = "github.com/crossplane-contrib/provider-mongodbatlas"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithShortName("mongodbatlas"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		advancedcluster.Configure,
		auditing.Configure,
		cluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
