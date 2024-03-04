package downloader

var (
	CloudStaticServer = "https://oss.kubeclipper.io/packages"
)

type Options struct {
	Address       string `json:"address" yaml:"address"`
	TLSCertFile   string `json:"tlsCertFile" yaml:"tlsCertFile"`
	TLSPrivateKey string `json:"tlsPrivateKey" yaml:"tlsPrivateKey"`
}

func NewOptions() *Options {
	return &Options{}
}

type ManifestElement struct {
	Name   string `json:"name"`
	Digest string `json:"digest"`
	Path   string `json:"path"`
}
