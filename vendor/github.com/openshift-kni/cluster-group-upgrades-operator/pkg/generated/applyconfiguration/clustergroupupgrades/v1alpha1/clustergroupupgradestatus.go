/*
Copyright 2021.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterGroupUpgradeStatusApplyConfiguration represents an declarative configuration of the ClusterGroupUpgradeStatus type for use
// with apply.
type ClusterGroupUpgradeStatusApplyConfiguration struct {
	PlacementBindings                     []string                                    `json:"placementBindings,omitempty"`
	PlacementRules                        []string                                    `json:"placementRules,omitempty"`
	CopiedPolicies                        []string                                    `json:"copiedPolicies,omitempty"`
	Conditions                            []v1.Condition                              `json:"conditions,omitempty"`
	RemediationPlan                       [][]string                                  `json:"remediationPlan,omitempty"`
	ManagedPoliciesNs                     map[string]string                           `json:"managedPoliciesNs,omitempty"`
	SafeResourceNames                     map[string]string                           `json:"safeResourceNames,omitempty"`
	ManagedPoliciesForUpgrade             []ManagedPolicyForUpgradeApplyConfiguration `json:"managedPoliciesForUpgrade,omitempty"`
	ManagedPoliciesCompliantBeforeUpgrade []string                                    `json:"managedPoliciesCompliantBeforeUpgrade,omitempty"`
	ManagedPoliciesContent                map[string]string                           `json:"managedPoliciesContent,omitempty"`
	Clusters                              []ClusterStateApplyConfiguration            `json:"clusters,omitempty"`
	Status                                *UpgradeStatusApplyConfiguration            `json:"status,omitempty"`
	Precaching                            *PrecachingStatusApplyConfiguration         `json:"precaching,omitempty"`
	Backup                                *BackupStatusApplyConfiguration             `json:"backup,omitempty"`
	ComputedMaxConcurrency                *int                                        `json:"computedMaxConcurrency,omitempty"`
}

// ClusterGroupUpgradeStatusApplyConfiguration constructs an declarative configuration of the ClusterGroupUpgradeStatus type for use with
// apply.
func ClusterGroupUpgradeStatus() *ClusterGroupUpgradeStatusApplyConfiguration {
	return &ClusterGroupUpgradeStatusApplyConfiguration{}
}

// WithPlacementBindings adds the given value to the PlacementBindings field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PlacementBindings field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithPlacementBindings(values ...string) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		b.PlacementBindings = append(b.PlacementBindings, values[i])
	}
	return b
}

// WithPlacementRules adds the given value to the PlacementRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PlacementRules field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithPlacementRules(values ...string) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		b.PlacementRules = append(b.PlacementRules, values[i])
	}
	return b
}

// WithCopiedPolicies adds the given value to the CopiedPolicies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CopiedPolicies field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithCopiedPolicies(values ...string) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		b.CopiedPolicies = append(b.CopiedPolicies, values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithConditions(values ...v1.Condition) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}

// WithRemediationPlan adds the given value to the RemediationPlan field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RemediationPlan field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithRemediationPlan(values ...[]string) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		b.RemediationPlan = append(b.RemediationPlan, values[i])
	}
	return b
}

// WithManagedPoliciesNs puts the entries into the ManagedPoliciesNs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ManagedPoliciesNs field,
// overwriting an existing map entries in ManagedPoliciesNs field with the same key.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithManagedPoliciesNs(entries map[string]string) *ClusterGroupUpgradeStatusApplyConfiguration {
	if b.ManagedPoliciesNs == nil && len(entries) > 0 {
		b.ManagedPoliciesNs = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ManagedPoliciesNs[k] = v
	}
	return b
}

// WithSafeResourceNames puts the entries into the SafeResourceNames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the SafeResourceNames field,
// overwriting an existing map entries in SafeResourceNames field with the same key.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithSafeResourceNames(entries map[string]string) *ClusterGroupUpgradeStatusApplyConfiguration {
	if b.SafeResourceNames == nil && len(entries) > 0 {
		b.SafeResourceNames = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.SafeResourceNames[k] = v
	}
	return b
}

// WithManagedPoliciesForUpgrade adds the given value to the ManagedPoliciesForUpgrade field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ManagedPoliciesForUpgrade field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithManagedPoliciesForUpgrade(values ...*ManagedPolicyForUpgradeApplyConfiguration) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithManagedPoliciesForUpgrade")
		}
		b.ManagedPoliciesForUpgrade = append(b.ManagedPoliciesForUpgrade, *values[i])
	}
	return b
}

// WithManagedPoliciesCompliantBeforeUpgrade adds the given value to the ManagedPoliciesCompliantBeforeUpgrade field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ManagedPoliciesCompliantBeforeUpgrade field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithManagedPoliciesCompliantBeforeUpgrade(values ...string) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		b.ManagedPoliciesCompliantBeforeUpgrade = append(b.ManagedPoliciesCompliantBeforeUpgrade, values[i])
	}
	return b
}

// WithManagedPoliciesContent puts the entries into the ManagedPoliciesContent field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ManagedPoliciesContent field,
// overwriting an existing map entries in ManagedPoliciesContent field with the same key.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithManagedPoliciesContent(entries map[string]string) *ClusterGroupUpgradeStatusApplyConfiguration {
	if b.ManagedPoliciesContent == nil && len(entries) > 0 {
		b.ManagedPoliciesContent = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ManagedPoliciesContent[k] = v
	}
	return b
}

// WithClusters adds the given value to the Clusters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Clusters field.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithClusters(values ...*ClusterStateApplyConfiguration) *ClusterGroupUpgradeStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithClusters")
		}
		b.Clusters = append(b.Clusters, *values[i])
	}
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithStatus(value *UpgradeStatusApplyConfiguration) *ClusterGroupUpgradeStatusApplyConfiguration {
	b.Status = value
	return b
}

// WithPrecaching sets the Precaching field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Precaching field is set to the value of the last call.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithPrecaching(value *PrecachingStatusApplyConfiguration) *ClusterGroupUpgradeStatusApplyConfiguration {
	b.Precaching = value
	return b
}

// WithBackup sets the Backup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Backup field is set to the value of the last call.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithBackup(value *BackupStatusApplyConfiguration) *ClusterGroupUpgradeStatusApplyConfiguration {
	b.Backup = value
	return b
}

// WithComputedMaxConcurrency sets the ComputedMaxConcurrency field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ComputedMaxConcurrency field is set to the value of the last call.
func (b *ClusterGroupUpgradeStatusApplyConfiguration) WithComputedMaxConcurrency(value int) *ClusterGroupUpgradeStatusApplyConfiguration {
	b.ComputedMaxConcurrency = &value
	return b
}