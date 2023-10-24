# \PublicSimulateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulatePayment**](PublicSimulateAPI.md#SimulatePayment) | **Post** /merchant/v1/simulate-payment | SimulatePayment is handler for simulate payment system.



## SimulatePayment

> SimulatePayment(ctx).Request(request).Execute()

SimulatePayment is handler for simulate payment system.



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
    request := *openapiclient.NewInternalWebControllersMerchantApiv1SimulateSimulateRequest() // InternalWebControllersMerchantApiv1SimulateSimulateRequest | Request payload to simulate payment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PublicSimulateAPI.SimulatePayment(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSimulateAPI.SimulatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalWebControllersMerchantApiv1SimulateSimulateRequest**](InternalWebControllersMerchantApiv1SimulateSimulateRequest.md) | Request payload to simulate payment | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

