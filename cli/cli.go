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

	flag.StringVar(&settings.Editor,
		"editor",
		getDefaultValue("EDITOR", "", func(v string) string { return v }),
		"Editor segment",
	)
	flag.StringVar(&settings.File,
		"file",
		getDefaultValue("FILE", "", func(v string) string { return v }),
		"File with scrollback content",
	)

	flag.StringVar(&settings.TTY,
		"tty",
		getDefaultValue("TTY", "", func(v string) string { return v }),
		"manually define the TTY being used instead of auto-detecting")

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
	flag.BoolVar(&settings.Verbose, "v", getDefaultValue("VERBOSE", false, func(v string) bool {
		b, _ := strconv.ParseBool(v)
		return b
	}), "Verbose output")

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
