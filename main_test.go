package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunScrollbackEditCmdReturnsAfterEditorExits(t *testing.T) {
	if os.Getenv("SCROLLBACKTAMER_EDITOR_HELPER") == "1" {
		marker := os.Getenv("SCROLLBACKTAMER_EDITOR_MARKER")
		if err := runScrollbackEditCmd("true"); err != nil {
			os.Exit(2)
		}
		if err := os.WriteFile(marker, []byte("returned"), 0600); err != nil {
			os.Exit(3)
		}
		return
	}

	marker := filepath.Join(t.TempDir(), "returned")
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunScrollbackEditCmdReturnsAfterEditorExits$")
	cmd.Env = append(os.Environ(),
		"SCROLLBACKTAMER_EDITOR_HELPER=1",
		"SCROLLBACKTAMER_EDITOR_MARKER="+marker,
	)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("helper failed: %v\n%s", err, output)
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatalf("editor command replaced the process instead of returning: %v", err)
	}
}
