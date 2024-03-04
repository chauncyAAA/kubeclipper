package filters

import (
	"bytes"
	"fmt"
	"net/http"
	rt "runtime"
	"time"

	restful "github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"

	"github.com/kubeclipper/kubeclipper/pkg/logger"
	"github.com/kubeclipper/kubeclipper/pkg/utils/netutil"
)

func LogRequestAndResponse(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	start := time.Now()
	chain.ProcessFilter(req, resp)

	// Always log error response
	if resp.StatusCode() > http.StatusBadRequest {
		logger.Warnf("%s - \"%s %s %s\" %d %d %dms",
			netutil.GetRequestIP(req.Request),
			req.Request.Method,
			req.Request.URL,
			req.Request.Proto,
			resp.StatusCode(),
			resp.ContentLength(),
			time.Since(start)/time.Millisecond,
		)
	}
	logger.Debugf(
		"%s - \"%s %s %s\" %d %d %dms",
		netutil.GetRequestIP(req.Request),
		req.Request.Method,
		req.Request.URL,
		req.Request.Proto,
		resp.StatusCode(),
		resp.ContentLength(),
		time.Since(start)/time.Millisecond,
	)
}

func LogStackOnRecover(panicReason interface{}, w http.ResponseWriter) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("recover from panic situation: - %v\r\n", panicReason))
	for i := 2; ; i++ {
		_, file, line, ok := rt.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", file, line))
	}
	logger.Error("recover from http server panic", zap.String("err", buffer.String()))

	headers := http.Header{}
	if ct := w.Header().Get("Content-Type"); len(ct) > 0 {
		headers.Set("Accept", ct)
	}

	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("Internal server error"))
}
