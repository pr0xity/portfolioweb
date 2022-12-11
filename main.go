package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello you!</h1>")
		fmt.Fprintf(w, "<p>This website will become my place to share my thought and project with the world!</p>")
	})

	r.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Resume</h1>")
		fmt.Fprintf(w, "<p>Here you will soon be able to se my awesome resume</p>")
	})

	http.ListenAndServe(":80", r)

}
