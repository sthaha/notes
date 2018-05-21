package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type revProxy struct {
	target *url.URL
	server *httputil.ReverseProxy
}

func newReverseProxy(target string) *revProxy {
	url, _ := url.Parse(target)
	return &revProxy{
		target: url,
		server: httputil.NewSingleHostReverseProxy(url),
	}
}

func (p *revProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reply := fmt.Sprintf("Reverse proxy: redirecting %q to: %q \n", r.URL, p.target)
	io.WriteString(w, reply)
	p.handle(w, r)
}

func (p *revProxy) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")
	p.server.Transport = &myTransport{}
	p.server.ServeHTTP(w, r)
}

type myTransport struct {
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	buf, _ := ioutil.ReadAll(req.Body)

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

	log.Printf(`>>> Request body
	%s
	>>> -----------------------------
	`, rdr1)

	req.Body = rdr2

	response, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Println("\n\ncame in error resp here", err)
		return nil, err //Server is not reachable. Server not working
	}

	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Println("\n\nerror in dumb response")
		// copying the response body did not work
		return nil, err
	}

	log.Printf(`>>> Response Body :
	%s
	----------------------
	`, string(body))
	return response, err
}

func startReverseProxy(ctx context.Context, wg *sync.WaitGroup, addr, target string) {
	proxy := newReverseProxy(target)
	startServer(ctx, wg, addr, proxy)
}
