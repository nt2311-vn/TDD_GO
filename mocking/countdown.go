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

func CountDown(w io.Writer) {
	for i := countStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(time.Second * 1)
	}

	fmt.Fprint(w, finalWord)
}
