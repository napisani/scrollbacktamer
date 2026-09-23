package tty

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTMuxCaptureFailureIsReturned(t *testing.T) {
	dir := t.TempDir()
	tmuxPath := filepath.Join(dir, "tmux")
	if err := os.WriteFile(tmuxPath, []byte("#!/bin/sh\necho capture failed >&2\nexit 1\n"), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)

	reader, err := (&TMux{}).GetScrollbackStream()
	if err != nil {
		t.Fatalf("GetScrollbackStream() start error = %v", err)
	}
	contents, err := io.ReadAll(reader)
	if err == nil || !strings.Contains(err.Error(), "capture failed") {
		t.Fatalf("ReadAll() contents = %q, error = %v; want capture failure", contents, err)
	}
	if strings.Contains(string(contents), "capture failed") {
		t.Fatalf("tmux stderr leaked into scrollback contents: %q", contents)
	}
}

func TestTMuxIsInTTY(t *testing.T) {
	tmux := &TMux{}

	tests := []struct {
		name    string
		envVars map[string]string
		want    bool
		wantErr bool
	}{
		{
			name: "in tmux",
			envVars: map[string]string{
				"TMUX":      "/tmp/tmux-1000/default,1234,0",
				"TMUX_PANE": "%1",
			},
			want:    true,
			wantErr: false,
		},
		{
			name:    "not in tmux",
			envVars: map[string]string{},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for k, v := range tt.envVars {
				os.Setenv(k, v)
				defer os.Unsetenv(k)
			}

			got, err := tmux.IsInTTY()
			if (err != nil) != tt.wantErr {
				t.Errorf("TMux.IsInTTY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TMux.IsInTTY() = %v, want %v", got, tt.want)
			}
		})
	}
}
