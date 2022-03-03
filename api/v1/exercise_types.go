/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExerciseSpec defines the desired state of Exercise
type ExerciseSpec struct {
	//+kubebuilder:validation:Minimum=0
	// Size is the size of the memcached deployment
	Size           int32  `json:"size"`
	Image          string `json:"image"`
	ServiceAccount string `json:"serviceAccount"`
}

// ExerciseStatus defines the observed state of Exercise
type ExerciseStatus struct {
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Exercise is the Schema for the exercises API
//+kubebuilder:subresource:status
type Exercise struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExerciseSpec   `json:"spec,omitempty"`
	Status ExerciseStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExerciseList contains a list of Exercise
type ExerciseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Exercise `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Exercise{}, &ExerciseList{})
}
