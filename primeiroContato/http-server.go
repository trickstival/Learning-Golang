package main

import (
	"fmt"
	"net/http"
)

func serve() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/hello", handlerH)

	http.ListenAndServe(":8080", nil)
}

func handler(writter http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writter, "Handler called\n")
}

func handlerH(writter http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writter, "Handler 2 called\n")
}
