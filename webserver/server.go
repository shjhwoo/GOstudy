package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	mux.HandleFunc("/bar",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bar")
	})
	return mux
}

func main(){
	http.ListenAndServe(":3000",MakeWebHandler())
}