package main

import (
	"fmt"
	"io"
	"os"
)

func CountDown(w io.Writer) {
	fmt.Fprint(w, "3")
}

func main() {
	CountDown(os.Stdout)
}
