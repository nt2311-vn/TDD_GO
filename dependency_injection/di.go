package dependencyinjection

import (
	"bytes"
	"fmt"
)

func Greet(buf *bytes.Buffer, name string) {
	fmt.Fprintf(buf, "Hello, %s", name)
}
