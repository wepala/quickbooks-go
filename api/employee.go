// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/wepala/quickbooks-go/api/core"
)

type EmployeeCreateRequest struct {
	Active           *bool                              `json:"Active,omitempty" url:"-"`
	BillableTime     *bool                              `json:"BillableTime,omitempty" url:"-"`
	DisplayName      *string                            `json:"DisplayName,omitempty" url:"-"`
	FamilyName       *string                            `json:"FamilyName,omitempty" url:"-"`
	GivenName        *string                            `json:"GivenName,omitempty" url:"-"`
	Id               *string                            `json:"Id,omitempty" url:"-"`
	MetaData         *EmployeeCreateRequestMetaData     `json:"MetaData,omitempty" url:"-"`
	PrimaryAddr      *EmployeeCreateRequestPrimaryAddr  `json:"PrimaryAddr,omitempty" url:"-"`
	PrimaryPhone     *EmployeeCreateRequestPrimaryPhone `json:"PrimaryPhone,omitempty" url:"-"`
	PrintOnCheckName *string                            `json:"PrintOnCheckName,omitempty" url:"-"`
	Ssn              *string                            `json:"SSN,omitempty" url:"-"`
	SyncToken        *string                            `json:"SyncToken,omitempty" url:"-"`
	Domain           *string                            `json:"domain,omitempty" url:"-"`
	Sparse           *bool                              `json:"sparse,omitempty" url:"-"`
}

type EmployeeCreateRequestMetaData struct {
	CreateTime      *string `json:"CreateTime,omitempty" url:"CreateTime,omitempty"`
	LastUpdatedTime *string `json:"LastUpdatedTime,omitempty" url:"LastUpdatedTime,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EmployeeCreateRequestMetaData) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EmployeeCreateRequestMetaData) UnmarshalJSON(data []byte) error {
	type unmarshaler EmployeeCreateRequestMetaData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EmployeeCreateRequestMetaData(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EmployeeCreateRequestMetaData) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EmployeeCreateRequestPrimaryAddr struct {
	City                   *string `json:"City,omitempty" url:"City,omitempty"`
	CountrySubDivisionCode *string `json:"CountrySubDivisionCode,omitempty" url:"CountrySubDivisionCode,omitempty"`
	Id                     *string `json:"Id,omitempty" url:"Id,omitempty"`
	Line1                  *string `json:"Line1,omitempty" url:"Line1,omitempty"`
	PostalCode             *string `json:"PostalCode,omitempty" url:"PostalCode,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EmployeeCreateRequestPrimaryAddr) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EmployeeCreateRequestPrimaryAddr) UnmarshalJSON(data []byte) error {
	type unmarshaler EmployeeCreateRequestPrimaryAddr
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EmployeeCreateRequestPrimaryAddr(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EmployeeCreateRequestPrimaryAddr) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EmployeeCreateRequestPrimaryPhone struct {
	FreeFormNumber *string `json:"FreeFormNumber,omitempty" url:"FreeFormNumber,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EmployeeCreateRequestPrimaryPhone) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EmployeeCreateRequestPrimaryPhone) UnmarshalJSON(data []byte) error {
	type unmarshaler EmployeeCreateRequestPrimaryPhone
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EmployeeCreateRequestPrimaryPhone(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EmployeeCreateRequestPrimaryPhone) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}
