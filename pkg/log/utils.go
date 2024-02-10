package log

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

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
	source := a.Value.Any().(*slog.Source)

	pwd, err := os.Getwd()
	if err != nil {
		return a
	}

	relPath, err := filepath.Rel(pwd, source.File)
	if err != nil {
		return a
	}

	formattedPath := fmt.Sprintf("%s:%d", relPath, source.Line)

	return slog.Attr{
		Key:   a.Key,
		Value: slog.StringValue(formattedPath),
	}
}
