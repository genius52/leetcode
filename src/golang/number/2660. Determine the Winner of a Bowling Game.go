package number

func isWinner(player1 []int, player2 []int) int {
	var l int = len(player1)
	var ten1 int = 0
	var ten2 int = 0
	var score1 int = 0
	var score2 int = 0
	for i := 0; i < l; i++ {
		if ten1 > 0 {
			score1 += 2 * player1[i]
			ten1--
		} else {
			score1 += player1[i]
		}
		if ten2 > 0 {
			score2 += 2 * player2[i]
			ten2--
		} else {
			score2 += player2[i]
		}
		if player1[i] == 10 {
			ten1 = 2
		}
		if player2[i] == 10 {
			ten2 = 2
		}
	}
	if score1 == score2 {
		return 0
	} else if score1 > score2 {
		return 1
	} else {
		return 2
	}
}
