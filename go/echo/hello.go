package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on 8080.")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
