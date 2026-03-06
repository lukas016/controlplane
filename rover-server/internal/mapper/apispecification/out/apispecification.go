// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package out

import (
	"github.com/lukas016/controlplane/rover-server/internal/api"
	"github.com/lukas016/controlplane/rover-server/internal/mapper"
	"github.com/lukas016/controlplane/rover-server/internal/mapper/status"
	roverv1 "github.com/lukas016/controlplane/rover/api/v1"
	"github.com/pkg/errors"
)

func MapResponse(in *roverv1.ApiSpecification, inFile map[string]any) (res api.ApiSpecificationResponse, err error) {
	if in == nil {
		return res, errors.New("input api specification crd is nil")
	}

	if inFile == nil {
		return res, errors.New("input api specification is nil")
	}

	res = api.ApiSpecificationResponse{
		Category:      in.Spec.Category,
		Id:            mapper.MakeResourceId(in),
		Name:          in.Name,
		Specification: inFile,
		VendorApi:     in.Spec.XVendor,
	}
	res.Status = status.MapStatus(in.Status.Conditions)

	return
}
