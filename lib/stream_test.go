package lib

import (
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestWriteLastNLines(t *testing.T) {
	input := strings.NewReader("line1\nline2\nline3\nline4\n")
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	reader := io.Reader(input)
	err = writeLastNLines(&reader, tmpfile, 2)
	if err != nil {
		t.Errorf("writeLastNLines() error = %v", err)
	}

	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := "line3\nline4\n"
	if string(content) != expected {
		t.Errorf("writeLastNLines() content = %v, want %v", string(content), expected)
	}
}

func TestWriteLastNSegments(t *testing.T) {
	input := strings.NewReader("cmd1\noutput1\n\ncmd2\noutput2\n\ncmd3\noutput3\n")
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	reader := io.Reader(input)
	terminator := regexp.MustCompile(`^cmd`)
	err = writeLastNSegments(&reader, tmpfile, 2, *terminator)
	if err != nil {
		t.Errorf("writeLastNSegments() error = %v", err)
	}

	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := "cmd2\noutput2\ncmd3\noutput3\n"
	if string(content) != expected {
		t.Errorf("writeLastNSegments() content = %v, want %v", string(content), expected)
	}
}
