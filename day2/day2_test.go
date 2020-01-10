package main

import (
	"reflect"
	"testing"
)

func TestDay2(t *testing.T) {
	compareAndAssert := func(t *testing.T, got, want, given []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v, given %v", got, want, given)
		}
	}

	copyGiven := func(given []int) []int {
		givenCopy := make([]int, len(given))
		copy(givenCopy, given)
		return givenCopy
	}

	t.Run("Empty program", func(t *testing.T) {
		given := []int{}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{}
		compareAndAssert(t, got, want, given)
	})

	t.Run("terminator only", func(t *testing.T) {
		given := []int{99}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{99}
		compareAndAssert(t, got, want, given)
	})

	t.Run("simple addiction", func(t *testing.T) {
		given := []int{1, 0, 0, 3, 99}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{1, 0, 0, 2, 99}
		compareAndAssert(t, got, want, given)
	})

	t.Run("simple multiplication", func(t *testing.T) {
		given := []int{2, 0, 0, 3, 99}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{2, 0, 0, 4, 99}
		compareAndAssert(t, got, want, given)
	})

	t.Run("two instructions", func(t *testing.T) {
		given := []int{1, 0, 0, 3, 2, 4, 4, 7, 99}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{1, 0, 0, 2, 2, 4, 4, 4, 99}
		compareAndAssert(t, got, want, given)
	})

	t.Run("half after a terminator", func(t *testing.T) {
		given := []int{1, 0, 0, 3, 2, 4, 4, 7, 99, 0, 0, 0, 1, 8, 8, 2}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{1, 0, 0, 2, 2, 4, 4, 4, 99, 0, 0, 0, 1, 8, 8, 2}
		compareAndAssert(t, got, want, given)
	})

	t.Run("first example", func(t *testing.T) {
		given := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
		compareAndAssert(t, got, want, given)
	})

	t.Run("second example", func(t *testing.T) {
		given := []int{1, 0, 0, 0, 99}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{2, 0, 0, 0, 99}
		compareAndAssert(t, got, want, given)
	})

	t.Run("third example", func(t *testing.T) {
		given := []int{2, 3, 0, 3, 99}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{2, 3, 0, 6, 99}
		compareAndAssert(t, got, want, given)
	})

	t.Run("fourth example", func(t *testing.T) {
		given := []int{2, 4, 4, 5, 99, 0}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{2, 4, 4, 5, 99, 9801}
		compareAndAssert(t, got, want, given)
	})

	t.Run("fifth example", func(t *testing.T) {
		given := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
		input := copyGiven(given)
		got := Day2(input)
		want := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
		compareAndAssert(t, got, want, given)
	})

}
