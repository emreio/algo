package tournamentwinner

const HOME_TEAM_WON = 1

/*
TEST INPUT

// competitions := [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
*/
func TournamentWinner(competitions [][]string, results []int) string {

	maxPoint := 0
	winnerTeam := ""

	teams := make(map[string]int)

	for i, j := range competitions {
		result := results[i]
		var winnerIndex int

		if result == HOME_TEAM_WON {
			winnerIndex = 0
		} else {
			winnerIndex = 1
		}

		teams[j[winnerIndex]] += 3
		if teams[j[winnerIndex]] > maxPoint {
			maxPoint = teams[j[winnerIndex]]
			winnerTeam = j[winnerIndex]
		}
	}
	return winnerTeam
}
