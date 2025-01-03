package tty

import (
	"fmt"
	"io"
)

type TTY interface {
	IsInTTY() (bool, error)
	GetScrollbackStream() (io.Reader, error)
}

var ttyImpls = []TTY{
	&TMux{},
}

func GetTTY() (TTY, error) {
	for _, impl := range ttyImpls {
		if ok, err := impl.IsInTTY(); ok {
			if err != nil {
				return nil, fmt.Errorf("failed to check if in tty: %w", err)
			}
			return impl, nil
		}
	}
	return nil, fmt.Errorf("The current TTY could not be identified or is not supported")
}
