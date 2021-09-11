package arraysandslices

import (
	"reflect"
	"testing"
)

func TestArrayAndSlices(t *testing.T) {

	t.Run("should sum all 5 items inside the array", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers[:])
		want := 15

		t.Logf("Numbers: %v", numbers)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("should sum all items inside the slice", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		t.Logf("Numbers: %v", numbers)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should sum all elements inside every array/slice passed", func(t *testing.T) {

		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAll(numbers1, numbers2)
		want := []int{3, 9}

		t.Logf("The numbers #1: %v", numbers1)
		t.Logf("The numbers #2: %v", numbers2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("should sum all tails from the arrays given", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
