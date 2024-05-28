package internal

import (
	"encoding/json"
	"os"
)

type ProxyConfig struct {
	Proxy struct {
		ProxyServer string `json:"proxy_server"`
	} `json:"Proxy"`
}

func ReadProxyConfig(filename string) (*ProxyConfig, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var proxyConfig ProxyConfig
    err = json.Unmarshal(data, &proxyConfig)
    if err != nil {
        return nil, err
    }

    return &proxyConfig, nil
}