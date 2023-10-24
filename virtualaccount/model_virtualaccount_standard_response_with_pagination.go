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

// checks if the InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination{}

// InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination struct for InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination
type InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination struct {
	Code    *int32                                                                       `json:"code,omitempty"`
	Data    []InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse    `json:"data,omitempty"`
	Message *string                                                                      `json:"message,omitempty"`
	Meta    *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination `json:"meta,omitempty"`
	Status  *string                                                                      `json:"status,omitempty"`
}

// NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination instantiates a new InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination() *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination {
	this := InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination{}
	return &this
}

// NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPaginationWithDefaults instantiates a new InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPaginationWithDefaults() *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination {
	this := InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetCode() int32 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetData() []InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse {
	if o == nil || utils.IsNil(o.Data) {
		var ret []InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetDataOk() ([]InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse and assigns it to the Data field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) SetData(v []InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) {
	o.Data = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) SetMessage(v string) {
	o.Message = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetMeta() InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination {
	if o == nil || utils.IsNil(o.Meta) {
		var ret InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetMetaOk() (*InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) HasMeta() bool {
	if o != nil && !utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination and assigns it to the Meta field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) SetMeta(v InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) {
	o.Meta = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) SetStatus(v string) {
	o.Status = &v
}

func (o InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) ToMap() (map[string]interface{}, error) {
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

type NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination struct {
	value *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) Get() *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) Set(val *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination(val *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination {
	return &NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
