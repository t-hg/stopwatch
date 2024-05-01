package tui

import "github.com/t-hg/stopwatch/curses"

func Init() {
	curses.Initscr()
	curses.Noecho()
	curses.Cbreak()
}

func End() {
	curses.Nocbreak()
	curses.Echo()
	curses.Endwin()
}
