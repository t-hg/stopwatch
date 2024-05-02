package ui

import (
	"strings"

	"github.com/t-hg/stopwatch/curses"
)

var memoizedText string

func (*ui) Render(text string) {
	if memoizedText == text {
		return
	}
	memoizedText = text
	maxWidth := curses.GetMaxX()
	maxHeight := curses.GetMaxY()
	textWidth := 0
	textHeight := 0
	lines := strings.Split(text, "\n")
	textHeight = len(lines)
	for _, line := range lines {
		if len(line) > textWidth {
			textWidth = len(line)
		}
	}
	startX := maxWidth/2 - textWidth/2
	startY := maxHeight/2 - textHeight/2
	curses.Clear()
	for i, line := range lines {
		curses.MvAddStr(startY+i, startX, line)
	}
	curses.Refresh()
}
