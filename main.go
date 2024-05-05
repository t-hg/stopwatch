package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/t-hg/stopwatch/curses"
	"github.com/t-hg/stopwatch/style"
)

func render(text string) {
	// find y, x so that given
	// text is centered
	maxLineLen := 0
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		runes := []rune(line)
		if len(runes) > maxLineLen {
			maxLineLen = len(runes)
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
	// flags
	flagStyle := flag.Int("style", 1, "different styles (1, 2 or 3)")
	flag.Parse()

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
	text := `
<space> - start/stop
    <r> - reset
    <q> - quit
`
	// styling to be used
	var charset []string
	switch *flagStyle {
	case 1:
		charset = style.Charset1
	case 2:
		charset = style.Charset2
	case 3:
		charset = style.Charset3
	}

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
		case '\x20':
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
			text = style.Apply("0", charset)
			text = fmt.Sprintf("%s.%d", text, 0)
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
			} else {
				text = fmt.Sprintf("%d", seconds)
			}
			text = style.Apply(text, charset)
			text = fmt.Sprintf("%s.%d", text, tenth)
		}

		// display text
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
