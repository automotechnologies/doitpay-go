# InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**AmountMax** | Pointer to **float32** |  | [optional] 
**AmountMin** | Pointer to **float32** |  | [optional] 
**BusinessId** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer**](InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer.md) |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**IsReusable** | Pointer to **bool** |  | [optional] 
**PaymentMethodCode** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ReferenceInternalId** | Pointer to **string** |  | [optional] 
**VirtualAccountSuffix** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest

`func NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest() *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest`

NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest instantiates a new InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequestWithDefaults

`func NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequestWithDefaults() *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest`

NewInternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequestWithDefaults instantiates a new InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountMax

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMax() float32`

GetAmountMax returns the AmountMax field if non-nil, zero value otherwise.

### GetAmountMaxOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMaxOk() (*float32, bool)`

GetAmountMaxOk returns a tuple with the AmountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMax

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetAmountMax(v float32)`

SetAmountMax sets AmountMax field to given value.

### HasAmountMax

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasAmountMax() bool`

HasAmountMax returns a boolean if a field has been set.

### GetAmountMin

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMin() float32`

GetAmountMin returns the AmountMin field if non-nil, zero value otherwise.

### GetAmountMinOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetAmountMinOk() (*float32, bool)`

GetAmountMinOk returns a tuple with the AmountMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMin

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetAmountMin(v float32)`

SetAmountMin sets AmountMin field to given value.

### HasAmountMin

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasAmountMin() bool`

HasAmountMin returns a boolean if a field has been set.

### GetBusinessId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetBusinessId() int32`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetBusinessIdOk() (*int32, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetBusinessId(v int32)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomer

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCustomer() InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetCustomerOk() (*InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetCustomer(v InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsClosed

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsReusable

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsReusable() bool`

GetIsReusable returns the IsReusable field if non-nil, zero value otherwise.

### GetIsReusableOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetIsReusableOk() (*bool, bool)`

GetIsReusableOk returns a tuple with the IsReusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReusable

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetIsReusable(v bool)`

SetIsReusable sets IsReusable field to given value.

### HasIsReusable

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasIsReusable() bool`

HasIsReusable returns a boolean if a field has been set.

### GetPaymentMethodCode

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetPaymentMethodCode() string`

GetPaymentMethodCode returns the PaymentMethodCode field if non-nil, zero value otherwise.

### GetPaymentMethodCodeOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetPaymentMethodCodeOk() (*string, bool)`

GetPaymentMethodCodeOk returns a tuple with the PaymentMethodCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodCode

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetPaymentMethodCode(v string)`

SetPaymentMethodCode sets PaymentMethodCode field to given value.

### HasPaymentMethodCode

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasPaymentMethodCode() bool`

HasPaymentMethodCode returns a boolean if a field has been set.

### GetReferenceId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReferenceInternalId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceInternalId() string`

GetReferenceInternalId returns the ReferenceInternalId field if non-nil, zero value otherwise.

### GetReferenceInternalIdOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetReferenceInternalIdOk() (*string, bool)`

GetReferenceInternalIdOk returns a tuple with the ReferenceInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInternalId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetReferenceInternalId(v string)`

SetReferenceInternalId sets ReferenceInternalId field to given value.

### HasReferenceInternalId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasReferenceInternalId() bool`

HasReferenceInternalId returns a boolean if a field has been set.

### GetVirtualAccountSuffix

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetVirtualAccountSuffix() string`

GetVirtualAccountSuffix returns the VirtualAccountSuffix field if non-nil, zero value otherwise.

### GetVirtualAccountSuffixOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) GetVirtualAccountSuffixOk() (*string, bool)`

GetVirtualAccountSuffixOk returns a tuple with the VirtualAccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountSuffix

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) SetVirtualAccountSuffix(v string)`

SetVirtualAccountSuffix sets VirtualAccountSuffix field to given value.

### HasVirtualAccountSuffix

`func (o *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) HasVirtualAccountSuffix() bool`

HasVirtualAccountSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


