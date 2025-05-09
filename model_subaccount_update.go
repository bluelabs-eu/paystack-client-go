/*
Paystack

The OpenAPI specification of the Paystack API that merchants and developers can harness to build financial solutions in Africa.

API version: 1.0.0
Contact: techsupport@paystack.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SubaccountUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountUpdate{}

// SubaccountUpdate struct for SubaccountUpdate
type SubaccountUpdate struct {
	// Name of business for subaccount
	BusinessName *string `json:"business_name,omitempty"`
	// Bank code for the bank. You can get the list of Bank Codes by calling the List Banks endpoint.
	SettlementBank *string `json:"settlement_bank,omitempty"`
	// Bank account number
	AccountNumber *string `json:"account_number,omitempty"`
	// Activate or deactivate a subaccount
	Active *bool `json:"active,omitempty"`
	// Customer's phone number
	PercentageCharge *float32 `json:"percentage_charge,omitempty"`
	// A description for this subaccount
	Description *string `json:"description,omitempty"`
	// A contact email for the subaccount
	PrimaryContactEmail *string `json:"primary_contact_email,omitempty"`
	// The name of the contact person for this subaccount
	PrimaryContactName *string `json:"primary_contact_name,omitempty"`
	// A phone number to call for this subaccount
	PrimaryContactPhone *string `json:"primary_contact_phone,omitempty"`
	// Stringified JSON object of custom data
	Metadata *string `json:"metadata,omitempty"`
}

// NewSubaccountUpdate instantiates a new SubaccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountUpdate() *SubaccountUpdate {
	this := SubaccountUpdate{}
	return &this
}

// NewSubaccountUpdateWithDefaults instantiates a new SubaccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountUpdateWithDefaults() *SubaccountUpdate {
	this := SubaccountUpdate{}
	return &this
}

// GetBusinessName returns the BusinessName field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetBusinessName() string {
	if o == nil || IsNil(o.BusinessName) {
		var ret string
		return ret
	}
	return *o.BusinessName
}

// GetBusinessNameOk returns a tuple with the BusinessName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetBusinessNameOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessName) {
		return nil, false
	}
	return o.BusinessName, true
}

// HasBusinessName returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasBusinessName() bool {
	if o != nil && !IsNil(o.BusinessName) {
		return true
	}

	return false
}

// SetBusinessName gets a reference to the given string and assigns it to the BusinessName field.
func (o *SubaccountUpdate) SetBusinessName(v string) {
	o.BusinessName = &v
}

// GetSettlementBank returns the SettlementBank field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetSettlementBank() string {
	if o == nil || IsNil(o.SettlementBank) {
		var ret string
		return ret
	}
	return *o.SettlementBank
}

// GetSettlementBankOk returns a tuple with the SettlementBank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetSettlementBankOk() (*string, bool) {
	if o == nil || IsNil(o.SettlementBank) {
		return nil, false
	}
	return o.SettlementBank, true
}

// HasSettlementBank returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasSettlementBank() bool {
	if o != nil && !IsNil(o.SettlementBank) {
		return true
	}

	return false
}

// SetSettlementBank gets a reference to the given string and assigns it to the SettlementBank field.
func (o *SubaccountUpdate) SetSettlementBank(v string) {
	o.SettlementBank = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *SubaccountUpdate) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SubaccountUpdate) SetActive(v bool) {
	o.Active = &v
}

// GetPercentageCharge returns the PercentageCharge field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetPercentageCharge() float32 {
	if o == nil || IsNil(o.PercentageCharge) {
		var ret float32
		return ret
	}
	return *o.PercentageCharge
}

// GetPercentageChargeOk returns a tuple with the PercentageCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetPercentageChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentageCharge) {
		return nil, false
	}
	return o.PercentageCharge, true
}

// HasPercentageCharge returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasPercentageCharge() bool {
	if o != nil && !IsNil(o.PercentageCharge) {
		return true
	}

	return false
}

// SetPercentageCharge gets a reference to the given float32 and assigns it to the PercentageCharge field.
func (o *SubaccountUpdate) SetPercentageCharge(v float32) {
	o.PercentageCharge = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubaccountUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetPrimaryContactEmail returns the PrimaryContactEmail field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetPrimaryContactEmail() string {
	if o == nil || IsNil(o.PrimaryContactEmail) {
		var ret string
		return ret
	}
	return *o.PrimaryContactEmail
}

// GetPrimaryContactEmailOk returns a tuple with the PrimaryContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetPrimaryContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactEmail) {
		return nil, false
	}
	return o.PrimaryContactEmail, true
}

// HasPrimaryContactEmail returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasPrimaryContactEmail() bool {
	if o != nil && !IsNil(o.PrimaryContactEmail) {
		return true
	}

	return false
}

// SetPrimaryContactEmail gets a reference to the given string and assigns it to the PrimaryContactEmail field.
func (o *SubaccountUpdate) SetPrimaryContactEmail(v string) {
	o.PrimaryContactEmail = &v
}

// GetPrimaryContactName returns the PrimaryContactName field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetPrimaryContactName() string {
	if o == nil || IsNil(o.PrimaryContactName) {
		var ret string
		return ret
	}
	return *o.PrimaryContactName
}

// GetPrimaryContactNameOk returns a tuple with the PrimaryContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetPrimaryContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactName) {
		return nil, false
	}
	return o.PrimaryContactName, true
}

// HasPrimaryContactName returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasPrimaryContactName() bool {
	if o != nil && !IsNil(o.PrimaryContactName) {
		return true
	}

	return false
}

// SetPrimaryContactName gets a reference to the given string and assigns it to the PrimaryContactName field.
func (o *SubaccountUpdate) SetPrimaryContactName(v string) {
	o.PrimaryContactName = &v
}

// GetPrimaryContactPhone returns the PrimaryContactPhone field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetPrimaryContactPhone() string {
	if o == nil || IsNil(o.PrimaryContactPhone) {
		var ret string
		return ret
	}
	return *o.PrimaryContactPhone
}

// GetPrimaryContactPhoneOk returns a tuple with the PrimaryContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetPrimaryContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactPhone) {
		return nil, false
	}
	return o.PrimaryContactPhone, true
}

// HasPrimaryContactPhone returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasPrimaryContactPhone() bool {
	if o != nil && !IsNil(o.PrimaryContactPhone) {
		return true
	}

	return false
}

// SetPrimaryContactPhone gets a reference to the given string and assigns it to the PrimaryContactPhone field.
func (o *SubaccountUpdate) SetPrimaryContactPhone(v string) {
	o.PrimaryContactPhone = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SubaccountUpdate) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountUpdate) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SubaccountUpdate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *SubaccountUpdate) SetMetadata(v string) {
	o.Metadata = &v
}

func (o SubaccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessName) {
		toSerialize["business_name"] = o.BusinessName
	}
	if !IsNil(o.SettlementBank) {
		toSerialize["settlement_bank"] = o.SettlementBank
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.PercentageCharge) {
		toSerialize["percentage_charge"] = o.PercentageCharge
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PrimaryContactEmail) {
		toSerialize["primary_contact_email"] = o.PrimaryContactEmail
	}
	if !IsNil(o.PrimaryContactName) {
		toSerialize["primary_contact_name"] = o.PrimaryContactName
	}
	if !IsNil(o.PrimaryContactPhone) {
		toSerialize["primary_contact_phone"] = o.PrimaryContactPhone
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableSubaccountUpdate struct {
	value *SubaccountUpdate
	isSet bool
}

func (v NullableSubaccountUpdate) Get() *SubaccountUpdate {
	return v.value
}

func (v *NullableSubaccountUpdate) Set(val *SubaccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountUpdate(val *SubaccountUpdate) *NullableSubaccountUpdate {
	return &NullableSubaccountUpdate{value: val, isSet: true}
}

func (v NullableSubaccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


