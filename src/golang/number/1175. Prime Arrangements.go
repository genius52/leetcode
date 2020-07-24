package number


func NumPrimeArrangements(n int) int {
	var LIMIT int = 1e9 + 7
	var prime_numbers []int = []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97}
	var prime_cnt int = 0
	var pos int = 0
	for(pos <= 24  && prime_numbers[pos] <= n){
		pos++
		prime_cnt++
	}
	var res int = 1
	for i := 1;i <= prime_cnt;i++{
		res = res * i % LIMIT
	}
	for i := 1;i <= (n - prime_cnt);i++{
		res = res * i % LIMIT
	}
	return res
}