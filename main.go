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

func render(text string) {
	// find y, x so that given
	// text is centered
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

	// print lines respectively
	curses.Clear()
	for idx, line := range lines {
		curses.MvAddStr(y+idx, x, line)
	}
	curses.Refresh()
}

func main() {
	// setup
	curses.InitScr()
	curses.Cbreak()
	curses.NoEcho()
	curses.CursSet(0)
	curses.NoDelay(true)

	// signal handlers
	sigint := make(chan os.Signal, 1)
	sigwinch := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigwinch, syscall.SIGWINCH)

	// text to be displayed
	text := startupMessage

	// loop variables
	var running bool
	var start int64
	var elapsed int64

loop:
	for {
		// handle signals
		select {
		case <-sigint:
			break loop
		case <-sigwinch:
			curses.EndWin()
			curses.Refresh()
		default:
		}

		// handle character input
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
			text = "0.0"
		case 'q':
			break loop
		}

		// update watch
		if running {
			now := time.Now().UnixMilli()
			elapsed = now - start
			hours := (elapsed / 1000 / 60 / 60) % 24
			minutes := (elapsed / 1000 / 60) % 60
			seconds := (elapsed / 1000) % 60
			tenth := (elapsed / 100) % 10
			if hours > 0 && hours >= 10 {
				text = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
			} else if hours > 0 && hours < 10 {
				text = fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
			} else if minutes > 0 && minutes >= 10 {
				text = fmt.Sprintf("%02d:%02d", minutes, seconds)
			} else if minutes > 0 && minutes < 10 {
				text = fmt.Sprintf("%d:%02d", minutes, seconds)
			}  else {
				text = fmt.Sprintf("%d", seconds)
			}
			text = fmt.Sprintf("%s.%d", text, tenth)
		}

		// display text
		// TODO: styling
		render(text)

		// little time interval
		// to avoid busy wait
		time.Sleep(50 * time.Millisecond)
	}

	// cleanup
	curses.CursSet(2)
	curses.Echo()
	curses.NoCbreak()
	curses.EndWin()
}
