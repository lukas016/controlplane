// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	adminapi "github.com/lukas016/controlplane/admin/api/v1"
	apiapi "github.com/lukas016/controlplane/api/api/v1"
	applicationapi "github.com/lukas016/controlplane/application/api/v1"
	approvalapi "github.com/lukas016/controlplane/approval/api/v1"
	gatewayapi "github.com/lukas016/controlplane/gateway/api/v1"
	identityapi "github.com/lukas016/controlplane/identity/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

func RegisterSchemesOrDie(scheme *runtime.Scheme) {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(approvalapi.AddToScheme(scheme))
	utilruntime.Must(gatewayapi.AddToScheme(scheme))
	utilruntime.Must(adminapi.AddToScheme(scheme))
	utilruntime.Must(apiapi.AddToScheme(scheme))
	utilruntime.Must(identityapi.AddToScheme(scheme))
	utilruntime.Must(applicationapi.AddToScheme(scheme))
}
