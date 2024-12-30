package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ensureTemplated(editor string) string {
	if strings.Contains(editor, "%s") {
		return editor
	}
	return editor + " %s"
}

func isBinaryAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func detectEditorsOnPath() (string, error) {
	editors := []string{"vim", "vi", "emacs", "nano", "ed"}
	for _, editor := range editors {
		if isBinaryAvailable(editor) {
			return editor, nil
		}
	}
	return "", fmt.Errorf("no editor found, set either $EDITOR or $SCROLLBACK_EDITOR")
}

func getFallbackEditor() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor, err := detectEditorsOnPath()
		if err != nil {
			return "", err
		}
		return editor, nil
	}
	return ensureTemplated(editor), nil
}

// export SCROLLBACK_EDITOR='nvim +"term cat %s"'  +"execute ':normal! G'"
func GetEditorCommand() (string, error) {
	editor := os.Getenv("SCROLLBACK_EDITOR")
	if editor == "" {
		return getFallbackEditor()
	}
	return ensureTemplated(editor), nil
}