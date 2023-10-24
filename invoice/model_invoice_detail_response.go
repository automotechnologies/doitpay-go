/*
Doitpay API

This is the payment gateway API server for Doitpay.

API version: 1.0
Contact: support@doitpay.co.id
*/

package invoice

import (
	"encoding/json"

	"github.com/automotechnologies/doitpay-go/utils"
)

// checks if the InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse{}

// InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse struct for InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse
type InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse struct {
	Amount               *float32                                                  `json:"amount,omitempty"`
	AmountCurrency       *string                                                   `json:"amount_currency,omitempty"`
	Country              *string                                                   `json:"country,omitempty"`
	Customer             *InternalWebControllersMerchantApiv1InvoiceCustomer       `json:"customer,omitempty"`
	CustomerNotification []string                                                  `json:"customer_notification,omitempty"`
	ExternalId           *string                                                   `json:"external_id,omitempty"`
	InvoiceLinkId        *string                                                   `json:"invoice_link_id,omitempty"`
	Items                []InternalWebControllersMerchantApiv1InvoiceItem          `json:"items,omitempty"`
	MerchantLogo         *string                                                   `json:"merchant_logo,omitempty"`
	MerchantName         *string                                                   `json:"merchant_name,omitempty"`
	Notes                *string                                                   `json:"notes,omitempty"`
	PaymentMethods       []InternalWebControllersMerchantApiv1InvoicePaymentMethod `json:"payment_methods,omitempty"`
	RedirectUrl          *InternalWebControllersMerchantApiv1InvoiceRedirectURL    `json:"redirect_url,omitempty"`
	Status               *string                                                   `json:"status,omitempty"`
	TimeBegin            *string                                                   `json:"time_begin,omitempty"`
	TimeEnd              *string                                                   `json:"time_end,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse() *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse {
	this := InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponseWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponseWithDefaults() *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse {
	this := InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmount() float32 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmountOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetAmountCurrency returns the AmountCurrency field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmountCurrency() string {
	if o == nil || utils.IsNil(o.AmountCurrency) {
		var ret string
		return ret
	}
	return *o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmountCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AmountCurrency) {
		return nil, false
	}
	return o.AmountCurrency, true
}

// HasAmountCurrency returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasAmountCurrency() bool {
	if o != nil && !utils.IsNil(o.AmountCurrency) {
		return true
	}

	return false
}

// SetAmountCurrency gets a reference to the given string and assigns it to the AmountCurrency field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetAmountCurrency(v string) {
	o.AmountCurrency = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetCountry(v string) {
	o.Country = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomer() InternalWebControllersMerchantApiv1InvoiceCustomer {
	if o == nil || utils.IsNil(o.Customer) {
		var ret InternalWebControllersMerchantApiv1InvoiceCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomerOk() (*InternalWebControllersMerchantApiv1InvoiceCustomer, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasCustomer() bool {
	if o != nil && !utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given InternalWebControllersMerchantApiv1InvoiceCustomer and assigns it to the Customer field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetCustomer(v InternalWebControllersMerchantApiv1InvoiceCustomer) {
	o.Customer = &v
}

// GetCustomerNotification returns the CustomerNotification field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomerNotification() []string {
	if o == nil || utils.IsNil(o.CustomerNotification) {
		var ret []string
		return ret
	}
	return o.CustomerNotification
}

// GetCustomerNotificationOk returns a tuple with the CustomerNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomerNotificationOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.CustomerNotification) {
		return nil, false
	}
	return o.CustomerNotification, true
}

// HasCustomerNotification returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasCustomerNotification() bool {
	if o != nil && !utils.IsNil(o.CustomerNotification) {
		return true
	}

	return false
}

// SetCustomerNotification gets a reference to the given []string and assigns it to the CustomerNotification field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetCustomerNotification(v []string) {
	o.CustomerNotification = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetExternalId() string {
	if o == nil || utils.IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetExternalIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasExternalId() bool {
	if o != nil && !utils.IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetInvoiceLinkId returns the InvoiceLinkId field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetInvoiceLinkId() string {
	if o == nil || utils.IsNil(o.InvoiceLinkId) {
		var ret string
		return ret
	}
	return *o.InvoiceLinkId
}

// GetInvoiceLinkIdOk returns a tuple with the InvoiceLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetInvoiceLinkIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InvoiceLinkId) {
		return nil, false
	}
	return o.InvoiceLinkId, true
}

// HasInvoiceLinkId returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasInvoiceLinkId() bool {
	if o != nil && !utils.IsNil(o.InvoiceLinkId) {
		return true
	}

	return false
}

// SetInvoiceLinkId gets a reference to the given string and assigns it to the InvoiceLinkId field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetInvoiceLinkId(v string) {
	o.InvoiceLinkId = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetItems() []InternalWebControllersMerchantApiv1InvoiceItem {
	if o == nil || utils.IsNil(o.Items) {
		var ret []InternalWebControllersMerchantApiv1InvoiceItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetItemsOk() ([]InternalWebControllersMerchantApiv1InvoiceItem, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InternalWebControllersMerchantApiv1InvoiceItem and assigns it to the Items field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetItems(v []InternalWebControllersMerchantApiv1InvoiceItem) {
	o.Items = v
}

// GetMerchantLogo returns the MerchantLogo field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantLogo() string {
	if o == nil || utils.IsNil(o.MerchantLogo) {
		var ret string
		return ret
	}
	return *o.MerchantLogo
}

// GetMerchantLogoOk returns a tuple with the MerchantLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantLogoOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantLogo) {
		return nil, false
	}
	return o.MerchantLogo, true
}

// HasMerchantLogo returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasMerchantLogo() bool {
	if o != nil && !utils.IsNil(o.MerchantLogo) {
		return true
	}

	return false
}

// SetMerchantLogo gets a reference to the given string and assigns it to the MerchantLogo field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetMerchantLogo(v string) {
	o.MerchantLogo = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantName() string {
	if o == nil || utils.IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasMerchantName() bool {
	if o != nil && !utils.IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetNotes() string {
	if o == nil || utils.IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetNotesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasNotes() bool {
	if o != nil && !utils.IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetNotes(v string) {
	o.Notes = &v
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetPaymentMethods() []InternalWebControllersMerchantApiv1InvoicePaymentMethod {
	if o == nil || utils.IsNil(o.PaymentMethods) {
		var ret []InternalWebControllersMerchantApiv1InvoicePaymentMethod
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetPaymentMethodsOk() ([]InternalWebControllersMerchantApiv1InvoicePaymentMethod, bool) {
	if o == nil || utils.IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasPaymentMethods() bool {
	if o != nil && !utils.IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []InternalWebControllersMerchantApiv1InvoicePaymentMethod and assigns it to the PaymentMethods field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetPaymentMethods(v []InternalWebControllersMerchantApiv1InvoicePaymentMethod) {
	o.PaymentMethods = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetRedirectUrl() InternalWebControllersMerchantApiv1InvoiceRedirectURL {
	if o == nil || utils.IsNil(o.RedirectUrl) {
		var ret InternalWebControllersMerchantApiv1InvoiceRedirectURL
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetRedirectUrlOk() (*InternalWebControllersMerchantApiv1InvoiceRedirectURL, bool) {
	if o == nil || utils.IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given InternalWebControllersMerchantApiv1InvoiceRedirectURL and assigns it to the RedirectUrl field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetRedirectUrl(v InternalWebControllersMerchantApiv1InvoiceRedirectURL) {
	o.RedirectUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTimeBegin returns the TimeBegin field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeBegin() string {
	if o == nil || utils.IsNil(o.TimeBegin) {
		var ret string
		return ret
	}
	return *o.TimeBegin
}

// GetTimeBeginOk returns a tuple with the TimeBegin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeBeginOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TimeBegin) {
		return nil, false
	}
	return o.TimeBegin, true
}

// HasTimeBegin returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasTimeBegin() bool {
	if o != nil && !utils.IsNil(o.TimeBegin) {
		return true
	}

	return false
}

// SetTimeBegin gets a reference to the given string and assigns it to the TimeBegin field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetTimeBegin(v string) {
	o.TimeBegin = &v
}

// GetTimeEnd returns the TimeEnd field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeEnd() string {
	if o == nil || utils.IsNil(o.TimeEnd) {
		var ret string
		return ret
	}
	return *o.TimeEnd
}

// GetTimeEndOk returns a tuple with the TimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeEndOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TimeEnd) {
		return nil, false
	}
	return o.TimeEnd, true
}

// HasTimeEnd returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasTimeEnd() bool {
	if o != nil && !utils.IsNil(o.TimeEnd) {
		return true
	}

	return false
}

// SetTimeEnd gets a reference to the given string and assigns it to the TimeEnd field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetTimeEnd(v string) {
	o.TimeEnd = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.AmountCurrency) {
		toSerialize["amount_currency"] = o.AmountCurrency
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !utils.IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !utils.IsNil(o.CustomerNotification) {
		toSerialize["customer_notification"] = o.CustomerNotification
	}
	if !utils.IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !utils.IsNil(o.InvoiceLinkId) {
		toSerialize["invoice_link_id"] = o.InvoiceLinkId
	}
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !utils.IsNil(o.MerchantLogo) {
		toSerialize["merchant_logo"] = o.MerchantLogo
	}
	if !utils.IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if !utils.IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !utils.IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !utils.IsNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.TimeBegin) {
		toSerialize["time_begin"] = o.TimeBegin
	}
	if !utils.IsNil(o.TimeEnd) {
		toSerialize["time_end"] = o.TimeEnd
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse struct {
	value *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) Get() *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) Set(val *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse(val *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse {
	return &NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
