package filters

import (
	restful "github.com/emicklei/go-restful/v3"
	"k8s.io/apiserver/pkg/authentication/authenticator"

	"github.com/kubeclipper/kubeclipper/pkg/logger"
	"github.com/kubeclipper/kubeclipper/pkg/server/request"
	"github.com/kubeclipper/kubeclipper/pkg/server/restplus"
)

func WithAuthentication(authRequest authenticator.Request) restful.FilterFunction {
	if authRequest == nil {
		logger.Warn("Authentication is disabled")
		return nil
	}
	return func(req *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		resp, ok, err := authRequest.AuthenticateRequest(req.Request)
		if err != nil || !ok {
			restplus.HandleUnauthorized(response, req, err)
			return
		}
		req.Request = req.Request.WithContext(request.WithUser(req.Request.Context(), resp.User))
		chain.ProcessFilter(req, response)
	}
}
