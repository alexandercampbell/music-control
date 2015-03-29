package main

import (
	"fmt"
	"net/http"
)

func launchMusic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GOOD MORNING")
}

func main() {
	http.HandleFunc("/", launchMusic)
	http.ListenAndServe(":8000", nil)
}
