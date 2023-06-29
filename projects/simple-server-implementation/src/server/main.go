package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	
}