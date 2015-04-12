package main

import "os/exec"

type Info struct {
	// map of humanreadable to ID
	Buttons  map[string]string
	Uname    string
	WhatTime string
	Whoami   string
	PS       string
}

func get(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	out, _ := cmd.Output()
	return string(out)
}

func getInfo() *Info {
	return &Info{
		Buttons:  buttons,
		Uname:    get("uname", "-a"),
		WhatTime: get("date"),
		Whoami:   get("whoami"),
		PS:       get("bash", "-c", "ps ux | grep vlc"),
	}
}
