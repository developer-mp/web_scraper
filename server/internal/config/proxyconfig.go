package internal

import (
	config "server/internal"
)

type ProxyConfig struct {
	Proxy struct {
		ProxyServer string `json:"proxy_server"`
	} `json:"Proxy"`
}

func ReadProxyConfig(filename string) (*ProxyConfig, error) {
	var proxyConfig ProxyConfig
	err := config.ReadConfig(filename, &proxyConfig)
	if err != nil {
		return nil, err
	}
	return &proxyConfig, nil
}