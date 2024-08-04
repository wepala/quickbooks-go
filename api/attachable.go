// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/wepala/quickbooks-go/api/core"
)

type AttachableCreateRequest struct {
	Minorversion  *string                                     `json:"-" url:"minorversion,omitempty"`
	Operation     *string                                     `json:"-" url:"operation,omitempty"`
	AttachableRef []*AttachableCreateRequestAttachableRefItem `json:"AttachableRef,omitempty" url:"-"`
	Id            *string                                     `json:"Id,omitempty" url:"-"`
	MetaData      *AttachableCreateRequestMetaData            `json:"MetaData,omitempty" url:"-"`
	Note          *string                                     `json:"Note,omitempty" url:"-"`
	SyncToken     *string                                     `json:"SyncToken,omitempty" url:"-"`
	Domain        *string                                     `json:"domain,omitempty" url:"-"`
	Sparse        *bool                                       `json:"sparse,omitempty" url:"-"`
}

type AttachableReadbyidRequest struct {
	Minorversion *string `json:"-" url:"minorversion,omitempty"`
}

type AttachableCreateRequestAttachableRefItem struct {
	EntityRef     *AttachableCreateRequestAttachableRefItemEntityRef `json:"EntityRef,omitempty" url:"EntityRef,omitempty"`
	IncludeOnSend *bool                                              `json:"IncludeOnSend,omitempty" url:"IncludeOnSend,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AttachableCreateRequestAttachableRefItem) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AttachableCreateRequestAttachableRefItem) UnmarshalJSON(data []byte) error {
	type unmarshaler AttachableCreateRequestAttachableRefItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AttachableCreateRequestAttachableRefItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AttachableCreateRequestAttachableRefItem) String() string {
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

type AttachableCreateRequestMetaData struct {
	CreateTime      *string `json:"CreateTime,omitempty" url:"CreateTime,omitempty"`
	LastUpdatedTime *string `json:"LastUpdatedTime,omitempty" url:"LastUpdatedTime,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AttachableCreateRequestMetaData) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AttachableCreateRequestMetaData) UnmarshalJSON(data []byte) error {
	type unmarshaler AttachableCreateRequestMetaData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AttachableCreateRequestMetaData(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AttachableCreateRequestMetaData) String() string {
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