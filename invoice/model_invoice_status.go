package invoice

import (
	"encoding/json"

	"fmt"
)

// InvoiceStatus Representing the status of an invoice.
type InvoiceStatus string

// List of InvoiceStatus
const (
	INVOICE_STATUS_PENDING         InvoiceStatus = "PENDING"
	INVOICE_STATUS_PAID            InvoiceStatus = "PAID"
	INVOICE_STATUS_CANCELED        InvoiceStatus = "CANCELLED"
	INVOICE_STATUS_REFUNDED        InvoiceStatus = "REFUNDED"
	INVOICE_STATUS_PARTIAL_SETTLED InvoiceStatus = "PARTIAL_SETTLED"
	INVOICE_STATUS_EXPIRED         InvoiceStatus = "EXPIRED"
	INVOICE_STATUS_UNKNOWN         InvoiceStatus = "UNKNOWN"
)

// All allowed values of InvoiceStatus enum
var AllowedInvoiceStatusEnumValues = []InvoiceStatus{
	"PENDING",
	"PAID",
	"CANCELLED",
	"REFUNDED",
	"PARTIAL_SETTLED",
	"EXPIRED",
	"UNKNOWN",
}

func (v *InvoiceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvoiceStatus(value)
	for _, existing := range AllowedInvoiceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = INVOICE_STATUS_UNKNOWN
	return nil
}

// NewInvoiceStatusFromValue returns a pointer to a valid InvoiceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvoiceStatusFromValue(v string) (*InvoiceStatus, error) {
	ev := InvoiceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvoiceStatus: valid values are %v", v, AllowedInvoiceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvoiceStatus) IsValid() bool {
	for _, existing := range AllowedInvoiceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v InvoiceStatus) String() string {
	return string(v)
}

// Ptr returns reference to InvoiceStatus value
func (v InvoiceStatus) Ptr() *InvoiceStatus {
	return &v
}

type NullableInvoiceStatus struct {
	value *InvoiceStatus
	isSet bool
}

func (v NullableInvoiceStatus) Get() *InvoiceStatus {
	return v.value
}

func (v *NullableInvoiceStatus) Set(val *InvoiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceStatus(val *InvoiceStatus) *NullableInvoiceStatus {
	return &NullableInvoiceStatus{value: val, isSet: true}
}

func (v NullableInvoiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
