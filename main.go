package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gnolang/gnopls/cmd"
	slogmulti "github.com/samber/slog-multi"
)

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	logFilePath := filepath.Join(userHomeDir, "gnopls.log")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logger := slog.New(
		slogmulti.Fanout(
			slog.NewTextHandler(logFile, &slog.HandlerOptions{}),   // pass to first handler: logstash over tcp
			slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}), // then to second handler: stderr
		),
	)
	slog.SetDefault(logger)
	cmd.Execute()
}
