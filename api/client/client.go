// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	api "github.com/wepala/quickbooks-go/api"
	account "github.com/wepala/quickbooks-go/api/account"
	attachable "github.com/wepala/quickbooks-go/api/attachable"
	bill "github.com/wepala/quickbooks-go/api/bill"
	billpayment "github.com/wepala/quickbooks-go/api/billpayment"
	changedatacapture "github.com/wepala/quickbooks-go/api/changedatacapture"
	class "github.com/wepala/quickbooks-go/api/class"
	companyinfo "github.com/wepala/quickbooks-go/api/companyinfo"
	core "github.com/wepala/quickbooks-go/api/core"
	creditmemo "github.com/wepala/quickbooks-go/api/creditmemo"
	customer "github.com/wepala/quickbooks-go/api/customer"
	department "github.com/wepala/quickbooks-go/api/department"
	deposit "github.com/wepala/quickbooks-go/api/deposit"
	employee "github.com/wepala/quickbooks-go/api/employee"
	estimate "github.com/wepala/quickbooks-go/api/estimate"
	exchangerate "github.com/wepala/quickbooks-go/api/exchangerate"
	invoice "github.com/wepala/quickbooks-go/api/invoice"
	item "github.com/wepala/quickbooks-go/api/item"
	journalentry "github.com/wepala/quickbooks-go/api/journalentry"
	option "github.com/wepala/quickbooks-go/api/option"
	payment "github.com/wepala/quickbooks-go/api/payment"
	paymentmethod "github.com/wepala/quickbooks-go/api/paymentmethod"
	preferences "github.com/wepala/quickbooks-go/api/preferences"
	purchase "github.com/wepala/quickbooks-go/api/purchase"
	purchaseorder "github.com/wepala/quickbooks-go/api/purchaseorder"
	query "github.com/wepala/quickbooks-go/api/query"
	refundreceipt "github.com/wepala/quickbooks-go/api/refundreceipt"
	reports "github.com/wepala/quickbooks-go/api/reports"
	salesreceipt "github.com/wepala/quickbooks-go/api/salesreceipt"
	taxagency "github.com/wepala/quickbooks-go/api/taxagency"
	taxcode "github.com/wepala/quickbooks-go/api/taxcode"
	taxrate "github.com/wepala/quickbooks-go/api/taxrate"
	taxservice "github.com/wepala/quickbooks-go/api/taxservice"
	term "github.com/wepala/quickbooks-go/api/term"
	timeactivity "github.com/wepala/quickbooks-go/api/timeactivity"
	transfer "github.com/wepala/quickbooks-go/api/transfer"
	vendor "github.com/wepala/quickbooks-go/api/vendor"
	vendorcredit "github.com/wepala/quickbooks-go/api/vendorcredit"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Account           *account.Client
	VendorCredit      *vendorcredit.Client
	Attachable        *attachable.Client
	Bill              *bill.Client
	Billpayment       *billpayment.Client
	ChangeDataCapture *changedatacapture.Client
	Class             *class.Client
	Companyinfo       *companyinfo.Client
	Creditmemo        *creditmemo.Client
	Customer          *customer.Client
	Department        *department.Client
	Deposit           *deposit.Client
	Employee          *employee.Client
	Estimate          *estimate.Client
	Exchangerate      *exchangerate.Client
	Invoice           *invoice.Client
	Item              *item.Client
	Journalentry      *journalentry.Client
	Payment           *payment.Client
	Paymentmethod     *paymentmethod.Client
	Preferences       *preferences.Client
	Purchase          *purchase.Client
	Purchaseorder     *purchaseorder.Client
	Query             *query.Client
	Refundreceipt     *refundreceipt.Client
	Reports           *reports.Client
	Salesreceipt      *salesreceipt.Client
	Taxagency         *taxagency.Client
	Taxcode           *taxcode.Client
	Taxrate           *taxrate.Client
	Taxservice        *taxservice.Client
	Term              *term.Client
	Timeactivity      *timeactivity.Client
	Transfer          *transfer.Client
	Vendor            *vendor.Client
	Vendorcredit      *vendorcredit.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:            options.ToHeader(),
		Account:           account.NewClient(opts...),
		VendorCredit:      vendorcredit.NewClient(opts...),
		Attachable:        attachable.NewClient(opts...),
		Bill:              bill.NewClient(opts...),
		Billpayment:       billpayment.NewClient(opts...),
		ChangeDataCapture: changedatacapture.NewClient(opts...),
		Class:             class.NewClient(opts...),
		Companyinfo:       companyinfo.NewClient(opts...),
		Creditmemo:        creditmemo.NewClient(opts...),
		Customer:          customer.NewClient(opts...),
		Department:        department.NewClient(opts...),
		Deposit:           deposit.NewClient(opts...),
		Employee:          employee.NewClient(opts...),
		Estimate:          estimate.NewClient(opts...),
		Exchangerate:      exchangerate.NewClient(opts...),
		Invoice:           invoice.NewClient(opts...),
		Item:              item.NewClient(opts...),
		Journalentry:      journalentry.NewClient(opts...),
		Payment:           payment.NewClient(opts...),
		Paymentmethod:     paymentmethod.NewClient(opts...),
		Preferences:       preferences.NewClient(opts...),
		Purchase:          purchase.NewClient(opts...),
		Purchaseorder:     purchaseorder.NewClient(opts...),
		Query:             query.NewClient(opts...),
		Refundreceipt:     refundreceipt.NewClient(opts...),
		Reports:           reports.NewClient(opts...),
		Salesreceipt:      salesreceipt.NewClient(opts...),
		Taxagency:         taxagency.NewClient(opts...),
		Taxcode:           taxcode.NewClient(opts...),
		Taxrate:           taxrate.NewClient(opts...),
		Taxservice:        taxservice.NewClient(opts...),
		Term:              term.NewClient(opts...),
		Timeactivity:      timeactivity.NewClient(opts...),
		Transfer:          transfer.NewClient(opts...),
		Vendor:            vendor.NewClient(opts...),
		Vendorcredit:      vendorcredit.NewClient(opts...),
	}
}

// Multiple operations using batch query
// Content-Type:application/json
// Method - POST
func (c *Client) Batch(
	ctx context.Context,
	companyid string,
	request *api.BatchRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://quickbooks.api.intuit.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v3/company/%v/batch", companyid)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
		},
	); err != nil {
		return err
	}
	return nil
}
