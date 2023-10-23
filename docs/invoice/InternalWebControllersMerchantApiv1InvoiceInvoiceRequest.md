# InternalWebControllersMerchantApiv1InvoiceInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**AmountCurrency** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**InternalWebControllersMerchantApiv1InvoiceCustomer**](InternalWebControllersMerchantApiv1InvoiceCustomer.md) |  | [optional] 
**CustomerNotification** | Pointer to **[]string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]InternalWebControllersMerchantApiv1InvoiceItem**](InternalWebControllersMerchantApiv1InvoiceItem.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**PaymentMethods** | Pointer to [**[]InternalWebControllersMerchantApiv1InvoicePaymentMethod**](InternalWebControllersMerchantApiv1InvoicePaymentMethod.md) |  | [optional] 
**RedirectUrl** | Pointer to [**InternalWebControllersMerchantApiv1InvoiceRedirectURL**](InternalWebControllersMerchantApiv1InvoiceRedirectURL.md) |  | [optional] 
**TimeBegin** | Pointer to **string** |  | [optional] 
**TimeEnd** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequest

`func NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequest() *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest`

NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequest instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequestWithDefaults

`func NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequestWithDefaults() *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest`

NewInternalWebControllersMerchantApiv1InvoiceInvoiceRequestWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmountCurrency() string`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetAmountCurrencyOk() (*string, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetAmountCurrency(v string)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetCountry

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCustomer

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomer() InternalWebControllersMerchantApiv1InvoiceCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomerOk() (*InternalWebControllersMerchantApiv1InvoiceCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetCustomer(v InternalWebControllersMerchantApiv1InvoiceCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerNotification

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomerNotification() []string`

GetCustomerNotification returns the CustomerNotification field if non-nil, zero value otherwise.

### GetCustomerNotificationOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetCustomerNotificationOk() (*[]string, bool)`

GetCustomerNotificationOk returns a tuple with the CustomerNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotification

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetCustomerNotification(v []string)`

SetCustomerNotification sets CustomerNotification field to given value.

### HasCustomerNotification

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasCustomerNotification() bool`

HasCustomerNotification returns a boolean if a field has been set.

### GetExternalId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetItems

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetItems() []InternalWebControllersMerchantApiv1InvoiceItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetItemsOk() (*[]InternalWebControllersMerchantApiv1InvoiceItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetItems(v []InternalWebControllersMerchantApiv1InvoiceItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNotes

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetPaymentMethods() []InternalWebControllersMerchantApiv1InvoicePaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetPaymentMethodsOk() (*[]InternalWebControllersMerchantApiv1InvoicePaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetPaymentMethods(v []InternalWebControllersMerchantApiv1InvoicePaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetRedirectUrl() InternalWebControllersMerchantApiv1InvoiceRedirectURL`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetRedirectUrlOk() (*InternalWebControllersMerchantApiv1InvoiceRedirectURL, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetRedirectUrl(v InternalWebControllersMerchantApiv1InvoiceRedirectURL)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetTimeBegin

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeBegin() string`

GetTimeBegin returns the TimeBegin field if non-nil, zero value otherwise.

### GetTimeBeginOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeBeginOk() (*string, bool)`

GetTimeBeginOk returns a tuple with the TimeBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBegin

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetTimeBegin(v string)`

SetTimeBegin sets TimeBegin field to given value.

### HasTimeBegin

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasTimeBegin() bool`

HasTimeBegin returns a boolean if a field has been set.

### GetTimeEnd

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


