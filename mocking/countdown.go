package mocking

import (
	"fmt"
	"io"
)

func CountDown(w io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
	}

	fmt.Fprint(w, "Go!")
}
