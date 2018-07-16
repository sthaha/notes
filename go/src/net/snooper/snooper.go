package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"time"
)

func forward(serviceURL string) http.HandlerFunc {
	u, err := url.Parse(serviceURL)

	if err != nil {
		log.Fatalf("Failed to parse: %q - error: %s \n", serviceURL, err)
		// NOTE: fatal exits so the return is added only for readability
		return nil
	}
	proxy := httputil.NewSingleHostReverseProxy(u)

	return proxy.ServeHTTP // (w, r)

}

func snoop(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req, _ := httputil.DumpRequest(r, true)

		log.Printf(`
-----------------------------------------------------------
>>> %s
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`, string(req))

		// switch out response writer for a recorder for all subsequent handlers
		rec := httptest.NewRecorder()
		// ensure it isn't gzipped so that the response can just be dumped
		r.Header.Set("Accept-Encoding", "deflate")

		next(rec, r)
		res, _ := httputil.DumpResponse(rec.Result(), true)

		log.Printf(`
>>> %s
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
%s
===========================================================`,
			r.URL.String(), string(res))

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
	http.HandleFunc("/api/tenants/", snoop(forward("http://localhost:9222")))
	http.HandleFunc("/", snoop(wait(60*time.Second)))
	log.Fatal(http.ListenAndServe(":9002", nil))
}
