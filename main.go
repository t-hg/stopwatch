package main

import (
	"os"
	"os/signal"

	"github.com/t-hg/stopwatch/stopwatch"
	"github.com/t-hg/stopwatch/tui"
)

func main() {
	running := true

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		select {
		case <- interrupt:
			running = false
		}
	}()

	tui.Init()
	defer tui.End()

	sw := stopwatch.New()
	sw.Start()

	for running {
		select {
		case text := <- sw.Display():
			tui.Render(text)
		}
	}
}
