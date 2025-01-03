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
