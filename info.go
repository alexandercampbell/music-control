package main

import "os/exec"

type Info struct {
	Uname  string
	Whoami string
	PS     string
}

func get(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	out, _ := cmd.Output()
	return string(out)
}

func getInfo() *Info {
	return &Info{
		Uname:  get("uname", "-a"),
		Whoami: get("whoami"),
		PS:     get("ps", "u"),
	}
}
