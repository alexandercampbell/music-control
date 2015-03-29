package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	indexTemplate, _ = template.ParseFiles("index.template")
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, info)
}

func main() {
	http.HandleFunc("/", indexPage)
	fmt.Println("------------------------------------")
	fmt.Println("Music Controller is running on :8000")
	fmt.Println("------------------------------------")
	fmt.Println(get("ifconfig"))
	http.ListenAndServe(":8000", nil)
}
