package proxy

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	config "server/internal"
	"sync"
	"time"
)

var proxyList []string
var proxyIndex int
var mutex sync.Mutex

func InitProxy(configFile string) error {
	cfg, err := config.ReadProxyConfig(configFile)
	if err != nil {
		return err
	}

	proxyList = []string{cfg.Proxy.ProxyServer}
	if len(proxyList) == 0 {
		return errors.New("no proxies found in configuration")
	}

	proxyIndex = 0
	return nil
}

func getNextProxy() string {
    mutex.Lock()
    defer mutex.Unlock()
    proxy := proxyList[proxyIndex]
    proxyIndex = (proxyIndex + 1) % len(proxyList)
    return proxy
}

func createHttpClient(proxyURL string) (*http.Client, error) {
    proxy, err := url.Parse(proxyURL)
    if err != nil {
        return nil, err
    }

    transport := &http.Transport{
        Proxy: http.ProxyURL(proxy),
    }

    return &http.Client{
        Transport: transport,
        Timeout:   10 * time.Second,
    }, nil
}

func ScrapeWithProxy(url string) (*http.Response, error) {
    proxy := getNextProxy()
    client, err := createHttpClient(proxy)
    if err != nil {
        return nil, fmt.Errorf("error creating HTTP client: %v", err)
    }

    resp, err := client.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making request with proxy %s: %v", proxy, err)
    }

    return resp, nil
}
