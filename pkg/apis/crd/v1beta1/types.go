package v1beta1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Crd struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CrdSpec   `json:"spec,omitempty"`
	Status CrdStatus `json:"status,omitempty"`
}

type CrdSpec struct {
}

type CrdStatus struct {
	Status     string         `json:"status,omitempty"`
	Conditions []CrdCondition `json:"conditions,omitempty"`
}

type CrdCondition struct {
	LastProbeTime string `json:"lastProbeTime,omitempty"`
	Status        string `json:"status,omitempty"`
	Reason        string `json:"reason,omitempty"`
}

// +genclient:nonNamespaced
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CrdList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`

	Items []Crd `json:"items"`
}
