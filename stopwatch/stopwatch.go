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
	running       bool
	control       chan string
	display       chan string
	timeReference int64
	timeElapsed   int64
}

func New() Stopwatch {
	s := &stopwatch{
		running:       false,
		control:       make(chan string),
		display:       make(chan string),
		timeReference: 0,
		timeElapsed:   0,
	}

	go func() {
		for {
			select {
			case ctrl := <-s.control:
				switch ctrl {
				case "start":
					s.running = true
					s.timeReference = time.Now().UnixNano() - s.timeElapsed
				case "stop":
					s.running = false
				case "reset":
					if !s.running {
						s.timeReference = 0
						s.timeElapsed = 0
					}
				default:
					panic("Unknown control: " + ctrl)
				}
			default:
				if s.running {
					s.timeElapsed = time.Now().UnixNano() - s.timeReference
					hours := (s.timeElapsed / 1000000000 / 60 / 60) % 24
					minutes := (s.timeElapsed / 1000000000 / 60) % 60
					seconds := (s.timeElapsed / 1000000000) % 60
					tenth := (s.timeElapsed / 100000000) % 10
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
