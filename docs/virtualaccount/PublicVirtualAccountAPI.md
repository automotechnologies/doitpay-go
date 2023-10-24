# \PublicVirtualAccountAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualAccount**](PublicVirtualAccountAPI.md#CreateVirtualAccount) | **Post** /merchant/v1/virtual-account | Create virtual account data.
[**GetVirtualAccountById**](PublicVirtualAccountAPI.md#GetVirtualAccountById) | **Get** /merchant/v1/virtual-account/{virtualAccountId} | Fetch virtual account data by ID.
[**GetVirtualAccountByNumber**](PublicVirtualAccountAPI.md#GetVirtualAccountByNumber) | **Get** /merchant/v1/virtual-account/number/{virtualAccountNumber} | Fetch virtual account data by virtual account number.
[**GetVirtualAccounts**](PublicVirtualAccountAPI.md#GetVirtualAccounts) | **Get** /merchant/v1/virtual-account | List all created virtual account data.



## CreateVirtualAccount

> InternalWebControllersMerchantApiv1VirtualaccountStandardResponse CreateVirtualAccount(ctx).Request(request).Execute()

Create virtual account data.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/automotechnologies/doitpay"
)

func main() {
    request := *openapiclient.NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest() // InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest | Request payload to create virtual account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicVirtualAccountAPI.CreateVirtualAccount(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicVirtualAccountAPI.CreateVirtualAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualAccount`: InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicVirtualAccountAPI.CreateVirtualAccount`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualAccountById

> InternalWebControllersMerchantApiv1VirtualaccountStandardResponse GetVirtualAccountById(ctx, virtualAccountId).Execute()

Fetch virtual account data by ID.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/automotechnologies/doitpay"
)

func main() {
    virtualAccountId := int32(56) // int32 | Virtual Account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicVirtualAccountAPI.GetVirtualAccountById(context.Background(), virtualAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicVirtualAccountAPI.GetVirtualAccountById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualAccountById`: InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicVirtualAccountAPI.GetVirtualAccountById`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualAccountByNumber

> InternalWebControllersMerchantApiv1VirtualaccountStandardResponse GetVirtualAccountByNumber(ctx, virtualAccountNumber).Execute()

Fetch virtual account data by virtual account number.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/automotechnologies/doitpay"
)

func main() {
    virtualAccountNumber := "virtualAccountNumber_example" // string | Virtual Account Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicVirtualAccountAPI.GetVirtualAccountByNumber(context.Background(), virtualAccountNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicVirtualAccountAPI.GetVirtualAccountByNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualAccountByNumber`: InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicVirtualAccountAPI.GetVirtualAccountByNumber`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualAccounts

> InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination GetVirtualAccounts(ctx).Page(page).Limit(limit).Statuses(statuses).Execute()

List all created virtual account data.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/automotechnologies/doitpay"
)

func main() {
    page := int32(56) // int32 | Page number (optional) (default to 1)
    limit := int32(56) // int32 | Page limit (optional) (default to 10)
    statuses := "statuses_example" // string | Comma-separated list of statuses to filter by. Example: ?statuses=ACTIVE,PENDING (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicVirtualAccountAPI.GetVirtualAccounts(context.Background()).Page(page).Limit(limit).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicVirtualAccountAPI.GetVirtualAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualAccounts`: InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination
    fmt.Fprintf(os.Stdout, "Response from `PublicVirtualAccountAPI.GetVirtualAccounts`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

