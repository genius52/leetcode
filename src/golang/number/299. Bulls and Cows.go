package number

import "strconv"

//Input: secret = "1807", guess = "7810"
//Output: "1A3B"
//Input: secret = "1123", guess = "0111"
//Output: "1A1B"
func GetHint(secret string, guess string) string {
	var secret_record [10]int
	var guess_record [10]int
	var a_cnt int = 0
	for i,_ := range secret{
		if secret[i] == guess[i]{
			a_cnt++
		}else{
			secret_record[secret[i] - '0']++
			guess_record[guess[i] - '0']++
		}
	}
	var b_cnt int = 0
	for i := 0;i < 10;i++{
		b_cnt += min_int(secret_record[i],guess_record[i])
	}
	res := strconv.Itoa(a_cnt) + "A" + strconv.Itoa(b_cnt) + "B"
	return res
}