package cli

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/napisani/scrollbacktamer/lib"
)

func parseUnits(units string) (lib.ScrollbackUnit, error) {
	switch units {
	case "lines":
		return lib.ScrollbackUnitLines, nil
	case "commands":
		return lib.ScrollbackUnitCommands, nil
	default:
		return "", fmt.Errorf("invalid units: %s", units)
	}
}

func ParseCLIArgs() (*lib.Settings, error) {
	settings := &lib.Settings{}
	var units string
	var terminator string

	flag.StringVar(&settings.Editor, "editor", "", "Editor command")
	flag.IntVar(&settings.LastN, "last", -1, "last N lines or commands")
	flag.StringVar(&units, "units", string(lib.ScrollbackUnitLines), "Scrollback units (lines or commands)")
	flag.StringVar(&terminator, "terminator", "exit", "Scrollback terminator string")

	flag.Parse()

	parsedUnit, err := parseUnits(units)
	if err != nil {
		return nil, err
	}
	settings.Units = parsedUnit

	if terminator != "" {
		settings.ScrollbackTerminator = regexp.MustCompile(terminator)
	}

	return settings, nil
}
