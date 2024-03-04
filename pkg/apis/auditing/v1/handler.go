package v1

import (
	"net/http"

	apimachineryErrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/kubeclipper/kubeclipper/pkg/utils/strutil"

	restful "github.com/emicklei/go-restful/v3"

	"github.com/kubeclipper/kubeclipper/pkg/models/platform"
	"github.com/kubeclipper/kubeclipper/pkg/query"
	"github.com/kubeclipper/kubeclipper/pkg/server/restplus"
)

type handler struct {
	operator platform.Operator
}

func newHandler(operator platform.Operator) *handler {
	return &handler{
		operator: operator,
	}
}

func (h *handler) ListEvents(req *restful.Request, resp *restful.Response) {
	q := query.ParseQueryParameter(req)
	result, err := h.operator.ListEventsEx(req.Request.Context(), q)
	if err != nil {
		restplus.HandleInternalError(resp, req, err)
		return
	}
	_ = resp.WriteHeaderAndEntity(http.StatusOK, result)
}

func (h *handler) DescribeEvent(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter(query.ParameterName)
	resourceVersion := strutil.StringDefaultIfEmpty("0", req.QueryParameter(query.ParameterResourceVersion))
	c, err := h.operator.GetEventEx(req.Request.Context(), name, resourceVersion)
	if err != nil {
		if apimachineryErrors.IsNotFound(err) {
			restplus.HandleNotFound(resp, req, err)
			return
		}
		restplus.HandleInternalError(resp, req, err)
		return
	}
	_ = resp.WriteHeaderAndEntity(http.StatusOK, c)
}
