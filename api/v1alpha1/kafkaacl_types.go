// Copyright (c) 2020 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KafkaACLSpec defines the desired state of KafkaACL
type KafkaACLSpec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Format="^[a-zA-Z0-9_-]*$"
	// x-kubernetes-immutable: true
	// Project to link the Kafka ACL to
	Project string `json:"project"`

	// +kubebuilder:validation:MaxLength=63
	// x-kubernetes-immutable: true
	// Service to link the Kafka ACL to
	ServiceName string `json:"service_name"`

	// +kubebuilder:validation:Enum=admin;read;readwrite;write
	// x-kubernetes-immutable: true
	// Kafka permission to grant (admin, read, readwrite, write)
	Permission string `json:"permission"`

	// x-kubernetes-immutable: true
	// Topic name pattern for the ACL entry
	Topic string `json:"topic"`

	// x-kubernetes-immutable: true
	// Username pattern for the ACL entry
	Username string `json:"username"`
}

// KafkaACLStatus defines the observed state of KafkaACL
type KafkaACLStatus struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Format="^[a-zA-Z0-9_-]*$"
	// Project to link the Kafka ACL to
	Project string `json:"project"`

	// +kubebuilder:validation:MaxLength=63
	// Service to link the Kafka ACL to
	ServiceName string `json:"service_name"`

	// +kubebuilder:validation:Enum=admin;read;readwrite;write
	// Kafka permission to grant (admin, read, readwrite, write)
	Permission string `json:"permission"`

	// Topic name pattern for the ACL entry
	Topic string `json:"topic"`

	// Username pattern for the ACL entry
	Username string `json:"username"`

	// Kafka ACL ID
	Id string `json:"id"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KafkaACL is the Schema for the kafkaacls API
type KafkaACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaACLSpec   `json:"spec,omitempty"`
	Status KafkaACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaACLList contains a list of KafkaACL
type KafkaACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaACL `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KafkaACL{}, &KafkaACLList{})
}
