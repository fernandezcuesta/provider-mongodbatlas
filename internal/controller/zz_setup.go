// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	advancedcluster "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/advancedcluster"
	auditing "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/auditing"
	cluster "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/cluster"
	accesslistapikey "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/org/accesslistapikey"
	apikey "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/project/apikey"
	providerconfig "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		advancedcluster.Setup,
		auditing.Setup,
		cluster.Setup,
		accesslistapikey.Setup,
		apikey.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
