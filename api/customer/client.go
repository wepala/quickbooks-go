// This file was auto-generated by Fern from our API Definition.

package customer

import (
	context "context"
	api "github.com/wepala/quickbooks-go/api"
	core "github.com/wepala/quickbooks-go/api/core"
	option "github.com/wepala/quickbooks-go/api/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

// Create a customer
// Method : POST
func (c *Client) Create(
	ctx context.Context,
	companyid string,
	request *api.CustomerCreateRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://{{baseurl}}"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v3/company/%v/customer", companyid)

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

// Read a customer entry by Id
// Method : GET
func (c *Client) Readbyid(
	ctx context.Context,
	companyid string,
	request *api.CustomerReadbyidRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://{{baseurl}}"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v3/company/%v/customer/63", companyid)

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
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}
