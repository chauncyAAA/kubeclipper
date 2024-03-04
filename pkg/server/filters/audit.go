package filters

import (
	"strings"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientrest"

	restful "github.com/emicklei/go-restful/v3"

	"github.com/kubeclipper/kubeclipper/pkg/auditing"
	"github.com/kubeclipper/kubeclipper/pkg/logger"
	"github.com/kubeclipper/kubeclipper/pkg/server/request"
)

func WithAudit(p auditing.Interface) restful.FilterFunction {
	if p == nil {
		logger.Warn("Authentication is disabled")
		return nil
	}
	a := PathExclude{
		prefixes: []string{"oauth/"},
	}

	return func(req *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		if !p.Enabled() || clientrest.IsInformerRawQuery(req.Request) {
			chain.ProcessFilter(req, response)
			return
		}
		pth := strings.TrimPrefix(req.Request.URL.Path, "/")
		info, ok := request.InfoFrom(req.Request.Context())
		if !ok || (!info.IsResourceRequest && !a.hasPrefix(pth)) {
			chain.ProcessFilter(req, response)
			return
		}
		e := p.LogRequestObject(req.Request, info)
		if e != nil {
			respCapture := auditing.NewResponseCapture(response.ResponseWriter)
			response.ResponseWriter = respCapture
			chain.ProcessFilter(req, response)
			go p.LogResponseObject(e, respCapture)
			return
		}
		chain.ProcessFilter(req, response)
	}
}
