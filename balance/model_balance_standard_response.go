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

// checks if the InternalWebControllersMerchantApiv1BalanceStandardResponse type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1BalanceStandardResponse{}

// InternalWebControllersMerchantApiv1BalanceStandardResponse struct for InternalWebControllersMerchantApiv1BalanceStandardResponse
type InternalWebControllersMerchantApiv1BalanceStandardResponse struct {
	Code    *int32                                                  `json:"code,omitempty"`
	Data    *InternalWebControllersMerchantApiv1BalanceBalance      `json:"data,omitempty"`
	Message *string                                                 `json:"message,omitempty"`
	Meta    *InternalWebControllersMerchantApiv1BalanceStandardMeta `json:"meta,omitempty"`
	Status  *string                                                 `json:"status,omitempty"`
}

// NewInternalWebControllersMerchantApiv1BalanceStandardResponse instantiates a new InternalWebControllersMerchantApiv1BalanceStandardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1BalanceStandardResponse() *InternalWebControllersMerchantApiv1BalanceStandardResponse {
	this := InternalWebControllersMerchantApiv1BalanceStandardResponse{}
	return &this
}

// NewInternalWebControllersMerchantApiv1BalanceStandardResponseWithDefaults instantiates a new InternalWebControllersMerchantApiv1BalanceStandardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1BalanceStandardResponseWithDefaults() *InternalWebControllersMerchantApiv1BalanceStandardResponse {
	this := InternalWebControllersMerchantApiv1BalanceStandardResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetCode() int32 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetData() InternalWebControllersMerchantApiv1BalanceBalance {
	if o == nil || utils.IsNil(o.Data) {
		var ret InternalWebControllersMerchantApiv1BalanceBalance
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetDataOk() (*InternalWebControllersMerchantApiv1BalanceBalance, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given InternalWebControllersMerchantApiv1BalanceBalance and assigns it to the Data field.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) SetData(v InternalWebControllersMerchantApiv1BalanceBalance) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetMeta() InternalWebControllersMerchantApiv1BalanceStandardMeta {
	if o == nil || utils.IsNil(o.Meta) {
		var ret InternalWebControllersMerchantApiv1BalanceStandardMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetMetaOk() (*InternalWebControllersMerchantApiv1BalanceStandardMeta, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) HasMeta() bool {
	if o != nil && !utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InternalWebControllersMerchantApiv1BalanceStandardMeta and assigns it to the Meta field.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) SetMeta(v InternalWebControllersMerchantApiv1BalanceStandardMeta) {
	o.Meta = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalWebControllersMerchantApiv1BalanceStandardResponse) SetStatus(v string) {
	o.Status = &v
}

func (o InternalWebControllersMerchantApiv1BalanceStandardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1BalanceStandardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !utils.IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1BalanceStandardResponse struct {
	value *InternalWebControllersMerchantApiv1BalanceStandardResponse
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1BalanceStandardResponse) Get() *InternalWebControllersMerchantApiv1BalanceStandardResponse {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceStandardResponse) Set(val *InternalWebControllersMerchantApiv1BalanceStandardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1BalanceStandardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceStandardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1BalanceStandardResponse(val *InternalWebControllersMerchantApiv1BalanceStandardResponse) *NullableInternalWebControllersMerchantApiv1BalanceStandardResponse {
	return &NullableInternalWebControllersMerchantApiv1BalanceStandardResponse{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1BalanceStandardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1BalanceStandardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
