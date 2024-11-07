package advancedcluster

import (
	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_advanced_cluster", func(r *config.Resource) {
		r.ShortGroup = "mongodbatlas"
		r.Kind = "AdvancedCluster"
	})
}
