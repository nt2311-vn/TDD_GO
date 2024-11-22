package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{{33, "London"}, {34, "Reykjavík"}},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{{33, "London"}, {34, "Reykjavík"}},
			[]string{"London", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{"Cow": "Moo", "Sheep": "Baa"},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})

		t.Run("with maps", func(t *testing.T) {
			aMap := map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			}
			var got []string

			walk(aMap, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, "Moo")
			assertContains(t, got, "Baa")
		})
	}
}

func assertContains(t testing.TB, haystacks []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystacks {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but did not", haystacks, needle)
	}
}
