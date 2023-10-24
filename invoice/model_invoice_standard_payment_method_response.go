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

// checks if the InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse{}

// InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse struct for InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse
type InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse struct {
	Code    *int32                                                    `json:"code,omitempty"`
	Data    []InternalWebControllersMerchantApiv1InvoicePaymentMethod `json:"data,omitempty"`
	Message *string                                                   `json:"message,omitempty"`
	Meta    *InternalWebControllersMerchantApiv1InvoiceStandardMeta   `json:"meta,omitempty"`
	Status  *string                                                   `json:"status,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse instantiates a new InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse() *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse {
	this := InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponseWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponseWithDefaults() *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse {
	this := InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetCode() int32 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetData() []InternalWebControllersMerchantApiv1InvoicePaymentMethod {
	if o == nil || utils.IsNil(o.Data) {
		var ret []InternalWebControllersMerchantApiv1InvoicePaymentMethod
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetDataOk() ([]InternalWebControllersMerchantApiv1InvoicePaymentMethod, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []InternalWebControllersMerchantApiv1InvoicePaymentMethod and assigns it to the Data field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) SetData(v []InternalWebControllersMerchantApiv1InvoicePaymentMethod) {
	o.Data = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetMeta() InternalWebControllersMerchantApiv1InvoiceStandardMeta {
	if o == nil || utils.IsNil(o.Meta) {
		var ret InternalWebControllersMerchantApiv1InvoiceStandardMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetMetaOk() (*InternalWebControllersMerchantApiv1InvoiceStandardMeta, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) HasMeta() bool {
	if o != nil && !utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InternalWebControllersMerchantApiv1InvoiceStandardMeta and assigns it to the Meta field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) SetMeta(v InternalWebControllersMerchantApiv1InvoiceStandardMeta) {
	o.Meta = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) SetStatus(v string) {
	o.Status = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) ToMap() (map[string]interface{}, error) {
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

type NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse struct {
	value *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) Get() *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) Set(val *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse(val *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) *NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse {
	return &NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
