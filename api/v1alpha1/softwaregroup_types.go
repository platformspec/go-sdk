/*
Copyright 2024.

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

package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type SoftwareGroupSpecPackage struct {
	// The unique name of the package or release within this software group.
	Name string `json:"name"`

	// Defines the format and deployment method for this package (e.g., helm, docker).
	Engine string `json:"engine"`

	// Provides configuration options specific to the package engine.
	Config apiextensionsv1.JSON `json:"config,omitempty"`
}

// SoftwareGroupSpec defines the desired state of SoftwareGroup.
type SoftwareGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of individual software components included in this group. Each package specifies its name, type, and configuration details:
	Packages []SoftwareGroupSpecPackage `json:"packages"`
}

// SoftwareGroupStatus defines the observed state of SoftwareGroup.
type SoftwareGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SoftwareGroup is the Schema for the softwaregroups API.
type SoftwareGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SoftwareGroupSpec   `json:"spec,omitempty"`
	Status SoftwareGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SoftwareGroupList contains a list of SoftwareGroup.
type SoftwareGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SoftwareGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SoftwareGroup{}, &SoftwareGroupList{})
}
