package curses

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lcurses
// #include <curses.h>
import "C"

func Initscr() {
	C.initscr()
}

func Endwin() {
	C.endwin()
}

func Noecho() {
	C.noecho()
}

func Echo() {
	C.echo()
}

func Cbreak() {
	C.cbreak()
}

func Nocbreak() {
	C.nocbreak()
}

func Getmaxx() int {
	return int(C.getmaxx(C.stdscr))
}

func Getmaxy() int {
	return int(C.getmaxy(C.stdscr))
}

func Clear() {
	C.clear()
}

func Mvaddstr(y int, x int, str string) {
	C.mvaddstr(C.int(y), C.int(x), C.CString(str))
}

func Refresh() {
	C.refresh()
}
