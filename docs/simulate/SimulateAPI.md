# doitpay\SimulateAPI

All URIs are relative to *https://api.doitpay.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulatePayment**](SimulateAPI.md#SimulatePayment) | **Post** /merchant/v1/simulate-payment | Simulate payment of the system.



## SimulatePayment

Simulate payment of the system.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    doitpay "github.com/automotechnologies/doitpay-go"
    simulate "github.com/automotechnologies/doitpay-go/simulate"
)

func main() {
    doitpayClient := doitpay.NewAPIClient("API-KEY")

    accountNumber := "12345678"
 paymentIdentifier := "sample payment identifier"

 request := simulate.InternalWebControllersMerchantApiv1SimulateSimulateRequest{ // [REQUIRED]
  AccountNumber: &accountNumber,
  PaymentIdentifier: &paymentIdentifier,
 }

    doitpayClient := doitpay.NewAPIClient("API-KEY")

    r, err := apiClient.SimulateAPI.SimulatePayment(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.SimulatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any path parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSimulatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalWebControllersMerchantApiv1SimulateSimulateRequest**](InternalWebControllersMerchantApiv1SimulateSimulateRequest.md) | Request payload to simulate payment | 

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) 
[[Back to README]](../../README.md)

