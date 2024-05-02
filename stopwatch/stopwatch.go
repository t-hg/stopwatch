package stopwatch

import (
	"fmt"
	"time"
)

type StyleFunc func(string) string

type Stopwatch interface {
	Start()
	Stop()
	Reset()
	Display() chan string
	SetStyleFunc(fn StyleFunc)
	IsRunning() bool
}

var defaultStyleFunc = func(text string) string {
	return text
}

type stopwatch struct {
	running       bool
	control       chan string
	display       chan string
	timeReference int64
	timeElapsed   int64
	styleFunc     StyleFunc
}

func New() Stopwatch {
	s := &stopwatch{
		running:       false,
		control:       make(chan string),
		display:       make(chan string),
		timeReference: 0,
		timeElapsed:   0,
		styleFunc:     defaultStyleFunc,
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
					s.timeElapsed = 0
					s.timeReference = time.Now().UnixNano() - s.timeElapsed
					if !s.running {
						s.display <- s.styleFunc("0") + ".0"
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
					stylized := ""
					if hours > 0 && hours >= 10 {
						stylized = s.styleFunc(fmt.Sprintf("%02d : %02d : %02d", hours, minutes, seconds))
					} else if hours > 0 && hours < 10 {
						stylized = s.styleFunc(fmt.Sprintf("%d : %02d : %02d", hours, minutes, seconds))
					} else if minutes > 0 && minutes >= 10 {
						stylized = s.styleFunc(fmt.Sprintf("%02d : %02d", minutes, seconds))
					} else if minutes > 0 && minutes < 10 {
						stylized = s.styleFunc(fmt.Sprintf("%d : %02d", minutes, seconds))
					} else if seconds > 10 {
						stylized = s.styleFunc(fmt.Sprintf("%02d", seconds))
					} else {
						stylized = s.styleFunc(fmt.Sprintf("%d", seconds))
					}
					stylized += fmt.Sprintf(".%d", tenth)
					s.display <- stylized
				}
				time.Sleep(time.Millisecond)
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

func (s *stopwatch) SetStyleFunc(fn StyleFunc) {
	s.styleFunc = fn
}

func (s *stopwatch) IsRunning() bool {
	return s.running
}
