package lib

import (
	"os"
	"testing"
)

func TestEnsureTemplated(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "already templated",
			input:    "vim %s",
			expected: "vim %s",
		},
		{
			name:     "not templated",
			input:    "vim",
			expected: "vim %s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ensureTemplated(tt.input)
			if result != tt.expected {
				t.Errorf("ensureTemplated() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetEditorCommand(t *testing.T) {
	tests := []struct {
		name     string
		settings Settings
		envVar   string
		wantErr  bool
	}{
		{
			name: "editor in settings",
			settings: Settings{
				Editor: "code",
			},
			wantErr: false,
		},
		{
			name:     "editor from env",
			settings: Settings{},
			envVar:   "vim",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envVar != "" {
				os.Setenv("EDITOR", tt.envVar)
				defer os.Unsetenv("EDITOR")
			}

			_, err := GetEditorCommand(&tt.settings)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEditorCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
