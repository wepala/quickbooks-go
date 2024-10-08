// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

type InvoiceCreateRequest struct {
	Operation    *string  `json:"-" url:"operation,omitempty"`
	Minorversion *string  `json:"-" url:"minorversion,omitempty"`
	Body         *Invoice `json:"-" url:"-"`
}

func (i *InvoiceCreateRequest) UnmarshalJSON(data []byte) error {
	body := new(Invoice)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	i.Body = body
	return nil
}

func (i *InvoiceCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Body)
}
