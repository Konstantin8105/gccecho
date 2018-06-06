package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const compiler string = "gcc"
const logFile string = "/tmp/gcc.log"

func main() {
	var cmd *exec.Cmd
	if len(os.Args) == 1 {
		cmd = exec.Command(compiler)
	} else {
		cmd = exec.Command(compiler, os.Args[1:]...)
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s", stderr.String())
		return
	}

	fmt.Println(stdout.String())

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(fmt.Sprintln("Dir: ", exPath)); err != nil {
		panic(err)
	}
	if _, err = f.WriteString(fmt.Sprintln("ARG: ", os.Args)); err != nil {
		panic(err)
	}
	if _, err = f.Write(stdout.Bytes()); err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}
}
