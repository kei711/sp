package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
)

func Executor(in string) {
	if in == "" {
		return
	}

	switch true {
	case in == cmdExit || in == cmdQuit:
		fmt.Println("Bye!")
		os.Exit(0)
		break
	case in == cmdClearCache:
		os.Remove(getCachePath())
		saveCommandXML()
		reloadSymfonyXML()
		break
	default:
		var commandName string
		inputCommands := strings.Split(in, " ")
		if len(inputCommands) > 1 {
			commandName = inputCommands[0]
		} else {
			commandName = in
		}
		commandName = strings.TrimSpace(commandName)

		if s.Contains(commandName) {
			execCommand(inputCommands)
		} else {
			fmt.Println("Command \"" + commandName + "\", Does not exist.")
		}

		break
	}
}

func execCommand(inputCommands []string) {
	_, _, exitCode, err := runCommand(inputCommands)
	if err != nil {
		log.Fatal(err)
	}
	if exitCode != 0 {
		fmt.Print("\r")
	}
}

func runCommand(commands []string) (stdout, stderr string, exitCode int, err error) {
	args := append([]string{selectedCommand}, commands...)
	args = append(args, "--ansi")
	cmd := exec.Command("php", args...)

	indicator := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	indicator.Start()

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	var bufout, buferr bytes.Buffer
	outReader2 := io.TeeReader(outReader, &bufout)
	errReader2 := io.TeeReader(errReader, &buferr)

	if err = cmd.Start(); err != nil {
		return
	}

	go printOutputWithHeader(outReader2)
	go printOutputWithHeader(errReader2)

	err = cmd.Wait()
	indicator.Stop()

	stdout = bufout.String()
	stderr = buferr.String()

	if err != nil {
		if err2, ok := err.(*exec.ExitError); ok {
			if s, ok := err2.Sys().(syscall.WaitStatus); ok {
				err = nil
				exitCode = s.ExitStatus()
			}
		}
	}

	return
}

func printOutputWithHeader(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("\r%s\n", scanner.Text())
	}
}
