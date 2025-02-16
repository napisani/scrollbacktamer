package lib

import (
	"fmt"
	"regexp"
)

type ScrollbackUnit string

const (
	ScrollbackUnitLines    ScrollbackUnit = "lines"
	ScrollbackUnitSegments ScrollbackUnit = "segments"
)

type Settings struct {
	ScrollbackTerminator *regexp.Regexp
	Units                ScrollbackUnit
	LastN                int
	Editor               string
	File                 string
	Verbose              bool
	TTY                  string
}

func (s *Settings) String() string {
	return fmt.Sprintf("Settings{ScrollbackTerminator: %v, Units: %v, LastN: %v, Editor: %v, File: %v, Verbose: %v}", s.ScrollbackTerminator, s.Units, s.LastN, s.Editor, s.File, s.Verbose)
}

func ValidateSettings(settings *Settings) error {
	if settings.Units != ScrollbackUnitLines && settings.Units != ScrollbackUnitSegments {
		return fmt.Errorf("invalid units: %s", settings.Units)
	}

	if settings.Units == ScrollbackUnitSegments && settings.ScrollbackTerminator == nil {
		return fmt.Errorf("a terminator is required for segment-based scrollback editing")
	}

	return nil
}
