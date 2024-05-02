package ui

import "github.com/t-hg/stopwatch/curses"

type UI interface {
	Init()
	End()
	Render(string)
	MapKey(rune, func())
}

type ui struct {
	keyMappings map[rune]func()
}

func New() UI {
	return &ui{
		keyMappings: make(map[rune]func()),
	}
}

func (ui *ui) Init() {
	curses.InitScr()
	curses.NoEcho()
	curses.Cbreak()
	curses.CursSet(0)

	go func() {
		for {
			char := curses.GetCh()
			if fn, ok := ui.keyMappings[char]; ok {
				fn()
			}
		}
	}()
}

func (*ui) End() {
	curses.CursSet(2)
	curses.NoCbreak()
	curses.Echo()
	curses.EndWin()
}

func (ui *ui) MapKey(key rune, fn func()) {
	ui.keyMappings[key] = fn
}
