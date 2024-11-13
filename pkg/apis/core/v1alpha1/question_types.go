package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// QuestionSpec defines the desired state of Question
type QuestionSpec struct {
	// Question is the question to be asked
	// +kubebuilder:validation:Required
	Question string `json:"question"`
}

// QuestionStatus defines the observed state of Question
type QuestionStatus struct {
	// Phase indicates the current phase of the Question (e.g., Pending, Answered)
	Phase string `json:"phase,omitempty"`

	// Answer provides the answer for the question
	Answer string `json:"answer,omitempty"`

	// Message provides additional information about the status
	Message string `json:"message,omitempty"`

	// LastUpdated records the time when the status was last updated
	// +optional
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Question is the Schema for the questions API
type Question struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuestionSpec   `json:"spec,omitempty"`
	Status QuestionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// QuestionList contains a list of Question
type QuestionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Question `json:"items"`
}

func init() { //nolint:gochecknoinits
	SchemeBuilder.Register(&Question{}, &QuestionList{})
}
