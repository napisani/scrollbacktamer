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
	fmt.Println(os.Args)
	settings, err := cli.ParseCLIArgs()
	if err != nil {
		panic(err)
	}
	tty := lib.GetTTY()
	fmt.Println(tty)
	reader, err := tty.GetScrollbackStream()
	if err != nil {
		panic(err)
	}

	editorCmd, err := lib.GetEditorCommand()
	if err != nil {
		panic(err)
	}

	fileName, err := lib.GetTempFileName()
	if err != nil {
		panic(err)
	}

	fmt.Println(settings)

	lib.WriteStream(&reader, fileName, settings)

	defer os.Remove(fileName)
	cmd := fmt.Sprintf(editorCmd, fileName)
	fmt.Println(cmd)
	err = runScrollbackEditCmd(cmd)
	if err != nil {
		panic(err)
	}

}
