// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/wepala/quickbooks-go/api/core"
)

type PreferenceReadRequest struct {
	Minorversion *string `json:"-" url:"minorversion,omitempty"`
}

type PreferenceUpdateRequest struct {
	Minorversion            *string                                         `json:"-" url:"minorversion,omitempty"`
	AccountingInfoPrefs     *PreferenceUpdateRequestAccountingInfoPrefs     `json:"AccountingInfoPrefs,omitempty" url:"-"`
	CurrencyPrefs           *PreferenceUpdateRequestCurrencyPrefs           `json:"CurrencyPrefs,omitempty" url:"-"`
	EmailMessagesPrefs      *PreferenceUpdateRequestEmailMessagesPrefs      `json:"EmailMessagesPrefs,omitempty" url:"-"`
	Id                      *string                                         `json:"Id,omitempty" url:"-"`
	MetaData                *PreferenceUpdateRequestMetaData                `json:"MetaData,omitempty" url:"-"`
	OtherPrefs              *PreferenceUpdateRequestOtherPrefs              `json:"OtherPrefs,omitempty" url:"-"`
	ProductAndServicesPrefs *PreferenceUpdateRequestProductAndServicesPrefs `json:"ProductAndServicesPrefs,omitempty" url:"-"`
	ReportPrefs             *PreferenceUpdateRequestReportPrefs             `json:"ReportPrefs,omitempty" url:"-"`
	SalesFormsPrefs         *PreferenceUpdateRequestSalesFormsPrefs         `json:"SalesFormsPrefs,omitempty" url:"-"`
	SyncToken               *string                                         `json:"SyncToken,omitempty" url:"-"`
	TaxPrefs                *PreferenceUpdateRequestTaxPrefs                `json:"TaxPrefs,omitempty" url:"-"`
	TimeTrackingPrefs       *PreferenceUpdateRequestTimeTrackingPrefs       `json:"TimeTrackingPrefs,omitempty" url:"-"`
	VendorAndPurchasesPrefs *PreferenceUpdateRequestVendorAndPurchasesPrefs `json:"VendorAndPurchasesPrefs,omitempty" url:"-"`
	Domain                  *string                                         `json:"domain,omitempty" url:"-"`
	Sparse                  *bool                                           `json:"sparse,omitempty" url:"-"`
}

type PreferenceUpdateRequestAccountingInfoPrefs struct {
	ClassTrackingPerTxn     *bool   `json:"ClassTrackingPerTxn,omitempty" url:"ClassTrackingPerTxn,omitempty"`
	ClassTrackingPerTxnLine *bool   `json:"ClassTrackingPerTxnLine,omitempty" url:"ClassTrackingPerTxnLine,omitempty"`
	CustomerTerminology     *string `json:"CustomerTerminology,omitempty" url:"CustomerTerminology,omitempty"`
	DepartmentTerminology   *string `json:"DepartmentTerminology,omitempty" url:"DepartmentTerminology,omitempty"`
	TrackDepartments        *bool   `json:"TrackDepartments,omitempty" url:"TrackDepartments,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestAccountingInfoPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestAccountingInfoPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestAccountingInfoPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestAccountingInfoPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestAccountingInfoPrefs) String() string {
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

type PreferenceUpdateRequestCurrencyPrefs struct {
	HomeCurrency         *PreferenceUpdateRequestCurrencyPrefsHomeCurrency `json:"HomeCurrency,omitempty" url:"HomeCurrency,omitempty"`
	MultiCurrencyEnabled *bool                                             `json:"MultiCurrencyEnabled,omitempty" url:"MultiCurrencyEnabled,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestCurrencyPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestCurrencyPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestCurrencyPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestCurrencyPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestCurrencyPrefs) String() string {
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

type PreferenceUpdateRequestEmailMessagesPrefs struct {
	EstimateMessage     *PreferenceUpdateRequestEmailMessagesPrefsEstimateMessage     `json:"EstimateMessage,omitempty" url:"EstimateMessage,omitempty"`
	InvoiceMessage      *PreferenceUpdateRequestEmailMessagesPrefsInvoiceMessage      `json:"InvoiceMessage,omitempty" url:"InvoiceMessage,omitempty"`
	SalesReceiptMessage *PreferenceUpdateRequestEmailMessagesPrefsSalesReceiptMessage `json:"SalesReceiptMessage,omitempty" url:"SalesReceiptMessage,omitempty"`
	StatementMessage    *PreferenceUpdateRequestEmailMessagesPrefsStatementMessage    `json:"StatementMessage,omitempty" url:"StatementMessage,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestEmailMessagesPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestEmailMessagesPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestEmailMessagesPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestEmailMessagesPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestEmailMessagesPrefs) String() string {
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

type PreferenceUpdateRequestMetaData struct {
	CreateTime      *string `json:"CreateTime,omitempty" url:"CreateTime,omitempty"`
	LastUpdatedTime *string `json:"LastUpdatedTime,omitempty" url:"LastUpdatedTime,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestMetaData) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestMetaData) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestMetaData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestMetaData(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestMetaData) String() string {
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

type PreferenceUpdateRequestOtherPrefs struct {
	NameValue []*PreferenceUpdateRequestOtherPrefsNameValueItem `json:"NameValue,omitempty" url:"NameValue,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestOtherPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestOtherPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestOtherPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestOtherPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestOtherPrefs) String() string {
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

type PreferenceUpdateRequestProductAndServicesPrefs struct {
	ForPurchase              *bool `json:"ForPurchase,omitempty" url:"ForPurchase,omitempty"`
	ForSales                 *bool `json:"ForSales,omitempty" url:"ForSales,omitempty"`
	QuantityOnHand           *bool `json:"QuantityOnHand,omitempty" url:"QuantityOnHand,omitempty"`
	QuantityWithPriceAndRate *bool `json:"QuantityWithPriceAndRate,omitempty" url:"QuantityWithPriceAndRate,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestProductAndServicesPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestProductAndServicesPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestProductAndServicesPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestProductAndServicesPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestProductAndServicesPrefs) String() string {
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

type PreferenceUpdateRequestReportPrefs struct {
	CalcAgingReportFromTxnDate *bool   `json:"CalcAgingReportFromTxnDate,omitempty" url:"CalcAgingReportFromTxnDate,omitempty"`
	ReportBasis                *string `json:"ReportBasis,omitempty" url:"ReportBasis,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestReportPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestReportPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestReportPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestReportPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestReportPrefs) String() string {
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

type PreferenceUpdateRequestSalesFormsPrefs struct {
	AllowDeposit               *bool                                                    `json:"AllowDeposit,omitempty" url:"AllowDeposit,omitempty"`
	AllowDiscount              *bool                                                    `json:"AllowDiscount,omitempty" url:"AllowDiscount,omitempty"`
	AllowEstimates             *bool                                                    `json:"AllowEstimates,omitempty" url:"AllowEstimates,omitempty"`
	AllowServiceDate           *bool                                                    `json:"AllowServiceDate,omitempty" url:"AllowServiceDate,omitempty"`
	AllowShipping              *bool                                                    `json:"AllowShipping,omitempty" url:"AllowShipping,omitempty"`
	CustomField                []*PreferenceUpdateRequestSalesFormsPrefsCustomFieldItem `json:"CustomField,omitempty" url:"CustomField,omitempty"`
	CustomTxnNumbers           *bool                                                    `json:"CustomTxnNumbers,omitempty" url:"CustomTxnNumbers,omitempty"`
	DefaultCustomerMessage     *string                                                  `json:"DefaultCustomerMessage,omitempty" url:"DefaultCustomerMessage,omitempty"`
	DefaultDiscountAccount     *string                                                  `json:"DefaultDiscountAccount,omitempty" url:"DefaultDiscountAccount,omitempty"`
	DefaultTerms               *PreferenceUpdateRequestSalesFormsPrefsDefaultTerms      `json:"DefaultTerms,omitempty" url:"DefaultTerms,omitempty"`
	ETransactionAttachPdf      *bool                                                    `json:"ETransactionAttachPDF,omitempty" url:"ETransactionAttachPDF,omitempty"`
	ETransactionEnabledStatus  *string                                                  `json:"ETransactionEnabledStatus,omitempty" url:"ETransactionEnabledStatus,omitempty"`
	ETransactionPaymentEnabled *bool                                                    `json:"ETransactionPaymentEnabled,omitempty" url:"ETransactionPaymentEnabled,omitempty"`
	IpnSupportEnabled          *bool                                                    `json:"IPNSupportEnabled,omitempty" url:"IPNSupportEnabled,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestSalesFormsPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestSalesFormsPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestSalesFormsPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestSalesFormsPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestSalesFormsPrefs) String() string {
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

type PreferenceUpdateRequestTaxPrefs struct {
	TaxGroupCodeRef *PreferenceUpdateRequestTaxPrefsTaxGroupCodeRef `json:"TaxGroupCodeRef,omitempty" url:"TaxGroupCodeRef,omitempty"`
	UsingSalesTax   *bool                                           `json:"UsingSalesTax,omitempty" url:"UsingSalesTax,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestTaxPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestTaxPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestTaxPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestTaxPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestTaxPrefs) String() string {
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

type PreferenceUpdateRequestTimeTrackingPrefs struct {
	BillCustomers           *bool   `json:"BillCustomers,omitempty" url:"BillCustomers,omitempty"`
	MarkTimeEntriesBillable *bool   `json:"MarkTimeEntriesBillable,omitempty" url:"MarkTimeEntriesBillable,omitempty"`
	ShowBillRateToAll       *bool   `json:"ShowBillRateToAll,omitempty" url:"ShowBillRateToAll,omitempty"`
	UseServices             *bool   `json:"UseServices,omitempty" url:"UseServices,omitempty"`
	WorkWeekStartDate       *string `json:"WorkWeekStartDate,omitempty" url:"WorkWeekStartDate,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestTimeTrackingPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestTimeTrackingPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestTimeTrackingPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestTimeTrackingPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestTimeTrackingPrefs) String() string {
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

type PreferenceUpdateRequestVendorAndPurchasesPrefs struct {
	BillableExpenseTracking *bool                                                              `json:"BillableExpenseTracking,omitempty" url:"BillableExpenseTracking,omitempty"`
	PoCustomField           []*PreferenceUpdateRequestVendorAndPurchasesPrefsPoCustomFieldItem `json:"POCustomField,omitempty" url:"POCustomField,omitempty"`
	TrackingByCustomer      *bool                                                              `json:"TrackingByCustomer,omitempty" url:"TrackingByCustomer,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefs) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefs) UnmarshalJSON(data []byte) error {
	type unmarshaler PreferenceUpdateRequestVendorAndPurchasesPrefs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreferenceUpdateRequestVendorAndPurchasesPrefs(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreferenceUpdateRequestVendorAndPurchasesPrefs) String() string {
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
