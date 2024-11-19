package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("simple repeat character a", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		assertCorrectStr(t, repeated, expected)
	})
}

func assertCorrectStr(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
