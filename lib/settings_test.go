package lib

import (
	"regexp"
	"testing"
)

func TestValidateSettings(t *testing.T) {
	tests := []struct {
		name     string
		settings Settings
		wantErr  bool
	}{
		{
			name: "valid lines settings",
			settings: Settings{
				Units: ScrollbackUnitLines,
				LastN: 10,
			},
			wantErr: false,
		},
		{
			name: "valid segments settings",
			settings: Settings{
				Units:                ScrollbackUnitSegments,
				ScrollbackTerminator: regexp.MustCompile(`^$`),
				LastN:               10,
			},
			wantErr: false,
		},
		{
			name: "invalid units",
			settings: Settings{
				Units: "invalid",
				LastN: 10,
			},
			wantErr: true,
		},
		{
			name: "segments without terminator",
			settings: Settings{
				Units: ScrollbackUnitSegments,
				LastN: 10,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSettings(&tt.settings)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSettings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
