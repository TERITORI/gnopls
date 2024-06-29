package tools

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Build a Gno package: gno transpile -gobuild <dir>.
// TODO: Remove this in the favour of directly using tools/transpile.go
func Build(rootDir string) ([]byte, error) {
	args := []string{"transpile", "-skip-imports", "-gobuild", filepath.Join(rootDir)}
	srcPathsStr := os.Getenv("GNOSRCPATHS")
	if srcPathsStr != "" {
		args = append(args, "-extra-dirs", srcPathsStr)
	}
	slog.Info("will run", slog.String("cmd", strings.Join(append([]string{"gno"}, args...), " ")))
	return exec.Command("gno", args...).CombinedOutput()
}
