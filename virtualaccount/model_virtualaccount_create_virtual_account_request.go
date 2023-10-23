/*
Doitpay API

This is the payment gateway API server for Doitpay.

API version: 1.0
Contact: support@doitpay.co.id
*/

package virtualaccount

import (
	"encoding/json"

	"github.com/automotechnologies/doitpay-go/utils"
)

// checks if the InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest{}

// InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest struct for InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest
type InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest struct {
	Amount               *float32                                                                 `json:"amount,omitempty"`
	AmountMax            *float32                                                                 `json:"amount_max,omitempty"`
	AmountMin            *float32                                                                 `json:"amount_min,omitempty"`
	BusinessId           *int32                                                                   `json:"business_id,omitempty"`
	Currency             *string                                                                  `json:"currency,omitempty"`
	Customer             *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer `json:"customer,omitempty"`
	ExpirationDate       *string                                                                  `json:"expiration_date,omitempty"`
	IsClosed             *bool                                                                    `json:"is_closed,omitempty"`
	IsReusable           *bool                                                                    `json:"is_reusable,omitempty"`
	PaymentMethodCode    *string                                                                  `json:"payment_method_code,omitempty"`
	ReferenceId          *string                                                                  `json:"reference_id,omitempty"`
	ReferenceInternalId  *string                                                                  `json:"reference_internal_id,omitempty"`
	VirtualAccountSuffix *string                                                                  `json:"virtual_account_suffix,omitempty"`
}

// NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest instantiates a new InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest() *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest {
	this := InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest{}
	return &this
}

// NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequestWithDefaults instantiates a new InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequestWithDefaults() *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest {
	this := InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmount() float32 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetAmount(v float32) {
	o.Amount = &v
}

// GetAmountMax returns the AmountMax field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMax() float32 {
	if o == nil || utils.IsNil(o.AmountMax) {
		var ret float32
		return ret
	}
	return *o.AmountMax
}

// GetAmountMaxOk returns a tuple with the AmountMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMaxOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.AmountMax) {
		return nil, false
	}
	return o.AmountMax, true
}

// HasAmountMax returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasAmountMax() bool {
	if o != nil && !utils.IsNil(o.AmountMax) {
		return true
	}

	return false
}

// SetAmountMax gets a reference to the given float32 and assigns it to the AmountMax field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetAmountMax(v float32) {
	o.AmountMax = &v
}

// GetAmountMin returns the AmountMin field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMin() float32 {
	if o == nil || utils.IsNil(o.AmountMin) {
		var ret float32
		return ret
	}
	return *o.AmountMin
}

// GetAmountMinOk returns a tuple with the AmountMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMinOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.AmountMin) {
		return nil, false
	}
	return o.AmountMin, true
}

// HasAmountMin returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasAmountMin() bool {
	if o != nil && !utils.IsNil(o.AmountMin) {
		return true
	}

	return false
}

// SetAmountMin gets a reference to the given float32 and assigns it to the AmountMin field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetAmountMin(v float32) {
	o.AmountMin = &v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetBusinessId() int32 {
	if o == nil || utils.IsNil(o.BusinessId) {
		var ret int32
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetBusinessIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasBusinessId() bool {
	if o != nil && !utils.IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given int32 and assigns it to the BusinessId field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetBusinessId(v int32) {
	o.BusinessId = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCustomer() InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer {
	if o == nil || utils.IsNil(o.Customer) {
		var ret InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCustomerOk() (*InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasCustomer() bool {
	if o != nil && !utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer and assigns it to the Customer field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetCustomer(v InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) {
	o.Customer = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetExpirationDate() string {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetExpirationDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasExpirationDate() bool {
	if o != nil && !utils.IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetIsClosed returns the IsClosed field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsClosed() bool {
	if o == nil || utils.IsNil(o.IsClosed) {
		var ret bool
		return ret
	}
	return *o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsClosedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsClosed) {
		return nil, false
	}
	return o.IsClosed, true
}

// HasIsClosed returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasIsClosed() bool {
	if o != nil && !utils.IsNil(o.IsClosed) {
		return true
	}

	return false
}

// SetIsClosed gets a reference to the given bool and assigns it to the IsClosed field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetIsClosed(v bool) {
	o.IsClosed = &v
}

// GetIsReusable returns the IsReusable field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsReusable() bool {
	if o == nil || utils.IsNil(o.IsReusable) {
		var ret bool
		return ret
	}
	return *o.IsReusable
}

// GetIsReusableOk returns a tuple with the IsReusable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsReusableOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsReusable) {
		return nil, false
	}
	return o.IsReusable, true
}

// HasIsReusable returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasIsReusable() bool {
	if o != nil && !utils.IsNil(o.IsReusable) {
		return true
	}

	return false
}

// SetIsReusable gets a reference to the given bool and assigns it to the IsReusable field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetIsReusable(v bool) {
	o.IsReusable = &v
}

// GetPaymentMethodCode returns the PaymentMethodCode field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetPaymentMethodCode() string {
	if o == nil || utils.IsNil(o.PaymentMethodCode) {
		var ret string
		return ret
	}
	return *o.PaymentMethodCode
}

// GetPaymentMethodCodeOk returns a tuple with the PaymentMethodCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetPaymentMethodCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentMethodCode) {
		return nil, false
	}
	return o.PaymentMethodCode, true
}

// HasPaymentMethodCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasPaymentMethodCode() bool {
	if o != nil && !utils.IsNil(o.PaymentMethodCode) {
		return true
	}

	return false
}

// SetPaymentMethodCode gets a reference to the given string and assigns it to the PaymentMethodCode field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetPaymentMethodCode(v string) {
	o.PaymentMethodCode = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetReferenceInternalId returns the ReferenceInternalId field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceInternalId() string {
	if o == nil || utils.IsNil(o.ReferenceInternalId) {
		var ret string
		return ret
	}
	return *o.ReferenceInternalId
}

// GetReferenceInternalIdOk returns a tuple with the ReferenceInternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceInternalIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceInternalId) {
		return nil, false
	}
	return o.ReferenceInternalId, true
}

// HasReferenceInternalId returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasReferenceInternalId() bool {
	if o != nil && !utils.IsNil(o.ReferenceInternalId) {
		return true
	}

	return false
}

// SetReferenceInternalId gets a reference to the given string and assigns it to the ReferenceInternalId field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetReferenceInternalId(v string) {
	o.ReferenceInternalId = &v
}

// GetVirtualAccountSuffix returns the VirtualAccountSuffix field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetVirtualAccountSuffix() string {
	if o == nil || utils.IsNil(o.VirtualAccountSuffix) {
		var ret string
		return ret
	}
	return *o.VirtualAccountSuffix
}

// GetVirtualAccountSuffixOk returns a tuple with the VirtualAccountSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetVirtualAccountSuffixOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VirtualAccountSuffix) {
		return nil, false
	}
	return o.VirtualAccountSuffix, true
}

// HasVirtualAccountSuffix returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasVirtualAccountSuffix() bool {
	if o != nil && !utils.IsNil(o.VirtualAccountSuffix) {
		return true
	}

	return false
}

// SetVirtualAccountSuffix gets a reference to the given string and assigns it to the VirtualAccountSuffix field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetVirtualAccountSuffix(v string) {
	o.VirtualAccountSuffix = &v
}

func (o InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.AmountMax) {
		toSerialize["amount_max"] = o.AmountMax
	}
	if !utils.IsNil(o.AmountMin) {
		toSerialize["amount_min"] = o.AmountMin
	}
	if !utils.IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !utils.IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !utils.IsNil(o.ExpirationDate) {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if !utils.IsNil(o.IsClosed) {
		toSerialize["is_closed"] = o.IsClosed
	}
	if !utils.IsNil(o.IsReusable) {
		toSerialize["is_reusable"] = o.IsReusable
	}
	if !utils.IsNil(o.PaymentMethodCode) {
		toSerialize["payment_method_code"] = o.PaymentMethodCode
	}
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !utils.IsNil(o.ReferenceInternalId) {
		toSerialize["reference_internal_id"] = o.ReferenceInternalId
	}
	if !utils.IsNil(o.VirtualAccountSuffix) {
		toSerialize["virtual_account_suffix"] = o.VirtualAccountSuffix
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest struct {
	value *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) Get() *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) Set(val *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest(val *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) *NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest {
	return &NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
