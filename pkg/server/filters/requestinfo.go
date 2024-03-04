package filters

import (
	restful "github.com/emicklei/go-restful/v3"

	"github.com/kubeclipper/kubeclipper/pkg/logger"
	"github.com/kubeclipper/kubeclipper/pkg/server/request"
	"github.com/kubeclipper/kubeclipper/pkg/server/restplus"
)

func WithRequestInfo(resolver request.Resolver) restful.FilterFunction {
	if resolver == nil {
		logger.Warn("RequestInfo is disabled")
		return nil
	}
	return func(r *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		ctx := r.Request.Context()
		info, err := resolver.NewRequestInfo(r.Request)
		if err != nil {
			restplus.HandleInternalError(response, r, err)
			return
		}
		r.Request = r.Request.WithContext(request.WithInfo(ctx, info))
		chain.ProcessFilter(r, response)
	}
}
