package tools

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Transpile a Gno package: gno transpile <dir>.
func Transpile(rootDir string) ([]byte, error) {
	args := []string{"transpile", "-skip-imports", filepath.Join(rootDir)}
	srcPathsStr := os.Getenv("GNOSRCPATHS")
	if srcPathsStr != "" {
		args = append(args, "-extra-dirs", srcPathsStr)
	}
	return exec.Command("gno", args...).CombinedOutput()
}
