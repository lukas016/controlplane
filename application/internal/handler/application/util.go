// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package application

import (
	"context"

	admin "github.com/lukas016/controlplane/admin/api/v1"
	application "github.com/lukas016/controlplane/application/api/v1"
	"github.com/lukas016/controlplane/common/pkg/client"
	"github.com/lukas016/controlplane/common/pkg/types"
	identity "github.com/lukas016/controlplane/identity/api/v1"
	"github.com/pkg/errors"
)

func MakeClientName(obj *application.Application) string {
	return obj.Spec.Team + "--" + obj.Name
}

func GetZone(ctx context.Context, client client.ScopedClient, zoneRef types.ObjectRef) (*admin.Zone, error) {
	zone := &admin.Zone{}
	err := client.Get(ctx, zoneRef.K8s(), zone)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get zone %s", zoneRef.Name)
	}
	return zone, nil

}

func GetIdpClient(ctx context.Context, client client.ScopedClient, obj *application.Application, clientName string, namespace string) (*identity.Client, error) {

	clientRef := &types.ObjectRef{
		Name:      clientName,
		Namespace: namespace,
	}

	idpClient := &identity.Client{}

	err := client.Get(ctx, clientRef.K8s(), idpClient)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get zone")
	}
	return idpClient, nil

}
