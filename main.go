package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("config/env/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func main() {
	PORT := os.Getenv("APP_PORT")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<h1>Hello you!</h1>
<p>This website will become my place to share my thought and project with the world!</p>`)
	})

	r.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<h1>Resume</h1>
<p>Here you will soon be able to se my awesome resume</p>`)
	})

	fmt.Printf("Starting server on port %s", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), r); err != nil {
		panic(fmt.Sprintf("Could not start http server: %v", err))
	}
}
