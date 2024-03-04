package staticresource

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/kubeclipper/kubeclipper/pkg/logger"
	"github.com/kubeclipper/kubeclipper/pkg/service"
	"github.com/kubeclipper/kubeclipper/pkg/simple/staticserver"
)

var _ service.Interface = (*Service)(nil)

type Service struct {
	server *http.Server
	path   string
}

func NewService(opts *staticserver.Options) (service.Interface, error) {
	httpSrv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", opts.BindAddress, opts.InsecurePort),
	}
	if opts.SecurePort != 0 {
		certificate, err := tls.LoadX509KeyPair(opts.TLSCertFile, opts.TLSPrivateKey)
		if err != nil {
			return nil, err
		}
		httpSrv.TLSConfig.Certificates = []tls.Certificate{certificate}
		httpSrv.Addr = fmt.Sprintf("%s:%d", opts.BindAddress, opts.SecurePort)
	}
	return &Service{
		server: httpSrv,
		path:   opts.Path,
	}, nil
}

func (s *Service) PrepareRun(stopCh <-chan struct{}) error {
	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		return os.MkdirAll(s.path, os.ModeDir|0755)
	}
	s.server.Handler = http.StripPrefix("/", http.FileServer(http.Dir(s.path)))
	return nil
}

func (s *Service) Run(stopCh <-chan struct{}) error {
	logger.Info("Static resource server start", zap.String("addr", s.server.Addr), zap.String("path", s.path))
	go func() {
		<-stopCh
		_ = s.server.Shutdown(context.TODO())
	}()
	go func() {
		var err error
		if s.server.TLSConfig != nil {
			err = s.server.ListenAndServeTLS("", "")
		} else {
			err = s.server.ListenAndServe()
		}
		logger.Error("static resource server exit", zap.Error(err))
	}()

	return nil
}

func (s *Service) Close() {
	s.server.Close()
}
