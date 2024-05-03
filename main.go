package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/t-hg/stopwatch/curses"
)

const startupMessage = `'s' - start/stop
'r' - reset
'q' - quit`

func initialize() {
	curses.InitScr()
	curses.Cbreak()
	curses.NoEcho()
	curses.CursSet(0)
	curses.NoDelay(true)
}

func cleanup() {
	curses.CursSet(2)
	curses.Echo()
	curses.NoCbreak()
	curses.EndWin()
}

func print(text string) {
	curses.Clear()
	maxLineLen := 0
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if len(line) > maxLineLen {
			maxLineLen = len(line)
		}
	}
	maxY := curses.GetMaxY()
	maxX := curses.GetMaxX()
	y := maxY/2 - len(lines)/2
	x := maxX/2 - maxLineLen/2
	for idx, line := range lines {
		curses.MvAddStr(y+idx, x, line)
	}
	curses.Refresh()
}

func main() {
	initialize()
	sigint := make(chan os.Signal, 1)
	sigwinch := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigwinch, syscall.SIGWINCH)
	text := startupMessage
	var running bool
	var start int64
	var elapsed int64
loop:
	for {
		select {
		case <-sigint:
			break loop
		case <-sigwinch:
			curses.EndWin()
			curses.Refresh()
		default:
		}
		switch curses.GetCh() {
		case 's':
			if running {
				running = false
			} else {
				running = true
				start = time.Now().UnixMilli() - elapsed
			}
		case 'r':
			if running {
				start = time.Now().UnixMilli()
			} else {
				start = 0
			}
			elapsed = 0
		case 'q':
			break loop
		}
		if running {
			now := time.Now().UnixMilli()
			elapsed = now - start
			text = fmt.Sprintf("%d", elapsed)
		}
		print(text)
		time.Sleep(50 * time.Millisecond)
	}
	cleanup()
}
