// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/wepala/quickbooks-go/api/core"
)

type VendorCreateRequest struct {
	AcctNum          *string                              `json:"AcctNum,omitempty" url:"-"`
	Active           *bool                                `json:"Active,omitempty" url:"-"`
	Balance          *float64                             `json:"Balance,omitempty" url:"-"`
	BillAddr         *VendorCreateRequestBillAddr         `json:"BillAddr,omitempty" url:"-"`
	CompanyName      *string                              `json:"CompanyName,omitempty" url:"-"`
	CurrencyRef      *VendorCreateRequestCurrencyRef      `json:"CurrencyRef,omitempty" url:"-"`
	DisplayName      *string                              `json:"DisplayName,omitempty" url:"-"`
	FamilyName       *string                              `json:"FamilyName,omitempty" url:"-"`
	GivenName        *string                              `json:"GivenName,omitempty" url:"-"`
	Id               *string                              `json:"Id,omitempty" url:"-"`
	Mobile           *VendorCreateRequestMobile           `json:"Mobile,omitempty" url:"-"`
	PrimaryEmailAddr *VendorCreateRequestPrimaryEmailAddr `json:"PrimaryEmailAddr,omitempty" url:"-"`
	PrimaryPhone     *VendorCreateRequestPrimaryPhone     `json:"PrimaryPhone,omitempty" url:"-"`
	PrintOnCheckName *string                              `json:"PrintOnCheckName,omitempty" url:"-"`
	Suffix           *string                              `json:"Suffix,omitempty" url:"-"`
	SyncToken        *string                              `json:"SyncToken,omitempty" url:"-"`
	TaxIdentifier    *string                              `json:"TaxIdentifier,omitempty" url:"-"`
	Title            *string                              `json:"Title,omitempty" url:"-"`
	Vendor1099       *bool                                `json:"Vendor1099,omitempty" url:"-"`
	WebAddr          *VendorCreateRequestWebAddr          `json:"WebAddr,omitempty" url:"-"`
	Domain           *string                              `json:"domain,omitempty" url:"-"`
	Sparse           *bool                                `json:"sparse,omitempty" url:"-"`
}

type VendorCreateRequestBillAddr struct {
	City                   *string `json:"City,omitempty" url:"City,omitempty"`
	Country                *string `json:"Country,omitempty" url:"Country,omitempty"`
	CountrySubDivisionCode *string `json:"CountrySubDivisionCode,omitempty" url:"CountrySubDivisionCode,omitempty"`
	Id                     *string `json:"Id,omitempty" url:"Id,omitempty"`
	Line1                  *string `json:"Line1,omitempty" url:"Line1,omitempty"`
	Line2                  *string `json:"Line2,omitempty" url:"Line2,omitempty"`
	Line3                  *string `json:"Line3,omitempty" url:"Line3,omitempty"`
	PostalCode             *string `json:"PostalCode,omitempty" url:"PostalCode,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VendorCreateRequestBillAddr) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VendorCreateRequestBillAddr) UnmarshalJSON(data []byte) error {
	type unmarshaler VendorCreateRequestBillAddr
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VendorCreateRequestBillAddr(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VendorCreateRequestBillAddr) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type VendorCreateRequestCurrencyRef struct {
	Name  *string `json:"name,omitempty" url:"name,omitempty"`
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VendorCreateRequestCurrencyRef) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VendorCreateRequestCurrencyRef) UnmarshalJSON(data []byte) error {
	type unmarshaler VendorCreateRequestCurrencyRef
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VendorCreateRequestCurrencyRef(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VendorCreateRequestCurrencyRef) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type VendorCreateRequestMobile struct {
	FreeFormNumber *string `json:"FreeFormNumber,omitempty" url:"FreeFormNumber,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VendorCreateRequestMobile) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VendorCreateRequestMobile) UnmarshalJSON(data []byte) error {
	type unmarshaler VendorCreateRequestMobile
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VendorCreateRequestMobile(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VendorCreateRequestMobile) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type VendorCreateRequestPrimaryEmailAddr struct {
	Address *string `json:"Address,omitempty" url:"Address,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VendorCreateRequestPrimaryEmailAddr) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VendorCreateRequestPrimaryEmailAddr) UnmarshalJSON(data []byte) error {
	type unmarshaler VendorCreateRequestPrimaryEmailAddr
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VendorCreateRequestPrimaryEmailAddr(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VendorCreateRequestPrimaryEmailAddr) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type VendorCreateRequestPrimaryPhone struct {
	FreeFormNumber *string `json:"FreeFormNumber,omitempty" url:"FreeFormNumber,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VendorCreateRequestPrimaryPhone) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VendorCreateRequestPrimaryPhone) UnmarshalJSON(data []byte) error {
	type unmarshaler VendorCreateRequestPrimaryPhone
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VendorCreateRequestPrimaryPhone(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VendorCreateRequestPrimaryPhone) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type VendorCreateRequestWebAddr struct {
	Uri *string `json:"URI,omitempty" url:"URI,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VendorCreateRequestWebAddr) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VendorCreateRequestWebAddr) UnmarshalJSON(data []byte) error {
	type unmarshaler VendorCreateRequestWebAddr
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VendorCreateRequestWebAddr(value)

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VendorCreateRequestWebAddr) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}
