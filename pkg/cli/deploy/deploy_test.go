package deploy

import (
	"testing"

	"github.com/google/uuid"

	"github.com/kubeclipper/kubeclipper/cmd/kcctl/app/options"
)

func TestDeployOptions_getEtcdTemplateContent(t *testing.T) {
	d := NewDeployOptions(options.IOStreams{})
	d.deployConfig.ServerIPs = []string{"192.168.234.3", "192.168.234.4", "192.168.234.5"}
	d.servers = map[string]string{
		"192.168.234.3": "master1",
		"192.168.234.4": "master2",
		"192.168.234.5": "master3",
	}

	for _, s := range d.deployConfig.ServerIPs {
		t.Log(d.getEtcdTemplateContent(s))
	}
}

func TestDeployOptions_getKcServerConfigTemplateContent(t *testing.T) {
	d := NewDeployOptions(options.IOStreams{})
	d.deployConfig.ServerIPs = []string{"192.168.234.3", "192.168.234.4", "192.168.234.5"}
	d.servers = map[string]string{
		"192.168.234.3": "master1",
		"192.168.234.4": "master2",
		"192.168.234.5": "master3",
	}

	for _, s := range d.deployConfig.ServerIPs {
		t.Log(d.deployConfig.GetKcServerConfigTemplateContent(s))
	}
}

func TestDeployOptions_getKcAgentConfigTemplateContent(t *testing.T) {
	d := NewDeployOptions(options.IOStreams{})
	d.deployConfig.ServerIPs = []string{"192.168.234.3", "192.168.234.4", "192.168.234.5"}
	d.servers = map[string]string{
		"192.168.234.3": "master1",
		"192.168.234.4": "master2",
		"192.168.234.5": "master3",
	}
	metadata := options.Metadata{
		Region:  d.deployConfig.DefaultRegion,
		FloatIP: "1.1.1.1",
	}
	for range d.deployConfig.ServerIPs {
		metadata.AgentID = uuid.New().String()
		t.Log(d.deployConfig.GetKcAgentConfigTemplateContent(metadata))
	}
}

func TestDeployOptions_getKcConsoleTemplateContent(t *testing.T) {
	d := NewDeployOptions(options.IOStreams{})
	d.deployConfig.ServerIPs = []string{"192.168.234.3", "192.168.234.4", "192.168.234.5"}
	d.servers = map[string]string{
		"192.168.234.3": "master1",
		"192.168.234.4": "master2",
		"192.168.234.5": "master3",
	}

	t.Log(d.getKcConsoleTemplateContent())
}
