package structmethodsinterface

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	assertCorrectResult(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	assertCorrectResult(t, got, want)
}

func assertCorrectResult(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
