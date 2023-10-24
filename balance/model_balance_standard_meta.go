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

// checks if the InternalWebControllersMerchantApiv1BalanceStandardMeta type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1BalanceStandardMeta{}

// InternalWebControllersMerchantApiv1BalanceStandardMeta struct for InternalWebControllersMerchantApiv1BalanceStandardMeta
type InternalWebControllersMerchantApiv1BalanceStandardMeta struct {
	ProcessingTime *string `json:"processing_time,omitempty"`
	VersionApi     *string `json:"version_api,omitempty"`
}

// NewInternalWebControllersMerchantApiv1BalanceStandardMeta instantiates a new InternalWebControllersMerchantApiv1BalanceStandardMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1BalanceStandardMeta() *InternalWebControllersMerchantApiv1BalanceStandardMeta {
	this := InternalWebControllersMerchantApiv1BalanceStandardMeta{}
	return &this
}

// NewInternalWebControllersMerchantApiv1BalanceStandardMetaWithDefaults instantiates a new InternalWebControllersMerchantApiv1BalanceStandardMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1BalanceStandardMetaWithDefaults() *InternalWebControllersMerchantApiv1BalanceStandardMeta {
	this := InternalWebControllersMerchantApiv1BalanceStandardMeta{}
	return &this
}

// GetProcessingTime returns the ProcessingTime field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) GetProcessingTime() string {
	if o == nil || utils.IsNil(o.ProcessingTime) {
		var ret string
		return ret
	}
	return *o.ProcessingTime
}

// GetProcessingTimeOk returns a tuple with the ProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) GetProcessingTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProcessingTime) {
		return nil, false
	}
	return o.ProcessingTime, true
}

// HasProcessingTime returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) HasProcessingTime() bool {
	if o != nil && !utils.IsNil(o.ProcessingTime) {
		return true
	}

	return false
}

// SetProcessingTime gets a reference to the given string and assigns it to the ProcessingTime field.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) SetProcessingTime(v string) {
	o.ProcessingTime = &v
}

// GetVersionApi returns the VersionApi field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) GetVersionApi() string {
	if o == nil || utils.IsNil(o.VersionApi) {
		var ret string
		return ret
	}
	return *o.VersionApi
}

// GetVersionApiOk returns a tuple with the VersionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) GetVersionApiOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VersionApi) {
		return nil, false
	}
	return o.VersionApi, true
}

// HasVersionApi returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) HasVersionApi() bool {
	if o != nil && !utils.IsNil(o.VersionApi) {
		return true
	}

	return false
}

// SetVersionApi gets a reference to the given string and assigns it to the VersionApi field.
func (o *InternalWebControllersMerchantApiv1BalanceStandardMeta) SetVersionApi(v string) {
	o.VersionApi = &v
}

func (o InternalWebControllersMerchantApiv1BalanceStandardMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1BalanceStandardMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ProcessingTime) {
		toSerialize["processing_time"] = o.ProcessingTime
	}
	if !utils.IsNil(o.VersionApi) {
		toSerialize["version_api"] = o.VersionApi
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1BalanceStandardMeta struct {
	value *InternalWebControllersMerchantApiv1BalanceStandardMeta
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1BalanceStandardMeta) Get() *InternalWebControllersMerchantApiv1BalanceStandardMeta {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceStandardMeta) Set(val *InternalWebControllersMerchantApiv1BalanceStandardMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1BalanceStandardMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceStandardMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1BalanceStandardMeta(val *InternalWebControllersMerchantApiv1BalanceStandardMeta) *NullableInternalWebControllersMerchantApiv1BalanceStandardMeta {
	return &NullableInternalWebControllersMerchantApiv1BalanceStandardMeta{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1BalanceStandardMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceStandardMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
