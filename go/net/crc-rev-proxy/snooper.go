package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strings"
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

// NewSingleHostReverseProxy returns a new ReverseProxy that routes
// URLs to the scheme, host, and base path provided in target. If the
// target's path is "/base" and the incoming request was for "/dir",
// the target request will be for /base/dir.
// NewSingleHostReverseProxy does not rewrite the Host header.
// To rewrite Host headers, use ReverseProxy directly with a custom
// Director policy.
func newHostRewriterProxy(target *url.URL) *httputil.ReverseProxy {
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
		req.Header.Set("Host", req.URL.Host)

		reqStr, _ := httputil.DumpRequest(req, true)

		log.Printf(`
	Director
	-----------------------
	
	%s
	----------------------
	`, reqStr)

	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &httputil.ReverseProxy{Director: director, Transport: tr}
}

func forward(serviceURL string, retry int) http.HandlerFunc {
	u, err := url.Parse(serviceURL)
	if err != nil {
		log.Fatalf("Failed to parse: %q - error: %s \n", serviceURL, err)
		// NOTE: fatal exits so the return is added only for readability
		return nil
	}

	log.Printf("Creating proxy for %s\n", u.String())

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(">>> [forward]: Host: %q | forwarding %q", r.Header["Host"], r.URL.String())

		req, _ := httputil.DumpRequest(r, true)

		log.Printf(`
-----------------------------------------------------------
Forward: Request %s
...............................................
%s

-----------------------------------------------------------
`, r.URL, string(req))
		target, _ := url.Parse(serviceURL)
		target.Host = r.Host
		service := newHostRewriterProxy(target)
		service.ServeHTTP(w, r)
	}

}

func snoop(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req, _ := httputil.DumpRequest(r, true)

		log.Printf(`
-----------------------------------------------------------
Snoop: Request
...............................................
%s
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`, string(req))

		// switch out response writer for a recorder for all subsequent handlers
		rec := httptest.NewRecorder()
		// ensure it isn't gzipped so that the response can just be dumped
		//r.Header.Set("Accept-Encoding", "deflate")

		next(rec, r)
		res, _ := httputil.DumpResponse(rec.Result(), false)

		log.Printf(`
...............................................
Snoop Response:
...............................................
%s
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
%s
===========================================================`,
			r.URL.String(), res)

		// copy everything from response recorder
		// to actual response writer
		for k, v := range rec.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		rec.Body.WriteTo(w)

	}
}

func wait(duration time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(">>> [wait]: request %q  waiting %v", r.URL.String(), duration)
		select {
		case <-r.Context().Done():
			log.Printf(">>> [wait]: request %q got cancelled", r.URL.String())
		case <-time.After(duration):
			log.Printf(">>> [wait]: request %q waiting DONE", r.URL.String())
		}
	}
}

func main() {
	http.HandleFunc("/",
		snoop(
			//forward("https://192.168.130.11/", 3),
			//oauth-openshift.apps-crc.testing
			forward("https://oauth-openshift.apps-crc.testing/", 3),
			//forward("https://console-openshift-console.apps-crc.testing/", 3),
		),
	)
	// http.HandleFunc("/", snoop(wait(60*time.Second)))
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
	//log.Fatal(http.ListenAndServe(":9002", nil))
}
