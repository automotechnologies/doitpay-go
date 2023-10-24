/*
Doitpay API

This is the payment gateway API server for Doitpay.

API version: 1.0
Contact: support@doitpay.co.id
*/

package simulate

import (
	"encoding/json"

	"github.com/automotechnologies/doitpay-go/utils"
)

// checks if the InternalWebControllersMerchantApiv1SimulateSimulateRequest type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1SimulateSimulateRequest{}

// InternalWebControllersMerchantApiv1SimulateSimulateRequest struct for InternalWebControllersMerchantApiv1SimulateSimulateRequest
type InternalWebControllersMerchantApiv1SimulateSimulateRequest struct {
	AccountNumber     *string `json:"account_number,omitempty"`
	PaymentIdentifier *string `json:"payment_identifier,omitempty"`
}

// NewInternalWebControllersMerchantApiv1SimulateSimulateRequest instantiates a new InternalWebControllersMerchantApiv1SimulateSimulateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1SimulateSimulateRequest() *InternalWebControllersMerchantApiv1SimulateSimulateRequest {
	this := InternalWebControllersMerchantApiv1SimulateSimulateRequest{}
	return &this
}

// NewInternalWebControllersMerchantApiv1SimulateSimulateRequestWithDefaults instantiates a new InternalWebControllersMerchantApiv1SimulateSimulateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1SimulateSimulateRequestWithDefaults() *InternalWebControllersMerchantApiv1SimulateSimulateRequest {
	this := InternalWebControllersMerchantApiv1SimulateSimulateRequest{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) GetAccountNumber() string {
	if o == nil || utils.IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) HasAccountNumber() bool {
	if o != nil && !utils.IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetPaymentIdentifier returns the PaymentIdentifier field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) GetPaymentIdentifier() string {
	if o == nil || utils.IsNil(o.PaymentIdentifier) {
		var ret string
		return ret
	}
	return *o.PaymentIdentifier
}

// GetPaymentIdentifierOk returns a tuple with the PaymentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) GetPaymentIdentifierOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentIdentifier) {
		return nil, false
	}
	return o.PaymentIdentifier, true
}

// HasPaymentIdentifier returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) HasPaymentIdentifier() bool {
	if o != nil && !utils.IsNil(o.PaymentIdentifier) {
		return true
	}

	return false
}

// SetPaymentIdentifier gets a reference to the given string and assigns it to the PaymentIdentifier field.
func (o *InternalWebControllersMerchantApiv1SimulateSimulateRequest) SetPaymentIdentifier(v string) {
	o.PaymentIdentifier = &v
}

func (o InternalWebControllersMerchantApiv1SimulateSimulateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1SimulateSimulateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !utils.IsNil(o.PaymentIdentifier) {
		toSerialize["payment_identifier"] = o.PaymentIdentifier
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest struct {
	value *InternalWebControllersMerchantApiv1SimulateSimulateRequest
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest) Get() *InternalWebControllersMerchantApiv1SimulateSimulateRequest {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest) Set(val *InternalWebControllersMerchantApiv1SimulateSimulateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1SimulateSimulateRequest(val *InternalWebControllersMerchantApiv1SimulateSimulateRequest) *NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest {
	return &NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1SimulateSimulateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
