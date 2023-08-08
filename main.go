package main

import (
	sortedsquaredarray "algo/sortedsquaredarray"
	"algo/tournamentwinner"
	"fmt"
)

func main() {
	inputArray := []int{1, 2, 3, 5, 6, 8, 9}
	sortedsquaredarray.SortedSquaredArray(inputArray)

	competitions := [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
	fmt.Println(tournamentwinner.TournamentWinner(competitions, []int{0, 0, 1}))
}
