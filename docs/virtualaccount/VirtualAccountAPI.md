# doitpay\VirtualAccountAPI

All URIs are relative to *https://api.doitpay.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualAccount**](VirtualAccountAPI.md#CreateVirtualAccount) | **Post** /merchant/v1/virtual-account | Create virtual account data
[**GetVirtualAccountById**](VirtualAccountAPI.md#GetVirtualAccountById) | **Get** /merchant/v1/virtual-account/{virtualAccountId} | Fetch virtual account data by ID
[**GetVirtualAccountByNumber**](VirtualAccountAPI.md#GetVirtualAccountByNumber) | **Get** /merchant/v1/virtual-account/number/{virtualAccountNumber} | Fetch virtual account data by virtual account number
[**GetVirtualAccounts**](VirtualAccountAPI.md#GetVirtualAccounts) | **Get** /merchant/v1/virtual-account | List all created virtual account data



## CreateVirtualAccount

Create virtual account data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    virtualaccount "github.com/automotechnologies/doitpay-go/virtualaccount"
)

func main() {
    doitpayClient := doitpay.NewClient("API-KEY")

    amount := float32(200.00)
 amountMax := float32(300.00)
 amountMin := float32(100.00)
 businessId := int32(101)
 currency := "USD"
 name := "John Doe"
 email := "john.doe@example.com"
 phone := "555-555-5555"
 expirationDate := "2024-10-25"
 isClosed := false
 isReusable := true
 paymentMethodCode := "PMC789"
 referenceId := "REF12345"
 referenceInternalId := "INT67890"
 virtualAccountSuffix := "VAS456"

 customer := virtualaccount.InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer{
  Name:  &name,
  Email: &email,
  Phone: &phone,
 }

 request := virtualaccount.InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest{
  Amount:               &amount,
  AmountMax:            &amountMax,
  AmountMin:            &amountMin,
  BusinessId:           &businessId,
  Currency:             &currency,
  Customer:             &customer,
  ExpirationDate:       &expirationDate,
  IsClosed:             &isClosed,
  IsReusable:           &isReusable,
  PaymentMethodCode:    &paymentMethodCode,
  ReferenceId:          &referenceId,
  ReferenceInternalId:  &referenceInternalId,
  VirtualAccountSuffix: &virtualAccountSuffix,
 }

    resp, r, err := doitpayClient.VirtualAccountAPI.CreateVirtualAccount(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAccountAPI.CreateVirtualAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `CreateVirtualAccount`
    fmt.Fprintf(os.Stdout, "Response from `VirtualAccountAPI.CreateVirtualAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest**](InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest.md) | Request payload to create virtual account | 

### Return type

[**InternalWebControllersMerchantApiv1VirtualaccountStandardResponse**](InternalWebControllersMerchantApiv1VirtualaccountStandardResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## GetVirtualAccountById

Fetch virtual account data by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
)

func main() {
    doitpayClient := doitpay.NewClient("API-KEY")

    virtualAccountID := int32(123) // int32 | Virtual Account ID

    resp, r, err := doitpayClient.VirtualAccountAPI.GetVirtualAccountById(context.Background(), virtualAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAccountAPI.GetVirtualAccountById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `GetVirtualAccountById`
    fmt.Fprintf(os.Stdout, "Response from `VirtualAccountAPI.GetVirtualAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualAccountId** | **int32** | Virtual Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalWebControllersMerchantApiv1VirtualaccountStandardResponse**](InternalWebControllersMerchantApiv1VirtualaccountStandardResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## GetVirtualAccountByNumber

Fetch virtual account data by virtual account number



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
)

func main() {
    doitpayClient := doitpay.NewClient("API-KEY")

    virtualAccountNumber := "12345678" // string | Virtual Account Number

    resp, r, err := doitpayClient.VirtualAccountAPI.GetVirtualAccountByNumber(context.Background(), virtualAccountNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAccountAPI.GetVirtualAccountByNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `GetVirtualAccountByNumber`
    fmt.Fprintf(os.Stdout, "Response from `VirtualAccountAPI.GetVirtualAccountByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualAccountNumber** | **string** | Virtual Account Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualAccountByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalWebControllersMerchantApiv1VirtualaccountStandardResponse**](InternalWebControllersMerchantApiv1VirtualaccountStandardResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)


## GetVirtualAccounts

List all created virtual account data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
)

func main() {
    doitpayClient := doitpay.NewClient("API-KEY")

    page := int32(1) // int32 | Page number (optional) (default to 1)
    limit := int32(20) // int32 | Page limit (optional) (default to 10)
    statuses := "ACTIVE,PENDING" // string | Comma-separated list of statuses to filter by. Example: ?statuses=ACTIVE,PENDING (optional)

    resp, r, err := doitpayClient.VirtualAccountAPI.GetVirtualAccounts(context.Background()).Page(page).Limit(limit).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAccountAPI.GetVirtualAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `GetVirtualAccounts`
    fmt.Fprintf(os.Stdout, "Response from `VirtualAccountAPI.GetVirtualAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Page limit | [default to 10]
 **statuses** | **string** | Comma-separated list of statuses to filter by. Example: ?statuses&#x3D;ACTIVE,PENDING | 

### Return type

[**InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination**](InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)

