package main

import (
	"fmt"

	"github.com/t-hg/stopwatch/stopwatch"
)

func main() {
	s := stopwatch.New()
	s.Start()
		
	for {
		select {
		case d := <- s.Display():
			fmt.Println(d)
		}
	}
}
