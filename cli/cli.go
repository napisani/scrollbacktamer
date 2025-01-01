package cli

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/napisani/scrollbacktamer/lib"
)

func parseUnits(units string) (lib.ScrollbackUnit, error) {
	switch units {
	case "lines":
		return lib.ScrollbackUnitLines, nil
	case "segments":
		return lib.ScrollbackUnitSegments, nil
	default:
		return "", fmt.Errorf("invalid units: %s", units)
	}
}

const PREFIX = "SBTAMER_"

func getDefaultValue[T any](key string, defaultValue T, toVal func(string) T) T {
	value := os.Getenv(PREFIX + key)
	if value == "" {
		return defaultValue
	}
	return toVal(value)
}
func ParseCLIArgs() (*lib.Settings, error) {
	settings := &lib.Settings{}
	var units string
	var terminator string

	// export SBTAMER_EDITOR='nvim +"term cat %s"  +"execute \":normal! G\""'
	flag.StringVar(&settings.Editor,
		"editor",
		getDefaultValue("EDITOR", "", func(v string) string { return v }),
		"Editor segment",
	)
	flag.IntVar(&settings.LastN, "last", getDefaultValue("LAST", -1, func(v string) int {
		i, _ := strconv.Atoi(v)
		return i
	}), "last N lines or segments")
	flag.StringVar(&units, "units", getDefaultValue("UNITS", string(lib.ScrollbackUnitLines), func(v string) string { return v }),
		"Scrollback units (lines or segments)",
	)
	flag.StringVar(&terminator, "terminator",
		getDefaultValue("TERMINATOR", "", func(v string) string { return v }),
		"Scrollback terminator string")

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
