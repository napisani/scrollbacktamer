package tty

import (
	"io"
	"os"
	"os/exec"
)

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

	pr, pw := io.Pipe()

	cmd.Stdout = pw
	cmd.Stderr = pw

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	go func() {
		cmd.Wait()
		pw.Close()
	}()

	return pr, nil
}

func (t *TMux) GetName() string {
  return "tmux"
}
