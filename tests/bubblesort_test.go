package tests

import (
	"algo/bubblesort"
	"testing"

	"golang.org/x/exp/slices"
)

func TestBubbleSort(t *testing.T) {
	_input := []int{6, 4, 5, 1, 1, 8, 9, 3, -1, -2, 3, -5, 0}
	expected := []int{-5, -2, -1, 0, 1, 1, 3, 3, 4, 5, 6, 8, 9}

	if !slices.Equal(expected, bubblesort.BubbleSortArray(_input)) {
		t.Fail()
	}
}
