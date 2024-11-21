package mocking

import (
	"fmt"
	"io"
)

func CountDown(w io.Writer) {
	fmt.Fprintf(w, "%d", 3)
}
