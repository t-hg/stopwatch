package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/t-hg/stopwatch/curses"
)

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

func main() {
	go handleSIGINT()
	go handleSIGWINCH()
	initialize()
	curses.GetCh()
	cleanup()
}
