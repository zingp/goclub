package main

import (
	//"io"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte ("Hello world!"))
}

func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":8080", nil)
} 
