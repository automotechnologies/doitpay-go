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

// checks if the InternalWebControllersMerchantApiv1InvoicePaymentMethod type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoicePaymentMethod{}

// InternalWebControllersMerchantApiv1InvoicePaymentMethod struct for InternalWebControllersMerchantApiv1InvoicePaymentMethod
type InternalWebControllersMerchantApiv1InvoicePaymentMethod struct {
	Code   *string `json:"code,omitempty"`
	Status *string `json:"status,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoicePaymentMethod instantiates a new InternalWebControllersMerchantApiv1InvoicePaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoicePaymentMethod() *InternalWebControllersMerchantApiv1InvoicePaymentMethod {
	this := InternalWebControllersMerchantApiv1InvoicePaymentMethod{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoicePaymentMethodWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoicePaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoicePaymentMethodWithDefaults() *InternalWebControllersMerchantApiv1InvoicePaymentMethod {
	this := InternalWebControllersMerchantApiv1InvoicePaymentMethod{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) GetCode() string {
	if o == nil || utils.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) GetCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) SetCode(v string) {
	o.Code = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InternalWebControllersMerchantApiv1InvoicePaymentMethod) SetType(v string) {
	o.Type = &v
}

func (o InternalWebControllersMerchantApiv1InvoicePaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoicePaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod struct {
	value *InternalWebControllersMerchantApiv1InvoicePaymentMethod
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod) Get() *InternalWebControllersMerchantApiv1InvoicePaymentMethod {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod) Set(val *InternalWebControllersMerchantApiv1InvoicePaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoicePaymentMethod(val *InternalWebControllersMerchantApiv1InvoicePaymentMethod) *NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod {
	return &NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoicePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
