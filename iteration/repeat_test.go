package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("simple repeat character a", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"

		assertCorrectStr(t, repeated, expected)
	})

	t.Run("repeat a phrase", func(t *testing.T) {
		repeated := Repeat("yahoo", 5)
		expected := "yahooyahooyahooyahooyahoo"

		assertCorrectStr(t, repeated, expected)
	})
}

func assertCorrectStr(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}
