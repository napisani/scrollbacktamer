package tty

import (
	"fmt"
	"io"
)

type TTY interface {
	IsInTTY() (bool, error)
	GetScrollbackStream() (io.Reader, error)
	GetName() string
}

var ttyImpls = []TTY{
	&TMux{},
}

func GetTTY(force_tty string) (TTY, error) {
	if force_tty != "" {
		for _, impl := range ttyImpls {
			if impl.GetName() == force_tty {
				return impl, nil
			}
		}
		return nil, fmt.Errorf("The specified TTY: '%s' is not supported", force_tty)
	}

	for _, impl := range ttyImpls {
		if ok, err := impl.IsInTTY(); ok {
			if err != nil {
				return nil, fmt.Errorf("failed to check if in tty: %w", err)
			}
			return impl, nil
		}
	}
	firstImpl := ttyImpls[0]
	return firstImpl, nil
	// return nil, fmt.Errorf("The current TTY could not be identified or is not supported")
}
