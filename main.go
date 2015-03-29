package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

var (
	indexTemplate, _ = template.ParseFiles("index.template")
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, getInfo())
}

func startBBC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request to launch BBC Radio from", r.RemoteAddr)
	exec.Command("bash", "scripts/bbc.sh").Start()
}

func startPolice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request to launch police scanner from", r.RemoteAddr)
	exec.Command("bash", "scripts/police.sh").Start()
}

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/index.html", indexPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/start-bbc", startBBC)
	http.HandleFunc("/start-police", startPolice)

	fmt.Println(get("ifconfig"))
	fmt.Println("------------------------------------")
	fmt.Println("Music Controller is running on :8000")
	fmt.Println("------------------------------------")

	http.ListenAndServe(":8000", nil)
}
