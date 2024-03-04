package deploy

import (
	"testing"

	"sigs.k8s.io/yaml"

	"github.com/kubeclipper/kubeclipper/cmd/kcctl/app/options"
)

func TestDeployOptions_GenDefaultConfig(t *testing.T) {
	omitempty, err := options.Omitempty([]byte(configTemplate))
	if err != nil {
		t.Fatal(err)
	}
	d := options.NewDeployOptions()
	err = yaml.Unmarshal(omitempty, d)
	if err != nil {
		return
	}
	marshal, err := yaml.Marshal(d)
	if err != nil {
		return
	}
	t.Log(string(marshal))
}
