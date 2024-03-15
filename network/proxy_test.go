package network

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"

	"golang.org/x/net/proxy"
)

// http 代理测试
func TestHttpProxy(t *testing.T) {
	url, _ := url.Parse("http://admin:123456@10.248.162.60:8080")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(url),
		},
	}
	result, err := client.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", result)
}

// socks5 代理测试
func TestSocks5Proxy(t *testing.T) {
	url, _ := url.Parse("socks5://admin:123456@10.248.162.60:8080")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(url),
		},
	}
	result, err := client.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", result)
}

// socks5h 代理测试
func TestSocks5HProxy(t *testing.T) {
	dialer, err := proxy.SOCKS5("tcp", "10.248.162.60:8080", &proxy.Auth{
		User:     "admin",
		Password: "123456",
	}, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
	}
	client := http.Client{
		Transport: &http.Transport{DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.Dial(network, address)
		}},
	}
	result, err := client.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", result)
}
