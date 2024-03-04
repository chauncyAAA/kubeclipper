package k8s

const (
	K8s       = "k8s"
	CniCalico = "calico"

	NodeRoleMaster = "master"
	NodeRoleWorker = "worker"

	APIServerDomainPrefix = "apiserver."

	KubeConfigDir            = ".kube"
	CniDefaultConfigDir      = "/etc/cni"
	K8SDefaultConfigDir      = "/etc/kubernetes"
	ManifestDir              = "/tmp/.k8s"
	DockershimDefaultDataDir = "/var/lib/dockershim"
	EtcdDefaultDataDir       = "/var/lib/etcd"
	KubeletDefaultDataDir    = "/var/lib/kubelet"
	KubeletSystemdDir        = "/etc/systemd/system"
	Kubelet10KubeadmDir      = "/etc/systemd/system/kubelet.service.d"
	KubeBinaryDir            = "/usr/bin"
	KubeManifestsDir         = "/etc/kubernetes/manifests"
	// KubeletSystemdResolverConfig specifies the default resolver config when systemd service is active
	KubeletSystemdResolverConfig = "/run/systemd/resolve/resolv.conf"
	KubeletDefaultResolvConf     = "/etc/resolv.conf"
)
