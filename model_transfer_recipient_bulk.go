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

// checks if the TransferRecipientBulk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferRecipientBulk{}

// TransferRecipientBulk struct for TransferRecipientBulk
type TransferRecipientBulk struct {
	// A list of transfer recipient object. Each object should contain type, name, and bank_code.  Any Create Transfer Recipient param can also be passed.
	Batch []TransferRecipientCreate `json:"batch"`
}

type _TransferRecipientBulk TransferRecipientBulk

// NewTransferRecipientBulk instantiates a new TransferRecipientBulk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRecipientBulk(batch []TransferRecipientCreate) *TransferRecipientBulk {
	this := TransferRecipientBulk{}
	this.Batch = batch
	return &this
}

// NewTransferRecipientBulkWithDefaults instantiates a new TransferRecipientBulk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRecipientBulkWithDefaults() *TransferRecipientBulk {
	this := TransferRecipientBulk{}
	return &this
}

// GetBatch returns the Batch field value
func (o *TransferRecipientBulk) GetBatch() []TransferRecipientCreate {
	if o == nil {
		var ret []TransferRecipientCreate
		return ret
	}

	return o.Batch
}

// GetBatchOk returns a tuple with the Batch field value
// and a boolean to check if the value has been set.
func (o *TransferRecipientBulk) GetBatchOk() ([]TransferRecipientCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Batch, true
}

// SetBatch sets field value
func (o *TransferRecipientBulk) SetBatch(v []TransferRecipientCreate) {
	o.Batch = v
}

func (o TransferRecipientBulk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRecipientBulk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["batch"] = o.Batch
	return toSerialize, nil
}

func (o *TransferRecipientBulk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"batch",
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

	varTransferRecipientBulk := _TransferRecipientBulk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferRecipientBulk)

	if err != nil {
		return err
	}

	*o = TransferRecipientBulk(varTransferRecipientBulk)

	return err
}

type NullableTransferRecipientBulk struct {
	value *TransferRecipientBulk
	isSet bool
}

func (v NullableTransferRecipientBulk) Get() *TransferRecipientBulk {
	return v.value
}

func (v *NullableTransferRecipientBulk) Set(val *TransferRecipientBulk) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRecipientBulk) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRecipientBulk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRecipientBulk(val *TransferRecipientBulk) *NullableTransferRecipientBulk {
	return &NullableTransferRecipientBulk{value: val, isSet: true}
}

func (v NullableTransferRecipientBulk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRecipientBulk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


