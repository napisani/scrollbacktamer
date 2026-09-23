package tty

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	go func() {
		if err := cmd.Wait(); err != nil {
			message := strings.TrimSpace(stderr.String())
			if message != "" {
				err = fmt.Errorf("tmux capture-pane: %s: %w", message, err)
			} else {
				err = fmt.Errorf("tmux capture-pane: %w", err)
			}
			_ = pw.CloseWithError(err)
			return
		}
		_ = pw.Close()
	}()

	return pr, nil
}

func (t *TMux) GetName() string {
	return "tmux"
}
