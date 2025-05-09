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
	"bytes"
	"fmt"
)

// checks if the TransactionPartialDebit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionPartialDebit{}

// TransactionPartialDebit struct for TransactionPartialDebit
type TransactionPartialDebit struct {
	// Customer's email address
	Email string `json:"email"`
	// Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	Amount int32 `json:"amount"`
	// Valid authorization code to charge
	AuthorizationCode string `json:"authorization_code"`
	// The transaction currency
	Currency string `json:"currency"`
	// Unique transaction reference. Only -, ., = and alphanumeric characters allowed.
	Reference *string `json:"reference,omitempty"`
	// Minimum amount to charge
	AtLeast *string `json:"at_least,omitempty"`
}

type _TransactionPartialDebit TransactionPartialDebit

// NewTransactionPartialDebit instantiates a new TransactionPartialDebit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionPartialDebit(email string, amount int32, authorizationCode string, currency string) *TransactionPartialDebit {
	this := TransactionPartialDebit{}
	this.Email = email
	this.Amount = amount
	this.AuthorizationCode = authorizationCode
	this.Currency = currency
	return &this
}

// NewTransactionPartialDebitWithDefaults instantiates a new TransactionPartialDebit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionPartialDebitWithDefaults() *TransactionPartialDebit {
	this := TransactionPartialDebit{}
	return &this
}

// GetEmail returns the Email field value
func (o *TransactionPartialDebit) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TransactionPartialDebit) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TransactionPartialDebit) SetEmail(v string) {
	o.Email = v
}

// GetAmount returns the Amount field value
func (o *TransactionPartialDebit) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionPartialDebit) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionPartialDebit) SetAmount(v int32) {
	o.Amount = v
}

// GetAuthorizationCode returns the AuthorizationCode field value
func (o *TransactionPartialDebit) GetAuthorizationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value
// and a boolean to check if the value has been set.
func (o *TransactionPartialDebit) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationCode, true
}

// SetAuthorizationCode sets field value
func (o *TransactionPartialDebit) SetAuthorizationCode(v string) {
	o.AuthorizationCode = v
}

// GetCurrency returns the Currency field value
func (o *TransactionPartialDebit) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionPartialDebit) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionPartialDebit) SetCurrency(v string) {
	o.Currency = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TransactionPartialDebit) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartialDebit) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TransactionPartialDebit) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TransactionPartialDebit) SetReference(v string) {
	o.Reference = &v
}

// GetAtLeast returns the AtLeast field value if set, zero value otherwise.
func (o *TransactionPartialDebit) GetAtLeast() string {
	if o == nil || IsNil(o.AtLeast) {
		var ret string
		return ret
	}
	return *o.AtLeast
}

// GetAtLeastOk returns a tuple with the AtLeast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPartialDebit) GetAtLeastOk() (*string, bool) {
	if o == nil || IsNil(o.AtLeast) {
		return nil, false
	}
	return o.AtLeast, true
}

// HasAtLeast returns a boolean if a field has been set.
func (o *TransactionPartialDebit) HasAtLeast() bool {
	if o != nil && !IsNil(o.AtLeast) {
		return true
	}

	return false
}

// SetAtLeast gets a reference to the given string and assigns it to the AtLeast field.
func (o *TransactionPartialDebit) SetAtLeast(v string) {
	o.AtLeast = &v
}

func (o TransactionPartialDebit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionPartialDebit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["amount"] = o.Amount
	toSerialize["authorization_code"] = o.AuthorizationCode
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.AtLeast) {
		toSerialize["at_least"] = o.AtLeast
	}
	return toSerialize, nil
}

func (o *TransactionPartialDebit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"amount",
		"authorization_code",
		"currency",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransactionPartialDebit := _TransactionPartialDebit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionPartialDebit)

	if err != nil {
		return err
	}

	*o = TransactionPartialDebit(varTransactionPartialDebit)

	return err
}

type NullableTransactionPartialDebit struct {
	value *TransactionPartialDebit
	isSet bool
}

func (v NullableTransactionPartialDebit) Get() *TransactionPartialDebit {
	return v.value
}

func (v *NullableTransactionPartialDebit) Set(val *TransactionPartialDebit) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionPartialDebit) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionPartialDebit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionPartialDebit(val *TransactionPartialDebit) *NullableTransactionPartialDebit {
	return &NullableTransactionPartialDebit{value: val, isSet: true}
}

func (v NullableTransactionPartialDebit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionPartialDebit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


