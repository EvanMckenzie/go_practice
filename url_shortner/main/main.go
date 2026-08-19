package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	urlshortner "url-shortner"
)

func main() {
	mux := defaultMux()

	// build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	// set up yaml file flag
	file := flag.String("filename", "data.yaml", "YAML file containing paths")
	flag.Parse()

	yaml, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	// build the YAMLHandler using the mapHandler as the fallback
	dataHandler, err := urlshortner.DataHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dataHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
