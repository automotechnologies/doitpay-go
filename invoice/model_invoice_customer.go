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

// checks if the InternalWebControllersMerchantApiv1InvoiceCustomer type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1InvoiceCustomer{}

// InternalWebControllersMerchantApiv1InvoiceCustomer struct for InternalWebControllersMerchantApiv1InvoiceCustomer
type InternalWebControllersMerchantApiv1InvoiceCustomer struct {
	Addresses     *string `json:"addresses,omitempty"`
	City          *string `json:"city,omitempty"`
	Country       *string `json:"country,omitempty"`
	CustomerRefId *string `json:"customer_ref_id,omitempty"`
	Email         *string `json:"email,omitempty"`
	Name          *string `json:"name,omitempty"`
	Notes         *string `json:"notes,omitempty"`
	Phone         *string `json:"phone,omitempty"`
	PostalCode    *int32  `json:"postal_code,omitempty"`
}

// NewInternalWebControllersMerchantApiv1InvoiceCustomer instantiates a new InternalWebControllersMerchantApiv1InvoiceCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1InvoiceCustomer() *InternalWebControllersMerchantApiv1InvoiceCustomer {
	this := InternalWebControllersMerchantApiv1InvoiceCustomer{}
	return &this
}

// NewInternalWebControllersMerchantApiv1InvoiceCustomerWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1InvoiceCustomerWithDefaults() *InternalWebControllersMerchantApiv1InvoiceCustomer {
	this := InternalWebControllersMerchantApiv1InvoiceCustomer{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetAddresses() string {
	if o == nil || utils.IsNil(o.Addresses) {
		var ret string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetAddressesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasAddresses() bool {
	if o != nil && !utils.IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given string and assigns it to the Addresses field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetAddresses(v string) {
	o.Addresses = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetCity() string {
	if o == nil || utils.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetCityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasCity() bool {
	if o != nil && !utils.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetCountry(v string) {
	o.Country = &v
}

// GetCustomerRefId returns the CustomerRefId field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetCustomerRefId() string {
	if o == nil || utils.IsNil(o.CustomerRefId) {
		var ret string
		return ret
	}
	return *o.CustomerRefId
}

// GetCustomerRefIdOk returns a tuple with the CustomerRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetCustomerRefIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CustomerRefId) {
		return nil, false
	}
	return o.CustomerRefId, true
}

// HasCustomerRefId returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasCustomerRefId() bool {
	if o != nil && !utils.IsNil(o.CustomerRefId) {
		return true
	}

	return false
}

// SetCustomerRefId gets a reference to the given string and assigns it to the CustomerRefId field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetCustomerRefId(v string) {
	o.CustomerRefId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetNotes() string {
	if o == nil || utils.IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetNotesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasNotes() bool {
	if o != nil && !utils.IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetNotes(v string) {
	o.Notes = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetPhone() string {
	if o == nil || utils.IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetPhoneOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasPhone() bool {
	if o != nil && !utils.IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetPhone(v string) {
	o.Phone = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetPostalCode() int32 {
	if o == nil || utils.IsNil(o.PostalCode) {
		var ret int32
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) GetPostalCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) HasPostalCode() bool {
	if o != nil && !utils.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given int32 and assigns it to the PostalCode field.
func (o *InternalWebControllersMerchantApiv1InvoiceCustomer) SetPostalCode(v int32) {
	o.PostalCode = &v
}

func (o InternalWebControllersMerchantApiv1InvoiceCustomer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1InvoiceCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !utils.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !utils.IsNil(o.CustomerRefId) {
		toSerialize["customer_ref_id"] = o.CustomerRefId
	}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !utils.IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !utils.IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1InvoiceCustomer struct {
	value *InternalWebControllersMerchantApiv1InvoiceCustomer
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceCustomer) Get() *InternalWebControllersMerchantApiv1InvoiceCustomer {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceCustomer) Set(val *InternalWebControllersMerchantApiv1InvoiceCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1InvoiceCustomer(val *InternalWebControllersMerchantApiv1InvoiceCustomer) *NullableInternalWebControllersMerchantApiv1InvoiceCustomer {
	return &NullableInternalWebControllersMerchantApiv1InvoiceCustomer{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1InvoiceCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1InvoiceCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
