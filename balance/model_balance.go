/*
Doitpay API

This is the payment gateway API server for Doitpay.

API version: 1.0
Contact: support@doitpay.co.id
*/

package balance

import (
	"encoding/json"

	"github.com/automotechnologies/doitpay-go/utils"
)

// checks if the InternalWebControllersMerchantApiv1BalanceBalance type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1BalanceBalance{}

// InternalWebControllersMerchantApiv1BalanceBalance struct for InternalWebControllersMerchantApiv1BalanceBalance
type InternalWebControllersMerchantApiv1BalanceBalance struct {
	Gross *float32 `json:"gross,omitempty"`
	Net   *float32 `json:"net,omitempty"`
}

// NewInternalWebControllersMerchantApiv1BalanceBalance instantiates a new InternalWebControllersMerchantApiv1BalanceBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1BalanceBalance() *InternalWebControllersMerchantApiv1BalanceBalance {
	this := InternalWebControllersMerchantApiv1BalanceBalance{}
	return &this
}

// NewInternalWebControllersMerchantApiv1BalanceBalanceWithDefaults instantiates a new InternalWebControllersMerchantApiv1BalanceBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1BalanceBalanceWithDefaults() *InternalWebControllersMerchantApiv1BalanceBalance {
	this := InternalWebControllersMerchantApiv1BalanceBalance{}
	return &this
}

// GetGross returns the Gross field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) GetGross() float32 {
	if o == nil || utils.IsNil(o.Gross) {
		var ret float32
		return ret
	}
	return *o.Gross
}

// GetGrossOk returns a tuple with the Gross field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) GetGrossOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Gross) {
		return nil, false
	}
	return o.Gross, true
}

// HasGross returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) HasGross() bool {
	if o != nil && !utils.IsNil(o.Gross) {
		return true
	}

	return false
}

// SetGross gets a reference to the given float32 and assigns it to the Gross field.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) SetGross(v float32) {
	o.Gross = &v
}

// GetNet returns the Net field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) GetNet() float32 {
	if o == nil || utils.IsNil(o.Net) {
		var ret float32
		return ret
	}
	return *o.Net
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) GetNetOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Net) {
		return nil, false
	}
	return o.Net, true
}

// HasNet returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) HasNet() bool {
	if o != nil && !utils.IsNil(o.Net) {
		return true
	}

	return false
}

// SetNet gets a reference to the given float32 and assigns it to the Net field.
func (o *InternalWebControllersMerchantApiv1BalanceBalance) SetNet(v float32) {
	o.Net = &v
}

func (o InternalWebControllersMerchantApiv1BalanceBalance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1BalanceBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Gross) {
		toSerialize["gross"] = o.Gross
	}
	if !utils.IsNil(o.Net) {
		toSerialize["net"] = o.Net
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1BalanceBalance struct {
	value *InternalWebControllersMerchantApiv1BalanceBalance
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1BalanceBalance) Get() *InternalWebControllersMerchantApiv1BalanceBalance {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceBalance) Set(val *InternalWebControllersMerchantApiv1BalanceBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1BalanceBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1BalanceBalance(val *InternalWebControllersMerchantApiv1BalanceBalance) *NullableInternalWebControllersMerchantApiv1BalanceBalance {
	return &NullableInternalWebControllersMerchantApiv1BalanceBalance{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1BalanceBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
