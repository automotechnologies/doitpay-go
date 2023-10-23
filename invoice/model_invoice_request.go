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

// checks if the InternalWebControllersMerchantApiv1InvoiceInvoiceRequest type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceInvoiceRequest{}

// InternalWebControllersMerchantApiv1InvoiceInvoiceRequest Invoice Request
type InternalWebControllersMerchantApiv1InvoiceInvoiceRequest struct {
	Amount               *float32                                                  `json:"amount,omitempty"`
	AmountCurrency       *string                                                   `json:"amount_currency,omitempty"`
	Country              *string                                                   `json:"country,omitempty"`
	Customer             *InternalWebControllersMerchantApiv1InvoiceCustomer       `json:"customer,omitempty"`
	CustomerNotification []string                                                  `json:"customer_notification,omitempty"`
	ExternalId           *string                                                   `json:"external_id,omitempty"`
	Items                []InternalWebControllersMerchantApiv1InvoiceItem          `json:"items,omitempty"`
	Notes                *string                                                   `json:"notes,omitempty"`
	PaymentMethods       []InternalWebControllersMerchantApiv1InvoicePaymentMethod `json:"payment_methods,omitempty"`
	RedirectUrl          *InternalWebControllersMerchantApiv1InvoiceRedirectURL    `json:"redirect_url,omitempty"`
	TimeBegin            *string                                                   `json:"time_begin,omitempty"`
	TimeEnd              *string                                                   `json:"time_end,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequest instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequest() *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest {
	this := InternalWebControllersMerchantApiv1InvoiceInvoiceRequest{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequestWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequestWithDefaults() *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest {
	this := InternalWebControllersMerchantApiv1InvoiceInvoiceRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmount() float32 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmountOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetAmount(v float32) {
	o.Amount = &v
}

// GetAmountCurrency returns the AmountCurrency field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmountCurrency() string {
	if o == nil || utils.IsNil(o.AmountCurrency) {
		var ret string
		return ret
	}
	return *o.AmountCurrency
}

// GetAmountCurrencyOk returns a tuple with the AmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmountCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AmountCurrency) {
		return nil, false
	}
	return o.AmountCurrency, true
}

// HasAmountCurrency returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasAmountCurrency() bool {
	if o != nil && !utils.IsNil(o.AmountCurrency) {
		return true
	}

	return false
}

// SetAmountCurrency gets a reference to the given string and assigns it to the AmountCurrency field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetAmountCurrency(v string) {
	o.AmountCurrency = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetCountry(v string) {
	o.Country = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomer() InternalWebControllersMerchantApiv1InvoiceCustomer {
	if o == nil || utils.IsNil(o.Customer) {
		var ret InternalWebControllersMerchantApiv1InvoiceCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomerOk() (*InternalWebControllersMerchantApiv1InvoiceCustomer, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasCustomer() bool {
	if o != nil && !utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given InternalWebControllersMerchantApiv1InvoiceCustomer and assigns it to the Customer field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetCustomer(v InternalWebControllersMerchantApiv1InvoiceCustomer) {
	o.Customer = &v
}

// GetCustomerNotification returns the CustomerNotification field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomerNotification() []string {
	if o == nil || utils.IsNil(o.CustomerNotification) {
		var ret []string
		return ret
	}
	return o.CustomerNotification
}

// GetCustomerNotificationOk returns a tuple with the CustomerNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomerNotificationOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.CustomerNotification) {
		return nil, false
	}
	return o.CustomerNotification, true
}

// HasCustomerNotification returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasCustomerNotification() bool {
	if o != nil && !utils.IsNil(o.CustomerNotification) {
		return true
	}

	return false
}

// SetCustomerNotification gets a reference to the given []string and assigns it to the CustomerNotification field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetCustomerNotification(v []string) {
	o.CustomerNotification = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetExternalId() string {
	if o == nil || utils.IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasExternalId() bool {
	if o != nil && !utils.IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetItems() []InternalWebControllersMerchantApiv1InvoiceItem {
	if o == nil || utils.IsNil(o.Items) {
		var ret []InternalWebControllersMerchantApiv1InvoiceItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetItemsOk() ([]InternalWebControllersMerchantApiv1InvoiceItem, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InternalWebControllersMerchantApiv1InvoiceItem and assigns it to the Items field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetItems(v []InternalWebControllersMerchantApiv1InvoiceItem) {
	o.Items = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetNotes() string {
	if o == nil || utils.IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetNotesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasNotes() bool {
	if o != nil && !utils.IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetPaymentMethods() []InternalWebControllersMerchantApiv1InvoicePaymentMethod {
	if o == nil || utils.IsNil(o.PaymentMethods) {
		var ret []InternalWebControllersMerchantApiv1InvoicePaymentMethod
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetPaymentMethodsOk() ([]InternalWebControllersMerchantApiv1InvoicePaymentMethod, bool) {
	if o == nil || utils.IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasPaymentMethods() bool {
	if o != nil && !utils.IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []InternalWebControllersMerchantApiv1InvoicePaymentMethod and assigns it to the PaymentMethods field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetPaymentMethods(v []InternalWebControllersMerchantApiv1InvoicePaymentMethod) {
	o.PaymentMethods = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetRedirectUrl() InternalWebControllersMerchantApiv1InvoiceRedirectURL {
	if o == nil || utils.IsNil(o.RedirectUrl) {
		var ret InternalWebControllersMerchantApiv1InvoiceRedirectURL
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetRedirectUrlOk() (*InternalWebControllersMerchantApiv1InvoiceRedirectURL, bool) {
	if o == nil || utils.IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given InternalWebControllersMerchantApiv1InvoiceRedirectURL and assigns it to the RedirectUrl field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetRedirectUrl(v InternalWebControllersMerchantApiv1InvoiceRedirectURL) {
	o.RedirectUrl = &v
}

// GetTimeBegin returns the TimeBegin field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeBegin() string {
	if o == nil || utils.IsNil(o.TimeBegin) {
		var ret string
		return ret
	}
	return *o.TimeBegin
}

// GetTimeBeginOk returns a tuple with the TimeBegin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeBeginOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TimeBegin) {
		return nil, false
	}
	return o.TimeBegin, true
}

// HasTimeBegin returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasTimeBegin() bool {
	if o != nil && !utils.IsNil(o.TimeBegin) {
		return true
	}

	return false
}

// SetTimeBegin gets a reference to the given string and assigns it to the TimeBegin field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetTimeBegin(v string) {
	o.TimeBegin = &v
}

// GetTimeEnd returns the TimeEnd field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeEnd() string {
	if o == nil || utils.IsNil(o.TimeEnd) {
		var ret string
		return ret
	}
	return *o.TimeEnd
}

// GetTimeEndOk returns a tuple with the TimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeEndOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TimeEnd) {
		return nil, false
	}
	return o.TimeEnd, true
}

// HasTimeEnd returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasTimeEnd() bool {
	if o != nil && !utils.IsNil(o.TimeEnd) {
		return true
	}

	return false
}

// SetTimeEnd gets a reference to the given string and assigns it to the TimeEnd field.
func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetTimeEnd(v string) {
	o.TimeEnd = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
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
	if !utils.IsNil(o.TimeBegin) {
		toSerialize["time_begin"] = o.TimeBegin
	}
	if !utils.IsNil(o.TimeEnd) {
		toSerialize["time_end"] = o.TimeEnd
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest struct {
	value *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest) Get() *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest) Set(val *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest(val *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest {
	return &NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
