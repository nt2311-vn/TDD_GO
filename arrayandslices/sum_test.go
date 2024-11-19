package arrayandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectResult(t, numbers, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 50}

		got := Sum(numbers)
		want := 56

		assertCorrectResult(t, numbers, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum all list of collection of two numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sum all list of collection of four numbers", func(t *testing.T) {
		got := SumAll([]int{0, 3, 2, 4}, []int{1, 1, 1, 2}, []int{6, 7, 5, 4})
		want := []int{9, 5, 22}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sum all with a collection is empty slice", func(t *testing.T) {
		got := SumAll([]int{}, []int{6, 9})
		want := []int{0, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum all tails of collection two numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sum all tails of collection four numbers", func(t *testing.T) {
		got := SumAllTails([]int{0, 3, 2, 4}, []int{1, 1, 1, 2}, []int{6, 7, 5, 4})
		want := []int{9, 4, 16}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sum all tails with a collection is empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{6, 9})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertCorrectResult(t testing.TB, numbers []int, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got '%d' but want '%d' input numbers %v", numbers, got, want)
	}
}
