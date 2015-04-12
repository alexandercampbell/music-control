package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func readButtonsFromDirectory(directoryPath string) map[string]string {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatalf("Error trying to read %q: %s", directoryPath, err)
	}

	buttons := map[string]string{}
	for _, file := range files {
		properName := file.Name()
		if strings.HasSuffix(properName, ".sh") {
			properName = properName[:len(properName)-3]
		}
		properName = strings.Replace(properName, "_", " ", -1)
		properName = strings.Title(properName)

		buttons[properName] = file.Name()
	}
	return buttons
}

const scriptsDirectory = "./scripts/"

var buttons = readButtonsFromDirectory(scriptsDirectory)
