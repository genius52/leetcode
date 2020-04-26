package number

//Input: s = "011101"
//Output: 5
//Explanation:
//All possible ways of splitting s into two non-empty substrings are:
//left = "0" and right = "11101", score = 1 + 4 = 5
//left = "01" and right = "1101", score = 1 + 3 = 4
//left = "011" and right = "101", score = 1 + 2 = 3
//left = "0111" and right = "01", score = 1 + 1 = 2
//left = "01110" and right = "1", score = 2 + 1 = 3
func MaxScore(s string) int {
	l := len(s)
	var dp_zero []int = make([]int,l)
	var dp_one []int = make([]int,l)
	var zero_cnt int = 0
	var one_cnt int = 0
	for i, c := range s{
		if c == '0'{
			zero_cnt++
		}else{
			one_cnt++
		}
		dp_zero[i] = zero_cnt
		dp_one[i] = one_cnt
	}
	var max int = 0
	for i := 0;i < l - 1;i++{
		tmp := dp_zero[i] + (one_cnt - dp_one[i])
		if tmp > max{
			max = tmp
		}
	}
	return max
}