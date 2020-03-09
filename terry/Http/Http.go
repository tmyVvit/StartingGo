package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello %s", request.FormValue("name"))
	})
	_ = http.ListenAndServe(":8080", nil)
}