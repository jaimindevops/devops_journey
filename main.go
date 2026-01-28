package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the "Name" of the machine (Pod) running this code
	hostname, _ := os.Hostname()
	
	// Print it to the user
	fmt.Fprintf(w, "Hello Network Engineer! I am running on Pod: %s\n", hostname)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
