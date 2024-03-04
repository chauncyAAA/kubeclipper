package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false

type Domain struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec,omitempty"`
	Status            DomainStatus `json:"status,omitempty"`
}

type DomainSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Records map[string]Record `json:"records,omitempty"` // key: rr.domain value: record
	// +optional
	SyncCluster []string `json:"syncCluster,omitempty"`
}

type DomainStatus struct {
	Count int64 `json:"count"` // update by informer
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DomainList contains a list of Domain
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

type Record struct {
	Domain      string        `json:"domain,omitempty"` // domain e.g. baidu.com
	RR          string        `json:"rr,omitempty"`     // resource record e.g. www
	CreateTime  metav1.Time   `json:"createTime,omitempty"`
	Description string        `json:"description,omitempty"`
	ParseRecord []ParseRecord `json:"parseRecord,omitempty"`
}

type ParseRecord struct {
	Type string `json:"type,omitempty"` // resolve record. A or AAAA
	IP   string `json:"ip,omitempty"`   // ipv4 or ipv6
}
