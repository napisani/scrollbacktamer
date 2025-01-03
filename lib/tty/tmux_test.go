package tty

import (
	"os"
	"testing"
)

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
