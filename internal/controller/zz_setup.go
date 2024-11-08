// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	compliancepolicy "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/backup/compliancepolicy"
	schedule "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/backup/schedule"
	snapshot "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/backup/snapshot"
	advancedcluster "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/advancedcluster"
	alertconfiguration "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/alertconfiguration"
	auditing "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/auditing"
	cluster "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/mongodbatlas/cluster"
	archive "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/online/archive"
	accesslistapikey "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/org/accesslistapikey"
	apikey "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/project/apikey"
	providerconfig "github.com/crossplane-contrib/provider-mongodbatlas/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		compliancepolicy.Setup,
		schedule.Setup,
		snapshot.Setup,
		advancedcluster.Setup,
		alertconfiguration.Setup,
		auditing.Setup,
		cluster.Setup,
		archive.Setup,
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
