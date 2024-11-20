package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		want := ErrNotFound
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new key", func(t *testing.T) {
		k, v := "test", "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Add(k, v)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, k, v)
	})

	t.Run("Add exist key", func(t *testing.T) {
		k, v := "test", "this is just a test"
		dictionary := Dictionary{k: v}

		err := dictionary.Add(k, v)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, k, v)
	})
}

func TestUpdate(t *testing.T) {
	k, v := "test", "this is just a test"
	dict := Dictionary{k: v}
	newVal := "new definition"

	dict.Update(k, newVal)
	assertDefinition(t, dict, k, newVal)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should fin added word:", err)
	}

	assertStrings(t, got, value)
}
