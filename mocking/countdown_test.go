package mocking

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := bytes.Buffer{}

	CoutDown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
