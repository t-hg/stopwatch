package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/t-hg/stopwatch/curses"
)

const startup_message = 
	"'s' - start/stop\n" + 
	"'r' - reset\n" +
	"'q' - quit"

func initialize() {
	curses.InitScr()
	curses.Cbreak()
	curses.NoEcho()
	curses.CursSet(0)
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
	maxLineLen := 0
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if len(line) > maxLineLen {
			maxLineLen = len(line)
		}
	}
	maxY := curses.GetMaxY()
	maxX := curses.GetMaxX()
	y := maxY / 2 - len(lines) / 2
	x := maxX / 2 - maxLineLen / 2
	for idx, line := range lines {
		curses.MvAddStr(y+idx, x, line)	
	}
}

func main() {
	go handleSIGINT()
	go handleSIGWINCH()
	initialize()
	print(startup_message)
	curses.GetCh()
	cleanup()
}
