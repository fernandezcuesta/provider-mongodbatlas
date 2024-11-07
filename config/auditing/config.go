package auditing

import (
	"github.com/crossplane-contrib/provider-mongodbatlas/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_auditing", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha3
	})
}
