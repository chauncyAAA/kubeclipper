package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Event struct {
	metav1.TypeMeta          `json:",inline"`
	metav1.ObjectMeta        `json:"metadata,omitempty"`
	AuditID                  string `json:"auditID,omitempty"`
	RequestURI               string `json:"requestURI"`
	UserID                   string `json:"userID,omitempty"`
	Username                 string `json:"username,omitempty"`
	Verb                     string `json:"verb"`
	Type                     string `json:"type,omitempty"`
	SourceIP                 string `json:"sourceIP"`
	UserAgent                string `json:"userAgent,omitempty"`
	Success                  bool   `json:"success"`
	RequestReceivedTimestamp metav1.MicroTime
	StageTimestamp           metav1.MicroTime
	Resource                 string `json:"resource"`
	ResourceName             string `json:"resourceName"`
	Subresource              string `json:"subresource"`
	ResourceAPIGroup         string `json:"resourceAPIGroup"`
	ResourceAPIVersion       string `json:"resourceAPIVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EventList contains a list of EventList
type EventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Event `json:"items"`
}
