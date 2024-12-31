package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Units string

const (
	UnitsLines    Units = "lines"
	UnitsCommands Units = "commands"
)

type Settings struct {
	ScrollbackTerminator string
	Units                Units
	LastN                int
}

func GetTempFileName() (string, error) {
	epoch := time.Now().Unix()
	file := filepath.Join(os.TempDir(), fmt.Sprintf("sbt-%d.scrollback", epoch))
	return file, nil
}

func WriteStream(r *io.Reader, filename string, settings *Settings) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the contents of the reader to the file
	_, err = io.Copy(file, *r)
	return err
}
