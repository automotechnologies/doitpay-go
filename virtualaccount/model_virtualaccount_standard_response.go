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

// checks if the InternalWebControllersMerchantApiv1VirtualaccountStandardResponse type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1VirtualaccountStandardResponse{}

// InternalWebControllersMerchantApiv1VirtualaccountStandardResponse struct for InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
type InternalWebControllersMerchantApiv1VirtualaccountStandardResponse struct {
	Code    *int32                                                                   `json:"code,omitempty"`
	Data    *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse `json:"data,omitempty"`
	Message *string                                                                  `json:"message,omitempty"`
	Meta    *InternalWebControllersMerchantApiv1VirtualaccountStandardMeta           `json:"meta,omitempty"`
	Status  *string                                                                  `json:"status,omitempty"`
}

// NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponse instantiates a new InternalWebControllersMerchantApiv1VirtualaccountStandardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponse() *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse {
	this := InternalWebControllersMerchantApiv1VirtualaccountStandardResponse{}
	return &this
}

// NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithDefaults instantiates a new InternalWebControllersMerchantApiv1VirtualaccountStandardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithDefaults() *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse {
	this := InternalWebControllersMerchantApiv1VirtualaccountStandardResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetCode() int32 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetData() InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse {
	if o == nil || utils.IsNil(o.Data) {
		var ret InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetDataOk() (*InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse and assigns it to the Data field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) SetData(v InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetMeta() InternalWebControllersMerchantApiv1VirtualaccountStandardMeta {
	if o == nil || utils.IsNil(o.Meta) {
		var ret InternalWebControllersMerchantApiv1VirtualaccountStandardMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetMetaOk() (*InternalWebControllersMerchantApiv1VirtualaccountStandardMeta, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) HasMeta() bool {
	if o != nil && !utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InternalWebControllersMerchantApiv1VirtualaccountStandardMeta and assigns it to the Meta field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) SetMeta(v InternalWebControllersMerchantApiv1VirtualaccountStandardMeta) {
	o.Meta = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) SetStatus(v string) {
	o.Status = &v
}

func (o InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) ToMap() (map[string]interface{}, error) {
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

type NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse struct {
	value *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse) Get() *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse) Set(val *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse(val *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse) *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse {
	return &NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
