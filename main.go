package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/t-hg/stopwatch/stopwatch"
	"github.com/t-hg/stopwatch/tui"
)

func main() {
	running := true

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
				tui.End()
				tui.Init()
			}
		}
	}()

	tui.Init()
	defer tui.End()

	sw := stopwatch.New()
	sw.Start()

	for running {
		select {
		case text := <-sw.Display():
			tui.Render(text)
		}
	}
}
