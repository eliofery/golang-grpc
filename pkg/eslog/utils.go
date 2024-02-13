// Package eslog ...
// nolint
package eslog

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const timestampFormat = "2006-01-02 15:04:05.999999999 -0700 -07"

// replaceLevel ...
func replaceLevel(a slog.Attr) slog.Attr {
	l := a.Value.Any().(slog.Level)

	levelLabel, ok := LevelNames[l]
	if !ok {
		levelLabel = l.String()
	}

	a.Value = slog.StringValue(levelLabel)

	return a
}

// replaceSource путь до файла в котором был вызван лог
// Реализация с использованием ASCII
// absPath, err := filepath.Abs(source.File)
// formattedPath := fmt.Sprintf("\x1b]8;;file://%v\x1b\\%s:%v\x1b]8;;\x1b\\", absPath, relPath, source.Line)
func replaceSource(a slog.Attr) slog.Attr {
	source := a.Value.Any().(slog.Source)

	pwd, err := os.Getwd()
	if err != nil {
		return a
	}

	relPath, err := filepath.Rel(pwd, source.File)
	if err != nil {
		return a
	}

	basePath := filepath.Base(relPath)

	formattedPath := fmt.Sprintf("%s:%d", basePath, source.Line)

	return slog.Attr{
		Key:   a.Key,
		Value: slog.StringValue(formattedPath),
	}
}

// replaceTime ...
func replaceTime(a slog.Attr) slog.Attr {
	t, err := time.Parse(timestampFormat, a.Value.String())
	if err != nil {
		return a
	}

	formattedTime := t.Format(time.DateTime)
	a.Value = slog.StringValue(formattedTime)

	return a
}

// stackFormat ...
func stackFormat(stack string) []string {
	stack = strings.ReplaceAll(stack, "\t", " ")
	s := strings.Split(stack, "\n")

	if s[len(s)-1] == "" {
		s = s[:len(s)-1]
	}

	return s
}
