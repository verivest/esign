// Copyright 2017 James Cote and Liberty Fund, Inc.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go-swagger; DO NOT EDIT.

// Package connect implements the DocuSign SDK
// category Connect.
// 
// The Connect category enables your application to be called via HTTPS when an event of interest occurs.
// 
// Use the Connect service to "end the polling madness." With Connect, there is no need for your application to poll DocuSign every 15 minutes to learn the latest about your envelopes.
// 
// Instead, you register your interest in one or more types of envelope or recipient events. Then, when an interesting event occurs, the DocuSign platform will contact your application with the event's details and data. You can register interest in envelopes sent by particular users in your account, or for envelopes sent by any user.
// 
// ## Incoming Connect Calls
// To use the Connect service, your application needs to provide an https url that can be called from the public internet. If your application runs on a server behind your organization's firewall, then you will need to create a "pinhole" in the firewall to allow the incoming Connect calls from DocuSign to reach your application. Other techniques for receiving the incoming calls including proxy servers and DMZ networking can also be used. 
// 
// ## Per-envelope Connect Configuration
// Instead of registering a general Connect configuration and listener, an individual envelope can have its own Connect configuration. See the `eventNotification` field for envelopes.
// 
// ## Categories
// Use the Connect category for:
// 
// * Programmatically creating Connect configurations. Connect configurations can be created manually by using the DocuSign web service, or programmatically via the API. Configurations created via the API can be seen and updated from the web service.
// * Retrieving and managing the event log for your Connect configurations. 
// * Requesting that an event be re-published to the listener.
// Api documentation may be found at:
// https://docs.docusign.com/esign/restapi/Connect
package connect

import (
    "fmt"
    "net/url"
    "strings"
    "time"
    
    "golang.org/x/net/context"
    
    "mystuff/esign"
    "mystuff/esign/model"
)

// Service generates DocuSign Connect Category API calls
type Service struct {
    credential esign.Credential 
}

// New initializes a connect service using cred to authorize calls.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

// DeleteEventFailureLog deletes a Connect failure log entry.
// SDK Method Connect::deleteEventFailureLog
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/deleteFailure
func (s *Service) DeleteEventFailureLog(failureID string) *DeleteEventFailureLogCall {
    return &DeleteEventFailureLogCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "DELETE",
            Path: "connect/failures/{failureId}",
            PathParameters: map[string]string{ 
                "{failureId}": failureID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// DeleteEventFailureLogCall implements DocuSign API SDK Connect::deleteEventFailureLog
type DeleteEventFailureLogCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *DeleteEventFailureLogCall) Do(ctx context.Context)  error {
    
    return op.Call.Do(ctx, nil)
}

// ListEventFailureLogs gets the Connect failure log information.
// SDK Method Connect::listEventFailureLogs
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/listFailures
func (s *Service) ListEventFailureLogs() *ListEventFailureLogsCall {
    return &ListEventFailureLogsCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "connect/failures",
            QueryOpts: make(url.Values),
        },
    }
}

// ListEventFailureLogsCall implements DocuSign API SDK Connect::listEventFailureLogs
type ListEventFailureLogsCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListEventFailureLogsCall) Do(ctx context.Context)  (*model.ConnectLogs, error) {
    var res *model.ConnectLogs
    return res, op.Call.Do(ctx, &res)
}

// FromDate start of the search date range. Only returns templates created on or after this date/time. If no value is specified, there is no limit on the earliest date created.
func (op *ListEventFailureLogsCall) FromDate(val time.Time) *ListEventFailureLogsCall {
    op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
    return op
}

// ToDate end of the search date range. Only returns templates created up to this date/time. If no value is provided, this defaults to the current date.
func (op *ListEventFailureLogsCall) ToDate(val time.Time) *ListEventFailureLogsCall {
    op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
    return op
}

// DeleteEventLog deletes a specified Connect log entry.
// SDK Method Connect::deleteEventLog
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/delete
func (s *Service) DeleteEventLog(logID string) *DeleteEventLogCall {
    return &DeleteEventLogCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "DELETE",
            Path: "connect/logs/{logId}",
            PathParameters: map[string]string{ 
                "{logId}": logID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// DeleteEventLogCall implements DocuSign API SDK Connect::deleteEventLog
type DeleteEventLogCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *DeleteEventLogCall) Do(ctx context.Context)  error {
    
    return op.Call.Do(ctx, nil)
}

// DeleteEventLogs gets a list of Connect log entries.
// SDK Method Connect::deleteEventLogs
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/deleteList
func (s *Service) DeleteEventLogs() *DeleteEventLogsCall {
    return &DeleteEventLogsCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "DELETE",
            Path: "connect/logs",
            QueryOpts: make(url.Values),
        },
    }
}

// DeleteEventLogsCall implements DocuSign API SDK Connect::deleteEventLogs
type DeleteEventLogsCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *DeleteEventLogsCall) Do(ctx context.Context)  error {
    
    return op.Call.Do(ctx, nil)
}

// GetEventLog get the specified Connect log entry.
// SDK Method Connect::getEventLog
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/get
func (s *Service) GetEventLog(logID string) *GetEventLogCall {
    return &GetEventLogCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "connect/logs/{logId}",
            PathParameters: map[string]string{ 
                "{logId}": logID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// GetEventLogCall implements DocuSign API SDK Connect::getEventLog
type GetEventLogCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *GetEventLogCall) Do(ctx context.Context)  (*model.ConnectLog, error) {
    var res *model.ConnectLog
    return res, op.Call.Do(ctx, &res)
}

// AdditionalInfo when true, the connectDebugLog information is included in the response.
func (op *GetEventLogCall) AdditionalInfo() *GetEventLogCall {
    op.QueryOpts.Set("additional_info", "true")
    return op
}

// ListEventLogs gets the Connect log.
// SDK Method Connect::listEventLogs
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/list
func (s *Service) ListEventLogs() *ListEventLogsCall {
    return &ListEventLogsCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "connect/logs",
            QueryOpts: make(url.Values),
        },
    }
}

// ListEventLogsCall implements DocuSign API SDK Connect::listEventLogs
type ListEventLogsCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListEventLogsCall) Do(ctx context.Context)  (*model.ConnectLogs, error) {
    var res *model.ConnectLogs
    return res, op.Call.Do(ctx, &res)
}

// FromDate start of the search date range. Only returns templates created on or after this date/time. If no value is specified, there is no limit on the earliest date created.
func (op *ListEventLogsCall) FromDate(val time.Time) *ListEventLogsCall {
    op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
    return op
}

// ToDate end of the search date range. Only returns templates created up to this date/time. If no value is provided, this defaults to the current date.
func (op *ListEventLogsCall) ToDate(val time.Time) *ListEventLogsCall {
    op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
    return op
}

// RetryEventForEnvelopes republishes Connect information for multiple envelopes.
// SDK Method Connect::retryEventForEnvelopes
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/retryForEnvelopes
func (s *Service) RetryEventForEnvelopes(connectFailureFilter *model.ConnectFailureFilter) *RetryEventForEnvelopesCall {
    return &RetryEventForEnvelopesCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "PUT",
            Path: "connect/envelopes/retry_queue",
            Payload: connectFailureFilter,
            QueryOpts: make(url.Values),
        },
    }
}

// RetryEventForEnvelopesCall implements DocuSign API SDK Connect::retryEventForEnvelopes
type RetryEventForEnvelopesCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *RetryEventForEnvelopesCall) Do(ctx context.Context)  (*model.ConnectFailureResults, error) {
    var res *model.ConnectFailureResults
    return res, op.Call.Do(ctx, &res)
}

// RetryEventForEnvelope republishes Connect information for the specified envelope.
// SDK Method Connect::retryEventForEnvelope
// https://docs.docusign.com/esign/restapi/Connect/ConnectEvents/retryForEnvelope
func (s *Service) RetryEventForEnvelope(envelopeID string) *RetryEventForEnvelopeCall {
    return &RetryEventForEnvelopeCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "PUT",
            Path: "connect/envelopes/{envelopeId}/retry_queue",
            PathParameters: map[string]string{ 
                "{envelopeId}": envelopeID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// RetryEventForEnvelopeCall implements DocuSign API SDK Connect::retryEventForEnvelope
type RetryEventForEnvelopeCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *RetryEventForEnvelopeCall) Do(ctx context.Context)  (*model.ConnectFailureResults, error) {
    var res *model.ConnectFailureResults
    return res, op.Call.Do(ctx, &res)
}

// DeleteConfiguration deletes the specified connect configuration.
// SDK Method Connect::deleteConfiguration
// https://docs.docusign.com/esign/restapi/Connect/ConnectConfigurations/delete
func (s *Service) DeleteConfiguration(connectID string) *DeleteConfigurationCall {
    return &DeleteConfigurationCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "DELETE",
            Path: "connect/{connectId}",
            PathParameters: map[string]string{ 
                "{connectId}": connectID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// DeleteConfigurationCall implements DocuSign API SDK Connect::deleteConfiguration
type DeleteConfigurationCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *DeleteConfigurationCall) Do(ctx context.Context)  error {
    
    return op.Call.Do(ctx, nil)
}

// GetConfiguration get information on a Connect Configuration
// SDK Method Connect::getConfiguration
// https://docs.docusign.com/esign/restapi/Connect/ConnectConfigurations/get
func (s *Service) GetConfiguration(connectID string) *GetConfigurationCall {
    return &GetConfigurationCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "connect/{connectId}",
            PathParameters: map[string]string{ 
                "{connectId}": connectID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// GetConfigurationCall implements DocuSign API SDK Connect::getConfiguration
type GetConfigurationCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *GetConfigurationCall) Do(ctx context.Context)  (*model.ConnectConfigResults, error) {
    var res *model.ConnectConfigResults
    return res, op.Call.Do(ctx, &res)
}

// ListConfigurations get Connect Configuration Information
// SDK Method Connect::listConfigurations
// https://docs.docusign.com/esign/restapi/Connect/ConnectConfigurations/list
func (s *Service) ListConfigurations() *ListConfigurationsCall {
    return &ListConfigurationsCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "connect",
            QueryOpts: make(url.Values),
        },
    }
}

// ListConfigurationsCall implements DocuSign API SDK Connect::listConfigurations
type ListConfigurationsCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListConfigurationsCall) Do(ctx context.Context)  (*model.ConnectConfigResults, error) {
    var res *model.ConnectConfigResults
    return res, op.Call.Do(ctx, &res)
}

// ListUsers returns users from the configured Connect service.
// SDK Method Connect::listUsers
// https://docs.docusign.com/esign/restapi/Connect/ConnectConfigurations/listUsers
func (s *Service) ListUsers(connectID string) *ListUsersCall {
    return &ListUsersCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "connect/{connectId}/users",
            PathParameters: map[string]string{ 
                "{connectId}": connectID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// ListUsersCall implements DocuSign API SDK Connect::listUsers
type ListUsersCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListUsersCall) Do(ctx context.Context)  (*model.IntegratedUserInfoList, error) {
    var res *model.IntegratedUserInfoList
    return res, op.Call.Do(ctx, &res)
}

// Count optional. Number of items to return.
func (op *ListUsersCall) Count(val int) *ListUsersCall {
    op.QueryOpts.Set("count", fmt.Sprintf("%d", val ))
    return op
}

// EmailSubstring filters the returned user records by the email address or a sub-string of email address.
func (op *ListUsersCall) EmailSubstring(val string) *ListUsersCall {
    op.QueryOpts.Set("email_substring", val)
    return op
}

// ListIncludedUsers set the call query parameter list_included_users
func (op *ListUsersCall) ListIncludedUsers() *ListUsersCall {
    op.QueryOpts.Set("list_included_users", "true")
    return op
}

// StartPosition is the position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (op *ListUsersCall) StartPosition(val int) *ListUsersCall {
    op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val ))
    return op
}

// Status filters the results by user status.
// You can specify a comma-separated
// list of the following statuses:
// 
// * ActivationRequired 
// * ActivationSent 
// * Active
// * Closed 
// * Disabled
func (op *ListUsersCall) Status(val ...string) *ListUsersCall {
    op.QueryOpts.Set("status", strings.Join(val,","))
    return op
}

// UserNameSubstring filters the user records returned by the user name or a sub-string of user name.
func (op *ListUsersCall) UserNameSubstring(val string) *ListUsersCall {
    op.QueryOpts.Set("user_name_substring", val)
    return op
}

// CreateConfiguration creates a connect configuration for the specified account.
// SDK Method Connect::createConfiguration
// https://docs.docusign.com/esign/restapi/Connect/ConnectConfigurations/create
func (s *Service) CreateConfiguration(connectConfigurations *model.ConnectCustomConfiguration) *CreateConfigurationCall {
    return &CreateConfigurationCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "POST",
            Path: "connect",
            Payload: connectConfigurations,
            QueryOpts: make(url.Values),
        },
    }
}

// CreateConfigurationCall implements DocuSign API SDK Connect::createConfiguration
type CreateConfigurationCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *CreateConfigurationCall) Do(ctx context.Context)  (*model.ConnectCustomConfiguration, error) {
    var res *model.ConnectCustomConfiguration
    return res, op.Call.Do(ctx, &res)
}

// UpdateConfiguration updates a specified Connect configuration.
// SDK Method Connect::updateConfiguration
// https://docs.docusign.com/esign/restapi/Connect/ConnectConfigurations/update
func (s *Service) UpdateConfiguration(connectConfigurations *model.ConnectCustomConfiguration) *UpdateConfigurationCall {
    return &UpdateConfigurationCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "PUT",
            Path: "connect",
            Payload: connectConfigurations,
            QueryOpts: make(url.Values),
        },
    }
}

// UpdateConfigurationCall implements DocuSign API SDK Connect::updateConfiguration
type UpdateConfigurationCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *UpdateConfigurationCall) Do(ctx context.Context)  (*model.ConnectCustomConfiguration, error) {
    var res *model.ConnectCustomConfiguration
    return res, op.Call.Do(ctx, &res)
}
