package web

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"
)

func TLSClientConfig(caFile string) (*tls.Config, error) {
	pem, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	cas := x509.NewCertPool()
	cas.AppendCertsFromPEM(pem)

	tlsConfig := &tls.Config{
		RootCAs: cas,
	}

	return tlsConfig, nil
}

func Client(ca string, timeout time.Duration) (*http.Client, error) {
	tlsConfig, err := TLSClientConfig(ca)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return client, nil
}
