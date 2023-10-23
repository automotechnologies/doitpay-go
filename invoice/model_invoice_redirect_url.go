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

// checks if the InternalWebControllersMerchantApiv1InvoiceRedirectURL type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceRedirectURL{}

// InternalWebControllersMerchantApiv1InvoiceRedirectURL struct for InternalWebControllersMerchantApiv1InvoiceRedirectURL
type InternalWebControllersMerchantApiv1InvoiceRedirectURL struct {
	Cancel  *string `json:"cancel,omitempty"`
	Success *string `json:"success,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceRedirectURL instantiates a new InternalWebControllersMerchantApiv1InvoiceRedirectURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceRedirectURL() *InternalWebControllersMerchantApiv1InvoiceRedirectURL {
	this := InternalWebControllersMerchantApiv1InvoiceRedirectURL{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceRedirectURLWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceRedirectURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceRedirectURLWithDefaults() *InternalWebControllersMerchantApiv1InvoiceRedirectURL {
	this := InternalWebControllersMerchantApiv1InvoiceRedirectURL{}
	return &this
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) GetCancel() string {
	if o == nil || utils.IsNil(o.Cancel) {
		var ret string
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) GetCancelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Cancel) {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) HasCancel() bool {
	if o != nil && !utils.IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given string and assigns it to the Cancel field.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) SetCancel(v string) {
	o.Cancel = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) GetSuccess() string {
	if o == nil || utils.IsNil(o.Success) {
		var ret string
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) GetSuccessOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) HasSuccess() bool {
	if o != nil && !utils.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given string and assigns it to the Success field.
func (o *InternalWebControllersMerchantApiv1InvoiceRedirectURL) SetSuccess(v string) {
	o.Success = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceRedirectURL) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceRedirectURL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cancel) {
		toSerialize["cancel"] = o.Cancel
	}
	if !utils.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL struct {
	value *InternalWebControllersMerchantApiv1InvoiceRedirectURL
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL) Get() *InternalWebControllersMerchantApiv1InvoiceRedirectURL {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL) Set(val *InternalWebControllersMerchantApiv1InvoiceRedirectURL) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceRedirectURL(val *InternalWebControllersMerchantApiv1InvoiceRedirectURL) *NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL {
	return &NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceRedirectURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
