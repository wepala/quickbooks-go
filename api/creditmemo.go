// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/wepala/quickbooks-go/api/core"
)

type CreditmemoCreateRequest struct {
	Operation *string     `json:"-" url:"operation,omitempty"`
	Body      *CreditMemo `json:"-" url:"-"`
}

func (c *CreditmemoCreateRequest) UnmarshalJSON(data []byte) error {
	body := new(CreditMemo)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.Body = body
	return nil
}

func (c *CreditmemoCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

type CreditmemoCreateResponse struct {
	CreditMemo *CreditMemo `json:"CreditMemo,omitempty" url:"CreditMemo,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreditmemoCreateResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreditmemoCreateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreditmemoCreateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreditmemoCreateResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreditmemoCreateResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}
