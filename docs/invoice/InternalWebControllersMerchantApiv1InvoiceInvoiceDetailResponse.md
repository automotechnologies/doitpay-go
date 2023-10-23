# InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**AmountCurrency** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**InternalWebControllersMerchantApiv1InvoiceCustomer**](InternalWebControllersMerchantApiv1InvoiceCustomer.md) |  | [optional] 
**CustomerNotification** | Pointer to **[]string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**InvoiceLinkId** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]InternalWebControllersMerchantApiv1InvoiceItem**](InternalWebControllersMerchantApiv1InvoiceItem.md) |  | [optional] 
**MerchantLogo** | Pointer to **string** |  | [optional] 
**MerchantName** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**PaymentMethods** | Pointer to [**[]InternalWebControllersMerchantApiv1InvoicePaymentMethod**](InternalWebControllersMerchantApiv1InvoicePaymentMethod.md) |  | [optional] 
**RedirectUrl** | Pointer to [**InternalWebControllersMerchantApiv1InvoiceRedirectURL**](InternalWebControllersMerchantApiv1InvoiceRedirectURL.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeBegin** | Pointer to **string** |  | [optional] 
**TimeEnd** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse

`func NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse() *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse`

NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponseWithDefaults

`func NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponseWithDefaults() *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse`

NewInternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponseWithDefaults instantiates a new InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountCurrency

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmountCurrency() string`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetAmountCurrencyOk() (*string, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetAmountCurrency(v string)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### GetCountry

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCustomer

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomer() InternalWebControllersMerchantApiv1InvoiceCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomerOk() (*InternalWebControllersMerchantApiv1InvoiceCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetCustomer(v InternalWebControllersMerchantApiv1InvoiceCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerNotification

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomerNotification() []string`

GetCustomerNotification returns the CustomerNotification field if non-nil, zero value otherwise.

### GetCustomerNotificationOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetCustomerNotificationOk() (*[]string, bool)`

GetCustomerNotificationOk returns a tuple with the CustomerNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotification

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetCustomerNotification(v []string)`

SetCustomerNotification sets CustomerNotification field to given value.

### HasCustomerNotification

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasCustomerNotification() bool`

HasCustomerNotification returns a boolean if a field has been set.

### GetExternalId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInvoiceLinkId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetInvoiceLinkId() string`

GetInvoiceLinkId returns the InvoiceLinkId field if non-nil, zero value otherwise.

### GetInvoiceLinkIdOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetInvoiceLinkIdOk() (*string, bool)`

GetInvoiceLinkIdOk returns a tuple with the InvoiceLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLinkId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetInvoiceLinkId(v string)`

SetInvoiceLinkId sets InvoiceLinkId field to given value.

### HasInvoiceLinkId

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasInvoiceLinkId() bool`

HasInvoiceLinkId returns a boolean if a field has been set.

### GetItems

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetItems() []InternalWebControllersMerchantApiv1InvoiceItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetItemsOk() (*[]InternalWebControllersMerchantApiv1InvoiceItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetItems(v []InternalWebControllersMerchantApiv1InvoiceItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMerchantLogo

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantLogo() string`

GetMerchantLogo returns the MerchantLogo field if non-nil, zero value otherwise.

### GetMerchantLogoOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantLogoOk() (*string, bool)`

GetMerchantLogoOk returns a tuple with the MerchantLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLogo

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetMerchantLogo(v string)`

SetMerchantLogo sets MerchantLogo field to given value.

### HasMerchantLogo

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasMerchantLogo() bool`

HasMerchantLogo returns a boolean if a field has been set.

### GetMerchantName

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetNotes

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetPaymentMethods() []InternalWebControllersMerchantApiv1InvoicePaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetPaymentMethodsOk() (*[]InternalWebControllersMerchantApiv1InvoicePaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetPaymentMethods(v []InternalWebControllersMerchantApiv1InvoicePaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetRedirectUrl() InternalWebControllersMerchantApiv1InvoiceRedirectURL`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetRedirectUrlOk() (*InternalWebControllersMerchantApiv1InvoiceRedirectURL, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetRedirectUrl(v InternalWebControllersMerchantApiv1InvoiceRedirectURL)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeBegin

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeBegin() string`

GetTimeBegin returns the TimeBegin field if non-nil, zero value otherwise.

### GetTimeBeginOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeBeginOk() (*string, bool)`

GetTimeBeginOk returns a tuple with the TimeBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBegin

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetTimeBegin(v string)`

SetTimeBegin sets TimeBegin field to given value.

### HasTimeBegin

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasTimeBegin() bool`

HasTimeBegin returns a boolean if a field has been set.

### GetTimeEnd

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


