// This file was auto-generated by Fern from our API Definition.

package journalentry

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

// Create an journalentry object
// Method : POST
func (c *Client) Create(
	ctx context.Context,
	companyid string,
	request *api.JournalentryCreateRequest,
	opts ...option.RequestOption,
) (*api.JournalentryCreateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://quickbooks.api.intuit.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v3/company/%v/journalentry", companyid)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.JournalentryCreateResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Read an journalentry object by Id
// Method : POST
func (c *Client) Readbyid(
	ctx context.Context,
	companyid string,
	id string,
	opts ...option.RequestOption,
) (*api.JournalEntry, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://quickbooks.api.intuit.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/v3/company/%v/journalentry/%v",
		companyid,
		id,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *api.JournalEntry
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
