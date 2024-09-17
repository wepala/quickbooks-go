// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/wepala/quickbooks-go/api/core"
)

type AccountReadallRequest struct {
	Query *string `json:"-" url:"query,omitempty"`
}

type AccountReadallResponse struct {
	QueryResponse *AccountReadallResponseQueryResponse `json:"QueryResponse,omitempty" url:"QueryResponse,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AccountReadallResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AccountReadallResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AccountReadallResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccountReadallResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccountReadallResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}
