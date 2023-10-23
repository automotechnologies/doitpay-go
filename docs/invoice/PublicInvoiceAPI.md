# \PublicInvoiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoice**](PublicInvoiceAPI.md#CreateInvoice) | **Post** /merchant/v1/invoice | Create invoice for a payment.
[**DownloadInvoice**](PublicInvoiceAPI.md#DownloadInvoice) | **Get** /merchant/v1/invoice/download/csv | Download CSV file.
[**ExpireInvoice**](PublicInvoiceAPI.md#ExpireInvoice) | **Post** /merchant/v1/invoice/{invoice_id}/expire | Expire invoice by invoice ID.
[**GetInvoiceById**](PublicInvoiceAPI.md#GetInvoiceById) | **Get** /merchant/v1/invoice/{invoice_id} | Fetch invoice data by ID.
[**GetInvoices**](PublicInvoiceAPI.md#GetInvoices) | **Get** /merchant/v1/invoice | List all created invoice data
[**GetPaymentMethodById**](PublicInvoiceAPI.md#GetPaymentMethodById) | **Get** /merchant/v1/invoice/{invoice_id}/payment-method | Fetch payment method by invoice ID.
[**UpdateInvoiceById**](PublicInvoiceAPI.md#UpdateInvoiceById) | **Put** /merchant/v1/invoice/{invoice_id}/update | Update pending invoice



## CreateInvoice

> InternalWebControllersMerchantApiv1InvoiceStandardResponse CreateInvoice(ctx).Request(request).Execute()

Create invoice for a payment.



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
    request := *openapiclient.NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequest() // InternalWebControllersMerchantApiv1InvoiceInvoiceRequest | Request payload to create invoice

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicInvoiceAPI.CreateInvoice(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicInvoiceAPI.CreateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoice`: InternalWebControllersMerchantApiv1InvoiceStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicInvoiceAPI.CreateInvoice`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadInvoice

> string DownloadInvoice(ctx).Page(page).Limit(limit).Statuses(statuses).Execute()

Download CSV file.



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
    resp, r, err := apiClient.PublicInvoiceAPI.DownloadInvoice(context.Background()).Page(page).Limit(limit).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicInvoiceAPI.DownloadInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadInvoice`: string
    fmt.Fprintf(os.Stdout, "Response from `PublicInvoiceAPI.DownloadInvoice`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireInvoice

> InternalWebControllersMerchantApiv1InvoiceStandardResponse ExpireInvoice(ctx, invoiceId).Execute()

Expire invoice by invoice ID.



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
    invoiceId := "invoiceId_example" // string | Invoice ID to fetch payment method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicInvoiceAPI.ExpireInvoice(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicInvoiceAPI.ExpireInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpireInvoice`: InternalWebControllersMerchantApiv1InvoiceStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicInvoiceAPI.ExpireInvoice`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceById

> InternalWebControllersMerchantApiv1InvoiceStandardResponse GetInvoiceById(ctx, invoiceId).Execute()

Fetch invoice data by ID.



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
    invoiceId := "invoiceId_example" // string | Invoice ID to fetch data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicInvoiceAPI.GetInvoiceById(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicInvoiceAPI.GetInvoiceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceById`: InternalWebControllersMerchantApiv1InvoiceStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicInvoiceAPI.GetInvoiceById`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination GetInvoices(ctx).Page(page).Limit(limit).Statuses(statuses).Execute()

List all created invoice data



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
    resp, r, err := apiClient.PublicInvoiceAPI.GetInvoices(context.Background()).Page(page).Limit(limit).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicInvoiceAPI.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination
    fmt.Fprintf(os.Stdout, "Response from `PublicInvoiceAPI.GetInvoices`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethodById

> InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse GetPaymentMethodById(ctx, invoiceId).Execute()

Fetch payment method by invoice ID.



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
    invoiceId := "invoiceId_example" // string | Invoice ID to fetch payment method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicInvoiceAPI.GetPaymentMethodById(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicInvoiceAPI.GetPaymentMethodById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentMethodById`: InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicInvoiceAPI.GetPaymentMethodById`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceById

> InternalWebControllersMerchantApiv1InvoiceStandardResponse UpdateInvoiceById(ctx, invoiceId).Request(request).Execute()

Update pending invoice



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
    invoiceId := "invoiceId_example" // string | invoice id to update
    request := *openapiclient.NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequest() // InternalWebControllersMerchantApiv1InvoiceInvoiceRequest | Invoice Request Body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicInvoiceAPI.UpdateInvoiceById(context.Background(), invoiceId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicInvoiceAPI.UpdateInvoiceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInvoiceById`: InternalWebControllersMerchantApiv1InvoiceStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicInvoiceAPI.UpdateInvoiceById`: %v\n", resp)
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

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

