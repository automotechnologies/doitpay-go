package virtualaccount

import (
	"encoding/json"

	"fmt"
)

// VirtualAccountStatus Representing the status of a virtual account.
type VirtualAccountStatus string

// List of VirtualAccountStatus
const (
	VIRTUAL_ACCOUNT_STATUS_UNKNOWN  VirtualAccountStatus = "UNKNOWN"
	VIRTUAL_ACCOUNT_STATUS_INACTIVE VirtualAccountStatus = "INACTIVE"
	VIRTUAL_ACCOUNT_STATUS_PENDING  VirtualAccountStatus = "PENDING"
	VIRTUAL_ACCOUNT_STATUS_ACTIVE   VirtualAccountStatus = "ACTIVE"
	VIRTUAL_ACCOUNT_STATUS_PAID     VirtualAccountStatus = "PAID"
	VIRTUAL_ACCOUNT_STATUS_EXPIRED  VirtualAccountStatus = "EXPIRED"
)

// All allowed values of VirtualAccountStatus enum
var AllowedVirtualAccountStatusEnumValues = []VirtualAccountStatus{
	"UNKNOWN",
	"INACTIVE",
	"PENDING",
	"ACTIVE",
	"PAID",
	"EXPIRED",
}

func (v *VirtualAccountStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualAccountStatus(value)
	for _, existing := range AllowedVirtualAccountStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = VIRTUAL_ACCOUNT_STATUS_UNKNOWN
	return nil
}

// NewVirtualAccountStatusFromValue returns a pointer to a valid VirtualAccountStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualAccountStatusFromValue(v string) (*VirtualAccountStatus, error) {
	ev := VirtualAccountStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualAccountStatus: valid values are %v", v, AllowedVirtualAccountStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualAccountStatus) IsValid() bool {
	for _, existing := range AllowedVirtualAccountStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v VirtualAccountStatus) String() string {
	return string(v)
}

// Ptr returns reference to VirtualAccountStatus value
func (v VirtualAccountStatus) Ptr() *VirtualAccountStatus {
	return &v
}

type NullableVirtualAccountStatus struct {
	value *VirtualAccountStatus
	isSet bool
}

func (v NullableVirtualAccountStatus) Get() *VirtualAccountStatus {
	return v.value
}

func (v *NullableVirtualAccountStatus) Set(val *VirtualAccountStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountStatus(val *VirtualAccountStatus) *NullableVirtualAccountStatus {
	return &NullableVirtualAccountStatus{value: val, isSet: true}
}

func (v NullableVirtualAccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
