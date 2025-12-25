package groq

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

type ClientConfig struct {
	DialTimeout           time.Duration
	TLSHandShakeTimeout   time.Duration
	ResponseHeaderTimeout time.Duration
	SkipTLSVerify         bool
	ProxyURL              *url.URL
}

// returns a default client config for debugging
func DefaultClientConfig() ClientConfig {
	proxy, _ := url.Parse("http://127.0.0.1:8080")
	return ClientConfig{
		DialTimeout:           5 * time.Second,
		TLSHandShakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ProxyURL:              proxy,
		SkipTLSVerify:         true,
	}
}

func (c ClientConfig) CreateNewClient() *http.Client {
	customTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: c.DialTimeout,
		}).DialContext,
		TLSHandshakeTimeout:   c.TLSHandShakeTimeout,
		ResponseHeaderTimeout: c.ResponseHeaderTimeout,
	}
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: c.SkipTLSVerify}
	if c.ProxyURL != nil {
		customTransport.Proxy = http.ProxyURL(c.ProxyURL)
	}
	return &http.Client{Transport: customTransport}
}
