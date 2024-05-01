package tui

import (
	"strings"

	"github.com/t-hg/stopwatch/curses"
)

func Render(text string) {
	maxWidth := curses.Getmaxx()
	maxHeight := curses.Getmaxy()
	textWidth := 0
	textHeight := 0
	lines := strings.Split(text, "\n")
	textHeight = len(lines)
	for _, line := range lines {
		if len(line) > textWidth {
			textWidth = len(line)
		}
	}
	startX := maxWidth / 2 - textWidth / 2
	startY := maxHeight / 2 - textHeight / 2
	curses.Clear()
	for i, line := range lines {
		curses.Mvaddstr(startY + i, startX, line)
	}
	curses.Refresh()
}
