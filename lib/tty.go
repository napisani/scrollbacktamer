package lib

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type ScrollbackReadSettings struct {
	ANSIColors bool
	LineLimit  int
}

type TTY interface {
	IsInTTY() (bool, error)
	GetScrollbackStream() (io.Reader, error)
}

type TMux struct {
}

func (t *TMux) IsInTTY() (bool, error) {
	// get env variable TMUX
	tmuxEnv := os.Getenv("TMUX")
	tmuxPaneEnv := os.Getenv("TMUX_PANE")
	if tmuxEnv == "" || tmuxPaneEnv == "" {
		return false, nil
	}
	return true, nil
}

func (t *TMux) GetScrollbackStream() (io.Reader, error) {
	cmd := exec.Command("tmux", "capture-pane", "-p", "-S", "-", "-e")

	// Create a pipe to capture both stdout and stderr
	pr, pw := io.Pipe()

	// Set both stdout and stderr to the pipe writer
	cmd.Stdout = pw
	cmd.Stderr = pw

	// Start the command
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	// Close the pipe writer in a separate goroutine when the command finishes
	go func() {
		cmd.Wait()
		pw.Close()
	}()

	// Return the pipe reader
	return pr, nil
}

var ttyImpls = []TTY{
	&TMux{},
}

func GetTTY() (TTY, error) {
	for _, impl := range ttyImpls {
		if ok, err := impl.IsInTTY(); ok {
			if err != nil {
				return nil, fmt.Errorf("failed to check if in tty: %w", err)
			}
			return impl, nil
		}
	}
	return nil, fmt.Errorf("The current TTY could not be identified or is not supported")
}
