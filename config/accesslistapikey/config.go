package accesslistapikey

import (
	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_access_list_api_key", func(r *config.Resource) {
		r.ShortGroup = "org.mongodbatlas"
		r.Kind = "AccessListApiKey"
	})
}
