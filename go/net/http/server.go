package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func good(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "bad request")
}

func goodEnc(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusBadRequest)
	foo := &struct {
		Msg string `json:"msg"`
	}{"bad request from foo"}
	json.NewEncoder(w).Encode(foo)

}

func encFailOk(w http.ResponseWriter, r *http.Request) {
	type Foo struct {
		name string
		num  json.Number
	}
	foo := Foo{"john doe", json.Number(`invalid`)}
	err := json.NewEncoder(w).Encode(foo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad request")
	}
}

func wrongOrderWH(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bad request fails so does content type") // WRONG order
	w.WriteHeader(http.StatusBadRequest)                        // writes 200 OK instead of 400
	w.Header().Set("Content-Type", "application/json")
}

func wrongOrderEncode(w http.ResponseWriter, r *http.Request) {
	foo := &struct {
		Msg string `json:"msg"`
	}{"bad request from foo"}
	json.NewEncoder(w).Encode(foo)

	// writes 200 OK instead of 400
	w.WriteHeader(http.StatusBadRequest) // WRONG ORDER

}

func main() {
	http.HandleFunc("/good/ok", good)
	http.HandleFunc("/good/enc", goodEnc)
	http.HandleFunc("/good/enc-fail", encFailOk)
	http.HandleFunc("/bad/wh", wrongOrderWH)
	http.HandleFunc("/bad/enc", wrongOrderEncode)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
