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

// checks if the GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader type satisfies the utils.MappedNullable interface at compile time
var _ utils.MappedNullable = &GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader{}

// GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader struct for GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader
type GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader struct {
	CurrentPage *int32 `json:"current_page,omitempty"`
	PerPage     *int32 `json:"per_page,omitempty"`
	TotalData   *int32 `json:"total_data,omitempty"`
	TotalPages  *int32 `json:"total_pages,omitempty"`
}

// NewGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader instantiates a new GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader() *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader {
	this := GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader{}
	return &this
}

// NewGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeaderWithDefaults instantiates a new GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeaderWithDefaults() *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader {
	this := GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetCurrentPage() int32 {
	if o == nil || utils.IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetCurrentPageOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) HasCurrentPage() bool {
	if o != nil && !utils.IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetPerPage() int32 {
	if o == nil || utils.IsNil(o.PerPage) {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetPerPageOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) HasPerPage() bool {
	if o != nil && !utils.IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) SetPerPage(v int32) {
	o.PerPage = &v
}

// GetTotalData returns the TotalData field value if set, zero value otherwise.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetTotalData() int32 {
	if o == nil || utils.IsNil(o.TotalData) {
		var ret int32
		return ret
	}
	return *o.TotalData
}

// GetTotalDataOk returns a tuple with the TotalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetTotalDataOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalData) {
		return nil, false
	}
	return o.TotalData, true
}

// HasTotalData returns a boolean if a field has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) HasTotalData() bool {
	if o != nil && !utils.IsNil(o.TotalData) {
		return true
	}

	return false
}

// SetTotalData gets a reference to the given int32 and assigns it to the TotalData field.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) SetTotalData(v int32) {
	o.TotalData = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetTotalPages() int32 {
	if o == nil || utils.IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) GetTotalPagesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) HasTotalPages() bool {
	if o != nil && !utils.IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CurrentPage) {
		toSerialize["current_page"] = o.CurrentPage
	}
	if !utils.IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	if !utils.IsNil(o.TotalData) {
		toSerialize["total_data"] = o.TotalData
	}
	if !utils.IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	return toSerialize, nil
}

type NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader struct {
	value *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader
	isSet bool
}

func (v NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) Get() *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader {
	return v.value
}

func (v *NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) Set(val *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader(val *GithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) *NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader {
	return &NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader{value: val, isSet: true}
}

func (v NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComAutomotechnologiesDoitpayInternalWebRenderAppPaginationHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
