# InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**AmountMax** | Pointer to **float32** |  | [optional] 
**AmountMin** | Pointer to **float32** |  | [optional] 
**Customer** | Pointer to [**InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer**](InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer.md) |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**IsReusable** | Pointer to **bool** |  | [optional] 
**PaymentMethodBank** | Pointer to **string** |  | [optional] 
**PaymentMethodCode** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ReferenceInternalId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VirtualAccountLinkId** | Pointer to **string** |  | [optional] 
**VirtualAccountNumber** | Pointer to **string** |  | [optional] 
**VirtualAccountSuffix** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse

`func NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse() *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse`

NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse instantiates a new InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponseWithDefaults

`func NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponseWithDefaults() *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse`

NewInternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponseWithDefaults instantiates a new InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountMax

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetAmountMax() float32`

GetAmountMax returns the AmountMax field if non-nil, zero value otherwise.

### GetAmountMaxOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetAmountMaxOk() (*float32, bool)`

GetAmountMaxOk returns a tuple with the AmountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMax

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetAmountMax(v float32)`

SetAmountMax sets AmountMax field to given value.

### HasAmountMax

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasAmountMax() bool`

HasAmountMax returns a boolean if a field has been set.

### GetAmountMin

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetAmountMin() float32`

GetAmountMin returns the AmountMin field if non-nil, zero value otherwise.

### GetAmountMinOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetAmountMinOk() (*float32, bool)`

GetAmountMinOk returns a tuple with the AmountMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMin

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetAmountMin(v float32)`

SetAmountMin sets AmountMin field to given value.

### HasAmountMin

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasAmountMin() bool`

HasAmountMin returns a boolean if a field has been set.

### GetCustomer

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetCustomer() InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetCustomerOk() (*InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetCustomer(v InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsClosed

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsReusable

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetIsReusable() bool`

GetIsReusable returns the IsReusable field if non-nil, zero value otherwise.

### GetIsReusableOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetIsReusableOk() (*bool, bool)`

GetIsReusableOk returns a tuple with the IsReusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReusable

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetIsReusable(v bool)`

SetIsReusable sets IsReusable field to given value.

### HasIsReusable

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasIsReusable() bool`

HasIsReusable returns a boolean if a field has been set.

### GetPaymentMethodBank

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetPaymentMethodBank() string`

GetPaymentMethodBank returns the PaymentMethodBank field if non-nil, zero value otherwise.

### GetPaymentMethodBankOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetPaymentMethodBankOk() (*string, bool)`

GetPaymentMethodBankOk returns a tuple with the PaymentMethodBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodBank

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetPaymentMethodBank(v string)`

SetPaymentMethodBank sets PaymentMethodBank field to given value.

### HasPaymentMethodBank

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasPaymentMethodBank() bool`

HasPaymentMethodBank returns a boolean if a field has been set.

### GetPaymentMethodCode

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetPaymentMethodCode() string`

GetPaymentMethodCode returns the PaymentMethodCode field if non-nil, zero value otherwise.

### GetPaymentMethodCodeOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetPaymentMethodCodeOk() (*string, bool)`

GetPaymentMethodCodeOk returns a tuple with the PaymentMethodCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodCode

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetPaymentMethodCode(v string)`

SetPaymentMethodCode sets PaymentMethodCode field to given value.

### HasPaymentMethodCode

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasPaymentMethodCode() bool`

HasPaymentMethodCode returns a boolean if a field has been set.

### GetReferenceId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReferenceInternalId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetReferenceInternalId() string`

GetReferenceInternalId returns the ReferenceInternalId field if non-nil, zero value otherwise.

### GetReferenceInternalIdOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetReferenceInternalIdOk() (*string, bool)`

GetReferenceInternalIdOk returns a tuple with the ReferenceInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInternalId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetReferenceInternalId(v string)`

SetReferenceInternalId sets ReferenceInternalId field to given value.

### HasReferenceInternalId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasReferenceInternalId() bool`

HasReferenceInternalId returns a boolean if a field has been set.

### GetStatus

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualAccountLinkId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetVirtualAccountLinkId() string`

GetVirtualAccountLinkId returns the VirtualAccountLinkId field if non-nil, zero value otherwise.

### GetVirtualAccountLinkIdOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetVirtualAccountLinkIdOk() (*string, bool)`

GetVirtualAccountLinkIdOk returns a tuple with the VirtualAccountLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountLinkId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetVirtualAccountLinkId(v string)`

SetVirtualAccountLinkId sets VirtualAccountLinkId field to given value.

### HasVirtualAccountLinkId

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasVirtualAccountLinkId() bool`

HasVirtualAccountLinkId returns a boolean if a field has been set.

### GetVirtualAccountNumber

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetVirtualAccountNumber() string`

GetVirtualAccountNumber returns the VirtualAccountNumber field if non-nil, zero value otherwise.

### GetVirtualAccountNumberOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetVirtualAccountNumberOk() (*string, bool)`

GetVirtualAccountNumberOk returns a tuple with the VirtualAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountNumber

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetVirtualAccountNumber(v string)`

SetVirtualAccountNumber sets VirtualAccountNumber field to given value.

### HasVirtualAccountNumber

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasVirtualAccountNumber() bool`

HasVirtualAccountNumber returns a boolean if a field has been set.

### GetVirtualAccountSuffix

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetVirtualAccountSuffix() string`

GetVirtualAccountSuffix returns the VirtualAccountSuffix field if non-nil, zero value otherwise.

### GetVirtualAccountSuffixOk

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) GetVirtualAccountSuffixOk() (*string, bool)`

GetVirtualAccountSuffixOk returns a tuple with the VirtualAccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountSuffix

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) SetVirtualAccountSuffix(v string)`

SetVirtualAccountSuffix sets VirtualAccountSuffix field to given value.

### HasVirtualAccountSuffix

`func (o *InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) HasVirtualAccountSuffix() bool`

HasVirtualAccountSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


