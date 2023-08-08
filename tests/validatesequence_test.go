package tests

import (
	validatesequence "algo/ValidateSequence"
	"testing"
)

func TestValidateSequence(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 10, 8}
	subArray := []int{1, 6, -1, 10}

	if !validatesequence.IsValidSubsequence(array, subArray) {
		t.Fail()
	}
}
