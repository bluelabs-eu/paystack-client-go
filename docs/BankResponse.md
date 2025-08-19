# BankResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]FullBankDetails**](FullBankDetails.md) |  | [optional] 

## Methods

### NewBankResponse

`func NewBankResponse() *BankResponse`

NewBankResponse instantiates a new BankResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankResponseWithDefaults

`func NewBankResponseWithDefaults() *BankResponse`

NewBankResponseWithDefaults instantiates a new BankResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BankResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BankResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *BankResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BankResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BankResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BankResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *BankResponse) GetData() []FullBankDetails`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BankResponse) GetDataOk() (*[]FullBankDetails, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BankResponse) SetData(v []FullBankDetails)`

SetData sets Data field to given value.

### HasData

`func (o *BankResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


