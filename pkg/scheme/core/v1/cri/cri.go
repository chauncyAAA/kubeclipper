package cri

import (
	"fmt"

	"github.com/kubeclipper/kubeclipper/pkg/component"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

func init() {
	if err := component.RegisterAgentStep(
		fmt.Sprintf(component.RegisterStepKeyFormat, criContainerd, criVersion, component.TypeStep),
		&ContainerdRunnable{}); err != nil {
		panic(err)
	}

	if err := component.RegisterAgentStep(
		fmt.Sprintf(component.RegisterStepKeyFormat, criDocker, criVersion, component.TypeStep),
		&DockerRunnable{}); err != nil {
		panic(err)
	}

	if err := component.RegisterAgentStep(
		ContainerdRegistryConfigureIdentity,
		&ContainerdRegistryConfigure{}); err != nil {
		panic(err)
	}

	if err := component.RegisterAgentStep(
		DockerInsecureRegistryConfigureIdentity,
		&DockerInsecureRegistryConfigure{}); err != nil {
		panic(err)
	}
}

const (
	criDocker     = "docker"
	criContainerd = "containerd"
	criVersion    = "v1"
)

const (
	// dockerDefaultVersion    = "20.10.13"
	dockerDefaultConfigDir = "/etc/docker"
	dockerDefaultDataDir   = "/var/lib/docker"
	// dockerDefaultSystemdDir = "/etc/systemd/system"
	dockerDefaultCriDir = "/etc/containerd"

	// containerdDefaultVersion    = "1.6.4"
	containerdDefaultConfigDir         = "/etc/containerd"
	ContainerdDefaultRegistryConfigDir = "/etc/containerd/certs.d"
	// containerdDefaultSystemdDir = "/etc/systemd/system"
	containerdDefaultDataDir = "/var/lib/containerd"
)

var (
	DockerInsecureRegistryConfigureIdentity = fmt.Sprintf(
		component.RegisterStepKeyFormat, criDocker, criVersion, component.TypeRegistryConfigure)
	ContainerdRegistryConfigureIdentity = fmt.Sprintf(
		component.RegisterStepKeyFormat, criContainerd, criVersion, component.TypeRegistryConfigure)
)

var k8sMatchPauseVersion = map[string]string{
	"118": "3.2",
	"119": "3.2",
	"120": "3.2",
	"121": "3.4.1",
	"122": "3.5",
	"123": "3.6",
	"124": "3.7",
	"125": "3.8",
	"126": "3.9",
	"127": "3.9",
	"128": "3.9",
}

var _ component.StepRunnable = (*ContainerdRunnable)(nil)
var _ component.StepRunnable = (*DockerRunnable)(nil)

type Base struct {
	Version     string            `json:"version,omitempty"`
	Offline     bool              `json:"offline"`
	DataRootDir string            `json:"rootDir"`
	Registies   []v1.RegistrySpec `json:"registry,omitempty"`
}
