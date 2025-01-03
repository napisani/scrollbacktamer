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
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}

	defer file.Close()

	switch settings.Units {
	case ScrollbackUnitLines:
		if settings.LastN > 0 {
			return writeLastNLines(r, file, settings.LastN)
		}
		return writeAllLines(r, file)
	case ScrollbackUnitSegments:
		if settings.LastN > 0 {
			return writeLastNSegments(r, file, settings.LastN, *settings.ScrollbackTerminator)
		}
		return writeAllLines(r, file)
	default:
		return fmt.Errorf("invalid units: %s", settings.Units)
	}

}

func writeLastNLines(r *io.Reader, w *os.File, n int) error {
	reader := bufio.NewReader(*r)
	lines := make([]string, 0, n)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read line: %w", err)
		}
		lines = append(lines, line)
		if len(lines) > n {
			lines = lines[1:]
		}
	}

	for _, line := range lines {
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeAllLines(r *io.Reader, w *os.File) error {
	_, err := io.Copy(w, *r)
	return err
}

func writeLastNSegments(r *io.Reader, w *os.File, n int, reg regexp.Regexp) error {
	cmdSegments := [][]string{}
	reader := bufio.NewReader(*r)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read line: %w", err)
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
			cmdSegments[lastIdx] = append(cmdSegments[lastIdx], line)
		}
	}

	for _, segment := range cmdSegments {
		for _, line := range segment {
			_, err := w.WriteString(line)
			if err != nil {
				return fmt.Errorf("failed to write line: %w", err)
			}
		}
	}
	return nil

}
