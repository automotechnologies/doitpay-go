# doitpay\InvoiceAPI

All URIs are relative to *https://api.doitpay.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoice**](InvoiceAPI.md#CreateInvoice) | **Post** /merchant/v1/invoice | Create invoice for a payment
[**DownloadInvoice**](InvoiceAPI.md#DownloadInvoice) | **Get** /merchant/v1/invoice/download/csv | Download invoices as a CSV file
[**ExpireInvoice**](InvoiceAPI.md#ExpireInvoice) | **Post** /merchant/v1/invoice/{invoice_id}/expire | Expire an invoice by invoice ID
[**GetInvoiceById**](InvoiceAPI.md#GetInvoiceById) | **Get** /merchant/v1/invoice/{invoice_id} | Fetch invoice data by invoice ID
[**GetInvoices**](InvoiceAPI.md#GetInvoices) | **Get** /merchant/v1/invoice | List all created invoices data
[**GetPaymentMethodById**](InvoiceAPI.md#GetPaymentMethodById) | **Get** /merchant/v1/invoice/{invoice_id}/payment-method | Fetch payment method by invoice ID
[**UpdateInvoiceById**](InvoiceAPI.md#UpdateInvoiceById) | **Put** /merchant/v1/invoice/{invoice_id}/update | Update a pending invoice data



## CreateInvoice

Create invoice for a payment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    doitpay "github.com/automotechnologies/doitpay-go"
    invoice "github.com/automotechnologies/doitpay-go/invoice"
)

func main() {
    doitpayClient := doitpay.NewClient("API-KEY")

    amount := float32(100.50)
	amountCurrency := "USD"
	country := "US"
	addresses := "123 Main St"
	city := "Springfield"
	customerCountry := "US"
	customerRefId := "CUST12345"
	email := "john.doe@example.com"
	name := "John Doe"
	customerNotes := "Preferred customer"
	phone := "555-555-5555"
	postalCode := int32(62704)

	customer := &invoice.InternalWebControllersMerchantApiv1InvoiceCustomer{ // [OPTIONAL]
		Addresses:     &addresses,
		City:          &city,
		Country:       &customerCountry,
		CustomerRefId: &customerRefId,
		Email:         &email,
		Name:          &name,
		Notes:         &customerNotes,
		Phone:         &phone,
		PostalCode:    &postalCode,
	}

	customerNotification := []string{"email", "sms"}
	externalId := "EXT12345"
	sku1 := "SKU123"
	itemName1 := "Item1"
	itemNotes1 := "This is a test item"
	price1 := float32(20.00)
	quantity1 := int32(2)

	item1 := invoice.InternalWebControllersMerchantApiv1InvoiceItem{
		SKU:      &sku1,
		Name:     &itemName1,
		Notes:    &itemNotes1,
		Price:    &price1,
		Quantity: &quantity1,
	}

	sku2 := "SKU456"
	itemName2 := "Item2"
	itemNotes2 := "This is another test item"
	price2 := float32(15.00)
	quantity2 := int32(3)

	item2 := invoice.InternalWebControllersMerchantApiv1InvoiceItem{
		SKU:      &sku2,
		Name:     &itemName2,
		Notes:    &itemNotes2,
		Price:    &price2,
		Quantity: &quantity2,
	}

	items := []invoice.InternalWebControllersMerchantApiv1InvoiceItem{item1, item2}

	notes := "Thank you for your business."
	code := "PM123"
	status := "active"
	paymentType := "credit_card"

	paymentMethod := invoice.InternalWebControllersMerchantApiv1InvoicePaymentMethod{
		Code:   &code,
		Status: &status,
		Type:   &paymentType,
	}

	paymentMethods := []invoice.InternalWebControllersMerchantApiv1InvoicePaymentMethod{paymentMethod}

	cancelUrl := "https://example.com/cancel"
	successUrl := "https://example.com/success"

	redirectUrl := &invoice.InternalWebControllersMerchantApiv1InvoiceRedirectURL{
		Cancel:  &cancelUrl,
		Success: &successUrl,
	}

	timeBegin := "2023-10-25T12:00:00Z" // [OPTIONAL]
	timeEnd := "2023-10-25T18:00:00Z" // [OPTIONAL]

	request := invoice.InternalWebControllersMerchantApiv1InvoiceInvoiceRequest{
		Amount:               &amount,
		AmountCurrency:       &amountCurrency,
		Country:              &country,
		Customer:             customer,
		CustomerNotification: customerNotification,
		ExternalId:           &externalId,
		Items:                items,
		Notes:                &notes,
		PaymentMethods:       paymentMethods,
		RedirectUrl:          redirectUrl,
		TimeBegin:            &timeBegin,
		TimeEnd:              &timeEnd,
	}

	resp, r, err := doitpayClient.InvoiceAPI.CreateInvoice(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.CreateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

    // response from `CreateInvoice`
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalWebControllersMerchantApiv1InvoiceInvoiceRequest**](InternalWebControllersMerchantApiv1InvoiceInvoiceRequest.md) | Request payload to create invoice | 

### Return type

[**InternalWebControllersMerchantApiv1InvoiceStandardResponse**](InternalWebControllersMerchantApiv1InvoiceStandardResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## DownloadInvoice

Download invoices as a CSV file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    invoice "github.com/automotechnologies/doitpay-go/invoice"
)

func main() {
    doitpayClient := doitpay.NewAPIClient("API-KEY")

    page := int32(56) // int32 | Page number (optional) (default to 1)
    limit := int32(56) // int32 | Page limit (optional) (default to 10)
    statuses := "ACTIVE,PENDING" // string | Comma-separated list of statuses to filter by. Example: ?statuses=ACTIVE,PENDING (optional)

    resp, r, err := doitpayClient.InvoiceAPI.DownloadInvoice(context.Background()).Page(page).Limit(limit).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.DownloadInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `DownloadInvoice`
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.DownloadInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Page limit | [default to 10]
 **statuses** | **string** | Comma-separated list of statuses to filter by. Example: ?statuses&#x3D;ACTIVE,PENDING | 

### Return type

**string**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## ExpireInvoice

Expire an invoice by invoice ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    invoice "github.com/automotechnologies/doitpay-go/invoice"
)

func main() {
    doitpayClient := doitpay.NewAPIClient("API-KEY")

    invoiceID := "123" // string | Invoice ID to fetch payment method

    resp, r, err := apiClient.InvoiceAPI.ExpireInvoice(context.Background(), invoiceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.ExpireInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `ExpireInvoice`
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.ExpireInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Invoice ID to fetch payment method | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalWebControllersMerchantApiv1InvoiceStandardResponse**](InternalWebControllersMerchantApiv1InvoiceStandardResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## GetInvoiceById

Fetch invoice data by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    invoice "github.com/automotechnologies/doitpay-go/invoice"
)

func main() {
    doitpayClient := doitpay.NewAPIClient("API-KEY")

    invoiceID := "123" // string | Invoice ID to fetch data

    resp, r, err := doitpayClient.InvoiceAPI.GetInvoiceById(context.Background(), invoiceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetInvoiceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `GetInvoiceById`
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.GetInvoiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Invoice ID to fetch data | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalWebControllersMerchantApiv1InvoiceStandardResponse**](InternalWebControllersMerchantApiv1InvoiceStandardResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## GetInvoices

List all created invoice data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    invoice "github.com/automotechnologies/doitpay-go/invoice"
)

func main() {
    doitpayClient := doitpay.NewAPIClient("API-KEY")

    page := int32(56) // int32 | Page number (optional) (default to 1)
    limit := int32(56) // int32 | Page limit (optional) (default to 10)
    statuses := "ACTIVE,PENDING" // string | Comma-separated list of statuses to filter by. Example: ?statuses=ACTIVE,PENDING (optional)

    resp, r, err := doitpayClient.InvoiceAPI.GetInvoices(context.Background()).Page(page).Limit(limit).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `GetInvoices`
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Page limit | [default to 10]
 **statuses** | **string** | Comma-separated list of statuses to filter by. Example: ?statuses&#x3D;ACTIVE,PENDING | 

### Return type

[**InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination**](InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## GetPaymentMethodById

Fetch payment method by invoice ID.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    invoice "github.com/automotechnologies/doitpay-go/invoice"
)

func main() {
    doitpayClient := doitpay.NewAPIClient("API-KEY")

    invoiceID := "123" // string | Invoice ID to fetch payment method

    resp, r, err := doitpayClient.InvoiceAPI.GetPaymentMethodById(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.GetPaymentMethodById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `GetPaymentMethodById`
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.GetPaymentMethodById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Invoice ID to fetch payment method | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse**](InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## UpdateInvoiceById

Update pending invoice data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    invoice "github.com/automotechnologies/doitpay-go/invoice"
)

func main() {
    invoiceID := "123" // string | invoice id to update

    amount := float32(100.50)
	amountCurrency := "USD"
	country := "US"
	addresses := "123 Main St"
	city := "Springfield"
	customerCountry := "US"
	customerRefId := "CUST12345"
	email := "john.doe@example.com"
	name := "John Doe"
	customerNotes := "Preferred customer"
	phone := "555-555-5555"
	postalCode := int32(62704)

	customer := &invoice.InternalWebControllersMerchantApiv1InvoiceCustomer{ // [OPTIONAL]
		Addresses:     &addresses,
		City:          &city,
		Country:       &customerCountry,
		CustomerRefId: &customerRefId,
		Email:         &email,
		Name:          &name,
		Notes:         &customerNotes,
		Phone:         &phone,
		PostalCode:    &postalCode,
	}

	customerNotification := []string{"email", "sms"}
	externalId := "EXT12345"
	sku1 := "SKU123"
	itemName1 := "Item1"
	itemNotes1 := "This is a test item"
	price1 := float32(20.00)
	quantity1 := int32(2)

	item1 := invoice.InternalWebControllersMerchantApiv1InvoiceItem{
		SKU:      &sku1,
		Name:     &itemName1,
		Notes:    &itemNotes1,
		Price:    &price1,
		Quantity: &quantity1,
	}

	sku2 := "SKU456"
	itemName2 := "Item2"
	itemNotes2 := "This is another test item"
	price2 := float32(15.00)
	quantity2 := int32(3)

	item2 := invoice.InternalWebControllersMerchantApiv1InvoiceItem{
		SKU:      &sku2,
		Name:     &itemName2,
		Notes:    &itemNotes2,
		Price:    &price2,
		Quantity: &quantity2,
	}

	items := []invoice.InternalWebControllersMerchantApiv1InvoiceItem{item1, item2}

	notes := "Thank you for your business."
	code := "PM123"
	status := "active"
	paymentType := "credit_card"

	paymentMethod := invoice.InternalWebControllersMerchantApiv1InvoicePaymentMethod{
		Code:   &code,
		Status: &status,
		Type:   &paymentType,
	}

	paymentMethods := []invoice.InternalWebControllersMerchantApiv1InvoicePaymentMethod{paymentMethod}

	cancelUrl := "https://example.com/cancel"
	successUrl := "https://example.com/success"

	redirectUrl := &invoice.InternalWebControllersMerchantApiv1InvoiceRedirectURL{
		Cancel:  &cancelUrl,
		Success: &successUrl,
	}

	timeBegin := "2023-10-25T12:00:00Z" // [OPTIONAL]
	timeEnd := "2023-10-25T18:00:00Z" // [OPTIONAL]

	request := invoice.InternalWebControllersMerchantApiv1InvoiceInvoiceRequest{
		Amount:               &amount,
		AmountCurrency:       &amountCurrency,
		Country:              &country,
		Customer:             customer,
		CustomerNotification: customerNotification,
		ExternalId:           &externalId,
		Items:                items,
		Notes:                &notes,
		PaymentMethods:       paymentMethods,
		RedirectUrl:          redirectUrl,
		TimeBegin:            &timeBegin,
		TimeEnd:              &timeEnd,
	}

    doitpayClient := doitpay.NewClient("API-KEY")

    resp, r, err := doitpayClient.InvoiceAPI.UpdateInvoiceById(context.Background(), invoiceID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.UpdateInvoiceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `UpdateInvoiceById`
    fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.UpdateInvoiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | invoice id to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**InternalWebControllersMerchantApiv1InvoiceInvoiceRequest**](InternalWebControllersMerchantApiv1InvoiceInvoiceRequest.md) | Invoice Request Body | 

### Return type

[**InternalWebControllersMerchantApiv1InvoiceStandardResponse**](InternalWebControllersMerchantApiv1InvoiceStandardResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)

