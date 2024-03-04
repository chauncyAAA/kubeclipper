package registry

import (
	"github.com/kubeclipper/kubeclipper/pkg/cli/printer"
)

type Image struct {
	Name string   `json:"name" yaml:"name"`
	Tags []string `json:"tags" yaml:"tags"`
}

func (i *Image) JSONPrint() ([]byte, error) {
	return printer.JSONPrinter(i)
}

func (i *Image) YAMLPrint() ([]byte, error) {
	return printer.YAMLPrinter(i)
}

func (i *Image) TablePrint() ([]string, [][]string) {
	headers := []string{"name", "tags"}
	var data [][]string
	for index, v := range i.Tags {
		if index == 0 {
			data = append(data, []string{i.Name, v})
		} else {
			data = append(data, []string{"", v})
		}
	}
	return headers, data
}

type Repositories struct {
	Repositories []string `json:"repositories" yaml:"repositories"`
}

func (i *Repositories) JSONPrint() ([]byte, error) {
	return printer.JSONPrinter(i)
}

func (i *Repositories) YAMLPrint() ([]byte, error) {
	return printer.YAMLPrinter(i)
}

func (i *Repositories) TablePrint() ([]string, [][]string) {
	headers := []string{"repositories"}
	var data [][]string
	for _, v := range i.Repositories {
		data = append(data, []string{v})
	}
	return headers, data
}
