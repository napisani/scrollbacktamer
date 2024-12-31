package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func GetTempFileName() (string, error) {
	epoch := time.Now().Unix()
	file := filepath.Join(os.TempDir(), fmt.Sprintf("sbt-%d.scrollback", epoch))
	return file, nil
}

func WriteStream(r *io.Reader, filename string, settings *Settings) error {
	switch settings.Units {
	case ScrollbackUnitLines:
		if settings.LastN > 0 {
			return writeLastNLines(r, filename, settings.LastN)
		}
		return writeEntireStream(r, filename)
	case ScrollbackUnitCommands:
		if settings.LastN > 0 {
			return writeLastNCommands(r, filename, settings.LastN, *settings.ScrollbackTerminator)
		}
		return writeEntireStream(r, filename)
	default:
		return fmt.Errorf("invalid units: %s", settings.Units)
	}
}

func writeLastNLines(r *io.Reader, filename string, n int) error {
	fmt.Println("writeLastNLines")
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(*r)
	lines := make([]string, 0, n)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		lines = append(lines, line)
		if len(lines) > n {
			lines = lines[1:]
		}
	}

	for _, line := range lines {
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeEntireStream(r *io.Reader, filename string) error {
	fmt.Println("writeEntireStream")
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the entire stream
	_, err = io.Copy(file, *r)
	if err != nil {
		return err
	}

	return nil
}

func writeLastNCommands(r *io.Reader, filename string, n int, reg regexp.Regexp) error {
  fmt.Println("writeLastNCommands")
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	cmdSegments := [][]string{}
	reader := bufio.NewReader(*r)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if reg.MatchString(line) {
			cmdSegments = append(cmdSegments, []string{line})
			if len(cmdSegments) > n {
				cmdSegments = cmdSegments[1:]
			}
		} else {
			if len(cmdSegments) == 0 {
				cmdSegments = append(cmdSegments, []string{})
			}
			lastIdx := len(cmdSegments) - 1
			lastSegment := cmdSegments[lastIdx]
			lastSegment = append(lastSegment, line)
		}
	}

	for _, segment := range cmdSegments {
		for _, line := range segment {
			_, err := file.WriteString(line)
			if err != nil {
				return err
			}
		}
	}
	return nil

}
