package runtime

import (
	restful "github.com/emicklei/go-restful/v3"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	APIRootPath = "/api"
)

const (
	MimeMergePatchJSON    = "application/merge-patch+json"
	MimeJSONPatchJSON     = "application/json-patch+json"
	MimeMultipartFormData = "multipart/form-data"
)

type ContainerBuilder []func(c *restful.Container) error

func init() {
	restful.RegisterEntityAccessor(MimeMergePatchJSON, restful.NewEntityAccessorJSON(restful.MIME_JSON))
	restful.RegisterEntityAccessor(MimeJSONPatchJSON, restful.NewEntityAccessorJSON(restful.MIME_JSON))
}

func NewWebService(gv schema.GroupVersion) *restful.WebService {
	webservice := restful.WebService{}
	webservice.Path(APIRootPath + "/" + gv.String()).
		Produces(restful.MIME_JSON)

	return &webservice
}

func (cb *ContainerBuilder) AddToContainer(c *restful.Container) error {
	for _, f := range *cb {
		if err := f(c); err != nil {
			return err
		}
	}
	return nil
}

func (cb *ContainerBuilder) Register(funcs ...func(*restful.Container) error) {
	for _, f := range funcs {
		*cb = append(*cb, f)
	}
}
