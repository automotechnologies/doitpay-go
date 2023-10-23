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

// checks if the InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer{}

// InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer struct for InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer
type InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer instantiates a new InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer() *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer {
	this := InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer{}
	return &this
}

// NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomerWithDefaults instantiates a new InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomerWithDefaults() *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer {
	this := InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) GetPhone() string {
	if o == nil || utils.IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) GetPhoneOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) HasPhone() bool {
	if o != nil && !utils.IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) SetPhone(v string) {
	o.Phone = &v
}

func (o InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer struct {
	value *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) Get() *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) Set(val *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer(val *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) *NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer {
	return &NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
