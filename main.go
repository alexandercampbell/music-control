package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"path"
	"strings"
)

var indexTemplate, _ = template.ParseFiles("index.template")

func indexPage(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, getInfo())
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("")
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/index.html", indexPage)
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	log.Println("------------------------------------")
	log.Println("          MUSIC CONTROLLER          ")
	log.Println("------------------------------------")

	for _, buttonID := range buttons {
		// for each button script, we set up an endpoint at
		// /triggers/<buttonID> that will launch the script
		httpPath := "/triggers/" + buttonID
		log.Printf("Handling %-20s @ %s", buttonID, httpPath)

		// required to make the anonymous closure work
		buttonID := buttonID

		http.HandleFunc(httpPath, func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Recieved request to launch %q from %s",
				buttonID, r.RemoteAddr)

			// execute the script corresponding to the buttonID
			scriptPath := path.Join(scriptsDirectory, buttonID)
			outBytes, err := exec.Command("bash", scriptPath).Output()
			if err != nil {
				log.Printf("Error executing %q: %q", scriptPath, err)
			}

			out := string(outBytes)
			out = strings.Replace(out, "\n", "\n\t", -1)
			log.Printf("Output:\n%q", out)
		})
	}

	log.Println()
	log.Println("Music Controller is running on :8000")

	http.ListenAndServe(":8000", nil)
}
