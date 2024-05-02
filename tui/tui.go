package tui

import "github.com/t-hg/stopwatch/curses"

func Init() {
	curses.InitScr()
	curses.NoEcho()
	curses.Cbreak()
	curses.CursSet(0)
}

func End() {
	curses.CursSet(2)
	curses.NoCbreak()
	curses.Echo()
	curses.EndWin()
}
