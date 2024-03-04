package component

import (
	"errors"
	"strings"
)

var _tmpl = defaultTmpl()

var (
	ErrTemplateExist     = errors.New("component template already exist")
	ErrTemplateKeyFormat = errors.New("component template key must be name/version/templateName")
)

const (
	RegisterTemplateKeyFormat = "%s/%s/%s"
	RegisterStepKeyFormat     = "%s/%s/%s"
)

type tmpl struct {
	template map[string]TemplateRender
}

func defaultTmpl() tmpl {
	return tmpl{template: map[string]TemplateRender{}}
}

func RegisterTemplate(kv string, t TemplateRender) error {
	if !checkTemplateKey(kv) {
		return ErrTemplateKeyFormat
	}
	return _tmpl.registerTemplate(kv, t)
}

func LoadTemplate(kv string) (TemplateRender, bool) {
	return _tmpl.load(kv)
}

func (h *tmpl) load(kv string) (TemplateRender, bool) {
	c, exist := h.template[kv]
	return c, exist
}

func (h *tmpl) registerTemplate(kv string, p TemplateRender) error {
	_, exist := h.template[kv]
	if exist {
		return ErrTemplateExist
	}
	h.template[kv] = p
	return nil
}

func checkTemplateKey(kv string) bool {
	parts := strings.Split(kv, "/")
	return len(parts) == 3
}
