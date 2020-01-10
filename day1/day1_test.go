package main

import (
	"reflect"
	"runtime"
	"testing"
)

type functionToTest func(...int) int

func TestDay1(t *testing.T) {
	executeAndCheck(Day1, t, 2, 12)
	executeAndCheck(Day1, t, 2, 14)
	executeAndCheck(Day1, t, 654, 1969)
	executeAndCheck(Day1, t, 33583, 100756)
	executeAndCheck(Day1, t, 2+2+654+33583, 12, 14, 1969, 100756)
}

func executeAndCheck(foo functionToTest, t *testing.T, want int, input ...int) {
	got := foo(input...)

	funcName := runtime.FuncForPC(reflect.ValueOf(foo).Pointer()).Name()
	if got != want {
		t.Errorf("%v == %d, want %d", funcName, got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	executeAndCheck(Day1Part2, t, 0, 0)
	executeAndCheck(Day1Part2, t, 2, 12)
	executeAndCheck(Day1Part2, t, 2, 14)
	executeAndCheck(Day1Part2, t, 966, 1969)
	executeAndCheck(Day1Part2, t, 50346, 100756)
	executeAndCheck(Day1Part2, t, 2+2+966+50346, 12, 14, 1969, 100756)
}
