package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type SenzMsg struct {
	Uid string
	Msg string
}

func main() {
	http.HandleFunc("/promize", promize)
	address := ":7070"
	println("starting server on address" + address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		println(err.Error)
		os.Exit(1)
	}
}

func promize(w http.ResponseWriter, r *http.Request) {
	// read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	println(string(b))

	// read headers
	rIp := r.Header.Get("X-Real-IP")
	fIp := r.Header.Get("X-Forwarded-For")
	verified := r.Header.Get("VERIFIED")
	dn := r.Header.Get("DN")

	println(rIp)
	println(fIp)
	println(verified)
	println(dn)
  
  	// serd response
  	j := `{"status": "CREATED"}`
  	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(j))
	
  	return
}