package number

//Input: 3
//Output: false
//Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
//Choosing any x with 0 < x < N and N % x == 0.
//Replacing the number N on the chalkboard with N - x.
//Return True if and only if Alice wins the game, assuming both players play optimally.
//When get one lose the game
func dp_DivisorGame(N int,Alice_round bool)bool{
	for i := 1;i < N;i++{
		if((N % i) == 0){
			var ret = dp_DivisorGame(N - i,!Alice_round)
			if(!ret){
				return true
			}
		}
	}
	return false
}

func DivisorGame(N int) bool {
	return dp_DivisorGame(N,true)
}