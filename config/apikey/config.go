package apikey

import (
	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_api_key", func(r *config.Resource) {
		r.ShortGroup = "project.mongodbatlas"
		r.Kind = "ApiKey"
	})
}
