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

const startup_message = "'s' - start/stop\n" +
	"'r' - reset\n" +
	"'q' - quit"

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

func handleSIGINT() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT)
	<-signals
	cleanup()
	os.Exit(0)
}

func handleSIGWINCH() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGWINCH)
	for {
		<-signals
		cleanup()
		initialize()
		curses.Refresh()
	}
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
	go handleSIGINT()
	go handleSIGWINCH()
	initialize()
	print(startup_message)
	var running bool
	var start int64
	var elapsed int64
loop:
	for {
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
			print(fmt.Sprintf("%d", elapsed))
		}
		time.Sleep(50 * time.Millisecond)
	}
	cleanup()
}
