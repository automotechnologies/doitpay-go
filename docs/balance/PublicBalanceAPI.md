# \PublicBalanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalance**](PublicBalanceAPI.md#GetBalance) | **Get** /merchant/v1/balance | Get Balance.



## GetBalance

> InternalWebControllersMerchantApiv1BalanceStandardResponse GetBalance(ctx).Execute()

Get Balance.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicBalanceAPI.GetBalance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicBalanceAPI.GetBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalance`: InternalWebControllersMerchantApiv1BalanceStandardResponse
    fmt.Fprintf(os.Stdout, "Response from `PublicBalanceAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


### Return type

[**InternalWebControllersMerchantApiv1BalanceStandardResponse**](InternalWebControllersMerchantApiv1BalanceStandardResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

