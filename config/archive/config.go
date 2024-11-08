package archive

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_online_archive", func(r *config.Resource) {
		r.ShortGroup = "online.mongodbatlas"
		r.Kind = "Archive"
		r.ExternalName.GetExternalNameFn = getExternalName
		r.ExternalName.GetIDFn = getID
	})
}

func getExternalName(tfstate map[string]any) (string, error) {
	aid, ok := tfstate["archive_id"]
	if !ok {
		return "", errors.New("archive_id in tfstate cannot be empty")
	}
	return aid.(string), nil
}

func getID(_ context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
	p, ok := parameters["project_id"]
	if !ok {
		return "", errors.New("project_id cannot be empty")
	}
	c, ok := parameters["cluster_name"]
	if !ok {
		return "", errors.New("cluster_name cannot be empty")
	}
	return fmt.Sprintf("%s-%s-%s", p.(string), c.(string), externalName), nil
}
