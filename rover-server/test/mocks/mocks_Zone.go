// Copyright 2025 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	adminv1 "github.com/lukas016/controlplane/admin/api/v1"
	"github.com/lukas016/controlplane/common-server/pkg/store"
	"github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/mock"
)

func NewZoneStoreMock(testing ginkgo.FullGinkgoTInterface) store.ObjectStore[*adminv1.Zone] {
	mockStore := NewMockObjectStore[*adminv1.Zone](testing)
	ConfigureZoneStoreMock(testing, mockStore)
	return mockStore
}

func ConfigureZoneStoreMock(testing ginkgo.FullGinkgoTInterface, mockedStore *MockObjectStore[*adminv1.Zone]) {
	zone := GetZone(testing, zoneFileName)
	mockedStore.EXPECT().Get(
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("string"),
		mock.Anything,
	).Return(zone, nil).Maybe()
}
