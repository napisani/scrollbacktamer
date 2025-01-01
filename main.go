package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/shlex"
	"github.com/napisani/scrollbacktamer/cli"
	"github.com/napisani/scrollbacktamer/lib"
)

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

	// Find the full path of the command
	path, err := exec.LookPath(cmdOnly)
	if err != nil {
		return err
	}
	return syscall.Exec(path, args, os.Environ())
}

func main() {
	settings, err := cli.ParseCLIArgs()
	if err != nil {
		panic(fmt.Errorf("failed to parse CLI args: %w", err))
	}
	tty, err := lib.GetTTY()
	if err != nil {
		panic(fmt.Errorf("failed to get tty: %w", err))
	}
	reader, err := tty.GetScrollbackStream()
	if err != nil {
		panic(fmt.Errorf("failed to get scrollback stream: %w", err))

	}

	editorCmd, err := lib.GetEditorCommand(settings)
	if err != nil {
		panic(fmt.Errorf("failed to get the editor to use for scrollback editing: %w", err))
	}

	fileName, err := lib.GetTempFileName()
	if err != nil {
		panic(fmt.Errorf("failed to get a temporary file name: %w", err))
	}

	lib.WriteStream(&reader, fileName, settings)

	defer os.Remove(fileName)
	cmd := fmt.Sprintf(editorCmd, fileName)
	err = runScrollbackEditCmd(cmd)
	if err != nil {
		panic(fmt.Errorf("failed to run the editor: %w", err))
	}

}
