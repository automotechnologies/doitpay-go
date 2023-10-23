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

// checks if the InternalWebControllersMerchantApiv1InvoiceStandardMeta type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceStandardMeta{}

// InternalWebControllersMerchantApiv1InvoiceStandardMeta struct for InternalWebControllersMerchantApiv1InvoiceStandardMeta
type InternalWebControllersMerchantApiv1InvoiceStandardMeta struct {
	ProcessingTime *string `json:"processing_time,omitempty"`
	VersionApi     *string `json:"version_api,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceStandardMeta instantiates a new InternalWebControllersMerchantApiv1InvoiceStandardMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceStandardMeta() *InternalWebControllersMerchantApiv1InvoiceStandardMeta {
	this := InternalWebControllersMerchantApiv1InvoiceStandardMeta{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceStandardMetaWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceStandardMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceStandardMetaWithDefaults() *InternalWebControllersMerchantApiv1InvoiceStandardMeta {
	this := InternalWebControllersMerchantApiv1InvoiceStandardMeta{}
	return &this
}

// GetProcessingTime returns the ProcessingTime field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) GetProcessingTime() string {
	if o == nil || utils.IsNil(o.ProcessingTime) {
		var ret string
		return ret
	}
	return *o.ProcessingTime
}

// GetProcessingTimeOk returns a tuple with the ProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) GetProcessingTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProcessingTime) {
		return nil, false
	}
	return o.ProcessingTime, true
}

// HasProcessingTime returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) HasProcessingTime() bool {
	if o != nil && !utils.IsNil(o.ProcessingTime) {
		return true
	}

	return false
}

// SetProcessingTime gets a reference to the given string and assigns it to the ProcessingTime field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) SetProcessingTime(v string) {
	o.ProcessingTime = &v
}

// GetVersionApi returns the VersionApi field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) GetVersionApi() string {
	if o == nil || utils.IsNil(o.VersionApi) {
		var ret string
		return ret
	}
	return *o.VersionApi
}

// GetVersionApiOk returns a tuple with the VersionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) GetVersionApiOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VersionApi) {
		return nil, false
	}
	return o.VersionApi, true
}

// HasVersionApi returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) HasVersionApi() bool {
	if o != nil && !utils.IsNil(o.VersionApi) {
		return true
	}

	return false
}

// SetVersionApi gets a reference to the given string and assigns it to the VersionApi field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardMeta) SetVersionApi(v string) {
	o.VersionApi = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceStandardMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceStandardMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ProcessingTime) {
		toSerialize["processing_time"] = o.ProcessingTime
	}
	if !utils.IsNil(o.VersionApi) {
		toSerialize["version_api"] = o.VersionApi
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta struct {
	value *InternalWebControllersMerchantApiv1InvoiceStandardMeta
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta) Get() *InternalWebControllersMerchantApiv1InvoiceStandardMeta {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta) Set(val *InternalWebControllersMerchantApiv1InvoiceStandardMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceStandardMeta(val *InternalWebControllersMerchantApiv1InvoiceStandardMeta) *NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta {
	return &NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
