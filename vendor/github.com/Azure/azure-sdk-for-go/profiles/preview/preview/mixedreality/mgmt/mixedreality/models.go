// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package mixedreality

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/mixedreality/mgmt/2019-02-28/mixedreality"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type NameAvailability = original.NameAvailability

const (
	False NameAvailability = original.False
	True  NameAvailability = original.True
)

type NameUnavailableReason = original.NameUnavailableReason

const (
	AlreadyExists NameUnavailableReason = original.AlreadyExists
	Invalid       NameUnavailableReason = original.Invalid
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SpatialAnchorsAccount = original.SpatialAnchorsAccount
type SpatialAnchorsAccountKeyRegenerateRequest = original.SpatialAnchorsAccountKeyRegenerateRequest
type SpatialAnchorsAccountKeys = original.SpatialAnchorsAccountKeys
type SpatialAnchorsAccountList = original.SpatialAnchorsAccountList
type SpatialAnchorsAccountListIterator = original.SpatialAnchorsAccountListIterator
type SpatialAnchorsAccountListPage = original.SpatialAnchorsAccountListPage
type SpatialAnchorsAccountProperties = original.SpatialAnchorsAccountProperties
type SpatialAnchorsAccountsClient = original.SpatialAnchorsAccountsClient
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSpatialAnchorsAccountListIterator(page SpatialAnchorsAccountListPage) SpatialAnchorsAccountListIterator {
	return original.NewSpatialAnchorsAccountListIterator(page)
}
func NewSpatialAnchorsAccountListPage(getNextPage func(context.Context, SpatialAnchorsAccountList) (SpatialAnchorsAccountList, error)) SpatialAnchorsAccountListPage {
	return original.NewSpatialAnchorsAccountListPage(getNextPage)
}
func NewSpatialAnchorsAccountsClient(subscriptionID string) SpatialAnchorsAccountsClient {
	return original.NewSpatialAnchorsAccountsClient(subscriptionID)
}
func NewSpatialAnchorsAccountsClientWithBaseURI(baseURI string, subscriptionID string) SpatialAnchorsAccountsClient {
	return original.NewSpatialAnchorsAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleNameAvailabilityValues() []NameAvailability {
	return original.PossibleNameAvailabilityValues()
}
func PossibleNameUnavailableReasonValues() []NameUnavailableReason {
	return original.PossibleNameUnavailableReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
