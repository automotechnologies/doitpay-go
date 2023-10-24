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

// checks if the InternalWebControllersMerchantApiv1InvoiceStandardResponse type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceStandardResponse{}

// InternalWebControllersMerchantApiv1InvoiceStandardResponse struct for InternalWebControllersMerchantApiv1InvoiceStandardResponse
type InternalWebControllersMerchantApiv1InvoiceStandardResponse struct {
	Code    *int32                                                           `json:"code,omitempty"`
	Data    *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse `json:"data,omitempty"`
	Message *string                                                          `json:"message,omitempty"`
	Meta    *InternalWebControllersMerchantApiv1InvoiceStandardMeta          `json:"meta,omitempty"`
	Status  *string                                                          `json:"status,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceStandardResponse instantiates a new InternalWebControllersMerchantApiv1InvoiceStandardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceStandardResponse() *InternalWebControllersMerchantApiv1InvoiceStandardResponse {
	this := InternalWebControllersMerchantApiv1InvoiceStandardResponse{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceStandardResponseWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceStandardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceStandardResponseWithDefaults() *InternalWebControllersMerchantApiv1InvoiceStandardResponse {
	this := InternalWebControllersMerchantApiv1InvoiceStandardResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetCode() int32 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetData() InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse {
	if o == nil || utils.IsNil(o.Data) {
		var ret InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetDataOk() (*InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse and assigns it to the Data field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) SetData(v InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetMeta() InternalWebControllersMerchantApiv1InvoiceStandardMeta {
	if o == nil || utils.IsNil(o.Meta) {
		var ret InternalWebControllersMerchantApiv1InvoiceStandardMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetMetaOk() (*InternalWebControllersMerchantApiv1InvoiceStandardMeta, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) HasMeta() bool {
	if o != nil && !utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InternalWebControllersMerchantApiv1InvoiceStandardMeta and assigns it to the Meta field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) SetMeta(v InternalWebControllersMerchantApiv1InvoiceStandardMeta) {
	o.Meta = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardResponse) SetStatus(v string) {
	o.Status = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceStandardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceStandardResponse) ToMap() (map[string]interface{}, error) {
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

type NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse struct {
	value *InternalWebControllersMerchantApiv1InvoiceStandardResponse
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse) Get() *InternalWebControllersMerchantApiv1InvoiceStandardResponse {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse) Set(val *InternalWebControllersMerchantApiv1InvoiceStandardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceStandardResponse(val *InternalWebControllersMerchantApiv1InvoiceStandardResponse) *NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse {
	return &NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
