// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package remoteapisubscription

import (
	"context"
	"fmt"

	adminapi "github.com/lukas016/controlplane/admin/api/v1"
	apiapi "github.com/lukas016/controlplane/api/api/v1"
	approvalapi "github.com/lukas016/controlplane/approval/api/v1"
	"github.com/lukas016/controlplane/common/pkg/client"
	"github.com/lukas016/controlplane/common/pkg/types"
	gatewayapi "github.com/lukas016/controlplane/gateway/api/v1"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func CalculateRemoteOrgZone(remoteOrg *adminapi.RemoteOrganization) types.ObjectRef {
	return types.ObjectRef{
		Name:      fmt.Sprintf("%s-%s", remoteOrg.Spec.Id, remoteOrg.Spec.Zone.Name),
		Namespace: remoteOrg.Spec.Zone.Namespace,
	}
}

func fillApprovalInfo(ctx context.Context, obj *apiapi.RemoteApiSubscription, apiSubscription *apiapi.ApiSubscription) (err error) {
	if apiSubscription.Status.Approval == nil {
		return nil
	}

	c := client.ClientFromContextOrDie(ctx)

	approval := &approvalapi.Approval{}
	err = c.Get(ctx, apiSubscription.Status.Approval.K8s(), approval)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return errors.Wrapf(err, "failed to get approval %s", apiSubscription.Status.Approval.Name)
		}
		return nil
	}
	obj.Status.Approval = &apiapi.ApprovalInfo{
		ApprovalState: approval.Spec.State.String(),
		Message:       "", // todo - resolve later, should be taken from decisions
	}
	return
}

func fillApprovalRequestInfo(ctx context.Context, obj *apiapi.RemoteApiSubscription, apiSubscription *apiapi.ApiSubscription) (err error) {
	if apiSubscription.Status.ApprovalRequest == nil {
		return nil
	}

	c := client.ClientFromContextOrDie(ctx)

	approvalRequest := &approvalapi.ApprovalRequest{}
	err = c.Get(ctx, apiSubscription.Status.ApprovalRequest.K8s(), approvalRequest)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return errors.Wrapf(err, "failed to get approval %s", apiSubscription.Status.Approval.Name)
		}
		return nil
	}
	obj.Status.ApprovalRequest = &apiapi.ApprovalInfo{
		ApprovalState: approvalRequest.Spec.State.String(),
		Message:       "", // todo - resolve later, should be taken from decisions
	}
	return
}

func fillRouteInfo(ctx context.Context, obj *apiapi.RemoteApiSubscription, apiSubscription *apiapi.ApiSubscription) (err error) {
	if apiSubscription.Status.Route == nil {
		return nil
	}

	c := client.ClientFromContextOrDie(ctx)
	downstreamRoute := &gatewayapi.Route{}
	err = c.Get(ctx, apiSubscription.Status.Route.K8s(), downstreamRoute)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return errors.Wrapf(err, "failed to get route %s", apiSubscription.Status.Route.Name)
		}
		return nil
	}
	// TODO: This is shit. What if we have multiple downstreams? Why is it like this?
	obj.Status.GatewayUrl = "https://" + downstreamRoute.Spec.Downstreams[0].Host + downstreamRoute.Spec.Downstreams[0].Path
	return
}
