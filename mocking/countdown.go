package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countStart = 3
	finalWord  = "Go!"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func CountDown(w io.Writer, s *SpySleeper) {
	for i := countStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(time.Second * 1)
		s.Sleep()
	}

	fmt.Fprint(w, finalWord)
}
