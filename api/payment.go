// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

type PaymentCreateRequest struct {
	Operation *string  `json:"-" url:"operation,omitempty"`
	Body      *Payment `json:"-" url:"-"`
}

func (p *PaymentCreateRequest) UnmarshalJSON(data []byte) error {
	body := new(Payment)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	p.Body = body
	return nil
}

func (p *PaymentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Body)
}
