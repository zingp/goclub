package main

import (
	"net/http"
)

func main {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)

}