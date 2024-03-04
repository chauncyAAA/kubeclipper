package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false

type PlatformSetting struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Template          DockerRegistry `json:"template,omitempty"`
	Terminal          WebTerminal    `json:"terminal,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// UserList contains a list of User

type PlatformSettingList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlatformSetting `json:"items"`
}

type DockerRegistry struct {
	InsecureRegistry []InsecureRegistry `json:"insecureRegistry"`
}

type InsecureRegistry struct {
	Host        string `json:"host"`
	Description string `json:"description,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	// +optional
	CreateAt metav1.Time `json:"createAt,omitempty"`
}

type WebTerminal struct {
	PrivateKey string `json:"privateKey,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
}
