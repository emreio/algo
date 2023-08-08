package tests

import (
	"algo/tournamentwinner"
	"testing"
)

func TestTournamentWinner(t *testing.T) {
	competitions := [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
	expected := "Python"
	if expected != tournamentwinner.TournamentWinner(competitions, []int{0, 0, 1}) {
		t.Fail()
	}
}
