package common

const (
	LabelHostname           = "kubeclipper.io/hostname"
	LabelOSStable           = "kubeclipper.io/os"
	LabelArchStable         = "kubeclipper.io/arch"
	LabelTopologyZone       = "topology.kubeclipper.io/zone"
	LabelTopologyRegion     = "topology.kubeclipper.io/region"
	LabelNodeRole           = "kubeclipper.io/nodeRole"
	LabelNodeDisable        = "kubeclipper.io/nodeDisable"
	LabelCreator            = "kubeclipper.io/creator"
	LabelUsername           = "kubeclipper.io/username"
	LabelClusterName        = "kubeclipper.io/cluster"
	LabelBackupName         = "kubeclipper.io/backup"
	LabelRecoveryName       = "kubeclipper.io/recovery"
	LabelOperationAction    = "kubeclipper.io/operation"
	LabelOperationName      = "kubeclipper.io/operation-name"
	LabelOperationRetry     = "kubeclipper.io/operation-retry-times"
	LabelOperationIntent    = "kubeclipper.io/operation-intent"
	LabelOperationSponsor   = "kubeclipper.io/operation-sponsor"
	LabelTimeoutSeconds     = "kubeclipper.io/timeout"
	LabelRoleTemplate       = "kubeclipper.io/role-template"
	LabelHidden             = "kubeclipper.io/hidden"
	LabelUserReference      = "iam.kubeclipper.io/user-ref"
	LabelExternalIP         = "kubeclipper.io/externalIP"
	LabelExternalPort       = "kubeclipper.io/externalPort"
	LabelExternalDomain     = "kubeclipper.io/externalDomain"
	LabelExternalDomainPort = "kubeclipper.io/externalDomainPort"
	LabelUpgradeVersion     = "kubeclipper.io/upgrade-version"
	LabelBackupPoint        = "kubeclipper.io/backupPoint"
	LabelCronBackupDisable  = "kubeclipper.io/cronBackupDisable"
	LabelCronBackupEnable   = "kubeclipper.io/cronBackupEnable"

	LabelClusterProviderType = "kubeclipper.io/clusterProviderType"
	LabelClusterProviderName = "kubeclipper.io/clusterProviderName"
)

const (
	ResourceKindGlobalRole = "GlobalRole"
)

const (
	AnnotationAggregationRoles = "kubeclipper.io/aggregation-roles"
	RegoOverrideAnnotation     = "kubeclipper.io/rego-override"
	RoleAnnotation             = "iam.kubeclipper.io/role"
	AnnotationInternal         = "kubeclipper.io/internal"
	AnnotationHidden           = "kubeclipper.io/hidden"

	AnnotationMetadataFloatIP        = "metadata.kubeclipper.io/floatIP"
	AnnotationMetadataProxyServer    = "metadata.kubeclipper.io/proxyServer"
	AnnotationMetadataProxyAPIServer = "metadata.kubeclipper.io/proxyAPIServer"
	AnnotationMetadataProxySSH       = "metadata.kubeclipper.io/proxySSH"

	// AnnotationOnlyInstallKubernetesComp mean not install cni when create cluster
	AnnotationOnlyInstallKubernetesComp = "kubeclipper.io/only-install-kubernetes-component"
)

type NodeRole string // master/worker/ingress(worker)

const (
	NodeRoleMaster NodeRole = "master"
	NodeRoleWorker NodeRole = "worker"
)

func (nr NodeRole) String() string {
	return string(nr)
}

const (
	LabelIDP       = "iam.kubeclipper.io/idp"
	LabelOriginUID = "iam.kubeclipper.io/origin-uid"
)

const (
	// eg: cinder/v1
	LabelComponentName    = "kubeclipper.io/componentName"
	LabelComponentVersion = "kubeclipper.io/componentVersion"
	// eg: storage
	LabelCategory = "kubeclipper.io/category"
	// eg: name
	AnnotationActualName  = "kubeclipper.io/actual-name"
	AnnotationDisplayName = "kubeclipper.io/display-name"
	AnnotationDescription = "kubeclipper.io/description"
	AnnotationOffline     = "kubeclipper.io/offline"

	AnnotationProviderSyncTime  = "kubeclipper.io/providerSyncTime"
	AnnotationProviderNodeID    = "kubeclipper.io/providerNodeID"    // provider's nodeID,just mark
	AnnotationOriginNode        = "kubeclipper.io/originNode"        // mark node is from kc or provider,add when join node to provider cluster
	AnnotationProviderClusterID = "kubeclipper.io/providerClusterID" // match kc cluster to provider cluster
)

const OperationIntent = "termination"
