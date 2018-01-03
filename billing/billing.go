// Copyright 2017 James Cote and Liberty Fund, Inc.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go-swagger; DO NOT EDIT.

// Package billing implements the DocuSign SDK
// category Billing.
// 
// Use the Billing category to manage the following billing related tasks:
// 
// * Retrieve and update billing plan information.
// * Retrieve invoices.
// * Retrieve and update payment information.
// Api documentation may be found at:
// https://docs.docusign.com/esign/restapi/Billing
package billing

import (
    "net/url"
    "time"
    
    "golang.org/x/net/context"
    
    "mystuff/esign"
    "mystuff/esign/model"
)

// Service generates DocuSign Billing Category API calls
type Service struct {
    credential esign.Credential 
}

// New initializes a billing service using cred to authorize calls.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

// GetInvoice retrieves a billing invoice.
// SDK Method Billing::getInvoice
// https://docs.docusign.com/esign/restapi/Billing/Invoices/get
func (s *Service) GetInvoice(invoiceID string) *GetInvoiceCall {
    return &GetInvoiceCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "billing_invoices/{invoiceId}",
            PathParameters: map[string]string{ 
                "{invoiceId}": invoiceID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// GetInvoiceCall implements DocuSign API SDK Billing::getInvoice
type GetInvoiceCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *GetInvoiceCall) Do(ctx context.Context)  (*model.BillingInvoice, error) {
    var res *model.BillingInvoice
    return res, op.Call.Do(ctx, &res)
}

// ListInvoices get a List of Billing Invoices
// SDK Method Billing::listInvoices
// https://docs.docusign.com/esign/restapi/Billing/Invoices/list
func (s *Service) ListInvoices() *ListInvoicesCall {
    return &ListInvoicesCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "billing_invoices",
            QueryOpts: make(url.Values),
        },
    }
}

// ListInvoicesCall implements DocuSign API SDK Billing::listInvoices
type ListInvoicesCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListInvoicesCall) Do(ctx context.Context)  (*model.BillingInvoicesResponse, error) {
    var res *model.BillingInvoicesResponse
    return res, op.Call.Do(ctx, &res)
}

// FromDate specifies the date/time of the earliest invoice in the account to retrieve.
func (op *ListInvoicesCall) FromDate(val time.Time) *ListInvoicesCall {
    op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
    return op
}

// ToDate specifies the date/time of the latest invoice in the account to retrieve.
func (op *ListInvoicesCall) ToDate(val time.Time) *ListInvoicesCall {
    op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
    return op
}

// ListInvoicesPastDue get a list of past due invoices.
// SDK Method Billing::listInvoicesPastDue
// https://docs.docusign.com/esign/restapi/Billing/Invoices/listPastDue
func (s *Service) ListInvoicesPastDue() *ListInvoicesPastDueCall {
    return &ListInvoicesPastDueCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "billing_invoices_past_due",
            QueryOpts: make(url.Values),
        },
    }
}

// ListInvoicesPastDueCall implements DocuSign API SDK Billing::listInvoicesPastDue
type ListInvoicesPastDueCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListInvoicesPastDueCall) Do(ctx context.Context)  (*model.BillingInvoicesSummary, error) {
    var res *model.BillingInvoicesSummary
    return res, op.Call.Do(ctx, &res)
}

// GetPayment gets billing payment information for a specific payment.
// SDK Method Billing::getPayment
// https://docs.docusign.com/esign/restapi/Billing/Payments/get
func (s *Service) GetPayment(paymentID string) *GetPaymentCall {
    return &GetPaymentCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "billing_payments/{paymentId}",
            PathParameters: map[string]string{ 
                "{paymentId}": paymentID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// GetPaymentCall implements DocuSign API SDK Billing::getPayment
type GetPaymentCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *GetPaymentCall) Do(ctx context.Context)  (*model.BillingPaymentItem, error) {
    var res *model.BillingPaymentItem
    return res, op.Call.Do(ctx, &res)
}

// ListPayments gets payment information for one or more payments.
// SDK Method Billing::listPayments
// https://docs.docusign.com/esign/restapi/Billing/Payments/list
func (s *Service) ListPayments() *ListPaymentsCall {
    return &ListPaymentsCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "billing_payments",
            QueryOpts: make(url.Values),
        },
    }
}

// ListPaymentsCall implements DocuSign API SDK Billing::listPayments
type ListPaymentsCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListPaymentsCall) Do(ctx context.Context)  (*model.BillingPaymentsResponse, error) {
    var res *model.BillingPaymentsResponse
    return res, op.Call.Do(ctx, &res)
}

// FromDate specifies the date/time of the earliest payment in the account to retrieve.
func (op *ListPaymentsCall) FromDate(val time.Time) *ListPaymentsCall {
    op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
    return op
}

// ToDate specifies the date/time of the latest payment in the account to retrieve.
func (op *ListPaymentsCall) ToDate(val time.Time) *ListPaymentsCall {
    op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
    return op
}

// MakePayment posts a payment to a past due invoice.
// SDK Method Billing::makePayment
// https://docs.docusign.com/esign/restapi/Billing/Payments/create
func (s *Service) MakePayment(billingPaymentRequest *model.BillingPaymentRequest) *MakePaymentCall {
    return &MakePaymentCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "POST",
            Path: "billing_payments",
            Payload: billingPaymentRequest,
            QueryOpts: make(url.Values),
        },
    }
}

// MakePaymentCall implements DocuSign API SDK Billing::makePayment
type MakePaymentCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *MakePaymentCall) Do(ctx context.Context)  (*model.BillingPaymentResponse, error) {
    var res *model.BillingPaymentResponse
    return res, op.Call.Do(ctx, &res)
}

// GetPlan get Account Billing Plan
// SDK Method Billing::getPlan
// https://docs.docusign.com/esign/restapi/Billing/BillingPlans/getAccountPlan
func (s *Service) GetPlan() *GetPlanCall {
    return &GetPlanCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "billing_plan",
            QueryOpts: make(url.Values),
        },
    }
}

// GetPlanCall implements DocuSign API SDK Billing::getPlan
type GetPlanCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *GetPlanCall) Do(ctx context.Context)  (*model.AccountBillingPlanResponse, error) {
    var res *model.AccountBillingPlanResponse
    return res, op.Call.Do(ctx, &res)
}

// IncludeCreditCardInformation when set to **true**, excludes credit card information from the response.
func (op *GetPlanCall) IncludeCreditCardInformation() *GetPlanCall {
    op.QueryOpts.Set("include_credit_card_information", "true")
    return op
}

// IncludeMetadata when set to **true**, the `canUpgrade` and `renewalStatus` properities are included the response and an array of `supportedCountries` property is added to the `billingAddress` information.
func (op *GetPlanCall) IncludeMetadata() *GetPlanCall {
    op.QueryOpts.Set("include_metadata", "true")
    return op
}

// IncludeSuccessorPlans when set to **true**, excludes successor information from the response.
func (op *GetPlanCall) IncludeSuccessorPlans() *GetPlanCall {
    op.QueryOpts.Set("include_successor_plans", "true")
    return op
}

// GetCreditCardInfo get metadata for a given credit card.
// SDK Method Billing::getCreditCardInfo
// https://docs.docusign.com/esign/restapi/Billing/BillingPlans/getCreditCard
func (s *Service) GetCreditCardInfo() *GetCreditCardInfoCall {
    return &GetCreditCardInfoCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "billing_plan/credit_card",
            QueryOpts: make(url.Values),
        },
    }
}

// GetCreditCardInfoCall implements DocuSign API SDK Billing::getCreditCardInfo
type GetCreditCardInfoCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *GetCreditCardInfoCall) Do(ctx context.Context)  (*model.CreditCardInformation, error) {
    var res *model.CreditCardInformation
    return res, op.Call.Do(ctx, &res)
}

// UpdatePlan updates the account billing plan.
// SDK Method Billing::updatePlan
// https://docs.docusign.com/esign/restapi/Billing/BillingPlans/update
func (s *Service) UpdatePlan(billingPlanInformation *model.BillingPlanInformation) *UpdatePlanCall {
    return &UpdatePlanCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "PUT",
            Path: "billing_plan",
            Payload: billingPlanInformation,
            QueryOpts: make(url.Values),
        },
    }
}

// UpdatePlanCall implements DocuSign API SDK Billing::updatePlan
type UpdatePlanCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *UpdatePlanCall) Do(ctx context.Context)  (*model.BillingPlanUpdateResponse, error) {
    var res *model.BillingPlanUpdateResponse
    return res, op.Call.Do(ctx, &res)
}

// PreviewBillingPlan when set to **true**, updates the account using a preview billing plan.
func (op *UpdatePlanCall) PreviewBillingPlan() *UpdatePlanCall {
    op.QueryOpts.Set("preview_billing_plan", "true")
    return op
}

// GetBillingPlan get the billing plan details.
// SDK Method Billing::getBillingPlan
// https://docs.docusign.com/esign/restapi/Billing/BillingPlans/get
func (s *Service) GetBillingPlan(billingPlanID string) *GetBillingPlanCall {
    return &GetBillingPlanCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "/v2/billing_plans/{billingPlanId}",
            PathParameters: map[string]string{ 
                "{billingPlanId}": billingPlanID,
            },
            QueryOpts: make(url.Values),
        },
    }
}

// GetBillingPlanCall implements DocuSign API SDK Billing::getBillingPlan
type GetBillingPlanCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *GetBillingPlanCall) Do(ctx context.Context)  (*model.BillingPlanResponse, error) {
    var res *model.BillingPlanResponse
    return res, op.Call.Do(ctx, &res)
}

// ListBillingPlans gets the list of available billing plans.
// SDK Method Billing::listBillingPlans
// https://docs.docusign.com/esign/restapi/Billing/BillingPlans/list
func (s *Service) ListBillingPlans() *ListBillingPlansCall {
    return &ListBillingPlansCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "GET",
            Path: "/v2/billing_plans",
            QueryOpts: make(url.Values),
        },
    }
}

// ListBillingPlansCall implements DocuSign API SDK Billing::listBillingPlans
type ListBillingPlansCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *ListBillingPlansCall) Do(ctx context.Context)  (*model.BillingPlansResponse, error) {
    var res *model.BillingPlansResponse
    return res, op.Call.Do(ctx, &res)
}

// PurchaseEnvelopes reserverd: Purchase additional envelopes.
// SDK Method Billing::purchaseEnvelopes
// https://docs.docusign.com/esign/restapi/Billing/BillingPlans/purchaseEnvelopes
func (s *Service) PurchaseEnvelopes(purchasedEnvelopesInformation *model.PurchasedEnvelopesInformation) *PurchaseEnvelopesCall {
    return &PurchaseEnvelopesCall{
        &esign.Call{
            Credential: s.credential,
    		Method:  "PUT",
            Path: "billing_plan/purchased_envelopes",
            Payload: purchasedEnvelopesInformation,
            QueryOpts: make(url.Values),
        },
    }
}

// PurchaseEnvelopesCall implements DocuSign API SDK Billing::purchaseEnvelopes
type PurchaseEnvelopesCall struct {
    *esign.Call
}

// Do executes the call.  A nil context will return error.
func (op *PurchaseEnvelopesCall) Do(ctx context.Context)  error {
    
    return op.Call.Do(ctx, nil)
}
