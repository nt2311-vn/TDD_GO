package structmethodsinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assertCorrectResult(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		assertCorrectResult(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Area()
		want := 314.1592653589793

		assertCorrectResult(t, got, want)
	})
}

func assertCorrectResult(t testing.TB, got, want float64) {
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
