package component

import (
	"errors"
	"strings"

	"k8s.io/apimachinery/pkg/labels"
)

const (
	RegisterFormat = "%s/%s"

	TypeStep              = "step"
	TypeTemplate          = "template"
	TypeRegistryConfigure = "registryConfigure"
)

var (
	_components = defaultComponentHandler()
)

var (
	ErrExist     = errors.New("component already exist")
	ErrKeyFormat = errors.New("component key must be name/version")
)

// TODO: Refactor to sync map
type handler struct {
	componentsMap map[string]Interface
}

func defaultComponentHandler() handler {
	return handler{componentsMap: map[string]Interface{}}
}

// Register KV must format at componentName/version
func Register(kv string, p Interface) error {
	if !checkComponentKey(kv) {
		return ErrKeyFormat
	}
	return _components.registerComponent(kv, p)
}

func Load(kv string) (Interface, bool) {
	return _components.load(kv)
}

func GetJSONSchemas(selector labels.Selector, lang Lang) []Meta {
	selectAll := selector.Empty()
	var result []Meta
	for _, v := range _components.componentsMap {
		if selectAll {
			result = append(result, v.GetComponentMeta(lang))
			continue
		}
		if selector.Matches(labels.Set(map[string]string{
			"category": v.GetComponentMeta(lang).Category,
			"name":     v.GetComponentMeta(lang).Name,
		})) {
			result = append(result, v.GetComponentMeta(lang))
		}
	}
	return result
}

func (h *handler) load(kv string) (Interface, bool) {
	c, exist := h.componentsMap[kv]
	return c, exist
}

func (h *handler) registerComponent(kv string, p Interface) error {
	_, exist := h.componentsMap[kv]
	if exist {
		return ErrExist
	}
	h.componentsMap[kv] = p
	return nil
}

func checkComponentKey(kv string) bool {
	parts := strings.Split(kv, "/")
	return len(parts) == 2
}
