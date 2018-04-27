package main

import (
	"fmt"
	"os"
	"os/exec"
)

const compiler string = "gcc"
const logFile string = "/tmp/gcc.log"

func main() {
	output, err := exec.Command(compiler, os.Args...).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	fmt.Println(string(output))

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintln("ARG: ", os.Args)); err != nil {
		panic(err)
	}
	if _, err = f.Write(output); err != nil {
		panic(err)
	}
}
