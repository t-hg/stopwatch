package stopwatch

import (
	"fmt"
	"time"
)

type Stopwatch interface {
	Start()
	Stop()
	Reset()
	Display() chan string
}

type stopwatch struct {
	running   bool
	control   chan string
	display   chan string
	millisecs int
}

func New() Stopwatch {
	s := &stopwatch{
		running:   false,
		control:   make(chan string),
		display:   make(chan string),
		millisecs: 0,
	}

	go func() {
		for {
			select {
			case ctrl := <-s.control:
				switch ctrl {
				case "start":
					s.running = true
				case "stop":
					s.running = false
				case "reset":
					if !s.running {
						s.millisecs = 0
					}
				default:
					// ignore
				}
			default:
				if s.running {
					s.millisecs += 1
					hours := (s.millisecs / 1000 / 60 / 60) % 24
					minutes := (s.millisecs / 1000 / 60) % 60
					seconds := (s.millisecs / 1000) % 60
					tenth := (s.millisecs / 100) % 10
					figletized := ""
					if hours > 0 && hours >= 10 {
						figletized = Figletize(fmt.Sprintf("%02d : %02d : %02d", hours, minutes, seconds))
					} else if hours > 0 && hours < 10 {
						figletized = Figletize(fmt.Sprintf("%d : %02d : %02d", hours, minutes, seconds))
					} else if minutes > 0 && minutes >= 10 {
						figletized = Figletize(fmt.Sprintf("%02d : %02d", minutes, seconds))
					} else if minutes > 0 && minutes < 10 {
						figletized = Figletize(fmt.Sprintf("%d : %02d", minutes, seconds))
					} else if seconds > 10 {
						figletized = Figletize(fmt.Sprintf("%02d", seconds))
					} else {
						figletized = Figletize(fmt.Sprintf("%d", seconds))
					}
					figletized += fmt.Sprintf(".%d", tenth)
					s.display <- figletized
				}
			}
			time.Sleep(time.Millisecond)
		}
	}()

	return s
}

func (s *stopwatch) Start() {
	s.control <- "start"
}

func (s *stopwatch) Stop() {
	s.control <- "stop"
}

func (s *stopwatch) Reset() {
	s.control <- "reset"
}

func (s *stopwatch) Display() chan string {
	return s.display
}
