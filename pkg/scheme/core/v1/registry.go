package v1

import (
	coordinationv1 "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	KindCluster   = "Cluster"
	KindConfigMap = "ConfigMap"
)

// GroupName is the group name used in this package
const GroupName = "core.kubeclipper.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// SchemeBuilder is the scheme builder with scheme init functions to run for this API package
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme is a common registration function for mapping packaged scoped group & version keys to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Cluster{},
		&ClusterList{},
		&Node{},
		&NodeList{},
		&Operation{},
		&OperationList{},
		&Region{},
		&RegionList{},
		&PlatformSetting{},
		&PlatformSettingList{},
		&Event{},
		&EventList{},
		&metav1.WatchEvent{},
		&coordinationv1.Lease{},
		&coordinationv1.LeaseList{},
		&Backup{},
		&BackupList{},
		&Recovery{},
		&RecoveryList{},
		&metav1.ListOptions{},
		&metav1.GetOptions{},
		&Domain{},
		&DomainList{},
		&BackupPoint{},
		&BackupPointList{},
		&CronBackup{},
		&CronBackupList{},
		&Template{},
		&TemplateList{},
		&ConfigMap{},
		&ConfigMapList{},
		&CloudProvider{},
		&CloudProviderList{},
		&Registry{},
		&RegistryList{},
	)
	return nil
}
