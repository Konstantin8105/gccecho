package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const compiler string = "gcc"
const logFile string = "/tmp/gcc.log"

func main() {
	cmd := exec.Command(compiler, os.Args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", stderr.String())
		return
	}

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintln("ARG: ", os.Args)); err != nil {
		panic(err)
	}
	if _, err = f.Write(stdout.Bytes()); err != nil {
		panic(err)
	}
}
