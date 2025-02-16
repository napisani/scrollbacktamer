package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/shlex"
	"github.com/napisani/scrollbacktamer/cli"
	"github.com/napisani/scrollbacktamer/lib"
	"github.com/napisani/scrollbacktamer/lib/tty"
)

func setupLogging(verbose bool) {
	if verbose {
		logFile, err := os.OpenFile("/tmp/sbtamer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		}
		log.SetOutput(logFile)
		log.Println("Logging started")
	} else {
		log.SetOutput(io.Discard)
	}
}

func runScrollbackEditCmd(cmd string) error {
	parts, err := shlex.Split(cmd)
	if err != nil {
		return fmt.Errorf("failed to parse command: %w", err)
	}
	if len(parts) == 0 {
		return fmt.Errorf("empty command")
	}
	cmdOnly := parts[0]
	args := parts

	path, err := exec.LookPath(cmdOnly)
	if err != nil {
		return err
	}
	return syscall.Exec(path, args, os.Environ())
}

func getReader(settings *lib.Settings) (io.Reader, error) {
	// if a file was passed just use the file without detecting the TTY
	if settings.File != "" {
		return os.Open(settings.File)
	}

	tty, err := tty.GetTTY(settings.TTY)
	if err != nil {
		return nil, fmt.Errorf("failed to get tty: %w", err)
	}
	return tty.GetScrollbackStream()
}

func main() {

	settings, err := cli.ParseCLIArgs()
	if err != nil {
		panic(fmt.Errorf("failed to parse CLI args: %w", err))
	}
	setupLogging(settings.Verbose)
	log.Println("Settings: ", settings)

	log.Println("Getting reader for scrollback")
	reader, err := getReader(settings)
	if err != nil {
		log.Fatalf("Failed to get reader: %v", err)
		panic(fmt.Errorf("failed to get reader: %w", err))
	}

	log.Println("Getting editor command")
	editorCmd, err := lib.GetEditorCommand(settings)
	if err != nil {
		log.Fatalf("Failed to get editor command: %v", err)
		panic(fmt.Errorf("failed to get the editor to use for scrollback editing: %w", err))
	}

	log.Println("Generating temporary file name")
	fileName, err := lib.GetTempFileName()
	if err != nil {
		log.Fatalf("Failed to get temporary file name: %v", err)
		panic(fmt.Errorf("failed to get a temporary file name: %w", err))
	}

	log.Printf("Writing stream to temporary file: %s", fileName)
	err = lib.WriteStream(&reader, fileName, settings)
	if err != nil {
		log.Fatalf("Failed to write stream: %v", err)
	}

	defer os.Remove(fileName)
	cmd := fmt.Sprintf(editorCmd, fileName)
	log.Printf("Running editor command: %s", cmd)
	err = runScrollbackEditCmd(cmd)
	if err != nil {
		log.Fatalf("Failed to run editor: %v", err)
		panic(fmt.Errorf("failed to run the editor: %w", err))
	}

}
