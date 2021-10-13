package common

import (
	"io/ioutil"
	"log"
	"os/exec"
)

type Cmd struct {
	path string
}

func NewCommand(path string) *Cmd {
	return &Cmd{path: path}
}

func (c *Cmd) Run(program string, args ...string) string {

	cmd := exec.Command(program, args...)
	cmd.Dir = c.path

	stderr, errStderr := cmd.StderrPipe()

	if errStderr != nil {
		log.Fatalln("error occurred at: " + errStderr.Error())
	}

	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		log.Fatalln("error occurred at: " + err.Error())
	}

	stderrOut, err := ioutil.ReadAll(stderr)

	if err != nil {
		log.Fatalln("error occurred at: " + err.Error())

	}

	out, _ := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if len(out) > 0 {
		return string(out)
	} else {
		return string(stderrOut)
	}

}
