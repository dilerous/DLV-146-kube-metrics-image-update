package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var defaultFalse = false
var defaultTrue = true

type CnvrgAppSpec struct {
	ClusterDomain string             `json:"clusterDomain,omitempty"`
	ImageHub      string             `json:"imageHub,omitempty"`
	Labels        map[string]string  `json:"labels,omitempty"`
	Annotations   map[string]string  `json:"annotations,omitempty"`
	ControlPlane  ControlPlane       `json:"controlPlane,omitempty"`
	Registry      Registry           `json:"registry,omitempty"`
	Dbs           AppDbs             `json:"dbs,omitempty"`
	Networking    CnvrgAppNetworking `json:"networking,omitempty"`
	Logging       CnvrgAppLogging    `json:"logging,omitempty"`
	Monitoring    CnvrgAppMonitoring `json:"monitoring,omitempty"`
	SSO           SSO                `json:"sso,omitempty"`
	Tenancy       Tenancy            `json:"tenancy,omitempty"`
	Proxy         Proxy              `json:"proxy,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.controlPlane.image`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Message",type=string,JSONPath=`.status.message`
// +kubebuilder:subresource:status
type CnvrgApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CnvrgAppSpec `json:"spec,omitempty"`
	Status Status       `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type CnvrgAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CnvrgApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CnvrgApp{}, &CnvrgAppList{})
}

func DefaultCnvrgAppSpec() CnvrgAppSpec {
	return CnvrgAppSpec{
		ClusterDomain: "",
		ImageHub:      "docker.io/cnvrg",
		ControlPlane:  controlPlaneDefault,
		Registry:      appRegistryDefault,
		Dbs:           appDbsDefaults,
		Logging:       cnvrgAppLoggingDefault,
		Networking:    cnvrgAppNetworkingDefault,
		Monitoring:    cnvrgAppMonitoringDefault,
		SSO:           ssoDefault,
		Tenancy:       tenancyDefault,
		Labels:        map[string]string{"owner": "cnvrg-control-plane"},
		Annotations:   nil,
		Proxy:         defaultProxy,
	}
}
