package examplelib

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func GetEgid() int {
	return os.Getegid()
}

func RunCommand(filename string, args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		panic(err)
	}
}
