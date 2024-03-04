package filters

import (
	"context"
	"errors"

	"github.com/kubeclipper/kubeclipper/pkg/server/request"

	restful "github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"

	"github.com/kubeclipper/kubeclipper/pkg/authorization/authorizer"
	"github.com/kubeclipper/kubeclipper/pkg/logger"
	"github.com/kubeclipper/kubeclipper/pkg/server/restplus"
)

func WithAuthorization(authorizers authorizer.Authorizer) restful.FilterFunction {
	if authorizers == nil {
		logger.Warn("Authorization is disabled")
		return nil
	}
	return func(req *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		ctx := req.Request.Context()
		attributes, err := getAuthorizerAttributes(ctx)
		if err != nil {
			restplus.HandleInternalError(response, req, err)
			return
		}
		authorized, reason, err := authorizers.Authorize(ctx, attributes)
		if authorized == authorizer.DecisionAllow {
			chain.ProcessFilter(req, response)
			return
		}
		if err != nil {
			restplus.HandleInternalError(response, req, err)
			return
		}
		logger.Debug("request forbidden", zap.String("url", req.Request.RequestURI),
			zap.String("reason", reason))

		restplus.HandleForbidden(response, req, errors.New(reason))
	}
}

func getAuthorizerAttributes(ctx context.Context) (authorizer.Attributes, error) {
	attribs := authorizer.AttributesRecord{}

	user, ok := request.UserFrom(ctx)
	if ok {
		attribs.User = user
	}

	requestInfo, found := request.InfoFrom(ctx)
	if !found {
		return nil, errors.New("no RequestInfo found in the context")
	}

	attribs.ResourceRequest = requestInfo.IsResourceRequest
	attribs.Path = requestInfo.Path
	attribs.Verb = requestInfo.Verb
	attribs.APIGroup = requestInfo.APIGroup
	attribs.APIVersion = requestInfo.APIVersion
	attribs.Resource = requestInfo.Resource
	attribs.Subresource = requestInfo.Subresource
	attribs.Name = requestInfo.Name

	return &attribs, nil
}
