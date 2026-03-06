// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package group

import (
	"context"
	"fmt"

	cclient "github.com/lukas016/controlplane/common/pkg/client"
	"github.com/lukas016/controlplane/common/pkg/condition"
	"github.com/lukas016/controlplane/common/pkg/handler"
	organizationv1 "github.com/lukas016/controlplane/organization/api/v1"
	"github.com/pkg/errors"
)

var _ handler.Handler[*organizationv1.Group] = &GroupHandler{}

type GroupHandler struct{}

func (h *GroupHandler) CreateOrUpdate(_ context.Context, groupObj *organizationv1.Group) error {
	groupObj.SetCondition(condition.NewDoneProcessingCondition("Created Group"))
	groupObj.SetCondition(condition.NewReadyCondition("Ready", "Group is ready"))
	return nil
}

func (h *GroupHandler) Delete(ctx context.Context, groupObj *organizationv1.Group) error {
	teams, err := GetTeamsForGroup(ctx, groupObj)
	if err != nil {
		groupObj.SetCondition(condition.NewBlockedCondition("Failed to get teams for group"))
		return errors.Wrap(err, fmt.Sprintf("failed to get teams for group %s", groupObj.GetName()))
	}

	k8sClient := cclient.ClientFromContextOrDie(ctx)
	for _, team := range teams.GetItems() {
		err = k8sClient.Delete(ctx, team)
		if err != nil {
			groupObj.SetCondition(condition.NewBlockedCondition("Failed to delete team"))
			return errors.Wrap(err, fmt.Sprintf("failed to delete team %s", team.GetName()))
		}
	}

	return nil
}
