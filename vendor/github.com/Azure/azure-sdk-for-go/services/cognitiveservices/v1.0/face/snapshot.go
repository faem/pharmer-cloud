package face

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// SnapshotClient is the an API for face detection, verification, and identification.
type SnapshotClient struct {
	BaseClient
}

// NewSnapshotClient creates an instance of the SnapshotClient client.
func NewSnapshotClient(endpoint string) SnapshotClient {
	return SnapshotClient{New(endpoint)}
}

// Apply submit an operation to apply a snapshot to current subscription. For each snapshot, only subscriptions
// included in the applyScope of Snapshot - Take can apply it.<br />
// The snapshot interfaces are for users to backup and restore their face data from one face subscription to another,
// inside same region or across regions. The workflow contains two phases, user first calls Snapshot - Take to create a
// copy of the source object and store it as a snapshot, then calls Snapshot - Apply to paste the snapshot to target
// subscription. The snapshots are stored in a centralized location (per Azure instance), so that they can be applied
// cross accounts and regions.<br />
// Applying snapshot is an asynchronous operation. An operation id can be obtained from the "Operation-Location" field
// in response header, to be used in OperationStatus - Get for tracking the progress of applying the snapshot. The
// target object id will be included in the "resourceLocation" field in OperationStatus - Get response when the
// operation status is "succeeded".<br />
// Snapshot applying time depends on the number of person and face entries in the snapshot object. It could be in
// seconds, or up to 1 hour for 1,000,000 persons with multiple faces.<br />
// Snapshots will be automatically expired and cleaned in 48 hours after it is created by Snapshot - Take. So the
// target subscription is required to apply the snapshot in 48 hours since its creation.<br />
// Applying a snapshot will not block any other operations against the target object, however it is not recommended
// because the correctness cannot be guaranteed during snapshot applying. After snapshot applying is completed, all
// operations towards the target object can work as normal. Snapshot also includes the training results of the source
// object, which means target subscription the snapshot applied to does not need re-train the target object before
// calling Identify/FindSimilar.<br />
// One snapshot can be applied multiple times in parallel, while currently only CreateNew apply mode is supported,
// which means the apply operation will fail if target subscription already contains an object of same type and using
// the same objectId. Users can specify the "objectId" in request body to avoid such conflicts.<br />
// * Free-tier subscription quota: 100 apply operations per month.
// * S0-tier subscription quota: 100 apply operations per day.
// Parameters:
// snapshotID - id referencing a particular snapshot.
// body - request body for applying a snapshot.
func (client SnapshotClient) Apply(ctx context.Context, snapshotID uuid.UUID, body ApplySnapshotRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.Apply")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.ObjectID", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.ObjectID", Name: validation.MaxLength, Rule: 64, Chain: nil},
					{Target: "body.ObjectID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("face.SnapshotClient", "Apply", err.Error())
	}

	req, err := client.ApplyPreparer(ctx, snapshotID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Apply", nil, "Failure preparing request")
		return
	}

	resp, err := client.ApplySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Apply", resp, "Failure sending request")
		return
	}

	result, err = client.ApplyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Apply", resp, "Failure responding to request")
	}

	return
}

// ApplyPreparer prepares the Apply request.
func (client SnapshotClient) ApplyPreparer(ctx context.Context, snapshotID uuid.UUID, body ApplySnapshotRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"snapshotId": autorest.Encode("path", snapshotID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/snapshots/{snapshotId}/apply", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ApplySender sends the Apply request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) ApplySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ApplyResponder handles the response to the Apply request. The method always
// closes the http.Response Body.
func (client SnapshotClient) ApplyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete delete an existing snapshot according to the snapshotId. All object data and information in the snapshot will
// also be deleted. Only the source subscription who took the snapshot can delete the snapshot. If the user does not
// delete a snapshot with this API, the snapshot will still be automatically deleted in 48 hours after creation.
// Parameters:
// snapshotID - id referencing a particular snapshot.
func (client SnapshotClient) Delete(ctx context.Context, snapshotID uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, snapshotID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SnapshotClient) DeletePreparer(ctx context.Context, snapshotID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"snapshotId": autorest.Encode("path", snapshotID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/snapshots/{snapshotId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SnapshotClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve information about a snapshot. Snapshot is only accessible to the source subscription who took it, and
// target subscriptions included in the applyScope in Snapshot - Take.
// Parameters:
// snapshotID - id referencing a particular snapshot.
func (client SnapshotClient) Get(ctx context.Context, snapshotID uuid.UUID) (result Snapshot, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, snapshotID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SnapshotClient) GetPreparer(ctx context.Context, snapshotID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"snapshotId": autorest.Encode("path", snapshotID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/snapshots/{snapshotId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SnapshotClient) GetResponder(resp *http.Response) (result Snapshot, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOperationStatus retrieve the status of a take/apply snapshot operation.
// Parameters:
// operationID - id referencing a particular take/apply snapshot operation.
func (client SnapshotClient) GetOperationStatus(ctx context.Context, operationID uuid.UUID) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.GetOperationStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationStatusPreparer(ctx, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "GetOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "GetOperationStatus", resp, "Failure responding to request")
	}

	return
}

// GetOperationStatusPreparer prepares the GetOperationStatus request.
func (client SnapshotClient) GetOperationStatusPreparer(ctx context.Context, operationID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"operationId": autorest.Encode("path", operationID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/operations/{operationId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationStatusSender sends the GetOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) GetOperationStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOperationStatusResponder handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (client SnapshotClient) GetOperationStatusResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all accessible snapshots with related information, including snapshots that were taken by the user, or
// snapshots to be applied to the user (subscription id was included in the applyScope in Snapshot - Take).
// Parameters:
// typeParameter - user specified object type as a search filter.
// applyScope - user specified snapshot apply scopes as a search filter. ApplyScope is an array of the target
// Azure subscription ids for the snapshot, specified by the user who created the snapshot by Snapshot - Take.
func (client SnapshotClient) List(ctx context.Context, typeParameter SnapshotObjectType, applyScope []uuid.UUID) (result ListSnapshot, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, typeParameter, applyScope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SnapshotClient) ListPreparer(ctx context.Context, typeParameter SnapshotObjectType, applyScope []uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	queryParameters := map[string]interface{}{}
	if len(string(typeParameter)) > 0 {
		queryParameters["type"] = autorest.Encode("query", typeParameter)
	}
	if applyScope != nil && len(applyScope) > 0 {
		queryParameters["applyScope"] = autorest.Encode("query", applyScope, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPath("/snapshots"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SnapshotClient) ListResponder(resp *http.Response) (result ListSnapshot, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Take submit an operation to take a snapshot of face list, large face list, person group or large person group, with
// user-specified snapshot type, source object id, apply scope and an optional user data.<br />
// The snapshot interfaces are for users to backup and restore their face data from one face subscription to another,
// inside same region or across regions. The workflow contains two phases, user first calls Snapshot - Take to create a
// copy of the source object and store it as a snapshot, then calls Snapshot - Apply to paste the snapshot to target
// subscription. The snapshots are stored in a centralized location (per Azure instance), so that they can be applied
// cross accounts and regions.<br />
// Taking snapshot is an asynchronous operation. An operation id can be obtained from the "Operation-Location" field in
// response header, to be used in OperationStatus - Get for tracking the progress of creating the snapshot. The
// snapshot id will be included in the "resourceLocation" field in OperationStatus - Get response when the operation
// status is "succeeded".<br />
// Snapshot taking time depends on the number of person and face entries in the source object. It could be in seconds,
// or up to several hours for 1,000,000 persons with multiple faces.<br />
// Snapshots will be automatically expired and cleaned in 48 hours after it is created by Snapshot - Take. User can
// delete the snapshot using Snapshot - Delete by themselves any time before expiration.<br />
// Taking snapshot for a certain object will not block any other operations against the object. All read-only
// operations (Get/List and Identify/FindSimilar/Verify) can be conducted as usual. For all writable operations,
// including Add/Update/Delete the source object or its persons/faces and Train, they are not blocked but not
// recommended because writable updates may not be reflected on the snapshot during its taking. After snapshot taking
// is completed, all readable and writable operations can work as normal. Snapshot will also include the training
// results of the source object, which means target subscription the snapshot applied to does not need re-train the
// target object before calling Identify/FindSimilar.<br />
// * Free-tier subscription quota: 100 take operations per month.
// * S0-tier subscription quota: 100 take operations per day.
// Parameters:
// body - request body for taking a snapshot.
func (client SnapshotClient) Take(ctx context.Context, body TakeSnapshotRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.Take")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.ObjectID", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.ObjectID", Name: validation.MaxLength, Rule: 64, Chain: nil},
					{Target: "body.ObjectID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil},
				}},
				{Target: "body.ApplyScope", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.UserData", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.UserData", Name: validation.MaxLength, Rule: 16384, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("face.SnapshotClient", "Take", err.Error())
	}

	req, err := client.TakePreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Take", nil, "Failure preparing request")
		return
	}

	resp, err := client.TakeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Take", resp, "Failure sending request")
		return
	}

	result, err = client.TakeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Take", resp, "Failure responding to request")
	}

	return
}

// TakePreparer prepares the Take request.
func (client SnapshotClient) TakePreparer(ctx context.Context, body TakeSnapshotRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPath("/snapshots"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TakeSender sends the Take request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) TakeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TakeResponder handles the response to the Take request. The method always
// closes the http.Response Body.
func (client SnapshotClient) TakeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update the information of a snapshot. Only the source subscription who took the snapshot can update the
// snapshot.
// Parameters:
// snapshotID - id referencing a particular snapshot.
// body - request body for updating a snapshot.
func (client SnapshotClient) Update(ctx context.Context, snapshotID uuid.UUID, body UpdateSnapshotRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, snapshotID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.SnapshotClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SnapshotClient) UpdatePreparer(ctx context.Context, snapshotID uuid.UUID, body UpdateSnapshotRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"snapshotId": autorest.Encode("path", snapshotID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/snapshots/{snapshotId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SnapshotClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
