package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/unhello/", unHello)
	http.ListenAndServe("localhost:4000", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Makhmudov Azizbek1")
}

func unHello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "UnhHello, my name is Makhmudov Azizbek")
}
