package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//handler functions
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Salamaleikum"))
}

func showCatalog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("test"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a catolog with test: %v", id)
}

func createCatalog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created products with various characteristics"))
}

func main() {
	//router || multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/catalog", showCatalog)
	mux.HandleFunc("/catalog/create/", createCatalog)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
