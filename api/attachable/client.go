// This file was auto-generated by Fern from our API Definition.

package attachable

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

// Create an attachable object
// Conent-Type:application/json
// Method - POST
func (c *Client) Create(
	ctx context.Context,
	companyid string,
	request *api.AttachableCreateRequest,
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
	endpointURL := core.EncodeURL(baseURL+"/v3/company/%v/attachable", companyid)

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

// Retrieve an attachable object by Id
// Accept:application/json
// Method - GET
func (c *Client) Readbyid(
	ctx context.Context,
	companyid string,
	request *api.AttachableReadbyidRequest,
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
	endpointURL := core.EncodeURL(baseURL+"/v3/company/%v/attachable/5000000000000029383", companyid)

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

// Uploading and linking new attachments
//
// https://developer.intuit.com/docs/0100_quickbooks_online/0200_dev_guides/accounting/attachments#Uploading_and_linking_new_attachments
//
// If the attachment is not in the Attachment list already, it's possible to upload it and link it to the object in one multipart operation.
//
// Operation: POST https://quickbooks.api.intuit.com/v3/company/companyID/upload
// Content type: multipart/form-data
//
// # Request body
//
// The following sample code shows the multipart request body for uploading a file and its supporting Attachable metatdata object, with the result of it being both added to the Attachment list and added to the object.
//
// The Attachable object accompanying this request supplies metadata and object information to which the attachment is linked.
// Each part of the multipart request is separated by a boundary. In the sample below, the string --YOjcLaTlykb6OxfYJx4O07j1MweeMFem is used. You can use any random and unique string.
// The file to be uploaded and its Attachable object are paired together via the name parameter in the part header for each one.
// The name parameter for the file part is of the form file_content_nn, where nn is a unique index number among the set of files being uploaded.
// The name parameter for the Attachable object is of the form file_metadata_nn, where nn corresponds to the file index number used with the content .
// The file or files are stored in the Attachment list with the name specified by the filename parameter.
// If the data supplied with the Attachable object cannot be validated, an error is returned and the file is not uploaded.
func (c *Client) UploadAttachments(
	ctx context.Context,
	companyid string,
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
	endpointURL := core.EncodeURL(baseURL+"/v3/company/%v/upload", companyid)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}
