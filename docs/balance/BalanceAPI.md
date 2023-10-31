# doitpay\BalanceAPI

All URIs are relative to *https://api.doitpay.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalance**](BalanceAPI.md#GetBalance) | **Get** /merchant/v1/balance | Obtains the balances for a business.



## GetBalance

Obtains the balances for a business.



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
    doitpayClient := doitpay.NewAPIClient("API-KEY")

    resp, r, err := doitpayClient.BalanceAPI.GetBalance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.GetBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

    // response from `GetBalance`
    fmt.Fprintf(os.Stdout, "Response from `BalanceAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern

### Return type

[**Balance**](InternalWebControllersMerchantApiv1BalanceStandardResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) 
[[Back to README]](../../README.md)

