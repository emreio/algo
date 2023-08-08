package tests

import (
	sortedsquaredarray "algo/sortedSquaredArray"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSortedSquaredArray(t *testing.T) {
	inputArray := []int{-7, -3, 1, 9, 22, 30}
	expected := []int{1, 9, 49, 81, 484, 900}
	if !slices.Equal(expected, sortedsquaredarray.SortedSquaredArray(inputArray)) {
		t.Fail()
	}
}
