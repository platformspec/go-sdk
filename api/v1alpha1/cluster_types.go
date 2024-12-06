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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterSpec defines the desired state of Cluster.
type ClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Specifies the cloud providers responsible for provisioning and managing the Kubernetes infrastructure. Each reference points to a Provider resource, defining details like account credentials and region.
	ProviderRefs []corev1.ObjectReference `json:"providerRefs"`

	// References an Environment resource, associating the cluster with its corresponding deployment stage (development, staging, production). This ensures that the cluster is configured with the appropriate settings for its intended purpose.
	EnvironmentRef corev1.ObjectReference `json:"environmentRef"`

	// Identifies the network or VPC to which this server will be connected.
	NetworkRefs []corev1.ObjectReference `json:"networkRefs,omitempty"`

	//  Links to a SoftwareGroup resource defining the base Kubernetes components and additional tools or software packages that will be included in the cluster.
	SoftwareGroupRefs []corev1.ObjectReference `json:"softwareGroupRefs,omitempty"`

	// Specifies the version to be deployed within cluster's engine, enabling you to manage different versions for various environments.
	Version string `json:"version"`

	// Indicates the geographical region where the cluster resources will be deployed.
	Region string `json:"region"`

	// Configuration parameters for the Provider provisioning or managing the Cluster.
	Config apiextensionsv1.JSON `json:"config,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Cluster is the Schema for the clusters API.
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Cluster.
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
