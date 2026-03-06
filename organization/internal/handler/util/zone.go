// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"context"

	adminv1 "github.com/lukas016/controlplane/admin/api/v1"
	cclient "github.com/lukas016/controlplane/common/pkg/client"
	"github.com/lukas016/controlplane/organization/internal/index"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetZoneObjWithTeamInfo(ctx context.Context) (*adminv1.Zone, error) {
	var teamApiZone *adminv1.Zone = nil
	zoneList := &adminv1.ZoneList{}
	clientFromContext := cclient.ClientFromContextOrDie(ctx)

	err := clientFromContext.List(ctx, zoneList, client.MatchingFields{index.FieldSpecTeamApis: "true"})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list zones")
	}

	for _, zone := range zoneList.GetItems() {
		if zone.(*adminv1.Zone).Spec.TeamApis != nil {
			teamApiZone = zone.(*adminv1.Zone).DeepCopy()
			break
		}
	}

	if teamApiZone == nil {
		return nil, errors.New("found no zone with team apis")
	}

	return teamApiZone, nil
}
