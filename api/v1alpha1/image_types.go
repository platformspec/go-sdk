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

type ImageSpecSpecReference struct {
	// Unique identifier of a pre-existing image (e.g., AMI ID for machine images). This field is used when you are referencing an image that already exists in your cloud provider's registry.
	Id string `json:"id"`

	// Specifies the region where the pre-existing image is located.
	Location string `json:"location,omitempty"`
}

type ImageSpecSpecBuilder struct {
	// Specifies the type of image builder to use for creating this image
	Driver string `json:"driver"`

	// Contains configuration parameters specific to the chosen builder driver
	Config apiextensionsv1.JSON `json:"config,omitempty"`
}

type ImageSpecSpec struct {
	// Indicates whether this image is the default choice for a given provider and environment combination. This helps simplify deployments by setting up common starting points.
	// +kubebuilder:default=false
	Default bool `json:"default,omitempty"`

	// References Provider resources specifying the cloud platform on which the image can be deployed (e.g., AWS, Azure).
	ProviderRefs []corev1.ObjectReference `json:"providerRefs"`

	// References Environment resources indicating the environments where this image is intended to be used. This ensures that the correct images are targeted for different deployment stages.
	EnvironmentRefs []corev1.ObjectReference `json:"environmentRefs"`

	// Represents the version of the image using semantic versioning (e.g., v1.28.13). This helps track and manage image updates effectively.
	Version string `json:"version"`

	Builder ImageSpecSpecBuilder `json:"builder,omitempty"`

	Reference ImageSpecSpecReference `json:"reference,omitempty"`
}

// ImageSpec defines the desired state of Image.
type ImageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Indicates the type of image being defined.
	// +kubebuilder:validation:Enum=machine;container
	Category string `json:"category"`

	// Contains configuration details specific to the image.
	Spec ImageSpecSpec `json:"spec"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Image is the Schema for the images API.
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageSpec   `json:"spec,omitempty"`
	Status ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Image.
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
