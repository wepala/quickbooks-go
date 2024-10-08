// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

type SalesreceiptCreateRequest struct {
	Operation *string  `json:"-" url:"operation,omitempty"`
	Include   *string  `json:"-" url:"include,omitempty"`
	Body      *Invoice `json:"-" url:"-"`
}

func (s *SalesreceiptCreateRequest) UnmarshalJSON(data []byte) error {
	body := new(Invoice)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.Body = body
	return nil
}

func (s *SalesreceiptCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}
