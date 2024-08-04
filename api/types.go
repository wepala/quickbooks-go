// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/wepala/quickbooks-go/api/core"
)

type BatchRequest struct {
	Minorversion     *string                             `json:"-" url:"minorversion,omitempty"`
	BatchItemRequest []*BatchRequestBatchItemRequestItem `json:"BatchItemRequest,omitempty" url:"-"`
}

type BatchRequestBatchItemRequestItem struct {
	Invoice      *BatchRequestBatchItemRequestItemInvoice      `json:"Invoice,omitempty" url:"Invoice,omitempty"`
	Query        *string                                       `json:"Query,omitempty" url:"Query,omitempty"`
	SalesReceipt *BatchRequestBatchItemRequestItemSalesReceipt `json:"SalesReceipt,omitempty" url:"SalesReceipt,omitempty"`
	Vendor       *BatchRequestBatchItemRequestItemVendor       `json:"Vendor,omitempty" url:"Vendor,omitempty"`
	BId          *string                                       `json:"bId,omitempty" url:"bId,omitempty"`
	Operation    *string                                       `json:"operation,omitempty" url:"operation,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BatchRequestBatchItemRequestItem) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchRequestBatchItemRequestItem) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchRequestBatchItemRequestItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchRequestBatchItemRequestItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchRequestBatchItemRequestItem) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type BatchRequestBatchItemRequestItemInvoice struct {
	Id        *string `json:"Id,omitempty" url:"Id,omitempty"`
	SyncToken *string `json:"SyncToken,omitempty" url:"SyncToken,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BatchRequestBatchItemRequestItemInvoice) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchRequestBatchItemRequestItemInvoice) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchRequestBatchItemRequestItemInvoice
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchRequestBatchItemRequestItemInvoice(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchRequestBatchItemRequestItemInvoice) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type BatchRequestBatchItemRequestItemSalesReceipt struct {
	Id          *string `json:"Id,omitempty" url:"Id,omitempty"`
	PrivateNote *string `json:"PrivateNote,omitempty" url:"PrivateNote,omitempty"`
	SyncToken   *string `json:"SyncToken,omitempty" url:"SyncToken,omitempty"`
	Domain      *string `json:"domain,omitempty" url:"domain,omitempty"`
	Sparse      *bool   `json:"sparse,omitempty" url:"sparse,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BatchRequestBatchItemRequestItemSalesReceipt) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchRequestBatchItemRequestItemSalesReceipt) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchRequestBatchItemRequestItemSalesReceipt
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchRequestBatchItemRequestItemSalesReceipt(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchRequestBatchItemRequestItemSalesReceipt) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type BatchRequestBatchItemRequestItemVendor struct {
	DisplayName *string `json:"DisplayName,omitempty" url:"DisplayName,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BatchRequestBatchItemRequestItemVendor) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchRequestBatchItemRequestItemVendor) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchRequestBatchItemRequestItemVendor
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchRequestBatchItemRequestItemVendor(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchRequestBatchItemRequestItemVendor) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type AttachableCreateRequestAttachableRefItemEntityRef struct {
	Type  *string `json:"type,omitempty" url:"type,omitempty"`
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *AttachableCreateRequestAttachableRefItemEntityRef) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AttachableCreateRequestAttachableRefItemEntityRef) UnmarshalJSON(data []byte) error {
	type unmarshaler AttachableCreateRequestAttachableRefItemEntityRef
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AttachableCreateRequestAttachableRefItemEntityRef(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AttachableCreateRequestAttachableRefItemEntityRef) String() string {
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

type BillpaymentCreateRequestCheckPaymentBankAccountRef struct {
	Name  *string `json:"name,omitempty" url:"name,omitempty"`
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BillpaymentCreateRequestCheckPaymentBankAccountRef) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BillpaymentCreateRequestCheckPaymentBankAccountRef) UnmarshalJSON(data []byte) error {
	type unmarshaler BillpaymentCreateRequestCheckPaymentBankAccountRef
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BillpaymentCreateRequestCheckPaymentBankAccountRef(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BillpaymentCreateRequestCheckPaymentBankAccountRef) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type BillpaymentCreateRequestLineItemLinkedTxnItem struct {
	TxnId   *string `json:"TxnId,omitempty" url:"TxnId,omitempty"`
	TxnType *string `json:"TxnType,omitempty" url:"TxnType,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (b *BillpaymentCreateRequestLineItemLinkedTxnItem) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BillpaymentCreateRequestLineItemLinkedTxnItem) UnmarshalJSON(data []byte) error {
	type unmarshaler BillpaymentCreateRequestLineItemLinkedTxnItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BillpaymentCreateRequestLineItemLinkedTxnItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties

	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BillpaymentCreateRequestLineItemLinkedTxnItem) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type PreferenceUpdateRequestCurrencyPrefsHomeCurrency struct {
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestCurrencyPrefsHomeCurrency) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestCurrencyPrefsHomeCurrency) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestCurrencyPrefsHomeCurrency
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestCurrencyPrefsHomeCurrency(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestCurrencyPrefsHomeCurrency) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestEmailMessagesPrefsEstimateMessage struct {
	Message *string `json:"Message,omitempty" url:"Message,omitempty"`
	Subject *string `json:"Subject,omitempty" url:"Subject,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsEstimateMessage) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsEstimateMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestEmailMessagesPrefsEstimateMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestEmailMessagesPrefsEstimateMessage(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsEstimateMessage) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestEmailMessagesPrefsInvoiceMessage struct {
	Message *string `json:"Message,omitempty" url:"Message,omitempty"`
	Subject *string `json:"Subject,omitempty" url:"Subject,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsInvoiceMessage) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsInvoiceMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestEmailMessagesPrefsInvoiceMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestEmailMessagesPrefsInvoiceMessage(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsInvoiceMessage) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestEmailMessagesPrefsSalesReceiptMessage struct {
	Message *string `json:"Message,omitempty" url:"Message,omitempty"`
	Subject *string `json:"Subject,omitempty" url:"Subject,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsSalesReceiptMessage) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsSalesReceiptMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestEmailMessagesPrefsSalesReceiptMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestEmailMessagesPrefsSalesReceiptMessage(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsSalesReceiptMessage) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestEmailMessagesPrefsStatementMessage struct {
	Message *string `json:"Message,omitempty" url:"Message,omitempty"`
	Subject *string `json:"Subject,omitempty" url:"Subject,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsStatementMessage) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsStatementMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestEmailMessagesPrefsStatementMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestEmailMessagesPrefsStatementMessage(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestEmailMessagesPrefsStatementMessage) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestOtherPrefsNameValueItem struct {
	Name  *string `json:"Name,omitempty" url:"Name,omitempty"`
	Value *string `json:"Value,omitempty" url:"Value,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestOtherPrefsNameValueItem) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestOtherPrefsNameValueItem) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestOtherPrefsNameValueItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestOtherPrefsNameValueItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestOtherPrefsNameValueItem) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestSalesFormsPrefsCustomFieldItem struct {
	CustomField []*PreferenceUpdateRequestSalesFormsPrefsCustomFieldItemCustomFieldItem `json:"CustomField,omitempty" url:"CustomField,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestSalesFormsPrefsCustomFieldItem) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestSalesFormsPrefsCustomFieldItem) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestSalesFormsPrefsCustomFieldItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestSalesFormsPrefsCustomFieldItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestSalesFormsPrefsCustomFieldItem) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestSalesFormsPrefsCustomFieldItemCustomFieldItem struct {
	BooleanValue *bool   `json:"BooleanValue,omitempty" url:"BooleanValue,omitempty"`
	Name         *string `json:"Name,omitempty" url:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" url:"Type,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestSalesFormsPrefsCustomFieldItemCustomFieldItem) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestSalesFormsPrefsCustomFieldItemCustomFieldItem) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestSalesFormsPrefsCustomFieldItemCustomFieldItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestSalesFormsPrefsCustomFieldItemCustomFieldItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestSalesFormsPrefsCustomFieldItemCustomFieldItem) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestSalesFormsPrefsDefaultTerms struct {
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestSalesFormsPrefsDefaultTerms) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestSalesFormsPrefsDefaultTerms) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestSalesFormsPrefsDefaultTerms
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestSalesFormsPrefsDefaultTerms(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestSalesFormsPrefsDefaultTerms) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestTaxPrefsTaxGroupCodeRef struct {
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestTaxPrefsTaxGroupCodeRef) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestTaxPrefsTaxGroupCodeRef) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestTaxPrefsTaxGroupCodeRef
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestTaxPrefsTaxGroupCodeRef(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestTaxPrefsTaxGroupCodeRef) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItem struct {
	CustomField []*PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItemCustomFieldItem `json:"CustomField,omitempty" url:"CustomField,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItem) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItem) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItem) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItemCustomFieldItem struct {
	BooleanValue *bool   `json:"BooleanValue,omitempty" url:"BooleanValue,omitempty"`
	Name         *string `json:"Name,omitempty" url:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" url:"Type,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItemCustomFieldItem) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItemCustomFieldItem) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItemCustomFieldItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItemCustomFieldItem(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItemCustomFieldItem) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}
