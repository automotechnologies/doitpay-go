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

// checks if the InternalWebControllersMerchantApiv1InvoiceItem type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceItem{}

// InternalWebControllersMerchantApiv1InvoiceItem struct for InternalWebControllersMerchantApiv1InvoiceItem
type InternalWebControllersMerchantApiv1InvoiceItem struct {
	SKU      *string  `json:"SKU,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Notes    *string  `json:"notes,omitempty"`
	Price    *float32 `json:"price,omitempty"`
	Quantity *int32   `json:"quantity,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceItem instantiates a new InternalWebControllersMerchantApiv1InvoiceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceItem() *InternalWebControllersMerchantApiv1InvoiceItem {
	this := InternalWebControllersMerchantApiv1InvoiceItem{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceItemWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceItemWithDefaults() *InternalWebControllersMerchantApiv1InvoiceItem {
	this := InternalWebControllersMerchantApiv1InvoiceItem{}
	return &this
}

// GetSKU returns the SKU field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetSKU() string {
	if o == nil || utils.IsNil(o.SKU) {
		var ret string
		return ret
	}
	return *o.SKU
}

// GetSKUOk returns a tuple with the SKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetSKUOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SKU) {
		return nil, false
	}
	return o.SKU, true
}

// HasSKU returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) HasSKU() bool {
	if o != nil && !utils.IsNil(o.SKU) {
		return true
	}

	return false
}

// SetSKU gets a reference to the given string and assigns it to the SKU field.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) SetSKU(v string) {
	o.SKU = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetNotes() string {
	if o == nil || utils.IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetNotesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) HasNotes() bool {
	if o != nil && !utils.IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) SetNotes(v string) {
	o.Notes = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetPrice() float32 {
	if o == nil || utils.IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetPriceOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) HasPrice() bool {
	if o != nil && !utils.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetQuantity() int32 {
	if o == nil || utils.IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) GetQuantityOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) HasQuantity() bool {
	if o != nil && !utils.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *InternalWebControllersMerchantApiv1InvoiceItem) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SKU) {
		toSerialize["SKU"] = o.SKU
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !utils.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !utils.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1InvoiceItem struct {
	value *InternalWebControllersMerchantApiv1InvoiceItem
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceItem) Get() *InternalWebControllersMerchantApiv1InvoiceItem {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceItem) Set(val *InternalWebControllersMerchantApiv1InvoiceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceItem(val *InternalWebControllersMerchantApiv1InvoiceItem) *NullableInternalWebControllersMerchantApiv1InvoiceItem {
	return &NullableInternalWebControllersMerchantApiv1InvoiceItem{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
