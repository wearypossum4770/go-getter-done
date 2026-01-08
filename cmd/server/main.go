package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go Task Manager!")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8275"
	}
	fmt.Printf("Server starting on port %s (T9: task)\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%q", port), nil))
}
