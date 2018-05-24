package main

import (
	"context"
	"io"
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

func modifyResponse(r *http.Response) error {
	log.Println(" > RevProxy: got response .. ")
	return nil
}

func newLoggedReverseProxy(target *url.URL) *httputil.ReverseProxy {
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
		Director:       director(target),
		Transport:      transport,
		ModifyResponse: modifyResponse,
	}
}

func newReverseProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}

type revProxy struct {
	targetURL *url.URL
}

type SpyResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewSpyResponseWriter(w http.ResponseWriter) *SpyResponseWriter {
	return &SpyResponseWriter{w, http.StatusOK}
}

func (spy *SpyResponseWriter) WriteHeader(code int) {
	spy.statusCode = code
	spy.ResponseWriter.WriteHeader(code)
}

func (p *revProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// add context
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	req := r.WithContext(ctx)

	defer cancel()

	proxy := newReverseProxy(p.targetURL)
	spy := NewSpyResponseWriter(w)
	proxy.ServeHTTP(spy, req)
	if spy.statusCode == http.StatusBadGateway {
		w.WriteHeader(200)
		io.WriteString(w, "From Proxy: all good to go")
	}
}

func startReverseProxy(ctx context.Context, wg *sync.WaitGroup, addr, target string) {
	targetURL, _ := url.Parse(target)
	myProxy := &revProxy{targetURL: targetURL}
	startServer(ctx, wg, addr, myProxy)
}
