package config

import (
	"github.com/razorpay/MachineRound/internal/constants"
	"net"
	"net/http"
	"time"
)

type HTTPClient struct {
	ConnectTimeoutMs        int `toml:"connectTimeoutMs"`
	ConnKeepAliveMs         int `toml:"connKeepAliveMs"`
	ExpectContinueTimeoutMs int `toml:"expectContinueTimeoutMs"`
	IdleConnTimeoutMs       int `toml:"idleConnTimeoutMs"`
	MaxAllIdleConns         int `toml:"maxAllIdleConns"`
	MaxHostIdleConns        int `toml:"maxHostIdleConns"`
	ResponseHeaderTimeoutMs int `toml:"responseHeaderTimeoutMs"`
	TLSHandshakeTimeoutMs   int `toml:"tlsHandshakeTimeoutMs"`
}

func GetNewHTTPClient(httpSettings *HTTPClient) *http.Client {
	tr := &http.Transport{
		ResponseHeaderTimeout: time.Duration(httpSettings.ResponseHeaderTimeoutMs) * time.Millisecond,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: time.Duration(httpSettings.ConnKeepAliveMs) * time.Millisecond,
			Timeout:   time.Duration(httpSettings.ConnectTimeoutMs) * time.Millisecond,
		}).DialContext,
		MaxIdleConns:          httpSettings.MaxAllIdleConns,
		IdleConnTimeout:       time.Duration(httpSettings.IdleConnTimeoutMs) * time.Millisecond,
		TLSHandshakeTimeout:   time.Duration(httpSettings.TLSHandshakeTimeoutMs) * time.Millisecond,
		MaxIdleConnsPerHost:   httpSettings.MaxHostIdleConns,
		ExpectContinueTimeout: time.Duration(httpSettings.ExpectContinueTimeoutMs) * time.Millisecond,
	}
	return &http.Client{
		Transport: tr,
		Timeout:   constants.REQUEST_TIMEOUT,
	}
}
