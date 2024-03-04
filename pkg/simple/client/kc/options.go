package kc

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"
)

var (
	JSONContentTypeHeader = map[string][]string{
		"Content-Type": {"application/json"},
	}
)

type Opt func(*Client) error

func WithEndpoint(endpoint string) Opt {
	return func(c *Client) error {
		endpointURL, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		c.host = endpointURL.Host
		c.scheme = endpointURL.Scheme
		// c.basePath = endpointURL.Path
		return nil
	}
}

func WithHost(h string) Opt {
	return func(c *Client) error {
		c.host = h
		return nil
	}
}

// WithHTTPClient overrides the client http client with the specified one
func WithHTTPClient(client *http.Client) Opt {
	return func(c *Client) error {
		if client != nil {
			c.client = client
		}
		return nil
	}
}

// WithScheme overrides the client scheme with the specified one
func WithScheme(scheme string) Opt {
	return func(c *Client) error {
		c.scheme = scheme
		return nil
	}
}

func WithBearerAuth(token string) Opt {
	return func(c *Client) error {
		c.bearerToken = token
		return nil
	}
}

func WithCAData(ca []byte) Opt {
	return func(client *Client) error {
		caPool := x509.NewCertPool()
		caPool.AppendCertsFromPEM(ca)
		client.caPool = caPool
		return nil
	}
}

func WithCertData(cert, key []byte) Opt {
	return func(client *Client) error {
		pair, err := tls.X509KeyPair(cert, key)
		if err != nil {
			return err
		}
		client.cliCert = &pair
		return nil
	}
}

func WithInsecureSkipTLSVerify() Opt {
	return func(client *Client) error {
		client.insecureSkipTLSVerify = true
		return nil
	}
}

func WithServerName(name string) Opt {
	return func(client *Client) error {
		client.tlsServerName = name
		return nil
	}
}
