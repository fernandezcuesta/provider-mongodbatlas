package alertconfiguration

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mongodbatlas_alert_configuration", func(r *config.Resource) {
		r.ShortGroup = "mongodbatlas"
		r.Kind = "AlertConfiguration"
		r.ExternalName = config.ParameterAsIdentifier("project_id")
		r.ExternalName.GetExternalNameFn = getExternalName
		r.ExternalName.GetIDFn = getID
	})
}

func getExternalName(tfstate map[string]any) (string, error) {
	aid, ok := tfstate["alert_configuration_id"]
	if !ok {
		return "", errors.New("alert_configuration_id in tfstate cannot be empty")
	}
	return aid, nil
}

func getID(_ context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
	p, ok := parameters["project_id"]
	if !ok {
		return "", errors.New("project_id cannot be empty")
	}
	return fmt.Sprintf("%s-%s", p.(string), externalName), nil
}
