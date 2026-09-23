package tty

import "testing"

func TestGetTTYReturnsErrorOutsideSupportedTTY(t *testing.T) {
	t.Setenv("TMUX", "")
	t.Setenv("TMUX_PANE", "")

	got, err := GetTTY("")
	if err == nil {
		t.Fatalf("GetTTY() = %#v, want unsupported TTY error", got)
	}
}
