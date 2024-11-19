package arrayandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("simple array into sum", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectResult(t, numbers, got, want)
	})
}

func assertCorrectResult(t testing.TB, numbers [5]int, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got '%d' but want '%d' input numbers %v", numbers, got, want)
	}
}
