// This file was auto-generated by Fern from our API Definition.

package api

type TermCreateRequest struct {
	Minorversion    *string  `json:"-" url:"minorversion,omitempty"`
	Active          *bool    `json:"Active,omitempty" url:"-"`
	DiscountPercent *float64 `json:"DiscountPercent,omitempty" url:"-"`
	DueDays         *float64 `json:"DueDays,omitempty" url:"-"`
	Id              *string  `json:"Id,omitempty" url:"-"`
	Name            *string  `json:"Name,omitempty" url:"-"`
	SyncToken       *string  `json:"SyncToken,omitempty" url:"-"`
	Type            *string  `json:"Type,omitempty" url:"-"`
	Domain          *string  `json:"domain,omitempty" url:"-"`
	Sparse          *bool    `json:"sparse,omitempty" url:"-"`
}

type TermReadbyidRequest struct {
	Minorversion *string `json:"-" url:"minorversion,omitempty"`
}