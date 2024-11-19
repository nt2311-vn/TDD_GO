package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("first case random number", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectResult(t, sum, expected)
	})

	t.Run("second case random number", func(t *testing.T) {
		sum := Add(1, 5)
		expected := 6

		assertCorrectResult(t, sum, expected)
	})
}

func assertCorrectResult(t testing.TB, sum, expected int) {
	t.Helper()

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
