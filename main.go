package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/t-hg/stopwatch/stopwatch"
	"github.com/t-hg/stopwatch/style"
	"github.com/t-hg/stopwatch/ui"
)

func main() {
	running := true

	ui := ui.New()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	sigwch := make(chan os.Signal)
	signal.Notify(sigwch, syscall.SIGWINCH)
	go func() {
		for {
			select {
			case <-interrupt:
				running = false
			case <-sigwch:
				ui.End()
				ui.Init()
			}
		}
	}()

	stopwatch := stopwatch.New()
	stopwatch.SetStyleFunc(style.Figletize)

	ui.Init()
	defer ui.End()
	ui.MapKey('s', func() {
		if stopwatch.IsRunning() {
			stopwatch.Stop()
		} else {
			stopwatch.Start()
		}
	})
	ui.MapKey('r', stopwatch.Reset)
	ui.MapKey('q', func() {
		running = false
	})

	ui.Render(`
's' - Start/Stop
'r' - Reset
'q' - Quit
`)

	for running {
		select {
		case text := <-stopwatch.Display():
			ui.Render(text)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
