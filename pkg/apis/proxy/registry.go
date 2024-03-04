package proxy

import (
	restful "github.com/emicklei/go-restful/v3"

	"github.com/kubeclipper/kubeclipper/pkg/models/cluster"
)

func AddToContainer(c *restful.Container, clusterReader cluster.ClusterReader) error {
	c.HandleWithFilter(ProxyHandlerPrefix, newHandler(clusterReader))
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)
	return nil
}
