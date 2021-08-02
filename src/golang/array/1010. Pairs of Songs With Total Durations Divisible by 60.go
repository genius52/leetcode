package array

//[30,20,150,100,40]
func numPairsDivisibleBy60(time []int) int {
	res := 0
	dict := make(map[int]int)
	l := len(time)
	for i := 0;i < l;i++{
		abs_val := time[i]%60
		if cnt,ok := dict[abs_val];ok && cnt > 0{
			res += dict[abs_val]
		}
		needed := (600 - time[i])%60
		if cnt,ok := dict[needed];ok && cnt > 0{
			dict[needed]++
		}else{
			dict[needed] = 1
		}
	}
	return res
}