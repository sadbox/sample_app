package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!<br>Just testing!")
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
