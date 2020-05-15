// Copyright © 2020 VMware
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	contourv1 "github.com/projectcontour/contour/apis/projectcontour/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ExtensionProtocolVersion is the version of the GRPC protocol used
// to request support services.
type ExtensionProtocolVersion string

// SupportProtocolVersion2 requests the "v2" support protocol version.
const SupportProtocolVersion2 ExtensionProtocolVersion = "v2"

// ExtensionServiceSpec defines the desired state of ExtensionService
type ExtensionServiceSpec struct {
	// +required
	Service contourv1.Service `json:"service"`

	// The load balancing policy for sending requests to the service.
	// +optional
	LoadBalancerPolicy *contourv1.LoadBalancerPolicy `json:"loadBalancerPolicy,omitempty"`

	// This field sets the version of the GRPC protocol that Envoy uses to
	// send requests to the support service. Since Contour always uses the
	// v2 Envoy API, this is currently fixed at "v2". However, other
	// protocol options will be available in future.
	//
	// +kubebuilder:validation:Enum=v2
	ProtocolVersion ExtensionProtocolVersion `json:"protocolVersion,omitempty"`
}

// ExtensionServiceStatus defines the observed state of ExtensionService
type ExtensionServiceStatus struct {
	Conditions []Condition `json:"conditions"`
}

// Condition ...
type Condition struct {
	// Type of condition in CamelCase or in foo.example.com/CamelCase.
	// Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be
	// useful (see .node.status.conditions), the ability to deconflict is important.
	// +required
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	// +required
	Status metav1.ConditionStatus `json:"status"`
	// If set, this represents the .metadata.generation that the condition was set based upon.
	// For instance, if .metadata.generation is currently 12, but the .status.condition[x].observedGeneration is 9, the condition is out of date
	// with respect to the current state of the instance.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// Last time the condition transitioned from one status to another.
	// This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.
	// +required
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// The reason for the condition's last transition in CamelCase.
	// The specific API may choose whether or not this field is considered a guaranteed API.
	// This field may not be empty.
	// +required
	Reason string `json:"reason"`
	// A human readable message indicating details about the transition.
	// This field may be empty.
	// +required
	Message string `json:"message"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=supportservice;supportservices

// ExtensionService is the Schema for the externalclusters API
type ExtensionService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExtensionServiceSpec   `json:"spec,omitempty"`
	Status ExtensionServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExtensionServiceList contains a list of ExtensionService
type ExtensionServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExtensionService `json:"items"`
}
