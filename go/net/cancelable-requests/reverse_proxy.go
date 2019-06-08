package main

import (
	// "bytes"
	"context"
	// "fmt"
	// "io"
	// "io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

type directorFunc func(req *http.Request)

func director(target *url.URL) directorFunc {
	targetQuery := target.RawQuery
	return func(req *http.Request) {
		log.Printf(" > RevProxy: CALLING Director %v", req.URL)
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

}

func newReverseProxy(target *url.URL) *httputil.ReverseProxy {

	// var transport http.RoundTripper = &http.Transport{
	// Proxy: http.ProxyFromEnvironment,
	// DialContext: (&net.Dialer{
	// Timeout:   3 * time.Second,
	// KeepAlive: 3 * time.Second,
	// DualStack: true,
	// }).DialContext,
	// MaxIdleConns:          100,
	// IdleConnTimeout:       3 * time.Second,
	// TLSHandshakeTimeout:   10 * time.Second,
	// ExpectContinueTimeout: 1 * time.Second,
	// }

	// copied from:

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			log.Println(" > RevProxy: CALLING Proxy")
			return http.ProxyFromEnvironment(req)
		},

		Dial: func(network, addr string) (net.Conn, error) {
			log.Println(" > RevProxy: CALLING Dial")

			dialer := &net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}

			conn, err := dialer.Dial(network, addr)
			if err != nil {
				log.Printf(" > RevProxy: Dial - Error: %v", err)
			}
			return conn, err
		},

		TLSHandshakeTimeout: 10 * time.Second,
	}

	return &httputil.ReverseProxy{
		Director:  director(target),
		Transport: transport,
	}
}

func startReverseProxy(ctx context.Context, wg *sync.WaitGroup, addr, target string) {
	targetURL, _ := url.Parse(target)

	proxy := newReverseProxy(targetURL)
	startServer(ctx, wg, addr, proxy)
}
