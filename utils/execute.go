// B"H
/*
Package utils NEEDS MORE COMMENTS
*/
package utils

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
// I need to read and maybe refactor this function
import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
// I need the above to read and maybe refactor this function
func ExecuteCommand(executable string, command []string) {

	cmd := exec.Command(executable, command...)
	// cmd.Dir = executionPath

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		waitGroup.Done()
	}()
	_, errStderr = io.Copy(stderr, stderrIn)
	waitGroup.Wait()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	// outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

}
