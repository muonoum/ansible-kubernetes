package web

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func Proxy(tlsConfig *tls.Config, url *url.URL, timeout time.Duration, keepalive time.Duration) http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.FlushInterval = 100 * time.Millisecond
	proxy.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
		Proxy:           http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepalive,
		}).Dial,
	}

	return proxy
}
