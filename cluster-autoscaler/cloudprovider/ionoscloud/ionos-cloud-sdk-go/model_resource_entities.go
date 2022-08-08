/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ResourceEntities struct for ResourceEntities
type ResourceEntities struct {
	Groups *ResourceGroups `json:"groups,omitempty"`
}

// NewResourceEntities instantiates a new ResourceEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceEntities() *ResourceEntities {
	this := ResourceEntities{}

	return &this
}

// NewResourceEntitiesWithDefaults instantiates a new ResourceEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceEntitiesWithDefaults() *ResourceEntities {
	this := ResourceEntities{}
	return &this
}

// GetGroups returns the Groups field value
// If the value is explicit nil, the zero value for ResourceGroups will be returned
func (o *ResourceEntities) GetGroups() *ResourceGroups {
	if o == nil {
		return nil
	}

	return o.Groups

}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceEntities) GetGroupsOk() (*ResourceGroups, bool) {
	if o == nil {
		return nil, false
	}

	return o.Groups, true
}

// SetGroups sets field value
func (o *ResourceEntities) SetGroups(v ResourceGroups) {

	o.Groups = &v

}

// HasGroups returns a boolean if a field has been set.
func (o *ResourceEntities) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

func (o ResourceEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableResourceEntities struct {
	value *ResourceEntities
	isSet bool
}

func (v NullableResourceEntities) Get() *ResourceEntities {
	return v.value
}

func (v *NullableResourceEntities) Set(val *ResourceEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceEntities(val *ResourceEntities) *NullableResourceEntities {
	return &NullableResourceEntities{value: val, isSet: true}
}

func (v NullableResourceEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
