# FullBankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Longcode** | Pointer to **string** |  | [optional] 
**SupportsTransfer** | Pointer to **bool** |  | [optional] 
**AvailableForDirectDebit** | Pointer to **bool** |  | [optional] 
**PayWithBank** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFullBankDetails

`func NewFullBankDetails() *FullBankDetails`

NewFullBankDetails instantiates a new FullBankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullBankDetailsWithDefaults

`func NewFullBankDetailsWithDefaults() *FullBankDetails`

NewFullBankDetailsWithDefaults instantiates a new FullBankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullBankDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullBankDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullBankDetails) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FullBankDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *FullBankDetails) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FullBankDetails) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FullBankDetails) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FullBankDetails) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetCode

`func (o *FullBankDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FullBankDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FullBankDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FullBankDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *FullBankDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullBankDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullBankDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FullBankDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FullBankDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullBankDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullBankDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FullBankDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCountry

`func (o *FullBankDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FullBankDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FullBankDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FullBankDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *FullBankDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FullBankDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FullBankDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FullBankDetails) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLongcode

`func (o *FullBankDetails) GetLongcode() string`

GetLongcode returns the Longcode field if non-nil, zero value otherwise.

### GetLongcodeOk

`func (o *FullBankDetails) GetLongcodeOk() (*string, bool)`

GetLongcodeOk returns a tuple with the Longcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongcode

`func (o *FullBankDetails) SetLongcode(v string)`

SetLongcode sets Longcode field to given value.

### HasLongcode

`func (o *FullBankDetails) HasLongcode() bool`

HasLongcode returns a boolean if a field has been set.

### GetSupportsTransfer

`func (o *FullBankDetails) GetSupportsTransfer() bool`

GetSupportsTransfer returns the SupportsTransfer field if non-nil, zero value otherwise.

### GetSupportsTransferOk

`func (o *FullBankDetails) GetSupportsTransferOk() (*bool, bool)`

GetSupportsTransferOk returns a tuple with the SupportsTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTransfer

`func (o *FullBankDetails) SetSupportsTransfer(v bool)`

SetSupportsTransfer sets SupportsTransfer field to given value.

### HasSupportsTransfer

`func (o *FullBankDetails) HasSupportsTransfer() bool`

HasSupportsTransfer returns a boolean if a field has been set.

### GetAvailableForDirectDebit

`func (o *FullBankDetails) GetAvailableForDirectDebit() bool`

GetAvailableForDirectDebit returns the AvailableForDirectDebit field if non-nil, zero value otherwise.

### GetAvailableForDirectDebitOk

`func (o *FullBankDetails) GetAvailableForDirectDebitOk() (*bool, bool)`

GetAvailableForDirectDebitOk returns a tuple with the AvailableForDirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForDirectDebit

`func (o *FullBankDetails) SetAvailableForDirectDebit(v bool)`

SetAvailableForDirectDebit sets AvailableForDirectDebit field to given value.

### HasAvailableForDirectDebit

`func (o *FullBankDetails) HasAvailableForDirectDebit() bool`

HasAvailableForDirectDebit returns a boolean if a field has been set.

### GetPayWithBank

`func (o *FullBankDetails) GetPayWithBank() bool`

GetPayWithBank returns the PayWithBank field if non-nil, zero value otherwise.

### GetPayWithBankOk

`func (o *FullBankDetails) GetPayWithBankOk() (*bool, bool)`

GetPayWithBankOk returns a tuple with the PayWithBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayWithBank

`func (o *FullBankDetails) SetPayWithBank(v bool)`

SetPayWithBank sets PayWithBank field to given value.

### HasPayWithBank

`func (o *FullBankDetails) HasPayWithBank() bool`

HasPayWithBank returns a boolean if a field has been set.

### GetActive

`func (o *FullBankDetails) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FullBankDetails) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FullBankDetails) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FullBankDetails) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIsDeleted

`func (o *FullBankDetails) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FullBankDetails) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FullBankDetails) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *FullBankDetails) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FullBankDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullBankDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullBankDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FullBankDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FullBankDetails) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FullBankDetails) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FullBankDetails) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FullBankDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


