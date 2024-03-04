package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false

type BackupPoint struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	StorageType       string    `json:"storageType,omitempty"`
	Description       string    `json:"description,omitempty"`
	FsConfig          *FsConfig `json:"fsConfig,omitempty"`
	S3Config          *S3Config `json:"s3Config,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupPointList contains a list of BackupPoint

type BackupPointList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPoint `json:"items"`
}

type FsConfig struct {
	BackupRootDir string `json:"backupRootDir,omitempty" yaml:"backupRootDir,omitempty"`
}

type S3Config struct {
	Bucket          string `json:"bucket" yaml:"bucket"`
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `json:"accessKeyID" yaml:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
	Region          string `json:"region" yaml:"region"`
	SSL             bool   `json:"ssl" yaml:"ssl"`
}
