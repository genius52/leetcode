package array

//Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
//Output: 16
//Explanation: The two words can be "abcw", "xtfn".
func MaxProduct(words []string) int {
	var l int = len(words)
	var record map[int]int = make(map[int]int)
	var max_len int = 0
	for i := 0;i < l;i++{
		var val int = 0
		for _,c := range words[i]{
			val |= 1 << (c - 'a')
		}
		for k,v := range record{
			if k & val == 0{
				cur_len := v * len(words[i])
				if cur_len > max_len{
					max_len = cur_len
				}
			}
		}
		if last_len,ok := record[val];ok{
			if last_len < len(words[i]){
				record[val] = len(words[i])
			}
		}else{
			record[val] = len(words[i])
		}
	}
	return max_len
}