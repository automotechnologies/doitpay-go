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

// checks if the InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination{}

// InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination struct for InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination
type InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination struct {
	Pagination     *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader `json:"pagination,omitempty"`
	ProcessingTime *string                                                                 `json:"processing_time,omitempty"`
	VersionApi     *string                                                                 `json:"version_api,omitempty"`
}

// NewInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination instantiates a new InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination() *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination {
	this := InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination{}
	return &this
}

// NewInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPaginationWithDefaults instantiates a new InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPaginationWithDefaults() *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination {
	this := InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) GetPagination() GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader {
	if o == nil || utils.IsNil(o.Pagination) {
		var ret GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) GetPaginationOk() (*GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader, bool) {
	if o == nil || utils.IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) HasPagination() bool {
	if o != nil && !utils.IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader and assigns it to the Pagination field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) SetPagination(v GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) {
	o.Pagination = &v
}

// GetProcessingTime returns the ProcessingTime field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) GetProcessingTime() string {
	if o == nil || utils.IsNil(o.ProcessingTime) {
		var ret string
		return ret
	}
	return *o.ProcessingTime
}

// GetProcessingTimeOk returns a tuple with the ProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) GetProcessingTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProcessingTime) {
		return nil, false
	}
	return o.ProcessingTime, true
}

// HasProcessingTime returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) HasProcessingTime() bool {
	if o != nil && !utils.IsNil(o.ProcessingTime) {
		return true
	}

	return false
}

// SetProcessingTime gets a reference to the given string and assigns it to the ProcessingTime field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) SetProcessingTime(v string) {
	o.ProcessingTime = &v
}

// GetVersionApi returns the VersionApi field value if set, zero value otherwise.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) GetVersionApi() string {
	if o == nil || utils.IsNil(o.VersionApi) {
		var ret string
		return ret
	}
	return *o.VersionApi
}

// GetVersionApiOk returns a tuple with the VersionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) GetVersionApiOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VersionApi) {
		return nil, false
	}
	return o.VersionApi, true
}

// HasVersionApi returns a boolean if a field has been set.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) HasVersionApi() bool {
	if o != nil && !utils.IsNil(o.VersionApi) {
		return true
	}

	return false
}

// SetVersionApi gets a reference to the given string and assigns it to the VersionApi field.
func (o *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) SetVersionApi(v string) {
	o.VersionApi = &v
}

func (o InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !utils.IsNil(o.ProcessingTime) {
		toSerialize["processing_time"] = o.ProcessingTime
	}
	if !utils.IsNil(o.VersionApi) {
		toSerialize["version_api"] = o.VersionApi
	}
	return toSerialize, nil
}

type NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination struct {
	value *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination
	isSet bool
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) Get() *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination {
	return v.value
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) Set(val *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination(val *InternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination {
	return &NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination{value: val, isSet: true}
}

func (v NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalWebControllersMerchantApiv1VirtualaccountStandardMetaWithPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
