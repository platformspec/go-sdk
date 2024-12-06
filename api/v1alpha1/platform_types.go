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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PlatformSpecDns defines the DNS provider and domain used by your platform.
type PlatformSpecDns struct {
	// Specifies the top-level domain name associated with this platform.
	Domain string `json:"domain"`

	// References a Provider resource defining the chosen DNS service (e.g., Route53).
	ProviderRef corev1.ObjectReference `json:"providerRef"`
}

// PlatformSpeceResources defines the resources required by the platform.
type PlatformSpecResources struct {

	// Defines different deployment environments (e.g., development, staging, production) for your platform's services and applications.
	Environments []corev1.ObjectReference `json:"environments,omitempty"`

	// References Provider resources that define the specific cloud platforms or services used within your environment (e.g., AWS, Azure, GCP).
	Providers []corev1.ObjectReference `json:"providers,omitempty"`

	// References Network resources that define the network configurations used by your platform.
	Networks []corev1.ObjectReference `json:"networks,omitempty"`

	// Defines Kubernetes clusters managed by this platform (if applicable). Each cluster can have its own configuration and deployment parameters.
	Clusters []corev1.ObjectReference `json:"clusters,omitempty"`

	// Lists virtual machines or servers managed within your platform, specifying their configurations and roles.
	Servers []corev1.ObjectReference `json:"servers,omitempty"`

	// Defines container images used for deploying applications or components within your platform.
	Images []corev1.ObjectReference `json:"images,omitempty"`

	// Groups together related software packages or dependencies required by various services or applications within your platform.
	SoftwareGroups []corev1.ObjectReference `json:"softwareGroups,omitempty"`

	// References Credential resources containing the necessary credentials for accessing and interacting with different cloud providers and services.
	Credentials []corev1.ObjectReference `json:"credentials,omitempty"`
}

// PlatformSpec defines the desired state of Platform resource.
type PlatformSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The organization responsible for managing this platform.
	Organization string `json:"organization"`

	// A brief description of the platform's purpose and functionalities.
	Description string `json:"description,omitempty"`

	// The name or team responsible for creating or maintaining this platform.
	Author string `json:"author,omitempty"`

	// A version of the platform, internal to the team defining and managing the platform.
	Version string `json:"version"`

	// An email address for contacting the platform's administrators, maintainers or support team.
	ContactEmail string `json:"contactEmail"`

	// Defines the DNS provider and domain used by your platform.
	Dns PlatformSpecDns `json:"dns"`

	// References to the various resources leveraged or managed by this platform:
	Resources PlatformSpecResources `json:"resources"`
}

// PlatformStatus defines the observed state of Platform.
type PlatformStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Platform is the Schema for the platforms API.
type Platform struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlatformSpec   `json:"spec,omitempty"`
	Status PlatformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlatformList contains a list of Platform.
type PlatformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Platform `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Platform{}, &PlatformList{})
}
