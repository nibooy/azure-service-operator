/*

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

// SqlServerSpec defines the desired state of SqlServer
type SqlServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location                string `json:"location"`
	ResourceGroup           string `json:"resourcegroup,omitempty"`
	AdminUser               string `json:"adminuser,omitempty"`
	AllowAzureServiceAccess bool   `json:"allowazureserviceaccess,omitempty"`
}

// SqlServerStatus defines the observed state of SqlServer
type SqlServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool   `json:"provisioning,omitempty"`
	Provisioned  bool   `json:"provisioned,omitempty"`
	State        string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// SqlServer is the Schema for the sqlservers API
type SqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqlServerSpec   `json:"spec,omitempty"`
	Status SqlServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlServerList contains a list of SqlServer
type SqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SqlServer{}, &SqlServerList{})
}

func (s *SqlServer) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}
