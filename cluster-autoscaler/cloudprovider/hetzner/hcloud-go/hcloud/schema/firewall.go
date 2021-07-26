/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package schema

import "time"

// Firewall defines the schema of a Firewall.
type Firewall struct {
	ID        int                `json:"id"`
	Name      string             `json:"name"`
	Labels    map[string]string  `json:"labels"`
	Created   time.Time          `json:"created"`
	Rules     []FirewallRule     `json:"rules"`
	AppliedTo []FirewallResource `json:"applied_to"`
}

// FirewallRule defines the schema of a Firewall rule.
type FirewallRule struct {
	Direction      string   `json:"direction"`
	SourceIPs      []string `json:"source_ips,omitempty"`
	DestinationIPs []string `json:"destination_ips,omitempty"`
	Protocol       string   `json:"protocol"`
	Port           *string  `json:"port,omitempty"`
}

// FirewallListResponse defines the schema of the response when listing Firewalls.
type FirewallListResponse struct {
	Firewalls []Firewall `json:"firewalls"`
}

// FirewallGetResponse defines the schema of the response when retrieving a single Firewall.
type FirewallGetResponse struct {
	Firewall Firewall `json:"firewall"`
}

// FirewallCreateRequest defines the schema of the request to create a Firewall.
type FirewallCreateRequest struct {
	Name    string             `json:"name"`
	Labels  *map[string]string `json:"labels,omitempty"`
	Rules   []FirewallRule     `json:"rules,omitempty"`
	ApplyTo []FirewallResource `json:"apply_to,omitempty"`
}

// FirewallResource defines the schema of a resource to apply the new Firewall on.
type FirewallResource struct {
	Type          string                         `json:"type"`
	Server        *FirewallResourceServer        `json:"server,omitempty"`
	LabelSelector *FirewallResourceLabelSelector `json:"label_selector,omitempty"`
}

// FirewallResourceLabelSelector defines the schema of a LabelSelector to apply a Firewall on.
type FirewallResourceLabelSelector struct {
	Selector string `json:"selector"`
}

// FirewallResourceServer defines the schema of a Server to apply a Firewall on.
type FirewallResourceServer struct {
	ID int `json:"id"`
}

// FirewallCreateResponse defines the schema of the response when creating a Firewall.
type FirewallCreateResponse struct {
	Firewall Firewall `json:"firewall"`
	Actions  []Action `json:"actions"`
}

// FirewallUpdateRequest defines the schema of the request to update a Firewall.
type FirewallUpdateRequest struct {
	Name   *string            `json:"name,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
}

// FirewallUpdateResponse defines the schema of the response when updating a Firewall.
type FirewallUpdateResponse struct {
	Firewall Firewall `json:"firewall"`
}

// FirewallActionSetRulesRequest defines the schema of the request when setting Firewall rules.
type FirewallActionSetRulesRequest struct {
	Rules []FirewallRule `json:"rules"`
}

// FirewallActionSetRulesResponse defines the schema of the response when setting Firewall rules.
type FirewallActionSetRulesResponse struct {
	Actions []Action `json:"actions"`
}

// FirewallActionApplyToResourcesRequest defines the schema of the request when applying a Firewall on resources.
type FirewallActionApplyToResourcesRequest struct {
	ApplyTo []FirewallResource `json:"apply_to"`
}

// FirewallActionApplyToResourcesResponse defines the schema of the response when applying a Firewall on resources.
type FirewallActionApplyToResourcesResponse struct {
	Actions []Action `json:"actions"`
}

// FirewallActionRemoveFromResourcesRequest defines the schema of the request when removing a Firewall from resources.
type FirewallActionRemoveFromResourcesRequest struct {
	RemoveFrom []FirewallResource `json:"remove_from"`
}

// FirewallActionRemoveFromResourcesResponse defines the schema of the response when removing a Firewall from resources.
type FirewallActionRemoveFromResourcesResponse struct {
	Actions []Action `json:"actions"`
}
