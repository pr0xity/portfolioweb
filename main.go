package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// const PORT = os.Getenv("APP_PORT") // This is how you would do it in production
// const PORT = "8091"

func init() {
	if err := godotenv.Load("config/env/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<h1>Hello you!</h1>
<p>This website will become my place to share my thought and project with the world!</p>`)
	})

	r.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<h1>Resume</h1>
<p>Here you will soon be able to se my awesome resume</p>`)
	})

	fmt.Printf("Starting server on port %s", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r); err != nil {
		panic(fmt.Sprintf("Could not start http server: %v", err))
	}
}
