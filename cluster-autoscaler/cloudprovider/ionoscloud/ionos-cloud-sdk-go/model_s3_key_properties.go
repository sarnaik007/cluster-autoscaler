/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// S3KeyProperties struct for S3KeyProperties
type S3KeyProperties struct {
	// secret of the s3 key
	SecretKey *string `json:"secretKey,omitempty"`
	// denotes if the s3 key is active or not
	Active *bool `json:"active,omitempty"`
}



// GetSecretKey returns the SecretKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *S3KeyProperties) GetSecretKey() *string {
	if o == nil {
		return nil
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3KeyProperties) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// SetSecretKey sets field value
func (o *S3KeyProperties) SetSecretKey(v string) {
	o.SecretKey = &v
}

// HasSecretKey returns a boolean if a field has been set.
func (o *S3KeyProperties) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}



// GetActive returns the Active field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *S3KeyProperties) GetActive() *bool {
	if o == nil {
		return nil
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3KeyProperties) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active, true
}

// SetActive sets field value
func (o *S3KeyProperties) SetActive(v bool) {
	o.Active = &v
}

// HasActive returns a boolean if a field has been set.
func (o *S3KeyProperties) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}


func (o S3KeyProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	

	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	
	return json.Marshal(toSerialize)
}

type NullableS3KeyProperties struct {
	value *S3KeyProperties
	isSet bool
}

func (v NullableS3KeyProperties) Get() *S3KeyProperties {
	return v.value
}

func (v *NullableS3KeyProperties) Set(val *S3KeyProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableS3KeyProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableS3KeyProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3KeyProperties(val *S3KeyProperties) *NullableS3KeyProperties {
	return &NullableS3KeyProperties{value: val, isSet: true}
}

func (v NullableS3KeyProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3KeyProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


