// This file was auto-generated by Fern from our API Definition.

package api

type BillCreateRequest struct {
	Minorversion *string `json:"-" url:"minorversion,omitempty"`
	Operation    *string `json:"-" url:"operation,omitempty"`
	Id           *string `json:"Id,omitempty" url:"-"`
	SyncToken    *string `json:"SyncToken,omitempty" url:"-"`
}

type BillGetbyidRequest struct {
	Minorversion *string `json:"-" url:"minorversion,omitempty"`
}