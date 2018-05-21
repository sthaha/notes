package main

import (
	// "bytes"
	"context"
	// "fmt"
	// "io"
	// "io/ioutil"
	// "log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
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

func newReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
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
	return &httputil.ReverseProxy{
		Director: director,
	}
}

func startReverseProxy(ctx context.Context, wg *sync.WaitGroup, addr, target string) {
	targetURL, _ := url.Parse(target)

	proxy := newReverseProxy(targetURL)
	startServer(ctx, wg, addr, proxy)
}
