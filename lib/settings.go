package lib

import (
	"regexp"
)

type ScrollbackUnit string

const (
	ScrollbackUnitLines    ScrollbackUnit = "lines"
	ScrollbackUnitCommands ScrollbackUnit = "commands"
)

type Settings struct {
	ScrollbackTerminator *regexp.Regexp
	Units                ScrollbackUnit
	LastN                int
	Editor               string
}
